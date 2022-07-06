package ws

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"unotv/go_utv_push/domain"

	"github.com/rs/zerolog/log"
)

//
func (ws *webservices) AkamaiPurgeUrls(akamai *domain.Akamai) (string, error) {
	log.Debug().Msg("-- Inicia AkamaiPurgeUrls [Akamai] --")
	url := ws.app.Cfg.ApiGoAkamai + "/purge_urls"
	log.Debug().Msg(url)
	jsonReq, err := json.Marshal(akamai)
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Error().Msg("Error AkamaiPurgaUrls: " + err.Error())
		return "", errors.New(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	log.Debug().Msg("API Response:" + resp.Status)
	return bodyString, nil

}
