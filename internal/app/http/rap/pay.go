package rap

type PayReq struct {
	OrderID int    `json:"order_id" binding:"required"`
	PayType string `json:"pay_type" binding:"required,oneof=wechat alipay"`
}
