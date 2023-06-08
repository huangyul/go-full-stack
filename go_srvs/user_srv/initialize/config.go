package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go_srvs/user_srv/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	isDebug := GetEnvInfo("GO_DEBUG")
	configFileName := fmt.Sprintf("user_srv/config-pro.yaml")
	if isDebug {
		configFileName = fmt.Sprintf("user_srv/config-debug.yaml")
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}

}
