package models

type Project struct {
	ProjectID        int    `json:"project_id"`
	ProjectCompanyID int    `json:"project_company_id"`
	ProjectName      string `json:"project_name"`
}
