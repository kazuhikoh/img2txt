package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s filepath", os.Args[0])
	}

	conf := func() image.Config {
		file, err := os.Open(os.Args[1])
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}

		c, _, err := image.DecodeConfig(file)
		if err != nil {
			log.Fatal(err)
		}

		return c
	}()

	img := func() image.Image {
		file, err := os.Open(os.Args[1])
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}

		i, _, err := image.Decode(file)
		if err != nil {
			log.Fatal(err)
		}

		return i
	}()

	table := make([][]color.Color, 0)
	for w := 0; w < conf.Width; w++ {
		row := make([]color.Color, 0)
		for h := 0; h < conf.Height; h++ {
			row = append(row, img.At(w, h))
		}
		table = append(table, row)
	}
	fmt.Println(table)
}
