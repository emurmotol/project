package service

import (
	"context"
	"errors"
	"time"

	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/emurmotol/project/auth_api/pkg/utils"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/spf13/viper"
)

// AuthApiService describes the service.
type AuthApiService interface {
	Login(ctx context.Context, payload *LoginInput) (data *LoginOutput, err error)
	Restricted(ctx context.Context) (data *RestrictedOutput, err error)
	HealthCheck(ctx context.Context) (status string, err error)
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginOutput struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresAt   int64  `json:"expires_at"`
}

type RestrictedOutput struct {
	Claims *stdjwt.StandardClaims `json:"claims"`
}

type basicAuthApiService struct{}

func (b *basicAuthApiService) Login(ctx context.Context, payload *LoginInput) (data *LoginOutput, err error) {
	now := time.Now()
	expiresAt := now.Local().Add(time.Second * viper.GetDuration("JWT_EXPIRES_IN_SECONDS")).Unix()
	claims := stdjwt.StandardClaims{
		Audience:  "AUDIENCE_HERE",
		ExpiresAt: expiresAt,
		Id:        "ID_HERE",
		IssuedAt:  now.Unix(),
		Issuer:    "ISSUER_HERE",
		NotBefore: 0,
		Subject:   "SUBJECT_HERE",
	}
	token, err := utils.GenerateJwtToken(claims, viper.GetString("JWT_PRIVATE_KEY"))
	if err != nil {
		return nil, err
	}

	data = &LoginOutput{
		AccessToken: token,
		TokenType:   viper.GetString("JWT_TOKEN_TYPE"),
		ExpiresAt:   time.Now().Unix(),
	}
	return data, nil
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

func (b *basicAuthApiService) Restricted(ctx context.Context) (data *RestrictedOutput, err error) {
	claims, ok := ctx.Value(jwt.JWTClaimsContextKey).(*stdjwt.StandardClaims)
	if !ok {
		return nil, errors.New("claims not ok")
	}
	return &RestrictedOutput{Claims: claims}, nil
}

func (b *basicAuthApiService) HealthCheck(ctx context.Context) (status string, err error) {
	return "OK", nil
}
