package module

import (
	"io"
	"net/url"

	"github.com/strugglerx/wechat/v2/utils"
)

/**
* @PROJECT_NAME wechat
* @author  Moqi
* @date  2021-05-20 10:41
* @Email:str@li.cm
**/

var OcrEntity = Ocr{}

type Ocr struct {
	App utils.App
}

func (a *Ocr) Init(app utils.App) *Ocr {
	a.App = app
	return a
}

//Bankcard 本接口提供基于小程序的银行卡 OCR 识别
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.bankcard.html
func (a *Ocr) Bankcard(imgUrl string) (utils.Response, error) {
	response, err := utils.PostBody("/cv/ocr/bankcard", []byte{}, utils.Query{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, a.App)
	return response, err
}

//BusinessLicense 本接口提供基于小程序的营业执照 OCR 识别
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.businessLicense.html
func (a *Ocr) BusinessLicense(imgUrl string) (utils.Response, error) {
	response, err := utils.PostBody("/cv/ocr/bizlicense", []byte{}, utils.Query{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, a.App)
	return response, err
}

//DriverLicense 本接口提供基于小程序的驾驶证 OCR 识别
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.driverLicense.html
func (a *Ocr) DriverLicense(imgUrl string) (utils.Response, error) {
	response, err := utils.PostBody("/cv/ocr/drivinglicense", []byte{}, utils.Query{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, a.App)
	return response, err
}

//Idcard 本接口提供基于小程序的身份证 OCR 识别
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.idcard.html
func (a *Ocr) Idcard(imgUrl string) (utils.Response, error) {
	response, err := utils.PostBody("/cv/ocr/idcard", []byte{}, utils.Query{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, a.App)
	return response, err
}

//PrintedText 本接口提供基于小程序的通用印刷体 OCR 识别
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.printedText.html
func (a *Ocr) PrintedText(imgUrl string) (utils.Response, error) {
	response, err := utils.PostBody("/cv/ocr/comm", []byte{}, utils.Query{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, a.App)
	return response, err
}

//VehicleLicense 本接口提供基于小程序的行驶证 OCR 识别
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.vehicleLicense.html
func (a *Ocr) VehicleLicense(imgUrl string) (utils.Response, error) {
	response, err := utils.PostBody("/cv/ocr/driving", []byte{}, utils.Query{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, a.App)
	return response, err
}

//BankcardBuffer 本接口提供基于小程序的银行卡 OCR 识别
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.bankcard.html
func (a *Ocr) BankcardBuffer(file io.Reader, fileName string, imgUrl string) (utils.Response, error) {
	response, err := utils.PostBufferFile("/cv/ocr/bankcard", "img", file, fileName, utils.Query{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, a.App)
	return response, err
}

//BusinessLicenseBuffer 本接口提供基于小程序的营业执照 OCR 识别
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.businessLicense.html
func (a *Ocr) BusinessLicenseBuffer(file io.Reader, fileName string, imgUrl string) (utils.Response, error) {
	response, err := utils.PostBufferFile("/cv/ocr/bizlicense", "img", file, fileName, utils.Query{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, a.App)
	return response, err
}

//DriverLicenseBuffer 本接口提供基于小程序的驾驶证 OCR 识别
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.driverLicense.html
func (a *Ocr) DriverLicenseBuffer(file io.Reader, fileName string, imgUrl string) (utils.Response, error) {
	response, err := utils.PostBufferFile("/cv/ocr/drivinglicense", "img", file, fileName, utils.Query{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, a.App)
	return response, err
}

//IdcardBuffer 本接口提供基于小程序的身份证 OCR 识别
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.idcard.html
func (a *Ocr) IdcardBuffer(file io.Reader, fileName string, imgUrl string) (utils.Response, error) {
	response, err := utils.PostBufferFile("/cv/ocr/idcard", "img", file, fileName, utils.Query{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, a.App)
	return response, err
}

//PrintedTextBuffer 本接口提供基于小程序的通用印刷体 OCR 识别
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.printedText.html
func (a *Ocr) PrintedTextBuffer(file io.Reader, fileName string, imgUrl string) (utils.Response, error) {
	response, err := utils.PostBufferFile("/cv/ocr/comm", "img", file, fileName, utils.Query{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, a.App)
	return response, err
}

//VehicleLicenseBuffer 本接口提供基于小程序的行驶证 OCR 识别
//http://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.vehicleLicense.html
func (a *Ocr) VehicleLicenseBuffer(file io.Reader, fileName string, imgUrl string) (utils.Response, error) {
	response, err := utils.PostBufferFile("/cv/ocr/driving", "img", file, fileName, utils.Query{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, a.App)
	return response, err
}
