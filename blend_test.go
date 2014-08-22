package blend_test

import (
	"fmt"
	"log"

	"github.com/mewmew/blend"
)

func ExampleBlend() {
	// Parse blend file header and block headers.
	b, err := blend.Parse("testdata/block.blend")
	if err != nil {
		log.Fatalln(err)
	}
	defer b.Close()

	// Parse DNA block.
	dna, err := b.GetDNA()
	if err != nil {
		log.Fatalln(err)
	}

	// Parse and output the body of block 11.
	blk := b.Blocks[11]
	err = blk.ParseBody(b.Hdr.Order, dna)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%#v\n", blk.Body)

	// Output:
	// &block.ScrVert{Next:0x0, Prev:0x4353b68, Newv:0x0, Vec:block.Vec2s{X:1920, Y:1053}, Flag:1, Editflag:1}
}
