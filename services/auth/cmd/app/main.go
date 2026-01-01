package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/sssoultrix/event-go/services/auth/internal/application/usecases/auth"
	"github.com/sssoultrix/event-go/services/auth/internal/config"
	"github.com/sssoultrix/event-go/services/auth/internal/infrastructure/cache/redis"
	userClient "github.com/sssoultrix/event-go/services/auth/internal/infrastructure/client"
	"github.com/sssoultrix/event-go/services/auth/internal/infrastructure/token"
	grpcTransport "github.com/sssoultrix/event-go/services/auth/internal/transport/grpc"
	authv1 "github.com/sssoultrix/event-go/services/auth/pkg/proto/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config/dev.yml", "path to config file")
	flag.Parse()
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	redisClient := redis.NewClient(cfg.Redis)
	if err := redisClient.Ping(context.Background(), 5*time.Second); err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	tokenStore := redis.NewTokenStore(redisClient.Client)

	tokenManager := token.NewJWTManager(cfg.JWT)

	usersServiceConn, err := grpc.Dial(
		cfg.Clients.UsersService.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to users service: %v", err)
	}
	defer usersServiceConn.Close()

	authUseCase := auth.NewAuthUseCase(
		userClient.NewUsersServiceClient(usersServiceConn),
		tokenManager,
		tokenStore,
	)

	grpcServer := grpc.NewServer()
	authService := grpcTransport.NewAuthService(authUseCase)
	authv1.RegisterAuthServiceServer(grpcServer, authService)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Starting auth service on port %s", cfg.GRPC.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
