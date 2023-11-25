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

	router.HandleFunc("/user/{id:[0-9]+}", Handlers.GetUserByIdHandler).Methods("GET")
	router.HandleFunc("/user", Handlers.GetAllUsersWhereParam).Methods("GET")

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
