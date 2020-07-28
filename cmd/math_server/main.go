package main

import (
	"fmt"
	"math_app/app/interface/http_server"
)

func main() {
	server := http_server.New(8080, handleError)
	err := server.Start()
	if err != nil {
		handleError(err)
	}
	select {}
}

func handleError(err error) {
	fmt.Println(err)
}
