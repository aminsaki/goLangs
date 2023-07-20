package main

func main() {

	//x := 20
	//y := 50
	//
	//if x > y {
	//	fmt.Println("x is big")
	//}
	//if x > 100 {
	//	fmt.Println("x isverybig")
	//} else {
	//	fmt.Println("x not that  big")
	//}
	println(switchs(2))
}

func switchs(x int) string {
	switch x {
	case 1:
		return "a"
	case 2:
		return "b"
	case 3:
		return "c"
	default:
		return "not code"
	}
}
