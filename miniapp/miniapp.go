package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/strugglerx/wechat/tool"
	"github.com/tidwall/gjson"
	"mime/multipart"
	"net/url"
	"os"
	"strings"
	"time"
)

type App struct {
	Appid  string
	Secret string
	Token  *tool.Token
}

//New 新建wechat
func New(appid, secret string) *App {
	app := &App{
		Appid:  appid,
		Secret: secret,
	}
	app.init()
	return app
}

func (a *App) init() {
	if a.Token == nil {
		a.Token = &tool.Token{
			Token:      "",
			UpdateTime: 0,
		}
	}
	responseString, _ := tool.Get("/cgi-bin/token", tool.MapStr{
		"appid":      a.Appid,
		"secret":     a.Secret,
		"grant_type": "client_credential",
	})
	a.Token.Token = gjson.Get(responseString, "access_token").String()
	if a.Token.Token == "" {
		panic("Wechat Package [" + a.Appid + "] : \n" + responseString)
	}
	a.Token.UpdateTime = int(time.Now().Unix())
}

/**
 * @PROJECT 会话状态
 * @Email:str@li.cm
 **/

//Session 获取session
func (a *App) Session(code string) (User, error) {
	var result SessResponse
	params := tool.MapStr{
		"appid":      a.Appid,
		"secret":     a.Secret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	responseString, err := tool.Get("/sns/jscode2session", params)
	if err != nil {
		return User{}, err
	}
	err = json.Unmarshal([]byte(responseString), &result)
	if err != nil {
		return User{}, err
	}
	if result.Errcode == 0 {
		result := User{Session: result.SessionKey, Openid: result.Openid, Appid: a.Appid, Unionid: result.Unionid, Status: true}
		return result, nil
	}
	return User{"", "", "", "", false}, errors.New(responseString)
}

/**
 * @PROJECT 获取accessToken Public
 * @Email:str@li.cm
 **/
func (a *App) GetAccessTokenPublic(force bool) string {
	if a.Token == nil {
		a.Token = &tool.Token{
			Token:      "",
			UpdateTime: 0,
		}
	}
	nowTime := int(time.Now().Unix())
	if nowTime-a.Token.UpdateTime >= 7000 || force {
		params := tool.MapStr{
			"appid":      a.Appid,
			"secret":     a.Secret,
			"grant_type": "client_credential",
		}
		responseString, _ := tool.Get("/cgi-bin/token", params)
		a.Token.Token = gjson.Get(responseString, "access_token").String()
		a.Token.UpdateTime = nowTime
		return a.Token.Token
	} else {
		return a.Token.Token
	}
}

/**
 * @PROJECT 获取accessToken
 * @Email:str@li.cm
 **/
func (a *App) GetAccessToken() *tool.Token {
	if a.Token == nil {
		a.Token = &tool.Token{
			Token:      "",
			UpdateTime: 0,
		}
	}
	nowTime := int(time.Now().Unix())
	if nowTime-a.Token.UpdateTime >= 7000 {
		params := tool.MapStr{
			"appid":      a.Appid,
			"secret":     a.Secret,
			"grant_type": "client_credential",
		}
		responseString, _ := tool.Get("/cgi-bin/token", params)
		a.Token.Token = gjson.Get(responseString, "access_token").String()
		a.Token.UpdateTime = nowTime
		return a.Token
	} else {
		return a.Token
	}
}

/**
 * @PROJECT 内容安全相关
 * @Email:str@li.cm
 **/

func (a *App) CheckText(body []byte) (interface{}, error) {
	var result TextResponse
	responseByte, err := tool.PostBody("/wxa/msg_sec_check", body, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *App) CheckImg(file multipart.File, fileHeader *multipart.FileHeader) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBufferFile("/wxa/img_sec_check", "media", file, fileHeader, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *App) CheckMedia(body []byte) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBody("/wxa/media_check_async", body, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

func (a *App) UploadTempMedia(fileUrlPath string) (string, error) {
	filePathsplit := strings.Split(fileUrlPath, "example")
	fileLocalPath := fmt.Sprintf("./public/resource/image/example%s", filePathsplit[1])
	fileData, err := os.Open(fileLocalPath)
	if err != nil {
		return "", err
	}
	defer fileData.Close()
	responseByte, err := tool.PostPathFile("/cgi-bin/media/upload", "media", fileData, fileUrlPath, tool.MapStr{
		"type": "image",
	}, tool.ContextApp(a))
	if err != nil {
		return "", err
	}
	return gjson.Get(string(responseByte), "media_id").String(), nil
}

/**
 * @PROJECT 客服消息发送
 * @Email:str@li.cm
 **/

func (a *App) SendTextMsg(content, openid string) int64 {
	var body SendText
	body.Touser = openid
	body.Msgtype = "text"
	body.Text.Content = content
	responseByte, _ := tool.PostBody("/cgi-bin/message/custom/send", tool.JsonToByte(body), tool.ContextApp(a))
	return gjson.Get(string(responseByte), "errcode").Int()
}

func (a *App) SendImageMsg(mediaId, openid string) int64 {
	var body SendImage
	body.Touser = openid
	body.Msgtype = "image"
	body.Image.MediaId = mediaId
	responseByte, _ := tool.PostBody("/cgi-bin/message/custom/send", tool.JsonToByte(body), tool.ContextApp(a))
	return gjson.Get(string(responseByte), "errcode").Int()
}

/**
 * 模版消息 订阅消息
 **/
func (a *App) PushSubscribe(body []byte) (interface{}, error) {
	var subscribe Subscribe
	err := json.Unmarshal(body, &subscribe)
	if err != nil {
		return nil, errors.New("unmarshal fail")
	}
	responseByte, err := tool.PostBody("/cgi-bin/message/subscribe/send", body, tool.ContextApp(a))
	var result interface{}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

/**
 * @PROJECT 订阅消息模版相关
 * @Email:str@li.cm
 **/

//推送订阅信息
//{
//	"touser": "OPENID",
//	"template_id": "TEMPLATE_ID",
//	"page": "index",
//	"miniprogram_state":"developer",
//	"lang":"zh_CN",
//	"data": {
//	"number01": {
//	"value": "339208499"
//	},
//	"date01": {
//			"value": "2015年01月05日"
//			},
//			"site01": {
//			"value": "TIT创意园"
//			} ,
//			"site02": {
//			"value": "广州市新港中路397号"
//			}
//	}
//}
func (a *App) PushScribeTemplate(body []byte) (interface{}, error) {
	var template Template
	err := json.Unmarshal(body, &template)
	if err != nil {
		return nil, errors.New("unmarshal fail")
	}
	responseByte, err := tool.PostBody("/cgi-bin/message/subscribe/send", body, tool.ContextApp(a))
	var result interface{}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

//获取当前帐号下的个人模板列表
func (a *App) GetSelfScribeTemplate() (interface{}, error) {
	var result interface{}
	responseString, err := tool.Get("/wxaapi/newtmpl/gettemplate", tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(responseString), &result)
	return result, err
}

//获取帐号所属类目下的公共模板标题
func (a *App) GetSelfScribeTemplateTitle(ids string, start, limit int) (interface{}, error) {
	var result interface{}
	responseString, err := tool.Get("/wxaapi/newtmpl/getpubtemplatetitles", tool.MapStr{
		"ids":   ids,
		"start": fmt.Sprintf("%d", start),
		"limit": fmt.Sprintf("%d", limit),
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(responseString), &result)
	return result, err
}

//获取当前账号下的个人模版关键词列表
func (a *App) GetSelfScribeTemplateKeywords(tid string) (interface{}, error) {
	var result interface{}
	responseString, err := tool.Get("/wxaapi/newtmpl/getpubtemplatekeywords", tool.MapStr{
		"tid": tid,
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(responseString), &result)
	return result, err
}

//获取当前账号类目
func (a *App) GetSelfScribeTemplateCategory() (interface{}, error) {
	var result interface{}
	responseString, err := tool.Get("/wxaapi/newtmpl/getcategory", tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(responseString), &result)
	return result, err
}

//删除账号下个人模版
func (a *App) GetSelfScribeTemplateDelete(priTmplId string) (interface{}, error) {
	var result interface{}
	body := mapInterface{
		"priTmplId": priTmplId,
	}
	responseByte, err := tool.PostBody("/wxaapi/newtmpl/deltemplate", tool.JsonToByte(body), tool.ContextApp(a))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

//增加账号个人模版
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.addTemplate.html
func (a *App) GetSelfScribeTemplateAdd(body []byte) (interface{}, error) {
	var result interface{}
	var subscribe SubscribeAdd
	err := json.Unmarshal(body, &subscribe)
	if err != nil {
		return nil, errors.New("unmarshal fail")
	}
	responseByte, err := tool.PostBody("/wxaapi/newtmpl/addtemplate", body, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

/**
 * @PROJECT 生成二维码
 * @Email:str@li.cm
 **/

//创建自定义scene的小程序码
func (a *App) CreateAcode(path, scene string) (interface{}, error) {
	var verify interface{}
	body := Acode{
		Scene: scene,
		Page:  path, //"pages/index/index",
		Width: 280,
	}
	responseByte, err := tool.PostBody("/wxa/getwxacodeunlimit", tool.JsonToByte(body), tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &verify)
	if err != nil {
		// 不为nil 为解析出图片
		return responseByte, err
	}
	//没有解析出图片
	return verify, nil
}

//永久小程序码
func (a *App) CreateForeverAcode(path string) (interface{}, error) {
	var verify interface{}
	body := Acode{
		Page:  path, //"pages/index/index",
		Width: 280,
	}
	responseByte, err := tool.PostBody("/wxa/getwxacode", tool.JsonToByte(body), tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &verify)
	if err != nil {
		// 不为nil 为解析出图片
		return responseByte, err
	}
	//没有解析出图片
	return verify, nil
}

//生成极多小程序码，暂时无限制
func (a *App) CreateMoreAcode(path, scene string) (interface{}, error) {
	var verify interface{}
	body := Acode{
		Scene: scene,
		Page:  path, //"pages/index/index",
		Width: 280,
	}
	responseByte, err := tool.PostBody("/wxa/getwxacodeunlimit", tool.JsonToByte(body), tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &verify)
	if err != nil {
		// 不为nil 为解析出图片
		return responseByte, err
	}
	//没有解析出图片
	return verify, nil
}

/**
 * @PROJECT 加密数据
 * @Email:str@li.cm
 **/

//解析30天步数等加密数据
//通过code
func (a *App) DecodeCryptoDataByCode(code, encryptedData, iv string) (interface{}, error) {
	userSession, _ := a.Session(code)
	wxCrypt := WxBizDataCrypt{
		AppId:      a.Appid,
		SessionKey: userSession.Session,
	}
	result, err := wxCrypt.Decrypt(encryptedData, iv, false)
	return result, err
}

//解析30天步数等加密数据
//通过sessionKey
func (a *App) DecodeCryptoDataBySessionKey(sessionKey, encryptedData, iv string) (interface{}, error) {
	wxCrypt := WxBizDataCrypt{
		AppId:      a.Appid,
		SessionKey: sessionKey,
	}
	result, err := wxCrypt.Decrypt(encryptedData, iv, false)
	return result, err
}

/**
 * @PROJECT 生物认证
 * @Email:str@li.cm
 **/

//生物认证
func (a *App) DecodeVerify(openId, jsonString, jsonSignature string) (interface{}, error) {
	var result interface{}
	body := mapInterface{
		"openid":         openId,
		"json_string":    jsonString,
		"json_signature": jsonSignature,
	}
	responseByte, err := tool.PostBody("/cgi-bin/soter/verify_signature", tool.JsonToByte(body), tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

/**
 * @PROJECT 小程序搜索
 * @Email:str@li.cm
 **/

//搜索小程序页面
func (a *App) SiteSearch(keyword, nextPageInfo string) (interface{}, error) {
	var result interface{}
	body := mapInterface{
		"keyword":        keyword,
		"next_page_info": nextPageInfo,
	}
	responseByte, err := tool.PostBody("/wxa/sitesearch", tool.JsonToByte(body), tool.ContextApp(a))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

//提交小程序页面
func (a *App) SiteSubmitPage(body []byte) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBody("/wxa/search/wxaapi_submitpages", tool.JsonToByte(body), tool.ContextApp(a))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

//小程序搜索图片
func (a *App) SiteImageSearch(body []byte) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBody("/wxa/imagesearch", tool.JsonToByte(body), tool.ContextApp(a))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, nil
}

/**
 * @PROJECT 图像处理
 * @Email:str@li.cm
 **/

//图片智能裁剪二进制
func (a *App) ImgAicropBinary(file multipart.File, fileHeader *multipart.FileHeader, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBufferFile("/cv/img/aicrop", "img", file, fileHeader, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//图片智能裁剪
func (a *App) ImgAicrop(body []byte, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBody("/cv/img/aicrop", body, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

//小程序的条码/二维码识别二进制
func (a *App) ImgScanQRCodeBinary(file multipart.File, fileHeader *multipart.FileHeader, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBufferFile("/cv/img/qrcode", "img", file, fileHeader, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//小程序的条码/二维码识别
func (a *App) ImgScanQRCode(body []byte, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBody("/cv/img/qrcode", body, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

//小程序的图片高清化能力二进制
func (a *App) ImgSuperresolutionBinary(file multipart.File, fileHeader *multipart.FileHeader, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBufferFile("/cv/img/superresolution", "img", file, fileHeader, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//小程序的图片高清化能力
func (a *App) ImgSuperresolution(body []byte, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBody("/cv/img/superresolution", body, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

/**
 * @PROJECT OCR相关
 * @Email:str@li.cm
 **/

//银行卡 OCR 识别binary
func (a *App) OcrBankcardBinary(file multipart.File, fileHeader *multipart.FileHeader, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBufferFile("/cv/ocr/bankcard", "img", file, fileHeader, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//银行卡 OCR 识别
func (a *App) OcrBankcard(body []byte, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBody("/cv/ocr/bankcard", body, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

//营业执照 OCR 识别binary
func (a *App) OcrBusinessLicenseBinary(file multipart.File, fileHeader *multipart.FileHeader, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBufferFile("/cv/ocr/bizlicense", "img", file, fileHeader, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//营业执照 OCR 识别
func (a *App) OcrBusinessLicense(body []byte, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBody("/cv/ocr/bizlicense", body, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
	}, tool.ContextApp(tool.ContextApp(a)))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

//驾驶证 OCR 识别binary
func (a *App) OcrDriverLicenseBinary(file multipart.File, fileHeader *multipart.FileHeader, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBufferFile("/cv/ocr/drivinglicense", "img", file, fileHeader, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//驾驶证 OCR 识别
func (a *App) OcrDriverLicense(body []byte, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBody("/cv/ocr/drivinglicense", body, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

//身份证 OCR 识别binary
func (a *App) OcrIdcardBinary(file multipart.File, fileHeader *multipart.FileHeader, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBufferFile("/cv/ocr/idcard", "img", file, fileHeader, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//身份证 OCR 识别
func (a *App) OcrIdcard(body []byte, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBody("/cv/ocr/idcard", body, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

//通用印刷体  OCR 识别binary
func (a *App) OcrPrintedTextBinary(file multipart.File, fileHeader *multipart.FileHeader, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBufferFile("/cv/ocr/comm", "img", file, fileHeader, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//通用印刷体  OCR 识别
func (a *App) OcrPrintedText(body []byte, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBody("/cv/ocr/comm", body, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}

//行驶证  OCR 识别binary
func (a *App) OcrVehicleLicenseBinary(file multipart.File, fileHeader *multipart.FileHeader, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBufferFile("/cv/ocr/driving", "img", file, fileHeader, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//行驶证  OCR 识别
func (a *App) OcrVehicleLicense(body []byte, imgUrl string) (interface{}, error) {
	var result interface{}
	responseByte, err := tool.PostBody("/cv/ocr/driving", body, tool.MapStr{
		"img_url": url.QueryEscape(imgUrl),
		"type":    "MODE",
	}, tool.ContextApp(a))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	return result, err
}
