package hendler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/unknwon/com"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mall/initializ"
	"mall/internal/app/http/rap"
	"mall/internal/app/logic/goods"
	"mall/internal/app/model"
	"mall/internal/pkg/errno"
	"mall/internal/pkg/response"
	"mall/internal/pkg/validators"
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

func GoodsAdd(ctx *gin.Context) {
	p := new(rap.GoodsAddReq)
	if err := ctx.ShouldBindJSON(p); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			response.ErrorWithMsg(ctx, errno.ErrParam, errs.Error())
			//请求参数有误,直接返回响应
			zap.L().Error("GoodsAdd with invalid param", zap.Error(err))
			return
		}

		//请求参数有误,直接返回响应
		jsonString, errJson := json.Marshal(validators.RemoveTopStruct(errs.Translate(validators.Trans)))
		if errJson != nil {
			response.ErrorWithMsg(ctx, errno.ErrParam, errJson.Error())
		}
		zap.L().Error("GoodsAdd with invalid param", zap.String("params", string(jsonString)))
		response.ErrorWithMsg(ctx, errno.ErrParam, string(jsonString))
		return
	}

	whereMap := make(map[string]interface{})
	whereMap["name"] = p.Name
	n, err := goods.GoodsWhereCount(ctx, whereMap)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrServer, err.Error())
		return
	}
	if n > 0 {
		response.ErrorWithMsg(ctx, errno.ErrGoodsExist, "")
		return
	}
	//插入商品
	goodsInsert := model.Goods{
		Name:    p.Name,
		Desc:    p.Desc,
		Price:   p.Price,
		Details: p.Details,
	}
	result := initializ.SQLiteDB().Create(&goodsInsert)

	if result.Error != nil {
		response.ErrorWithMsg(ctx, errno.ErrServer, result.Error.Error())
		return
	}

	if result.RowsAffected <= 0 {
		response.ErrorWithMsg(ctx, errno.ErrServer, "")
		return
	}

	response.SuccessWithData(ctx, gin.H{
		"goods_id": goodsInsert.GoodsID,
	})
	return
}
