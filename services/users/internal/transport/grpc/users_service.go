package grpc

import (
	"context"

	usersv1 "github.com/sssoultrix/event-go/services/auth/pkg/proto/users/v1"
)

type UsersService struct {
	usersv1.UnimplementedUsersServiceServer
}

func (u UsersService) CreateUser(ctx context.Context, request *usersv1.CreateUserRequest) (*usersv1.CreateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UsersService) Login(ctx context.Context, request *usersv1.LoginRequest) (*usersv1.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UsersService) DeleteUser(ctx context.Context, request *usersv1.DeleteUserRequest) (*usersv1.DeleteUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UsersService) mustEmbedUnimplementedUsersServiceServer() {
	//TODO implement me
	panic("implement me")
}
