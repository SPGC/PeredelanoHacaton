package main

import (
	"PeredelanoHakaton/Handlers"
	"net/http"

	"github.com/rs/cors"
)

func RunServer() {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // your front-end origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	defaultHandler := func(w http.ResponseWriter, r *http.Request) {
		//println(r.Method)
		switch r.Method {
		case "GET":
			Handlers.GetHandler(w, r, nil)
		case "POST":
			Handlers.PostHandler(w, r, nil)
		}

	}

	handlerWithCors := c.Handler(http.HandlerFunc(defaultHandler))

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
