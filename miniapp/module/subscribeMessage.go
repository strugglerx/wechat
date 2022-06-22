package module

import (
	"strconv"

	"github.com/strugglerx/wechat/v2/utils"
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

//AddTemplate 组合模板并添加至帐号下的个人模板库
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.addTemplate.html
func (a *SubscribeMessage) AddTemplate(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxaapi/newtmpl/addtemplate", body, a.App)
	return response, err
}

//DeleteTemplate 删除帐号下的个人模板
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.deleteTemplate.html
func (a *SubscribeMessage) DeleteTemplate(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxaapi/newtmpl/deltemplate", body, a.App)
	return response, err
}

//GetCategory 获取小程序账号的类目
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getCategory.html
func (a *SubscribeMessage) GetCategory() (utils.Response, error) {
	response, err := utils.Get("/wxaapi/newtmpl/getcategory", a.App)
	return response, err
}

//GetPubTemplateKeyWordsById 获取模板标题下的关键词列表
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateKeyWordsById.html
func (a *SubscribeMessage) GetPubTemplateKeyWordsById(tid string) (utils.Response, error) {
	param := utils.Query{
		"tid": tid,
	}
	response, err := utils.Get("/wxaapi/newtmpl/getpubtemplatekeywords", param, a.App)
	return response, err
}

//GetPubTemplateTitleList 获取帐号所属类目下的公共模板标题
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateTitleList.html
func (a *SubscribeMessage) GetPubTemplateTitleList(ids string, start, limit int) (utils.Response, error) {
	param := utils.Query{
		"ids":   ids,
		"start": strconv.Itoa(start),
		"limit": strconv.Itoa(limit),
	}
	response, err := utils.Get("/wxaapi/newtmpl/getpubtemplatetitles", param, a.App)
	return response, err
}

//GetTemplateList 获取当前帐号下的个人模板列表
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html
func (a *SubscribeMessage) GetTemplateList() (utils.Response, error) {
	response, err := utils.Get("/wxaapi/newtmpl/gettemplate", a.App)
	return response, err
}

//Send 发送订阅消息
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
func (a *SubscribeMessage) Send(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/cgi-bin/message/subscribe/send", body, a.App)
	return response, err
}
