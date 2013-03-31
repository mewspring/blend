// Package blend implements parsing of Blender files.
package blend

import (
	"encoding/binary"
	"fmt"
	"io"
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

// Parse parsers the provided blend file.
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

		b.Blocks = append(b.Blocks, blk)
	}

	return b, nil
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

	return hdr, nil
}
