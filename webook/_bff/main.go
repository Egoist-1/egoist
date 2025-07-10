package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	initViper()
	initLogger()
	//pprof
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	//prometheus
	go func() {
		fmt.Printf("prometheus start")
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":9091", nil)
		fmt.Printf("prometheus err :%v", err)
	}()
	app := InitApp()
	app.web.Run(":3001")
}

func initViper() {
	viper.SetConfigName("dev")      // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // path to look for the config file in
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
func initLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	zap.NewAtomicLevelAt(zapcore.DebugLevel)
	// 设置了全局的 logger，
	// 你在你的代码里面就可以直接使用 zap.XXX 来记录日志
	zap.ReplaceGlobals(logger)
}
