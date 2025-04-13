package authentication

import (
	"net/http"
	"template_backend/core/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

type JWTContent struct {
	ID string `json:"_id"`
	jwt.RegisteredClaims
}

// CreateJWT creates a JWT for a user
func CreateJWT(userId *string) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(config.GlobalConfig.Auth.TokenExpirationMinutes))
	content := &JWTContent{
		ID: *userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, content)
	log.Info().Msgf("Signing JWT with secret: %s", config.GlobalConfig.Auth.JWTSecret)
	signedJWT, err := token.SignedString([]byte(config.GlobalConfig.Auth.JWTSecret))

	return signedJWT, expirationTime, err
}

// VerifyJWT verifies a JWT
func VerifyJWT(jwtString *string) (*jwt.Token, *JWTContent, error) {
	content := &JWTContent{}

	token, err := jwt.ParseWithClaims(*jwtString, content, func(token *jwt.Token) (any, error) {
		return []byte(config.GlobalConfig.Auth.JWTSecret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Error().Int("status", http.StatusUnauthorized).Msg("verify JWT")
		}
		log.Error().Int("status", http.StatusBadRequest).Msg("verify JWT")
	}
	if !token.Valid {
		log.Error().Int("status", http.StatusUnauthorized).Msg("tampered JWT")
	}

	return token, content, err
}
