package main

import (
	"fmt"
	"math_app/app/interface/http_server"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		handleError(err)
		return
	}

	server := http_server.New(cfg.APIPort, handleError)
	err = server.Start()
	if err != nil {
		handleError(err)
	}
	select {}
}

func handleError(err error) {
	fmt.Println(err)
}
