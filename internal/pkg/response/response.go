package response

import (
	"github.com/gin-gonic/gin"
	"mall/internal/pkg/errno"
	"net/http"
)

type response struct {
	Code int         `json:"code"`
	Em   string      `json:"em,omitempty"`
	Et   string      `json:"et,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func display(c *gin.Context, code int, em string, et string, data interface{}) {
	respData := response{
		Code: code,
		Em:   em,
		Et:   et,
		Data: data,
	}
	c.JSON(http.StatusOK, respData)
}

// ErrorWithMsg 返回错误 附带更多信息
func ErrorWithMsg(c *gin.Context, code int, msg string) {
	if m, ok := errno.CodeMsgMap[code]; ok {
		msg = m + " " + msg
	}
	display(c, code, msg, errno.CodeAlertMap[code], struct{}{})
}

// SuccessWithData 返回成功并携带数据
func SuccessWithData(c *gin.Context, data interface{}) {
	display(c, errno.Success, errno.CodeMsgMap[errno.Success], errno.CodeAlertMap[errno.Success], data)
}

func Success(c *gin.Context) {
	display(c, errno.Success, errno.CodeMsgMap[errno.Success], errno.CodeAlertMap[errno.Success], struct{}{})
}
