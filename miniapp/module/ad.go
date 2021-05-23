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

var AdEntity = Ad{}

type Ad struct {
	App utils.App
}

func (a *Ad) Init(app utils.App) *Ad {
	a.App = app
	return a
}
