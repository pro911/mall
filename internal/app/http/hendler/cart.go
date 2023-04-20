package hendler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mall/internal/app/http/rap"
	"mall/internal/app/logic/cart"
	"mall/internal/app/logic/goods"
	"mall/internal/app/model"
	"mall/internal/pkg/errno"
	"mall/internal/pkg/response"
	"mall/internal/pkg/validators"
)

func AddCart(ctx *gin.Context) {
	p := new(rap.AddCartReq)
	if err := ctx.ShouldBindJSON(p); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			response.ErrorWithMsg(ctx, errno.ErrParam, errs.Error())
			//请求参数有误,直接返回响应
			zap.L().Error("AddCart with invalid param", zap.Error(err))
			return
		}

		//请求参数有误,直接返回响应
		jsonString, errJson := json.Marshal(validators.RemoveTopStruct(errs.Translate(validators.Trans)))
		if errJson != nil {
			response.ErrorWithMsg(ctx, errno.ErrParam, errJson.Error())
		}
		zap.L().Error("AddCart with invalid param", zap.String("params", string(jsonString)))
		response.ErrorWithMsg(ctx, errno.ErrParam, string(jsonString))
		return
	}

	//查询商品信息
	goods, err := goods.GoodsInfo(ctx, p.GoodsID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.ErrorWithMsg(ctx, errno.ErrRecordNotFound, "")
			return
		}
		response.ErrorWithMsg(ctx, errno.ErrServer, err.Error())
		return
	}

	//将商品添加到购物车
	if err = cart.AddCart(ctx, model.Cart{
		UserID:  ctx.GetInt("user_id"),
		GoodsID: goods.GoodsID,
	}); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrServer, err.Error())
		return
	}

	response.Success(ctx)
	return
}
