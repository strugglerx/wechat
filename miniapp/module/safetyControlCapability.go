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
func (a *SafetyControlCapability) GetUserRiskRank(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxa/getuserriskrank", body, a.App)
	return response, err
}
