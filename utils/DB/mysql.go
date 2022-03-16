package DB

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Mysql *gorm.DB
)

func init() {
	var err error
	dsn := "root:12345@tcp(127.0.0.1:3306)/temp?charset=utf8&parseTime=True&loc=Local"
	Mysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connect DB error")
		panic(err)
	}
}
