package Handlers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ParseIntError = errors.New("can't parse int")
	BadRequest    = errors.New("unsupported request")
	BadBody       = errors.New("can't parse body")
)

const (
	getAll      = "getAll"
	getById     = "getById"
	unsupported = "unsupported"
	newIssue    = "newIssue"
)

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {

	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("SELECT * FROM users WHERE id = %d", ids["id"])
	w.Write([]byte(sqlQuery))

}

func GetAllUsersWhereParam(w http.ResponseWriter, r *http.Request) {

	//queryParams :=
	//	r.URL.Query()
	println(r.URL.String())

}

//func GetHandler(w http.ResponseWriter, r *http.Request, dataBase *sql.DB) {
//	requestType, entity, id, query, err := ReadUrlGet(r.URL)
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
//	if requestType == getAll {
//		page, err := strconv.Atoi(query.Get("page"))
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//		}
//		limit, err := strconv.Atoi(query.Get("limit"))
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//		}
//		sqlQuery = fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", entity, limit, limit*(page-1))
//	} else if requestType == getById {
//		sqlQuery = fmt.Sprintf("SELECT * FROM %s WHERE id = %d", entity, id)
//
//	}
//	//w.Write([]byte(sqlQuery))
//	println(sqlQuery)
//}
//
//func PostHandler(w http.ResponseWriter, r *http.Request, dataBase *sql.DB) {
//	entity, err := ReadUrlPost(r.URL)
//
//	if err != nil {
//		if err == BadRequest {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//		} else {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//		}
//	}
//
//	if entity == newIssue {
//		var issue Entities.IssueCreationWrapper
//		body := r.Body
//		data := make([]byte, 1024)
//		sb := strings.Builder{}
//		var readBytes int
//		var err error = nil
//		for true {
//			if err != nil {
//				break
//			}
//			readBytes, err = body.Read(data)
//			if readBytes == 0 {
//				continue
//			}
//			sb.Write(data[:readBytes])
//		}
//		bodyData := sb.String()
//		err = json.Unmarshal([]byte(bodyData), &issue)
//		if err != nil {
//			http.Error(w, BadBody.Error(), http.StatusUnprocessableEntity)
//			return
//		}
//		sqlQuery := fmt.Sprintf("PROCEDURE_NAME (%s, %s, %s, %s, %s, %s, %s)",
//			issue.Issuer.Name,
//			issue.Issuer.ContactInfo,
//			issue.IssueMessage.Description,
//			issue.ProblemCompany.Name,
//			issue.ProblemCompany.Country,
//			issue.ProblemCompany.ContactInfo,
//			issue.ProblemCompany.OrgType,
//		)
//
//		w.Write([]byte(sqlQuery))
//		//println(sqlQuery)
//	}
//
//	//var sqlQuery string
//	//page, err := strconv.Atoi(query.Get("page"))
//	//if err != nil {
//	//	http.Error(w, err.Error(), http.StatusBadRequest)
//	//}
//	//limit, err := strconv.Atoi(query.Get("limit"))
//	//if err != nil {
//	//	http.Error(w, err.Error(), http.StatusBadRequest)
//	//}
//	//sqlQuery = fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", entity, limit, limit*(page-1))
//	//
//	//println(sqlQuery)
//}
