package v1

import "github.com/gin-gonic/gin"

var Home = &home{}

type home struct{}

func (this *home) GetSlideshow(c *gin.Context) {

}

func (this *home) GetList(c *gin.Context) {

}
