package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"go_srvs/user_srv/global"
	"go_srvs/user_srv/model"
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

	// 迁移表
	//_ = global.DB.AutoMigrate(&model.User{})
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
	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("user%d", i),
			Mobile:   fmt.Sprintf("1311222222%d", i),
			PassWord: pwd,
		}
		global.DB.Save(&user)
	}
}
