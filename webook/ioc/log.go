package ioc

import (
	"7day/webook/pkg/logger"
	"fmt"
	"go.uber.org/zap"
)

func InitLogger() logger.Logger {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintln("init logger error:", err))
	}
	return logger.NewZapLogger(l)
}
