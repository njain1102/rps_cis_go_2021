package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func user(c chan string){
    c<-"How to open an account"
    umsg:=<-c
    fmt.Println(umsg)
	if len(umsg)>0{
		defer waitGroup.Done()
	}
}

func bot(c chan string){

     msg:=<-c
     fmt.Println(msg)
    if len(msg)>0{
    	defer waitGroup.Done()
	}
     c<-"Please open https://hdfcbank.com/" +
     	"accounts guide"
}

func main() {
	var c chan string = make(chan string)
     waitGroup.Add(2)
	 go bot(c)
	 go user(c)
	waitGroup.Wait()

}
