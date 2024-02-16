package main

import (
	"github.com/XT4RM1NATOR/PostsProject/initializers"
	"github.com/XT4RM1NATOR/PostsProject/protos/user_service"
	"github.com/XT4RM1NATOR/PostsProject/user_service/repository"
	"github.com/XT4RM1NATOR/PostsProject/user_service/service"
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
	lis, err := net.Listen("tcp", ":"+os.Getenv("USER_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	repo := repository.NewUserRepository(DB)
	userService := service.NewUserService(repo)

	grpcServer := grpc.NewServer()

	s := Server{userService: userService}

	user_service.RegisterUserServiceServer(grpcServer, &s)
	log.Println("gRPC server started on port: ", os.Getenv("AUTH_SERVICE_PORT"))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌Failed to start the server❌: %v", err)
	}
}
