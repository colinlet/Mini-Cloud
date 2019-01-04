package v1

import "github.com/gin-gonic/gin"

var User = &user{}

type user struct{}

func (this *user) GetInfo(c *gin.Context) {

}
