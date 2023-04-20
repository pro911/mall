package hendler

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mall/initializ"
	"mall/internal/app/http/rap"
	"mall/internal/app/logic/order"
	"mall/internal/app/model"
	"mall/internal/pkg/errno"
	"mall/internal/pkg/response"
	"mall/internal/pkg/tools"
)

func Pay(ctx *gin.Context) {
	p := new(rap.PayReq)
	if err := ctx.ShouldBindJSON(p); err != nil {
		// 非validator.ValidationErrors类型错误直接返回
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		//请求参数有误,直接返回响应
		zap.L().Error("Pay with invalid param", zap.Error(err))
		return
	}
	zap.L().Info("Pay", zap.Int("order_id", p.OrderID), zap.String("pay_type", p.PayType), zap.Int("user_id", ctx.GetInt("user_id")))

	order, err := order.OrderInfo(ctx, p.OrderID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.ErrorWithMsg(ctx, errno.ErrReportNotFound, err.Error())
		return
	}

	//查看订单状体
	if order.Status == 1 {
		response.Success(ctx)
		return
	} else if order.Status == -1 {
		response.ErrorWithMsg(ctx, errno.ErrParam, "订单不存在或已被删除")
	}

	//执行支付
	if !tools.RandomBool(0.3) {
		response.ErrorWithMsg(ctx, errno.ErrPayFail, "")
	}

	//支付成功修改订单数据
	initializ.SQLiteDB().Model(model.Order{}).Where("order_id", order.OrderID).Updates(model.Order{
		PayType: p.PayType,
		Status:  1,
	})

	response.Success(ctx)
	return
}
