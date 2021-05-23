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

var CloudbaseEntity = Cloudbase{}

type Cloudbase struct {
	App utils.App
}

func (a *Cloudbase) Init(app utils.App) *Cloudbase {
	a.App = app
	return a
}

//AddDelayedFunctionTask 延时调用云函数
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.addDelayedFunctionTask.html
func (a *Cloudbase) AddDelayedFunctionTask(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/tcb/adddelayedfunctiontask", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//CreatePressureTest 创建压测任务
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.createPressureTest.html
func (a *Cloudbase) CreatePressureTest(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/tcb/createpressuretesttask", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//CreateSendSmsTask 创建发短信任务。发送的短信支持打开云开发静态网站 H5，进而在 H5 里可以打开小程序。
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.createSendSmsTask.html
func (a *Cloudbase) CreateSendSmsTask(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/tcb/createsendsmstask", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//DescribeExtensionUploadInfo 描述扩展上传文件信息
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.describeExtensionUploadInfo.html
func (a *Cloudbase) DescribeExtensionUploadInfo(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/tcb/describeextensionuploadinfo", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//DescribeSmsRecords 查询 2 个月内的短信记录
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.describeSmsRecords.html
func (a *Cloudbase) DescribeSmsRecords(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/tcb/describesmsrecords", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetOpenData 换取 cloudID 对应的开放数据
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.getOpenData.html
func (a *Cloudbase) GetOpenData(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/getopendata", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetPressureTestReport 获取压测报告
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.getPressureTestReport.html
func (a *Cloudbase) GetPressureTestReport(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/tcb/getpressuretestreport", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetPressureTestStatus 获取压测状态
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.getPressureTestStatus.html
func (a *Cloudbase) GetPressureTestStatus(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/tcb/getpressureteststatus", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetStatistics 获取云开发数据接口
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.getStatistics.html
func (a *Cloudbase) GetStatistics(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/tcb/getstatistics", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetVoIPSign 获取实时语音签名
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.getVoIPSign.html
func (a *Cloudbase) GetVoIPSign(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/getvoipsign", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//Report 云开发通用上报接口
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.report.html
func (a *Cloudbase) Report(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/tcb/cloudbasereport", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//SendSms 发送支持打开云开发静态网站的短信，该 H5 可以打开小程序。
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.sendSms.html
func (a *Cloudbase) SendSms(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/tcb/sendsms", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//SendSmsV2 发送携带 URL Link 的短信
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/cloudbase/cloudbase.sendSmsV2.html
func (a *Cloudbase) SendSmsV2(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/tcb/sendsmsv2", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}
