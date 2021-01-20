package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo os.FileInfo
	ferr      error
)

func main() {
	// Stat returns file info. It will return
	// an error if there is no file.
	fileInfo, ferr = os.Stat("test.txt")
	if ferr != nil {
		log.Fatal(ferr)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}