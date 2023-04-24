package initialize

import (
	"fmt"
	"go.uber.org/zap"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"mxshop-api/user-web/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user-web/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("user-web/%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	// 读取yaml文件
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息: %v", global.ServerConfig)

	// 动态监听变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置信息变化：%s", e.Name)
		err := v.ReadInConfig()
		if err != nil {
			panic(err)
		}
		err = v.Unmarshal(global.ServerConfig)
		if err != nil {
			panic(err)
		}
		zap.S().Infof("配置信息: %v", global.ServerConfig)
	})
}
