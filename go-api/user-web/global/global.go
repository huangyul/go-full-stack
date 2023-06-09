package global

import (
	ut "github.com/go-playground/universal-translator"
	"go-api/user-web/config"
	"go-api/user-web/proto"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	Trans        ut.Translator

	UserSrvClient proto.UserClient
)
