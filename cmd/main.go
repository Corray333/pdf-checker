package main

import (
	"testp/internal/app"

	_ "github.com/lib/pq"
)

func main() {

	app := app.NewApp()

	app.Log.Info("Server is starting...")

	if err := app.Server.ListenAndServe(); err != nil {
		app.Log.Error("Server is not started")
	}

	app.Log.Error("Server is stopped")
}
