package hendler

import (
	"github.com/gin-gonic/gin"
	"mall/internal/app/logic/coupon"
	"mall/internal/pkg/response"
)

// CouponList 获取优惠券列表
func CouponList(ctx *gin.Context) {
	response.SuccessWithData(ctx, coupon.CouponList(ctx, ctx.GetInt("user_id")))
	return
}
