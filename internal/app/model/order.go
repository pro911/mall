package model

import "time"

type Order struct {
	OrderID       int       `json:"order_id" gorm:"primaryKey"`
	UserID        int       `json:"user_id" gorm:"index"`
	MarkPrice     int       `json:"mark_price"`     //标价
	Price         int       `json:"price"`          //价格/分 真实价格
	DiscountPrice int       `json:"discount_price"` //优惠减免的钱 /分
	CouponID      int       `json:"coupon_id"`      //优惠券id
	Status        int8      `json:"status"`         //0待支付 1已支付 -1删除
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
