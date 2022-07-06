package router

import (
	"github.com/etcsilver/go-skeleton.git/pkg/application"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Get(app *application.Application) *gin.Engine {
	log.Info().Msg("Definiendo metodos...")
	if app.CfgApp.Ambiente == "produccion" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()

	//Example handler
	r.GET("/hello/:name", func(c *gin.Context) { handler.Hello(app, c) })
	return r
}
