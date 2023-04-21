package rap

type GoodsAddReq struct {
	Name    string `json:"name" binding:"required,min=1,max=120"`
	Desc    string `json:"desc" binding:"required,min=1,max=120"`
	Price   int    `json:"price" binding:"required,min=1"`
	Details string `json:"details" binding:"required"`
}
