package module

import (
	"encoding/json"
	"errors"

	"github.com/strugglerx/wechat/v2/utils"
)

/**
* @PROJECT_NAME wechat
* @author  Moqi
* @date  2021-05-20 10:41
* @Email:str@li.cm
**/

var OperationEntity = Operation{}

type Operation struct {
	App utils.App
}

func (a *Operation) Init(app utils.App) *Operation {
	a.App = app
	return a
}

//GetDomainInfo 查询域名配置
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getDomainInfo.html
func (a *Operation) GetDomainInfo(action string) (utils.Response, error) {
	response, err := utils.Get("/wxa/getwxadevinfo", utils.Query{
		"action": action,
	}, a.App)
	return response, err
}

//GetFeedback 获取用户反馈列表
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getFeedback.html
func (a *Operation) GetFeedback(types, page, num string) (utils.Response, error) {
	response, err := utils.Get("/wxaapi/feedback/list", utils.Query{
		"type": types,
		"page": page,
		"num":  num,
	}, a.App)
	return response, err
}

//GetFeedbackmedia
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getFeedbackmedia.html
func (a *Operation) GetFeedbackmedia(recordId, mediaId string) (utils.Response, error) {
	response, err := utils.Get("/cgi-bin/media/getfeedbackmedia", utils.Query{
		"record_id": recordId,
		"media_id":  mediaId,
	}, a.App)
	return response, err
}

//GetJsErrDetail 错误查询详情
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getJsErrDetail.html
func (a *Operation) GetJsErrDetail(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxaapi/log/jserr_detail", body, a.App)
	return response, err
}

//GetJsErrList 错误查询列表
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getJsErrList.html
func (a *Operation) GetJsErrList(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxaapi/log/jserr_list", body, a.App)
	return response, err
}

//GetJsErrSearch
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getJsErrSearch.html
func (a *Operation) GetJsErrSearch(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxaapi/log/jserr_search", body, a.App)
	return response, err
}

//GetPerformance 性能监控
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getPerformance.html
func (a *Operation) GetPerformance(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxaapi/log/get_performance", body, a.App)
	return response, err
}

//GetSceneList 获取访问来源
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getSceneList.html
func (a *Operation) GetSceneList() (utils.Response, error) {
	response, err := utils.Get("/wxaapi/log/get_scene", a.App)
	return response, err
}

//GetVersionList 获取客户端版本
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getVersionList.html
func (a *Operation) GetVersionList() (utils.Response, error) {
	response, err := utils.Get("/wxaapi/log/get_client_version", a.App)
	return response, err
}

//RealtimelogSearch 实时日志查询
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.realtimelogSearch.html
func (a *Operation) RealtimelogSearch(body []byte) (utils.Response, error) {
	var query utils.Query
	err := json.Unmarshal(body, &query)
	if err != nil {
		return nil, errors.New("params fail")
	}

	response, err := utils.Get("/wxaapi/userlog/userlog_search", query, a.App)
	return response, err
}
