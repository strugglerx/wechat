package miniapp

import (
	"regexp"
	"testing"
	"time"

	"github.com/strugglerx/wechat/utils"
)

const appid = "wxb6e0730e1f5c8a98"
const secret = "afbc1b1c8beefd3e5e9048faad4dd078"

func TestWxSession(t *testing.T) {
	app := New(appid, secret)
	for i := 1; i <= 10; i++ {
		t.Logf("[%d] %s", i, app.GetAccessToken().Token)
	}
}

func TestApp_CreateAcode(t *testing.T) {
	app := New(appid, secret)
	app.GetAccessToken().Token = "xxxxxx"
	app.Wxacode().GetUnlimited("pages/index/index", "1")
}

func TestApp_OcrBusinessLicense(t *testing.T) {
	url := "https://image.platform.smartfacade.com.cn/tmp_4984410b35ba74caf4855b2200c862a043090648502553fb.jpg"
	app := New(appid, secret)
	result, err := app.Ocr().BusinessLicense(url)
	t.Log(app.GetAccessToken().Token)
	t.Log(result, err)
	app.GetAccessToken().Token = "dsadasdasdsd"
	result, err = app.Ocr().BusinessLicense(url)
	t.Log(app.GetAccessToken().Token)
	t.Log(result, err)
}

func TestApp_Ocr(t *testing.T) {
	//url := "https://image.platform.smartfacade.com.cn/tmp_f9992847c6952413266f4c81071e2ef610135abd61d354af.jpg"
	//app := New(appid, secret)
	//result, err := app.Ocr().PrintedText(url)
	//t.Log(app.GetAccessToken().Token)
	//
	//te, err := json.Marshal(result)
	//type OcrResponse struct {
	//	Errcode int    `json:"errcode"`
	//	Errmsg  string `json:"errmsg"`
	//	ImgSize struct {
	//		H int `json:"h"`
	//		W int `json:"w"`
	//	} `json:"img_size"`
	//	Items []struct {
	//		Pos struct {
	//			LeftBottom struct {
	//				X int `json:"x"`
	//				Y int `json:"y"`
	//			} `json:"left_bottom"`
	//			LeftTop struct {
	//				X int `json:"x"`
	//				Y int `json:"y"`
	//			} `json:"left_top"`
	//			RightBottom struct {
	//				X int `json:"x"`
	//				Y int `json:"y"`
	//			} `json:"right_bottom"`
	//			RightTop struct {
	//				X int `json:"x"`
	//				Y int `json:"y"`
	//			} `json:"right_top"`
	//		} `json:"pos"`
	//		Text string `json:"text"`
	//	} `json:"items"`
	//}
	//var r OcrResponse
	//err = json.Unmarshal(te, &r)
	//t.Log(err)
	//text := ""
	//if r.Errcode == 0 {
	//	for _, v := range r.Items {
	//		text += v.Text
	//	}
	//}
	//t.Log(text)
	text := "港澳居民来往内地通行证蒋元红JIANG,YUANHONG出生日期性刷女1982.04.20有效期限2021.12.13-2031.12.12签发机关中华人民共和国出入境管理局证件号码换证次数01H10481010THIS CARD IS INTENDED FOR ITS HOLDER TO TRAVEL TO THE MAINLAND OF CHINA"
	rgx, _ := regexp.Compile(`内地通行证([^\w]+)[A-Z]+.*?(H[\d]{8})`)
	list := rgx.FindStringSubmatch(text)
	t.Log(list)
	//text := "中国联通14:3350%/核酸检测记录·●核酸检测记录心刷新检测完成检测中采样时间2022-02-2321:05检测时间2022-02-2408:00检测机构深圳海普洛斯医学检验实验室数据来源广东省卫生健康委员会检测结果阴性采样时间2022-02-2121:56检测时间2022-02-2200:04外省核酸查不到?点这里>>"
	//text := "14:28<核酸检测记录核酸检测记录O刷新检测完成检测中杨曼妮采样时间2022-02-24 16:25检测时间2022-02-24 23:50检测机构珠海横琴铂华医学检验实验室数据来源广东省卫生健康委员会检测结果阴性杨曼妮采样时间2022-02-23 18:57检测时间2022-02-24 15:34检测机构深圳海普洛斯医学检验实验室数据来源广东省卫生健康委员会检测结果阴性外省核酸查不到?点这里>>"
	//(采样时间|检测时间)
	//rgx, _ := regexp.Compile(`(采样时间|检测时间)(\d{4}-\d{2}-\d{2}[ ]?\d{2}:\d{2})`)
	//list := rgx.FindAllStringSubmatch(text, -1)
	//result := map[string]string{}
	//timergx, _ := regexp.Compile(`(\d{4}-\d{2}-\d{2}\d{2}:\d{2})`)
	//def := map[string]string{
	//	"采样时间": "collect_time",
	//	"检测时间": "testing_time",
	//}
	//for _, v := range list[0:2] {
	//	t.Log(v)
	//	time := v[2]
	//	if timergx.MatchString(time) {
	//		time = time[0:10] + " " + time[10:]
	//	}
	//	result[def[v[1]]] = time
	//}
	//t.Log(result)
}

