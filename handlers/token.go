package handlers

import (
	"net/http"
	"time"

	"authsvc/auth"
	"authsvc/db/models"
	"database/sql"
	"github.com/gin-gonic/gin"
)

var DB *sql.DB
var JWTSecret string

func Issue(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка: user_id не задан"})
		return
	}
	ip := c.ClientIP()
	access, jti, err := auth.CreateAccess(userID, ip, JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании access token"})
		return
	}
	raw, hash, err := auth.CreateRefresh()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании refresh token"})
		return
	}
	expires := time.Now().Add(auth.RefreshTTL)
	if _, err := DB.Exec(`INSERT INTO refresh_tokens(user_id, token_hash, access_jti, ip, created_at, expires_at) VALUES($1,$2,$3,$4,NOW(),$5)`, userID, hash, jti, ip, expires); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении refresh token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": access, "refresh_token": raw})
}

func Refresh(c *gin.Context) {
	var in struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	if c.BindJSON(&in) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка: неверный формат JSON"})
		return
	}
	claims, err := auth.ParseAccess(in.AccessToken, JWTSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка: недействительный access token"})
		return
	}
	var rt models.RefreshToken
	if err := DB.QueryRow(`SELECT id,user_id,token_hash,access_jti,ip,created_at,expires_at FROM refresh_tokens WHERE access_jti=$1`, claims.ID).Scan(&rt.ID, &rt.UserID, &rt.TokenHash, &rt.AccessJTI, &rt.IP, &rt.CreatedAt, &rt.ExpiresAt); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка: refresh token не найден"})
		return
	}
	if time.Now().After(rt.ExpiresAt) || auth.CheckRefresh(in.RefreshToken, rt.TokenHash) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка: недействующий refresh token"})
		return
	}
	newIP := c.ClientIP()
	if newIP != rt.IP {
		// лог (так проще)
	}
	DB.Exec(`DELETE FROM refresh_tokens WHERE id=$1`, rt.ID)
	access2, jti2, err := auth.CreateAccess(rt.UserID, newIP, JWTSecret)
	raw2, hash2, err := auth.CreateRefresh()
	expires2 := time.Now().Add(auth.RefreshTTL)
	DB.Exec(`INSERT INTO refresh_tokens(user_id, token_hash, access_jti, ip, created_at, expires_at) VALUES($1,$2,$3,$4,NOW(),$5)`, rt.UserID, hash2, jti2, newIP, expires2)
	c.JSON(http.StatusOK, gin.H{"access_token": access2, "refresh_token": raw2})
}
