package user

import (
	"context"
	"mall/initializ"
	"mall/internal/app/http/rap"
	"mall/internal/app/model"
)

func Login(ctx context.Context, req *rap.LoginReq) (user model.User, err error) {
	//查询用户是否存在
	err = initializ.SQLiteDB().
		Where("username", req.Username).
		Where("password", req.Password).
		First(&user).Error
	return
}

func AddUser(ctx context.Context, req *rap.LoginReq) error {
	result := initializ.SQLiteDB().Create(&model.User{
		Username: req.Username,
		Password: req.Password,
	})
	return result.Error
}
