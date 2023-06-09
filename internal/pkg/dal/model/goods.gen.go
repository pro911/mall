// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGood = "goods"

// Good mapped from table <goods>
type Good struct {
	GoodsID   int32     `gorm:"column:goods_id;primaryKey" json:"goods_id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Desc      string    `gorm:"column:desc;not null" json:"desc"`
	Price     int32     `gorm:"column:price;not null" json:"price"`
	Details   string    `gorm:"column:details;not null" json:"details"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName Good's table name
func (*Good) TableName() string {
	return TableNameGood
}
