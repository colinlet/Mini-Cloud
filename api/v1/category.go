package v1

import (
	"github.com/colinlet/mini_cloud/models"
	"github.com/colinlet/mini_cloud/pkg/e"
	"github.com/colinlet/mini_cloud/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Category = &category{}

type category struct{}

func (*category) GetList(c *gin.Context) {
	pid := c.Query("pid")
	if pid == "" {
		pid = "0"
	}

	data := make(map[string]interface{})
	list := models.Category.GetList(pid)

	sec, _ := setting.Cfg.GetSection("app")
	for key, value := range list {
		if value.Img != "" {
			list[key].Img = sec.Key("CDN").MustString("") + value.Img
		}
	}

	data["list"] = list
	data["total"] = models.Category.GetTotal(pid)

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
