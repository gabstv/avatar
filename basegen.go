package avatar

import (
	"fmt"
	"image"
	"image/color"

	"github.com/gabstv/avatar/helpers"

	"github.com/anthonynsimon/bild/blend"
	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/transform"
	"github.com/aquilax/go-perlin"
)

type DefaultGen struct {
	cfg Config
	dig *DefaultDigest
}

func (g *DefaultGen) Generate(cfg Config) image.Image {
	if cfg.Width < 1 || cfg.Height < 1 {
		cfg.Width = 100
		cfg.Height = 100
	}
	g.cfg = cfg
	g.dig = &DefaultDigest{
		Seed: cfg.Seed,
	}
	return g.generate()
}

func (g *DefaultGen) generate() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, g.cfg.Width, g.cfg.Height))
	bgc := g.newbgcolor()
	helpers.Fill(img, bgc)
	c := &helpers.ImageContainer{img}
	g.mutator1(c)
	g.mutator2(c)
	g.mutator3(c, 50, 22, true)
	g.mutator3(c, 65, 345, false)
	return c.Image
}

func (g *DefaultGen) mutator1(c *helpers.ImageContainer) {
	var scale float64 = 44
	img := c.Image
	img2 := image.NewRGBA(img.Rect)
	mutator_r := float64(helpers.ReadByte(g.dig)) * 0.44
	mutator_g := float64(helpers.ReadByte(g.dig)) * 0.44
	mutator_b := float64(helpers.ReadByte(g.dig)) * 0.44
	sz := img2.Bounds().Size()
	p0 := perlin.NewPerlin(2.0, 2.0, 8, g.dig.i64seed)
	bx := float64(helpers.ReadInt64(g.dig) / 64)
	by := float64(helpers.ReadInt64(g.dig) / 64)
	fmt.Println(bx, by)
	fmt.Println(g.dig.i64seed)
	for y := 0; y < sz.Y; y++ {
		for x := 0; x < sz.X; x++ {
			noise := p0.Noise2D((bx+float64(x))/scale, (by+float64(y))/scale)
			if noise > 0.1 {
				c := color.RGBA{}
				c.A = 255
				c.R = byte(mutator_r * noise)
				c.G = byte(mutator_g * noise)
				c.B = byte(mutator_b * noise)
				img2.SetRGBA(x, y, c)
			}
		}
	}
	img2 = blur.Box(img2, 2.0)
	img2 = effect.Dilate(img2, 4)
	c.Image = blend.Add(img, img2)
}

func (g *DefaultGen) mutator2(c *helpers.ImageContainer) {
	var scale float64 = 11
	img := c.Image
	img2 := image.NewRGBA(img.Rect)
	helpers.Fill(img2, helpers.WhiteColor())
	mutator_r := float64(helpers.ReadByte(g.dig)) * 0.86
	mutator_g := float64(helpers.ReadByte(g.dig)) * 0.86
	mutator_b := float64(helpers.ReadByte(g.dig)) * 0.86
	sz := img2.Bounds().Size()
	p0 := perlin.NewPerlin(2.0, 2.0, 16, g.dig.i64seed)
	bx := float64(helpers.ReadInt64(g.dig)/111) - 14
	by := float64(helpers.ReadInt64(g.dig)/111) + 29
	fmt.Println(bx, by)
	fmt.Println(g.dig.i64seed)
	for y := 0; y < sz.Y; y++ {
		for x := 0; x < sz.X; x++ {
			noise := p0.Noise2D((bx+float64(x))/scale, (by+float64(y))/scale)
			c := color.RGBA{}
			c.A = 255
			c.R = byte(mutator_r * noise)
			c.G = byte(mutator_g * noise)
			c.B = byte(mutator_b * noise)
			img2.SetRGBA(x, y, c)
		}
	}
	img2 = blur.Box(img2, 3.0)
	img2 = effect.Dilate(img2, 10)
	c.Image = blend.Multiply(img, img2)
}

func (g *DefaultGen) mutator3(c *helpers.ImageContainer, intensity uint8, seedy int64, paintblack bool) {
	scale := 1.23
	p0 := perlin.NewPerlin(2.0, 2.0, 3, g.dig.i64seed-seedy)
	img2 := image.NewRGBA(image.Rect(0, 0, 8, 8))
	sz := img2.Bounds().Size()
	for y := 0; y < sz.Y; y++ {
		for x := 0; x < sz.X; x++ {
			noise := p0.Noise2D(float64(x)/scale, float64(y)/scale)
			if noise > 0.2 {
				img2.SetRGBA(x, y, helpers.WhiteColorA(intensity))
			} else if paintblack {
				img2.SetRGBA(x, y, helpers.BlackColorA(intensity))
			}
		}
	}
	sz2 := c.Image.Bounds().Size()
	img2 = transform.Resize(img2, sz2.X, sz2.Y, transform.NearestNeighbor)
	c.Image = blend.Overlay(c.Image, img2)
}

func (g *DefaultGen) newbgcolor() color.RGBA {
	c := color.RGBA{}
	c.A = 255
	sd := g.cfg.Seed
	if len(sd) < 1 {
		bb := make([]byte, 3)
		g.dig.Read(bb)
		sd = string(bb)
	} else {
		for len(sd) < 3 {
			sd = sd + sd
		}
	}
	sdb := []byte(sd)
	c.R = sdb[0] * (helpers.ReadByte(g.dig) / 4)
	c.G = sdb[1] * (helpers.ReadByte(g.dig) / 4)
	c.B = sdb[2] * (helpers.ReadByte(g.dig) / 4)
	return c
}
