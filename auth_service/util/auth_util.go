package util

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var (
	accessTokenDuration  = time.Minute * 15
	refreshTokenDuration = time.Hour * 24 * 7
	jwtSecretKey         = os.Getenv("AUTH_SERVICE_JWT_KEY")
)

// Claims represents the claims of the JWT token.
type Claims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// GeneratePasswordHash generates a hash from a plaintext password.
func GeneratePasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswordAndHash compares a plaintext password with its hash.
func ComparePasswordAndHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// GenerateTokens generates access and refresh tokens for a given user ID and role.
func GenerateTokens(userID int64, role string) (string, string, error) {
	accessToken, err := generateToken(userID, role, accessTokenDuration)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateToken(userID, role, refreshTokenDuration)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// generateToken generates a JWT token with the provided claims and expiration duration.
func generateToken(userID int64, role string, duration time.Duration) (string, error) {
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

// ParseToken parses and validates a JWT token string.
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

// Middleware to extract JWT token from context
func ExtractTokenFromContext(ctx context.Context) (string, error) {
	token, ok := ctx.Value("token").(string)
	if !ok {
		return "", jwt.ErrECDSAVerification
	}
	return token, nil
}
