package v1

import (
	"Mini-Cloud/models"
	"Mini-Cloud/pkg/e"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"

	"Mini-Cloud/pkg/setting"
	"fmt"
	"net/http"
)

var User = &user{}

type user struct{}

func (this *user) Login(c *gin.Context) {

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
	type wxS struct {
		Errcode    int32  `json:"errcode"`
		Errmsg     string `json:"errmsg"`
		Openid     string `json:"openid"`
		SessionKey string `json:"session_key"`
	}
	resp, _ := http.Get(uri)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	data := new(wxS)
	json.Unmarshal(body, &data)

	//测试数据
	data.SessionKey = jsCode[0:6]
	data.Openid = jsCode[0:6]

	//登录注册
	user := models.User.GetByOpenid(data.Openid)

	if user.Id > 0 {
		fmt.Println(user)
	}

	data.SessionKey = jsCode[0:6]

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     e.GetMsg(code),
		"session": data.SessionKey,
	})
}

func (this *user) GetInfo(c *gin.Context) {

}
