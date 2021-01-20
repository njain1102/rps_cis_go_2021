package main

import(
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"strconv"
)
func main() {

	   rfile,err :=os.Create("RandomData.txt",)
	   if err!=nil{
	   	fmt.Println(err)
	   }

	   defer rfile.Close()
	   for i:=0;i<100;i++{
		   rfile.WriteString(strconv.Itoa(rand.Int())+"\n")
	   }
	w := bufio.NewWriter(rfile)
	for i:=0;i<100;i++{
		n4,_:=w.WriteString(strconv.Itoa(rand.Int())+"\n")
		fmt.Printf("wrote %d bytes\n", n4)
	}


}
