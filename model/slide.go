package models

// 幻灯片
type Slide struct {
	ID          int64  `db:"id"json:"id"`
	TargetUrl   string `db:"target_url" json:"target_url"`
	CarouselUrl string `db:"carousel_url" json:"carousel_url"`
}
