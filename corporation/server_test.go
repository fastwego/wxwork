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
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"

	"github.com/fastwego/wechat4work/corporation/type/type_message"
)

var MockCorporation *Corporation
var MockApp *App

func TestMain(m *testing.M) {
	MockCorporation = New(CorporationConfig{
		Corpid: "CROPID",
	})
	MockApp = MockCorporation.NewApp(AppConfig{
		AgentId:        "AGENTID",
		Secret:         "SERET",
		Token:          "Token",
		EncodingAESKey: "EncodingAESKey",
	})
	os.Exit(m.Run())
}

func TestServer_ParseXML(t *testing.T) {

	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		args    args
		wantM   interface{}
		wantErr bool
	}{
		{
			name: "case1",
			args: args{body: []byte(`
			<xml>
			  <ToUserName><![CDATA[toUser]]></ToUserName>
			  <FromUserName><![CDATA[fromUser]]></FromUserName>
			  <CreateTime>1348831860</CreateTime>
			  <MsgType><![CDATA[text]]></MsgType>
			  <Content><![CDATA[this is a test]]></Content>
			  <MsgId>1234567890123456</MsgId>
			</xml>
			`)},
			wantM: type_message.MessageText{
				Message: type_message.Message{
					ToUserName:   "toUser",
					FromUserName: "fromUser",
					CreateTime:   "1348831860",
					MsgType:      "text",
				},
				MsgId:   "1234567890123456",
				Content: "this is a test",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM, err := MockApp.Server.ParseXML(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("ParseXML() gotM = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func TestReplyMessage(t *testing.T) {
	tests := []struct {
		name     string
		wantEcho string
	}{
		{
			name:     "case1",
			wantEcho: `<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>12345678</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[你好]]></Content></xml>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 回复文本消息
			msg := type_message.ReplyMessageText{
				ReplyMessage: type_message.ReplyMessage{
					ToUserName:   "toUser",
					FromUserName: "fromUser",
					CreateTime:   "12345678",
					MsgType:      type_message.ReplyMsgTypeText,
				},
				Content: "你好",
			}
			data, err := xml.Marshal(msg)
			fmt.Println(string(data), err)

			if tt.wantEcho != string(data) {
				t.Errorf("\nwant %s \nget %s", tt.wantEcho, string(data))
			}
		})
	}
}

func TestServer_EchoStr(t *testing.T) {
	MockApp.Config.Token = "QDG6eK"
	MockApp.Config.AgentId = "wx5823bf96d3bd56c7"
	MockApp.Config.EncodingAESKey="jWmYm7qr5nMoAUwZRjGtBxmz3KA1tkAj3ykkR6q2B2C"
	s := &Server{
		Ctx: MockApp,
	}

	// Mock Server
	MockSvrHandler := http.NewServeMux()
	MockSvrHandler.HandleFunc("/echostr", s.EchoStr)

	tests := []struct {
		name     string
		args     url.Values
		wantEcho string
	}{
		{
			name: "case1",
			args: url.Values{
				"timestamp": []string{"1409659589"},
				"nonce":     []string{"263014780"},
				"echostr":   []string{"P9nAzCzyDtyTWESHep1vC5X9xho/qYX3Zpb4yKa9SKld1DsH3Iyt3tP3zNdtp+4RPcs8TgAE7OaBO+FZXvnaqQ=="},
				"msg_signature": []string{"5c45ff5e21c57e6ad56bac8758b79b1d9ac89fd3"},
			},
			wantEcho: "1616140317555161061",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest(http.MethodGet, "/echostr?"+tt.args.Encode(), nil)
			w := httptest.NewRecorder()
			MockSvrHandler.ServeHTTP(w, r)
			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Response code is %v", resp.StatusCode)
			}

			echo := string(w.Body.Bytes())
			fmt.Println(echo)

			if tt.wantEcho != echo {
				t.Errorf("want %s but get %s", tt.wantEcho, echo)
			}
		})
	}
}
