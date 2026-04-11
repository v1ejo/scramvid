package main

import (
	"log"
	"os"
)


func main() {
	if err := os.RemoveAll("video/"); err != nil {
		log.Fatal(err)
	}
}