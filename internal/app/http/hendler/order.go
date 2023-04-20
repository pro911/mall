package hendler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mall/initializ"
	"mall/internal/app/http/rap"
	"mall/internal/app/logic/cart"
	"mall/internal/app/logic/coupon"
	"mall/internal/app/logic/goods"
	"mall/internal/app/model"
	"mall/internal/pkg/errno"
	"mall/internal/pkg/response"
	"strconv"
)

func OrderCreate(ctx *gin.Context) {
	p := new(rap.OrderCreateReq)
	if err := ctx.ShouldBindJSON(p); err != nil {
		// 非validator.ValidationErrors类型错误直接返回
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		//请求参数有误,直接返回响应
		zap.L().Error("OrderCreate with invalid param", zap.Error(err))
		return
	}
	zap.L().Info("OrderCreate", zap.Int("coupon_id", p.CouponID), zap.Int("user_id", ctx.GetInt("user_id")))
	//获取优惠券面值
	couponInfo, _ := coupon.CouponInfo(ctx, p.CouponID)
	//couponInfo, err := couponInfo.CouponInfo(ctx, p.CouponID)
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		response.ErrorWithMsg(ctx, errno.ErrRecordNotFound, "")
	//		return
	//	}
	//	response.ErrorWithMsg(ctx, errno.ErrServer, err.Error())
	//	return
	//}
	zap.L().Info("OrderCreate", zap.Any("couponInfo", couponInfo))
	//获取购物车内数据
	carts := cart.MyCartList(ctx, ctx.GetInt("user_id"))
	if carts == nil {
		response.ErrorWithMsg(ctx, errno.ErrCartEmpty, "")
		return
	}
	zap.L().Info("OrderCreate", zap.Any("couponInfo", couponInfo))

	//获取一个map用来存储goods_ids
	var goodsIDs []int
	var ClearCartIDs []int
	for k, v := range carts {
		goodsIDs = append(goodsIDs, v.GoodsID)
		zap.L().Info("OrderCreate", zap.Any(strconv.Itoa(k), v))
	}

	//获取购物车内的商品信息
	goodsList := goods.GoodsIn(ctx, goodsIDs)
	if goodsList == nil {
		for k, v := range carts {
			ClearCartIDs = append(ClearCartIDs, v.CartID)
			zap.L().Info("OrderCreate", zap.Any(strconv.Itoa(k), v))
		}
	}

	//处理成映射关系
	goodsKv := make(map[int]model.Goods)
	for _, m := range goodsList {
		goodsKv[m.GoodsID] = m
	}

	//计算购物车有效商品的标价总和
	var markPrice, discountPrice, price int
	var validCarts []model.Cart
	for _, v := range carts {
		c, ok := goodsKv[v.GoodsID]
		if !ok {
			ClearCartIDs = append(ClearCartIDs, v.CartID)
		} else {
			markPrice += c.Price
			v.Price = c.Price
			v.Status = 1
			validCarts = append(validCarts, v)
		}
	}

	//清理购物车无效数据
	if len(ClearCartIDs) > 0 {
		initializ.SQLiteDB().Model(model.Cart{}).Where("cart_id in ?", ClearCartIDs).Updates(model.Cart{Status: -1})
	}

	zap.L().Info("OrderCreate", zap.Int("markPrice", markPrice), zap.Int("couponInfo.Satisfy", couponInfo.Satisfy), zap.Int("couponInfo.Minus", couponInfo.Minus))

	//判断优惠券的使用
	if couponInfo.CouponID > 0 {
		if markPrice < couponInfo.Satisfy {
			response.ErrorWithMsg(ctx, errno.ErrCouponDoNotCondition, "")
			return
		}
		discountPrice = couponInfo.Minus
	}
	price = markPrice - discountPrice

	// 开始事务
	tx := initializ.SQLiteDB().Begin()

	//生成订单
	order := model.Order{
		UserID:        ctx.GetInt("user_id"),
		MarkPrice:     markPrice,
		Price:         price,
		DiscountPrice: discountPrice,
		CouponID:      couponInfo.CouponID,
	}
	result := initializ.SQLiteDB().Create(&order)

	if result.Error != nil {
		tx.Rollback()
		response.ErrorWithMsg(ctx, errno.ErrServer, result.Error.Error())
		return
	}

	for _, v := range validCarts {
		initializ.SQLiteDB().Model(model.Cart{}).Where("cart_id", v.CartID).Updates(model.Cart{
			OrderID: order.OrderID,
			Price:   v.Price,
			Status:  v.Status,
		})
	}
	tx.Commit()

	zap.L().Info("OrderCreate", zap.Int("price", price))
	response.SuccessWithData(ctx, order)
	return
}
