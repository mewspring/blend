package block_test

import (
	"fmt"
	"log"

	"github.com/mewmew/blend"
	"github.com/mewmew/blend/block"
)

func ExamplePointer() {
	// Parse blend file header and file blocks.
	b, err := blend.ParseAll("../testdata/block.blend")
	if err != nil {
		log.Fatalln(err)
	}
	for _, blk := range b.Blocks {
		lamp, ok := blk.Body.(*block.Lamp)
		if !ok {
			continue
		}
		// Get access to the actual data instead of the stored memory address.
		data, err := lamp.Curfalloff.Data()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", data)
		break
	}

	// Output:
	// &block.CurveMapping{Flag:1, Cur:0, Preset:0, Changed_timestamp:0, Curr:block.Rctf{Xmin:0, Xmax:1, Ymin:0, Ymax:1}, Clipr:block.Rctf{Xmin:0, Xmax:1, Ymin:0, Ymax:1}, Cm:[4]block.CurveMap{block.CurveMap{Totpoint:2, Flag:1, Range:256, Mintable:0, Maxtable:1, Ext_in:[2]float32{-0.70710677, 0.7071067}, Ext_out:[2]float32{-0.7071067, 0.70710677}, Curve:0x4a0d1c8, Table:0x0, Premultable:0x0}, block.CurveMap{Totpoint:0, Flag:0, Range:0, Mintable:0, Maxtable:0, Ext_in:[2]float32{0, 0}, Ext_out:[2]float32{0, 0}, Curve:0x0, Table:0x0, Premultable:0x0}, block.CurveMap{Totpoint:0, Flag:0, Range:0, Mintable:0, Maxtable:0, Ext_in:[2]float32{0, 0}, Ext_out:[2]float32{0, 0}, Curve:0x0, Table:0x0, Premultable:0x0}, block.CurveMap{Totpoint:0, Flag:0, Range:0, Mintable:0, Maxtable:0, Ext_in:[2]float32{0, 0}, Ext_out:[2]float32{0, 0}, Curve:0x0, Table:0x0, Premultable:0x0}}, Black:[3]float32{0, 0, 0}, White:[3]float32{1, 1, 1}, Bwmul:[3]float32{1, 1, 1}, Sample:[3]float32{0, 0, 0}}
}
