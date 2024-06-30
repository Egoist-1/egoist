package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrNotFindRecord = gorm.ErrRecordNotFound
	ErrDataDuplicate = errors.New("400 用户已存在")
)

//go:generate mockgen -source=./user_dao.go -package=daomocks -destination=mocks/user_dao.mock.go
type UserDAO interface {
	Create(ctx context.Context, user User) error
	GetByEmail(ctx context.Context, email string) (User, error)
	FindById(ctx context.Context, id int64) (User, error)
	FindByPhone(ctx context.Context, phone string) (User, error)
}

func NewUserDAOIml(db *gorm.DB) UserDAO {
	return &UserDAOIml{db: db}
}

type UserDAOIml struct {
	db *gorm.DB
}

func (dao *UserDAOIml) FindByPhone(ctx context.Context, phone string) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Where("phone = ?", phone).First(user).Error
	return user, err
}

func (dao *UserDAOIml) FindById(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Model(User{}).Where("id = ?", id).First(&u).Error
	return u, err
}

func (dao *UserDAOIml) GetByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Model(User{}).Where("email = ?", email).First(&u).Error
	return u, err
}

func (dao *UserDAOIml) Create(ctx context.Context, user User) error {
	now := time.Now().UnixMilli()
	user.Utime = now
	user.Ctime = now
	err := dao.db.WithContext(ctx).Create(&user).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			return errors.New("邮箱冲突")
		}
	}
	return err
}

type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	//唯一索引可以有多个空值 但是不能有多个""
	//所以string类型在使用手机号或者邮箱注册时会有""的情况
	//这时使用sql.nullString
	Email    sql.NullString `gorm:"unique"`
	Password string
	Phone    sql.NullString `gorm:"unique"`
	//存储的毫秒值
	Ctime int64
	Utime int64
}

