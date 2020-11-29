package dao

import "github.com/jinzhu/gorm"
import  _ "github.com/go-sql-driver/mysql"

var(
	DB *gorm.DB
)

func InitMysql() (err error){
	dsn := "root:jfdxa1304@tcp(127.0.0.1:3306)/todoList?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err!=nil{
		return
	}
	err = DB.DB().Ping()
	return
}