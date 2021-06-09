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
app := miniapp.New(appid,secret)

//hook
var cacheToken utils.Token
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

//To Do
```

# License

`Go Wechat` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.