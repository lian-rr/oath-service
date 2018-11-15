package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func validate(w http.ResponseWriter, r *http.Request) {
	params := getParams(r)

	token := params["token"]

	if jwt, ok := tokens[token]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(jwt)
		return
	}

	errorResponse("Provided token is not valid", http.StatusBadRequest, w)
}

func invalidate(w http.ResponseWriter, r *http.Request) {
	params := getParams(r)

	token := params["token"]
	fmt.Printf("Deleting token with value: %s \n", token)
	delete(tokens, token)
	w.WriteHeader(http.StatusOK)
}
