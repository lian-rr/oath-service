package main

import (
	"fmt"

	"github.com/lian-rr/oath-service/model"
)

//Temporary DB xD
var users map[string]model.User

var tokens map[string]JWT

func init() {
	users = make(map[string]model.User)
	tokens = make(map[string]JWT)
}

func main() {
	loadUserDB()

	startServer()
}

func loadUserDB() {
	fmt.Println("Loading user data")
	users["1"] = model.User{Id: "1", Email: "test", Password: "1234"}
}

func saveToken(jwt JWT) {
	tokens[jwt.Token] = jwt
}
