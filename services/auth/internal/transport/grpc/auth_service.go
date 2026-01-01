package grpc

import (
	"context"

	"github.com/sssoultrix/event-go/services/auth/internal/application/usecases/auth"
	"github.com/sssoultrix/event-go/services/auth/internal/domain"
	authv1 "github.com/sssoultrix/event-go/services/auth/pkg/proto/auth/v1"
)

type AuthService struct {
	authv1.UnimplementedAuthServiceServer
	useCase auth.AuthUseCase
}

func NewAuthService(uc auth.AuthUseCase) *AuthService {
	return &AuthService{useCase: uc}
}

func (a *AuthService) Register(ctx context.Context, request *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	panic("implement me")
}

func (a *AuthService) Login(ctx context.Context, request *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	params := domain.LoginParams{
		Email:    request.GetEmail(),
		Password: request.GetPassword(),
	}

	tokenPair, err := a.useCase.Login(ctx, params)
	if err != nil {
		return nil, err
	}

	return &authv1.LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

func (a *AuthService) Access(ctx context.Context, request *authv1.AccessRequest) (*authv1.AccessResponse, error) {
	panic("implement me")
}

func (a *AuthService) Refresh(ctx context.Context, request *authv1.RefreshRequest) (*authv1.RefreshResponse, error) {
	tokenPair, err := a.useCase.Refresh(ctx, request.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &authv1.RefreshResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

func (a *AuthService) Logout(ctx context.Context, request *authv1.LogoutRequest) (*authv1.LogoutResponse, error) {
	if err := a.useCase.Logout(ctx, request.GetRefreshToken()); err != nil {
		return nil, err
	}

	return &authv1.LogoutResponse{Success: true}, nil
}

func (a AuthService) mustEmbedUnimplementedAuthServiceServer() {
	panic("implement me")
}
