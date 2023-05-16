package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-api/user-web/global"
	"go.uber.org/zap"
)

func InitConfig() {
	// 获取环境变量， 判断使用那个配置文件
	viper.AutomaticEnv()
	debug := viper.GetBool("GO_DEBUG")
	configFileName := "user-web/config-pro.yaml"
	if debug {
		configFileName = "user-web/config-debug.yaml"
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}

	zap.S().Infof("配置信息：%v", global.ServerConfig)

	// viper 动态监听配置文件的变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件产生变化: %s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
		zap.S().Infof("配置信息: %v", global.ServerConfig)
	})
}
