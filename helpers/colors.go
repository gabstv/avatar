package helpers

import (
	"image/color"
)

func WhiteColor() color.RGBA {
	return color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
}

func WhiteColorA(a uint8) color.RGBA {
	return color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: a,
	}
}

func BlackColor() color.RGBA {
	return color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	}
}

func BlackColorA(a uint8) color.RGBA {
	return color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: a,
	}
}
