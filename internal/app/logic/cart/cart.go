package cart

import (
	"github.com/gin-gonic/gin"
	"mall/initializ"
	"mall/internal/app/model"
)

func AddCart(ctx *gin.Context, cart model.Cart) error {
	result := initializ.SQLiteDB().Create(&cart)
	return result.Error
}

func MyCartList(ctx *gin.Context, userID int) (carts []model.Cart) {
	initializ.SQLiteDB().Where("user_id", userID).Where("status", 0).Order("cart_id").Find(&carts)
	return
}
