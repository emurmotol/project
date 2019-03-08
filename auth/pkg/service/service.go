package service

import (
	"context"
	"errors"
	"time"

	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/emurmotol/project/auth/pkg/grpc/pb"
	"github.com/emurmotol/project/auth/pkg/utils"
	"github.com/spf13/viper"
)

// AuthService describes the service.
type AuthService interface {
	Login(ctx context.Context, username string, password string) (accessToken string, tokenType string, expiresAt int64, err error)
	Restricted(ctx context.Context) (claims *utils.JWTClaims, err error)
	HealthCheck(ctx context.Context) (status string, err error)
}

type basicAuthService struct{}

func (b *basicAuthService) Login(ctx context.Context, username string, password string) (accessToken string, tokenType string, expiresAt int64, err error) {
	userClient := utils.GetUserClient(ctx)
	reply, err := userClient.GetUserForAuth(ctx, &pb.GetUserForAuthRequest{Username: username})
	if err != nil {
		return "", "", int64(0), err
	}
	user := reply.User

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", "", int64(0), errors.New("wrong password")
	}
	now := time.Now()
	expiresAt = now.Local().Add(time.Second * viper.GetDuration("JWT_EXPIRES_IN_SECONDS")).Unix()
	claims := &utils.JWTClaims{
		StandardClaims: stdjwt.StandardClaims{
			Audience:  "0.0.0.0",
			ExpiresAt: expiresAt,
			Id:        user.ID,
			IssuedAt:  now.Unix(),
			Issuer:    "0.0.0.0",
			NotBefore: 0,
			Subject:   user.Role,
		},
	}
	token, err := utils.GenerateJWTToken(claims, viper.GetString("JWT_PRIVATE_KEY"))
	if err != nil {
		return "", "", int64(0), err
	}
	return token, viper.GetString("JWT_TOKEN_TYPE"), expiresAt, nil
}

// NewBasicAuthService returns a naive, stateless implementation of AuthService.
func NewBasicAuthService() AuthService {
	return &basicAuthService{}
}

// New returns a AuthService with all of the expected middleware wired in.
func New(middleware []Middleware) AuthService {
	var svc AuthService = NewBasicAuthService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicAuthService) Restricted(ctx context.Context) (claims *utils.JWTClaims, err error) {
	return utils.GetClaims(ctx), nil
}

func (b *basicAuthService) HealthCheck(ctx context.Context) (status string, err error) {
	return "OK", nil
}
