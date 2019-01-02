package v1

import (
	"Mini-Cloud/models"
	"Mini-Cloud/pkg/e"
	"Mini-Cloud/pkg/setting"
	"Mini-Cloud/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Category = &category{}

type category struct{}

func (*category) GetList(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS

	data["lists"] = models.Category.GetList(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.Category.GetTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
