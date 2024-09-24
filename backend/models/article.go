package models

type Article struct {
	ID         int    `json:"id"`
	Title      string `json:"title" validate:"required,min=1,max=50"`
	CoverImage string `json:"cover_image"`
	Intro      string `json:"intro"`
	Keywords   string `json:"keywords"`
	Content    string `json:"content"`
	Views      int    `json:"views"`
	CreatorID  int    `json:"creator_id"`
	CreateTime string `json:"create_time"`
	Status     string `json:"status" validate:"required,oneof=0 1 2"`
}
