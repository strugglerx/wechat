package miniapp

import (
	"testing"
	"time"

	"github.com/strugglerx/wechat/utils"
)

const appid = "wxb6e0730e1f5c8a1f1"
const secret = "afbc1b1c8beefd3e5e9048faad4dd0787"

func TestWxSession(t *testing.T) {
	app := New(&App{
		Appid:  appid,
		Secret: secret,
		Verify: true,
	})
	for i := 1; i <= 10; i++ {
		t.Logf("[%d] %s", i, app.GetAccessToken().Token)
	}
}

func TestApp_CreateAcode(t *testing.T) {
	app := App{
		Appid:  appid,
		Secret: secret,
	}
	app.GetAccessToken().Token = "xxxxxx"
	result, err := app.Wxacode().GetUnlimited(utils.JsonToByte(map[string]string{
		"path":        "/pages/index1/index",
		"env_version": "release",
		"scene":       "1",
	}))
	t.Log(result, err)
}

func TestApp_OcrBusinessLicense(t *testing.T) {
	url := "https://image.platform.smartfacade.com.cn/tmp_4984410b35ba74caf4855b2200c862a043090648502553fb.jpg"
	app := App{
		Appid:  appid,
		Secret: secret,
	}
	result, err := app.Ocr().BusinessLicense(url)
	t.Log(app.GetAccessToken().Token)
	t.Log(result, err)
	app.GetAccessToken().Token = "dsadasdasdsd"
	result, err = app.Ocr().BusinessLicense(url)
	t.Log(app.GetAccessToken().Token)
	t.Log(result, err)
}

func TestApp_OcrBusinessLicenseWithHook(t *testing.T) {
	var cacheToken utils.Token
	url := "https://image.platform.smartfacade.com.cn/tmp_4984410b35ba74caf4855b2200c862a043090648502553fb.jpg"
	app := &App{
		Appid:  appid,
		Secret: secret,
		Read: func(appid string) *utils.Token {
			return &cacheToken
		},
		Write: func(appid, accessToken string) *utils.Token {
			cacheToken.Token = accessToken
			cacheToken.UpdateTime = int(time.Now().Unix())
			return &cacheToken
		},
	}
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
	app := App{
		Appid:  appid,
		Secret: secret,
	}
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
	app := App{
		Appid:  appid,
		Secret: secret,
	}
	result, err := app.Auth().Code2Session(code)
	t.Log(result, err)
}

func TestApp_UrlScheme(t *testing.T) {
	app := App{
		Appid:  appid,
		Secret: secret,
	}
	result, err := app.UrlScheme().Generate([]byte(`{
		"jump_wxa":{
			"path":"/pages/logs/logs",
			"query":""
		}
	}`))
	t.Log(err)
	resultMap, err := result.Map()
	t.Logf("%+v", resultMap["openlink"])
}

func TestApp_UrlLink(t *testing.T) {
	app := App{
		Appid:  appid,
		Secret: secret,
	}
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
	app := App{
		Appid:  appid,
		Secret: secret,
	}
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
	app := App{
		Appid:  appid,
		Secret: secret,
	}
	result, err := app.PhoneNumber().GetPhoneNumber([]byte(`{
		"code":"03c52dedef3306d529d53bb31452ec9a2f46880b2040cec9d760876e821f9429"
	}`))
	t.Log(result, err)
}
