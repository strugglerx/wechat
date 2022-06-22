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
func (a *DataAnalysis) GetDailySummary(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/datacube/getweanalysisappiddailysummarytrend", body, a.App)
	return response, err
}

//GetPerformanceData
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getPerformanceData.html
func (a *DataAnalysis) GetPerformanceData(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxa/business/performance/boot", body, a.App)
	return response, err
}

//GetUserPortrait
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html
func (a *DataAnalysis) GetUserPortrait(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/datacube/getweanalysisappiduserportrait", body, a.App)
	return response, err
}

//GetVisitDistribution 获取用户小程序访问分布数据
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitDistribution.html
func (a *DataAnalysis) GetVisitDistribution(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/datacube/getweanalysisappidvisitdistribution", body, a.App)
	return response, err
}

//GetVisitPage
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitPage.html
func (a *DataAnalysis) GetVisitPage(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/datacube/getweanalysisappidvisitpage", body, a.App)
	return response, err
}
