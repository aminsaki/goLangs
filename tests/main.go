package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func IsEven(n int) bool {
	return n%2 == 0
}

func main() {
	result := Add(5, 6)
	fmt.Println(result)
	fmt.Println("==========================================")

	result2 := IsEven(51)
	fmt.Println(result2)
}
