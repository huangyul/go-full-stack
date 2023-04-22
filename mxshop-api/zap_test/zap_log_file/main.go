package main

import "go.uber.org/zap"

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./myproject.log",
		"stderr",
	}
	return cfg.Build()
}

func main() {
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	su := logger.Sugar()
	defer su.Sync()
	url := "http://"
	su.Infow("failed to fetch", zap.String("url", url))
}
