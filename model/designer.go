package models

//设计师
type Designer struct {
	Id             int64  `json:"id"`
	Uuid           string `json:"uuid,omitempty" gorm:"uuid" uri:"uuid" binding:"required"`
	Name           string `json:"name,omitempty"`
	NickNme        string `json:"nick_nme,omitempty"`
	Photo          string `json:"photo,omitempty"`
	Level          int    `json:"level,omitempty"`
	Motto          string `json:"motto,omitempty"`
	Description    string `json:"description,omitempty"`
	CreatedAt      int64  `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt      int64  `json:"updated_at,omitempty" gorm:"updated_at"`
	CreatedAdminId string `json:"created_admin_id,omitempty" gorm:"created_admin_id"`
	UpdatedAdminId string `json:"updated_admin_id,omitempty" gorm:"updated_admin_id"`
	FansCount      int64  `json:"fans_count" gorm:"fans_count"`
	ContentCount   int64  `json:"content_count" gorm:"content_count"`
}
