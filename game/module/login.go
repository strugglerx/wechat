package module

import (
	"encoding/json"
	"errors"

	"github.com/strugglerx/wechat/v2/utils"
)

/**
* @PROJECT_NAME wechat
* @author  Moqi
* @date  2021-05-20 10:41
* @Email:str@li.cm
**/

var LoginEntity = Login{}

type Login struct {
	App utils.App
}

func (a *Login) Init(app utils.App) *Login {
	a.App = app
	return a
}

//CheckSessionKey 校验服务器所保存的登录态 session_key 是否合法。为了保持 session_key 私密性，接口不明文传输 session_key，而是通过校验登录态签名完成。
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/login/auth.checkSessionKey.html
func (a *Login) CheckSessionKey(openid, signature, sigMethod string) (utils.Response, error) {
	params := utils.Query{
		"openid":     openid,
		"signature":  signature,
		"sig_method": sigMethod,
	}
	response, err := utils.Get("/wxa/checksession", params, a.App)
	return response, err
}

//Code2Session 登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。更多使用方法详见 小程序登录。
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (a *Login) Code2Session(code string) (User, error) {
	var result Session
	user := User{}
	params := utils.Query{
		"appid":      a.App.GetConfig().Appid,
		"secret":     a.App.GetConfig().Secret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	response, err := utils.Get("/sns/jscode2session", params)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return user, err
	}
	if result.Errcode == 0 {
		user := User{Session: result.SessionKey, Openid: result.Openid, Appid: a.App.GetConfig().Appid, Unionid: result.Unionid, Status: true}
		return user, nil
	}
	return user, errors.New(string(response))
}
