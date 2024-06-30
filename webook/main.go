package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	initViper()
	keys := viper.AllKeys()
	fmt.Println(keys)

	//otel

	//web
	app := InitWebServer()
	server := app.Web
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
	server.Run(":8080")
}
