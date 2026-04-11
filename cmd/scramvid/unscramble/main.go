package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/v1ejo/scramvid/internal/ffmpeg"
	"github.com/v1ejo/scramvid/internal/frames"
	"github.com/v1ejo/scramvid/internal/transform"
)

func main() {
	in := flag.String("in", "", "Video to unscramble")
	key := flag.String("key", "", "Secret to unscramble")
	out := flag.String("out", "", "Name of the output video")
	flag.Parse()

	if *in == "" || *key == "" || *out == "" {
		log.Fatalf("You must provide a input and output path and key")
	}

	err := ffmpeg.ExtractFrames(*in)
	if err != nil {
		log.Fatalf("Cannot extract frames")
	}
	err = ffmpeg.ExtractAudio(*in)
	if err != nil {
		log.Fatal("Cannot extract audio")
	}

	paths, err := filepath.Glob("video/frames/*.png")
	if err != nil {
		log.Fatal(err.Error())
	}

	pathOutputFrames := "video/scrambled"
	for _, path := range paths {
		name := filepath.Base(path)
		img, err := frames.OpenImage(path)
		if err != nil {
			log.Fatal(err)
		}
		outImg, err := transform.Unscramble(img, *key)
		if err != nil {
			log.Fatal(err)
		}
		err = frames.SaveImage(pathOutputFrames, outImg, name)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = ffmpeg.JoinFramesAndAudio(*out)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Finished")

	if err := os.RemoveAll("video/"); err != nil {
		log.Fatal(err)
	}
}
