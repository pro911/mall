package model

import "time"

type Comment struct {
	CommentID int       `json:"comment_id" gorm:"primaryKey"`
	GoodsID   int       `json:"goods_id" gorm:"index"`
	UserID    int       `json:"user_id" gorm:"index"`
	Desc      string    `json:"Desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
