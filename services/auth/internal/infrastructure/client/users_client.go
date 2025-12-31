package client

import usersv1 "github.com/sssoultrix/event-go/services/users/pkg/proto/users/v1"

type usersClient struct {
}

func (u usersClient) CreateUser(ctx context.Context, in *usersv1.CreateUserRequest, opts ...grpc.CallOption) (*usersv1.CreateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usersClient) Login(ctx context.Context, in *usersv1.LoginRequest, opts ...grpc.CallOption) (*usersv1.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usersClient) DeleteUser(ctx context.Context, in *usersv1.DeleteUserRequest, opts ...grpc.CallOption) (*usersv1.DeleteUserResponse, error) {
	//TODO implement me
	panic("implement me")
}
