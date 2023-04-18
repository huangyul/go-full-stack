package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"io"
	"strings"
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

	//
	//_ = db.AutoMigrate(&model.User{})

	// Using custom options
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(newPassword)
	fmt.Println(len(newPassword))

	passwordInfo := strings.Split(newPassword, "$")
	check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	fmt.Println(check) // true
}
