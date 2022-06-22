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

var DataAnalysisEntity = DataAnalysis{}

type DataAnalysis struct {
	App utils.App
}

func (a *DataAnalysis) Init(app utils.App) *DataAnalysis {
	a.App = app
	return a
}

//GetGameAnalysisData 获取小游戏分析数据。
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/data-analysis/analysis.getGameAnalysisData.html
func (a *DataAnalysis) GetGameAnalysisData(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/datacube/getgameanalysisdata", body, a.App)
	return response, err
}
