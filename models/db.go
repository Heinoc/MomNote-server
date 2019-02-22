package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sync"
)

/**
 * Created by Heinoc on 2018/8/29.
 */

var db *gorm.DB
var once sync.Once

/**
单例，获取数据库对象
 */
func DBUtil() *gorm.DB {
	once.Do(func() {
		d, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:23306)/mom_note?charset=utf8&parseTime=True&loc=Local")

		if err != nil {
			log.Fatal(err)
		}
		db = d

		/**
		model 注册
		 */
		db.AutoMigrate(
			&User{},
			&Record{},
		)

	})

	return db
}