package ginx

import (
	"7day/webook/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var L logger.Logger

func WarpToken[C jwt.Claims](fn func(ctx *gin.Context, c C) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val, ok := ctx.Get("claims")
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c, ok := val.(C)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		result, err := fn(ctx, c)
		if err != nil {
			L.Error("业务逻辑错误",
				logger.String("path", ctx.Request.URL.Path),
				logger.String("route", ctx.FullPath()),
				logger.Error(err),
			)
		}
		ctx.JSON(http.StatusOK, result)
	}
}

func WarpBodyAndToken[Req any, C jwt.Claims](fn func(ctx *gin.Context, req Req, c C) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		if err := ctx.Bind(&req); err != nil {
			return
		}
		val, ok := ctx.Get("claims")
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c, ok := val.(C)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		result, err := fn(ctx, req, c)
		if err != nil {
			L.Error("业务逻辑错误",
				logger.String("path", ctx.Request.URL.Path),
				logger.String("route", ctx.FullPath()),
				logger.Error(err),
			)
		}
		ctx.JSON(http.StatusOK, result)
	}
}
