package v1

import (
	"github.com/colinlet/mini_cloud/models"
	"github.com/colinlet/mini_cloud/pkg/e"
	"github.com/colinlet/mini_cloud/pkg/setting"
	"github.com/colinlet/mini_cloud/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Goods = &goods{}

type goods struct{}

func (*goods) GetList(c *gin.Context) {
	categoryId := c.Query("id")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	maps["category_id"] = categoryId
	maps["status"] = 1
	list := models.Goods.GetList(util.GetPage(c), setting.PageSize, maps)

	sec, _ := setting.Cfg.GetSection("app")
	for key, value := range list {
		if value.Img != "" {
			list[key].Img = sec.Key("CDN").MustString("") + value.Img
		}
	}

	data["list"] = list
	data["total"] = models.Goods.GetTotal(maps)

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func (*goods) Get(c *gin.Context) {
	goodsId := c.Query("id")

	maps := make(map[string]interface{})
	maps["id"] = goodsId
	data := models.Goods.Get(maps)

	sec, _ := setting.Cfg.GetSection("app")
	data.Img = sec.Key("CDN").MustString("")

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
