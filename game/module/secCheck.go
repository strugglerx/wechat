package module

import (
	"io"

	"github.com/strugglerx/wechat/v2/utils"
)

/**
* @PROJECT_NAME wechat
* @author  Moqi
* @date  2021-05-20 10:41
* @Email:str@li.cm
**/

var SecCheckEntity = SecCheck{}

type SecCheck struct {
	App utils.App
}

func (a *SecCheck) Init(app utils.App) *SecCheck {
	a.App = app
	return a
}

//ImgSecCheck 校验一张图片是否含有违法违规内容。详见内容安全解决方案
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.imgSecCheck.html
func (a *SecCheck) ImgSecCheck(file io.Reader, fileName string) (utils.Response, error) {
	response, err := utils.PostBufferFile("/wxa/img_sec_check", "media", file, fileName, a.App)
	return response, err
}

//MediaCheckAsync 异步校验图片/音频是否含有违法违规内容。
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.mediaCheckAsync.html
func (a *SecCheck) MediaCheckAsync(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxa/media_check_async", body, a.App)
	return response, err
}

//MsgSecCheck 检查一段文本是否含有违法违规内容。
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.msgSecCheck.html
func (a *SecCheck) MsgSecCheck(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxa/msg_sec_check", body, a.App)
	return response, err
}
