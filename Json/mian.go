package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var data = `{
  
  "user":"amin",
   "avg":"187",
   "amount":"2555.00"

}`

type Request struct {
	User   string `json:"user"`
	Avg    string `json:"avg"`
	Amount string `json:"amount"`
}

func main() {

	rander := bytes.NewBufferString(data)

	decode := json.NewDecoder(rander)

	request := &Request{}
	if err := decode.Decode(request); err != nil {

		log.Fatal("ERROR cat't decode - %S", err)
	}

	fmt.Printf("got: % +v\n", request)

	prevBalance := "2000000"

	var response = map[string]interface{}{

		"ok":      true,
		"code":    200,
		"balance": request.Amount + prevBalance,
	}

	encdoe := json.NewEncoder(os.Stdout)

	if err := encdoe.Encode(response); err != nil {
		log.Fatal("ERROR cat't decode - %S", err)
	}
	fmt.Printf("got: % +v\n", response)

}
