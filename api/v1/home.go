package v1

import (
	"github.com/colinlet/mini_cloud/entity"
	"github.com/colinlet/mini_cloud/models"
	"github.com/colinlet/mini_cloud/pkg/e"
	"github.com/colinlet/mini_cloud/pkg/setting"
	"github.com/colinlet/mini_cloud/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var Home = &home{}

type home struct{}

func (*home) GetSlideshow(c *gin.Context) {
	list := models.Slideshow.GetList()
	sec, _ := setting.Cfg.GetSection("app")
	for key, value := range list {
		if value.Img != "" {
			list[key].Img = sec.Key("CDN").MustString("") + value.Img
		}
	}
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"list": list,
	})
}

func (*home) GetList(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	maps["status"] = 1
	raws := models.Goods.GetList(util.GetPage(c), setting.PageSize*2, maps)

	//随机商品
	count := 0
	total := len(raws) / 2
	list := make([]entity.Goods, total)
	day, _ := strconv.Atoi(time.Now().Format("02"))
	sec, _ := setting.Cfg.GetSection("app")
	for _, value := range raws {
		if count < total && day%2 == int(value.Id)%2 {
			if value.Img != "" {
				value.Img = sec.Key("CDN").MustString("") + value.Img
			}
			list[count] = value
			count++
		}
	}

	data["list"] = list
	data["total"] = models.Goods.GetTotal(maps) / 2

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
