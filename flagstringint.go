package main

import (
	"flag"
	"fmt"
)

var (
	env  *string
	port *int
)
//init will be called before calling any other function inside the main package
func init() {
	env = flag.String("env", "development", "current environment")
	port = flag.Int("port", 3000, "port number")
}

func main() {

	flag.Parse()

	fmt.Println("env:", *env)
	fmt.Println("port:", *port)
}