package mp

/**
 * @PROJECT_NAME wechatmp
 * @author  Moqi
 * @date  2020-04-09 19:58
 * @Email:str@li.cm
 **/

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/strugglerx/wechat/tool"
	"github.com/tidwall/gjson"
	"mime/multipart"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)


//var H5Token_ *Token = &Token{"", 0}
//
//var H5Oauth2Token_ *OauthToken = &OauthToken{Token: "", UpdateTime: 0}

// ====== 初始化实例 ======

type Mp struct {
	Appid string
	Secret string
	Token *tool.Token
	Oauth2Token *OauthToken
}

/**
 * @param
 * @author struggler
 * @description 初始化wechatmp
 * @date 10:36 上午 2021/2/24
 * @return
 **/
func New(appid,secret string) *Mp {
	mp := &Mp{
		Appid: appid,
		Secret: secret,
	}
	//mp.init()
	return mp
}

func (m *Mp) init() {
	if m.Token == nil {
		m.Token = &tool.Token{
			Token:      "",
			UpdateTime: 0,
		}
	}
	if m.Oauth2Token == nil {
		m.Oauth2Token = &OauthToken{
			Token: "", UpdateTime: 0,
		}
	}
	params := tool.MapStr{
		"appid":      m.Appid,
		"secret":     m.Secret,
		"grant_type": "client_credential",
	}
	responseString, _ := tool.Get("/cgi-bin/token",params)
	m.Token.Token = gjson.Get(responseString, "access_token").String()
	if m.Token.Token == "" {
		panic("WechatMp Package [" + m.Appid + "] : \n" + responseString)
	}
	m.Token.UpdateTime = int(time.Now().Unix())
}

/**
 * @PROJECT 服务端接口
 * @Email:str@li.cm
 **/

