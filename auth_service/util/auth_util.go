package util

import (
	"errors"
	userPb "github.com/XT4RM1NATOR/PostsProject/protos/user_service"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

var (
	AccessTokenDuration  = time.Hour * 24 * calculateTokenDuration("DAYS_ACCESS_TOKEN")
	RefreshTokenDuration = time.Hour * 24 * calculateTokenDuration("DAYS_REFRESH_TOKEN")
	jwtSecretKey         = os.Getenv("AUTH_SERVICE_JWT_KEY")
)

type Claims struct {
	UserID int
	Role   string
	jwt.StandardClaims
}

func GenerateTokens(userID int, role string) (string, string, error) {
	accessToken, err := GenerateToken(userID, role, AccessTokenDuration)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateToken(userID, role, RefreshTokenDuration)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func GenerateToken(userID int, role string, duration time.Duration) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func calculateTokenDuration(tokenType string) time.Duration {
	daysToken, err := strconv.Atoi(os.Getenv(tokenType))
	if err != nil {
		return time.Duration(1)
	}
	return time.Duration(daysToken)
}

func GetUserRole(role string) (userPb.Roles, error) {
	var roleEnum userPb.Roles
	switch role {
	case "user":
		roleEnum = userPb.Roles_user
	case "admin":
		roleEnum = userPb.Roles_admin
	default:
		return userPb.Roles_user, errors.New("invalid role")
	}
	return roleEnum, nil
}
