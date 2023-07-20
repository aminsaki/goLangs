package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	response, err := http.Get("https://httpbin.org/get")

	if err != nil {
		log.Fatalf("ERROR cant't call htttpbin.orgin")
	}

	defer response.Body.Close()

	io.Copy(os.Stdout, response.Body)
}
