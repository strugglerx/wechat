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

//Generate JSON
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html
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
