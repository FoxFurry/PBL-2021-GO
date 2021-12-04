package main

import (
	"foxy/application"
	"foxy/config"
	"foxy/internal/db"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	config.LoadConfig("config.yaml")
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	app := application.NewFoxyApp()

	go app.Start()

	<-sigChan
	app.Shutdown()
	db.GetDB().Close()

	return
}
