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

var PluginManagementEntity = PluginManagement{}

type PluginManagement struct {
	App utils.App
}

func (a *PluginManagement) Init(app utils.App) *PluginManagement {
	a.App = app
	return a
}

//ApplyPlugin 向插件开发者发起使用插件的申请
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.applyPlugin.html
func (a *PluginManagement) ApplyPlugin(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/plugin", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetPluginDevApplyList
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.getPluginDevApplyList.html
func (a *PluginManagement) GetPluginDevApplyList(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/devplugin", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//GetPluginList 查询已添加的插件
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.getPluginList.html
func (a *PluginManagement) GetPluginList(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/plugin", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//SetDevPluginApplyStatus 或
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.setDevPluginApplyStatus.html
func (a *PluginManagement) SetDevPluginApplyStatus(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/devplugin", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}

//UnbindPlugin 删除已添加的插件
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.unbindPlugin.html
func (a *PluginManagement) UnbindPlugin(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/plugin", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result , nil
}
