package models

type User struct {
	ID         int64  `json:"id"`
	Uuid       string `json:"uuid"`
	Name       string `json:"name"`
	NickName   string `json:"nick_name"`
	Avatar     string `json:"avatar"`
	Mobile     string `json:"mobile"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	RegisterAt int64  `json:"register_at"`
	RegisterIp string `json:"register_ip"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}
