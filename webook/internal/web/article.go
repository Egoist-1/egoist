package web

import (
	"7day/webook/internal/domain"
	"7day/webook/internal/service"
	jwt "7day/webook/internal/web/jwt"
	"7day/webook/pkg/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewArticle(svc service.ArticleSVC) *ArticleHandler {
	return &ArticleHandler{
		svc: svc,
	}
}

func (a *ArticleHandler) RegisterRouter(server *gin.Engine) {
	ag := server.Group("article")
	ag.POST("edit", a.edit)
	ag.POST("publish", a.publish)
	ag.POST("withdraw", a.withdraw)
	ag.POST("list", a.list)
	ag.POST("/detail/:id", a.detail)
	pub := ag.Group("pub")
	pub.POST("/detail/id")
	pub.POST("/like",a.like)

}

type ArticleHandler struct {
	svc service.ArticleSVC
}

func (a *ArticleHandler) edit(ctx *gin.Context) {

	var req artReq
	err := ctx.Bind(&req)
	if err != nil {
		return
	}
	claims := ctx.MustGet("claims").(*jwt.UserClaims)
	id, err := a.svc.Save(ctx, domain.Article{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Content,
		Author: domain.Author{
			Id: claims.Id,
		},
	})
	ResultJSON(ctx, "成功", id, err, func() {})
}

func (a *ArticleHandler) publish(ctx *gin.Context) {
	var req artReq
	err := ctx.Bind(&req)
	if err != nil {
		return
	}
	claims := ctx.MustGet("claims").(*jwt.UserClaims)
	id, err := a.svc.Publish(ctx, req.toDomain(claims.Id))
	ResultJSON(ctx, "发布成功", id, err, func() {})
}

func (a *ArticleHandler) withdraw(ctx *gin.Context) {
	var art artReq
	err := ctx.Bind(&art)
	if err != nil {
		return
	}
	claims := ctx.MustGet("claims").(*jwt.UserClaims)
	err = a.svc.Withdraw(ctx, domain.Article{
		Id: art.Id,
		Author: domain.Author{
			Id: claims.Id,
		},
	})
	ResultJSON(ctx, "修改成功", nil, err, func() {})
}

func (a *ArticleHandler) list(ctx *gin.Context) {
	var req reqList
	err := ctx.Bind(&req)
	if err != nil {
		return
	}
	claims := ctx.MustGet("claims")
	user, ok := claims.(*jwt.UserClaims)
	if !ok {
		return
	}
	arts, err := a.svc.List(ctx, user.Id, req.Offset, req.Limit)
	vos := tools.Map[domain.Article, ArticleVO](arts, func(idx int, s domain.Article) ArticleVO {
		return ArticleVO{
			Id:       s.Id,
			Title:    s.Title,
			Abstract: s.Abstract(),
			Ctime:    s.Ctime,
			Utime:    s.Utime,
			Status:   s.Status,
		}
	})
	ResultJSON(ctx, "ok", vos, err, func() {})
}

func (a *ArticleHandler) detail(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return
	}
	detail, err := a.svc.GetDetail(ctx, int64(id))
	claims := ctx.MustGet("claims")
	u := claims.(*jwt.UserClaims)
	if err != nil || detail.Id != u.Id {
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
		return
	}
	result := ArticleVO{
		Id:       detail.Id,
		Title:    detail.Title,
		Content:  detail.Content,
		Ctime:    detail.Ctime,
		Utime:    detail.Utime,
		Status:   detail.Status,
	}
	ResultJSON(ctx,"ok", result, err, func() {})
}

func (a *ArticleHandler) like(ctx *gin.Context) {
	
	
}

func (r *artReq) toDomain(id int64) domain.Article {
	return domain.Article{
		Id:      r.Id,
		Title:   r.Title,
		Content: r.Content,
		Author: domain.Author{
			Id: id,
		},
	}

}
