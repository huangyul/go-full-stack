package main

import (
	"go_srvs/user_srv/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// 连接数据库

func main() {
	dsn := "root:123456@tcp(47.106.214.127)/go_user_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger, NamingStrategy: schema.NamingStrategy{SingularTable: true}})

	if err != nil {
		panic(err)
	}

	// 迁移表
	_ = db.AutoMigrate(&model.User{})
}
