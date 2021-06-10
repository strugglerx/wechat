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

var StorageEntity = Storage{}

type Storage struct {
	App utils.App
}

func (a *Storage) Init(app utils.App) *Storage {
	a.App = app
	return a
}

//RemoveUserStorage 删除已经上报到微信的key-value数据
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/data/storage.removeUserStorage.html
func (a *Storage) RemoveUserStorage(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/remove_user_storage", body, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//SetUserInteractiveData 删除已经上报到微信的key-value数据
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/data/storage.setUserInteractiveData.html
func (a *Storage) SetUserInteractiveData(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/setuserinteractivedata", body, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//SetUserStorage 删除已经上报到微信的key-value数据
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/data/storage.setUserStorage.html
func (a *Storage) SetUserStorage(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/set_user_storage", body, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
