package application

import "github.com/etcsilver/go-skeleton.git/pkg/confapp"

type Application struct {
	CfgApp *confapp.Conf
}

func Get() (*Application, error) {

	cfgApp := confapp.Get()

	return &Application{

		CfgApp: cfgApp,
	}, nil
}