func TestApp_OcrBusinessLicenseWithHook(t *testing.T) {
	var cacheToken utils.Token
	url := "https://image.platform.smartfacade.com.cn/tmp_4984410b35ba74caf4855b2200c862a043090648502553fb.jpg"
	app := New(appid, secret, func(appidAndAccessToken ...string) *utils.Token {
		if contextToken, err := utils.ExtractAppidAndAccessToken(appidAndAccessToken...); err == nil {
			// write token logic
			cacheToken.Token = contextToken.Token
			cacheToken.UpdateTime = int(time.Now().Unix())
			return &cacheToken
		}
		//read token logic
		return &cacheToken
	})
	result, err := app.Ocr().BusinessLicense(url)
	t.Log(app.GetAccessToken().Token)
	t.Log(result, err)
	app.GetAccessToken().Token = "dsadasdasdsd"
	result, err = app.Ocr().BusinessLicense(url)
	t.Log(app.GetAccessToken().Token)
	t.Log(result, err)
}

func TestApp_CustomerServiceMessage(t *testing.T) {
	url := "https://image.platform.smartfacade.com.cn/tmp_4984410b35ba74caf4855b2200c862a043090648502553fb.jpg"
	app := New(appid, secret)
	CustomerServiceMessage := app.CustomerServiceMessage()
	result, err := CustomerServiceMessage.UploadTempMedia(url)
	t.Log(result, err)
	//media,err := CustomerServiceMessage.GetTempMedia(result)
	//file,_ := os.Create("test.png")
	//io.Copy(file,bytes.NewReader(media.([]byte)))
	//t.Log(err)
}

func TestApp_Auth(t *testing.T) {
	code := "xxxxxxxxxxxxx"
	app := New(appid, secret)
	result, err := app.Auth().Code2Session(code)
	t.Log(result, err)
}

func TestApp_UrlScheme(t *testing.T) {
	app := New(appid, secret)
	result, err := app.UrlScheme().Generate([]byte(`{
		"jump_wxa":{
			"path":"/pages/logs/logs",
			"query":""
		}
	}`))
	t.Log(result, err)
}

func TestApp_UrlLink(t *testing.T) {
	app := New(appid, secret)
	result, err := app.UrlLink().Generate([]byte(`{
		"path":"/package/pages/epidemic_report/index",
		"query":"street_id=2",
		"is_expire":true,
		"expire_type":1,
		"expire_interval":179
	}`))
	t.Log(result, err)
}

func TestApp_DecodedData(t *testing.T) {
	app := New(appid, secret)
	sessionKey := "6CbG7oWDQqcZ9nQLxsygUg=="
	encryptedData := "4VcFzKSTX4ete1V10o/6DQhCLLy3J2G5KQTxhitIg9SOp3StA/YOpdDm5yBAz7ECzNgOAhF7saSpg/U90SFlFAdwmBnjngMFiy711YWYAUcyyFKe43LGWInqMh8UIerRVyaakYYbzCl6NqbNypsLgqH14XvoUAed9YHJ1LtX4JbfCmcQ/4jQczBUq+uDtpXvAwu0g5ZEkCIYAWb8+aNlTeOzYTgK5ZjGu8wVk1JzhBHM/XA3CfEPMecjPWqwpN9uwEpGGYRikx7vikj2xfu6TPlZCDQ/3xRsAWW0CUvpZAWlIT9pmRMvJL6qxOoisPgoWRfQ9J6lLrkYCk1v/ImevFnp9/PO59qpedLnHAtPzwr/G1kUd7HQGZRbj8BtkUBmRRofzzm1oe2nbK4KPx96Lg=="
	iv := "D8nMjb2umX+ee+yH0Tt16w=="
	t.Log(app.DecodedData().DecodeCryptoData(sessionKey, encryptedData, iv))

	if true {
		sessionKey := "NM5LYVuxjbyFexm38w1YBA=="
		encryptedData := "u0p3mhZx6SJ0Bm0Hn09JCro3GEmEUyTc6avmBaNnRa4bVETFmWL1vSZEyWAGw5HsTSJRLycD+1nNyz5v6IB0CzPu4wBApeBdx4JqR/ApuqW1sSRGHl0eGYxbHSwTjY9lj7BQA9JfoGD1WqpAiPX+UaQRDz4LBO8ocZU9RWWnB+d1+koBzb1v9fdx+D2id2d6RSdbxHNuZOBzER9cikZtpQt1meRn2e0UakZZ2BPuKwi/DDsZnckibvxVGVDgdefGvI1rxqsgNyU1CkQAFprO35Uaji9fz41t6vH95chGsd5cjU0XUtyMJt7kblpUst0iPPGJUhPdnuo2b3AtDZ4W19lPrn5OVBKp+qsdKDgUshKqqfGGh/ROgYBnRIQi+WxC/qb+V4Rv0mSf5SjiMYWhkQ=="
		iv := "ucPPsWmmrZ1wAtiA7L3xow=="
		t.Log(app.DecodedData().DecodeCryptoData(sessionKey, encryptedData, iv))
	}
}

func TestApp_PhoneNumber(t *testing.T) {
	app := New(appid, secret)
	result, err := app.PhoneNumber().GetPhoneNumber([]byte(`{
		"code":"03c52dedef3306d529d53bb31452ec9a2f46880b2040cec9d760876e821f9429"
	}`))
	t.Log(result, err)
}
