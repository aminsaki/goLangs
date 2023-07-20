package main

import "fmt"

type Users struct {
	x int
	y int
}

func (u Users) Sum() error {

	if u.x > u.y {
		return fmt.Errorf("x big y")
	}
	return fmt.Errorf("x big y")

}

func main() {
	us := Users{5, 6}
	result := us.Sum()
	fmt.Println(result)

}
