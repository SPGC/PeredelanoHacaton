package Entities

type Organisation struct {
	Id          int    `json:"id,omitempty"`
	Country     string `json:"country"`
	Name        string `json:"name"`
	ContactInfo string `json:"contact_info"`
	OrgType     string `json:"org_type"`
}

type Issue struct {
	Id             int    `json:"id,omitempty"`
	Status         string `json:"status,omitempty"`
	Description    string `json:"description"`
	OrganisationId int    `json:"organisation_id,omitempty"`
}

type Message struct {
	Id      int    `json:"id,omitempty"`
	Data    string `json:"data"`
	Date    string `json:"date"`
	IssueId string `json:"issue_id"`
}

type User struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	ContactInfo string `json:"contact_info"`
}

type IssueTemplate struct {
	Issuer User
}
