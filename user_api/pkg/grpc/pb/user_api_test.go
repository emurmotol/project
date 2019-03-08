package pb

import (
	context "context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"syreclabs.com/go/faker"
)

var (
	testUserApiClient UserApiClient
	testAuthApiClient AuthApiClient
	authorization     string
	authUser          *User
	authUserPassword  = "secret"
)

func setTestAuthApiClient() {
	conn, err := grpc.Dial(":8085", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	testAuthApiClient = NewAuthApiClient(conn)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		conn.Close()
	}()
}

func setTestUserApiClient() {
	conn, err := grpc.Dial(":8087", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	testUserApiClient = NewUserApiClient(conn)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		conn.Close()
	}()
}

func setTestAuthorization() {
	ctx := context.Background()

	authUser = &User{
		Username: "foobar1",
	}
	reply, err := testAuthApiClient.Login(ctx, &LoginRequest{Username: authUser.Username, Password: authUserPassword})
	if err != nil {
		panic(err)
	}
	authorization = fmt.Sprintf("%s %s", reply.TokenType, reply.AccessToken)
}

func TestMain(m *testing.M) {
	setTestAuthApiClient()
	setTestUserApiClient()
	setTestAuthorization()
	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	username := faker.Internet().UserName()
	email := faker.Internet().SafeEmail()
	password := "secret"
	md := metadata.Pairs("authorization", authorization)
	ctx = metadata.NewOutgoingContext(ctx, md)
	reply, err := testUserApiClient.CreateUser(ctx, &CreateUserRequest{
		Username: username,
		Email:    email,
		Password: password,
		Role:     "user",
	})
	if err != nil {
		assert.Nil(err)
		return
	}
	assert.NotNil(reply)
}

func TestGetByUsername(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	md := metadata.Pairs("authorization", authorization)
	ctx = metadata.NewOutgoingContext(ctx, md)
	reply, err := testUserApiClient.GetByUsername(ctx, &GetByUsernameRequest{Username: authUser.Username})
	if err != nil {
		assert.Nil(err)
		return
	}
	assert.NotNil(reply)
}

func TestGetUserForAuth(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	reply, err := testUserApiClient.GetUserForAuth(ctx, &GetUserForAuthRequest{Username: authUser.Username})
	if err != nil {
		assert.Nil(err)
		return
	}
	assert.NotNil(reply)
}
