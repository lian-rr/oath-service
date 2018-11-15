package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lian-rr/oath-service/model"
)

func getUser(w http.ResponseWriter, r *http.Request) {

	params := getParams(r)

	id := params["id"]

	fmt.Printf("Getting data for user with id: %s\n", id)

	user, err := uData(id)

	if err != nil {
		m := fmt.Sprintf("User with id: %s not found.", id)

		fmt.Printf(m)
		errorResponse(m, http.StatusNoContent, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapUserResp(user))
}

func mapUserResp(u model.User) UserResponse {
	return UserResponse{Id: u.Id, Email: u.Email}
}
