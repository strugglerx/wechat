package miniapp

import (
	"testing"
	"time"

	"github.com/strugglerx/wechat/utils"
)

const appid = "wx044f45b5fd8fb87c"
const secret = "ab99034aed77b4a43b1738babf2040ce"

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
			"path":"/pages/index/index",
			"query":""
		}
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
