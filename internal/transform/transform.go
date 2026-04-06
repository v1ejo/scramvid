package transform

import (
	"crypto/sha256"
	"image"
	_ "image/png"
	"os"
)

const tileSize int = 16

func Scramble(imagePath string, key string) (image.Image, error) {
	f, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	width, height := getImageDimesions(img)
	_ = width / tileSize
	_ = height / tileSize

	return nil, nil
}

func getImageDimesions(img image.Image) (width, height int) {
	bounds := img.Bounds()
	width = bounds.Dx()
	height = bounds.Dy()
	return
}

func gridToIndex(row, col, tilesX int) int {
	return row*tilesX + col
}

func indexToGrid(idx int, tilesX int) (row, col int) {
	row = idx / tilesX
	col = idx % tilesX
	return
}

func generateSeed(key string) int {
	h := sha256.New()
	h.Write([]byte(key))
	bs := h.Sum(nil)
	sum := 0
	for _, b := range bs {
		sum += int(b)
	}
	return sum
}

func baseList(n int) []int {
	base := make([]int, n)
	for i := 0; i < n; i++ {
		base[i] = i
	}
	return base
}
