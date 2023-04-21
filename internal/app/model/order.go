package model

import "time"

type Order struct {
	OrderID       int       `json:"order_id" gorm:"primaryKey"`
	UserID        int       `json:"user_id" gorm:"index"`
	MarkPrice     int       `json:"mark_price" gorm:"comment=标价"`               //标价
	Price         int       `json:"price" gorm:"comment=真实价格,折扣后的价格支付价格"`       //价格/分 真实价格
	DiscountPrice int       `json:"discount_price" gorm:"comment=优惠的价格"`        //优惠减免的钱 /分
	CouponID      int       `json:"coupon_id" gorm:"comment=优惠券ID"`             //优惠券id
	Status        int8      `json:"status" gorm:"comment=状态: 0待支付 1已支付 -1删除失效"` //0待支付 1已支付 -1删除
	PayType       string    `json:"pay_type" gorm:"comment=支付类型:wechat,alipay"` //支付类型
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
