package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Test struct {
	gorm.Model
	Id   uint
	Name string
}

func main() {
	User, Password, Host, Port, Db := "root", "123456", "39.100.126.197", 3306, "daidai"

	connArgs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", User, Password, Host, Port, Db)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connArgs, // data source name
		DefaultStringSize:         256,      // default size for string fields
		DisableDatetimePrecision:  true,     // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,     // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,     // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,    // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Test{})

	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		db.Create(&Test{Name: "111"})
		context.String(http.StatusOK, "hello gin")
	})

	r.Run(":9999")
}
