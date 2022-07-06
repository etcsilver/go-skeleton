package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Get(app *application.Application) *gin.Engine {
	log.Info().Msg("Definiendo metodos...")
	if app.Cfg.Ambiente == "produccion" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()
	
	//Example handler
	//r.GET("/publica_nota/:idNota", func(c *gin.Context) { home.PublicaCalendario(app, c) })
	//Metrics
	//r.GET("/metrics", prometheusHandler())
	/*
		router.HandleFunc("/publica_files", file.PublicaFiles(app)).Methods("POST")
		router.HandleFunc("/borra_files", file.BorraFiles(app)).Methods("POST")
	*/
	return r
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}