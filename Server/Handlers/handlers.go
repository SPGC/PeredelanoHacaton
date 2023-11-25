package Handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var (
	ParseIntError = errors.New("can't parse int")
	BadRequest    = errors.New("unsupported request")
)

const (
	getAll      = "getAll"
	getById     = "getById"
	unsupported = "unsupported"
)

func ReadUrlGet(query *url.URL) (string, string, int, url.Values, error) {

	path := query.Path
	tokenizedPath := strings.Split(path, "/")
	var entity string
	id := -1
	var err error = nil
	if len(tokenizedPath) == 2 {
		entity = strings.Split(tokenizedPath[1], "?")[0]
		return getAll, entity, id, query.Query(), err
	} else if len(tokenizedPath) == 3 {
		entity = tokenizedPath[1]
		idToken := strings.Split(tokenizedPath[2], "?")[0]
		id, err = strconv.Atoi(idToken)
		if err != nil {
			return unsupported, "", id, nil, ParseIntError
		}
		return getById, entity, id, query.Query(), err
	} else {
		return unsupported, "", id, nil, BadRequest
	}
}

func ReadUrlPost(query *url.URL) (string, error) {

	path := query.Path
	tokenizedPath := strings.Split(path, "/")
	var entity string
	var err error = nil
	if len(tokenizedPath) == 3 {
		entity = tokenizedPath[1]
		return entity, err
	} else {
		return "", BadRequest
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request, dataBase *sql.DB) {
	requestType, entity, id, query, err := ReadUrlGet(r.URL)

	if err != nil {
		if err == ParseIntError {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else if err == BadRequest {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	var sqlQuery string
	if requestType == getAll {
		page, err := strconv.Atoi(query.Get("page"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		limit, err := strconv.Atoi(query.Get("limit"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		sqlQuery = fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", entity, limit, limit*(page-1))
	} else if requestType == getById {
		sqlQuery = fmt.Sprintf("SELECT * FROM %s WHERE id = %d", entity, id)

	}
	println(sqlQuery)
}

//func PostHandler(w http.ResponseWriter, r *http.Request, dataBase *sql.DB) {
//	entity, err := ReadUrlPost(r.URL)
//
//	if err != nil {
//		if err == ParseIntError {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//		} else if err == BadRequest {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//		} else {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//		}
//	}
//
//	var sqlQuery string
//	page, err := strconv.Atoi(query.Get("page"))
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//	}
//	limit, err := strconv.Atoi(query.Get("limit"))
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//	}
//	sqlQuery = fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", entity, limit, limit*(page-1))
//
//	println(sqlQuery)
//}
