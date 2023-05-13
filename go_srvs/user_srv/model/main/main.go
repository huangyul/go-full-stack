package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"io"
	"strings"
)

// 生成md5算法
func genMd5(code string) string {
	md := md5.New()
	_, _ = io.WriteString(md, code)
	return hex.EncodeToString(md.Sum(nil))
}

// 连接数据库

func main() {
	//dsn := "root:123456@tcp(47.106.214.127)/go_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	//
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags),
	//	logger.Config{
	//		SlowThreshold: time.Second,
	//		LogLevel:      logger.Info,
	//		Colorful:      true,
	//	},
	//)
	//
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger, NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	//
	//if err != nil {
	//	panic(err)
	//}

	// 迁移表
	//_ = db.AutoMigrate(&model.User{})
	// Using the default options
	//salt, encodedPwd := password.Encode("generic password", nil)
	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println(check) // true

	// Using custom options
	options := &password.Options{16, 100, 32, sha256.New}
	salt, encodedPwd := password.Encode("generic password", options)
	pwd := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(pwd)
	pwdInfo := strings.Split(pwd, "$")
	check := password.Verify("generic password", pwdInfo[2], pwdInfo[3], options)
	fmt.Println(check) // true
}
