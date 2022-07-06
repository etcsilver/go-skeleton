package confwf

import (
	"embed"
	"flag"
)

var config embed.FS

type Conf struct {
	//PathFileParalimpicos         string
	//DominioParalimpicos          string

}

func Get() *Conf {
	conf := &Conf{}
	//flag.StringVar(&conf.PathFileParalimpicos, "Path Paralimpicos", os.Getenv("PAR_PATH_FILES"), "Path Paralimpicos")
	//flag.StringVar(&conf.DominioParalimpicos, "Dominio Paralimpicos", os.Getenv("PAR_DOMINIO"), "Dominio Paralimpicos")

	flag.Parse()
	return conf
}
