package routers

import (
	"github.com/gin-gonic/gin"
	"mall/internal/app/http/hendler"
	"mall/internal/pkg/response"
	"mall/internal/pkg/tools"
	"mall/middlewares"
)

func Init() *gin.Engine {
	r := gin.New()

	r.Use(middlewares.Cors(), middlewares.GinLogger(), middlewares.GinRecovery(true))

	r.GET("/online", func(ctx *gin.Context) {
		response.SuccessWithData(ctx, gin.H{
			"bool": tools.RandomBool(0.9),
		})
		return
	})

	apiV1 := r.Group("/api")
	{
		apiV1.GET("/swagger/:t/:f", hendler.Swagger) //t: json yaml

		//登录
		apiV1.POST("/login", hendler.Login)

		goods := apiV1.Group("/goods")
		{
			//商品列表
			goods.GET("/list", hendler.GoodsList)
			//商品详情
			goods.GET("/info", hendler.GoodsInfo)
			//商品评论
			goods.GET("/comment", hendler.GoodsComment)

			goods.Use(middlewares.JWT()).POST("/add", hendler.GoodsAdd)
		}

		apiV1.Use(middlewares.JWT())
		{
			//加入购物车
			apiV1.POST("/cart", hendler.AddCart)

			//获取优惠券
			apiV1.GET("/coupon/list", hendler.CouponList)

			//生成订单
			apiV1.POST("/order/create", hendler.OrderCreate)

			//支付
			apiV1.POST("/pay", hendler.Pay)
		}

	}
	return r
}
