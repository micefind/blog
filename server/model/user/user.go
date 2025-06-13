package user

type User struct {
	Id         int    `json:"id" `
	Username   string `json:"username" validate:"required,min=1,max=12,excludesall= " label:"用户名"`
	Password   string `json:"password" validate:"omitempty,min=6,max=30,excludesall= ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=0123456789,containsany=_.!@#$%^&*()" label:"密码"`
	Name       string `json:"name"`
	Mobile     string `json:"mobile" label:"手机号"`
	Email      string `json:"email" validate:"omitempty,email" label:"邮箱"`
	Avatar     string `json:"avatar"`
	Role       int `json:"role"`
	CreateTime string `json:"create_time"`
	IsDeleted  int    `json:"is_deleted"`
}

type Query struct {
	PageNum   *int   `json:"page_num"`
	PageSize  *int   `json:"page_size"`
	Username  string `json:"username"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	IsDeleted *int   `json:"is_deleted"`
}

type List struct {
	Total int64  `json:"total"`
	List  []User `json:"list"`
}

type ChangePassword struct {
	Id          int    `json:"id"`
	Password    string `json:"password" validate:"required" label:"旧密码"`
	NewPassword string `json:"new_password" validate:"required,min=6,max=30,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=0123456789,containsany=_.!@#$%^&*()" label:"新密码"`
}
