package main

import "fmt"

//
//type Animal interface {
//	Sound() string
//}
//
//type Cat struct {
//}
//
//func (c Cat) Sound() string {
//	return "Meow"
//}
//
//type Dog struct{}
//
//func (d Dog) Sound() string {
//	return "Woof"
//}
//func MakeSound(a Animal) {
//	fmt.Println(a.Sound())
//}

//type Sms interface {
//	Send(number int, messages string)
//}
//
//type Nanger struct {
//}
//
//func (n Nanger) Send(number int, messages string) {
//
//	fmt.Printf("Sending SMS to number %d: %s\n", number, messages)
//}
//
//func MakeSms(s Sms, number int, message string) {
//	s.Send(number, message)
//}

type User interface {
	Show(id int)
}

type Res struct {
}

func (r Res) Show(id int) error {
	if id > 10 {
		return fmt.Errorf("no")
	}
	return fmt.Errorf("yes")
}

func MakeUser(u User, id int) {
	u.Show(id)
}

func main() {
	us := Res{}
	fmt.Println(us.Show(20))

	//n := Nanger{}
	//MakeSms(n, 9904289707, "This is one message")

	//animals := []Animal{Cat{}, Dog{}}
	//
	//for _, animal := range animals {
	//	MakeSound(animal)
	//}

}
