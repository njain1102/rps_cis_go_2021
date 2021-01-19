package main

import "fmt"

func main() {
//order of execution

	defer fmt.Println("ends...")
	fmt.Println("Starts")
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
}