package module

import (
	"encoding/json"
	"github.com/strugglerx/wechat/utils"
	"io"
)

/**
 * @PROJECT_NAME wechat
 * @author  Moqi
 * @date  2021-05-20 10:41
 * @Email:str@li.cm
 **/

var SearchEntity = Search{}

type Search struct {
	App utils.App
}

func (a *Search) Init(app utils.App) *Search {
	a.App = app
	return a
}

//ImageSearch 本接口提供基于小程序的站内搜商品图片搜索能力
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.imageSearch.html
func (a *Search) ImageSearch(file io.Reader,fileName string) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBufferFile("/wxa/imagesearch","img",file,fileName,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	return result , err
}

//SiteSearch 小程序内部搜索API提供针对页面的查询能力，小程序开发者输入搜索词后，将返回自身小程序和搜索词相关的页面。因此，利用该接口，开发者可以查看指定内容的页面被微信平台的收录情况；同时，该接口也可供开发者在小程序内应用，给小程序用户提供搜索能力。
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.siteSearch.html
func (a *Search) SiteSearch(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/sitesearch", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	return result , err
}

//SubmitPages 小程序开发者可以通过本接口提交小程序页面url及参数信息(不要推送webview页面)，让微信可以更及时的收录到小程序的页面信息，开发者提交的页面信息将可能被用于小程序搜索结果展示。
//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.submitPages.html
func (a *Search) SubmitPages(body []byte) (interface{}, error) {
	var result interface{}
	response, err := utils.PostBody("/wxa/search/wxaapi_submitpages", body,a.App)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	return result , err
}
