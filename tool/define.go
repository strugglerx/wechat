package tool

/**
 * @PROJECT_NAME wechat
 * @author  Moqi
 * @date  2021-04-12 08:51
 * @Email:str@li.cm
 **/

const domain Domain = "https://api.weixin.qq.com"
const expiredToken = "42001,40001"

type App interface {
	GetAccessToken() (*Token)
}

type Token struct {
	Token      string
	UpdateTime int
}

type Domain string

type MapStr map[string]string

type MapInterface map[string]interface{}
