package pb

import (
	context "context"
	fmt "fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var svc UserApiClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial(":8087", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	svc = NewUserApiClient(conn)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		conn.Close()
	}()
	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	prefix := time.Now().Unix()
	reply, err := svc.CreateUser(ctx, &CreateUserRequest{
		Username: fmt.Sprintf("foobar1_%d", prefix),
		// Username: "foobar1",
		Email: fmt.Sprintf("foobar1_%d@gmail.com", prefix),
		// Email:    "foobar1@gmail.com",
		Password: "secret",
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
	md := metadata.Pairs("authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIwLjAuMC4wIiwiZXhwIjoxNTUyMDUxODIzLCJqdGkiOiIzZGM1ZDFhOC1kODkzLTQxY2EtYmFkMS03MjdmMWVhZDVhNWYiLCJpYXQiOjE1NTE5NjU0MjMsImlzcyI6IjAuMC4wLjAiLCJzdWIiOiJ1c2VyIn0.iHArP4tzk2FDBf1DG4JuODKcla-VTPPT-1gPQRzwVE1L8C6AdS5Z5FKSd6nVt3gsRX8wBJPXu4vZ_xHdoNV9v9b4h2S0bpm85qh_shl10al5CZJwS26SZ_e7GHK-sa_nXYXeZ9vTXFPPXUwMJedA1YIWWwT47_nSKtIOfP5Ma6kNufll0G9Pu0c_EuBhpV8KaybIey3X9Avf2gGY4I89KGiCMHdw2A2d1KI2ZKt7NvJrEfCVXK_Jqf8TmcBOSL9Yd4l7LderrA6rCES7HXQg8GraLvY3uKPCc7kBiA-zPl6U8jhxBzVuVWK495SwkHVRmbESLk3TQD10fGtIsQfn_A")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	reply, err := svc.GetByUsername(ctx, &GetByUsernameRequest{Username: "foobar1"})
	if err != nil {
		assert.Nil(err)
		return
	}
	assert.NotNil(reply)
}

func TestGetUserForAuth(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	reply, err := svc.GetUserForAuth(ctx, &GetUserForAuthRequest{Username: "foobar1"})
	if err != nil {
		assert.Nil(err)
		return
	}
	assert.NotNil(reply)
}
