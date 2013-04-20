// Package blend implements parsing of Blender files.
package blend

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/mewmew/blend/block"
)

// Blend represents the information contained within a blend file. It contains a
// file header and a slice of file blocks.
type Blend struct {
	Hdr    *Header
	Blocks []*block.Block
	// io.Closer of the underlying file.
	io.Closer
}

// Parse parsers the provided blend file. An *io.SectionReader is stored in each
// individual block body. By default, blocks must be parsed manually using the
// ParseBody method. Use ParseAll for a complete parsing of the blend file,
// including all block bodies.
//
// Caller should close b when done reading from it.
func Parse(filePath string) (b *Blend, err error) {
	b = new(Blend)
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	b.Closer = f

	// Parse file header.
	b.Hdr, err = ParseHeader(f)
	if err != nil {
		return nil, err
	}

	// Parse file blocks.
	for {
		blk, err := block.Parse(f, b.Hdr.Order, b.Hdr.PtrSize)
		if err != nil {
			return nil, err
		}
		if blk.Hdr.Code == block.CodeENDB {
			break
		}
		_, ok := block.Addr[blk.Hdr.OldAddr]
		if ok {
			return nil, fmt.Errorf("blend.Parse: multiple occurances of struct address %p.", blk.Hdr.OldAddr)
		}
		block.Addr[blk.Hdr.OldAddr] = blk

		b.Blocks = append(b.Blocks, blk)
	}

	return b, nil
}

// ParseAll parses the blend file and all block bodies.
//
// ParseAll closes b when done reading from it.
func ParseAll(filePath string) (b *Blend, err error) {
	// Parse file header and block headers.
	b, err = Parse(filePath)
	if err != nil {
		return nil, err
	}
	defer b.Close()

	// Parse DNA block.
	dna, err := b.GetDNA()
	if err != nil {
		return nil, err
	}

	// Parse block bodies.
	for _, blk := range b.Blocks {
		blk.ParseBody(b.Hdr.Order, dna)
	}

	return b, nil
}

// GetDNA locates, parses and returns the DNA block.
func (b *Blend) GetDNA() (dna *block.DNA, err error) {
	for _, blk := range b.Blocks {
		dna, ok := blk.Body.(*block.DNA)
		if ok {
			// DNA block already parsed.
			return dna, nil
		}
		if blk.Hdr.Code == block.CodeDNA1 {
			// Parse the DNA block body and store it in blk.Body.
			r, ok := blk.Body.(io.Reader)
			if !ok {
				return nil, errors.New("Blend.GetDNA: unable to locate DNA block body reader.")
			}
			dna, err = block.ParseDNA(r, b.Hdr.Order)
			if err != nil {
				return nil, err
			}
			blk.Body = dna
			return dna, nil
		}
	}
	return nil, errors.New("Blend.GetDNA: unable to locate DNA block.")
}

// A Header is present at the beginning of each blend file.
type Header struct {
	// Pointer size.
	PtrSize int
	// Byte order.
	Order binary.ByteOrder
	// Blender version.
	Ver int
}

// ParseHeader parses and returns the header of a blend file.
//
// Example file header:
//    "BLENDER_V100"
//    //  0-6   magic ("BLENDER")
//    //    7   pointer size ("_" or "-")
//    //    8   endianness ("V" or "v")
//    // 9-11   version ("100")
func ParseHeader(r io.Reader) (hdr *Header, err error) {
	buf := make([]byte, 12)
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}

	// File identifier.
	magic := string(buf[0:7])
	if magic != "BLENDER" {
		return nil, fmt.Errorf("blend.ParseHeader: invalid file identifier %q.", magic)
	}

	// Pointer size.
	hdr = new(Header)
	size := buf[7]
	switch size {
	case '_':
		// _ = 4 byte pointer
		hdr.PtrSize = 4
	case '-':
		// - = 8 byte pointer
		hdr.PtrSize = 8
	default:
		return nil, fmt.Errorf("blend.ParseHeader: invalid pointer size character '%c'.", size)
	}

	// Byte order.
	order := buf[8]
	switch order {
	case 'v':
		// v = little endian
		hdr.Order = binary.LittleEndian
	case 'V':
		// V = big endian
		hdr.Order = binary.BigEndian
	default:
		return nil, fmt.Errorf("blend.ParseHeader: invalid byte order character '%c'.", order)
	}

	// Version.
	hdr.Ver, err = strconv.Atoi(string(buf[9:12]))
	if err != nil {
		return nil, fmt.Errorf("blend.ParseHeader: invalid version; %s", err)
	}
	if WarnVersion {
		if hdr.Ver != block.BlenderVer {
			log.Printf("Warning: Version mismatch between file (v%d) and block package (v%d).\n", hdr.Ver, block.BlenderVer)
			if hdr.Ver < block.BlenderVer {
				log.Println("The file's Blender version is too old. Use Blender [1] to open and resave the file.")
				log.Println("[1]: http://www.blender.org/")
			} else {
				log.Println("The block package Blender version is too old. Use blendef [1] to regenerate the block package.")
				log.Println("[1]: go get github.com/mewmew/blend/cmd/blendef")
			}
		}
	}

	return hdr, nil
}

// When WarnVersion is set to true, ParseHeader will report warnings if the
// Blender version of the blend file and the block package differs.
var WarnVersion = true
