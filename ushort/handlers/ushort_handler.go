package handlers

import (
	"encoding/json"
	"net/http"

	"ushort/bl"
	"ushort/responses"

	log "github.com/sirupsen/logrus"
)

// Encode encodes url to short url
func Encode() func(resp http.ResponseWriter, req *http.Request) {
	log.Info("in Encode method")
	return func(resp http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			log.Fatal("error while parsing request - ", err.Error())
			resp.WriteHeader(503)
			resp.Write([]byte("error while parsing request - " + err.Error()))
			return
		}

		inputURL := req.FormValue("url")

		result, err := bl.EncodeToShortUrl(inputURL)

		if err != nil {
			resp.WriteHeader(503)
			resp.Write([]byte("cannot encode the url"))
			return
		}

		er := responses.NewEncoderResponse(inputURL, result)

		resp.Header().Set("Content-Type", "application/json")
		json.NewEncoder(resp).Encode(er)

		log.Info("return from Encode method")
	}
}

// Decode decodes the short url to long url
func Decode() func(resp http.ResponseWriter, req *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {

	}
}
