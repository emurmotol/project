package service

import (
	"context"

	"github.com/emurmotol/project/user_api/pkg/utils"
)

// UserApiService describes the service.
type UserApiService interface {
	GetByUsername(ctx context.Context, username string) (user User, err error)
	CreateUser(ctx context.Context, username string, email string, password string, role string) (user User, err error)
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type basicUserApiService struct{}

func (b *basicUserApiService) GetByUsername(ctx context.Context, username string) (user User, err error) {
	db := utils.GetDB(ctx)

	if err := db.First(&user, User{Username: username}).Error; err != nil {
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
	db := utils.GetDB(ctx)

	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return User{}, err
	}

	user = User{
		Username: username,
		Email:    email,
		Password: hashPassword,
		Role:     role,
	}
	if err := db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
