package main

import (
	"github.com/XT4RM1NATOR/PostsProject/grpc_gateway/internal/routes"
	"github.com/XT4RM1NATOR/PostsProject/initializers"
	"github.com/XT4RM1NATOR/PostsProject/shared"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {

	authClient, err := shared.GetAuthServiceClient()
	if err != nil {
		log.Fatal("Failed connecting to the service")
	}

	r := gin.Default()

	routes.AuthRoutes(r, authClient)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start Gin server: %v", err)
	}

}
