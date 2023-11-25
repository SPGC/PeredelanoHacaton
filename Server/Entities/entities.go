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
	Validation     bool   `json:"validation,omitempty"`
}

type Message struct {
	Id      int    `json:"id,omitempty"`
	Data    string `json:"data"`
	Date    string `json:"date"`
	IssueId int    `json:"issue_id"`
}

type User struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	ContactInfo string `json:"contact_info"`
}

type IssueCreationWrapper struct {
	Issuer         User         `json:"issuer"`
	ProblemCompany Organisation `json:"company"`
	IssueMessage   Issue        `json:"message"`
}
