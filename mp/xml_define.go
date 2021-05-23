package mp

import "encoding/xml"

/**
 * @PROJECT_NAME wechat
 * @author  Moqi
 * @date  2021-05-23 18:53
 * @Email:str@li.cm
 **/

type CDATA struct {
	Text string `xml:",cdata"`
}

//返回信息的结构体
/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[text]]></MsgType>
  <Content><![CDATA[你好]]></Content>
</xml>
*/
type SendXmlText struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Content      CDATA    `xml:"Content"`
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[news]]></MsgType>
  <ArticleCount>1</ArticleCount>
  <Articles>
    <item>
      <Title><![CDATA[title1]]></Title>
      <Description><![CDATA[description1]]></Description>
      <PicUrl><![CDATA[picurl]]></PicUrl>
      <Url><![CDATA[url]]></Url>
    </item>
  </Articles>
</xml>
*/
type SendArticle struct {
	XMLName      xml.Name                  `xml:"xml"`
	ToUserName   CDATA                     `xml:"ToUserName"`
	FromUserName CDATA                     `xml:"FromUserName"`
	CreateTime   int                       `xml:"CreateTime"`
	MsgType      CDATA                     `xml:"MsgType"`
	ArticleCount int                       `xml:"ArticleCount"`
	Articles     []SendArticleItem `xml:"Articles>item"`
}

type SendArticleItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       CDATA    `xml:"Title"`
	Description CDATA    `xml:"Description"`
	PicUrl      CDATA    `xml:"PicUrl"`
	Url         CDATA    `xml:"Url"`
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[image]]></MsgType>
  <Image>
    <MediaId><![CDATA[media_id]]></MediaId>
  </Image>
</xml>
*/
type SendXmlImage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Image        struct {
		MediaId CDATA `xml:"MediaId"`
	} `xml:"Image"`
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[voice]]></MsgType>
  <Voice>
    <MediaId><![CDATA[media_id]]></MediaId>
  </Voice>
</xml>
*/
type SendXmlVoice struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Voice        struct {
		MediaId CDATA `xml:"MediaId"`
	} `xml:"Voice"`
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[video]]></MsgType>
  <Video>
    <MediaId><![CDATA[media_id]]></MediaId>
    <Title><![CDATA[title]]></Title>
    <Description><![CDATA[description]]></Description>
  </Video>
</xml>
*/
type SendXmlVideo struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Video        struct {
		MediaId     CDATA `xml:"MediaId"`
		Title       CDATA `xml:"Title"`
		Description CDATA `xml:"Description"`
	} `xml:"Video"`
}