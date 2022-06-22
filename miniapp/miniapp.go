package miniapp

import (
	"fmt"
	"time"

	"github.com/strugglerx/wechat/v2/miniapp/module"
	"github.com/strugglerx/wechat/v2/utils"
	"github.com/tidwall/gjson"
)

type App struct {
	Appid  string
	Secret string
	Token  *utils.Token
	Verify bool        //检查Appid和Secret 是否是错的
	Write  utils.Write //读
	Read   utils.Read  //写
}

//New 新建wechat Hook 处理accesstoken的逻辑
func New(app *App) *App {
	if app.Verify {
		app.init()
	}
	return app
}

func (a *App) GetConfig() utils.Config {
	return utils.Config{
		Appid:  a.Appid,
		Secret: a.Secret,
	}
}

func (a *App) init() {
	if a.Token == nil {
		a.Token = &utils.Token{
			Token:      "",
			UpdateTime: 0,
		}
	}
	response, _ := utils.Get("/cgi-bin/token", utils.Query{
		"appid":      a.Appid,
		"secret":     a.Secret,
		"grant_type": "client_credential",
	})
	a.Token.Token = gjson.Get(string(response), "access_token").String()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if a.Token.Token == "" {
		panic("Wechat Package [" + a.Appid + "] : \n" + string(response))
	}
	// Hook Logic
	if a.Write != nil {
		a.Write(a.Appid, a.Token.Token)
	}
	a.Token.UpdateTime = int(time.Now().Unix())
}

//GetAccessTokenSafety 获取公开的accessToken
func (a *App) GetAccessTokenSafety(reflush bool) string {
	return a.GetAccessToken(reflush).Token
}

//GetAccessToken 获取accessToken 不建议暴露使用
func (a *App) GetAccessToken(reflush ...bool) *utils.Token {
	//hook
	if a.Read != nil && a.Write != nil {
		return a.GetAccessTokenWithHook(reflush...)
	}
	if a.Token == nil {
		a.Token = &utils.Token{
			Token:      "",
			UpdateTime: 0,
		}
	}
	nowTime := int(time.Now().Unix())
	doReflush := false
	if len(reflush) > 0 {
		doReflush = reflush[0]
	}
	if nowTime-a.Token.UpdateTime >= 7000 || doReflush {
		params := utils.Query{
			"appid":      a.Appid,
			"secret":     a.Secret,
			"grant_type": "client_credential",
		}
		response, _ := utils.Get("/cgi-bin/token", params)
		a.Token.Token = gjson.Get(string(response), "access_token").String()
		a.Token.UpdateTime = nowTime
		return a.Token
	} else {
		return a.Token
	}
}

//GetAccessTokenWithHook 获取accessToken 不建议暴露使用
func (a *App) GetAccessTokenWithHook(reflush ...bool) *utils.Token {
	token := a.Read(a.Appid)
	if token == nil {
		token = &utils.Token{
			Token:      "",
			UpdateTime: 0,
		}
	}
	nowTime := int(time.Now().Unix())
	doReflush := false
	if len(reflush) > 0 {
		doReflush = reflush[0]
	}
	if nowTime-token.UpdateTime >= 7000 || doReflush {
		params := utils.Query{
			"appid":      a.Appid,
			"secret":     a.Secret,
			"grant_type": "client_credential",
		}
		response, _ := utils.Get("/cgi-bin/token", params)
		token.Token = gjson.Get(string(response), "access_token").String()
		token.UpdateTime = nowTime
		a.Token = a.Write(a.Appid, token.Token)
		return token
	} else {
		return token
	}
}

//Auth 用户
func (a *App) Auth() *module.Auth {
	return module.AuthEntity.Init(a)
}

//CustomerServiceMessage 客服消息
func (a *App) CustomerServiceMessage() *module.CustomerServiceMessage {
	return module.CustomerServiceMessageEntity.Init(a)
}

//UrlScheme scheme
func (a *App) UrlScheme() *module.UrlScheme {
	return module.UrlSchemeEntity.Init(a)
}

//UrlLink url link
func (a *App) UrlLink() *module.UrlLink {
	return module.UrlLinkEntity.Init(a)
}

//Wxacode 小程序码
func (a *App) Wxacode() *module.Wxacode {
	return module.WxacodeEntity.Init(a)
}

//Soter 生物认证
func (a *App) Soter() *module.Soter {
	return module.SoterEntity.Init(a)
}

//SubscribeMessage 订阅消息
func (a *App) SubscribeMessage() *module.SubscribeMessage {
	return module.SubscribeMessageEntity.Init(a)
}

//Ocr ocr
func (a *App) Ocr() *module.Ocr {
	return module.OcrEntity.Init(a)
}

//Img 图像处理
func (a *App) Img() *module.Img {
	return module.ImgEntity.Init(a)
}

//DecodedData 解密数据
func (a *App) DecodedData() *module.DecodedData {
	return module.DecodedDataEntity.Init(a)
}

//NearbyPoi 附近的小程序
func (a *App) NearbyPoi() *module.NearbyPoi {
	return module.NearbyPoiEntity.Init(a)
}

//Operation 运维中心
func (a *App) Operation() *module.Operation {
	return module.OperationEntity.Init(a)
}

//SafetyControlCapability 安全风控
func (a *App) SafetyControlCapability() *module.SafetyControlCapability {
	return module.SafetyControlCapabilityEntity.Init(a)
}

//Search 小程序搜索
func (a *App) Search() *module.Search {
	return module.SearchEntity.Init(a)
}

//ServiceMarket 服务市场
func (a *App) ServiceMarket() *module.ServiceMarket {
	return module.ServiceMarketEntity.Init(a)
}

//Cloudbase 云开发
func (a *App) Cloudbase() *module.Cloudbase {
	return module.CloudbaseEntity.Init(a)
}

//DataAnalysis 数据分析
func (a *App) DataAnalysis() *module.DataAnalysis {
	return module.DataAnalysisEntity.Init(a)
}

//PhoneNumber 手机号
func (a *App) PhoneNumber() *module.PhoneNumber {
	return module.PhoneNumberEntity.Init(a)
}

//Custom 自定义
func (a *App) Custom() *module.Custom {
	return module.CustomEntity.Init(a)
}
