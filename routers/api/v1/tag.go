package v1

import (
	"Mini-Cloud/models"
	"Mini-Cloud/pkg/e"
	"Mini-Cloud/pkg/setting"
	"Mini-Cloud/pkg/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
