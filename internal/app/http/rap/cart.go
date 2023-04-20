package rap

type AddCartReq struct {
	GoodsID int `json:"goods_id" binding:"required"`
}
