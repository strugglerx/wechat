package module

import (
	"encoding/json"
	"github.com/strugglerx/wechat/utils"
	"io"
)

/**
* @PROJECT_NAME wechat
* @author  Moqi
* @date  2021-05-20 10:41
* @Email:str@li.cm
**/

var ImgEntity = Img{}

type Img struct {
	App utils.App
}

func (a *Img) Init(app utils.App) *Img {
	a.App = app
	return a
}

//AiCrop 本接口提供基于小程序的图片智能裁剪能力。
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.aiCrop.html
func (a *Img) AiCrop(img_url string) (interface{}, error) {
	var result interface{}
	param := utils.Query{
		"img_url":img_url,
	}
	response, err := utils.PostBody("/cv/img/aicrop",[]byte{}, param,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return response, err
	}
	return result , nil
}

//AiCropBuffer 本接口提供基于小程序的图片智能裁剪能力。
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.aiCrop.html
func (a *Img) AiCropBuffer(file io.Reader,fileName string) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBufferFile("/cv/img/aicrop", "img",file,fileName,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return response, err
	}
	return result , nil
}

//ScanQRCode 本接口提供基于小程序的条码/二维码识别的API。
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.scanQRCode.html
func (a *Img) ScanQRCode(img_url string) (interface{}, error) {
	var result interface{}
	param := utils.Query{
		"img_url":img_url,
	}
	response, err := utils.PostBody("/cv/img/qrcode",[]byte{}, param,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return response, err
	}
	return result , nil
}

//ScanQRCode 本接口提供基于小程序的条码/二维码识别的API。
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.scanQRCode.html
func (a *Img) ScanQRCodeBuffer(file io.Reader,fileName string) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBufferFile("/cv/img/qrcode","img",file,fileName,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return response, err
	}
	return result , nil
}

//Superresolution 本接口提供基于小程序的图片高清化能力。
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.superresolution.html
func (a *Img) Superresolution(img_url string) (interface{}, error) {
	var result interface{}
	param := utils.Query{
		"img_url":img_url,
	}
	response, err := utils.PostBody("/cv/img/superresolution",[]byte{}, param,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return response, err
	}
	return result , nil
}

//Superresolution 本接口提供基于小程序的图片高清化能力。
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.superresolution.html
func (a *Img) SuperresolutionBuffer(file io.Reader,fileName string) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBufferFile("/cv/img/superresolution","img",file,fileName,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return response, err
	}
	return result , nil
}
