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

var CustomEntity = Custom{}

type Custom struct {
	App utils.App
}

func (a *Custom) Init(app utils.App) *Custom {
	a.App = app
	return a
}

func (a *Custom) Get(path string, params utils.Query, withAccessToken ...bool) ([]byte, error) {
	if len(withAccessToken) > 0 {
		response, err := utils.Get(path, params, a.App)
		return response, err
	}
	response, err := utils.Get(path, params)
	return response, err
}

func (a *Custom) PostBody(path string, body []byte, withAccessToken ...bool) ([]byte, error) {
	if len(withAccessToken) > 0 {
		response, err := utils.PostBody(path, body, a.App)
		return response, err
	}
	response, err := utils.PostBody(path, body)
	return response, err
}
