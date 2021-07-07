package models

//内容
type Content struct {
	Id             int64  `json:"id",gorm:"id"`
	Uuid           string `json:"uuid",gorm:"uuid"`
	PcCoverPicture string `json:"pc_cover_picture,omitempty",gorm:"pc_cover_picture"`
	MbCoverPicture string `json:"mb_cover_picture,omitempty",gorm:"mb_cover_picture"`
	Title          string `json:"title,omitempty",gorm:"title"`
	SubTitle       string `json:"sub_title,omitempty",gorm:"sub_title"`
	SearchCount    int    `json:"search_count,omitempty",gorm:"search_count"`
	Description    string `json:"description,omitempty",gorm:"description"`
	Carousel       string `json:"carousel,omitempty",gorm:"carousel"`
	City           string `json:"city,omitempty",gorm:"city"`
	Size           int    `json:"size",gorm:"size"`
	Status         int    `json:"status,omitempty",gorm:"status"`
	CreatedAt      int64  `json:"created_at,omitempty",gorm:"created_at"`
	UpdatedAt      int64  `json:"updated_at,omitempty",gorm:"updated_at"`
	CreatedAdminId string `json:"created_admin_id,omitempty",gorm:"created_admin_id"`
	UpdatedAdminId string `json:"updated_admin_id,omitempty",gorm:"updated_admin_id"`
	CategoryName   string `json:"category_name,omitempty",gorm:"category_name"` //添加自定义字段
}
