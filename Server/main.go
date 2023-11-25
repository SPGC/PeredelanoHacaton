package main

import (
	"PeredelanoHakaton/Handlers"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	router := mux.NewRouter()

	router.HandleFunc("/users/{id:[0-9]+}", Handlers.GetUserById).Methods("GET")
	router.HandleFunc("/users", Handlers.GetAllUsersWhereParam).Methods("GET")
	router.HandleFunc("/organisations/{id:[0-9]+}", Handlers.GetOrganisationById).Methods("GET")
	router.HandleFunc("/organisations", Handlers.GetAllOrganisationWhereParam).Methods("GET")
	router.HandleFunc("/issues/{id:[0-9]+}", Handlers.GetIssueById).Methods("GET")
	router.HandleFunc("/issues", Handlers.GetAllIssuesWhereParam).Methods("GET")
	router.HandleFunc("/messages/{id:[0-9]+}", Handlers.GetMessageById).Methods("GET")
	router.HandleFunc("/messages", Handlers.GetAllMessagesWhereParam).Methods("GET")

	router.HandleFunc("/issues", Handlers.PostIssue).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // Frontend origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	handlerWithCors := c.Handler(router)

	err := http.ListenAndServe(":8080", handlerWithCors)

	if err != nil {
		panic(err)
	}
}

func main() {
	println("Server started")
	RunServer()
}
