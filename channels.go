package main
import (
	"fmt"
	"time"
)
func pinger(c chan string) {
	for i := 0; ; i++ {
		//send message
		c <- "ping"
		//time.Sleep(time.Second * 1)
	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func printer1(c chan string) {
	for {
		//receive message
		msg := <- c
		fmt.Println("Printer1",msg)
		time.Sleep(time.Second * 1)
	}
}

func printer2(c chan string) {
	for {
		//receive message
		msg := <- c
		fmt.Println("Printer2",msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	//creating channel
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer1(c)
	go printer2(c)
	var input string
	fmt.Scanln(&input)
}