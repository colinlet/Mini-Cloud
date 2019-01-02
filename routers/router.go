package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"Mini-Cloud/pkg/setting"
	"Mini-Cloud/routers/api/v1"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/getCategory", v1.Category.GetList)

		apiv1.GET("/test", v1.GetTest)

		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	//图片资源
	router.StaticFS("/images", http.Dir("./images"))

	return router
}
