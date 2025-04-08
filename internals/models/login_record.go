package models

import (
	"time"
)

type LoginRecord struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	IPAddress string    `json:"ip_address"`
	LoginTime time.Time `json:"login_time"`
	CreatedAt time.Time `json:"created_at"`
}
