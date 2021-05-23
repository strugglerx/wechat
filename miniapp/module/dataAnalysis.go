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

var DataAnalysisEntity = DataAnalysis{}

type DataAnalysis struct {
	App utils.App
}

func (a *DataAnalysis) Init(app utils.App) *DataAnalysis {
	a.App = app
	return a
}

//GetDailySummary 获取用户访问小程序数据概况
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getDailySummary.html
func (a *DataAnalysis) GetDailySummary(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/datacube/getweanalysisappiddailysummarytrend", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetPerformanceData
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getPerformanceData.html
func (a *DataAnalysis) GetPerformanceData(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/business/performance/boot", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetUserPortrait
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html
func (a *DataAnalysis) GetUserPortrait(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/datacube/getweanalysisappiduserportrait", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetVisitDistribution 获取用户小程序访问分布数据
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitDistribution.html
func (a *DataAnalysis) GetVisitDistribution(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/datacube/getweanalysisappidvisitdistribution", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetVisitPage
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitPage.html
func (a *DataAnalysis) GetVisitPage(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/datacube/getweanalysisappidvisitpage", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}
