package service

import (
	"context"
	pb "github.com/XT4RM1NATOR/PostsProject/protos/auth_service" // Import generated gRPC client code
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthenticateHandler(client pb.AuthServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req pb.AuthRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := client.Authenticate(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}

func RegisterHandler(client pb.AuthServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req pb.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := client.Register(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}
