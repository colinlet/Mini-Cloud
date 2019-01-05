package v1

import (
	"Mini-Cloud/entity"
	"Mini-Cloud/models"
	"Mini-Cloud/pkg/e"
	"Mini-Cloud/pkg/setting"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var User = &user{}

type user struct{}

func (*user) Login(c *gin.Context) {
	jsCode := c.Query("code")
	sec, _ := setting.Cfg.GetSection("app")
	uri := fmt.Sprintf(
		"%s?appid=%s&secret=%s&js_code=%s&grant_type=%s",
		sec.Key("WX_CTS").MustString(""),
		sec.Key("WX_APPID").MustString(""),
		sec.Key("WX_SECRET").MustString(""),
		jsCode,
		"authorization_code",
	)

	//获取微信信息
	resp, _ := http.Get(uri)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	data := new(entity.Code2Session)
	json.Unmarshal(body, &data)

	//测试数据
	data.SessionKey = jsCode[0:12]
	data.Openid = jsCode[len(jsCode)-12 : len(jsCode)-0]

	//登录注册
	user := models.User.GetByOpenid(data.Openid)
	if user.Id < 1 {
		userData := &entity.User{
			Openid:  data.Openid,
			Session: data.SessionKey,
		}
		models.User.Insert(userData)
	} else {
		models.User.Update(user.Id, map[string]interface{}{"session": data.SessionKey})
	}

	//返回数据
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     e.GetMsg(code),
		"session": data.SessionKey,
	})
}

func (*user) GetInfo(c *gin.Context) {
	session := c.Query("session")
	info := models.User.GetInfo(session)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": info,
	})
}

func (*user) SetInfo(c *gin.Context) {
	session := c.PostForm("session")
	avatarUrl := c.PostForm("avatarUrl")
	nickName := c.PostForm("nickName")

	user := models.User.GetBySession(session)
	models.User.UpdateInfo(user.Id, map[string]interface{}{"avatar_url": avatarUrl, "nick_name": nickName})

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
