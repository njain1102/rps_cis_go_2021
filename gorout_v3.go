package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var wgg sync.WaitGroup

func walkDir(dir string) {
defer wgg.Done()

visit := func(path string, f os.FileInfo, err error) error {
if f.IsDir() && path != dir {
wgg.Add(1)
go walkDir(path)
return filepath.SkipDir
}
if f.Mode().IsRegular() {
fmt.Printf("Visited: %s File name: %s Size: %d bytes\n",
path, f.Name(), f.Size())
}
return nil
}

filepath.Walk(dir, visit)
}

func main() {
flag.Parse()
root := "./" //flag.Arg(0)

wgg.Add(1)
walkDir(root)
wgg.Wait()
}