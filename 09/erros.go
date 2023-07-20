package main

import "fmt"

func checkI(i int) error {

	if i > 5 {
		return fmt.Errorf("this is one test true")
	}
	return fmt.Errorf("this is one test false")

}

func main() {
	fmt.Println(checkI(10))
}
