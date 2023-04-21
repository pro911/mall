package model

import (
	"time"
)

type User struct {
	UserID    int       `json:"user_id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex,comment=用户名称"`
	Password  string    `json:"password" gorm:"comment=密码"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
