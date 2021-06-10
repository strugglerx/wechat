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
func (a *UrlScheme) Generate(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/generatescheme", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}
