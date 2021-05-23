package module

import (
	"encoding/json"
	"github.com/strugglerx/wechat/utils"
)

/**
 * @PROJECT_NAME wechat
 * @author  Moqi
 * @date  2021-05-20 10:41
 * @Email:str@li.cm
 **/

var UrlLinkEntity = UrlLink{}

type UrlLink struct {
	App utils.App
}

func (a *UrlLink) Init(app utils.App) *UrlLink {
	a.App = app
	return a
}

//Generate 获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景。通过该接口，可以选择生成到期失效和永久有效的小程序链接，目前仅针对国内非个人主体的小程序开放，详见获取 URL Link。
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html
func (a *UrlLink) Generate(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/generate_urllink", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}
