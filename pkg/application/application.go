package application

import (
	"go_strtucut/pkg/confwf"
	"unotv/go_utv_utils/pkg/config"
)

type Application struct {
	Cfg   *config.Config
	CfgWF *confwf.Conf
}

func Get() (*Application, error) {
	cfg := config.Get()
	cfgWF := confwf.Get()

	return &Application{
		Cfg:   cfg,
		CfgWF: cfgWF,
	}, nil
}
