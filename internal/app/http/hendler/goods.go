package hendler

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"gorm.io/gorm"
	"mall/internal/app/logic/goods"
	"mall/internal/pkg/errno"
	"mall/internal/pkg/response"
)

func GoodsList(ctx *gin.Context) {
	k := ctx.DefaultQuery("keyword", "")
	response.SuccessWithData(ctx, goods.GoodsList(ctx, k))
	return
}

func GoodsInfo(ctx *gin.Context) {
	goodsID := com.StrTo(ctx.DefaultQuery("goods_id", "0")).MustInt()
	if goodsID == 0 {
		response.ErrorWithMsg(ctx, errno.ErrParam, "")
		return
	}

	//查询商品是否存在
	i, err := goods.GoodsInfo(ctx, goodsID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.ErrorWithMsg(ctx, errno.ErrRecordNotFound, "")
			return
		}
		response.ErrorWithMsg(ctx, errno.ErrServer, err.Error())
		return
	}
	response.SuccessWithData(ctx, i)
	return
}

func GoodsComment(ctx *gin.Context) {
	goodsID := com.StrTo(ctx.DefaultQuery("goods_id", "0")).MustInt()
	if goodsID == 0 {
		response.ErrorWithMsg(ctx, errno.ErrParam, "")
		return
	}

	//查询商品是否存在
	i, err := goods.GoodsComment(ctx, goodsID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.ErrorWithMsg(ctx, errno.ErrRecordNotFound, "")
			return
		}
		response.ErrorWithMsg(ctx, errno.ErrServer, err.Error())
		return
	}

	response.SuccessWithData(ctx, i)
	return
}
