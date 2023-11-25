package main

import (
	"PeredelanoHakaton/Handlers"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func RunServer() {
	router := mux.NewRouter()

	dbConnection, err := sql.Open("postgres", "user=postgres dbname=gerahelperdb password=12345678 host=localhost sslmode=disable")
	if err != nil {
		println("Can't access data base")
	}
	defer func(dbConnection *sql.DB) {
		err := dbConnection.Close()
		if err != nil {
			println("Error in closing DB connection")
		}
	}(dbConnection)
	err = dbConnection.Ping()
	if err != nil {
		println("Can't access data base")
	}

	db := Handlers.DBWrapper{Db: dbConnection}

	router.HandleFunc("/", Handlers.Ping).Methods("GET")

	router.HandleFunc("/users/{id:[0-9]+}", db.GetUserById).Methods("GET")
	router.HandleFunc("/organisations/{id:[0-9]+}", Handlers.GetOrganisationById).Methods("GET")
	router.HandleFunc("/issues/{id:[0-9]+}", Handlers.GetIssueById).Methods("GET")
	router.HandleFunc("/messages/{id:[0-9]+}", Handlers.GetMessageById).Methods("GET")

	router.HandleFunc("/users", db.GetAllUsersWhereParam).Methods("GET")
	router.HandleFunc("/organisations", Handlers.GetAllOrganisationWhereParam).Methods("GET")
	router.HandleFunc("/issues", Handlers.GetAllIssuesWhereParam).Methods("GET")
	router.HandleFunc("/messages", Handlers.GetAllMessagesWhereParam).Methods("GET")

	router.HandleFunc("/issues", Handlers.PostIssue).Methods("POST")
	router.HandleFunc("/organisations", Handlers.PostOrganisation).Methods("POST")
	router.HandleFunc("/messages", Handlers.PostMessage).Methods("POST")
	router.HandleFunc("/users", db.PostUser).Methods("POST")

	router.HandleFunc("/users/{id:[0-9]+}", db.DeleteUserById).Methods("DELETE")
	router.HandleFunc("/organisations/{id:[0-9]+}", Handlers.DeleteOrganisationById).Methods("DELETE")
	router.HandleFunc("/issues/{id:[0-9]+}", Handlers.DeleteIssueById).Methods("DELETE")
	router.HandleFunc("/messages/{id:[0-9]+}", Handlers.DeleteMessageById).Methods("DELETE")

	router.HandleFunc("/issues", Handlers.UpdateIssue).Methods("PUT")
	router.HandleFunc("/organisations", Handlers.UpdateOrganisation).Methods("PUT")
	router.HandleFunc("/messages", Handlers.UpdateMessage).Methods("PUT")
	router.HandleFunc("/users", db.UpdateUser).Methods("PUT")

	c := cors.New(cors.Options{
		//AllowedOrigins: []string{"http://localhost:5173"}, // Frontend origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	handlerWithCors := c.Handler(router)

	err = http.ListenAndServe(":8080", handlerWithCors)

	if err != nil {
		panic(err)
	}
}

func main() {
	println("Server started")
	RunServer()
}
