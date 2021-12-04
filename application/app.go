package application

import (
	"context"
	"fmt"
	"foxy/internal/http/controller"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type IApp interface {
	Start()
	Shutdown()
}

type foxyApp struct {
	server *http.Server
}

func NewFoxyApp() IApp {
	serverAddress := fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port"))

	ctrl := controller.NewFoxyController()

	app := foxyApp{
		server: &http.Server{
			Addr:    serverAddress,
			Handler: ctrl.GetRouter(),
		},
	}

	return &app
}

func (d *foxyApp) Start() {
	if err := d.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Error while running the server: %v", err)
	}
}

func (d *foxyApp) Shutdown() {
	if err := d.server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error while shutting down the server: %v", err)
	}
}
