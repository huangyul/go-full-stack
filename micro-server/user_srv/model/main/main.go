package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"micro-server/user_srv/model"
	"os"
	"strings"
	"time"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, err := io.WriteString(Md5, code)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {

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
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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

	//_ = DB.AutoMigrate(&model.User{})

	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(newPassword)
	fmt.Println(len(newPassword))

	passwordInfo := strings.Split(newPassword, "$")
	check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	fmt.Println(check) // true
	// 创建10个用户
	for i := 0; i < 10; i++ {
		var user = model.User{
			NickName: fmt.Sprintf("huang%d", i),
			Mobile:   fmt.Sprintf("1111111111%d", i),
			Password: newPassword,
		}
		DB.Create(&user)
	}

}
