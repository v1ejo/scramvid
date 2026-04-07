package main

import (
	"flag"
	"image/png"
	"log"
	"os"

	"github.com/v1ejo/scramvid/internal/frames"
	"github.com/v1ejo/scramvid/internal/transform"
)

func main() {
	in := flag.String("in", "", "PNG file for now")
	out := flag.String("out", "", "Output path PNG")
	key := flag.String("key", "secret", "Key to scramble")
	flag.Parse()

	if *in == "" || *out == "" {
		log.Fatal("You must give a input file name and output file name")
	}

	inImg, err := frames.OpenImage(*in)
	if err != nil {
		log.Fatal("Unable to open image")
	}

	outImg, err := transform.Scramble(inImg, *key)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	png.Encode(f, outImg)
}
