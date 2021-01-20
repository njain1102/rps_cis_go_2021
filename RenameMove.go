package main

import (
	"log"
	"os"
)

func main() {
	originalFile := "test.txt"
	newFile := "test2.txt"
	err := os.Rename(originalFile, newFile)
	if err != nil {
		log.Fatal(err)
	}
}