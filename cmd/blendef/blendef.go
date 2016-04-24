// blendef is a tool which regenerates the block structure definitions and the
// block parsing logic of the block package.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mewspring/blend"
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: blendef FILE.blend")
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	err := blendef(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
}

// blendef parses the provided blend file and generates the following Go files:
//
//    struct.go // structure definitions
//    parse.go  // block parser logic
func blendef(filePath string) (err error) {
	blend.WarnVersion = false
	b, err := blend.Parse(filePath)
	if err != nil {
		return err
	}
	defer b.Close()

	dna, err := b.GetDNA()
	if err != nil {
		return err
	}

	// Generate struct.go
	err = genStruct(b, dna)
	if err != nil {
		return err
	}

	// Generate parse.go
	err = genParse(b, dna)
	if err != nil {
		return err
	}

	return nil
}
