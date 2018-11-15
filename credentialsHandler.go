package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lian-rr/oath-service/model"
)

func authenticate(w http.ResponseWriter, r *http.Request) {

	var credentials UserCredentials

	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		m := "Not possible to decode the provided payload"
		fmt.Printf(m)
		errorResponse(m, http.StatusBadRequest, w)
		return
	}

	cUser := cCredentials(credentials)

	uD, err := uDataByEmail(cUser.Email)
	if err != nil {
		m := err.Error()
		fmt.Printf(m)
		errorResponse(m, http.StatusNoContent, w)
		return

	}

	if vCredentials(cUser, uD) {
		jwt := JWT{Token: generateToken(uD)}
		saveToken(jwt)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(jwt)
		return
	}

	errorResponse("Credentials not valid", http.StatusUnauthorized, w)

}

func vCredentials(u UserCredentials, uD model.User) bool {
	if len(u.Email) == 0 || len(u.Password) == 0 {
		return false
	}

	return u.Email == uD.Email && encriptP(u.Password) == uD.Password
}

//Should clean crendentials from strange characters
func cCredentials(u UserCredentials) UserCredentials {
	return UserCredentials{Email: u.Email, Password: u.Password}
}

//Should encript the password
func encriptP(p string) string {
	return p
}
