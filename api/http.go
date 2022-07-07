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
	requestHello := &domain.Hello{}
	if err := c.ShouldBindJSON(&requestHello); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorRespose{Code: http.StatusInternalServerError, Message: "Error json: " + err.Error()})
		return
	}

	saludo, err := h.service.Hello(requestHello.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Error: " + err.Error()})
		return
	}
	log.Debug().Msg("--== FIN Hello ==--")
	c.JSON(http.StatusOK, domain.SuccessRespose{Code: http.StatusOK, Message: saludo})
}
