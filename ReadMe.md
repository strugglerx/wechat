# Go Wechat

[![Go Doc](https://godoc.org/github.com/strugglerx/wechat?status.svg)](https://godoc.org/github.com/strugglerx/wechat)
[![Production Ready](https://img.shields.io/badge/production-ready-blue.svg)](https://github.com/strugglerx/wechat)
[![License](https://img.shields.io/github/license/strugglerx/wechat.svg?style=flat)](https://github.com/strugglerx/wechat)

# Installation
```
go get -u -v github.com/strugglerx/wechat
```
suggested using `go.mod`:
```
require github.com/strugglerx/wechat
```

# Usage
```golang

appid := "xx"
secret := "xx"

//sample
app := miniapp.App{Appid:  appid, Secret: secret,}

//verify 
app := New(&App{ Appid:  appid, Secret: secret, Verify: true,})

//hook
var cacheToken utils.Token
app := &miniapp.App{
    Appid:  appid,
    Secret: secret,
    Read: func(appid string) *utils.Token {return &cacheToken},
    Write: func(appid, accessToken string) *utils.Token {
            cacheToken.Token = accessToken
            cacheToken.UpdateTime = int(time.Now().Unix())
            return &cacheToken
        },
    }
	
//custom 
app.Custom.PostBody("/wxa/business/getuserphonenumber", []byte(`{
        "code":"03c52dedef3306d529d53bb31452ec9a2f46880b2040cec9d760876e821f9429"
    }`),true)

//To Do
```

# License

`Go Wechat` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.