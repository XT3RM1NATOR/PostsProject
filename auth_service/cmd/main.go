package main

import (
	"github.com/XT4RM1NATOR/PostsProject/auth_service/repository"
	"github.com/XT4RM1NATOR/PostsProject/auth_service/service"
	"github.com/XT4RM1NATOR/PostsProject/initializers"
	"github.com/XT4RM1NATOR/PostsProject/protos/auth_service"
	userPb "github.com/XT4RM1NATOR/PostsProject/protos/user_service"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"
	"net"
	"os"
)

var (
	DB         *sqlx.DB
	UserClient userPb.UserServiceClient
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

	repo := repository.NewAuthRepository(DB)
	authService := service.NewAuthService(repo)

	grpcServer := grpc.NewServer()

	auth_service.RegisterAuthServiceServer(grpcServer, &server{authService: authService})
	log.Println("gRPC server started on port: ", os.Getenv("AUTH_SERVICE_PORT"))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌Failed to start the server❌: %v", err)
	}

	userConn, err := grpc.Dial("localhost:"+os.Getenv("USER_SERVICE_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC user_server: %v", err)
	}

	UserClient = userPb.NewUserServiceClient(userConn)

}
