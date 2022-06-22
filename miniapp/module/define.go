package module

/**
 * @PROJECT_NAME wechat
 * @author  Moqi
 * @date  2021-05-20 14:02
 * @Email:str@li.cm
 **/

type mapInterface map[string]interface{}

type User struct {
	Session string `json:"session,omitempty"`
	Openid  string `json:"openid,omitempty"`
	Appid   string `json:"appid,omitempty"`
	Unionid string `json:"unionid,omitempty"`
}

type Session struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type Response struct {
	Errcode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type Template struct {
	Touser     string      `json:"touser,omitempty"`
	TemplateId string      `json:"template_id,omitempty"`
	Page       string      `json:"page,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

type Subscribe struct {
	Touser     string      `json:"touser,omitempty"`
	TemplateId string      `json:"template_id,omitempty"`
	Page       string      `json:"page,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

type CheckMedia struct {
	MediaUrl  string `json:"media_url"`
	MediaType int    `json:"media_type"`
}

type Text struct {
	Content string `json:"content"`
}

type SubscribeAdd struct {
	Tid       string `json:"tid,omitempty"`
	KidList   []int  `json:"kidList,omitempty"`
	SceneDesc string `json:"sceneDesc,omitempty"`
}

type TextResponse struct {
	Errcode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type SendText struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type SendImage struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Image   struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

type SendLink struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Link    struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		ThumbUrl    string `json:"thumb_url"`
	} `json:"link"`
}

type SendMini struct {
	Touser          string `json:"touser"`
	Msgtype         string `json:"msgtype"`
	Miniprogrampage struct {
		Title        string `json:"title"`
		Pagepath     string `json:"pagepath"`
		ThumbMediaId string `json:"thumb_media_id"`
	} `json:"miniprogrampage"`
}
