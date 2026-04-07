package main

import (
	"flag"
	"image"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/v1ejo/scramvid/internal/frames"
	"github.com/v1ejo/scramvid/internal/transform"
)

const (
	DECODE string = "decode"
	ENCODE string = "encode"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must provide a mode: encode or decode")
	}

	in := flag.String("in", "", "PNG file for now")
	out := flag.String("out", "", "Output path PNG")
	key := flag.String("key", "secret", "Key to scramble")
	mode := strings.ToLower(os.Args[1])

	if err := flag.CommandLine.Parse(os.Args[2:]); err != nil {
		log.Fatal(err)
	}

	var outImg image.Image

	if *in == "" || *out == "" {
		log.Fatal("You must give a input file name and output file name")
	}

	inImg, err := frames.OpenImage(*in)
	if err != nil {
		log.Fatal("Unable to open image")
	}

	switch mode {
	case ENCODE:
		outImg, err = transform.Scramble(inImg, *key)
		if err != nil {
			log.Fatal(err)
		}
	case DECODE:
		outImg, err = transform.Unscramble(inImg, *key)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalln("Invalid option")
	}

	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	png.Encode(f, outImg)
}
