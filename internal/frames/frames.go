package frames

import (
	"image"
	"image/png"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
)

func OpenImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func SaveImage(path string, img image.Image, name string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(path, name))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	return png.Encode(f, img)
}
