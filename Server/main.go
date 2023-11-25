package main

import (
	"PeredelanoHakaton/Handlers"
	"net/http"
)

func RunServer() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong\n"))
		//method, values := Handlers.ReadUrlGet(r.URL)
		//println(r.Method)
		Handlers.GetHandler(w, r, nil)
	}

	err := http.ListenAndServe(":8080", http.HandlerFunc(handler))
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
