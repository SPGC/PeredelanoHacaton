package Handlers

import (
	"net/http"
	"net/url"
	"strings"
)

func ReadUrl(query *url.URL) (string, url.Values) {

	method := query.Path
	method = strings.Split(method, "/")[1]

	return method, query.Query()
}

func getHandler(w http.ResponseWriter, r *http.Request) {

}
