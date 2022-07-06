package domain

import "github.com/etcsilver/go-skeleton.git/pkg/application"

type Service interface {
	Hello(name string) (error, string)
}

type Repository interface {
	//NotaUpdate(selector interface{}, update interface{}) error

}

//
type WebServices interface {
}

type service struct {
	app         *application.Application
	webservices WebServices
}

//NewWorkflowService...
func NewWorkflowService(app *application.Application, webservices WebServices) Service {
	return &service{app: app, webservices: webservices}
}

func (s *service) Hello(name string) (error, string) {
	return nil, "Hola " + name
}
