package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func switchassertprintValue(v interface{}) {
	switch v := v.(type) {
	case string:
		fmt.Printf("%v is a string\n", v)
	case int:
		fmt.Printf("%v is an int\n", v)
	default:
		fmt.Printf("The type of v is unknown\n")
	}
}

func main() {
	v := 10
	switchassertprintValue(v)
	do(21)
	do("hello")
	do(true)
}