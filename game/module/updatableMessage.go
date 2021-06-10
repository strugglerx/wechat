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

var UpdatableMessageEntity = UpdatableMessage{}

type UpdatableMessage struct {
	App utils.App
}

func (a *UpdatableMessage) Init(app utils.App) *UpdatableMessage {
	a.App = app
	return a
}

//CreateActivityId 创建被分享动态消息或私密消息的 activity_id。详见动态消息。
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.createActivityId.html
func (a *UpdatableMessage) CreateActivityId(openid, unionid string) (interface{}, error) {
	var result interface{}
	params := utils.Query{
		"openid":  openid,
		"unionid": unionid,
	}
	response, err := utils.Get("/cgi-bin/message/wxopen/activityid/create", params, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//SetUpdatableMsg 修改被分享的动态消息。详见动态消息。
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.setUpdatableMsg.html
func (a *UpdatableMessage) SetUpdatableMsg(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/cgi-bin/message/wxopen/updatablemsg/send", body, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
