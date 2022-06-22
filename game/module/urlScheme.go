package module

import (
	"github.com/strugglerx/wechat/utils"
)

/**
* @PROJECT_NAME wechat
* @author  Moqi
* @date  2021-05-20 10:41
* @Email:str@li.cm
**/

var UrlSchemeEntity = UrlScheme{}

type UrlScheme struct {
	App utils.App
}

func (a *UrlScheme) Init(app utils.App) *UrlScheme {
	a.App = app
	return a
}

//Generate JSON
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-scheme/urlscheme.generate.html
func (a *UrlScheme) Generate(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxa/generatescheme", body, a.App)
	return response, err
}
