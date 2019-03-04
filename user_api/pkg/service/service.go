package service

import (
	"context"

	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

type basicUserApiService struct {
	db *gorm.DB
}

func (b *basicUserApiService) GetByUsername(ctx context.Context, username string) (user User, err error) {
	if err := b.db.Find(&user, User{Username: username}).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

// NewBasicUserApiService returns a naive, stateless implementation of UserApiService.
func NewBasicUserApiService() UserApiService {
	db, err := gorm.Open("postgres", viper.GetString("POSTGRES_DATABASE"))
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	return &basicUserApiService{
		db: db,
	}
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
	user = User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}

	if err := b.db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
