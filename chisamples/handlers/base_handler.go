package handlers

import (
	"chisamples/bl"
	"chisamples/reqs"
	"encoding/json"
	"net/http"
)

// TokenizeTextHandler will tokenize the text
func TokenizeTextHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var tokenRequest reqs.TokenizerRequest

	if err := json.NewDecoder(r.Body).Decode(&tokenRequest); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tokens := bl.TokenizeText(tokenRequest.Text)

	respondWithJSON(w, http.StatusOK, tokens)
}

// WordCountHandler computes the wordcount for the text
func WordCountHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var wordCountRequest reqs.WordCountRequest

	if err := json.NewDecoder(r.Body).Decode(&wordCountRequest); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	wordCount := bl.WordCount(wordCountRequest.Text)

	respondWithJSON(w, http.StatusOK, wordCount)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
