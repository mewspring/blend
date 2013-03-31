package main

import (
	"flag"
	"log"

	"github.com/mewmew/blend"
)

func main() {
	flag.Parse()
	for _, filePath := range flag.Args() {
		err := blendef(filePath)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// blendef parses the provided blend file and generates the following Go files:
//
//    block_def.go   // structure definitions
//    block_parse.go // block parser logic
func blendef(filePath string) (err error) {
	b, err := blend.Parse(filePath)
	if err != nil {
		return err
	}
	err = genStruct(b)
	if err != nil {
		return err
	}
	err = genParse(b)
	if err != nil {
		return err
	}
	return nil
}
