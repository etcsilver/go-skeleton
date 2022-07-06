package confapp

import (
	"embed"
	"flag"
	"os"
)

var config embed.FS

type Conf struct {
	Ambiente string
}

func Get() *Conf {
	conf := &Conf{}
	flag.StringVar(&conf.Ambiente, "Ambiente", os.Getenv("AMBIENTE"), "Ambiente")

	flag.Parse()
	return conf
}
