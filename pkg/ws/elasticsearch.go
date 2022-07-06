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

// Funcion que indexa una nota en el elasticsearch
//
func (ws *webservices) ElasticIndexNota(nota *domain.Nota) (string, error) {
	log.Debug().Msg("-- Inicia ElasticIndexNota [Elasticsearch] --")
	url := ws.app.Cfg.ApiGoElastic + "/paralimpicos/index_nota"
	log.Debug().Msg(url)
	jsonReq, err := json.Marshal(nota)
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Error().Msg("Error ElasticIndexNota: " + err.Error())
		return "", errors.New(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	log.Debug().Msg("API Response:" + resp.Status)
	return bodyString, nil
}

// Funcion que borra una nota del elastic search
//
func (ws *webservices) ElasticDeleteNota(idNota string) (string, error) {
	log.Debug().Msg("-- Inicia ElasticDeleteNota [Elasticsearch] --")
	url := ws.app.Cfg.ApiGoElastic + "/paralimpicos/delete_nota/" + idNota
	log.Debug().Msg(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Msg("Error ElasticDeleteNota: " + err.Error())
		return "", errors.New(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	log.Debug().Msg("API Response:" + resp.Status)
	return bodyString, nil
}
