package module

import (
	"github.com/strugglerx/wechat/utils"
)

/**
* @PROJECT_NAME wechat
* @author  Moqi
* @date  2021-05-20 10:41
* @Email:str@li.cm
**/

var PhoneNumberEntity = PhoneNumber{}

type PhoneNumber struct {
	App utils.App
}

func (a *PhoneNumber) Init(app utils.App) *PhoneNumber {
	a.App = app
	return a
}

//GetPhoneNumber
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/phonenumber/phonenumber.getPhoneNumber.html
func (a *PhoneNumber) GetPhoneNumber(body []byte) (utils.Response, error) {
	response, err := utils.PostBody("/wxa/business/getuserphonenumber", body, a.App)
	return response, err
}
