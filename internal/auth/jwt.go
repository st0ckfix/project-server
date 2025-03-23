package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

// Token status constants
const (
	TOKEN_VALID          = 1
	TOKEN_INVALID_FORMAT = 2
	TOKEN_EXPIRED        = 3
	TOKEN_EMPTY          = 4
	TOKEN_BLACKLISTED    = 5
	TOKEN_NOT_FOUND      = 6
	TOKEN_UNKNOWN_ERROR  = 7
)

var jwtSecret []byte
var redisClient *redis.Client
var ctx = context.Background()

func init() {
	// Use environment variable or fallback to default
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "v4TLZ6MabrGhUvFbc8V7zBoAR9VTuqBARFnpM2vdD64" // Fallback secret
	}
	jwtSecret = []byte(secretKey)

	// Setup Redis for token blacklisting
	redisClient = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379", // Replace with your Redis address
	})
}

// GenerateAccessToken Tạo access token
func GenerateAccessToken(username string) (string, error) {

	claims := jwt.MapClaims{
		"username": fmt.Sprintf("%s", username),
		"exp":      time.Now().Add(time.Minute * 15).Unix(), // Hết hạn sau 15 phút
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// GenerateRefreshToken Tạo refresh token (hết hạn lâu hơn)
func GenerateRefreshToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": fmt.Sprintf("%s", username),
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // Hết hạn sau 7 ngày
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken decodes the token (used in middleware)
func ParseToken(tokenStr string) (jwt.MapClaims, int) {
	if tokenStr == "" {
		return nil, TOKEN_EMPTY
	}

	// Check blacklist first
	if IsTokenBlacklisted(tokenStr) {
		return nil, TOKEN_BLACKLISTED
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, TOKEN_INVALID_FORMAT
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, TOKEN_EXPIRED
		} else {
			return nil, TOKEN_UNKNOWN_ERROR
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, TOKEN_NOT_FOUND
	}

	return claims, TOKEN_VALID
}

// BlacklistToken adds a token to the blacklist
func BlacklistToken(tokenStr string) error {
	// Parse the token to get expiration time
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		// Even if we can't parse it, blacklist for a default period
		return redisClient.Set(ctx, "blacklist:"+tokenStr, "1", time.Hour*24).Err()
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Get expiration time from token
		if exp, ok := claims["exp"].(float64); ok {
			expTime := time.Unix(int64(exp), 0)
			ttl := expTime.Sub(time.Now())
			if ttl > 0 {
				return redisClient.Set(ctx, "blacklist:"+tokenStr, "1", ttl).Err()
			}
		}
	}

	// If we couldn't extract expiration time, blacklist for a day
	return redisClient.Set(ctx, "blacklist:"+tokenStr, "1", time.Hour*24).Err()
}

// IsTokenBlacklisted checks if a token is in the blacklist
func IsTokenBlacklisted(tokenStr string) bool {
	result, err := redisClient.Exists(ctx, "blacklist:"+tokenStr).Result()
	if err != nil {
		return false
	}
	return result > 0
}
