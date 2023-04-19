package user

import (
	"context"
	"mall/initializ"
	"mall/internal/app/http/rap"
	"mall/internal/app/model"
)

func Login(ctx context.Context, req *rap.LoginReq) (int64, error) {
	//查询用户是否存在
	var count int64
	result := initializ.SQLiteDB().
		Model(&model.User{}).
		Where("username", req.Username).
		Where("password", req.Password).
		Limit(1).
		Count(&count)
	return count, result.Error
}

func AddUser(ctx context.Context, req *rap.LoginReq) error {
	result := initializ.SQLiteDB().Create(&model.User{
		Username: req.Username,
		Password: req.Password,
	})
	return result.Error
}
