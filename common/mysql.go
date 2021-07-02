package common

import (
	"design-api/config"
	//_ "github.com/go-sql-driver/mysql"
	//_ "github.com/jmoiron/sqlx"
	//"github.com/jinzhu/gorm"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func init() {
	dsn := config.Config.Mysql.DbUsername+":"+config.Config.Mysql.DbPassword+"@tcp("+config.Config.Mysql.DbHost+":"+config.Config.Mysql.DbPort+")/"+config.Config.Mysql.DbDatabase+"?charset=utf8&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{
		NamingStrategy:schema.NamingStrategy{
			TablePrefix:"q_",
		},
	})

	if err == nil && db != nil {
		Db = db
	} else {
		log.Println("mysql 连接错误")
	}
}

/**
获取表名
*/
func Table(prefix, table string) string {
	return prefix + table
}
