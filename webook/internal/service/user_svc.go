package service

import (
	"7day/webook/internal/domain"
	"7day/webook/internal/repository"
	"7day/webook/internal/repository/dao/user"
	"7day/webook/internal/service/sms"
	"7day/webook/pkg/logger"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

//go:generate mockgen -source=./user_svc.go -package=svcmocks -destination=mocks/user_svc.mock.go
type UserSVC interface {
	Signup(ctx context.Context, user domain.User) error
	Login(ctx context.Context, user domain.User) (domain.User, error)
	Profile(ctx context.Context, id int64) (domain.User, error)
	SendSMS(ctx context.Context, phone string) error
	FindOrCreate(ctx context.Context, user domain.User) (domain.User, error)
	Verify(ctx context.Context, phone string, code string) error
}

func NewUserSvcIml(repo repository.UserRepo, l logger.Logger, sms sms.SMS, codeRepo repository.CodeRepo) UserSVC {
	return &UserSvcIml{
		repo:     repo,
		repoCode: codeRepo,
		sms:      sms,
		l:        l,
		biz:      "user",
	}
}

type UserSvcIml struct {
	repo     repository.UserRepo
	repoCode repository.CodeRepo
	sms      sms.SMS
	l        logger.Logger
	biz      string
}

func (svc *UserSvcIml) FindOrCreate(ctx context.Context, user domain.User) (domain.User, error) {
	user, err := svc.repo.FindByPhone(ctx, user.Phone)
	// nil 会走这里,其他错误也会走这里,没发现记录会继续往下走
	if err != dao.ErrNotFindRecord {
		return user, err
	}
	err = svc.repo.Create(ctx, user)
	if err != nil || err != dao.ErrDataDuplicate {
		return domain.User{}, err
	}
	return svc.repo.FindByPhone(ctx, user.Phone)
}

func (svc *UserSvcIml) Verify(ctx context.Context, phone string, code string) error {
	return svc.repoCode.Verify(ctx, svc.genertLoginSMS(svc.biz, phone), code)
}

func (svc *UserSvcIml) SendSMS(ctx context.Context, phone string) error {
	random := rand.Intn(1000000)
	code := fmt.Sprintf("%06d", random)
	key := svc.genertLoginSMS(svc.biz, phone)

	err := svc.repoCode.Store(ctx, key, code)
	if err != nil {
		return err
	}
	err = svc.sms.Send(ctx, phone, svc.biz, code)
	if err != nil {
		svc.l.Error("发送验证码失败")
		return err
	}
	return nil
}

func (u *UserSvcIml) Profile(ctx context.Context, id int64) (domain.User, error) {
	return u.repo.FindById(ctx, id)
}

func (u *UserSvcIml) Login(ctx context.Context, user domain.User) (domain.User, error) {
	resultU, err := u.repo.GetByEmail(ctx, user.Email)
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(resultU.Password), []byte(user.Password))
	if err != nil {
		return domain.User{}, errors.New("密码错误")
	}
	return resultU, err
}

func (u *UserSvcIml) Signup(ctx context.Context, user domain.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		u.l.Error("密码加密失败")
		return err
	}
	user.Password = string(password)
	err = u.repo.Create(ctx, user)
	return err
}

func (svc *UserSvcIml) genertLoginSMS(biz, phone string) string {

	return fmt.Sprintf("%v:%v", svc.biz, phone)
}
