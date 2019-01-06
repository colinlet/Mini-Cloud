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

func (*user) AddAddress(c *gin.Context) {
	session := c.PostForm("session")
	userName := c.PostForm("user_name")
	postalCode := c.PostForm("postal_code")
	provinceName := c.PostForm("province_name")
	cityName := c.PostForm("city_name")
	countyName := c.PostForm("county_name")
	detailInfo := c.PostForm("detail_info")
	nationalCode := c.PostForm("national_code")
	telNumber := c.PostForm("tel_number")

	user := models.User.GetBySession(session)

	models.Address.UpdateStatus(user.Id, map[string]interface{}{"is_use": 0})

	address := &entity.UserAddress{
		Mid:          user.Id,
		UserName:     userName,
		PostalCode:   postalCode,
		ProvinceName: provinceName,
		CityName:     cityName,
		CountyName:   countyName,
		DetailInfo:   detailInfo,
		NationalCode: nationalCode,
		TelNumber:    telNumber,
		IsUse:        1,
	}
	models.Address.Insert(address)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func (*user) GetAddress(c *gin.Context) {
	session := c.Query("session")
	user := models.User.GetBySession(session)
	list := models.Address.GetList(user.Id)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"list": list,
	})
}

func (*user) ChooseAddress(c *gin.Context) {
	session := c.PostForm("session")
	id := c.PostForm("id")

	user := models.User.GetBySession(session)
	models.Address.UpdateStatus(user.Id, map[string]interface{}{"is_use": 0})
	models.Address.Choose(id)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func (*user) CurrentAddress(c *gin.Context) {
	session := c.Query("session")
	user := models.User.GetBySession(session)

	data := models.Address.GetCurrent(user.Id)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code":       code,
		"msg":        e.GetMsg(code),
		"data":       data,
		"pay_notice": "支付还未接入，暂时无法结算",
	})
}
