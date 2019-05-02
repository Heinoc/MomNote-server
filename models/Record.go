package models

import (
	"math"
	"time"
)

/**
 * author: heinoc
 */

type Record struct {
	ID        uint       `gorm:"primary_key" json:"-" form:"id" binding:""`
	CreatedAt time.Time  `json:"createdTime"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`

	User   User   `json:"-"`
	UserId string `json:"userId"`

	//体重
	Weight string `json:"weight"`
	//臂围
	Armline string `json:"armline"`
	//腰围
	Waistline string `json:"waistline"`
	//胸围
	Bust string `json:"bust"`
	//臀围
	Hipline string `json:"hipline"`
	//大腿围
	Thighline string `json:"thighline"`
}

func (c *Record) Insert() error {

	return DBUtil().Create(c).Error

}

//查询指定pageNum的pageSize容量的记录列表，如果pageSize=0，则返回全部列表
func FindAllRecords(userId string, pageNum, pageSize int64) (totalPages int64, list []Record, err error) {
	if pageSize > 0 {
		//查询总页数
		var count int64
		DBUtil().Find(&list).Count(&count)
		totalPages = int64(math.Ceil(float64(count) / float64(pageSize)))

		err = DBUtil().Order("id desc").
			Where("user_id = ?", userId).
			Offset((pageNum - 1) * pageSize).
			Limit(pageSize).
			Find(&list).
			Error

	} else {
		err = DBUtil().Order("id desc").Where("user_id = ?", userId).Find(&list).Error
	}

	return

}
