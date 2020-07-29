package main

import (
	"github.com/sirupsen/logrus"
	"math_app/app/interface/http_server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := newLogger()
	errorHandler := newErrorHandler(logger)

	cfg, err := loadConfig()
	if err != nil {
		errorHandler(err)
		return
	}

	server := http_server.New(cfg.APIPort, errorHandler)
	err = server.Start()
	if err != nil {
		errorHandler(err)
	}

	waitForOSSignal()

	err = server.Stop()
	if err != nil {
		errorHandler(err)
	}
}

func newLogger() logrus.FieldLogger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	return logger
}

func newErrorHandler(logger logrus.FieldLogger) func(error) {
	return func(err error) {
		logger.Error(err.Error())
	}
}

func waitForOSSignal() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)
	<-signalChan
}
