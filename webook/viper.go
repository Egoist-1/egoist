package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func initViper() {
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./webook/config")

	//实时监听配置文件
	//只能告诉你文件变了,不能告诉你那些文件内容变了
	// 读取配置到 viper 里面，或者你可以理解为加载到内存里面
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintln("配置文件错误:", err))
	}
}
