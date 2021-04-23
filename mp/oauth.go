package mp

import (
	"errors"
	"github.com/strugglerx/wechat/tool"
	"github.com/tidwall/gjson"
	"time"
)

//获取Oauth2accessToken(网页版本) //不可以缓存，每个用户的登陆凭证不一样
func (m *Mp) GetOauthAccessToken(code string) *OauthToken {
	if m.Oauth2Token == nil {
		m.Oauth2Token = &OauthToken{
			Token: "", UpdateTime: 0,
		}
	}
	params := param{
		"appid":      m.Appid,
		"secret":     m.Secret,
		"code":       code,
		"grant_type": "authorization_code",
	}
	responseString,_ := tool.Get("/sns/oauth2/access_token",params)
	m.Oauth2Token = &OauthToken{
		Token:        gjson.Get(responseString, "access_token").String(),
		ExpiresIn:    int(gjson.Get(responseString, "expires_in").Int()),
		RefreshToken: gjson.Get(responseString, "refresh_token").String(),
		OpenId:       gjson.Get(responseString, "openid").String(),
		Scope:        gjson.Get(responseString, "scope").String(),
		UpdateTime:   int(time.Now().Unix()),
	}
	return m.Oauth2Token
}

//刷新Oauth2accessToken
func (m *Mp) RefreshOauthAccessToken(refreshToken string) (interface{}, error) {
	params := param{
		"appid":         m.Appid,
		"refresh_token": refreshToken,
		"grant_type":    "client_credential",
	}
	responseString, err := tool.Get("/cgi-bin/token",params)
	if err != nil {
		return nil, err
	}
	return responseString, nil
}

//获取oauth用户真实信息
func (m *Mp) GetOauthUserInfo(token, openid string) (string, error) {
	params := param{
		"access_token": token,
		"openid":       openid,
		"lang":         "zh_CN",
	}
	responseString, err := tool.Get("/sns/userinfo",params)
	if err != nil {
		return "", err
	}
	return responseString, nil
}

//验证Oauth2accessToken是否有效
func (m *Mp) VerifyOauthToken(openid string) (interface{}, error) {
	params := param{
		"access_token": m.Oauth2Token.Token,
		"openid":       openid,
	}
	responseString, err := tool.Get("/sns/auth",params)
	if err != nil || gjson.Get(responseString, "errcode").Int() != 0 {
		return nil, errors.New("remain")
	}
	return responseString, nil
}
