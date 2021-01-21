package main

import "fmt"

func printValue(v interface{}) {
	//type assertion by v.(int)
	fmt.Printf("The value of v is: %v", v.(int))
}

func main() {
	v := 10
	printValue(v)
}