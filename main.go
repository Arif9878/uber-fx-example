package main

import (
	// stdhttp "net/http"

	"github.com/Arif9878/uber-fx-example/lib/config"
	"github.com/Arif9878/uber-fx-example/lib/db"
	"github.com/Arif9878/uber-fx-example/lib/http"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		db.Module,
		http.Module,
		// fx.Invoke(func(cfg *config.Config, handler stdhttp.Handler) error {
		// 	go stdhttp.ListenAndServe(cfg.HTTP.ListenAddress, handler)
		// 	return nil
		// }),
	)
	app.Run()
}
