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
func (a *TemplateMessage) AddTemplate(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/cgi-bin/wxopen/template/add", body, a.App)
	return response, err
}

//DeleteTemplate
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.deleteTemplate.html
func (a *TemplateMessage) DeleteTemplate(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/cgi-bin/wxopen/template/del", body, a.App)
	return response, err
}

//GetTemplateLibraryById
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.getTemplateLibraryById.html
func (a *TemplateMessage) GetTemplateLibraryById(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/cgi-bin/wxopen/template/library/get", body, a.App)
	return response, err
}

//GetTemplateLibraryList
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.getTemplateLibraryList.html
func (a *TemplateMessage) GetTemplateLibraryList(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/cgi-bin/wxopen/template/library/list", body, a.App)
	return response, err
}

//GetTemplateList
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.getTemplateList.html
func (a *TemplateMessage) GetTemplateList(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/cgi-bin/wxopen/template/list", body, a.App)
	return response, err
}

//Send
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.send.html
func (a *TemplateMessage) Send(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/cgi-bin/message/wxopen/template/send", body, a.App)
	return response, err
}
