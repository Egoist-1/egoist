//wireinject

package main

import (
	"github.com/google/wire"
	ioc2 "webook/_bff/ioc"
	"webook/_bff/web"
	"webook/sms/_internal/service/sms/memory"
	repository2 "webook/user/_internal/repository"
	"webook/user/_internal/repository/cache"
	"webook/user/_internal/repository/dao"
	service2 "webook/user/_internal/service"
)

var email = wire.NewSet()
var user = wire.NewSet(
	web.NewUserHandle,
	service2.NewUserServiceImpl,
	repository2.NewUserRepo,
	dao.NewUserDao,
)
var code = wire.NewSet(
	service2.NewCodeService,
	repository2.NewCodeRepo,
	cache.NewCodeCacheRedis,
)
var sms = wire.NewSet(
	memory.NewMemory,
)

func InitApp() *App {
	wire.Build(
		user,
		code,
		sms,
		email,
		ioc2.InitWebServer,
		ioc2.InitGorm,
		wire.Struct(new(App), "*"),
	)
	return new(App)
}
