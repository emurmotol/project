package service

import (
	"context"
	"errors"
	"time"

	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/emurmotol/project/auth_api/pkg/utils"
	"github.com/go-kit/kit/auth/jwt"
)

// AuthApiService describes the service.
type AuthApiService interface {
	Login(ctx context.Context, payload *LoginInput) (data *LoginOutput, err error)
	Restricted(ctx context.Context) (claims *stdjwt.StandardClaims, err error)
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

type basicAuthApiService struct{}

func (b *basicAuthApiService) Login(ctx context.Context, payload *LoginInput) (data *LoginOutput, err error) {
	// TODO implement the business logic of Login

	now := time.Now()
	expiresAt := now.Local().Add(time.Second * time.Duration(86400)).Unix()
	claims := stdjwt.StandardClaims{
		Audience:  "AUDIENCE_HERE",
		ExpiresAt: expiresAt,
		Id:        "USER_ID_HERE",
		IssuedAt:  now.Unix(),
		Issuer:    "ISSUER_HERE",
		NotBefore: 0,
		Subject:   "SUBJECT_HERE",
	}
	token, err := utils.GenerateJwtToken(claims, "/go/src/github.com/emurmotol/project/auth_api/certs/jwt.key")
	if err != nil {
		return nil, err
	}

	data = &LoginOutput{
		AccessToken: token,
		TokenType:   "Bearer",
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

func (b *basicAuthApiService) Restricted(ctx context.Context) (claims *stdjwt.StandardClaims, err error) {
	// TODO implement the business logic of Restricted
	claims, ok := ctx.Value(jwt.JWTClaimsContextKey).(*stdjwt.StandardClaims)
	if !ok {
		return nil, errors.New("claims not ok")
	}
	return claims, nil
}

func (b *basicAuthApiService) HealthCheck(ctx context.Context) (status string, err error) {
	// TODO implement the business logic of HealthCheck
	return "OK", nil
}
