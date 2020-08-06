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

package corporation

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	eventtype "github.com/fastwego/wechat4work/corporation/type/type_event"
	messagetype "github.com/fastwego/wechat4work/corporation/type/type_message"

	"github.com/fastwego/wechat4work/util"
)

/*
响应微信请求 或 推送消息/事件 的服务器
*/
type Server struct {
	Ctx *App
}


/*
EchoStr 服务器接口校验

msg_signature=ASDFQWEXZCVAQFASDFASDFSS
&timestamp=13500001234
&nonce=123412323
&echostr=ENCRYPT_STR
 */
func (s *Server) EchoStr(writer http.ResponseWriter, request *http.Request) {
	echoStr := request.URL.Query().Get("echostr")
	strs := []string{
		request.URL.Query().Get("timestamp"),
		request.URL.Query().Get("nonce"),
		s.Ctx.Config.Token,
		echoStr,
	}
	sort.Strings(strs)

	if s.Ctx.Corporation.Logger != nil {
		s.Ctx.Corporation.Logger.Println(strs)
	}

	h := sha1.New()
	_, _ = io.WriteString(h, strings.Join(strs, ""))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	if signature == request.URL.Query().Get("msg_signature") {

		// 解密 echoStr
		_, msg, appId, err := util.AESDecryptMsg(echoStr, s.Ctx.Config.EncodingAESKey)
		if err != nil {
			return
		}
		if string(appId) != s.Ctx.Config.AgentId {
			if s.Ctx.Corporation.Logger != nil {
				s.Ctx.Corporation.Logger.Printf("AgentId Not Match: %s != %s",appId,s.Ctx.Config.AgentId)
			}
			return
		}

		writer.Write(msg)
		if s.Ctx.Corporation.Logger != nil {
			s.Ctx.Corporation.Logger.Println("echostr ", string(msg))
		}
	}


}


/*
ParseXML 解析微信推送过来的消息/事件

POST /api/callback?msg_signature=ASDFQWEXZCVAQFASDFASDFSS
&timestamp=13500001234
&nonce=123412323

<xml>
   <ToUserName><![CDATA[toUser]]></ToUserName>
   <AgentID><![CDATA[toAgentID]]></AgentID>
   <Encrypt><![CDATA[msg_encrypt]]></Encrypt>
</xml>
 */
func (s *Server) ParseXML(body []byte) (m interface{}, err error) {

	if s.Ctx.Corporation.Logger != nil {
		s.Ctx.Corporation.Logger.Println(string(body))
	}

	// 是否加密消息
	encryptMsg := messagetype.EncryptMessage{}
	err = xml.Unmarshal(body, &encryptMsg)
	if err != nil {
		return
	}

	// 需要解密
	if encryptMsg.Encrypt != "" {
		var xmlMsg []byte
		_, xmlMsg, _, err = util.AESDecryptMsg(encryptMsg.Encrypt, s.Ctx.Config.EncodingAESKey)
		if err != nil {
			return
		}
		body = xmlMsg

		if s.Ctx.Corporation.Logger != nil {
			s.Ctx.Corporation.Logger.Println("AESDecryptMsg ", string(body))
		}

	}

	message := messagetype.Message{}
	err = xml.Unmarshal(body, &message)
	fmt.Println(message)
	if err != nil {
		return
	}

	switch message.MsgType {
	case messagetype.MsgTypeText:
		msg := messagetype.MessageText{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeImage:
		msg := messagetype.MessageImage{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeVoice:
		msg := messagetype.MessageVoice{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeVideo:
		msg := messagetype.MessageVideo{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeLocation:
		msg := messagetype.MessageLocation{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeLink:
		msg := messagetype.MessageLink{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeEvent:
		return parseEvent(body)
	}
	return
}

// ParseEvent 解析微信推送过来的事件
func parseEvent(body []byte) (m interface{}, err error) {
	event := eventtype.Event{}
	err = xml.Unmarshal(body, &event)
	if err != nil {
		return
	}
	switch event.Event {
	// 事件
	case eventtype.EventTypeChangeContact:
		msg := eventtype.EventChangeContactCreateParty{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	}

	return
}

// Response 响应微信消息
func (s *Server) Response(writer http.ResponseWriter, request *http.Request, reply interface{}) (err error) {

	output := []byte("") // 默认回复
	if reply != nil {
		output, err = xml.Marshal(reply)
		if err != nil {
			return
		}

		// 加密
		message := s.encryptReplyMessage(output)
		output, err = xml.Marshal(message)
		if err != nil {
			return
		}

	}

	_, err = writer.Write(output)

	if s.Ctx.Corporation.Logger != nil {
		s.Ctx.Corporation.Logger.Println("Response: ", string(output))
	}

	return
}

// encryptReplyMessage 加密回复消息
func (s *Server) encryptReplyMessage(rawXmlMsg []byte) (replyEncryptMessage messagetype.ReplyEncryptMessage) {
	cipherText := util.AESEncryptMsg([]byte(util.GetRandString(16)), rawXmlMsg, s.Ctx.Config.AgentId, s.Ctx.Config.EncodingAESKey)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := util.GetRandString(6)

	strs := []string{
		timestamp,
		nonce,
		s.Ctx.Config.Token,
		cipherText,
	}
	sort.Strings(strs)
	h := sha1.New()
	_, _ = io.WriteString(h, strings.Join(strs, ""))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	return messagetype.ReplyEncryptMessage{
		Encrypt:      cipherText,
		MsgSignature: signature,
		TimeStamp:    timestamp,
		Nonce:        nonce,
	}
}
