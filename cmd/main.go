package main

import (
	"go-unotv-skeleton/cmd/router"
	"go-unotv-skeleton/pkg/application"
	"os"
	"unotv/go_utv_utils/pkg/exithandler"
	"unotv/go_utv_utils/pkg/logging"
	"unotv/go_utv_utils/pkg/server"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// @title
// @version 1.0
// @description

// @contact.name API Support
// @contact.url http://www.unotv.com
// @contact.email fernando.aviles@clarosports.com

// @host dev-go-utv-name.utv-infra.com
// @BasePath /
// @schemes http
func main() {
	//Cargamos archivo env
	if err := godotenv.Load("/var/dev-repos/properties/unotv.env", ".env"); err != nil {
		log.Error().Msg("Failed to load env vars")
	}

	//Configuramos log
	logger := logging.Configure(logging.Config{
		ConsoleLoggingEnabled: true,
		EncodeLogsAsJson:      true,
		FileLoggingEnabled:    true,
		Directory:             os.Getenv("LOG_DIRECTORY"),
		Filename:              os.Getenv("LOG_FILENAME"),
		MaxSize:               1,
		MaxBackups:            2,
		MaxAge:                15})
	log.Logger = *logger.Logger

	//******
	app, err := application.Get()
	if err != nil {
		log.Error().Msg("Conf error: " + err.Error())
	}

	//***** Servidor
	srv := server.
		Get().
		WithAddr(app.Cfg.GetAPIPort()).
		WithRouter(router.Get(app))
	go func() {
		log.Info().Msg("starting server at " + app.Cfg.GetAPIPort())
		if err := srv.Start(); err != nil {
			log.Error().Msg("Server error: " + err.Error())
		}
	}()

	//###
	exithandler.Init(func() {
		if err := srv.Close(); err != nil {
			log.Error().Msg("Exithandler error: " + err.Error())
		}
	})

}
