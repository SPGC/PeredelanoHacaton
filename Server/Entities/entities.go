package Entities

type Organisation struct {
	Id             int    `json:"id,omitempty"`
	Country        string `json:"country"`
	Name           string `json:"name"`
	ContactInfo    string `json:"contact_info"`
	OrgType        string `json:"org_type"`
	AmountOfIssues int    `json:"amount_of_issues,omitempty"`
}

type Issue struct {
	Id               int    `json:"id,omitempty"`
	Status           string `json:"status,omitempty"`
	Description      string `json:"description"`
	OrganisationId   int    `json:"organisation_id,omitempty"`
	OrganisationName string `json:"organisation_name,omitempty"`
	Validation       bool   `json:"validation,omitempty"`
	UserId           int    `json:"user_id,omitempty"`
}

type Message struct {
	Id      int    `json:"id,omitempty"`
	Data    string `json:"data"`
	Date    string `json:"date"`
	IssueId int    `json:"issue_id"`
}

type User struct {
	Id             int    `json:"id,omitempty"`
	Name           string `json:"name"`
	ContactInfo    string `json:"contact_info"`
	AmountOfIssues int    `json:"amount_of_issues,omitempty"`
}

type IssueCreationWrapper struct {
	Issuer         User         `json:"issuer"`
	ProblemCompany Organisation `json:"company"`
	IssueMessage   Issue        `json:"message"`
}
