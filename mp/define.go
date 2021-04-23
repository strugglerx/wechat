package wechatmp

import (
	"encoding/xml"
	"github.com/strugglerx/wechat/tool"
)

/**
 * @PROJECT_NAME wechatmp
 * @author  Moqi
 * @date  2020-04-09 20:10
 * @Email:str@li.cm
 **/

// ====== 域名等固定参数 ======


//tickerDomain
const tickerDomain tool.Domain = "https://mp.weixin.qq.com"

// 解析query
type param map[string]string

// ====== 结构体 ======

type User struct {
	Session string `json:"session,omitempty"`
	Openid  string `json:"openid,omitempty"`
	Appid   string `json:"appid,omitempty"`
	Unionid string `json:"unionid,omitempty"`
	Status  bool   `json:"status,omitempty"`
}

type SessionResponse struct {
	Openid      string
	SessionKey string
	Unionid     string
	Errcode     int
	ErrMsg      string
}

type Template struct {
	Touser      string      `json:"touser,omitempty"`
	TemplateId string      `json:"template_id,omitempty"`
	Page        string      `json:"page,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

type Subscribe struct {
	Touser      string      `json:"touser,omitempty"`
	TemplateId string      `json:"template_id,omitempty"`
	Page        string      `json:"page,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

type OauthToken struct {
	Token        string
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
	UpdateTime   int
}

type Text struct {
	Content string `json:"content"`
}

type TextResponse struct {
	Errcode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type ShortUrl struct {
	Action  string `json:"action"`
	LongUrl string `json:"long_url"`
}

type Code struct {
	ExpireSeconds int    `json:"expire_seconds,omitempty"`
	ActionName    string `json:"action_name"`
	ActionInfo    struct {
		Scene struct {
			SceneID  int    `json:"scene_id,omitempty"`
			SceneStr string `json:"scene_str,omitempty"`
		} `json:"scene"`
	} `json:"action_info"`
}

type CodeResponse struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	URL           string `json:"url"`
}

type Remark struct {
	Openid string `json:"openid"`
	Remark string `json:"remark"`
}

type UserListResponse struct {
	Total int `json:"total"`
	Count int `json:"count"`
	Data struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`
}

type Comment struct {
	MsgDataID int `json:"msg_data_id"`
	Index     int `json:"index"`
	Begin     int `json:"begin"`
	Count     int `json:"count"`
	Type      int `json:"type"`
}

type SendJsonText struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type SendJsonImage struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Image   struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

type SendJsonVoice struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Voice   struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

type SendJsonVideo struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Video   struct {
		MediaId string `json:"media_id"`
		ThumbMediaId string `json:"thumb_media_id"`
		Title string `json:"title"`
		Description string `json:"description"`
	} `json:"video"`
}

type SendMusic struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Music   struct {
		Musicurl string `json:"musicurl"`
		Hqmusicurl string `json:"hqmusicurl"`
		ThumbMediaId string `json:"thumb_media_id"`
		Title string `json:"title"`
		Description string `json:"description"`
	} `json:"music"`
}

type SendNews struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	News   struct {
		Articles []struct{
			Url string `json:"url"`
			Picurl string `json:"picurl"`
			Title string `json:"title"`
			Description string `json:"description"`
		} `json:"articles"`
	} `json:"news"`
}


type SendMpNews struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Mpnews   struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

//客服状态
type SendCommand struct {
	Touser  string `json:"touser"`
	Command string `json:"command"`
}

type SendMsgMenu struct {
	Touser string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Msgmenu struct {
		HeadContent string `json:"head_content"`
		List []struct {
			ID string `json:"id"`
			Content string `json:"content"`
		} `json:"list"`
		TailContent string `json:"tail_content"`
	} `json:"msgmenu"`
}


type SendMini struct {
	Touser          string `json:"touser"`
	Msgtype         string `json:"msgtype"`
	Miniprogrampage struct {
		Title        string `json:"title"`
		Appid        string `json:"appid"`
		Pagepath     string `json:"pagepath"`
		ThumbMediaId string `json:"thumb_media_id"`
	} `json:"miniprogrampage"`
}

type News struct {
	Articles []struct {
		ThumbMediaID string `json:"thumb_media_id"`
		Author string `json:"author"`
		Title string `json:"title"`
		ContentSourceURL string `json:"content_source_url"`
		Content string `json:"content"`
		Digest string `json:"digest"`
		ShowCoverPic int `json:"show_cover_pic"`
		NeedOpenComment int `json:"need_open_comment"`
		OnlyFansCanComment int `json:"only_fans_can_comment"`
	} `json:"articles"`
}

type TagPushNews struct {
	Filter struct {
		IsToAll bool `json:"is_to_all"`
		TagID int `json:"tag_id"`
	} `json:"filter"`
	Mpnews struct {
		MediaID string `json:"media_id"`
	} `json:"mpnews"`
	Msgtype string `json:"msgtype"`
	SendIgnoreReprint int `json:"send_ignore_reprint"`
}

type TagPushText struct {
	Filter struct {
		IsToAll bool `json:"is_to_all"`
		TagID int `json:"tag_id"`
	} `json:"filter"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Msgtype string `json:"msgtype"`
}

type TagPushVoice struct {
	Filter struct {
		IsToAll bool `json:"is_to_all"`
		TagID int `json:"tag_id"`
	} `json:"filter"`
	Voice struct {
		MediaID string `json:"media_id"`
	} `json:"voice"`
	Msgtype string `json:"msgtype"`
}

type TagPushImage struct {
	Filter struct {
		IsToAll bool `json:"is_to_all"`
		TagID int `json:"tag_id"`
	} `json:"filter"`
	Images struct {
		MediaIds []string `json:"media_ids"`
		Recommend string `json:"recommend"`
		NeedOpenComment int `json:"need_open_comment"`
		OnlyFansCanComment int `json:"only_fans_can_comment"`
	} `json:"images"`
	Msgtype string `json:"msgtype"`
}

type TagPushVideo struct {
	MediaID string `json:"media_id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

// 服务号结构体
type ServerSendNews struct {
	Touser []string `json:"touser"`
	Mpnews struct {
		MediaID string `json:"media_id"`
	} `json:"mpnews"`
	Msgtype string `json:"msgtype"`
	SendIgnoreReprint int `json:"send_ignore_reprint"`
}

type ServerSendText struct {
	Touser []string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

type ServerSendVoice struct {
	Touser []string `json:"touser"`
	Voice struct {
		MediaID string `json:"media_id"`
	} `json:"voice"`
	Msgtype string `json:"msgtype"`
}

type ServerSendImage struct {
	Touser []string `json:"touser"`
	Images struct {
		MediaIds []string `json:"media_ids"`
		Recommend string `json:"recommend"`
		NeedOpenComment int `json:"need_open_comment"`
		OnlyFansCanComment int `json:"only_fans_can_comment"`
	} `json:"images"`
	Msgtype string `json:"msgtype"`
}

type DelPush struct {
	MsgID int `json:"msg_id"`
	ArticleIdx int `json:"article_idx"`
}

// ====== 公众号接受XML ======

//公众号消息相关
type CDATA struct {
	Text string `xml:",cdata"`
}

//返回信息的结构体
/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[text]]></MsgType>
  <Content><![CDATA[你好]]></Content>
</xml>
*/
type SendXmlText struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Content      CDATA    `xml:"Content"`
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[news]]></MsgType>
  <ArticleCount>1</ArticleCount>
  <Articles>
    <item>
      <Title><![CDATA[title1]]></Title>
      <Description><![CDATA[description1]]></Description>
      <PicUrl><![CDATA[picurl]]></PicUrl>
      <Url><![CDATA[url]]></Url>
    </item>
  </Articles>
</xml>
*/
type SendArticle struct {
	XMLName      xml.Name                  `xml:"xml"`
	ToUserName   CDATA                     `xml:"ToUserName"`
	FromUserName CDATA                     `xml:"FromUserName"`
	CreateTime   int                       `xml:"CreateTime"`
	MsgType      CDATA                     `xml:"MsgType"`
	ArticleCount int                       `xml:"ArticleCount"`
	Articles     []SendArticleItem `xml:"Articles>item"`
}

type SendArticleItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       CDATA    `xml:"Title"`
	Description CDATA    `xml:"Description"`
	PicUrl      CDATA    `xml:"PicUrl"`
	Url         CDATA    `xml:"Url"`
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[image]]></MsgType>
  <Image>
    <MediaId><![CDATA[media_id]]></MediaId>
  </Image>
</xml>
*/
type SendXmlImage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Image        struct {
		MediaId CDATA `xml:"MediaId"`
	} `xml:"Image"`
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[voice]]></MsgType>
  <Voice>
    <MediaId><![CDATA[media_id]]></MediaId>
  </Voice>
</xml>
*/
type SendXmlVoice struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Voice        struct {
		MediaId CDATA `xml:"MediaId"`
	} `xml:"Voice"`
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[video]]></MsgType>
  <Video>
    <MediaId><![CDATA[media_id]]></MediaId>
    <Title><![CDATA[title]]></Title>
    <Description><![CDATA[description]]></Description>
  </Video>
</xml>
*/
type SendXmlVideo struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Video        struct {
		MediaId     CDATA `xml:"MediaId"`
		Title       CDATA `xml:"Title"`
		Description CDATA `xml:"Description"`
	} `xml:"Video"`
}
