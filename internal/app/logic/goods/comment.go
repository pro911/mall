package goods

import (
	"github.com/gin-gonic/gin"
	"mall/initializ"
	"mall/internal/app/model"
)

func GoodsComment(ctx *gin.Context, goodsID int) (comments []model.Comment, err error) {
	r := initializ.SQLiteDB().Where("goods_id", goodsID).Find(&comments)
	return comments, r.Error
}
