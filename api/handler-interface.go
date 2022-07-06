package api

import "github.com/gin-gonic/gin"

type WorkflowHandler interface {
	PublicaNota(c *gin.Context)
	BorraNota(c *gin.Context)	
}
