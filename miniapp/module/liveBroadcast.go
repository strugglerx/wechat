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

var LivebroadcastEntity = Livebroadcast{}

type Livebroadcast struct {
	App utils.App
}

func (a *Livebroadcast) Init(app utils.App) *Livebroadcast {
	a.App = app
	return a
}
