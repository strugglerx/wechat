package mp

import (
	"errors"
	"github.com/strugglerx/wechat/utils"
	"github.com/tidwall/gjson"
	"time"
)

//GetOauthAccessToken 获取Oauth2accessToken(网页版本) //不可以缓存，每个用户的登陆凭证不一样
func (m *Mp) GetOauthAccessToken(code string) *OauthToken {
	if m.Oauth2Token == nil {
		m.Oauth2Token = &OauthToken{
			Token: "", UpdateTime: 0,
		}
	}
	params := utils.Query{
		"appid":      m.Appid,
		"secret":     m.Secret,
		"code":       code,
		"grant_type": "authorization_code",
	}
	response,_ := utils.Get("/sns/oauth2/access_token",params)
	m.Oauth2Token = &OauthToken{
		Token:        gjson.Get(string(response), "access_token").String(),
		ExpiresIn:    int(gjson.Get(string(response), "expires_in").Int()),
		RefreshToken: gjson.Get(string(response), "refresh_token").String(),
		OpenId:       gjson.Get(string(response), "openid").String(),
		Scope:        gjson.Get(string(response), "scope").String(),
		UpdateTime:   int(time.Now().Unix()),
	}
	return m.Oauth2Token
}

//RefreshOauthAccessToken 刷新Oauth2accessToken
func (m *Mp) RefreshOauthAccessToken(refreshToken string) (interface{}, error) {
	params := utils.Query{
		"appid":         m.Appid,
		"refresh_token": refreshToken,
		"grant_type":    "client_credential",
	}
	responseString, err := utils.Get("/cgi-bin/token",params)
	if err != nil {
		return nil, err
	}
	return responseString, nil
}

//GetOauthUserInfo 获取oauth用户真实信息
func (m *Mp) GetOauthUserInfo(token, openid string) (string, error) {
	params := utils.Query{
		"access_token": token,
		"openid":       openid,
		"lang":         "zh_CN",
	}
	response, err := utils.Get("/sns/userinfo",params)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

//VerifyOauthToken 验证Oauth2accessToken是否有效
func (m *Mp) VerifyOauthToken(openid string) (interface{}, error) {
	params := utils.Query{
		"access_token": m.Oauth2Token.Token,
		"openid":       openid,
	}
	response, err := utils.Get("/sns/auth",params)
	if err != nil || gjson.Get(string(response), "errcode").Int() != 0 {
		return nil, errors.New("remain")
	}
	return string(response), nil
}
