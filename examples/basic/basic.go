package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/gabstv/avatar"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: basic stringseed")
		os.Exit(1)
	}
	seed := os.Args[1]
	fmt.Println("basic", seed)

	g := avatar.DefaultGen{}
	img := g.Generate(avatar.Config{
		Seed: seed,
	})
	file, err := os.Create(seed + ".png")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Success!")
}
