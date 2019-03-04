package service

import (
	"context"
	"time"

	stdjwt "github.com/dgrijalva/jwt-go"
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
	userID := int64(1234567890)
	now := time.Now()
	expiresAt = now.Local().Add(time.Second * viper.GetDuration("JWT_EXPIRES_IN_SECONDS")).Unix()
	claims := &utils.JWTClaims{
		userID,
		stdjwt.StandardClaims{
			Audience:  "AUDIENCE_HERE",
			ExpiresAt: expiresAt,
			Id:        "ID_HERE",
			IssuedAt:  now.Unix(),
			Issuer:    "ISSUER_HERE",
			NotBefore: 0,
			Subject:   "SUBJECT_HERE",
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
