package models

import (
	"errors"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

/**
 * author: chenshuai09
 * created on: 2019-02-12
 */

type User struct {
	ID        uint `gorm:"primary_key" json:"-" form:"id" binding:""`
	CreatedAt time.Time
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`

	UserId string `json:"userId"`
	UserName string `json:"userName"`
	Phone string `json:"phone"`
	Password string `json:"-"`
	BirthDay *time.Time `json:"birthDay"`


}

func (c *User) Inser() error {

	if c.UserId == "" {
		c.UserId = xid.New().String()
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	c.Password = string(hash)

	return DBUtil().Create(c).Error

}

func FindUserByUserId(userId string) (user User, err error) {
	err = DBUtil().Find(&user, "user_id=?", userId).Error

	return
}

func VerifyUserByPwd(phone, userName, password string) (user User, err error) {
	err = DBUtil().Find(&user, "phone=? or user_name=?", phone, userName).Error

	// 正确密码验证
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("用户密码错误！")
	}

	return
}