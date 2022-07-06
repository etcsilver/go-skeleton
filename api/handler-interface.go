package api

import "github.com/gin-gonic/gin"

type WorkflowHandler interface {
	Hello(c *gin.Context)
}
