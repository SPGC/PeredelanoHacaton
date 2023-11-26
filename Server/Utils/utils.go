package Utils

import (
	"PeredelanoHakaton/Entities"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

const (
	ConnectionDied       = "connection to data base has died"
	CantRead             = "can't read from data base"
	BadBody              = "can't parse body"
	CantWrite            = "can't write to data base"
	JSONConversionFailed = "can't convert to JSON"
	UnsupportedType      = "unsupported type"
)

var (
	ErrorCantRead        = errors.New(CantRead)
	ErrorUnsupportedType = errors.New(UnsupportedType)
	MaxDelay             = time.Duration.Seconds(30)
)

func ReadItemFromDb(db *sql.DB, sqlQuery string, args ...interface{}) error {
	var rows *sql.Rows
	start := time.Now()
	var err error
	for true {
		rows, err = db.Query(sqlQuery)
		if err == nil {
			break
		}
		current := time.Now()
		if current.Sub(start).Seconds() > MaxDelay && err != nil {
			return ErrorCantRead
		}
	}

	for rows.Next() {
		err = rows.Scan(args...)
		if err != nil {
			//return ErrorCantRead
			return err
		}
	}

	return nil
}

func GetOrgNameById(db *sql.DB, id int) (string, error) {
	sqlQuery := fmt.Sprintf("SELECT name FROM organisations WHERE id=%d", id)

	rows, err := db.Query(sqlQuery, id)
	if err != nil {
		return "", ErrorCantRead
	}
	var result string
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return "", ErrorCantRead
		}
	}
	return result, nil
}

func GetOrgCountryById(db *sql.DB, id int) (string, error) {
	sqlQuery := fmt.Sprintf("SELECT country FROM organisations WHERE id=%d", id)

	rows, err := db.Query(sqlQuery, id)
	if err != nil {
		return "", ErrorCantRead
	}
	var result string
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return "", ErrorCantRead
		}
	}
	return result, nil
}

func GetUserIssuesList(db *sql.DB, id int) ([]Entities.Issue, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM issues WHERE user_id = %d", id)

	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, ErrorCantRead
	}
	issues := make([]Entities.Issue, 0)
	counter := 0
	for rows.Next() {
		issues = append(issues, Entities.Issue{})
		err = rows.Scan(
			&issues[counter].Id,
			&issues[counter].Status,
			&issues[counter].Description,
			&issues[counter].OrganisationId,
			&issues[counter].UserId,
			&issues[counter].Validation,
		)
		if err != nil {
			return nil, ErrorCantRead
		}
		issues[counter].OrganisationName, err = GetOrgNameById(db, issues[counter].OrganisationId)
		if err != nil {
			return nil, ErrorCantRead
		}
		counter++
	}
	return issues, nil
}

func GetEntityAmountOfIssuesById(db *sql.DB, id int, entityType string) (int, error) {
	var key string
	if entityType == "user" {
		key = "user_id"
	} else if entityType == "organisation" {
		key = "organisation_id"
	} else {
		return 0, ErrorUnsupportedType
	}
	var amount int
	sqlQuery := fmt.Sprintf("SELECT count(*) FROM issues WHERE %s=%d", key, id)
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return 0, ErrorCantRead
	}
	for rows.Next() {
		err = rows.Scan(&amount)
		if err != nil {
			return 0, ErrorCantRead
		}
	}
	return amount, nil
}

func GetUsersList(db *sql.DB, sqlQuery string) ([]Entities.User, error) {
	var rows *sql.Rows
	start := time.Now()
	var err error
	data := make([]Entities.User, 0)
	for true {
		data = make([]Entities.User, 0)
		flag := true
		rows, err = db.Query(sqlQuery)
		flag = flag && (err == nil)
		counter := 0
		for rows.Next() {
			data = append(data, Entities.User{})
			err = rows.Scan(&data[counter].Id, &data[counter].Name, &data[counter].ContactInfo)
			flag = flag && (err == nil)
			data[counter].AmountOfIssues, err = GetEntityAmountOfIssuesById(db, data[counter].Id, "user")
			flag = flag && (err == nil)
			counter++
		}
		if flag {
			return data, nil
		}
		current := time.Now()
		if current.Sub(start).Seconds() > MaxDelay && err != nil {
			return nil, ErrorCantRead
		}
	}
	return data, nil
}

