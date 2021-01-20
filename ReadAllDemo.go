package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"log"
)
func main() {

	 readFile,rerr:=os.Open("tickets.csv")

	 defer readFile.Close()
	 if rerr!=nil{
	 	log.Fatal(rerr)
	 }
	data,_:= ioutil.ReadAll(readFile)
	fmt.Println(data)
}
