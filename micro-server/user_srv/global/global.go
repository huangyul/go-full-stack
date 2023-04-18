package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func init() {
	// 连接数据库
	dsn := "root:root@tcp(47.106.214.127:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"

	// 日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
			//IgnoreRecordNotFoundError: true,
			//ParameterizedQueries:      true,
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 自动加表名前缀
			SingularTable: true, // 表名单数形式
			NameReplacer:  nil,
		},
	})
	if err != nil {
		panic(err)
	}
}
