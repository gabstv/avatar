package helpers

import (
	"encoding/binary"
	"image"
	"image/color"
	"io"
)

func ReadByte(r io.Reader) byte {
	b := make([]byte, 1)
	r.Read(b)
	return b[0]
}

func ReadInt(r io.Reader) int {
	b := make([]byte, 8)
	r.Read(b)
	i64, _ := binary.Varint(b)
	return int(i64)
}

func ReadInt64(r io.Reader) int64 {
	b := make([]byte, 8)
	r.Read(b)
	i64, _ := binary.Varint(b)
	return i64
}

func Fill(img *image.RGBA, c color.RGBA) {
	if img == nil {
		return
	}
	sz := img.Bounds().Size()
	for y := 0; y < sz.Y; y++ {
		for x := 0; x <= sz.X; x++ {
			img.SetRGBA(x, y, c)
		}
	}
}
