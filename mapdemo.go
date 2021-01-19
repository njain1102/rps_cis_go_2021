package main

import (
	"fmt"
	"math/rand"
	"strconv"
)
type DDate struct{
	day int
	month int
	year int
}
type Product struct{
	ProductId int64
	Name string
	DOP DDate
	Cost int32
}

func main() {

	//generate cart with 5 products
	products:=make(map[int]Product)
	fmt.Println("Enter number of products....")
	var count int
	fmt.Scanf("%d",&count)
	for i:=0;i<count;i++{
        products[i]=Product{int64(rand.Int31n(1000000)),
        	"Product"+strconv.Itoa(i),
        DDate{
        	rand.Intn(28),10,2020},
        875687+rand.Int31n(100)}
	}
	//fmt.Printf("Product Details%+v", products)
	//Key value separation
	for key,value:=range products{
		fmt.Printf("Key %d => %+v\n",key,value)
	}


}
