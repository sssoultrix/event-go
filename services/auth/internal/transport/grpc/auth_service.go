package grpc

import (
	"context"

	authv1 "github.com/sssoultrix/event-go/services/auth/pkg/proto/auth/v1"
)

type AuthService struct {
	authv1.UnimplementedAuthServiceServer
}

func (a AuthService) Register(ctx context.Context, request *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthService) Login(ctx context.Context, request *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthService) Access(ctx context.Context, request *authv1.AccessRequest) (*authv1.AccessResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthService) Refresh(ctx context.Context, request *authv1.RefreshRequest) (*authv1.RefreshResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthService) Logout(ctx context.Context, request *authv1.LogoutRequest) (*authv1.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthService) mustEmbedUnimplementedAuthServiceServer() {
	//TODO implement me
	panic("implement me")
}
