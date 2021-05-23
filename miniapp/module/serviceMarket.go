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

var ServiceMarketEntity = ServiceMarket{}

type ServiceMarket struct {
	App utils.App
}

func (a *ServiceMarket) Init(app utils.App) *ServiceMarket {
	a.App = app
	return a
}

//InvokeService 调用服务平台提供的服务
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/service-market/serviceMarket.invokeService.html
func (a *ServiceMarket) InvokeService(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/servicemarket", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	return result , err
}
