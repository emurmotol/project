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
	testUserClient UserClient
	testAuthClient AuthClient
	authorization  string
	authUser       *User
)

func setTestAuthClient() {
	conn, err := grpc.Dial(":8085", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	testAuthClient = NewAuthClient(conn)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		conn.Close()
	}()
}

func setTestUserClient() {
	conn, err := grpc.Dial(":8087", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	testUserClient = NewUserClient(conn)

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
		Username: "testuser",
		Password: "secret",
	}
	reply, err := testAuthClient.Login(ctx, &LoginRequest{Username: authUser.Username, Password: authUser.Password})
	if err != nil {
		panic(err)
	}
	authorization = fmt.Sprintf("%s %s", reply.TokenType, reply.AccessToken)
}

func TestMain(m *testing.M) {
	setTestAuthClient()
	setTestUserClient()
	setTestAuthorization()
	code := m.Run()
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	username := faker.Internet().UserName()
	email := faker.Internet().SafeEmail()
	password := "secret"
	md := metadata.Pairs("authorization", authorization)
	ctx = metadata.NewOutgoingContext(ctx, md)
	reply, err := testUserClient.CreateUser(ctx, &CreateUserRequest{
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
	reply, err := testUserClient.GetByUsername(ctx, &GetByUsernameRequest{Username: authUser.Username})
	if err != nil {
		assert.Nil(err)
		return
	}
	assert.NotNil(reply)
}

func TestGetUserForAuth(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	reply, err := testUserClient.GetUserForAuth(ctx, &GetUserForAuthRequest{Username: authUser.Username})
	if err != nil {
		assert.Nil(err)
		return
	}
	assert.NotNil(reply)
}
