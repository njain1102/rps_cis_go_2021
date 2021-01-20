package exportdemo

import "fmt"

type Date struct{
	Day int
	Month int
	Year int
}
type Customer struct{
	Name string
	Dob *Date
}

type ICustomer interface{
	ShowCustomer()
	ModifyCustomer(d *Date)

}

func (c *Customer) ShowCustomer(){
	fmt.Printf("%v",c.Name)
	fmt.Printf("%v/%v/%v",c.Dob.Day,c.Dob.Month,c.Dob.Year)
}
func (c *Customer) ModifyCustomer(d *Date){
	(*c).Dob=d
}
func (c *Customer) modifyCustomer(d *Date){
	(*c).Dob=d
}

var BankName string


