package main

import "fmt"

type Member struct {
	name string
	mtype string
}

type CISCustomer struct{

	name string
	mobileNo int64
}
type Modify interface {
	changeName(name *string)
}

//method
func (m *Member) changeName(name *string){
	(*m).name=*name
}
//method
func (customer *CISCustomer) changeName(name *string){
	customer.name=*name

}


func receiveInterface(m Modify){
	var name string ="Parameswari"
	var ptrname *string=&name
	m.changeName(ptrname)
}

func main() {
    member:=Member{"Ashok","silver"}
	customer:=CISCustomer{"Param",32435435}
	fmt.Println(customer)
    receiveInterface(&customer)
	receiveInterface(&member)
	fmt.Println(customer)
	fmt.Println(member)
}
