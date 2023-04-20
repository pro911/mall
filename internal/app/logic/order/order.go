package order

import (
	"github.com/gin-gonic/gin"
	"mall/initializ"
	"mall/internal/app/model"
)

// OrderInfo 获取订单信息
func OrderInfo(ctx *gin.Context, orderID int) (order model.Order, err error) {
	err = initializ.SQLiteDB().Where("order_id", orderID).First(&order).Error
	return
}

func Pay(ctx *gin.Context) {

}
