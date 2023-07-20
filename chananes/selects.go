package main

import (
	"fmt"
	"time"
)

func main() {

	//ch1, ch2 := make(chan int), make(chan int)
	//
	//go func() {
	//	ch1 <- 42
	//}()
	//
	//select {
	//case val := <-ch1:
	//	fmt.Printf("got %d from ch1\n", val)
	//
	//case val := <-ch2:
	//	fmt.Printf("got %d from ch2\n", val)
	//
	//}
	//fmt.Println("===========================================================")

	out := make(chan float64)

	go func() {
		time.Sleep(200 * time.Microsecond)
		out <- 3.14
	}()

	select {
	case val := <-out:
		fmt.Printf("got %f from ch1\n", val)

	case <-time.After(4000 * time.Microsecond):
		fmt.Println("timeout")
	}

}
