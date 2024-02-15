package routes

import (
	"github.com/XT4RM1NATOR/PostsProject/grpc_gateway/internal/service"
	pb "github.com/XT4RM1NATOR/PostsProject/protos/auth_service"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, grpcClient pb.AuthServiceClient) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/authenticate", service.AuthenticateHandler(grpcClient))
		authGroup.POST("/register", service.RegisterHandler(grpcClient))
	}
}
