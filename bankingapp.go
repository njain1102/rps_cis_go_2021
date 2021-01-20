package main

import (
	customer "../exportdemo"
)

func main() {

	date:=&customer.Date{2,2,1988}
	customerInstance:=&customer.Customer{
     "Parameswari",date}
	var icustomer customer.ICustomer=customerInstance
	icustomer.ShowCustomer()

}
