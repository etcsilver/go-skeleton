package domain

import "go-skeleton/pkg/application"

type Service interface {
	//PublicaNota(idNota string) error
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
	repoMongo   Repository
}

//NewWorkflowService...
func NewWorkflowService(app *application.Application, webservices WebServices, repository Repository) Service {
	return &service{app: app, webservices: webservices, repoMongo: repository}
}
