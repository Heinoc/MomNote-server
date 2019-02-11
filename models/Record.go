package models

import "time"

/**
 * author: chenshuai09
 */

type Record struct {
	ID        uint `gorm:"primary_key" json:"id" form:"id" binding:""`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	//    体重
	Weight string `json:"weight"`
	//    腰围
	Waistline string `json:"waistline"`
	//    臀围
	Hipline string `json:"hipline"`
	//    大腿围
	Thighline string `json:"thighline"`
}
