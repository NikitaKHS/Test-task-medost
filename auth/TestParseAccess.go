package auth

import "testing"

func TestParseAccess(t *testing.T) {
	secret := "testsecret"
	userID := "user-123"
	ip := "192.168.0.1"

	tokenStr, jti, err := CreateAccess(userID, ip, secret)
	if err != nil {
		t.Fatalf("ошибка при генерации токена: %v", err)
	}

	claims, err := ParseAccess(tokenStr, secret)
	if err != nil {
		t.Errorf("ошибка при парсинге токена: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("ожидался user_id %s, а получен %s", userID, claims.UserID)
	}
	if claims.IP != ip {
		t.Errorf("ожидался IP %s, а получен %s", ip, claims.IP)
	}
	if claims.ID != jti {
		t.Errorf("jti не совпадает")
	}
}
