package block

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// A Block contains a header and a type dependent body.
type Block struct {
	Hdr *Header
	// Body contains an *io.SectionReader which is later replaced with a concrete
	// block type after it has been parsed.
	Body interface{}
}

// Parse parses and returns a file block.
func Parse(f *os.File, order binary.ByteOrder, ptrSize int) (blk *Block, err error) {
	// Parse block header.
	blk = new(Block)
	blk.Hdr, err = ParseHeader(f, order, ptrSize)
	if err != nil {
		return nil, err
	}

	// Store section reader for block body.
	off, err := f.Seek(blk.Hdr.Size, os.SEEK_CUR)
	if err != nil {
		return nil, err
	}
	blk.Body = io.NewSectionReader(f, off-blk.Hdr.Size, blk.Hdr.Size)

	return blk, nil
}

// Header contains information about the block's type and size.
type Header struct {
	// Code provides a rough type description of the block.
	Code BlockCode
	// Total length of the data after the block header.
	Size int64
	// Memory address of the structure when it was written to disk.
	OldAddr uint64
	// Index in the Structure DNA.
	SDNAIndex int
	// Number of structures located in this block.
	Count int
}

// ParseHeader parses and returns a file block header.
//
// Example file block header:
//    44 41 54 41  E0 00 00 00  88 5E 9D 04  00 00 00 00    DATA.....^......
//    F8 00 00 00  0E 00 00 00                              ........

//    //   0-3   block code   ("DATA")
//    //   4-7   size         (0x000000E0 = 224)
//    //  8-15   old addr     (0x00000000049D5E88) // size depends on PtrSize.
//    // 16-19   sdna index   (0x000000F8 = 248)
//    // 20-23   count        (0x0000000E = 14)
func ParseHeader(r io.Reader, order binary.ByteOrder, ptrSize int) (hdr *Header, err error) {
	// Block code.
	buf := make([]byte, 4)
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}
	hdr = new(Header)
	code := string(buf)
	switch code {
	case "BR\x00\x00":
		hdr.Code = CodeBR
	case "CA\x00\x00":
		hdr.Code = CodeCA
	case "DATA":
		hdr.Code = CodeDATA
	case "DNA1":
		hdr.Code = CodeDNA1
	case "ENDB":
		hdr.Code = CodeENDB
	case "GLOB":
		hdr.Code = CodeGLOB
	case "LA\x00\x00":
		hdr.Code = CodeLA
	case "MA\x00\x00":
		hdr.Code = CodeMA
	case "ME\x00\x00":
		hdr.Code = CodeME
	case "OB\x00\x00":
		hdr.Code = CodeOB
	case "REND":
		hdr.Code = CodeREND
	case "SC\x00\x00":
		hdr.Code = CodeSC
	case "SN\x00\x00":
		hdr.Code = CodeSN
	case "TE\x00\x00":
		hdr.Code = CodeTE
	case "TEST":
		hdr.Code = CodeTEST
	case "WM\x00\x00":
		hdr.Code = CodeWM
	case "WO\x00\x00":
		hdr.Code = CodeWO
	default:
		return nil, fmt.Errorf("Header.ParseHeader: block code %q not yet implemented.", code)
	}

	// Block size.
	var x int32
	err = binary.Read(r, order, &x)
	if err != nil {
		return nil, err
	}
	hdr.Size = int64(x)

	// Old memory address.
	switch ptrSize {
	case 4:
		var x uint32
		err = binary.Read(r, order, &x)
		if err != nil {
			return nil, err
		}
		hdr.OldAddr = uint64(x)
	case 8:
		var x uint64
		err = binary.Read(r, order, &x)
		if err != nil {
			return nil, err
		}
		hdr.OldAddr = x
	}

	// SDNA index.
	err = binary.Read(r, order, &x)
	if err != nil {
		return nil, err
	}
	hdr.SDNAIndex = int(x)

	// Structure count.
	err = binary.Read(r, order, &x)
	if err != nil {
		return nil, err
	}
	hdr.Count = int(x)

	return hdr, nil
}

// BlockCode represents a rough type description of a block.
type BlockCode int

func (typ BlockCode) String() string {
	var m = map[BlockCode]string{
		CodeBR:   "BR",
		CodeCA:   "CA",
		CodeDATA: "DATA",
		CodeDNA1: "DNA1",
		CodeENDB: "ENDB",
		CodeGLOB: "GLOB",
		CodeLA:   "LA",
		CodeMA:   "MA",
		CodeME:   "ME",
		CodeOB:   "OB",
		CodeREND: "REND",
		CodeSC:   "SC",
		CodeSN:   "SN",
		CodeTE:   "TE",
		CodeTEST: "TEST",
		CodeWM:   "WM",
		CodeWO:   "WO",
	}
	s, ok := m[typ]
	if !ok {
		return "<unknown block code>"
	}
	return s
}

// Block codes.
const (
	CodeBR BlockCode = iota
	CodeCA
	CodeDATA
	CodeDNA1
	CodeENDB
	CodeGLOB
	CodeLA
	CodeMA
	CodeME
	CodeOB
	CodeREND
	CodeSC
	CodeSN
	CodeTE
	CodeTEST
	CodeWM
	CodeWO
)
