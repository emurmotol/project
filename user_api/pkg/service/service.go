package service

import (
	"context"
	"errors"

	"github.com/emurmotol/project/user_api/pkg/utils"
	"github.com/jinzhu/gorm"
)

// UserApiService describes the service.
type UserApiService interface {
	GetByUsername(ctx context.Context, username string) (user User, err error)
	CreateUser(ctx context.Context, username string, email string, password string, role string) (user User, err error)
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type basicUserApiService struct{}

func (b *basicUserApiService) GetByUsername(ctx context.Context, username string) (user User, err error) {
	db, ok := ctx.Value(utils.DBContextKey).(*gorm.DB)
	if !ok {
		return User{}, errors.New("db not ok")
	}

	if err := db.Find(&user, User{Username: username}).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

// NewBasicUserApiService returns a naive, stateless implementation of UserApiService.
func NewBasicUserApiService() UserApiService {
	return &basicUserApiService{}
}

// New returns a UserApiService with all of the expected middleware wired in.
func New(middleware []Middleware) UserApiService {
	var svc UserApiService = NewBasicUserApiService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicUserApiService) CreateUser(ctx context.Context, username string, email string, password string, role string) (user User, err error) {
	db, ok := ctx.Value(utils.DBContextKey).(*gorm.DB)
	if !ok {
		return User{}, errors.New("db not ok")
	}

	user = User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}
	if err := db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
