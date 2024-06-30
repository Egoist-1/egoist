package web

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}


func ResultJSON(ctx *gin.Context, success string, data any, err error, fn func()) {
	if err == nil {
		fn()
		ctx.JSON(http.StatusOK, Result{
			Code: 2,
			Msg:  success,
			Data: data,
		})
		return
	}
	preStr := []byte(err.Error())[0:3]
	resultErr := []byte(err.Error())[4:]
	err = errors.New(string(resultErr))
	switch string(preStr) {
	case "400":
		ctx.JSON(http.StatusOK, Result{
			Code: 4,
			Msg:  err.Error(),
			Data: nil,
		})
	case "500":
		fmt.Println("这是错误信息" + err.Error())
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
			Data: nil,
		})
	default:
		fmt.Println("这是错误信息" + err.Error())
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
			Data: nil,
		})
	}
}
