package app

import (
	"log/slog"
	"net/http"
	"testp/internal/config"
	"testp/internal/database"
	"testp/internal/handlers"
	"testp/internal/logger"
)

type App struct {
	*config.Config
	Log *slog.Logger
	*database.DB
	Server *http.Server
}

func NewApp() *App {
	var app App

	app.Config = config.MustLoad()
	app.Log = logger.SetupLogger(app.Env)
	app.DB = database.ConnectToDatabase()
	app.Server = &http.Server{
		Addr:         app.Config.Address,
		Handler:      handlers.SetupRouter(),
		ReadTimeout:  app.Config.Timeout,
		WriteTimeout: app.Config.Timeout,
		IdleTimeout:  app.Config.IdleTimeout,
	}

	return &app

}
