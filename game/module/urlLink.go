package module

import (
	"github.com/strugglerx/wechat/v2/utils"
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

//Generate JSON
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html
func (a *UrlLink) Generate(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxa/generate_urllink", body, a.App)
	return response, err
}
