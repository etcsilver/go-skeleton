package ws

import (
	"go-unotv-skeleton/domain"
	"go-unotv-skeleton/pkg/application"
)

type webservices struct {
	app *application.Application
}

//NewWorkflowService...
func NewWebServices(app *application.Application) domain.WebServices {
	return &webservices{app: app}
}
