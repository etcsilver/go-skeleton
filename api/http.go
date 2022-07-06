package api

import (
	"net/http"

	"github.com/etcsilver/go-skeleton.git/domain"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type handler struct {
	service domain.Service
}

//NewHandler ...
func NewHandler(service domain.Service) WorkflowHandler {

	return &handler{service: service}

}

// Publica Nota godoc
// @Summary
// @Description
// @Tags
// @Param        request   body domain.RequestNota  true  "id nota"
// @Success      200      {object} domain.SuccessRespose "Nota Publicada"
// @Failure      500      {object} domain.ErrorRespose "Error publicando la nota"
// @Router       /hello [get]
func (h *handler) Hello(c *gin.Context) {
	log.Debug().Msg("--== Inicia Hello ==--")
	name := c.Param("name")
	log.Debug().Msg("name: " + name)
	err, saludo := h.service.Hello(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Error: " + err.Error()})
		return
	}
	log.Debug().Msg("--== FIN Hello ==--")
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": saludo})
}
