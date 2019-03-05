package service

import (
	"context"
	"errors"
	"time"

	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/emurmotol/project/auth_api/pkg/grpc/pb"
	"github.com/emurmotol/project/auth_api/pkg/utils"
	"github.com/spf13/viper"
)

// AuthApiService describes the service.
type AuthApiService interface {
	Login(ctx context.Context, username string, password string) (accessToken string, tokenType string, expiresAt int64, err error)
	Restricted(ctx context.Context) (claims *utils.JWTClaims, err error)
	HealthCheck(ctx context.Context) (status string, err error)
}

type basicAuthApiService struct{}

func (b *basicAuthApiService) Login(ctx context.Context, username string, password string) (accessToken string, tokenType string, expiresAt int64, err error) {
	userApi := utils.GetUserApi(ctx)
	reply, err := userApi.GetUserForAuth(ctx, &pb.GetUserForAuthRequest{Username: username})
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

// NewBasicAuthApiService returns a naive, stateless implementation of AuthApiService.
func NewBasicAuthApiService() AuthApiService {
	return &basicAuthApiService{}
}

// New returns a AuthApiService with all of the expected middleware wired in.
func New(middleware []Middleware) AuthApiService {
	var svc AuthApiService = NewBasicAuthApiService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicAuthApiService) Restricted(ctx context.Context) (claims *utils.JWTClaims, err error) {
	return utils.GetClaims(ctx), nil
}

func (b *basicAuthApiService) HealthCheck(ctx context.Context) (status string, err error) {
	return "OK", nil
}
