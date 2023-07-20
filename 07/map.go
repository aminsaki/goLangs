package main

import "fmt"

func main() {

	bar := map[string]string{
		"amin":  "tru",
		"reza":  "false",
		"theri": "false",
		"saki":  "true",
	}
	delete(bar, "theri")
	//bar["ta"] = "false"

	fmt.Println(bar)

	//if bar["amin"] {
	//	fmt.Println("true")
	//} else {
	//	fmt.Println("false")
	//
	//}

	//fmt.Println(bar, map[string]string{"adfs": "true"})

}
