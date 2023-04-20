package model

import "time"

type Cart struct {
	CartID    int       `json:"cart_id" gorm:"primaryKey"`
	UserID    int       `json:"user_id" gorm:"index"`
	GoodsID   int       `json:"goods_id" gorm:"index"`
	OrderID   int       `json:"order_id"`
	Price     int       `json:"price"`  //价格/分
	Status    int8      `json:"status"` //0购物车内 1已生成订单 -1移除购物车
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
