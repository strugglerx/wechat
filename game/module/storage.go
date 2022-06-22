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
func (a *Storage) RemoveUserStorage(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxa/remove_user_storage", body, a.App)
	return response, err
}

//SetUserInteractiveData 删除已经上报到微信的key-value数据
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/data/storage.setUserInteractiveData.html
func (a *Storage) SetUserInteractiveData(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxa/setuserinteractivedata", body, a.App)
	return response, err
}

//SetUserStorage 删除已经上报到微信的key-value数据
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/data/storage.setUserStorage.html
func (a *Storage) SetUserStorage(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxa/set_user_storage", body, a.App)
	return response, err
}
