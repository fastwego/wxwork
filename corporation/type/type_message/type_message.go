// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package type_message

import "encoding/xml"

const (
	MsgTypeText     = "text"
	MsgTypeImage    = "image"
	MsgTypeVoice    = "voice"
	MsgTypeVideo    = "video"
	MsgTypeLocation = "location"
	MsgTypeLink     = "link"
	MsgTypeEvent    = "event"
)

type Message struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	MsgId        string
	AgentID      string
}

/*
加密消息格式
<xml>
   <ToUserName><![CDATA[toUser]]></ToUserName>
   <AgentID><![CDATA[toAgentID]]></AgentID>
   <Encrypt><![CDATA[msg_encrypt]]></Encrypt>
</xml>
*/
type EncryptMessage struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string
	Encrypt    string
	AgentID    string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1348831860</CreateTime>
  <MsgType><![CDATA[text]]></MsgType>
  <Content><![CDATA[this is a test]]></Content>
  <MsgId>1234567890123456</MsgId>
</xml>
*/
type MessageText struct {
	Message
	Content string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1348831860</CreateTime>
  <MsgType><![CDATA[image]]></MsgType>
  <PicUrl><![CDATA[this is a url]]></PicUrl>
  <MediaId><![CDATA[media_id]]></MediaId>
  <MsgId>1234567890123456</MsgId>
</xml>
*/
type MessageImage struct {
	Message
	PicUrl  string
	MediaId string
}

/*
<xml>
  <ToUserName>< ![CDATA[toUser] ]></ToUserName>
  <FromUserName>< ![CDATA[fromUser] ]></FromUserName>
  <CreateTime>1357290913</CreateTime>
  <MsgType>< ![CDATA[voice] ]></MsgType>
  <MediaId>< ![CDATA[media_id] ]></MediaId>
  <Format>< ![CDATA[Format] ]></Format>
  <Recognition>< ![CDATA[腾讯微信团队] ]></Recognition>
  <MsgId>1234567890123456</MsgId>
</xml>

*/
type MessageVoice struct {
	Message
	Format      string
	Recognition string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1357290913</CreateTime>
  <MsgType><![CDATA[video]]></MsgType>
  <MediaId><![CDATA[media_id]]></MediaId>
  <ThumbMediaId><![CDATA[thumb_media_id]]></ThumbMediaId>
  <MsgId>1234567890123456</MsgId>
</xml>

*/
type MessageVideo struct {
	Message
	MediaId      string
	ThumbMediaId string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1351776360</CreateTime>
  <MsgType><![CDATA[location]]></MsgType>
  <Location_X>23.134521</Location_X>
  <Location_Y>113.358803</Location_Y>
  <Scale>20</Scale>
  <Label><![CDATA[位置信息]]></Label>
  <MsgId>1234567890123456</MsgId>
</xml>

*/
type MessageLocation struct {
	Message
	Location_X string
	Location_Y string
	Scale      string
	Label      string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1351776360</CreateTime>
  <MsgType><![CDATA[link]]></MsgType>
  <Title><![CDATA[公众平台官网链接]]></Title>
  <Description><![CDATA[公众平台官网链接]]></Description>
  <Url><![CDATA[url]]></Url>
  <MsgId>1234567890123456</MsgId>
</xml>
*/
type MessageLink struct {
	Message
	Title       string
	Description string
	Url         string
}
