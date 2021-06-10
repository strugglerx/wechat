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

var LockStepEntity = LockStep{}

type LockStep struct {
	App utils.App
}

func (a *LockStep) Init(app utils.App) *LockStep {
	a.App = app
	return a
}

//CreateGameRoom 第三方后台创建帧同步游戏房间
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/lock-step/lock-step.createGameRoom.html
func (a *LockStep) CreateGameRoom(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/createwxagameroom", body, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//GetGameFrame 分片拉取对局游戏帧
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/lock-step/lock-step.getGameFrame.html
func (a *LockStep) GetGameFrame(accessInfo, beginFrameId, endFrameId string) (interface{}, error) {
	var result interface{}
	params := utils.Query{
		"access_info":    accessInfo,
		"begin_frame_id": beginFrameId,
		"end_frame_id":   endFrameId,
	}
	response, err := utils.Get("/wxa/getwxagameframe", params, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//GetGameIdentityInfo 获取对局玩家位次信息
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/lock-step/lock-step.getGameIdentityInfo.html
func (a *LockStep) GetGameIdentityInfo(accessInfo string) (interface{}, error) {
	var result interface{}
	params := utils.Query{
		"access_info": accessInfo,
	}
	response, err := utils.Get("/wxa/getwxagameidentityinfo", params, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//GetGameRoomInfo 获取指定房间信息
//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/lock-step/lock-step.getGameRoomInfo.html
func (a *LockStep) GetGameRoomInfo(accessInfo string) (interface{}, error) {
	var result interface{}
	params := utils.Query{
		"access_info": accessInfo,
	}
	response, err := utils.Get("/wxa/getwxagameroominfo", params, a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
