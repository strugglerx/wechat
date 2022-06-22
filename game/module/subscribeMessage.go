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

var SubscribeMessageEntity = SubscribeMessage{}

type SubscribeMessage struct {
	App utils.App
}

func (a *SubscribeMessage) Init(app utils.App) *SubscribeMessage {
	a.App = app
	return a
}

//Send 发送订阅消息
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
func (a *SubscribeMessage) Send(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/cgi-bin/message/subscribe/send", body, a.App)
	return response, err
}
