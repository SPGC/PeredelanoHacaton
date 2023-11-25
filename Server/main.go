package main

import (
	"PeredelanoHakaton/Handlers"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	GetAllWhereParam = 1
	GetById          = 2
)

func getEntity(url string) string {
	entityToken := strings.Split(url, "/")[1]
	entity := strings.Split(entityToken, "?")[0]
	return entity
}

func getGETRequestType(url string) int {
	if len(strings.Split(url, "/")) == 3 {
		return GetById
	}
	return GetAllWhereParam
}

func RunServer() {

	defaultHandler := mux.NewRouter()

	defaultHandler.HandleFunc("/users/{id:[0-9]+}", Handlers.GetUserByIdHandler).Methods("GET")
	defaultHandler.HandleFunc("/users", Handlers.GetAllUsersWhereParam).Methods("GET")

	err := http.ListenAndServe(":8080", defaultHandler)
	if err != nil {
		panic(err)
	}
}

func main() {
	println("Server started")
	RunServer()
}
