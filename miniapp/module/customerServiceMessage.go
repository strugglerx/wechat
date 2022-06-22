package module

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/strugglerx/wechat/utils"
	"github.com/tidwall/gjson"
)

/**
 * @PROJECT_NAME wechat
 * @author  Moqi
 * @date  2021-05-20 10:41
 * @Email:str@li.cm
 **/

var CustomerServiceMessageEntity = CustomerServiceMessage{}

type CustomerServiceMessage struct {
	App utils.App
	Dir string //素材服务器路径
}

func (a *CustomerServiceMessage) Init(app utils.App) *CustomerServiceMessage {
	a.App = app
	return a
}

//GetTempMedia 获取客服消息内的临时素材。即下载临时的多媒体文件。目前小程序仅支持下载图片文件。
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.getTempMedia.html
func (a *CustomerServiceMessage) GetTempMedia(mediaId string) (utils.Response, error) {
	params := utils.Query{
		"media_id": mediaId,
	}
	return utils.Get("/cgi-bin/media/get", params, a.App)
}

//SetTyping 下发客服当前输入状态给用户。详见 客服消息输入状态
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.setTyping.html
func (a *CustomerServiceMessage) SetTyping(openid, command string) (utils.Response, error) {
	params := utils.Query{
		"touser":  openid,
		"command": command,
	}
	response, err := utils.PostBody("/cgi-bin/message/custom/typing", utils.JsonToByte(params), a.App)
	return response, err
}

//UploadTempMedia 把媒体文件上传到微信服务器。目前仅支持图片。用于发送客服消息或被动回复用户消息。
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.uploadTempMedia.html
func (a *CustomerServiceMessage) UploadTempMedia(uri string) (string, error) { //网络资源
	if strings.HasPrefix(uri, "http") {
		response, err := utils.PostBufferFileWithField("/cgi-bin/media/upload", "media", bytes.NewReader(utils.FetchSource(uri)), uri, utils.Query{"type": "image"}, a.App)
		if gjson.Get(string(response), "errcode").Int() != 0 {
			return "", errors.New(string(response))
		}
		return gjson.Get(string(response), "media_id").String(), err
	}
	if a.Dir == "" {
		return "", errors.New("Must Set Dir")
	}
	filePath := fmt.Sprintf("%s/%s", a.Dir, uri)
	fileData, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer fileData.Close()
	responseByte, err := utils.PostPathFile("/cgi-bin/media/upload", "media", fileData, filePath, utils.Query{
		"type": "image",
	}, a.App)
	if err != nil {
		return "", err
	}
	return gjson.Get(string(responseByte), "media_id").String(), nil
}

//Send 发送客服消息给用户。详细规则见 发送客服消息
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html
func (a *CustomerServiceMessage) Send(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/cgi-bin/message/custom/send", body, a.App)
	return response, err
}
