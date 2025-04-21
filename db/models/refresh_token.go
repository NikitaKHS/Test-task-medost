package models

import "time"

type RefreshToken struct {
	ID        int64     `db:"id"`
	UserID    string    `db:"user_id"`
	TokenHash string    `db:"token_hash"`
	AccessJTI string    `db:"access_jti"`
	IP        string    `db:"ip"`
	CreatedAt time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expires_at"`
}
