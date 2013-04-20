// Package block implements parsing of blend file blocks.
//
// One unique feature of blend files is that they contain a full definition of
// every structure used in it's file blocks. The structure definitions are
// stored in the DNA block.
//
// All block structure definitions ("struct.go") and the block parsing logic
// ("parse.go") have been generating by parsing the DNA block of
// "testdata/block.blend".
//
// The tool which was used to generate these two files is located at:
//    github.com/mewmew/blend/cmd/blendef
//
// More complex blend files may contain structures which are not yet defined in
// this package. If so, use blendef to regenerate "struct.go" and "parse.go" for
// the given blend file.
package block

import (
	"encoding/binary"
	"io"
	"log"
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
	case "AR\x00\x00":
		hdr.Code = CodeAR
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
	case "IM\x00\x00":
		hdr.Code = CodeIM
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
	case "SR\x00\x00":
		hdr.Code = CodeSR
	case "TE\x00\x00":
		hdr.Code = CodeTE
	case "TEST":
		hdr.Code = CodeTEST
	case "TX\x00\x00":
		hdr.Code = CodeTX
	case "WM\x00\x00":
		hdr.Code = CodeWM
	case "WO\x00\x00":
		hdr.Code = CodeWO
	default:
		log.Printf("Header.ParseHeader: block code %q not yet implemented.\n", code)
		hdr.Code = CodeUnknown
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
		CodeAR:   "AR",
		CodeBR:   "BR",
		CodeCA:   "CA",
		CodeDATA: "DATA",
		CodeDNA1: "DNA1",
		CodeENDB: "ENDB",
		CodeGLOB: "GLOB",
		CodeIM:   "IM",
		CodeLA:   "LA",
		CodeMA:   "MA",
		CodeME:   "ME",
		CodeOB:   "OB",
		CodeREND: "REND",
		CodeSC:   "SC",
		CodeSN:   "SN",
		CodeSR:   "SR",
		CodeTE:   "TE",
		CodeTEST: "TEST",
		CodeTX:   "TX",
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
	CodeAR BlockCode = iota
	CodeBR
	CodeCA
	CodeDATA
	CodeDNA1
	CodeENDB
	CodeGLOB
	CodeIM
	CodeLA
	CodeMA
	CodeME
	CodeOB
	CodeREND
	CodeSC
	CodeSN
	CodeSR
	CodeTE
	CodeTEST
	CodeTX
	CodeWM
	CodeWO

	CodeUnknown BlockCode = -1
)
