package models

type Project struct {
	ID          int    `json:"id"`
	ProjectName string `json:"project_name" validate:"required,min=1,max=20"`
	Description string `json:"description" validate:"min=0,max=100"`
	Logo        string `json:"logo"`
	Url         string `json:"url"`
}
