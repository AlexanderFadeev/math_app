package main

import (
	"fmt"
	"math_app/app/interface/http_server"
	"os"
	"os/signal"
	"syscall"
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

	waitForOSSignal()

	err = server.Stop()
	if err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	fmt.Println(err)
}

func waitForOSSignal() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)
	<-signalChan
}
