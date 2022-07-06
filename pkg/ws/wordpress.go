package ws

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"source.marcaclaro.com/go_par_workflow/domain"
)

//
func (ws *webservices) GetInfoNota(idNota string) (*domain.Nota, error) {
	log.Debug().Msg("--- Inicia GetInfoNota")
	log.Debug().Msg("idNota: " + idNota)

	api := strings.ReplaceAll(ws.app.CfgWF.UrlApiWordpressParalimpicos, "&ID_NOTA&", idNota)
	log.Debug().Msg(api)
	url_api := ws.app.CfgWF.DominioWordpressParalimpicos + api
	log.Debug().Msg("URL: " + url_api)

	u, err := url.Parse(url_api)
	if err != nil {
		log.Error().Msg("Error en parse URL: " + err.Error())
		return nil, errors.New(err.Error())
	}

	text, err := getUrlContent(u.String())
	if err != nil {
		log.Error().Msg("Error getUrlContent: " + err.Error())
		return nil, errors.New(err.Error())
	}

	if text == "404" {
		log.Error().Msg(text + " Nota no encotrada")
		return nil, errors.New(text + " Nota no encotrada")
	}

	textBytes := []byte(text)
	nota := &domain.Nota{}
	err = json.Unmarshal(textBytes, &nota)
	if err != nil {
		log.Error().Msg("Error al parsear el json: " + err.Error())
		return nil, errors.New(err.Error())
	}
	return nota, nil
}

func getUrlContent(urlToGet string) (string, error) {
	var (
		err     error
		content []byte
		resp    *http.Response
	)

	// GET content of URL
	if resp, err = http.Get(urlToGet); err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if request was successful
	if resp.StatusCode != 200 {
		return strconv.Itoa(resp.StatusCode), err
	}

	// Read the body of the HTTP response
	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		return "", err
	}

	return string(content), err
}
