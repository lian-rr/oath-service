package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/lian-rr/oath-service/model"
)

type (
	JWT struct {
		Token string `json:"token"`
	}

	UserCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserResponse struct {
		Id    string `json:"id"`
		Email string `json:"email"`
	}

	ErrorResponse struct {
		Message string `json:"message"`
	}
)

func startServer() {
	router := mux.NewRouter()

	startHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func startHandlers(r *mux.Router) {
	fmt.Println("Starting handlers")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/tokens", authenticate).Methods("POST")
	r.HandleFunc("/tokens/{token}", validate).Methods("GET")
	r.HandleFunc("/tokens/{token}", invalidate).Methods("DELETE")
}

func errorResponse(m string, s int, w http.ResponseWriter) {
	w.WriteHeader(s)
	json.NewEncoder(w).Encode(ErrorResponse{Message: m})
}

func getParams(r *http.Request) map[string]string {
	return mux.Vars(r)
}

/*
* Generate token and save it
 */
func generateToken(u model.User) string {
	now := time.Now().Unix()
	return fmt.Sprintf("%s%x", u.Id, now)
}

/*
* Move to the data access layer
 */
func uData(id string) (model.User, error) {

	if user, ok := users[id]; ok {
		return user, nil
	}
	return model.User{}, errors.New("User not found")
}

func uDataByEmail(email string) (model.User, error) {
	for _, u := range users {
		if u.Email == email {
			return u, nil
		}
	}
	return model.User{}, errors.New("User not found")
}
