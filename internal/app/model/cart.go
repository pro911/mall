package model

import "time"

type Cart struct {
	CartID    int       `json:"cart_id" gorm:"primaryKey"`
	UserID    int       `json:"user_id" gorm:"index"`
	GoodsID   int       `json:"goods_id" gorm:"index"`
	Num       int       `json:"num"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
