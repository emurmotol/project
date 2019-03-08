package service

import (
	"context"
	"strings"
	"time"

	"github.com/emurmotol/project/user/pkg/utils"
	stdgrpc "google.golang.org/grpc"
)

// UserService describes the service.
type UserService interface {
	GetByUsername(ctx context.Context, username string) (user User, err error)
	CreateUser(ctx context.Context, username string, email string, password string, role string) (user User, err error)
	GetUserForAuth(ctx context.Context, username string) (user User, err error)
}

type CasbinRule struct {
	PType string `json:"p_type"`
	V0    string `json:"v0"`
	V1    string `json:"v1"`
	V2    string `json:"v2"`
	V3    string `json:"v3"`
	V4    string `json:"v4"`
	V5    string `json:"v5"`
}

func (cr CasbinRule) TableName() string {
	return "casbin_rule"
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

type basicUserService struct{}

func (b *basicUserService) GetByUsername(ctx context.Context, username string) (user User, err error) {
	db := utils.GetDB(ctx)

	if err := db.First(&user, User{Username: username}).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicUserService() UserService {
	return &basicUserService{}
}

// New returns a UserService with all of the expected middleware wired in.
func New(middleware []Middleware) UserService {
	var svc UserService = NewBasicUserService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicUserService) CreateUser(ctx context.Context, username string, email string, password string, role string) (user User, err error) {
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
	enforcer.AddPolicy(role, stdgrpc.ServerTransportStreamFromContext(ctx).Method(), "read")
	return user, nil
}

func (b *basicUserService) GetUserForAuth(ctx context.Context, username string) (user User, err error) {
	db := utils.GetDB(ctx)

	if err := db.Where(User{Username: username}).Or(User{Email: username}).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
