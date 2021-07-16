package models

//用户
type User struct {
	ID              *int64  `json:"id",gorm:"column:id;not null;primary_key;AUTO_INCREMENT;"`
	Uuid            string  `json:"uuid,omitempty",gorm:"uuid"`
	Name            string  `json:"name",gorm:"name"`
	Nickname        string  `json:"nickname",gorm:"nickname"`
	Avatar          string  `json:"avatar",gorm:"avatar"`
	Mobile          string  `json:"mobile,omitempty",gorm:"mobile"`
	Password        string  `json:"password,omitempty",gorm:"password"`
	Email           string  `json:"email,omitempty",gorm:"email"`
	EmailBindAt     *string `json:"email_bind_at,omitempty",gorm:"column:email_bind_at;default:null"`
	EmailVerifiedAt *string `json:"email_verified_at,omitempty",gorm:"column:email_verified_at;default:null"`
	WechatBindAt    *string `json:"wechat_bind_at,omitempty",gorm:"column:wechat_bind_at;default:null"`
	Status          int     `json:"status",gorm:"status"`
	Extra           *string `json:"extra,omitempty",gorm:"column:extra;default:null"`
	RegisterAt      int64   `json:"register_at,omitempty",gorm:"register_at"`
	RegisterIp      string  `json:"register_ip,omitempty",gorm:"register_ip"`
	LastLoginAt     int64   `json:"last_login_at,omitempty",gorm:"last_login_at"`
	LastLoginIp     string  `json:"last_login_ip,omitempty",gorm:"last_login_ip"`
	LoginCount      int     `json:"login_count,omitempty",gorm:"login_count"`
	RememberToken   string  `json:"remember_token,omitempty",gorm:"remember_token"`
	CreatedAt       int64   `json:"created_at,omitempty",gorm:"created_at"`
	UpdatedAt       int64   `json:"updated_at,omitempty",gorm:"updated_at"`
}