func GetAmountOfUsers(db *sql.DB) (int, error) {
	var amount int
	sqlQuery := "SELECT count(*) FROM users"
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return 0, ErrorCantRead
	}
	for rows.Next() {
		err = rows.Scan(&amount)
		if err != nil {
			return 0, ErrorCantRead
		}
	}
	return amount, nil
}

func GetOrganisationsList(db *sql.DB, sqlQuery string) ([]Entities.Organisation, error) {
	data := make([]Entities.Organisation, 0)
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, ErrorCantRead
	}
	counter := 0
	for rows.Next() {
		data = append(data, Entities.Organisation{})
		err = rows.Scan(
			&data[counter].Id,
			&data[counter].Country,
			&data[counter].Name,
			&data[counter].ContactInfo,
			&data[counter].OrgType,
		)
		if err != nil {
			return nil, ErrorCantRead
		}

		data[counter].AmountOfIssues, err = GetEntityAmountOfIssuesById(db, data[counter].Id, "organisation")

		if err != nil {
			return nil, ErrorCantRead
		}
		counter++
	}
	return data, nil
}

func GetAmountOfOrganisations(db *sql.DB) (int, error) {
	var amount int
	sqlQuery := "SELECT count(*) FROM organisations"
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return 0, ErrorCantRead
	}
	for rows.Next() {
		err = rows.Scan(&amount)
		if err != nil {
			return 0, ErrorCantRead
		}
	}
	return amount, nil
}

func GetIssuesList(db *sql.DB, sqlQuery string) ([]Entities.Issue, error) {
	var rows *sql.Rows
	start := time.Now()
	var err error
	data := make([]Entities.Issue, 0)
	for true {
		flag := true
		rows, err = db.Query(sqlQuery)
		flag = flag && (err == nil)
		counter := 0
		for rows.Next() {
			data = append(data, Entities.Issue{})
			err = rows.Scan(
				&data[counter].Id,
				&data[counter].Status,
				&data[counter].Description,
				&data[counter].OrganisationId,
				&data[counter].UserId,
				&data[counter].Validation,
			)
			flag = flag && (err == nil)

			data[counter].OrganisationName, err = GetOrgNameById(db, data[counter].OrganisationId)
			flag = flag && (err == nil)

			data[counter].OrganisationCountry, err = GetOrgCountryById(db, data[counter].OrganisationId)
			flag = flag && (err == nil)

			counter++
		}
		if flag {
			return data, nil
		}
		current := time.Now()
		if current.Sub(start).Seconds() > MaxDelay && err != nil {
			return nil, ErrorCantRead
		}
	}

	return data, nil
}

func GetAmountOfIssues(db *sql.DB) (int, error) {
	var amount int
	sqlQuery := "SELECT count(*) FROM issues"
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return 0, ErrorCantRead
	}
	for rows.Next() {
		err = rows.Scan(&amount)
		if err != nil {
			return 0, ErrorCantRead
		}
	}

	return amount, nil
}

func GetMessagesList(db *sql.DB, sqlQuery string) ([]Entities.Message, error) {
	data := make([]Entities.Message, 0)
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, ErrorCantRead
	}

	counter := 0
	for rows.Next() {
		data = append(data, Entities.Message{})
		err = rows.Scan(
			&data[counter].Id,
			&data[counter].Data,
			&data[counter].Date,
			&data[counter].IssueId,
		)
		if err != nil {
			return nil, ErrorCantRead
		}

		if err != nil {
			return nil, ErrorCantRead
		}
		counter++
	}
	return data, nil
}

func GetAmountOfMessages(db *sql.DB) (int, error) {
	var amount int
	sqlQuery := "SELECT count(*) FROM messages"
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return 0, ErrorCantRead
	}
	for rows.Next() {
		err = rows.Scan(&amount)
		if err != nil {
			return 0, ErrorCantRead
		}
	}
	return amount, nil
}
