package Handlers

import (
	"PeredelanoHakaton/Entities"
	"PeredelanoHakaton/Utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type DBWrapper struct {
	Db *sql.DB
}

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

func Ping(_ http.ResponseWriter, _ *http.Request) {
	log.Println("Pinged")
}

func (db DBWrapper) GetUserById(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}
	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("SELECT * FROM users WHERE id = %s", ids["id"])
	user := Entities.User{}
	err = Utils.ReadItemFromDb(db.Db, sqlQuery, &user.Id, &user.Name, &user.ContactInfo)
	if err != nil {
		http.Error(w, Utils.CantRead, http.StatusServiceUnavailable)
		return
	}

	issues, err := Utils.GetUserIssuesList(db.Db, user.Id)
	if err != nil {
		http.Error(w, Utils.CantRead, http.StatusServiceUnavailable)
		return
	}
	jsonIssues, err := json.Marshal(issues)
	if err != nil {
		http.Error(w, Utils.JSONConversionFailed, http.StatusServiceUnavailable)
		return
	}
	result := fmt.Sprintf(
		"{\"id\": %d, \"name\": \"%s\", \"contact_info\":\"%s\", \"issues\":%s}",
		user.Id,
		user.Name,
		user.ContactInfo,
		jsonIssues,
	)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(result))
	if err != nil {
		return
	}
	log.Println("Get user by ID")
}

