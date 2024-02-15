package main

import (
	"github.com/XT4RM1NATOR/PostsProject/grpc_gateway/internal/routes"
	"github.com/XT4RM1NATOR/PostsProject/initializers"
	pb "github.com/XT4RM1NATOR/PostsProject/protos/auth_service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	authConn, err := grpc.Dial("localhost:"+os.Getenv("AUTH_SERVICE_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC auth_server: %v", err)
	}
	defer authConn.Close()

	postConn, err := grpc.Dial("localhost:"+os.Getenv("POST_SERVICE_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC post_server: %v", err)
	}
	defer postConn.Close()

	userConn, err := grpc.Dial("localhost:"+os.Getenv("USER_SERVICE_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC user_server: %v", err)
	}
	defer userConn.Close()

	authClient := pb.NewAuthServiceClient(authConn)

	r := gin.Default()

	routes.AuthRoutes(r, authClient)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start Gin server: %v", err)
	}

}
