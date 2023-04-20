package coupon

import (
	"github.com/gin-gonic/gin"
	"mall/initializ"
	"mall/internal/app/model"
)

// CouponList 获取优惠券列表
func CouponList(ctx *gin.Context, userID int) (coupons []model.Coupon) {
	initializ.SQLiteDB().Order("coupon_id desc").Find(&coupons)
	return
}

// CouponInfo 优惠券详情
func CouponInfo(ctx *gin.Context, couponID int) (coupon model.Coupon, err error) {
	err = initializ.SQLiteDB().Where("coupon_id", couponID).First(&coupon).Error
	return
}
