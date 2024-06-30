package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	AtKey = []byte("95osj3fUD7fo0mlYdDbncXz4VD2igvf0")
	RtKey = []byte("95osj3fUD7fo0mlYdDbncXz4VD2igvfx")
)

type RedisJWTHandler struct {
	cmd redis.Cmdable
}

// 调用 refresh 和 user token的
func (r *RedisJWTHandler) SerLoginToken(ctx *gin.Context, uid int64) error {
	Ssid := uuid.New().String()
	err := r.SetJWTToken(ctx, uid, Ssid)
	if err != nil {
		return err
	}
	err = r.setRefreshToken(ctx, uid, Ssid)
	if err != nil {
		return err
	}
	return nil
}
func (r *RedisJWTHandler) setRefreshToken(ctx *gin.Context, uid int64, Ssid string) error {
	refreshClaims := RefreshClaims{
		Uid:  uid,
		Ssid: Ssid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 7)),
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS512, refreshClaims)
	token, err := claims.SignedString(RtKey)
	if err != nil {
		return err
	}
	ctx.Header("x-refresh-token", token)
	return nil
}

func (r *RedisJWTHandler) SetJWTToken(ctx *gin.Context, uid int64, Ssid string) error {
	claims := UserClaims{
		Id:   uid,
		Ssid: Ssid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(AtKey)
	if err != nil {
		return err
	}
	ctx.Header("x-jwt-token", tokenString)
	return nil
}

// 退出登录
func (r *RedisJWTHandler) ClearToken(ctx *gin.Context) error {
	ctx.Header("x-jwt-token", "")
	ctx.Header("x-refresh_token", "")
	claims := ctx.MustGet("claims").(UserClaims)
	return r.cmd.Set(ctx, fmt.Sprintf("users:Ssid:%s", claims.Ssid), "", time.Hour*24*7).Err()

}

// 检查是否退出
func (r *RedisJWTHandler) CheckSession(ctx *gin.Context, Ssid string) error {
	val, err := r.cmd.Exists(ctx, "users:Ssid:%s", Ssid).Result()
	switch err {
	case redis.Nil:
		return nil
	case nil:
		if val == 0 {
			return nil
		}
	default:
		return err
	}
	return nil
}

func (h *RedisJWTHandler) ExtractToken(ctx *gin.Context) string {
	// 我现在用 JWT 来校验
	tokenHeader := ctx.GetHeader("Authorization")
	return tokenHeader

}
func NewRedisJWTHandler(cmd redis.Cmdable) JWTHandler {
	return &RedisJWTHandler{cmd: cmd}
}
