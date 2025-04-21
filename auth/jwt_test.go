package auth

import (
	"testing"
)

func TestCreateAccess(t *testing.T) {
	secret := "testsecret"
	userID := "user-123"
	ip := "127.0.0.1"

	token, jti, err := CreateAccess(userID, ip, secret)
	if err != nil {
		t.Errorf("ошибка при создании access токена: %v", err)
	}
	if token == "" || jti == "" {
		t.Error("токен или jti пустые")
	}
}
