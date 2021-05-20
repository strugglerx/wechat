package miniapp

import (
	"testing"
)

func TestWxSession(t *testing.T) {
	appid := "wx18b97eec31b6db56"
	secret := "869aff2491fe005bfceb200e15679ffc"
	app := New(appid,secret)
	for i:=1;i<=2 ; i++ {
		t.Logf ("[%d] %s", i , app.GetAccessToken().Token)
	}
}

func TestApp_CreateAcode(t *testing.T) {
	appid := "wx18b97eec31b6db56"
	secret := "869aff2491fe005bfceb200e15679ffc"
	app := New(appid,secret)
	app.GetAccessToken().Token = "xxxxxx"
	app.CreateAcode("pages/index/index","1")
}

func TestWxUploadTempMedia(t *testing.T) {
	appid := "wx18b97eec31b6db56"
	secret := "869aff2491fe005bfceb200e15679ffc"
	app := &App{
		Appid: appid,
		Secret: secret,
	}
	t.Log(app.UploadTempMedia("http://10.10.10.70:8888/resource/image/example/1583205009_A1aMNEdPeZN57b6a30eb59d64a089d030f1183981213.jpg"))
}

func TestApp_OcrBusinessLicense(t *testing.T) {
	//https://image.platform.smartfacade.com.cn/tmp_4984410b35ba74caf4855b2200c862a043090648502553fb.jpg
	url := "https://image.platform.smartfacade.com.cn/tmp_4984410b35ba74caf4855b2200c862a043090648502553fb.jpg"
	appid := "wx18b97eec31b6db56"
	secret := "869aff2491fe005bfceb200e15679f7c"
	app := New(appid,secret)
	result,err :=app.OcrBusinessLicense([]byte{},url)
	t.Log(result,err)
}


