package main

import "fmt"

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

type person struct {
	firstName string
	lastName string
}

func (p person) fullName() {
	fmt.Printf("%s %s",p.firstName,p.lastName)
}
func main() {
	nums := []int{78, 109, 2, 563, 300}
	p := person {
		firstName: "RPS",
		lastName: "Consulting",
	}
	largest(nums)


	fmt.Printf("Welcome ")
	//defer called later
	defer p.fullName()
	fmt.Printf("Welcome ")
	//defer p.fullName()
	//fmt.Printf("Welcome ")
}