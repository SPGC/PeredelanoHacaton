package Handlers

import (
	"PeredelanoHakaton/Entities"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
	"strings"
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

func readBody(body io.ReadCloser) string {
	data := make([]byte, 1024)
	sb := strings.Builder{}
	var readBytes int
	var err error = nil
	for true {
		if err != nil {
			break
		}
		readBytes, err = body.Read(data)
		if readBytes == 0 {
			continue
		}
		sb.Write(data[:readBytes])
	}
	return sb.String()
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("SELECT * FROM users WHERE id = %s", ids["id"])
	w.Write([]byte(sqlQuery))

}

func GetAllUsersWhereParam(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	page, err := strconv.Atoi(queryParams.Get("page"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	limit, err := strconv.Atoi(queryParams.Get("limit"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sqlQuery := fmt.Sprintf("SELECT * FROM users LIMIT %d OFFSET %d", limit, limit*(page-1))
	println(r.URL.String())
	data := make([]Entities.User, limit)
	for i := 0; i < limit; i++ {
		data[i] = Entities.User{
			Id:          i,
			Name:        fmt.Sprintf("Foo %d", i),
			ContactInfo: "placeholder@mailmail.com",
		}
	}

	marshaled, err := json.Marshal(data)
	response := fmt.Sprintf("{meta: {size: %d}, data: %s}", 239, string(marshaled))
	w.Write([]byte(response))
	println(sqlQuery)
}

func GetOrganisationById(w http.ResponseWriter, r *http.Request) {

	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("SELECT * FROM organisations WHERE id = %s", ids["id"])
	w.Write([]byte(sqlQuery))

}

func GetAllOrganisationWhereParam(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	page, err := strconv.Atoi(queryParams.Get("page"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	limit, err := strconv.Atoi(queryParams.Get("limit"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sqlQuery := fmt.Sprintf("SELECT * FROM organisation LIMIT %d OFFSET %d", limit, limit*(page-1))
	println(r.URL.String())
	data := make([]Entities.Organisation, limit)
	for i := 0; i < limit; i++ {
		data[i] = Entities.Organisation{
			Id:          i,
			Name:        fmt.Sprintf("Foo %d", i),
			ContactInfo: "placeholder@mailmail.com",
			Country:     fmt.Sprintf("Country %d", i),
			OrgType:     fmt.Sprintf("Type %d", i),
		}
	}

	marshaled, err := json.Marshal(data)
	response := fmt.Sprintf("{meta: {size: %d}, data: %s}", 239, string(marshaled))
	w.Write([]byte(response))
	println(sqlQuery)
}

func GetIssueById(w http.ResponseWriter, r *http.Request) {

	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("SELECT * FROM issues WHERE id = %s", ids["id"])
	w.Write([]byte(sqlQuery))

}

func GetAllIssuesWhereParam(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	page, err := strconv.Atoi(queryParams.Get("page"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	limit, err := strconv.Atoi(queryParams.Get("limit"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sqlQuery := fmt.Sprintf("SELECT * FROM issues LIMIT %d OFFSET %d", limit, limit*(page-1))
	println(r.URL.String())
	data := make([]Entities.Issue, limit)
	for i := 0; i < limit; i++ {
		data[i] = Entities.Issue{
			Id:             i,
			Status:         fmt.Sprintf("Foo %d", i),
			Description:    fmt.Sprintf("Description %d", i),
			OrganisationId: i,
		}
	}

	marshaled, err := json.Marshal(data)
	response := fmt.Sprintf("{meta: {size: %d}, data: %s}", 239, string(marshaled))
	w.Write([]byte(response))
	println(sqlQuery)
}

func GetMessageById(w http.ResponseWriter, r *http.Request) {

	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("SELECT * FROM messages WHERE id = %d", ids["id"])
	w.Write([]byte(sqlQuery))

}

func GetAllMessagesWhereParam(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	page, err := strconv.Atoi(queryParams.Get("page"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	limit, err := strconv.Atoi(queryParams.Get("limit"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sqlQuery := fmt.Sprintf("SELECT * FROM messages LIMIT %d OFFSET %d", limit, limit*(page-1))
	println(r.URL.String())
	data := make([]Entities.Message, limit)
	for i := 0; i < limit; i++ {
		data[i] = Entities.Message{
			Id:      i,
			Date:    fmt.Sprintf("%d.%d.20%d", i),
			Data:    "Lorem ipsum",
			IssueId: i,
		}
	}

	marshaled, err := json.Marshal(data)
	response := fmt.Sprintf("{meta: {size: %d}, data: %s}", 239, string(marshaled))
	w.Write([]byte(response))
	println(sqlQuery)
}

func PostIssue(w http.ResponseWriter, r *http.Request) {

	var issue Entities.IssueCreationWrapper
	bodyData := readBody(r.Body)
	err := json.Unmarshal([]byte(bodyData), &issue)
	if err != nil {
		http.Error(w, BadBody.Error(), http.StatusUnprocessableEntity)
		return
	}
	sqlQuery := fmt.Sprintf("PROCEDURE_NAME (%s, %s, %s, %s, %s, %s, %s)",
		issue.Issuer.Name,
		issue.Issuer.ContactInfo,
		issue.IssueMessage.Description,
		issue.ProblemCompany.Name,
		issue.ProblemCompany.Country,
		issue.ProblemCompany.ContactInfo,
		issue.ProblemCompany.OrgType,
	)

	w.Write([]byte(sqlQuery))
	//println(sqlQuery)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user Entities.User
	bodyData := readBody(r.Body)
	err := json.Unmarshal([]byte(bodyData), &user)
	if err != nil {
		http.Error(w, BadBody.Error(), http.StatusUnprocessableEntity)
		return
	}
	sqlQuery := fmt.Sprintf("INSERT INTO users (name, contactInfo) VALUES (%s, %s)",
		user.Name, user.ContactInfo,
	)

	w.Write([]byte(sqlQuery))
	//println(sqlQuery)
}

func PostMessage(w http.ResponseWriter, r *http.Request) {
	var message Entities.Message
	bodyData := readBody(r.Body)
	err := json.Unmarshal([]byte(bodyData), &message)
	if err != nil {
		http.Error(w, BadBody.Error(), http.StatusUnprocessableEntity)
		return
	}
	sqlQuery := fmt.Sprintf("INSERT INTO messages (data, date, issue_id) VALUES (%s, %s, %d)",
		message.Data, message.Date, message.IssueId,
	)

	w.Write([]byte(sqlQuery))
	//println(sqlQuery)
}

func PostOrganisation(w http.ResponseWriter, r *http.Request) {
	var organisation Entities.Organisation
	bodyData := readBody(r.Body)
	err := json.Unmarshal([]byte(bodyData), &organisation)
	if err != nil {
		http.Error(w, BadBody.Error(), http.StatusUnprocessableEntity)
		return
	}
	sqlQuery := fmt.Sprintf("INSERT INTO organisations (country, name, contacts, type) VALUES (%s, %s, %s, %s)",
		organisation.Country, organisation.Name, organisation.ContactInfo, organisation.OrgType,
	)

	w.Write([]byte(sqlQuery))
	//println(sqlQuery)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {

	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("DELETE FROM users WHERE id = %s", ids["id"])
	w.Write([]byte(sqlQuery))

}

func DeleteOrganisationById(w http.ResponseWriter, r *http.Request) {

	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("DELETE FROM organisations WHERE id = %s", ids["id"])
	w.Write([]byte(sqlQuery))

}

func DeleteMessageById(w http.ResponseWriter, r *http.Request) {

	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("DELETE FROM messages WHERE id = %s", ids["id"])
	w.Write([]byte(sqlQuery))

}

func DeleteIssueById(w http.ResponseWriter, r *http.Request) {

	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("DELETE FROM issues WHERE id = %s", ids["id"])
	w.Write([]byte(sqlQuery))

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user Entities.User
	bodyData := readBody(r.Body)
	err := json.Unmarshal([]byte(bodyData), &user)
	if err != nil {
		http.Error(w, BadBody.Error(), http.StatusUnprocessableEntity)
		return
	}
	sqlQuery := fmt.Sprintf("UPDATE users SET name = %s, contactInfo = %s where id = %d",
		user.Name, user.ContactInfo, user.Id,
	)

	w.Write([]byte(sqlQuery))
	//println(sqlQuery)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	var message Entities.Message
	bodyData := readBody(r.Body)
	err := json.Unmarshal([]byte(bodyData), &message)
	if err != nil {
		http.Error(w, BadBody.Error(), http.StatusUnprocessableEntity)
		return
	}
	sqlQuery := fmt.Sprintf("UPDATE messages SET data = %s, date=%s, issue_id=%d where id=%d",
		message.Data, message.Date, message.IssueId, message.Id,
	)

	w.Write([]byte(sqlQuery))
	//println(sqlQuery)
}

func UpdateOrganisation(w http.ResponseWriter, r *http.Request) {
	var organisation Entities.Organisation
	bodyData := readBody(r.Body)
	err := json.Unmarshal([]byte(bodyData), &organisation)
	if err != nil {
		http.Error(w, BadBody.Error(), http.StatusUnprocessableEntity)
		return
	}
	sqlQuery := fmt.Sprintf("UPDATE organisations SET country = %s, name=%s, contacts=%s, type=%s WHERE id = %d",
		organisation.Country, organisation.Name, organisation.ContactInfo, organisation.OrgType, organisation.Id,
	)

	w.Write([]byte(sqlQuery))
	//println(sqlQuery)
}

func UpdateIssue(w http.ResponseWriter, r *http.Request) {
	var issue Entities.Issue
	bodyData := readBody(r.Body)
	err := json.Unmarshal([]byte(bodyData), &issue)
	if err != nil {
		http.Error(w, BadBody.Error(), http.StatusUnprocessableEntity)
		return
	}
	sqlQuery := fmt.Sprintf("UPDATE issues SET status = %s, description=%s, organisation_id=%d, validation=%t WHERE id = %d",
		issue.Status, issue.Description, issue.OrganisationId, issue.Validation, issue.Id,
	)

	w.Write([]byte(sqlQuery))
	//println(sqlQuery)
}
