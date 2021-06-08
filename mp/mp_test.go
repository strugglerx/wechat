package mp

import "testing"

func TestGetAccessToken(t *testing.T) {
	mp := &Mp{
		Appid:"wxac5621df42a3bc11",
		Secret: "b7c259a57d55dd109a01dc3b578cg316",
	}
	t.Logf("%+v",mp.GetAccessToken())
}

func TestCallbackIp(t *testing.T) {
	mp := &Mp{
		Appid:"wxac5621df42a3bc11",
		Secret: "b7c259a57d55dd109a01dc3b578cg316",
	}
	t.Log(mp.CallbackIp())
}

func TestSetMenu(t *testing.T) {
	mp := &Mp{
		Appid:"wxac5621df42a3bc11",
		Secret: "b7c259a57d55dd109a01dc3b578cg316",
	}
	t.Log(mp.SetMenu([]byte(` {
     "button":[
     {	
          "type":"click",
          "name":"今日歌曲",
          "key":"V1001_TODAY_MUSIC"
      },
      {
           "name":"菜单",
           "sub_button":[
           {	
               "type":"view",
               "name":"搜索",
               "url":"http://www.soso.com/"
            },
            {
               "type":"click",
               "name":"赞一下我们",
               "key":"V1001_GOOD"
            }]
       }]
 }`)))
}

func TestShortUrl(t *testing.T) {
	mp := &Mp{
		Appid:"wxac5621df42a3bc11",
		Secret: "b7c259a57d55dd109a01dc3b578cg316",
	}
	t.Log(mp.GetShortUrl("http://www.baidu.com"))
}

func TestCreateMoreAcode(t *testing.T) {
	mp := &Mp{
		Appid:"wxac5621df42a3bc11",
		Secret: "b7c259a57d55dd109a01dc3b578cg316",
	}
	t.Log(mp.CreateMoreAcode("QR_STR_SCENE", "helloworld"))
}

func TestGetUserList(t *testing.T) {
	mp := &Mp{
		Appid:"wxac5621df42a3bc11",
		Secret: "b7c259a57d55dd109a01dc3b578cg316",
	}
	t.Log(mp.GetUserList(""))
}

func TestGetUserDetail(t *testing.T) {
	mp := &Mp{
		Appid:"wxac5621df42a3bc11",
		Secret: "b7c259a57d55dd109a01dc3b578cg316",
	}
	t.Log(mp.GetUserDetail("o8zsstzrjSvlC9MHI6yfOUMNNi5Q"))
}

func TestGetUsersDetail(t *testing.T) {
	mp := &Mp{
		Appid:"wxac5621df42a3bc11",
		Secret: "b7c259a57d55dd109a01dc3b578cg316",
	}
	res, _ := mp.GetUsersDetail([]byte(`{
    "user_list": [
        {
            "openid": "o8zsstzrjSvlC9MHI6yfOUMNNi5Q", 
            "lang": "zh_CN"
        },
    ]
	}`))
	t.Log(res)
}

func TestPushTemplate(t *testing.T) {
	mp := &Mp{
		Appid:"wxac5621df42a3bc11",
		Secret: "b7c259a57d55dd109a01dc3b578cg316",
	}
	mp.init()
	mp.PushTemplate([]byte(`{
           "touser":"o8zsstzrjSvlC9MHI6yfOUMNNi5Q",
           "template_id":"TyU1XvqQVc9XEQDGxsR9bro2gkIigCpriLyMm4jvL2o",
           "url":"http://www.baidu.com",
           "data":{
                   "test1": {
                       "value":"恭喜你购买成功！",
                       "color":"#173177"
                   },
                   "test2":{
                       "value":"巧克力",
                       "color":"#173177"
					}
				}
    	}`))
}
