package auth

import (
	"time"

	"github.com/devproje/project-mirror/src/config"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID string
	jwt.StandardClaims
}

var (
	expTime = 5 * time.Minute
	JwtKey  = []byte(config.Get().SecretKey)
)

func (acc *Account) GetJwtToken() (string, error) {
	expTime := time.Now().Add(expTime)
	claims := &Claims{
		UserID: acc.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	JWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return JWT.SignedString(JwtKey)
}
