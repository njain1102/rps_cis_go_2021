package main

import "fmt"

type Employee struct{
Name string
}
func main() {

	var v interface{}=Employee{"Param"}
	res:=v.(Employee)
	fmt.Println(res)
}
