package main

import (
	"net/http"
)

func RunServer() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
		method, values := ReadUrl(r.URL)
		println(r.Method)
	}

	err := http.ListenAndServe(":8080", http.HandlerFunc(handler))
	if err != nil {
		panic(err)
	}
}

func main() {
	println("Server started")
	RunServer()
}
