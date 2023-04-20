package model

import "time"

type Coupon struct {
	CouponID  int       `json:"coupon_id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Satisfy   int       `json:"satisfy"` //满 /分
	Minus     int       `json:"minus"`   //减 /分
	Desc      string    `json:"Desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
