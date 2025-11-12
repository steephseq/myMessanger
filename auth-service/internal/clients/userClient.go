package clients

import (
	"context"
	"errors"
	"pkg/proto/userpb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	client userpb.UserServiceClient
	conn   *grpc.ClientConn
}

func NewUserClient(connAddr string) (*UserClient, error) {
	conn, err := grpc.NewClient(
		connAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	client := userpb.NewUserServiceClient(conn)

	return &UserClient{
		client: client,
		conn:   conn,
	}, nil
}

func (u *UserClient) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	return u.client.CreateUser(ctx, req)
}

func (u *UserClient) Close() error {
	if u.conn != nil {
		return u.conn.Close()
	}
	return errors.New("user client conntection is nil")
}
