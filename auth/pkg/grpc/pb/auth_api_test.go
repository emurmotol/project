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

func TestLogin(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	reply, err := testAuthClient.Login(ctx, &LoginRequest{Username: authUser.Username, Password: authUser.Password})
	if err != nil {
		assert.Nil(err)
		return
	}
	assert.NotNil(reply)
}

func TestRestricted(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	md := metadata.Pairs("authorization", authorization)
	ctx = metadata.NewOutgoingContext(ctx, md)
	reply, err := testAuthClient.Restricted(ctx, &RestrictedRequest{})
	if err != nil {
		assert.Nil(err)
		return
	}
	assert.NotNil(reply)
}

func TestHealthCheck(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	reply, err := testAuthClient.HealthCheck(ctx, &HealthCheckRequest{})
	if err != nil {
		assert.Nil(err)
		return
	}
	assert.NotNil(reply)
}
