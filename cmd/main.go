package main

import (
	"os"

	"github.com/etcsilver/go-skeleton.git/cmd/router"
	"github.com/etcsilver/go-skeleton.git/pkg/application"
	"github.com/etcsilver/go-skeleton.git/pkg/exithandler"
	"github.com/etcsilver/go-skeleton.git/pkg/logging"
	"github.com/etcsilver/go-skeleton.git/pkg/server"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// @title
// @version 1.0
// @description

// @contact.name API Support
// @contact.url
// @contact.email dilver@gmail.com

// @host
// @BasePath /
// @schemes http
func main() {
	//Cargamos archivo env
	if err := godotenv.Load(".env"); err != nil {
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
		WithAddr(":8080").
		WithRouter(router.Get(app))
	go func() {
		log.Info().Msg("starting server ")
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
