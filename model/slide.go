package models

// 幻灯片
type Slide struct {
	ID             int64  `json:"id",gorm:"id"`
	Device         int    `json:"device,omitempty",gorm:"device"`
	Title          string `json:"title,omitempty",gorm:"title"`
	TargetUrl      string `json:"target_url",gorm:"target_url"`
	CarouselUrl    string `json:"carousel_url",gorm:"carousel_url"`
	Sort           int    `json:"sort,omitempty",gorm:"sort"`
	Status         int    `json:"status,omitempty",gorm:"status"`
	Description    string `json:"description,omitempty",gorm:"description"`
	CreatedAt      int64  `json:"created_at,omitempty",gorm:"created_at"`
	UpdatedAt      int64  `json:"updated_at,omitempty",gorm:"updated_at"`
	CreatedAdminId string `json:"created_admin_id,omitempty",gorm:"created_admin_id"`
	UpdatedAdminId string `json:"updated_admin_id,omitempty",gorm:"updated_admin_id"`
}
