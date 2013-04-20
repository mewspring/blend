package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mewmew/blend"
	"github.com/mewmew/blend/block"
)

// predef maps C types to untyped Go types.
var predef = map[string]string{
	// int types.
	"char":    "int",
	"short":   "int",
	"int":     "int",
	"long":    "int",
	"int64_t": "int",

	// uint types.
	"uchar":    "uint",
	"ushort":   "uint",
	"ulong":    "uint",
	"uint64_t": "uint",

	// float types.
	"float":  "float",
	"double": "float",

	// empty struct types.
	"void": "struct{}",
}

/// TODO: use text/template.

// genStruct generates Go structure definitions by parsing the DNA data.
//
// The output is stored in "struct.go".
func genStruct(b *blend.Blend, dna *block.DNA) (err error) {
	f, err := os.Create("struct.go")
	if err != nil {
		return err
	}
	defer f.Close()

	// Map type sizes.
	size := make(map[string]int)
	for i, typ := range dna.Types {
		size[typ] = dna.TypeSizes[i]
	}

	// Map basic type definitions.
	basic := make(map[string]string)
	for _, typ := range dna.Types {
		def, ok := predef[typ]
		if ok {
			n := size[typ]
			switch n {
			case 0:
				basic[typ] = def
			case 1:
				basic[typ] = def + "8"
			case 2:
				basic[typ] = def + "16"
			case 4:
				basic[typ] = def + "32"
			case 8:
				basic[typ] = def + "64"
			default:
				return fmt.Errorf("genStruct: size %d of basic type %q not supported.", n, typ)
			}
		}
	}

	// Generate Go pointer definition.
	fmt.Fprintf(f, "// NOTE: generated automatically by blendef for Blender v%d.\n", b.Hdr.Ver)
	fmt.Fprintln(f)
	fmt.Fprintln(f, "package block")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "import (")
	fmt.Fprintln(f, `	"fmt"`)
	fmt.Fprintln(f, ")")
	fmt.Fprintln(f)
	fmt.Fprintln(f, `// BlenderVer is the version of Blender used when generating the files`)
	fmt.Fprintln(f, `// "parse.go" and "struct.go" of this package. Use blendef to regenerate these`)
	fmt.Fprintln(f, `// files if this version differs from the blend file's version.`)
	fmt.Fprintf(f, "const BlenderVer = %d\n", b.Hdr.Ver)
	fmt.Fprintln(f)
	fmt.Fprintln(f, "// Pointer is the memory address of a structure when it was written to disk.")
	switch b.Hdr.PtrSize {
	case 4:
		fmt.Fprintln(f, "type Pointer uint32")
	case 8:
		fmt.Fprintln(f, "type Pointer uint64")
	}
	fmt.Fprintln(f)
	fmt.Fprintln(f, pointerCode)

	// Generate Go structure definitions.
	for index, st := range dna.Structs {
		fmt.Fprintln(f, "// SDNA index:", index)
		fmt.Fprintf(f, "type %s struct {\n", strings.Title(st.Type))
		for _, field := range st.Fields {
			// Parse and capitalize field name.
			name, isFunc, ptrCount, arraySizes, err := parseName(field.Name)
			if err != nil {
				return err
			}
			name = strings.Title(name)

			typ := field.Type
			def, ok := basic[typ]
			if ok {
				// Use Go basic type names.
				typ = def
			} else {
				// Capitalize non-basic type names.
				typ = strings.Title(typ)
			}

			if isFunc {
				if field.Type == "void" {
					fmt.Fprintf(f, "\t%s func()\n", name)
				} else {
					fmt.Fprintf(f, "\t%s func() %s\n", name, typ)
				}
			} else {
				array := new(bytes.Buffer)
				for _, arraySize := range arraySizes {
					fmt.Fprintf(array, "[%d]", arraySize)
				}
				if ptrCount > 0 {
					ptr := strings.Repeat("*", ptrCount)
					fmt.Fprintf(f, "\t%s %sPointer // %s%s%s\n", name, array, array, ptr, typ)
				} else {
					fmt.Fprintf(f, "\t%s %s%s\n", name, array, typ)
				}
			}
		}
		fmt.Fprintln(f, "}")
		fmt.Fprintln(f)
	}

	// Map of structure type names.
	structs := make(map[string]bool)
	for _, st := range dna.Structs {
		structs[st.Type] = true
	}

	// Generate empty Go structure definitions.
	for _, typ := range dna.Types {
		_, ok := basic[typ]
		if ok {
			continue
		}
		_, ok = structs[typ]
		if ok {
			continue
		}
		fmt.Fprintf(f, "type %s struct{}\n", strings.Title(typ))
	}

	return nil
}

// parseName parses the provided string and extracts name, pointer count and
// array and function information.
//
// Example input strings:
//    "id_type"
//    "**links"
//    "*point_cache[2]"
//    "clip[6][4]"
//    "(*free_edit)()"
func parseName(s string) (name string, isFunc bool, ptrCount int, arraySizes []int, err error) {
	// Parse function pointer.
	if len(s) > 1 && s[0] == '(' && s[1] == '*' {
		p := s[2:]
		end := strings.Index(p, ")")
		if end == -1 {
			return "", false, 0, nil, fmt.Errorf("parseName: unmatched opening parenthesis in %q.", s)
		}
		name = p[:end]
		return name, true, 0, nil, nil
	}

	// Parse pointer count.
	p := s
	for i := 0; i < len(p); i++ {
		if p[i] != '*' {
			p = p[i:]
			break
		}
		ptrCount++
	}

	// Parse name.
	pos := strings.Index(p, "[")
	if pos == -1 {
		return p, false, ptrCount, nil, nil
	}
	name = p[:pos]
	p = p[pos:]

	// Parse array sizes.
	for {
		// Get start position.
		pos := strings.Index(p, "[")
		if pos == -1 {
			return name, false, ptrCount, arraySizes, nil
		}
		p = p[pos+1:]

		// Get end position.
		end := strings.Index(p, "]")
		if end == -1 {
			return "", false, 0, nil, fmt.Errorf("parseName: unmatched opening bracket in %q.", s)
		}
		num := p[:end]
		p = p[end+1:]

		arraySize, err := strconv.Atoi(num)
		if err != nil {
			return "", false, 0, nil, err
		}
		arraySizes = append(arraySizes, arraySize)
	}
}

const pointerCode = `// Addr is a map from the memory address of a structure (when it was written to
// disk) to it's file block.
var Addr = make(map[uint64]*Block)

// Data translates the memory address into a usable pointer and returns it.
func (addr Pointer) Data() (data interface{}, err error) {
	blk, ok := Addr[uint64(addr)]
	if !ok {
		return nil, fmt.Errorf("Pointer.Data: unable to locate data for pointer %p.", addr)
	}
	return blk.Body, nil
}
`
