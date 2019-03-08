package service

import (
	"context"
	"strings"
	"time"

	"github.com/emurmotol/project/user_api/pkg/utils"
)

// UserApiService describes the service.
type UserApiService interface {
	GetByUsername(ctx context.Context, username string) (user User, err error)
	CreateUser(ctx context.Context, username string, email string, password string, role string) (user User, err error)
	GetUserForAuth(ctx context.Context, username string) (user User, err error)
}

type User struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at"` // 2006-01-02 15:04:05.000000
	UpdatedAt time.Time  `json:"updated_at"` // 2006-01-02 15:04:05.000000
	DeletedAt *time.Time `json:"deleted_at"` // 2006-01-02 15:04:05.000000
}

func (u *User) BeforeCreate() (err error) {
	u.Email = strings.ToLower(u.Email)

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashPassword
	return nil
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

	user = User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}
	if err := db.Create(&user).Error; err != nil {
		return User{}, err
	}

	enforcer := utils.GetCasbinEnforcer(ctx)
	enforcer.AddPolicy(role, "object", "read")
	return user, nil
}

func (b *basicUserApiService) GetUserForAuth(ctx context.Context, username string) (user User, err error) {
	db := utils.GetDB(ctx)

	if err := db.Where(User{Username: username}).Or(User{Email: username}).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
