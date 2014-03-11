package block

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
)

// DNA stores information about the various structures contained within a blend
// file.
type DNA struct {
	// Names is a slice which contains the name of each structure field. It can
	// be accessed with the name index.
	Names []string
	// Types is a slice which contains the name of each type. It can be accessed
	// with the type index.
	Types []string
	// TypeSizes is a slice which contains the size of each type. It can be
	// accessed with the type index.
	TypeSizes []int
	// Structs is a slice which contains structure definitions. The SDNA index
	// can be used to access individual structures.
	Structs []DNAStruct
}

// DNAStruct stores information about a structure.
type DNAStruct struct {
	// Type is the type name of the structure.
	Type string
	// Fields is a slice which contains the structure's field definitions.
	Fields []DNAField
}

// DNAField stores information about a structure field.
type DNAField struct {
	// Type is the type name of the field.
	Type string
	// Name contains information about the field's name, pointer count and array
	// and function information.
	//
	// Below are a couple of examples:
	//    "id_type"
	//    "**links"
	//    "*point_cache[2]"
	//    "clip[6][4]"
	//    "(*free_edit)()"
	Name string
}

// ParseDNA parses and returns the body of the "DNA1" block.
func ParseDNA(r io.Reader, order binary.ByteOrder) (body *DNA, err error) {
	br := bufio.NewReader(r)
	rawID := make([]byte, 4)

	// Identifier.
	_, err = io.ReadFull(br, rawID)
	if err != nil {
		return nil, err
	}
	id := string(rawID)
	if id != "SDNA" {
		return nil, fmt.Errorf("block.ParseDNA: invalid identifier %q", id)
	}

	// Name identifier.
	_, err = io.ReadFull(br, rawID)
	if err != nil {
		return nil, err
	}
	nameID := string(rawID)
	if nameID != "NAME" {
		return nil, fmt.Errorf("block.ParseDNA: invalid name identifier %q", nameID)
	}

	// Name count.
	var x32 int32
	err = binary.Read(br, order, &x32)
	if err != nil {
		return nil, err
	}
	nameCount := int(x32)

	// Names.
	body = new(DNA)
	body.Names = make([]string, nameCount)
	var total int
	for i := range body.Names {
		buf, err := br.ReadSlice(0x00)
		if err != nil {
			return nil, err
		}
		total += len(buf)
		body.Names[i] = string(buf[:len(buf)-1])
	}
	err = align(br, total, 4)
	if err != nil {
		return nil, err
	}

	// Type identifier.
	_, err = io.ReadFull(br, rawID)
	if err != nil {
		return nil, err
	}
	typeID := string(rawID)
	if typeID != "TYPE" {
		return nil, fmt.Errorf("block.ParseDNA: invalid type identifier %q", typeID)
	}

	// Type count.
	err = binary.Read(br, order, &x32)
	if err != nil {
		return nil, err
	}
	typeCount := int(x32)

	// Types.
	body.Types = make([]string, typeCount)
	total = 0
	for i := range body.Types {
		buf, err := br.ReadSlice(0x00)
		if err != nil {
			return nil, err
		}
		total += len(buf)
		body.Types[i] = string(buf[:len(buf)-1])
	}
	err = align(br, total, 4)
	if err != nil {
		return nil, err
	}

	// Type length identifier.
	_, err = io.ReadFull(br, rawID)
	if err != nil {
		return nil, err
	}
	lenID := string(rawID)
	if lenID != "TLEN" {
		return nil, fmt.Errorf("block.ParseDNA: invalid type length identifier %q", lenID)
	}

	// Type sizes.
	var x16 int16
	body.TypeSizes = make([]int, typeCount)
	total = 0
	for i := range body.TypeSizes {
		err = binary.Read(br, order, &x16)
		if err != nil {
			return nil, err
		}
		total += 2
		body.TypeSizes[i] = int(x16)
	}
	err = align(br, total, 4)
	if err != nil {
		return nil, err
	}

	// Structure identifier.
	_, err = io.ReadFull(br, rawID)
	if err != nil {
		return nil, err
	}
	structID := string(rawID)
	if structID != "STRC" {
		return nil, fmt.Errorf("block.ParseDNA: invalid structure identifier %q.", structID)
	}

	// Structure count.
	err = binary.Read(br, order, &x32)
	if err != nil {
		return nil, err
	}
	structCount := int(x32)

	// Structures.
	body.Structs = make([]DNAStruct, structCount)
	for i := range body.Structs {
		// Structure type.
		err = binary.Read(br, order, &x16)
		if err != nil {
			return nil, err
		}
		body.Structs[i].Type = body.Types[x16]

		// Field count.
		err = binary.Read(br, order, &x16)
		if err != nil {
			return nil, err
		}
		fieldCount := int(x16)
		body.Structs[i].Fields = make([]DNAField, fieldCount)

		// Fields.
		for j := range body.Structs[i].Fields {
			// Field type.
			err = binary.Read(br, order, &x16)
			if err != nil {
				return nil, err
			}
			body.Structs[i].Fields[j].Type = body.Types[x16]

			// Field name.
			err = binary.Read(br, order, &x16)
			if err != nil {
				return nil, err
			}
			body.Structs[i].Fields[j].Name = body.Names[x16]
		}
	}

	return body, nil
}

// align advances the reader so that it is aligned with n, were total
// corresponds to the number of bytes read so far.
func align(r io.Reader, total int, n int) (err error) {
	i := total % n
	if i > 0 {
		_, err = io.CopyN(ioutil.Discard, r, int64(n-i))
		if err != nil {
			return err
		}
	}
	return nil
}
