package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewkiz/pkg/pathutil"
	"github.com/mewmew/blend"
	"github.com/mewmew/blend/block"
)

func main() {
	flag.Parse()
	for _, filePath := range flag.Args() {
		err := blendview(filePath)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// blendview parses the provided blend file and extract the embedded thumbnail.
func blendview(filePath string) (err error) {
	// Parse blend file header and block headers.
	b, err := blend.Parse(filePath)
	if err != nil {
		return err
	}
	defer b.Close()

	// Parse DNA block.
	dna, err := b.GetDNA()
	if err != nil {
		log.Fatalln(err)
	}

	for _, blk := range b.Blocks {
		// Locate the TEST block.
		if blk.Hdr.Code != block.CodeTEST {
			continue
		}
		err = blk.ParseBody(b.Hdr.Order, dna)
		if err != nil {
			return err
		}
		buf, ok := blk.Body.([]byte)
		if !ok {
			return fmt.Errorf("blendview: invalid body type of %q block; expected []byte, got %T.", blk.Hdr.Code, blk.Body)
		}
		t, err := NewThumb(buf)
		if err != nil {
			return err
		}
		thumbPath := pathutil.TrimExt(filePath) + ".png"
		err = imgutil.WriteFile(thumbPath, t)
		if err != nil {
			return err
		}
		fmt.Println("Created:", thumbPath)
		return nil
	}
	return fmt.Errorf("unable to locate %q block in %q.", block.CodeTEST, filePath)
}

// Thumb is a thumbnail image based on the body of a TEST block.
type Thumb struct {
	w, h int
	// Pixels are stored in backwards order with respect to normal image raster
	// scan order, starting in the lower right corner, going right to left, and
	// then row by row from the bottom to the top of the image.
	//
	// Each pixel is stored in RGBA order.
	pix []byte
}

// NewThumb returns an image.Image based on the body of a TEST block.
func NewThumb(buf []byte) (t *Thumb, err error) {
	if len(buf) < 8 {
		return nil, fmt.Errorf("NewThumb: %q block body length (%d) below 8.", block.CodeTEST, len(buf))
	}
	t = &Thumb{
		w:   int(binary.LittleEndian.Uint32(buf)),
		h:   int(binary.LittleEndian.Uint32(buf[4:])),
		pix: buf[8:],
	}
	// Verify length of pix buf.
	want := 4 * t.w * t.h
	if len(t.pix) != want {
		return nil, fmt.Errorf("NewThumb: invalid pix buf length; expected %d got %d.", want, len(t.pix))
	}
	return t, nil
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (t *Thumb) At(x, y int) color.Color {
	// Pixels are stored in backwards order with respect to normal image raster
	// scan order, starting in the lower right corner, going right to left, and
	// then row by row from the bottom to the top of the image.
	off := (len(t.pix) - 4)
	off -= 4 * (x + y*t.w)
	// Each pixel is stored in RGBA order.
	r := t.pix[off]
	g := t.pix[off+1]
	b := t.pix[off+2]
	a := t.pix[off+3]
	return color.RGBA{r, g, b, a}
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (t *Thumb) Bounds() image.Rectangle {
	return image.Rect(0, 0, t.w, t.h)
}

// ColorModel returns the Image's color model.
func (t *Thumb) ColorModel() color.Model {
	return color.RGBAModel
}
