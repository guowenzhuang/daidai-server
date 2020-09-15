package models

import (
	"daidai-server/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var (
	DB    *gorm.DB
	DbErr error
)

func init() {
	var (
		err                      error
		User, Password, Host, Db string
		Port                     int
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	User, Password, Host, Port, Db = sec.Key("USER").String(),
		sec.Key("PASSWORD").String(),
		sec.Key("HOST").String(),
		sec.Key("PORT").MustInt(),
		sec.Key("DATABASE").String()

	connArgs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", User, Password, Host, Port, Db)
	DB, DbErr = gorm.Open(mysql.New(mysql.Config{
		DSN:                       connArgs, // data source name
		DefaultStringSize:         256,      // default size for string fields
		DisableDatetimePrecision:  true,     // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,     // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,     // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,    // auto configure based on currently MySQL version
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})

	if DbErr != nil {
		log.Println("failed to connect database %v", DbErr)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("failed to connect database")
	}

	DB.AutoMigrate(&Users{})
	DB.AutoMigrate(&Demand{})

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("failed to connect database")
	}
	sqlDB.Close()
}
