package model

import "time"

type Goods struct {
	GoodsID   int       `json:"goods_id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"uniqueIndex,comment=商品名称"`
	Desc      string    `json:"Desc"  gorm:"comment=描述"`
	Price     int       `json:"price" gorm:"comment=价格/分"` //价格/分
	Details   string    `json:"details" gorm:"comment=详情内容body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
