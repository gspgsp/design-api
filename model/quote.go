package models

//报价
type Quote struct {
	Id        int64  `json:"id" gorm:"id"`
	Name      string `json:"name" gorm:"name" form:"name" binding:"required"`
	Phone     string `json:"phone" gorm:"phone" form:"phone" binding:"required,phoneValidator"`
	Address   string `json:"address" gorm:"address" form:"address" binding:"required"`
	Size      int    `json:"size" gorm:"size" form:"size" binding:"required,gt=0"`
	UserId    int    `json:"user_id" gorm:"user_id"`
	Status    int    `json:"status" gorm:"status"`
	CreatedAt int64  `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt int64  `json:"updated_at,omitempty" gorm:"updated_at"`
}
