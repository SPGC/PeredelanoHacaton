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

	defaultHandler.HandleFunc("/user/{id:[0-9]+}", Handlers.GetUserByIdHandler).Methods("GET")
	defaultHandler.HandleFunc("/user", Handlers.GetAllUsersWhereParam).Methods("GET")

	err := http.ListenAndServe(":8080", defaultHandler)
	if err != nil {
		panic(err)
	}
}

func main() {
	println("Server started")
	//connStr := "user=youruser dbname=yourdb password=yourpassword host=localhost sslmode=disable"
	//
	//// Установка соединения с базой данных
	//db, err := sql.Open("postgres", connStr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()
	//
	//// Проверка соединения с базой данных
	//err = db.Ping()
	//if err != nil {
	//	log.Fatal("Ошибка соединения:", err)
	//}
	RunServer()
}
