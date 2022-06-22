package module

import (
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
func (a *Wxacode) CreateQRCode(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/cgi-bin/wxaapp/createwxaqrcode", body, a.App)
	return response, err

}

//Get 获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制，详见获取二维码。
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func (a *Wxacode) Get(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxa/getwxacode", body, a.App)
	return response, err
}

//GetUnlimited 获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制。 更多用法详见 获取二维码。
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (a *Wxacode) GetUnlimited(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxa/getwxacodeunlimit", body, a.App)
	return response, err
}
