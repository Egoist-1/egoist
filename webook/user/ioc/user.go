package ioc

import (
	"github.com/google/wire"
	"webook/_bff/web"
	repository2 "webook/user/_internal/repository"
	"webook/user/_internal/repository/cache"
	"webook/user/_internal/repository/dao"
	service2 "webook/user/_internal/service"
)

var InitUser = wire.NewSet(
	web.NewUserHandle,
	service2.NewUserServiceImpl,
	repository2.NewUserRepo,
	dao.NewUserDao,
	service2.NewCodeService,
	repository2.NewCodeRepo,
	cache.NewCodeCacheRedis,
)
