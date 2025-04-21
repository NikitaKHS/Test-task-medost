package auth

import (
	"testing"
)

func TestCreateAndCheckRefresh(t *testing.T) {
	raw, hash, err := CreateRefresh()
	if err != nil {
		t.Errorf("ошибка при создании refresh токена: %v", err)
	}

	// Проверка правильного токена
	if CheckRefresh(raw, hash) != nil {
		t.Error("ожидалось совпадение refresh токена, но оно не прошло")
	}

	// Проверка поддельного токена
	if CheckRefresh("wrong-token", hash) == nil {
		t.Error("ожидалась ошибка при проверке неверного токена")
	}
}
