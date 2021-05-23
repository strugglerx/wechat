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

var SafetyControlCapabilityEntity = SafetyControlCapability{}

type SafetyControlCapability struct {
    App utils.App
}

func (a *SafetyControlCapability) Init(app utils.App) *SafetyControlCapability {
    a.App = app
    return a
}

//GetUserRiskRank 根据提交的用户信息数据获取用户的安全等级 risk_rank，无需用户授权。
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/safety-control-capability/riskControl.getUserRiskRank.html
func (a *SafetyControlCapability) GetUserRiskRank(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/getuserriskrank", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}
