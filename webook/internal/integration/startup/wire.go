//go:build wireinject

package startup

import (
	"7day/webook/internal/repository"
	"7day/webook/internal/repository/cache"
	"7day/webook/internal/repository/dao/article"
	dao "7day/webook/internal/repository/dao/user"
	"7day/webook/internal/service"
	"7day/webook/internal/service/sms"
	"7day/webook/internal/web"
	"7day/webook/internal/web/jwt"
	"7day/webook/ioc"
	"github.com/google/wire"
)

// web
var webConf = wire.NewSet(
	web.NewArticle,
	web.NewUserHandler,
)

// service
var serviceConf = wire.NewSet(
	service.NewUserSvcIml,
	service.NewArticleSVCImpl,
)

// SMS  短信
var SMS = wire.NewSet(
	sms.NewMemorySMS,
)

// Repo
var repo = wire.NewSet(
	repository.NewUserRepoIml,
	repository.NewCodeRepoImpl,
	repository.NewArticleRepoImpl,
)

// DAO
var daoConfi = wire.NewSet(
	dao.NewUserDAOIml,
	article.NewArticleDaoImpl,
)

// cache
var cacheConfig = wire.NewSet(
	cache.NewCodeCacheImpl,
	cache.NewArticleRedis,
)

// gin 初始化
var ginSet = wire.NewSet(
	jwt.NewRedisJWTHandler,
	ioc.InitWebServer,
	ioc.InitMiddlewares,
)

// 需要的 中间件
var middleware = wire.NewSet(
	ioc.InitDB,
	ioc.InitRedis,
	ioc.InitLogger,
)

func InitWebServer() *App {
	wire.Build(
		webConf,
		serviceConf,
		SMS,
		repo,
		daoConfi,
		cacheConfig,
		ginSet,
		middleware,
		wire.Struct(new(App), "*"),
	)

	return new(App)
}
