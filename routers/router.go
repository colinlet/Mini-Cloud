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
		apiv1.GET("/category/getList", v1.Category.GetList) //获取分类
		apiv1.GET("/goods/getList", v1.Goods.GetList)       //获取商品列表
		apiv1.GET("/goods", v1.Goods.Get)                   //获取商品
		apiv1.GET("/user/login", v1.User.Login)             //用户登录
		apiv1.GET("/user/info", v1.User.GetInfo)            //获取用户信息
		apiv1.POST("/user/info", v1.User.SetInfo)           //设置用户信息

		//首页
		apiv1.GET("/home/slideshow", v1.Home.GetSlideshow) //获取轮播图
		apiv1.GET("/home/list", v1.Home.GetList)           //获取首页商品

		//地址
		apiv1.GET("/address/list", nil) //获取地址
		//新增地址
		//编辑地址

		//获取订单
	}

	//图片资源
	router.StaticFS("/images", http.Dir("./images"))

	return router
}
