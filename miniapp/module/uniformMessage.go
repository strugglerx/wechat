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

var UniformMessageEntity = UniformMessage{}

type UniformMessage struct {
	App utils.App
}

func (a *UniformMessage) Init(app utils.App) *UniformMessage {
	a.App = app
	return a
}

//Send 下发小程序和公众号统一的服务消息
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html
func (a *UniformMessage) Send(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/cgi-bin/message/wxopen/template/uniform_send", body, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
