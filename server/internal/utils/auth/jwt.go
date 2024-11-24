package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type JWTConfig struct {
	JwtAccessTokenSecret          []byte
	JwtRefreshTokenSecret         []byte
	JwtAccessTokenExpirationTime  time.Duration
	JwtRefreshTokenExpirationTime time.Duration
}

var JwtConf JWTConfig

func init() {
	err := LoadJWTConfigs()
	if err != nil {
		log.Fatalf("jwt error occured: %v", err)
	}
}

func LoadJWTConfigs() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	accessTokenSecret := os.Getenv("JWT_ACCESS_TOKEN_SECRET")
	if accessTokenSecret == "" {
		return fmt.Errorf("access token secret not found: %w", err)
	}

	refreshTokenSecret := os.Getenv("JWT_REFRESH_TOKEN_SECRET")
	if refreshTokenSecret == "" {
		return fmt.Errorf("refresh token secret not found: %w", err)
	}

	accessTokenExpiration, err := time.ParseDuration(os.Getenv("JWT_ACCESS_TOKEN_EXPIRATION_TIME"))
	if err != nil {
		return fmt.Errorf("duration must be in a <integer><time-unit> format: %w", err)
	}

	refreshTokenExpiration, err := time.ParseDuration(os.Getenv(("JWT_REFRESH_TOKEN_EXPIRATION_TIME")))
	if err != nil {
		return fmt.Errorf("duration must be in a <integer><time-unit> format: %w", err)
	}

	JwtConf = JWTConfig{
		JwtAccessTokenSecret:          []byte(accessTokenSecret),
		JwtRefreshTokenSecret:         []byte(refreshTokenSecret),
		JwtAccessTokenExpirationTime:  accessTokenExpiration,
		JwtRefreshTokenExpirationTime: refreshTokenExpiration,
	}

	return nil
}

func GenerateToken(id uint, role string) (string, string, error) {
	accessTokenClaims := jwt.MapClaims{
		"user_id": id,
		"role":    role,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(JwtConf.JwtAccessTokenExpirationTime).Unix(),
	}

	refreshTokenClaims := jwt.MapClaims{
		"user_id": id,
		"role":    role,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(JwtConf.JwtRefreshTokenExpirationTime).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	accessTokenString, err := accessToken.SignedString(JwtConf.JwtAccessTokenSecret)
	if err != nil {
		return "", "", fmt.Errorf("error generating access token: %w", err)
	}

	refreshTokenString, err := refreshToken.SignedString(JwtConf.JwtRefreshTokenSecret)
	if err != nil {
		return "", "", fmt.Errorf("error generating refresh token: %w", err)
	}

	return accessTokenString, refreshTokenString, nil
}

func ValidateToken(tokenStr string, secret []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
