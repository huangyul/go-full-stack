package global

import (
	"go_srvs/user_srv/config"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
)

func init() {

}
