package main

import (
	"encoding/json"
	"net/http"
)

func isTokenValid(w http.ResponseWriter, r *http.Request) {
	params := getParams(r)

	token := params["token"]

	if jwt, ok := tokens[token]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(jwt)
		return
	}

	errorResponse("Provided token is not valid", http.StatusBadRequest, w)
}
