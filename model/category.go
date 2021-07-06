package models

//分类
type Category struct {
	ID             int64  `json:"id",gorm:"id"`
	Belong         string `json:"belong,omitempty",gorm:"belong"`
	Abbreviation   string `json:"abbreviation,omitempty",gorm:"abbreviation"`
	Name           string `json:"name",gorm:"name"`
	MbCoverPicture string `json:"mb_cover_picture",gorm:"mb_cover_picture"`
	PcCoverPicture string `json:"pc_cover_picture,omitempty",gorm:"pc_cover_picture"`
	Status         int    `json:"status,omitempty",gorm:"status"`
	CreatedAt      int64  `json:"created_at,omitempty",gorm:"created_at"`
	UpdatedAt      int64  `json:"updated_at,omitempty",gorm:"updated_at"`
	CreatedAdminId string `json:"created_admin_id,omitempty",gorm:"created_admin_id"`
	UpdatedAdminId string `json:"updated_admin_id,omitempty",gorm:"updated_admin_id"`
}
