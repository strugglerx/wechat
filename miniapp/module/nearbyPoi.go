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

var NearbyPoiEntity = NearbyPoi{}

type NearbyPoi struct {
	App utils.App
}

func (a *NearbyPoi) Init(app utils.App) *NearbyPoi {
	a.App = app
	return a
}

//Add 添加地点
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.add.html
func (a *NearbyPoi) Add(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/addnearbypoi", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//Delete 删除地点
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.delete.html
func (a *NearbyPoi) Delete(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/delnearbypoi", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetList 查看地点列表
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.getList.html
func (a *NearbyPoi) GetList(page,page_rows string) (interface{}, error) {
	var result interface{}
	response, err := utils.Get("/wxa/getnearbypoilist", utils.Query{
		"page":page,
		"page_rows":page_rows,
	},a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//SetShowStatus
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.setShowStatus.html
func (a *NearbyPoi) SetShowStatus(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/setnearbypoishowstatus", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}
