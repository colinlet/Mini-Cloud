package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/colinlet/Mini-Cloud/pkg/setting"
	"github.com/colinlet/Mini-Cloud/routers/api/v1"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/home/slideshow", v1.Home.GetSlideshow)        //获取轮播图
		apiv1.GET("/home/list", v1.Home.GetList)                  //获取首页商品
		apiv1.GET("/category/getList", v1.Category.GetList)       //获取分类
		apiv1.GET("/goods/getList", v1.Goods.GetList)             //获取商品列表
		apiv1.GET("/goods", v1.Goods.Get)                         //获取商品
		apiv1.GET("/user/login", v1.User.Login)                   //用户登录
		apiv1.GET("/user/info", v1.User.GetInfo)                  //获取用户信息
		apiv1.POST("/user/info", v1.User.SetInfo)                 //设置用户信息
		apiv1.GET("/user/address", v1.User.GetAddress)            //获取地址
		apiv1.POST("/user/address", v1.User.AddAddress)           //新增地址
		apiv1.POST("/user/chooseAddress", v1.User.ChooseAddress)  //选中地址
		apiv1.GET("/user/currentAddress", v1.User.CurrentAddress) //获取当前地址
	}

	router.StaticFS("/images", http.Dir("./images")) //图片资源

	return router
}
