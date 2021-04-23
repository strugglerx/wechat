package wechat

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

func TestWxUploadTempMedia(t *testing.T) {
	appid := "wx18b97eec31b6db56"
	secret := "869aff2491fe005bfceb200e15679ffc"
	app := &App{
		Appid: appid,
		Secret: secret,
	}
	t.Log(app.UploadTempMedia("http://10.10.10.70:8888/resource/image/example/1583205009_A1aMNEdPeZN57b6a30eb59d64a089d030f1183981213.jpg"))
}


