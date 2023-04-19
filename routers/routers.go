package routers

import (
	"github.com/gin-gonic/gin"
	"mall/internal/app/http/hendler"
	"mall/internal/pkg/response"
	"mall/middlewares"
)

func Init() *gin.Engine {
	r := gin.New()

	r.Use(middlewares.Cors(), middlewares.GinLogger(), middlewares.GinRecovery(true))

	r.GET("/online", func(ctx *gin.Context) {
		response.Success(ctx)
		return
	})

	apiV1 := r.Group("/api")
	{
		//登录
		apiV1.POST("/login", hendler.Login)
		//商品列表
		apiV1.GET("/goods/list", func(ctx *gin.Context) {
			response.Success(ctx)
			return
		})

		apiV1.Use(middlewares.JWT())
		{
			//加入购物车
			apiV1.POST("/cart", func(ctx *gin.Context) {
				response.Success(ctx)
				return
			})
		}
	}
	return r
}
