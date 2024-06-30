package web

import (
	"7day/webook/internal/domain"
	"7day/webook/internal/service"
	ijwt "7day/webook/internal/web/jwt"
	"errors"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ handler = (*UserHandler)(nil)

type UserHandler struct {
	emailMatch *regexp.Regexp
	svc        service.UserSVC
	ijwt.JWTHandler
}

func NewUserHandler(svc service.UserSVC, Ujwt ijwt.JWTHandler) *UserHandler {
	const (
		emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	return &UserHandler{
		emailMatch: emailExp,
		svc:        svc,
		JWTHandler: Ujwt,
	}
}

func (u *UserHandler) RegisterRouter(server *gin.Engine) {
	ug := server.Group("/users")
	ug.POST("/signup", u.signup)
	ug.POST("/login", u.login)
	//获取详情
	ug.GET("/profile", u.profile)
	ug.POST("/edit")
	ug.POST("/logout")
	ug.POST("/login_sms/code/send", u.sendSMS)
	ug.POST("/login_sms", u.loginSMS)
	ug.POST("/refresh_token")
}

func (u *UserHandler) signup(ctx *gin.Context) {
	type req struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	var r req
	err := ctx.Bind(&r)
	if err != nil {
		return
	}
	ok, err := u.emailMatch.MatchString(r.Email)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 500,
			Msg:  "系统错误",
			Data: nil,
		})
		return
	}
	if !ok {
		ctx.JSON(http.StatusOK, Result{
			Code: 400,
			Msg:  "邮箱格式错误",
			Data: nil,
		})
		return
	}
	if r.Password != r.ConfirmPassword {
		ctx.JSON(http.StatusOK, Result{
			Code: 400,
			Msg:  "密码不匹配",
			Data: nil,
		})
		return
	}
	err = u.svc.Signup(ctx, domain.User{
		Email:    r.Email,
		Password: r.Password,
	})
	switch err {
	case errors.New("邮箱冲突"):
		ctx.JSON(http.StatusOK, Result{
			Code: 400,
			Msg:  "邮箱已注册",
			Data: nil,
		})
		return
	case nil:
		ctx.JSON(http.StatusOK, Result{
			Code: 400,
			Msg:  "注册成功",
			Data: nil,
		})
		return
	default:
		ctx.JSON(http.StatusOK, Result{
			Code: 400,
			Msg:  "系统错误",
			Data: nil,
		})
		return
	}
}

func (u *UserHandler) login(ctx *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	err := ctx.Bind(&req)
	if err != nil {
		return
	}
	user, err := u.svc.Login(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	switch err {
	case nil:
		err := u.JWTHandler.SerLoginToken(ctx, user.Id)
		if err != nil {
			ctx.JSON(http.StatusOK, Result{
				Code: 5,
				Msg:  "系统错误",
				Data: nil,
			})
		}
		ctx.JSON(http.StatusOK, Result{
			Code: 2,
			Msg:  "登录成功",
			Data: nil,
		})
	case errors.New("密码错误"):
		ctx.JSON(http.StatusOK, Result{
			Code: 4,
			Msg:  "密码错误",
			Data: nil,
		})
	default:
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
			Data: nil,
		})
	}

}

func (web *UserHandler) profile(ctx *gin.Context) {
	result := Result{
		Code: 2,
		Msg:  "profile success",
		Data: nil,
	}
	claims, _ := ctx.Get("claims")
	c, ok := claims.(ijwt.UserClaims)
	if !ok {
		result.Code = 5
		result.Msg = "系统错误"
		ctx.JSON(http.StatusOK, result)
		return
	}
	user, err := web.svc.Profile(ctx, c.Id)
	if err != nil {
		result.Code = 5
		result.Msg = "系统错误"
		ctx.JSON(http.StatusOK, result)
		return
	}
	result.Data = user
	ctx.JSON(http.StatusOK, result)

}
func (web *UserHandler) sendSMS(ctx *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
	}
	var req Req
	err := ctx.Bind(&req)
	if err != nil {
		return
	}
	//没有校验密码格式
	err = web.svc.SendSMS(ctx, req.Phone)
	ResultJSON(ctx, "发送成功", nil, err, func() {})
}

func (web *UserHandler) loginSMS(ctx *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	var req Req
	err := ctx.Bind(&req)
	if err != nil {
		return
	}
	err = web.svc.Verify(ctx, req.Phone, req.Code)
	if err != nil {
		ResultJSON(ctx, "", nil, err, func() {

		})
	}
	u, err := web.svc.FindOrCreate(ctx, domain.User{
		Phone: req.Phone,
	})
	ResultJSON(ctx, "登录成功", nil, err, func() {
		err := web.JWTHandler.SerLoginToken(ctx, u.Id)
		if err != nil {
			ctx.JSON(http.StatusOK, Result{
				Code: 5,
				Msg:  "系统错误",
				Data: nil,
			})
			return
		}
	})

}
