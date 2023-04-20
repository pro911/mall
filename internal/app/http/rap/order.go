package rap

type OrderCreateReq struct {
	CouponID int `json:"coupon_id" binding:"omitempty"`
}
