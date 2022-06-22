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

var CloudbaseEntity = Cloudbase{}

type Cloudbase struct {
	App utils.App
}

func (a *Cloudbase) Init(app utils.App) *Cloudbase {
	a.App = app
	return a
}

//CreateSendSmsTask
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.createSendSmsTask.html
func (a *Cloudbase) CreateSendSmsTask(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/tcb/createsendsmstask", body, a.App)
	return response, err
}

//DescribeExtensionUploadInfo 描述扩展上传文件信息
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.describeExtensionUploadInfo.html
func (a *Cloudbase) DescribeExtensionUploadInfo(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/tcb/describeextensionuploadinfo", body, a.App)
	return response, err
}

//DescribeSmsRecords
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.describeSmsRecords.html
func (a *Cloudbase) DescribeSmsRecords(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/tcb/describesmsrecords", body, a.App)
	return response, err
}

//GetStatistics 获取云开发数据接口
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.getStatistics.html
func (a *Cloudbase) GetStatistics(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/tcb/getstatistics", body, a.App)
	return response, err
}

//Report 云开发通用上报接口
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.report.html
func (a *Cloudbase) Report(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/tcb/cloudbasereport", body)
	return response, err
}

//SendSms
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.sendSms.html
func (a *Cloudbase) SendSms(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/tcb/sendsms", body, a.App)
	return response, err
}

//SendSmsV2
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.sendSmsV2.html
func (a *Cloudbase) SendSmsV2(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/tcb/sendsmsv2", body, a.App)
	return response, err
}
