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
		//首页
		apiv1.GET("/home/slideshow", v1.Home.GetSlideshow) //获取轮播图
		apiv1.GET("/home/list", v1.Home.GetList)           //获取首页商品

		//分类
		apiv1.GET("/category/getList", v1.Category.GetList) //获取分类

		//地址
		apiv1.GET("/address/list", nil) //获取地址
		//新增地址
		//编辑地址

		//获取订单

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
