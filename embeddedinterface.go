package main

type IPayment interface{
	payment()
}

type ICancel interface{
	cancel()
}

type ECall interface{
	payment()
	cancel()
}

type Booking struct {
	n int
	amount int
}

type Trading struct{
	id int
	amount int64
}

type FundTransfer struct{
	payeeAccountNo int64
	amount int32
}

//methods

func(b *Booking) payment(){

}

func(t *Trading) payment(){

}

func(f *FundTransfer) payment(){

}


func(b *Booking) cancel(){

}


func main() {

	booking:=Booking{5,5789}
	trading:=Trading{4858,450000}
	fundTransfer:=FundTransfer{485824233,34000}
  //runtime polymorphism
	//interface 1
	//var ipayment IPayment = &booking
	//ipayment.payment()
	var ipayment IPayment= &trading
	ipayment.payment()
	ipayment= &fundTransfer
	ipayment.payment()

 //interface 2
  // var icancel ICancel=&booking
  //  icancel.cancel()

  //embedding
  //interface 3
  var ecall ECall=&booking
  ecall.payment()
  ecall.cancel()
	
}
