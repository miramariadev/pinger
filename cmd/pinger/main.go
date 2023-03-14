package main

import (
	"context"
	"flag"

	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"

	"gitlab.studionx.ru/id/pinger/app/errors"
	"gitlab.studionx.ru/id/pinger/app/pinger/external"
	"gitlab.studionx.ru/id/pinger/config"
)

const (
	defaultConfigPath = "init.toml"
	errorsChanLength  = 10
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	appConfig := getAppConfig()

	appErrorsChan := make(chan error, errorsChanLength)
	stopAppChan := make(chan bool, 1)
	systemTerminateSignalsChan := make(chan os.Signal, 1)

	signal.Notify(systemTerminateSignalsChan, os.Interrupt, syscall.SIGTERM)

	handlerErrors := errors.NewHandlerErrors(
		appErrorsChan,
		stopAppChan,
	)

	go handlerErrors.Handle()
	go handleSystemTerminateSignal(systemTerminateSignalsChan, stopAppChan)

	pingerServer := pinger.NewPingerServer(
		appConfig.Pinger.Addr,
		appErrorsChan,
	)

	go pingerServer.Run(ctx)

	<-stopAppChan

	cancel()
}

func getAppConfig() *config.Config {
	flag.Parse()

	configApp := config.NewConfig()

	_, err := toml.DecodeFile(defaultConfigPath, configApp)
	if err != nil {
		log.Fatalf("App config error: %s", err.Error())
	}
	err = configApp.IsValid()
	if err != nil {
		log.Fatalf("App config error: %s", err.Error())
	}

	return configApp
}

func handleSystemTerminateSignal(
	signalChan chan os.Signal,
	stopAppChan chan bool,
) {
	<-signalChan
	stopAppChan <- true
}
