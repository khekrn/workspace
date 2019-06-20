package handlers

import "net/http"

// TokenizeTextHandler will tokenize the text
func TokenizeTextHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
}
