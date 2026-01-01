package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sssoultrix/event-go/services/users/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load configuration
	cfg, err := config.Load("configs/prod.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to PostgreSQL
	// db, err := postgres.NewConnection(ctx, cfg.Postgres.DSN())
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }
	// defer db.Close()

	// Initialize repositories
	// userRepo := postgres.NewUserRepository(db)

	// Initialize use cases
	// usersUseCase := users.NewUsersUseCase(userRepo)

	// Create GRPC server
	grpcServer := grpc.NewServer()

	// Register services (you'll need to implement the GRPC service)
	// pb.RegisterUsersServiceServer(grpcServer, &grpcHandler{usersUseCase: usersUseCase})

	// Enable reflection for development
	reflection.Register(grpcServer)

	// Start GRPC server
	lis, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		log.Printf("Users service starting on port %s", cfg.GRPC.Port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down users service...")

	// Graceful shutdown
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		grpcServer.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		log.Println("Users service stopped gracefully")
	case <-ctx.Done():
		log.Println("Shutdown timeout, forcing exit")
	}
}
