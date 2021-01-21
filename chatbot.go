package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)
var c1 chan string = make(chan string)
var c2 chan string= make(chan string)
var c3 chan string= make(chan string)
var waitGroup sync.WaitGroup
//sender
func user(c chan <-string,id int){
    defer waitGroup.Done()
	c<-"How to open an account?"+strconv.Itoa(id)
}
//receiver
func bot(){

	for {
		select {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
		case msg3 := <- c3:
			fmt.Println(msg3)
		}
	}

}

func main() {
     waitGroup.Add(3)
	 go bot()
	 go user(c1,rand.Int())
	 go user(c2,rand.Int())
	 go user(c3,rand.Int())
	waitGroup.Wait()

}
