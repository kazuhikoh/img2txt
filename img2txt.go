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

	for h := 0; h < conf.Height; h++ {
		for w := 0; w < conf.Width; w++ {
			fmt.Print(colorize("■", img.At(w, h)))
		}
		fmt.Println()
	}
}

func colorize(word string, c color.Color) string {
	r, g, b, _ := c.RGBA()
	r = ((r + 1) / 256) - 1
	g = ((g + 1) / 256) - 1
	b = ((b + 1) / 256) - 1
	return "\033[38;2;" + fmt.Sprint(r) + ";" + fmt.Sprint(g) + ";" + fmt.Sprint(b) + "m" + word + "\033[0m"
}
