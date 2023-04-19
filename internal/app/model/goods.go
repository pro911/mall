package model

import "time"

type Goods struct {
	GoodsID   int       `json:"goods_id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"index"`
	Desc      string    `json:"Desc"`
	Price     int       `json:"price"` //价格/分
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
