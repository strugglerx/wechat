package game

import (
	"time"

	"github.com/strugglerx/wechat/game/module"
	"github.com/strugglerx/wechat/utils"
	"github.com/tidwall/gjson"
)

/**
 * @PROJECT_NAME wechat
 * @author  Moqi
 * @date  2021-06-10 18:41
 * @Email:str@li.cm
 **/

type App struct {
	Appid  string
	Secret string
	Token  *utils.Token
	Hook   utils.Hook
}

//New 新建wechat Hook 处理accesstoken的逻辑
func New(appid, secret string, Hook ...utils.Hook) *App {
	app := &App{
		Appid:  appid,
		Secret: secret,
	}
	if len(Hook) != 0 {
		app.Hook = Hook[0]
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
		panic("WechatGame Package [" + a.Appid + "] : \n" + string(response))
	}
	// Hook Logic
	if a.Hook != nil {
		a.Hook(a.Appid, a.Token.Token)
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
	if a.Hook != nil {
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

//GetAccessTokenWithHook 获取accessToken 不建议暴露使用
func (a *App) GetAccessTokenWithHook(reflush ...bool) *utils.Token {
	token := a.Hook()
	if token == nil {
		token = &utils.Token{
			Token:      "",
			UpdateTime: 0,
		}
	}
	nowTime := int(time.Now().Unix())
	doReflush := false
	if len(reflush) > 0 {
		doReflush = true
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
		a.Token = a.Hook(a.Appid, token.Token)
		return token
	} else {
		return token
	}
}

//Login 用户
func (a *App) Login() *module.Login {
	return module.LoginEntity.Init(a)
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

//UpdatableMessage 动态消息
func (a *App) UpdatableMessage() *module.UpdatableMessage {
	return module.UpdatableMessageEntity.Init(a)
}

//SubscribeMessage 订阅消息
func (a *App) SubscribeMessage() *module.SubscribeMessage {
	return module.SubscribeMessageEntity.Init(a)
}

//SafetyControlCapability 安全风控
func (a *App) SafetyControlCapability() *module.SafetyControlCapability {
	return module.SafetyControlCapabilityEntity.Init(a)
}

//Cloudbase 云开发
func (a *App) Cloudbase() *module.Cloudbase {
	return module.CloudbaseEntity.Init(a)
}

//LockStep 帧同步
func (a *App) LockStep() *module.LockStep {
	return module.LockStepEntity.Init(a)
}

//DataAnalysis 数据分析
func (a *App) DataAnalysis() *module.DataAnalysis {
	return module.DataAnalysisEntity.Init(a)
}

//Storage 开放数据
func (a *App) Storage() *module.Storage {
	return module.StorageEntity.Init(a)
}

//Gamematch 对局匹配
func (a *App) Gamematch() *module.Gamematch {
	return module.GamematchEntity.Init(a)
}

//Get
func (a *App) Get(path string,params utils.Query,withAccessToken ...bool) ([]byte,error)  {
	if len(withAccessToken)>0{
		response, err := utils.Get(path, params,utils.ContextApp(a))
		return response,err
	}
	response, err := utils.Get(path, params)
	return response,err
}

//Post
func (a *App) PostBody(path string,body []byte ,withAccessToken ...bool) ([]byte,error)  {
	if len(withAccessToken)>0{
		response, err := utils.PostBody(path, body,utils.ContextApp(a))
		return response,err
	}
	response, err := utils.PostBody(path, body)
	return response,err
}

