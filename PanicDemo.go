package main

import (
	"fmt"
	"log"
)

var names = []string{
"lobster",
"sea urchin",
"sea cucumber",
}


func accessElement() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	fmt.Println("My favorite sea creature is:", names[len(names)])
}
func main() {

	accessElement()
}