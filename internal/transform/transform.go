package transform

import (
	"crypto/sha256"
	"image"
	_ "image/png"
	"math/rand/v2"
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
	tilesX := width / tileSize
	tilesY := height / tileSize
	n := tilesX * tilesY
	base := baseSlice(n)
	perm := shuffleSlice(base, key)
	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	for src := 0; src < n; src++ {
		srcRow, srcCol := indexToGrid(src, tilesX)
		srcX, srcY := (srcCol * tileSize), (srcRow * tileSize)
		dst := perm[src]
		dstRow, dstCol := indexToGrid(dst, tilesX)
		dstX, dstY := (dstCol * tileSize), (dstRow * tileSize)
		for y  := 0; y  < tileSize; y++ {
			for x  := 0; x  < tileSize; x++ {
				color := img.At(srcX + x, srcY + y)
				newImage.Set(dstX + x, dstY + y, color)
			}
		}
	}

	return newImage, nil
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

func baseSlice(n int) []int {
	base := make([]int, n)
	for i := 0; i < n; i++ {
		base[i] = i
	}
	return base
}

func shuffleSlice(slice []int, key string) []int {
	if key == "" {
		key = "secret"
	}
	seed := generateSeed(key)
	s := rand.NewPCG(uint64(seed), 42)
	r := rand.New(s)

	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	return slice
}
