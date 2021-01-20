package main


import (
	"log"
	"os"
)

var (
	ffileInfo *os.FileInfo
	f_err      error
)

func main() {
	// Stat returns file info. It will return
	// an error if there is no file.
	ffileInfo, f_err := os.Stat("test2.txt")
	if f_err != nil {
		if os.IsNotExist(f_err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File does exist. File information:")
	log.Println(ffileInfo)
}