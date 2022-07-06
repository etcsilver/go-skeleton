package ws

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
	"source.marcaclaro.com/go_par_workflow/domain"
)

//
func (ws *webservices) RemplazaPaginaFindByTipo(tipoPagina string) ([]*domain.Replacement, error) {
	log.Debug().Msg("-- Inicia RemplazaPaginaFindByTipo [Mongo] --")

	url := ws.app.Cfg.ApiGoMongoDB + "/paralimpicos/replacement/findByTipo/" + tipoPagina
	log.Debug().Msg(url)

	resp, err := http.Get(url)
	if err != nil {
		log.Error().Msg("Error NotaSave: " + err.Error())
		return nil, errors.New(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)

	// Convert bodyString to json
	textBytes := []byte(bodyString)
	remplazaPagina := []*domain.Replacement{}
	json.Unmarshal(textBytes, &remplazaPagina)
	log.Debug().Msg("API Response:" + resp.Status)
	return remplazaPagina, nil

}

// Funcion que envia una nota para guarda en la base de datos
// Return string
func (ws *webservices) NotaSave(nota *domain.Nota) (string, error) {
	log.Debug().Msg("-- Inicia NotaSave [Mongo] --")

	url := ws.app.Cfg.ApiGoMongoDB + "/paralimpicos/nota/save"
	log.Debug().Msg(url)

	//Json a enviar
	jsonReq, err := json.Marshal(nota)
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Error().Msg("Error NotaSave: " + err.Error())
		return "", errors.New(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	log.Debug().Msg("API Response:" + resp.Status)
	return bodyString, nil
}

// Funcion que envia una nota para actualizar en la base de datos
// Return string
func (ws *webservices) NotaUpdate(nota *domain.Nota) (string, error) {
	log.Debug().Msg("-- Inicia NotaUpdate [Mongo] --")

	url := ws.app.Cfg.ApiGoMongoDB + "/paralimpicos/nota/update"
	log.Debug().Msg(url)

	jsonReq, err := json.Marshal(nota)
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))

	if err != nil {
		log.Error().Msg("Error NotaUpdate: " + err.Error())
		return "", errors.New(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	log.Debug().Msg("API Response:" + resp.Status)
	return bodyString, nil
}

// Funcion que envia una nota para actualizar en la base de datos
// Return string
func (ws *webservices) NotaUpsert(nota *domain.Nota) (string, error) {
	log.Debug().Msg("-- Inicia NotaUpsert [Mongo] --")

	url := ws.app.Cfg.ApiGoMongoDB + "/paralimpicos/nota/upsert"
	log.Debug().Msg(url)

	jsonReq, err := json.Marshal(nota)
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Error().Msg("Error NotaUpdate: " + err.Error())
		return "", errors.New(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	log.Debug().Msg("API Response:" + resp.Status)
	return bodyString, nil
}

// Funcion que borra una nota de la base de datos.
func (ws *webservices) NotaDelete(nota *domain.Nota) (string, error) {
	log.Debug().Msg("-- Inicia NotaDelete [Mongo] --")

	url := ws.app.Cfg.ApiGoMongoDB + "/paralimpicos/nota/delete"
	log.Debug().Msg(url)
	jsonReq, err := json.Marshal(nota)
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Error().Msg("Error NotaDelete: " + err.Error())
		return "", errors.New(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	log.Debug().Msg("API Response:" + resp.Status)
	return bodyString, nil
}

//
func (ws *webservices) NotaLoad(idNota string) (*domain.Nota, error) {
	log.Debug().Msg("-- Inicia NotaLoad [Mongo] --")

	url := ws.app.Cfg.ApiGoMongoDB + "/paralimpicos/nota/load/" + idNota
	log.Debug().Msg(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Msg("Error NotaLoad: " + err.Error())
		return nil, errors.New(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	log.Debug().Msg("API Response:" + resp.Status)

	// Convert response body to string
	bodyString := string(bodyBytes)

	// Convert bodyString to json
	textBytes := []byte(bodyString)
	nota := &domain.Nota{}
	json.Unmarshal(textBytes, &nota)
	return nota, nil
}
