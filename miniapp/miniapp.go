package miniapp

import (
	"github.com/strugglerx/wechat/miniapp/module"
	"github.com/strugglerx/wechat/utils"
	"github.com/tidwall/gjson"
	"time"
)

type App struct {
	Appid  string
	Secret string
	Token  *utils.Token
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
	if a.Token.Token == "" {
		panic("Wechat Package [" + a.Appid + "] : \n" + string(response))
	}
	a.Token.UpdateTime = int(time.Now().Unix())
}

//GetAccessTokenSafety 获取公开的accessToken
func (a *App) GetAccessTokenSafety(reflush bool) string {
	return  a.GetAccessToken(reflush).Token
}

//GetAccessToken 获取accessToken 不建议暴露使用
func (a *App) GetAccessToken(reflush ...bool) *utils.Token {
	if a.Token == nil {
		a.Token = &utils.Token{
			Token:      "",
			UpdateTime: 0,
		}
	}
	nowTime := int(time.Now().Unix())
	doReflush := false
	if len(reflush)>0 {
		doReflush = true
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

//Auth 用户
func (a *App) Auth() *module.Auth {
	return  module.AuthEntity.Init(a)
}

//CustomerServiceMessage 客服消息
func (a *App) CustomerServiceMessage() *module.CustomerServiceMessage {
	return  module.CustomerServiceMessageEntity.Init(a)
}

//UrlScheme scheme
func (a *App) UrlScheme() *module.UrlScheme {
	return  module.UrlSchemeEntity.Init(a)
}

//UrlLink url link
func (a *App) UrlLink() *module.UrlLink {
	return  module.UrlLinkEntity.Init(a)
}

//Wxacode 小程序码
func (a *App) Wxacode() *module.Wxacode {
	return  module.WxacodeEntity.Init(a)
}

//Soter 生物认证
func (a *App) Soter() *module.Soter {
	return  module.SoterEntity.Init(a)
}

//SubscribeMessage 订阅消息
func (a *App) SubscribeMessage() *module.SubscribeMessage {
	return  module.SubscribeMessageEntity.Init(a)
}

//Ocr ocr
func (a *App) Ocr() *module.Ocr {
	return  module.OcrEntity.Init(a)
}

//Img 图像处理
func (a *App) Img() *module.Img {
	return  module.ImgEntity.Init(a)
}

//DecodedData 解密数据
func (a *App) DecodedData() *module.DecodedData {
	return  module.DecodedDataEntity.Init(a)
}

//NearbyPoi 附近的小程序
func (a *App) NearbyPoi() *module.NearbyPoi {
	return  module.NearbyPoiEntity.Init(a)
}

//Operation 运维中心
func (a *App) Operation() *module.Operation {
	return  module.OperationEntity.Init(a)
}

//SafetyControlCapability 安全风控
func (a *App) SafetyControlCapability() *module.SafetyControlCapability {
	return  module.SafetyControlCapabilityEntity.Init(a)
}

//Search 小程序搜索
func (a *App) Search() *module.Search {
	return  module.SearchEntity.Init(a)
}

//ServiceMarket 服务市场
func (a *App) ServiceMarket() *module.ServiceMarket {
	return  module.ServiceMarketEntity.Init(a)
}

//Cloudbase 云开发
func (a *App) Cloudbase() *module.Cloudbase {
	return module.CloudbaseEntity.Init(a)
}