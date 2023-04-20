package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction() // 生产环境
	defer logger.Sync()
	url := "https://"
	// 使用sugar
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL", "url", url, "attempt", 3)
	sugar.Infof("Failed to fetch URL: %s", url)
	// 直接使用logger
	logger.Info("failed to fetch", zap.String("url", url), zap.Int("attempt", 3))
}
