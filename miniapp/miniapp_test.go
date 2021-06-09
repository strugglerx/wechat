package miniapp

import (
	"github.com/strugglerx/wechat/utils"
	"testing"
	"time"
)

const appid = "wx18b97eec31b6db56"
const secret = "869aff2491fe005bfceb200e15679f7c"

func TestWxSession(t *testing.T) {
	app := New(appid,secret)
	for i:=1;i<=10 ; i++ {
		t.Logf ("[%d] %s", i , app.GetAccessToken().Token)
	}
}

func TestApp_CreateAcode(t *testing.T) {
	app := New(appid,secret)
	app.GetAccessToken().Token = "xxxxxx"
	app.Wxacode().GetUnlimited("pages/index/index","1")
}

func TestApp_OcrBusinessLicense(t *testing.T) {
	url := "https://image.platform.smartfacade.com.cn/tmp_4984410b35ba74caf4855b2200c862a043090648502553fb.jpg"
	app := New(appid,secret)
	result,err :=app.Ocr().BusinessLicense(url)
	t.Log(app.GetAccessToken().Token)
	t.Log(result,err)
	app.GetAccessToken().Token = "dsadasdasdsd"
	result,err =app.Ocr().BusinessLicense(url)
	t.Log(app.GetAccessToken().Token)
	t.Log(result,err)
}

func TestApp_OcrBusinessLicenseWithHook(t *testing.T) {
	var cacheToken utils.Token
	url := "https://image.platform.smartfacade.com.cn/tmp_4984410b35ba74caf4855b2200c862a043090648502553fb.jpg"
	app := New(appid,secret, func(appidAndAccessToken ...string) *utils.Token {
		if contextToken,err := utils.ExtractAppidAndAccessToken(appidAndAccessToken...);err == nil{
			// write token logic
			cacheToken.Token = contextToken.Token
			cacheToken.UpdateTime =  int(time.Now().Unix())
			return &cacheToken
		}
		//read token logic
		return &cacheToken
	})
	result,err :=app.Ocr().BusinessLicense(url)
	t.Log(app.GetAccessToken().Token)
	t.Log(result,err)
	app.GetAccessToken().Token = "dsadasdasdsd"
	result,err =app.Ocr().BusinessLicense(url)
	t.Log(app.GetAccessToken().Token)
	t.Log(result,err)
}

func TestApp_CustomerServiceMessage(t *testing.T) {
	url := "https://image.platform.smartfacade.com.cn/tmp_4984410b35ba74caf4855b2200c862a043090648502553fb.jpg"
	app := New(appid,secret)
	CustomerServiceMessage := app.CustomerServiceMessage()
	result,err := CustomerServiceMessage.UploadTempMedia(url)
	t.Log(result,err)
	//media,err := CustomerServiceMessage.GetTempMedia(result)
	//file,_ := os.Create("test.png")
	//io.Copy(file,bytes.NewReader(media.([]byte)))
	//t.Log(err)
}

func TestApp_Auth(t *testing.T) {
	code := "xxxxxxxxxxxxx"
	app := New(appid,secret)
	result,err := app.Auth().Code2Session(code)
	t.Log(result,err)
}


func TestApp_UrlScheme(t *testing.T) {
	app := New(appid,secret)
	result,err := app.UrlScheme().Generate([]byte(`{
		"jump_wxa":{
			"path":"/pages/index/index",
			"query":""
		}
	}`))
	t.Log(result,err)
}



