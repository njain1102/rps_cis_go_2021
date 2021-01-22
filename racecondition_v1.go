package main

import (
	"fmt"
	"sync"
)

var Wait sync.WaitGroup
var Counter int = 0
var Lock sync.Mutex
func main() {

	for routine := 1; routine <= 2; routine++ {

		Wait.Add(1)
		go Routine(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {

	for count := 0; count < 2; count++ {

		//run go build -race before and after lock
		Lock.Lock()
		value := Counter

		//time.Sleep(1 * time.Nanosecond)
		value++
		Counter = value
		Lock.Unlock()
	}

	Wait.Done()
}

