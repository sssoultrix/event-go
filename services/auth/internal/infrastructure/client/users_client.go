package client

import (
	"context"

	usersv1 "github.com/sssoultrix/event-go/contracts/users/pkg/proto/users/v1"
	"github.com/sssoultrix/event-go/services/auth/internal/domain"
	"google.golang.org/grpc"
)

type usersClient struct {
	client usersv1.UsersServiceClient
}

func NewUsersServiceClient(conn *grpc.ClientConn) domain.UsersService {
	return &usersClient{
		client: usersv1.NewUsersServiceClient(conn),
	}
}

func (c *usersClient) Register(ctx context.Context, params domain.CreateUserParams) (*domain.User, error) {
	resp, err := c.client.CreateUser(ctx, &usersv1.CreateUserRequest{
		Email:           params.Email,
		Password:        params.Password,
		PasswordConfirm: params.PasswordConfirm,
	})
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        resp.UserId,
		Email:     resp.Email,
		CreatedAt: resp.CreatedAt.AsTime(),
	}, nil
}

func (c *usersClient) Login(ctx context.Context, params domain.LoginParams) (*domain.User, error) {
	resp, err := c.client.Login(ctx, &usersv1.LoginRequest{
		Email:    params.Email,
		Password: params.Password,
	})
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:    resp.UserId,
		Email: resp.Email,
	}, nil
}
