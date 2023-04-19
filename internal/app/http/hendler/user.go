package hendler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"mall/internal/app/http/rap"
	"mall/internal/app/logic/user"
	"mall/internal/pkg/errno"
	"mall/internal/pkg/response"
	"mall/internal/pkg/validators"
	"mall/middlewares"
)

func Login(ctx *gin.Context) {
	p := new(rap.LoginReq)
	if err := ctx.ShouldBindJSON(p); err != nil {

		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			response.ErrorWithMsg(ctx, errno.ErrParam, errs.Error())
			//请求参数有误,直接返回响应
			zap.L().Error("SignUp with invalid param", zap.Error(err))
			return
		}

		//请求参数有误,直接返回响应
		jsonString, errJson := json.Marshal(validators.RemoveTopStruct(errs.Translate(validators.Trans)))
		if errJson != nil {
			response.ErrorWithMsg(ctx, errno.ErrParam, errJson.Error())
		}
		zap.L().Error("SignUp with invalid param", zap.String("params", string(jsonString)))
		response.ErrorWithMsg(ctx, errno.ErrParam, string(jsonString))
		return
	}

	//查询用户是否存在
	b, err := user.Login(ctx, p)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrServer, err.Error())
		return
	}

	if b <= 0 {
		response.ErrorWithMsg(ctx, errno.ErrAuthFailed, "")
		return
	}

	//获取token
	token, err := middlewares.GenerateToken(p.Username, p.Password)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrServer, err.Error())
		return
	}

	response.SuccessWithData(ctx, gin.H{"token": token})
	return
}
