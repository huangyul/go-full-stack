package initialize

import (
	"fmt"
	"go_srvs/user_srv/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitDB() {
	//dsn := "root:123456@tcp(192.168.121.136:3306)/go_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", global.ServerConfig.MysqlInfo.User, global.ServerConfig.MysqlInfo.Password, global.ServerConfig.MysqlInfo.Host, global.ServerConfig.MysqlInfo.Port, global.ServerConfig.MysqlInfo.Name)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger, NamingStrategy: schema.NamingStrategy{SingularTable: true}})

	if err != nil {
		panic(err)
	}
}
