package main

import (
	"fmt"
	"time"
)

func main() {

	change := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			change <- i
			time.Sleep(time.Second)
		}
		close(change)
	}()

	for i := 0; i < 3; i++ {
		val := <-change
		fmt.Printf("reeived %d\n", val)
		//fmt.Println("here")
	}
}
