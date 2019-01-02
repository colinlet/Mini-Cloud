package v1

import (
	"Mini-Cloud/models"
	"Mini-Cloud/pkg/e"
	"Mini-Cloud/pkg/setting"
	"Mini-Cloud/pkg/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/errors/fmt"
	"net/http"
)

func GetTest(c *gin.Context) {

	fmt.Println("this is test")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS

	data["lists"] = models.GetTest(util.GetPage(c), setting.PageSize, maps)
	data["total"] = 0

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
