//go:build wireinject

package main

import (
	"github.com/google/wire"
	sms "webook/sms/ioc"
	user "webook/user/ioc"
)

var third = wire.NewSet(
	bff.InitRedis,
	bff.InitGorm,
)
var bff = wire.NewSet(
	bff.InitWebServer,
)

func InitApp() *App {
	wire.Build(
		bff,
		third,
		user.InitUser,
		sms.InitSMS,
		wire.Struct(new(App), "*"),
	)
	return new(App)
}
