package utils

import (
	"encoding/json"
	"fmt"
)

/**
 * @PROJECT_NAME wechat
 * @author  Moqi
 * @date  2021-04-12 08:51
 * @Email:str@li.cm
 **/

const domain Domain = "https://api.weixin.qq.com"

var expiredToken = MapBoolean{
	"42001": true, //40001 获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口
	"40001": true, //42001 access_token 超时，请检查 access_token 的有效期，请参考基础支持 - 获取 access_token 中，对 access_token 的详细机制说明
}

type App interface {
	GetAccessToken(reflush ...bool) *Token
	GetConfig() Config
}

type Config struct {
	Appid  string
	Secret string
}

type Token struct {
	Token      string
	UpdateTime int
}

type ContextToken struct {
	Appid string
	Token string
}

type Domain string

type MapStr map[string]string

type Query map[string]string

type MapBoolean map[string]bool

type MapInterface map[string]interface{}

type Write func(appid, accessToken string) *Token

type Read func(appid string) *Token

type JsonResponse struct {
	Errcode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type Response []byte

func (r *Response) Unmarshal(p interface{}) error {
	err := json.Unmarshal(*r, &p)
	return err
}

func (r *Response) Map() (MapInterface, error) {
	var p MapInterface
	err := json.Unmarshal(*r, &p)
	return p, err
}

func (r *Response) String() string {
	return string(*r)
}

type showError struct {
	errorCode int
	errorMsg  error
}

func (e showError) Error() string {
	return fmt.Sprintf("{code: %v, error: \"%v\"}", e.errorCode, e.errorMsg)
}
