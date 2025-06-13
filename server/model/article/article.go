package article

type Article struct {
	Id         int    `json:"id"`
	CateId     int    `json:"cate_id" validate:"required"`
	TagId      int    `json:"tag_id" validate:"required"`
	Title      string `json:"title" validate:"required,min=1,max=50" label:"标题"`
	Intro      string `json:"intro" validate:"required,min=1,max=200" label:"摘要"`
	Content    string `json:"content" validate:"required" label:"正文"`
	CoverImg   string `json:"cover_img"`
	Views      int    `json:"views"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
	CreatorId  int    `json:"creator_id"`
	UpdateTime string `json:"update_time"`
	UpdaterId  int    `json:"updater_id"`
	IsDeleted  int    `json:"is_deleted"`
}


type ArticleDetail struct {
	Article
	Creator  string `json:"creator"`
	Updater  string `json:"updater"`
	CateName string `json:"cate_name"`
	TagName  string `json:"tag_name"`
}

type Query struct {
	PageNum   *int   `json:"page_num"`
	PageSize  *int   `json:"page_size"`
	Keyword   string `json:"keyword"`
	CateId    *int   `json:"cate_id"`
	TagId     *int   `json:"tag_id"`
	IsDeleted *int   `json:"is_deleted"`
	Status    *int   `json:"status"`
}

type List struct {
	List  []ArticleDetail `json:"list"`
	Total int64           `json:"total"`
}
