package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func useranem(first string, last string) (string, string) {

	return first + last, last + first
}

func main() {

	//	fmt.Println(sum(1, 2))
	last_name, first_name := useranem("RSAS", "AMIN")

	fmt.Println(last_name, first_name)
}
