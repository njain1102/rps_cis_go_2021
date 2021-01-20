package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"bufio"
)
var fileInstance *os.File
var fileInfo os.FileInfo
var err error
//read json data and transfer to csv

func main() {
	defer fileInstance.Close()
	fileInstance, err = os.Create("tickets.csv")
       if err!=nil{
       	log.Fatal(err)
	   }
	   log.Println(fileInstance)
	fileInfo,err=os.Stat("tickets.csv")
	fmt.Println(fileInfo.Size())
    fileInstance.WriteString("ID\n")
	for i:=0;i<100000;i++{
		fileInstance.WriteString(strconv.Itoa(rand.Intn(10000))+"\n")
	}

    w:=bufio.NewWriter(fileInstance)
	for i:=0;i<100000;i++{
		w.WriteString(strconv.Itoa(rand.Intn(10000))+"\n")
	}

}
