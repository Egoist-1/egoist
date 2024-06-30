package cache

import (
	"7day/webook/pkg/logger"
	"context"
	_ "embed"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

//go:embed lua/storeLoginSMS.lua
var loginSMS string

//go:embed lua/verify.lua
var verify string

type CodeCache interface {
	Set(ctx context.Context, key, val string) error
	Verify(ctx context.Context, key string, val string) error
}

func NewCodeCacheImpl(redis redis.Cmdable, l logger.Logger) CodeCache {
	return &CodeCacheImpl{
		redis: redis,
		l:     l,
	}
}

type CodeCacheImpl struct {
	redis redis.Cmdable
	l     logger.Logger
}

func (cache CodeCacheImpl) Verify(ctx context.Context, key string, val string) error {
	result, err := cache.redis.Eval(ctx, verify, []string{key}, val).Int()
	if err != nil {
		return err
	}
	switch result {
	case 1:
		return nil
	case 0:
		return errors.New("400 输入错误,请重新输入")
	case -1:
		return errors.New("400 频繁输入,请重新获取验证码")
	default:
		return err
	}
}

func (cache CodeCacheImpl) Set(ctx context.Context, key, val string) error {
	//这里就要用lua脚本判断是否 发送频繁了
	result, err := cache.redis.Eval(ctx, loginSMS, []string{key}, val).Int()
	if err != nil {
		return err
	}
	switch result {
	case 0:
		return errors.New("400 发送太频繁,请稍后重试")
	case 1:
		return nil
	default:
		return errors.New("500 系统错误")
	}
}
