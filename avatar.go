package avatar

import (
	"bytes"
	"image/png"
)

func GenerateBasicPNG(seed string) []byte {
	g := DefaultGen{}
	img := g.Generate(Config{
		Seed: seed,
	})
	b := new(bytes.Buffer)
	png.Encode(b, img)
	return b.Bytes()
}
