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

var TemplateMessageEntity = TemplateMessage{}

type TemplateMessage struct {
    App utils.App
}

func (a *TemplateMessage) Init(app utils.App) *TemplateMessage {
    a.App = app
    return a
}

//AddTemplate 
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.addTemplate.html
func (a *TemplateMessage) AddTemplate(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/cgi-bin/wxopen/template/add", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//DeleteTemplate 
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.deleteTemplate.html
func (a *TemplateMessage) DeleteTemplate(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/cgi-bin/wxopen/template/del", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetTemplateLibraryById 
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.getTemplateLibraryById.html
func (a *TemplateMessage) GetTemplateLibraryById(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/cgi-bin/wxopen/template/library/get", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetTemplateLibraryList 
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.getTemplateLibraryList.html
func (a *TemplateMessage) GetTemplateLibraryList(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/cgi-bin/wxopen/template/library/list", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetTemplateList 
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.getTemplateList.html
func (a *TemplateMessage) GetTemplateList(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/cgi-bin/wxopen/template/list", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//Send 
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.send.html
func (a *TemplateMessage) Send(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/cgi-bin/message/wxopen/template/send", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}
