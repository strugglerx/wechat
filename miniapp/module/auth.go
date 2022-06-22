package module

import (
	"errors"

	"github.com/strugglerx/wechat/utils"
)

/**
 * @PROJECT_NAME wechat
 * @author  Moqi
 * @date  2021-05-20 10:41
 * @Email:str@li.cm
 **/

var AuthEntity = Auth{}

type Auth struct {
	App utils.App
}

func (a *Auth) Init(app utils.App) *Auth {
	a.App = app
	return a
}

//Code2Session 登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。更多使用方法详见 小程序登录。
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (a *Auth) Code2Session(code string) (Session, error) {
	var result Session
	params := utils.Query{
		"appid":      a.App.GetConfig().Appid,
		"secret":     a.App.GetConfig().Secret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	response, err := utils.Get("/sns/jscode2session", params)
	if err != nil {
		return result, err
	}
	err = response.Unmarshal(&result)
	if err != nil {
		return result, err
	}
	if result.Errcode == 0 {
		return result, nil
	}
	return result, errors.New(result.ErrMsg)
}

//getPaidUnionId 用户支付完成后，获取该用户的 UnionId，无需用户授权。本接口支持第三方平台代理查询。
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html
func (a *Auth) GetPaidUnionId(openid string) (unionId string, err error) {
	var result Session
	params := utils.Query{
		"openid": openid,
	}
	response, err := utils.Get("/wxa/getpaidunionid", params, a.App)
	if err != nil {
		return unionId, err
	}
	err = response.Unmarshal(&result)
	if err != nil {
		return unionId, err
	}
	if result.Errcode == 0 {
		return result.Unionid, nil
	}
	return unionId, errors.New(string(response))
}
