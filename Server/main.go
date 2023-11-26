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
	router.HandleFunc("/", Handlers.Ping).Methods("OPTIONS")

	router.HandleFunc("/users/{id:[0-9]+}", db.GetUserById).Methods("GET")
	router.HandleFunc("/organisations/{id:[0-9]+}", db.GetOrganisationById).Methods("GET")
	router.HandleFunc("/issues/{id:[0-9]+}", db.GetIssueById).Methods("GET")
	router.HandleFunc("/messages/{id:[0-9]+}", db.GetMessageById).Methods("GET")

	router.HandleFunc("/users", db.GetAllUsersWhereParam).Methods("GET")
	router.HandleFunc("/organisations", db.GetAllOrganisationWhereParam).Methods("GET")
	router.HandleFunc("/issues", db.GetAllIssuesWhereParam).Methods("GET")
	router.HandleFunc("/messages", db.GetAllMessagesWhereParam).Methods("GET")

	router.HandleFunc("/issues", db.PostIssue).Methods("POST")
	router.HandleFunc("/organisations", db.PostOrganisation).Methods("POST")
	router.HandleFunc("/messages", db.PostMessage).Methods("POST")
	router.HandleFunc("/users", db.PostUser).Methods("POST")

	router.HandleFunc("/users/{id:[0-9]+}", db.DeleteUserById).Methods("DELETE")
	router.HandleFunc("/organisations/{id:[0-9]+}", db.DeleteOrganisationById).Methods("DELETE")
	router.HandleFunc("/issues/{id:[0-9]+}", db.DeleteIssueById).Methods("DELETE")
	router.HandleFunc("/messages/{id:[0-9]+}", db.DeleteMessageById).Methods("DELETE")

	router.HandleFunc("/issues", db.UpdateIssue).Methods("PUT")
	router.HandleFunc("/organisations", db.UpdateOrganisation).Methods("PUT")
	router.HandleFunc("/messages", db.UpdateMessage).Methods("PUT")
	router.HandleFunc("/users", db.UpdateUser).Methods("PUT")

	c := cors.New(cors.Options{
		//AllowedOrigins: []string{"http://localhost:5173"}, // Frontend origin
		//AllowedOrigins: []string{"*"}, // Frontend origin
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
