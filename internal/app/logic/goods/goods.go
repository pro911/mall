package goods

import (
	"github.com/gin-gonic/gin"
	"mall/initializ"
	"mall/internal/app/model"
)

// GoodsList 商品列表
func GoodsList(ctx *gin.Context, k string) (goodsList []model.Goods) {
	queryStr := initializ.SQLiteDB()
	if k != "" {
		queryStr = queryStr.Where("name LIKE ?", "%"+k+"%")
	}
	queryStr.Order("goods_id desc").Find(&goodsList)
	return
}

// GoodsInfo 商品详情
func GoodsInfo(ctx *gin.Context, goodsID int) (goods model.Goods, err error) {
	err = initializ.SQLiteDB().Where("goods_id", goodsID).First(&goods).Error
	return
}

func GoodsIn(ctx *gin.Context, goodsIDs []int) (goodsList []model.Goods) {
	initializ.SQLiteDB().Where("goods_id in ?", goodsIDs).Find(&goodsList)
	return
}
