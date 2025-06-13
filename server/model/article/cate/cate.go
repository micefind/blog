package cate

type Cate struct {
	Id         int    `json:"id"`
	Name       string `json:"name" validate:"required,min=1,max=12,excludesall= "`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	CreatorId  int    `json:"creator_id"`
	UpdaterId  int    `json:"updater_id"`
}

type Query struct {
	PageNum  *int   `json:"page_num"`
	PageSize *int   `json:"page_size"`
	Name     string `json:"name"`
}

type CateDetail struct {
	Cate
	Creator string `json:"creator"`
	Updater string `json:"updater"`
}

type List struct {
	List  []CateDetail `json:"list"`
	Total int64        `json:"total"`
}
