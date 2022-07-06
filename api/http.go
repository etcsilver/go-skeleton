package api

import (
	"net/http"
	"olimpicos/go_par_workflow/domain"

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
// @Summary      Publica nota.
// @Description  Publica nota en el webserver.
// @Tags         workflow
// @Param        requestNota   body domain.RequestNota  true  "id nota"
// @Success      200      {object} domain.SuccessRespose "Nota Publicada"
// @Failure      500      {object} domain.ErrorRespose "Error publicando la nota"
// @Router       /publica_nota [post]
func (h *handler) PublicaNota(c *gin.Context) {
	log.Debug().Msg("--== Inicia PublicaNota ==--")
	idNota := c.Param("idNota")
	log.Debug().Msg("Par idNota: " + idNota)
	err := h.service.PublicaNota(idNota)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Error publicando la nota: " + err.Error()})
		return
	}
	log.Debug().Msg("--== FIN PublicaNota ==--")
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Nota Publicada"})
}

// Borra notas
// @Summary      Borra nota.
// @Description  Metodo que borra una nota del webserver.
// @Tags         workflow
// @Param        requestNota   body domain.RequestNota  true  "id nota"
// @Success      200     {object} domain.SuccessRespose
// @Failure      500     {object} domain.ErrorRespose
// @Router       /borra_nota [post]
func (h *handler) PublicaHomes(c *gin.Context) {
	log.Debug().Msg("--== Inicia PublicaHomes ==--")

	requestHomes := &domain.Homes{}
	if err := c.ShouldBindJSON(&requestHomes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Error json" + err.Error()})
		return
	}
	err := h.service.PublicaHomes(requestHomes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Error publicando homes: " + err.Error()})
		return
	}
	log.Debug().Msg("--== FIN PublicaHomes ==--")
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Home Publicado"})
}
