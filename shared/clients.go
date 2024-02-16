package shared

import (
	authPb "github.com/XT4RM1NATOR/PostsProject/protos/auth_service"
	userPb "github.com/XT4RM1NATOR/PostsProject/protos/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func GetAuthServiceClient() (authPb.AuthServiceClient, error) {
	authConn, err := grpc.Dial("localhost:"+os.Getenv("AUTH_SERVICE_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer authConn.Close()

	return authPb.NewAuthServiceClient(authConn), nil
}

func GetUserServiceClient() (userPb.UserServiceClient, error) {
	userConn, err := grpc.Dial("localhost:"+os.Getenv("USER_SERVICE_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer userConn.Close()

	return userPb.NewUserServiceClient(userConn), nil
}

func GetPostServiceClient() {
	//userConn, err := grpc.Dial("localhost:"+os.Getenv("POST_SERVICE_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	return nil, err
	//}
	//defer userConn.Close()
	//
	//return userPb.NewUserServiceClient(userConn), nil
}