func (db DBWrapper) GetAllUsersWhereParam(w http.ResponseWriter, r *http.Request) {

	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}
	queryParams := r.URL.Query()
	page, err := strconv.Atoi(queryParams.Get("page"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(queryParams.Get("limit"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlQuery := fmt.Sprintf("SELECT * FROM users LIMIT %d OFFSET %d", limit, limit*(page-1))
	data, err := Utils.GetUsersList(db.Db, sqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	marshaled, err := json.Marshal(data)
	amountOfUsers, err := Utils.GetAmountOfUsers(db.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	response := fmt.Sprintf("{\"meta\": {\"size\": %d}, \"data\": %s}", amountOfUsers, string(marshaled))
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(response))
	if err != nil {
		return
	}
	log.Println("Get users' list")
}

func (db DBWrapper) GetOrganisationById(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}

	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("SELECT * FROM organisations WHERE id = %s", ids["id"])
	organisation := Entities.Organisation{}
	err = Utils.ReadItemFromDb(
		db.Db,
		sqlQuery,
		&organisation.Id,
		&organisation.Country,
		&organisation.Name,
		&organisation.ContactInfo,
		&organisation.OrgType,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	organisation.AmountOfIssues, err = Utils.GetEntityAmountOfIssuesById(
		db.Db,
		organisation.Id,
		"organisation",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	orgJson, err := json.Marshal(organisation)
	if err != nil {
		http.Error(w, Utils.JSONConversionFailed, http.StatusServiceUnavailable)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(orgJson)
	if err != nil {
		return
	}
	log.Println("Get organisation by id")
}

func (db DBWrapper) GetAllOrganisationWhereParam(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	page, err := strconv.Atoi(queryParams.Get("page"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(queryParams.Get("limit"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlQuery := fmt.Sprintf("SELECT * FROM organisations LIMIT %d OFFSET %d", limit, limit*(page-1))
	data, err := Utils.GetOrganisationsList(db.Db, sqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	marshaled, err := json.Marshal(data)

	amountOfOrgs, err := Utils.GetAmountOfOrganisations(db.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	response := fmt.Sprintf("{\"meta\": {\"size\": %d}, \"data\": %s}", amountOfOrgs, string(marshaled))
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(response))
	if err != nil {
		return
	}
	log.Println("Get organisations' list")
}

func (db DBWrapper) GetIssueById(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}

	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("SELECT * FROM issues WHERE id = %s", ids["id"])
	issue := Entities.Issue{}
	err = Utils.ReadItemFromDb(
		db.Db,
		sqlQuery,
		&issue.Id,
		&issue.Status,
		&issue.Description,
		&issue.OrganisationId,
		&issue.UserId,
		&issue.Validation,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	issueJson, err := json.Marshal(issue)
	if err != nil {
		http.Error(w, Utils.JSONConversionFailed, http.StatusServiceUnavailable)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(issueJson)
	if err != nil {
		return
	}
	log.Println("Get issue by ID")
}

func (db DBWrapper) GetAllIssuesWhereParam(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	page, err := strconv.Atoi(queryParams.Get("page"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(queryParams.Get("limit"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlQuery := fmt.Sprintf("SELECT * FROM issues LIMIT %d OFFSET %d", limit, limit*(page-1))
	data, err := Utils.GetIssuesList(db.Db, sqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	marshaled, err := json.Marshal(data)
	if err != nil {
		http.Error(w, Utils.JSONConversionFailed, http.StatusServiceUnavailable)
		return
	}

	amountOfIssues, err := Utils.GetAmountOfIssues(db.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := fmt.Sprintf("{\"meta\": {\"size\": %d}, \"data\": %s}", amountOfIssues, string(marshaled))
	_, err = w.Write([]byte(response))
	if err != nil {
		return
	}

	log.Println("Get issues' list")
}

func (db DBWrapper) GetMessageById(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}

	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("SELECT * FROM messages WHERE id = %s", ids["id"])
	message := Entities.Message{}
	err = Utils.ReadItemFromDb(
		db.Db,
		sqlQuery,
		&message.Id,
		&message.Data,
		&message.Date,
		&message.IssueId,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	messageJson, err := json.Marshal(message)
	if err != nil {
		http.Error(w, Utils.JSONConversionFailed, http.StatusServiceUnavailable)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(messageJson)
	if err != nil {
		return
	}
	log.Println("Get message by ID")
}

func (db DBWrapper) GetAllMessagesWhereParam(w http.ResponseWriter, r *http.Request) {

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
	data, err := Utils.GetMessagesList(db.Db, sqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	marshaled, err := json.Marshal(data)
	if err != nil {
		http.Error(w, Utils.JSONConversionFailed, http.StatusServiceUnavailable)
		return
	}

	amountOfMessages, err := Utils.GetAmountOfMessages(db.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := fmt.Sprintf("{\"meta\": {\"size\": %d}, \"data\": %s}", amountOfMessages, string(marshaled))
	_, err = w.Write([]byte(response))
	if err != nil {
		return
	}
	log.Println("Get messages' list")
}

func (db DBWrapper) PostIssue(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}

	var issue Entities.IssueCreationWrapper
	bodyData := readBody(r.Body)
	err = json.Unmarshal([]byte(bodyData), &issue)
	if err != nil {
		http.Error(w, Utils.BadBody, http.StatusUnprocessableEntity)
		return
	}

	sqlQuery := fmt.Sprintf("CALL insert_issue ('%s', '%s', '%s', '%s', '%s', '%s', '%s')",
		issue.Issuer.Name,
		issue.Issuer.ContactInfo,
		issue.IssueMessage.Description,
		issue.ProblemCompany.Name,
		issue.ProblemCompany.Country,
		issue.ProblemCompany.ContactInfo,
		issue.ProblemCompany.OrgType,
	)
	_, err = db.Db.Exec(sqlQuery)
	if err != nil {
		http.Error(w, Utils.CantWrite, http.StatusServiceUnavailable)
		return
	}
	log.Println("Post issue")
}

func (db DBWrapper) PostUser(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}

	var user Entities.User
	bodyData := readBody(r.Body)
	err = json.Unmarshal([]byte(bodyData), &user)
	if err != nil {
		http.Error(w, Utils.BadBody, http.StatusUnprocessableEntity)
		return
	}

	sqlQuery := fmt.Sprintf("INSERT INTO users (name, Contact) VALUES ('%s', '%s')",
		user.Name, user.ContactInfo,
	)
	_, err = db.Db.Exec(sqlQuery)
	if err != nil {
		http.Error(w, Utils.CantWrite, http.StatusServiceUnavailable)
		return
	}
	log.Println("Post user")
}

func (db DBWrapper) PostMessage(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}

	var message Entities.Message
	bodyData := readBody(r.Body)
	err = json.Unmarshal([]byte(bodyData), &message)
	if err != nil {
		http.Error(w, Utils.BadBody, http.StatusUnprocessableEntity)
		return
	}

	sqlQuery := fmt.Sprintf("INSERT INTO messages (data, date, issue_id) VALUES ('%s', TO_DATE('%s', 'DD.MM.YYYY'), %d)",
		message.Data, message.Date, message.IssueId,
	)
	_, err = db.Db.Exec(sqlQuery)
	if err != nil {
		http.Error(w, Utils.CantWrite, http.StatusServiceUnavailable)
		return
	}
	log.Println("Post message")
}

func (db DBWrapper) PostOrganisation(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}

	var organisation Entities.Organisation
	bodyData := readBody(r.Body)
	err = json.Unmarshal([]byte(bodyData), &organisation)
	if err != nil {
		http.Error(w, Utils.BadBody, http.StatusUnprocessableEntity)
		return
	}

	sqlQuery := fmt.Sprintf("INSERT INTO organisations (country, name, contacts, type) VALUES ('%s', '%s', '%s', '%s')",
		organisation.Country, organisation.Name, organisation.ContactInfo, organisation.OrgType,
	)
	_, err = db.Db.Exec(sqlQuery)
	if err != nil {
		http.Error(w, Utils.CantWrite, http.StatusServiceUnavailable)
		return
	}
	log.Println("Post organisation")
}

func (db DBWrapper) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}

	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("DELETE FROM users WHERE id = %s", ids["id"])

	_, err = db.Db.Exec(sqlQuery)
	if err != nil {
		http.Error(w, Utils.CantWrite, http.StatusServiceUnavailable)
		return
	}
	log.Println("Delete user")
}

func (db DBWrapper) DeleteOrganisationById(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}
	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("DELETE FROM organisations WHERE id = %s", ids["id"])
	_, err = db.Db.Exec(sqlQuery)
	if err != nil {
		http.Error(w, Utils.CantWrite, http.StatusServiceUnavailable)
		return
	}
	log.Println("Delete organisation")
}

func (db DBWrapper) DeleteMessageById(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}
	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("DELETE FROM messages WHERE id = %s", ids["id"])
	_, err = db.Db.Exec(sqlQuery)
	if err != nil {
		http.Error(w, Utils.CantWrite, http.StatusServiceUnavailable)
		return
	}
	log.Println("Delete message")
}

func (db DBWrapper) DeleteIssueById(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}
	ids := mux.Vars(r)
	sqlQuery := fmt.Sprintf("DELETE FROM issues WHERE id = %s", ids["id"])
	_, err = db.Db.Exec(sqlQuery)
	if err != nil {
		http.Error(w, Utils.CantWrite, http.StatusServiceUnavailable)
		return
	}
	log.Println("Delete issue")
}

func (db DBWrapper) UpdateUser(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}
	var user Entities.User
	bodyData := readBody(r.Body)
	err = json.Unmarshal([]byte(bodyData), &user)
	if err != nil {
		http.Error(w, Utils.BadBody, http.StatusUnprocessableEntity)
		return
	}
	sqlQuery := fmt.Sprintf("UPDATE users SET name = '%s', contact = '%s' where id = %d",
		user.Name, user.ContactInfo, user.Id,
	)
	_, err = db.Db.Exec(sqlQuery)
	if err != nil {
		http.Error(w, Utils.CantWrite, http.StatusServiceUnavailable)
		return
	}
	log.Println("Update user")
}

func (db DBWrapper) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}
	var message Entities.Message
	bodyData := readBody(r.Body)
	err = json.Unmarshal([]byte(bodyData), &message)
	if err != nil {
		http.Error(w, Utils.BadBody, http.StatusUnprocessableEntity)
		return
	}
	sqlQuery := fmt.Sprintf("UPDATE messages SET data = %s, date=%s, issue_id=%d where id=%d",
		message.Data, message.Date, message.IssueId, message.Id,
	)
	_, err = db.Db.Exec(sqlQuery)
	if err != nil {
		http.Error(w, Utils.CantWrite, http.StatusUnprocessableEntity)
		return
	}
	log.Println("Update message")
}

func (db DBWrapper) UpdateOrganisation(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}

	var organisation Entities.Organisation
	bodyData := readBody(r.Body)
	err = json.Unmarshal([]byte(bodyData), &organisation)
	if err != nil {
		http.Error(w, Utils.BadBody, http.StatusUnprocessableEntity)
		return
	}
	sqlQuery := fmt.Sprintf("CALL update_organisation(%d,'%s','%s','%s','%s')",
		organisation.Id,
		organisation.Country,
		organisation.Name,
		organisation.ContactInfo,
		organisation.OrgType,
	)

	_, err = db.Db.Exec(sqlQuery)
	if err != nil {
		http.Error(w, Utils.CantWrite, http.StatusUnprocessableEntity)
		return
	}
	log.Println("Update organisation")
}

func (db DBWrapper) UpdateIssue(w http.ResponseWriter, r *http.Request) {
	err := db.Db.Ping()
	if err != nil {
		http.Error(w, Utils.ConnectionDied, http.StatusServiceUnavailable)
		return
	}
	var issue Entities.Issue
	bodyData := readBody(r.Body)
	err = json.Unmarshal([]byte(bodyData), &issue)
	if err != nil {
		http.Error(w, Utils.BadBody, http.StatusUnprocessableEntity)
		return
	}
	sqlQuery := fmt.Sprintf("UPDATE issues SET status = %s, description=%s, organisation_id=%d, validation=%t WHERE id = %d",
		issue.Status, issue.Description, issue.OrganisationId, issue.Validation, issue.Id,
	)

	_, err = db.Db.Exec(sqlQuery)
	if err != nil {
		http.Error(w, Utils.CantWrite, http.StatusUnprocessableEntity)
		return
	}
	log.Println("Update issue")
}
