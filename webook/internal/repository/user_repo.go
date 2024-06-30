package repository

import (
	"7day/webook/internal/domain"
	"7day/webook/internal/repository/dao/user"
	"7day/webook/pkg/logger"
	"context"
	"database/sql"
	"time"
)

type UserRepo interface {
	Create(ctx context.Context, user domain.User) error
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	FindById(ctx context.Context, id int64) (domain.User, error)
	FindByPhone(ctx context.Context, phone string) (domain.User, error)
}

func NewUserRepoIml(dao dao.UserDAO, l logger.Logger) UserRepo {
	return &userRepoIml{
		dao: dao,
		l:   l,
	}
}

type userRepoIml struct {
	dao dao.UserDAO
	l   logger.Logger
}

func (repo *userRepoIml) FindByPhone(ctx context.Context, phone string) (domain.User, error) {
	u, err := repo.dao.FindByPhone(ctx, phone)
	return repo.entityToDomain(u), err
}

func (repo *userRepoIml) FindById(ctx context.Context, id int64) (domain.User, error) {
	user, err := repo.dao.FindById(ctx, id)
	return repo.entityToDomain(user), err
}

func (repo *userRepoIml) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := repo.dao.GetByEmail(ctx, email)
	return repo.entityToDomain(u), err
}

func (repo *userRepoIml) Create(ctx context.Context, user domain.User) error {
	return repo.dao.Create(ctx, repo.domainToEntity(user))
}

func (repo *userRepoIml) entityToDomain(user dao.User) domain.User {
	return domain.User{
		Id:       user.Id,
		Email:    user.Email.String,
		Phone:    user.Phone.String,
		Password: user.Password,
		Ctime:    time.UnixMilli(user.Ctime),
	}
}

func (repo *userRepoIml) domainToEntity(user domain.User) dao.User {
	return dao.User{
		Id: user.Id,
		Email: sql.NullString{
			String: user.Email,
			Valid:  user.Email != "",
		},
		Password: user.Password,
		Phone: sql.NullString{
			String: user.Phone,
			Valid:  user.Phone != "",
		},
		Ctime: user.Ctime.UnixMilli(),
	}
}
