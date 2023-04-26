package hendler

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	"mall/internal/pkg/errno"
	"mall/internal/pkg/response"
	"mall/settings"
	"net/http"
	"os"
)

func Swagger(c *gin.Context) {
	t := c.Param("t")
	f := c.Param("f")
	if t != "yaml" && t != "json" {
		response.ErrorWithMsg(c, errno.ErrUrlTypeFailNotSupport, "")
		return
	}
	//获取文件地址
	filePath := fmt.Sprintf("%vstatic/json/swagger%v.json", settings.Conf.AppConfig.Resources, f)
	b, err := os.ReadFile(filePath) //读取文件内容
	if err != nil {
		response.ErrorWithMsg(c, errno.ErrReadUrlFail, "")
		return
	}

	if t == "yaml" {
		//将json转换成yaml
		b, err = yaml.JSONToYAML(b)
		if err != nil {
			response.ErrorWithMsg(c, errno.ErrReadUrlFail, "")
			return
		}
	}

	c.String(http.StatusOK, string(b))
	return
}
