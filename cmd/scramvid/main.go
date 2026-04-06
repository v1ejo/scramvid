package main

import (
	"flag"
	"image/png"
	"log"
	"os"

	"github.com/v1ejo/scramvid/internal/transform"
)

func main() {
	in := flag.String("in", "", "PNG file for now")
	key := flag.String("key", "", "Key to scramble")
	out := flag.String("out", "", "Output path PNG")
	flag.Parse()
	if *in == "" || *key == "" || *out == "" {
		log.Fatal("You must give a input file name, output file name and key")
	}

	image, err := transform.Scramble(*in, *key)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	png.Encode(f, image)
}
