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

var WxacodeEntity = Wxacode{}

type Wxacode struct {
	App utils.App
}

func (a *Wxacode) Init(app utils.App) *Wxacode {
	a.App = app
	return a
}

//CreateQRCode 获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制，详见获取二维码。
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html
func (a *Wxacode) CreateQRCode(path, scene string) (interface{}, error) {
	var result interface{}
	body := Acode{
		Scene: scene,
		Page:  path, //"pages/index/index",
		Width: 280,
	}
	response,err := utils.PostBody("/cgi-bin/wxaapp/createwxaqrcode", utils.JsonToByte(body) ,a.App)
	err = json.Unmarshal(response, &result)
	if err != nil {
		return response, err
	}
	return result, nil

}

//Get 获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制，详见获取二维码。
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func (a *Wxacode) Get(path, scene string) (interface{}, error) {
	var result interface{}
	body := Acode{
		Scene: scene,
		Page:  path, //"pages/index/index",
		Width: 280,
	}
	response,err := utils.PostBody("/wxa/getwxacode", utils.JsonToByte(body) ,a.App)
	err = json.Unmarshal(response, &result)
	if err != nil {
		return response, err
	}
	return result, nil
}

//GetUnlimited 获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制。 更多用法详见 获取二维码。
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (a *Wxacode) GetUnlimited(path, scene string) (interface{}, error) {
	var result interface{}
	body := Acode{
		Scene: scene,
		Page:  path, //"pages/index/index",
		Width: 280,
	}
	response,err := utils.PostBody("/wxa/getwxacodeunlimit", utils.JsonToByte(body) ,a.App)
	err = json.Unmarshal(response, &result)
	if err != nil {
		return response, err
	}
	return result, nil
}