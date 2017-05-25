package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
)

func main(){

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s filepath", os.Args[0])
	}

	file, err := os.Open(os.Args[1])
	defer file.Close()
	
	if err != nil {
		log.Fatal(err)
	}
	
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("(w,h)=(%d,%d)", config.Width, config.Height)

}
