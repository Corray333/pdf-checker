package main

import (
	"pdfchecker/internal/app"
)

func main() {

	app := app.NewApp()

	app.Log.Info("Server is starting...")

	if err := app.Server.ListenAndServe(); err != nil {
		app.Log.Error("Server is not started:", err)
	}

	app.Log.Error("Server is stopped")
}
