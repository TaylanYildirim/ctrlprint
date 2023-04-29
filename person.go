package models

type Person struct {
	PersonID         int    `json:"person_id"`
	PersonUsername   string `json:"person_username"`
	PersonCompanyID  int    `json:"person_company_id"`
	PersonLanguageID int    `json:"person_language_id"`
}
