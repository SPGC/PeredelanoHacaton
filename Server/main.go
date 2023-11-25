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

	//c := cors.New(cors.Options{
	//	AllowedOrigins: []string{"http://localhost:5173"}, // your front-end origin
	//	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	//})

	//defaultHandler := func(w http.ResponseWriter, r *http.Request) {
	//	//println(r.Method)
	//	url := r.URL.Path
	//	switch r.Method {
	//	case "GET":
	//		switch getEntity(url) {
	//		case "user":
	//			switch getGETRequestType(url) {
	//			case GetById:
	//
	//			case GetAllWhereParam:
	//
	//			}
	//		case "issue":
	//
	//		case "organisation":
	//
	//		case "message":
	//
	//		}
	//		Handlers.GetHandler(w, r, nil)
	//	case "POST":
	//		Handlers.PostHandler(w, r, nil)
	//	}
	//
	//}

	defaultHandler := mux.NewRouter()

	defaultHandler.HandleFunc("/user/{id:[0-9]+}", Handlers.GetUserByIdHandler)
	defaultHandler.HandleFunc("/user", Handlers.GetAllUsersWhereParam)

	//handlerWithCors := c.Handler(http.HandlerFunc(defaultHandler))

	defaultHandler.Methods("GET", "POST", "PUT", "DELETE", "OPTIONS").
		Headers("Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization")

	//err := http.ListenAndServe(":8080", handlerWithCors)
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
