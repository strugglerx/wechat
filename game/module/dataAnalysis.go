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

//GetGameAnalysisData 获取小游戏分析数据。
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/data-analysis/analysis.getGameAnalysisData.html
func (a *DataAnalysis) GetGameAnalysisData(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/datacube/getgameanalysisdata", body, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
