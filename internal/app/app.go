package app

import (
	"log/slog"
	"net/http"
	"pdfchecker/internal/config"
	"pdfchecker/internal/logger"
	"pdfchecker/internal/server"
)

type App struct {
	*config.Config
	Log *slog.Logger
	// *database.DB
	Server *http.Server
}

func NewApp() *App {
	var app App

	app.Config = config.MustLoad()
	app.Log = logger.SetupLogger(app.Env)
	// app.DB = database.ConnectToDatabase()
	app.Server = &http.Server{
		Addr:         app.Config.Address,
		Handler:      server.SetupRouter(app.Log),
		ReadTimeout:  app.Config.Timeout,
		WriteTimeout: app.Config.Timeout,
		IdleTimeout:  app.Config.IdleTimeout,
	}

	return &app

}
