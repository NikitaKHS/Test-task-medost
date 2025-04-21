package auth

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const AccessTTL = 15 * time.Minute

type Claims struct {
	UserID string `json:"user_id"`
	IP     string `json:"ip"`
	jwt.RegisteredClaims
}

func CreateAccess(userID, ip, secret string) (string, string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", "", err
	}
	jti := base64.RawURLEncoding.EncodeToString(b)
	claims := Claims{UserID: userID, IP: ip, RegisteredClaims: jwt.RegisteredClaims{ID: jti, IssuedAt: jwt.NewNumericDate(time.Now()), ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTTL))}}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := tok.SignedString([]byte(secret))
	return tokenStr, jti, err
}

func ParseAccess(tokenStr, secret string) (*Claims, error) {
	tok, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	if err != nil || !tok.Valid {
		return nil, err
	}
	return tok.Claims.(*Claims), nil
}
