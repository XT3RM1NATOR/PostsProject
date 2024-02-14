package main

import (
	"github.com/XT4RM1NATOR/PostsProject/initializers"
	"github.com/XT4RM1NATOR/PostsProject/protos/auth_service" // Import the package where AuthService is defined
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var (
	DB *sqlx.DB
)

func init() {
	initializers.LoadEnvVariables()
	DB = initializers.ConnectToDb()
}

func main() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("AUTH_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	authService := &auth_service.AuthService{}

	auth_service.RegisterAuthServiceServer(grpcServer, authService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌Failed to start the server❌: %v", err)
	}
}