//获取session
func (m *Mp) Session(code string) (User, error) {
	var res SessionResponse
	params := tool.MapStr{
		"appid":      m.Appid,
		"secret":     m.Secret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	responseString, _ := tool.Get("/sns/jscode2session",params)
	err := json.Unmarshal([]byte(responseString), &res)
	if err != nil {
		return User{},err
	}
	if res.Errcode == 0 {
		result := User{Session: res.SessionKey, Openid: res.Openid, Appid: m.Appid, Unionid: res.Unionid, Status: true}
		return result, nil
	}
	return User{"", "", "", "", false}, errors.New(responseString)
}

//获取accessToken
func (m *Mp) GetAccessToken() *tool.Token {
	if m.Token == nil {
		m.Token = &tool.Token{
			Token:      "",
			UpdateTime: 0,
		}
	}
	nowTime := int(time.Now().Unix())
	if nowTime - m.Token.UpdateTime >= 3600 {
		params := tool.MapStr{
			"appid":      m.Appid,
			"secret":     m.Secret,
			"grant_type": "client_credential",
		}
		responseString, _ := tool.Get("/cgi-bin/token",params)
		m.Token.Token = gjson.Get(responseString, "access_token").String()
		m.Token.UpdateTime = nowTime
		return m.Token
	} else {
		return m.Token
	}
}

//获取微信服务器ip
func (m *Mp) CallbackIp() (interface{}, error) {
	//封装json post请求
	responseString, err := tool.Get("/cgi-bin/getcallbackip",tool.ContextApp(m))
	if err != nil {
		return responseString, err
	}
	return gjson.Get(responseString, "ip_list").String(), nil
}

//设置菜单
func (m *Mp) SetMenu(body []byte) (interface{}, error) {
	var result TextResponse
	responseByte,err := tool.PostBody("/cgi-bin/menu/create",body,tool.ContextApp(m))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//获取普通菜单
func (m *Mp) GetSelfMenu() (interface{}, error) {
	var result interface{}
	responseString, err := tool.Get("/cgi-bin/get_current_selfmenu_info",tool.ContextApp(m))
	if err != nil {
		return result, err
	}
	return responseString, nil
}

//获取全部菜单(包括个性菜单)
func (m *Mp) GetMenu() (interface{}, error) {
	var result interface{}
	responseString, err := tool.Get("/cgi-bin/menu/get",tool.ContextApp(m))
	if err != nil {
		return result, err
	}
	return responseString, nil
}

//删除菜单
func (m *Mp) RemoveMenu() (interface{}, error) {
	var result interface{}
	//封装json post请求
	responseString, err := tool.Get("/cgi-bin/menu/delete",tool.ContextApp(m))
	if err != nil {
		return result, err
	}
	return responseString, nil
}

//设置个性化菜单
func (m *Mp) SetStyleMenu(body []byte) (interface{}, error) {
	var result TextResponse
	responseByte,err := tool.PostBody("/cgi-bin/menu/addconditional",body,tool.ContextApp(m))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//删除个性化菜单
func (m *Mp) DelStyleMenu(body []byte) (interface{}, error) {
	var result TextResponse
	responseByte,err := tool.PostBody("/cgi-bin/menu/delconditional",body,tool.ContextApp(m))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//测试个性化菜单
func (m *Mp) TestStyleMenu(body []byte) (interface{}, error) {
	var result TextResponse
	responseByte,err := tool.PostBody("/cgi-bin/menu/trymatch",body,tool.ContextApp(m))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

/**
 * @PROJECT 开放能力
 * @Email:str@li.cm
 **/

//长链接转短链接
//w.url.cn
func (m *Mp) GetShortUrl(longUrl string) (string, error) {
	bodyStruct := ShortUrl{
		Action:  "long2short",
		LongUrl: longUrl,
	}
	bodyByte, _ := json.Marshal(bodyStruct)
	responseByte,err := tool.PostBody("/cgi-bin/shorturl",bodyByte,tool.ContextApp(m))
	if err != nil {
		return "", err
	}
	if gjson.Get(string(responseByte), "errcode").Int() != 0 {
		return "", errors.New("request fail")
	}
	return gjson.Get(string(responseByte), "short_url").String(), nil
}

/**
 * @PROJECT 获取公众号自动回复规则
 * @Email:str@li.cm
 **/

//获取回复规则
func (m *Mp) GetReplyRules() (interface{}, error) {
	var result interface{}
	//封装json post请求
	responseString, err := tool.Get("/cgi-bin/get_current_autoreply_info",tool.ContextApp(m))
	if err != nil {
		return result, err
	}
	return responseString, nil
}

/**
 * @PROJECT 客服消息
 * @Email:str@li.cm
 **/

//发送文本消息
func (m *Mp) SendTextMsg(content, openid string) int64 {
	var body SendJsonText
	body.Touser = openid
	body.Msgtype = "text"
	body.Text.Content = content
	responseByte,err := tool.PostBody("/cgi-bin/message/custom/send",tool.JsonToByte(body),tool.ContextApp(m))
	if err != nil {
		return  1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

//发送图片
func (m *Mp) SendImageMsg(mediaId, openid string) int64 {
	var body SendJsonImage
	body.Touser = openid
	body.Msgtype = "image"
	body.Image.MediaId = mediaId
	responseByte,err := tool.PostBody("/cgi-bin/message/custom/send",tool.JsonToByte(body),tool.ContextApp(m))
	if err != nil {
		return  1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

//发送声音
func (m *Mp) SendVoiceMsg(mediaId, openid string) int64 {
	var body SendJsonVoice
	body.Touser = openid
	body.Msgtype = "voice"
	body.Voice.MediaId = mediaId
	responseByte,err := tool.PostBody("/cgi-bin/message/custom/send",tool.JsonToByte(body),tool.ContextApp(m))
	if err != nil {
		return  1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

//发送视频
func (m *Mp) SendVideoMsg(mediaId, openid, thumbMediaId, title, description string) int64 {
	var body SendJsonVideo
	body.Touser = openid
	body.Msgtype = "video"
	body.Video.MediaId = mediaId
	body.Video.ThumbMediaId = thumbMediaId
	body.Video.Title = title
	body.Video.Description = description
	responseByte,err := tool.PostBody("/cgi-bin/message/custom/send",tool.JsonToByte(body),tool.ContextApp(m))
	if err != nil {
		return  1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

//发送音乐
func (m *Mp) SendMusicMsg(openid, thumbMediaId, musicurl, hqmusicurl, title, description string) int64 {
	var body SendMusic
	body.Touser = openid
	body.Msgtype = "music"
	body.Music.Musicurl = musicurl
	body.Music.Hqmusicurl = hqmusicurl
	body.Music.ThumbMediaId = thumbMediaId
	body.Music.Title = title
	body.Music.Description = description
	responseByte,err := tool.PostBody("/cgi-bin/message/custom/send",tool.JsonToByte(body),tool.ContextApp(m))
	if err != nil {
		return  1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

//发送外链图文信息
func (m *Mp) SendNewsMsg(openid, title, description, url_, picurl string) int64 {
	var body SendNews
	body.Touser = openid
	body.Msgtype = "news"
	body.News.Articles = []struct {
		Url         string `json:"url"`
		Picurl      string `json:"picurl"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}{
		{
			Title:       title,
			Description: description,
			Url:         url_,
			Picurl:      picurl,
		},
	}
	responseByte,err := tool.PostBody("/cgi-bin/message/custom/send",tool.JsonToByte(body),tool.ContextApp(m))
	if err != nil {
		return  1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

//发送素材信息
func (m *Mp) SendMpnewsMsg(openid, mediaid string) int64 {
	var body SendMpNews
	body.Touser = openid
	body.Msgtype = "mpnews"
	body.Mpnews.MediaId = mediaid
	responseByte,err := tool.PostBody("/cgi-bin/message/custom/send",tool.JsonToByte(body),tool.ContextApp(m))
	if err != nil {
		return  1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

//客服发送状态
// command[ Typing CancelTyping ]
func (m *Mp) SendCommandMsg(openid, command string) int64 {
	var body SendCommand
	body.Touser = openid
	body.Command = command
	responseByte,err := tool.PostBody("/cgi-bin/message/custom/send",tool.JsonToByte(body),tool.ContextApp(m))
	if err != nil {
		return  1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

//发送小程序
func (m *Mp) SendMiniProgramMsg(openid, thumbMediaId, pagepath, appid, title string) int64 {
	var body SendMini
	body.Touser = openid
	body.Msgtype = "miniprogrampage"
	body.Miniprogrampage.Title = title
	body.Miniprogrampage.Appid = appid
	body.Miniprogrampage.Pagepath = pagepath
	body.Miniprogrampage.ThumbMediaId = thumbMediaId
	responseByte,err := tool.PostBody("/cgi-bin/message/custom/send",tool.JsonToByte(body),tool.ContextApp(m))
	if err != nil {
		return  1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

/**
 * @PROJECT 素材管理
 * @Email:str@li.cm
 **/

//上传临时素材Binary
//type :媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb）
func (m *Mp) UploadTempMediaBinary(file multipart.File, fileHeader *multipart.FileHeader, type_ string) (interface{}, error) {
	var result interface{}
	responseByte,err := tool.PostBufferFile("/cgi-bin/media/upload","media",file,fileHeader,tool.MapStr{
		"type":type_,
	},tool.ContextApp(m))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//上传临时素材
//type :媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb）
func (m *Mp) UploadTempMedia(fileUrlPath, type_ string) (string, error) {
	filePathsplit := strings.Split(fileUrlPath, "example")
	fileLocalPath := fmt.Sprintf("./public/resource/image/example%s", filePathsplit[1])
	fileData, err := os.Open(fileLocalPath)
	if err != nil {
		return "", err
	}
	defer fileData.Close()
	responseByte,err := tool.PostPathFile("/cgi-bin/media/upload","media",fileData,fileUrlPath,tool.MapStr{
		"type":type_,
	},tool.ContextApp(m))
	if err != nil {
		return "", err
	}
	return gjson.Get(string(responseByte), "media_id").String(), nil
}

//获取临时素材
//type :媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb）
func (m *Mp) GetTempMedia(mediaId string) ([]byte, error) {
	responseString, err := tool.Get("/cgi-bin/media/get",tool.MapStr{
		"media_id":mediaId,
	},tool.ContextApp(m))
	if err != nil {
		return []byte{}, err
	}
	return []byte(responseString), nil
}

//新增永久素材
func (m *Mp) SetForeverNews(body []byte) (interface{}, error) {
	var result TextResponse
	responseByte,err := tool.PostBody("/cgi-bin/material/add_news",body,tool.ContextApp(m))
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//上传永久素材
//type :媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb）
func (m *Mp) UploadForeverMedia(fileUrlPath, type_ string) (string, error) {
	filePathsplit := strings.Split(fileUrlPath, "example")
	fileLocalPath := fmt.Sprintf("./public/resource/image/example%s", filePathsplit[1])
	fileData, err := os.Open(fileLocalPath)
	if err != nil {
		return "", err
	}
	defer fileData.Close()
	responseByte,err := tool.PostPathFile("/cgi-bin/material/add_material","media",fileData,fileUrlPath,tool.MapStr{
		"type":type_,
	},tool.ContextApp(m))
	if err != nil {
		return "", err
	}
	return gjson.Get(string(responseByte), "media_id").String(), nil
}

//获取永久素材
//type :媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb）
func (m *Mp) GetForeverMedia(mediaId string) ([]byte, error) {
	responseString, err := tool.Get("/cgi-bin/material/get_material",tool.MapStr{
		"media_id":mediaId,
	},tool.ContextApp(m))
	if err != nil {
		return []byte{}, err
	}
	return []byte(responseString), nil
}

//删除永久素材
//{
//	"media_id" : "media id"
//}
func (m *Mp) DelForeverMedia(body []byte) (interface{}, error) {
	var result TextResponse
	responseByte,err := tool.PostBody("/cgi-bin/material/del_material",body,tool.ContextApp(m))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//获取永久素材总数
func (m *Mp) GetForeverMediaCount() ([]byte, error) {
	responseString, err := tool.Get("/cgi-bin/material/get_materialcount",tool.ContextApp(m))
	if err != nil {
		return []byte{}, err
	}
	return []byte(responseString), nil
}

//获取永久素材列表
//{
//	"type":TYPE,
//	"offset":OFFSET,
//	"count":COUNT
//}
func (m *Mp) GetForeverMediaList(body []byte) (interface{}, error) {
	var result interface{}
	responseByte,err := tool.PostBody("/cgi-bin/material/batchget_material",body,tool.ContextApp(m))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//上传图片
func (m *Mp) UploadImg(file multipart.File, fileHeader *multipart.FileHeader) (interface{}, error) {
	var result interface{}
	responseByte,err := tool.PostBufferFile("/cgi-bin/media/uploadimg","media",file,fileHeader,tool.ContextApp(m))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/**
 * @PROJECT 群发接口
 * @Email:str@li.cm
 **/

//上传图文素材
/*{
	"articles": [
	{
		"thumb_media_id":"qI6_Ze_6PtV7svjolgs-rN6stStuHIjs9_DidOHaj0Q-mwvBelOXCFZiq2OsIU-p",
		"author":"xxx",
		"title":"Happy Day",
		"content_source_url":"www.qq.com",
		"content":"content",
		"digest":"digest",
		"show_cover_pic":1,
		"need_open_comment":1,
		"only_fans_can_comment":1
	},
	{
		"thumb_media_id":"qI6_Ze_6PtV7svjolgs-rN6stStuHIjs9_DidOHaj0Q-mwvBelOXCFZiq2OsIU-p",
		"author":"xxx",
		"title":"Happy Day",
		"content_source_url":"www.qq.com",
		"content":"content",
		"digest":"digest",
		"show_cover_pic":0,
		"need_open_comment":1,
		"only_fans_can_comment":1
	}
]
}*/
func (m *Mp) CreateNews(news News) int64 {
	responseByte,err := tool.PostBody("/cgi-bin/media/uploadnews",tool.JsonToByte(news),tool.ContextApp(m))
	if err != nil {
		return 1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

/**
 * @PROJECT 服务号 专用通过openid群发消息
 * @Email:str@li.cm
 **/

func (m *Mp) PushServerNews(news ServerSendNews) int64 {
	responseByte,err := tool.PostBody("/cgi-bin/message/mass/send",tool.JsonToByte(news),tool.ContextApp(m))
	if err != nil {
		return 1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

func (m *Mp) PushServerText(news ServerSendText) int64 {
	responseByte,err := tool.PostBody("/cgi-bin/message/mass/send",tool.JsonToByte(news),tool.ContextApp(m))
	if err != nil {
		return 1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

func (m *Mp) PushServerVoice(news ServerSendVoice) int64 {
	responseByte,err := tool.PostBody("/cgi-bin/message/mass/send",tool.JsonToByte(news),tool.ContextApp(m))
	if err != nil {
		return 1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

func (m *Mp) PushServerImage(news ServerSendImage) int64 {
	responseByte,err := tool.PostBody("/cgi-bin/message/mass/send",tool.JsonToByte(news),tool.ContextApp(m))
	if err != nil {
		return 1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

//删除群发
func (m *Mp) PushDel(msgId, articleIdx int) int64 {
	var body DelPush
	body.MsgID = msgId
	body.ArticleIdx = articleIdx
	responseByte,err := tool.PostBody("/cgi-bin/message/mass/delete",tool.JsonToByte(body),tool.ContextApp(m))
	if err != nil {
		return 1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

/**
 * @PROJECT 认证号预览接口 通过openid群发消息 单日限制100次
 * @Email:str@li.cm
 **/

//发送文本消息
func (m *Mp) SendPreviewTextMsg(content, openid string) int64 {
	var body SendJsonText
	body.Touser = openid
	body.Msgtype = "text"
	body.Text.Content = content
	responseByte,err := tool.PostBody("/cgi-bin/message/mass/preview",tool.JsonToByte(body),tool.ContextApp(m))
	if err != nil {
		return 1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

//发送图片
func (m *Mp) SendPreviewImageMsg(mediaId, openid string) int64 {
	var body SendJsonImage
	body.Touser = openid
	body.Msgtype = "image"
	body.Image.MediaId = mediaId
	responseByte,err := tool.PostBody("/cgi-bin/message/mass/preview",tool.JsonToByte(body),tool.ContextApp(m))
	if err != nil {
		return 1
	}
	return gjson.Get(string(responseByte), "errcode").Int()
}

/**
 * @PROJECT 开放能力
 * @Email:str@li.cm
 **/

//内容安全检查
func (m *Mp) CheckText(body []byte) (interface{}, error) {
	var result TextResponse
	responseByte,err := tool.PostBody("/a/msg_sec_check",tool.JsonToByte(body),tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

func (m *Mp) CheckImg(file multipart.File, fileHeader *multipart.FileHeader) (interface{}, error) {
	var result interface{}
	responseByte,err := tool.PostBufferFile("/a/img_sec_check","media",file,fileHeader,tool.ContextApp(m))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/**
 * @PROJECT 模版消息
 * @Email:str@li.cm
 **/

//设置行业信息
//{
//"industry_id1":"1",
//"industry_id2":"4"
//}
func (m *Mp) SetIndustry(body []byte) (interface{}, error) {
	var result interface{}
	responseByte,err := tool.PostBody("/cgi-bin/example/api_set_industry",body,tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

//获取设置的行业信息
func (m *Mp) GetIndustry() (interface{}, error) {
	var result interface{}
	responseString, err := tool.Get("/cgi-bin/example/get_industry",tool.ContextApp(m))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(responseString), &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//获得模板ID
//{
//"template_id_short":"TM00015"
//}
func (m *Mp) GetTemplateId(body []byte) (interface{}, error) {
	var result TextResponse
	responseByte,err := tool.PostBody("/cgi-bin/example/api_add_template",body,tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

//获取模板列表
func (m *Mp) GetTemplateList() (interface{}, error) {
	var result interface{}
	responseString, err := tool.Get("/cgi-bin/example/get_all_private_template",tool.ContextApp(m))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(responseString), &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//删除模版
//{
//	"template_id" : "Dyvp3-Ff0cnail_CDSzk1fIc6-9lOkxsQE7exTJbwUE"
//}
func (m *Mp) DelTemplateId(body []byte) (interface{}, error) {
	var result TextResponse
	responseByte,err := tool.PostBody("/cgi-bin/example/del_private_template",body,tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

//发送模版信息
/*
 {
           "touser":"OPENID",
           "template_id":"ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
           "url":"http://weixin.qq.com/download",
           "miniprogram":{
             "appid":"xiaochengxuappid12345",
             "pagepath":"index?foo=bar"
           },
           "data":{
                   "first": {
                       "value":"恭喜你购买成功！",
                       "color":"#173177"
                   },
                   "keyword1":{
                       "value":"巧克力",
                       "color":"#173177"
                   },
                   "keyword2": {
                       "value":"39.8元",
                       "color":"#173177"
                   },
                   "keyword3": {
                       "value":"2014年9月22日",
                       "color":"#173177"
                   },
                   "remark":{
                       "value":"欢迎再次购买！",
                       "color":"#173177"
                   }
           }
       }
*/
func (m *Mp) PushTemplate(body []byte) (interface{}, error) {
	var result interface{}
	responseByte,err := tool.PostBody("/cgi-bin/message/example/send",body,tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

/**
 * @PROJECT 一次性订阅
 * @Email:str@li.cm
 **/

//发送订阅消息
/*
 {
    "touser":"OPENID",
    "template_id":"TEMPLATE_ID",
    "url":"URL",
    "miniprogram":{
    "appid":"xiaochengxuappid12345",
    "pagepath":"index?foo=bar"
},
    "scene":"SCENE",
    "title":"TITLE",
    "data":{
    "content":{
    "value":"VALUE",
    "color":"COLOR"
}
}
}
*/
func (m *Mp) PushSubscribeTemplate(body []byte) (interface{}, error) {
	var result interface{}
	responseByte,err := tool.PostBody("/cgi-bin/message/example/subscribe",body,tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

/**
 * @PROJECT 创建二维码
 * @Email:str@li.cm
 **/

//创建永久二维码 QR_LIMIT_SCENE QR_LIMIT_STR_SCENE
func (m *Mp) CreateForeverAcode(actionName, scene string) (interface{}, error) {
	var result CodeResponse
	var params Code
	if actionName == "QR_LIMIT_SCENE" {
		sceneId, _ := strconv.Atoi(scene)
		params = Code{
			ActionName: "QR_SCENE",
			ActionInfo: struct {
				Scene struct {
					SceneID  int    `json:"scene_id,omitempty"`
					SceneStr string `json:"scene_str,omitempty"`
				} `json:"scene"`
			}{
				Scene: struct {
					SceneID  int    `json:"scene_id,omitempty"`
					SceneStr string `json:"scene_str,omitempty"`
				}{
					SceneID: sceneId,
				},
			},
		}
	} else {
		params = Code{
			ActionName: "QR_LIMIT_STR_SCENE",
			ActionInfo: struct {
				Scene struct {
					SceneID  int    `json:"scene_id,omitempty"`
					SceneStr string `json:"scene_str,omitempty"`
				} `json:"scene"`
			}{
				Scene: struct {
					SceneID  int    `json:"scene_id,omitempty"`
					SceneStr string `json:"scene_str,omitempty"`
				}{
					SceneStr: scene,
				},
			},
		}

	}
	responseByte,err := tool.PostBody("/cgi-bin/qrcode/create",tool.JsonToByte(params),tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil|| result.Ticket == "" {
		return result, err
	}
	return m.GetTicketAcode(result.Ticket)
}

//创建临时二维码  QR_SCENE QR_STR_SCENE
func (m *Mp) CreateMoreAcode(actionName, scene string) (interface{}, error) {
	var result CodeResponse
	var params Code
	if actionName == "QR_SCENE" {
		sceneId, _ := strconv.Atoi(scene)
		params = Code{
			ExpireSeconds: 2592000,
			ActionName:    "QR_SCENE",
			ActionInfo: struct {
				Scene struct {
					SceneID  int    `json:"scene_id,omitempty"`
					SceneStr string `json:"scene_str,omitempty"`
				} `json:"scene"`
			}{
				Scene: struct {
					SceneID  int    `json:"scene_id,omitempty"`
					SceneStr string `json:"scene_str,omitempty"`
				}{
					SceneID: sceneId,
				},
			},
		}
	} else {
		params = Code{
			ExpireSeconds: 2592000,
			ActionName:    "QR_STR_SCENE",
			ActionInfo: struct {
				Scene struct {
					SceneID  int    `json:"scene_id,omitempty"`
					SceneStr string `json:"scene_str,omitempty"`
				} `json:"scene"`
			}{
				Scene: struct {
					SceneID  int    `json:"scene_id,omitempty"`
					SceneStr string `json:"scene_str,omitempty"`
				}{
					SceneStr: scene,
				},
			},
		}
	}
	responseByte,err := tool.PostBody("/cgi-bin/qrcode/create",tool.JsonToByte(params),tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil|| result.Ticket == "" {
		return result, err
	}
	return m.GetTicketAcode(result.Ticket)
}

//通过ticket换取真的二维码
func (m *Mp) GetTicketAcode(ticket string) (interface{}, error) {
	var result interface{}
	enTicket := url.QueryEscape(ticket)
	responseString, err := tool.Get("/cgi-bin/showqrcode",tool.MapStr{
		"ticket":enTicket,
	},tickerDomain)
	if err != nil {
		return result, err
	}
	//这里可以判断是创建出了图片
	//gfile.PutBytes("cover.png",[]byte(responseString))
	return []byte(responseString), nil
}

/**
 * @PROJECT 用户管理
 * @Email:str@li.cm
 **/

//用户标签管理
//增加tag
//{   "tag" : {     "name" : "广东"//标签名   } }
func (m *Mp) InsertTags(body []byte) (interface{}, error) {
	var result interface{}
	responseByte,err := tool.PostBody("/cgi-bin/tags/create",body,tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

//获取所有tag
func (m *Mp) GetTags() (interface{}, error) {
	var result interface{}
	responseString, err := tool.Get("/cgi-bin/tags/get",tool.ContextApp(m))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(responseString), &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//修改tag
//{   "tag" : {     "id" : 134,     "name" : "广东人"   } }
func (m *Mp) UpdateTags(body []byte) (interface{}, error) {
	var result interface{}
	responseByte,err := tool.PostBody("/cgi-bin/tags/update",body,tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

//删除tag
//{   "tag":{        "id" : 134   } }
func (m *Mp) RemoveTags(body []byte) (interface{}, error) {
	var result interface{}
	responseByte,err := tool.PostBody("/cgi-bin/tags/delete",body,tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

//获取标签下粉丝列表
//{   "tagid" : 134,   "next_openid":""//第一个拉取的OPENID，不填默认从头开始拉取 }
func (m *Mp) GetTagsUser(body []byte) (UserListResponse, error) {
	var result UserListResponse
	responseByte,err := tool.PostBody("/cgi-bin/user/tag/get",body,tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

//批量为用户打标签
//{
//    "openid_list" : [//粉丝列表
//    "ocYxcuAEy30bX0NXmGn4ypqx3tI0",
//    "ocYxcuBt0mRugKZ7tGAHPnUaOW7Y"   ],
//    "tagid" : 134
// }
func (m *Mp) SetTagsUser(body []byte) (interface{}, error) {
	var result interface{}
	responseByte,err := tool.PostBody("/cgi-bin/tags/members/batchtagging",body,tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

//批量为用户取消标签
//{
//"openid_list" : [//粉丝列表
//"ocYxcuAEy30bX0NXmGn4ypqx3tI0",
//"ocYxcuBt0mRugKZ7tGAHPnUaOW7Y"   ],
//"tagid" : 134
//}
func (m *Mp) CancelTagsUser(body []byte) (interface{}, error) {
	var result interface{}
	responseByte,err := tool.PostBody("/cgi-bin/tags/members/batchuntagging",body,tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

//获取用户身上的标签
//{   "openid" : "ocYxcuBt0mRugKZ7tGAHPnUaOW7Y" }
func (m *Mp) GetTagsOneUser(body []byte) (interface{}, error) {
	var result interface{}
	responseByte,err := tool.PostBody("/cgi-bin/tags/getidlist",body,tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

//为用户备注
func (m *Mp) SetUserRemark(openId, remark string) (interface{}, error) {
	var result interface{}
	bodyStruct := Remark{Openid: openId, Remark: remark}
	responseByte,err := tool.PostBody("/cgi-bin/users/info/updateremark",tool.JsonToByte(bodyStruct),tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

//获取用户列表
func (m *Mp) GetUserList(nextOpenId string) (UserListResponse, error) {
	var result UserListResponse
	responseString, err := tool.Get("/cgi-bin/user/get",tool.MapStr{
		"next_openid":nextOpenId,
	},tool.ContextApp(m))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(responseString), &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//获取用户详细信息(unionid)
func (m *Mp) GetUserDetail(openId string) (string, error) {
	responseString, err := tool.Get("/cgi-bin/user/info", tool.MapStr{
		"openid":       openId,
		"lang":         "zh_CN",
	},tool.ContextApp(m))
	if err != nil {
		return "", err
	}
	return responseString, nil
}

//批量获取用户详细信息
//{
//	"user_list": [
//		{
//		"openid": "otvxTs4dckWG7imySrJd6jSi0CWE",
//		"lang": "zh_CN"
//		},
//		{
//		"openid": "otvxTs_JZ6SEiP0imdhpi50fuSZg",
//		"lang": "zh_CN"
//		}
//	]
//}
func (m *Mp) GetUsersDetail(body []byte) (interface{}, error) {
	var result interface{}
	responseByte,err := tool.PostBody("/cgi-bin/user/info/batchget",body,tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}

/**
 * @PROJECT 图文消息留言管理
 * @Email:str@li.cm
 **/

//获取文章评论
func (m *Mp) GetCommentList(msgDataId, index, begin, count, type_ int) (interface{}, error) {
	var result interface{}
	bodyStruct := Comment{
		MsgDataID: msgDataId,
		Index:     index,
		Begin:     begin,
		Count:     count,
		Type:      type_,
	}
	responseByte,err := tool.PostBody("/cgi-bin/user/info/batchget",tool.JsonToByte(bodyStruct),tool.ContextApp(m))
	if err != nil {
		return result,err
	}
	err = json.Unmarshal(responseByte, &result)
	if err != nil {
		return result, err
	}
	return result,nil
}
