package main

import "fmt"

type Trade struct {
	Symbol string  ///public
	Volume int     ///public
	Price  float64 ///public
	Buy    bool    ///public
}

func main() {

	//fmt.Println(Trade{})
	//fmt.Println(Trade{"Apple", 10, 10.2, true})
	tr := Trade{"Apple", 10, 10.2, true}
	fmt.Println(tr.Symbol)

}
