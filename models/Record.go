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
}