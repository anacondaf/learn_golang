package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup {}

func loopFor10s(timing byte) {
	time.Sleep(5 * time.Second)

	fmt.Printf("%v\n", "I am done in goroutine")

	wg.Done()
}

func main() {
	//A
	fmt.Println("Hello Goroutine")

	//B
	wg.Add(1)

	go loopFor10s(10)

	//C
	var sum = 99+99
	fmt.Println(sum)
	
	wg.Wait()
}