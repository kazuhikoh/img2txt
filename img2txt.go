package main

import {
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
}

func main(){

	file, err := os.Open("sample.png")
	defer file.Close()
	
	if err != nil {
		log.Fatal(err)
		return
	}
	
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("(w,h)=(%d,%d)", config.Width, config.Height)

}
