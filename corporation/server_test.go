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
	"bytes"
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/fastwego/wxwork/corporation/type/type_event"
	"github.com/fastwego/wxwork/corporation/type/type_message"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sort"
	"strings"
	"testing"
)

var MockCorporation *Corporation
var MockApp *App

func TestMain(m *testing.M) {
	MockCorporation = New(Config{
		Corpid: "CROPID",
	})
	MockCorporation.SetLogger(nil)
	MockApp = MockCorporation.NewApp(AppConfig{
		AgentId:        "AGENTID",
		Secret:         "SERET",
		Token:          "Token",
		EncodingAESKey: "TfuodSKwmagZ0iCvQU2yfOWvOt8VLU5S5D85PcbOCMs",
	})
	os.Exit(m.Run())
}

func TestServer_ParseXML(t *testing.T) {

	type args struct {
		params url.Values
		body []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				body: []byte(`<xml><ToUserName><![CDATA[ww4ffb2457362fa91a]]></ToUserName><Encrypt><![CDATA[yzY09xkWVovVOxPTwSYfIFVRO92WzZiKxWXGZLeoAMYPqh1DLdt3Ph3vASJHpLZ3DZF8laEEx1A1ySJGogWBOKZn9QFqANkFj4GnWX2qheHLuCyCvHtHp/Jx6qnqQTRL0RpRvBa2OPCuMufQYEFi0UX4RpXsxS2oQnOCcnob3PJxIOFIs5kc4hxgRpbiAjy3f54tbwyJzzO2qq4ZpzpxOsmFUsJhuWROhK+ltW8VvSrVjhfeFZmcF3WLzk8ShSXowuQeF69kP80gZAbpdT/NFkeuP7L45+o5z/g5D+0xbalxWctGr5yM1C6p8ILJGDRUP2JdoKd3n+PBSKSQJBz98sgRJgjDZyCbxFa03yq77UbZn3sQKpwqFInqorHUvn47YWcQoSsohLiuhdB4m9zVa/sQ0f+cT9umemuQjvT7u6aNr1SAtYCN6kTkTh2gUMxhRfr+iy/kU7NK8dKQojBRXw==]]></Encrypt><AgentID><![CDATA[1000004]]></AgentID></xml>`),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			timestamp := "1596792211"
			nonce := "nonce"
			encryptMsg := "yzY09xkWVovVOxPTwSYfIFVRO92WzZiKxWXGZLeoAMYPqh1DLdt3Ph3vASJHpLZ3DZF8laEEx1A1ySJGogWBOKZn9QFqANkFj4GnWX2qheHLuCyCvHtHp/Jx6qnqQTRL0RpRvBa2OPCuMufQYEFi0UX4RpXsxS2oQnOCcnob3PJxIOFIs5kc4hxgRpbiAjy3f54tbwyJzzO2qq4ZpzpxOsmFUsJhuWROhK+ltW8VvSrVjhfeFZmcF3WLzk8ShSXowuQeF69kP80gZAbpdT/NFkeuP7L45+o5z/g5D+0xbalxWctGr5yM1C6p8ILJGDRUP2JdoKd3n+PBSKSQJBz98sgRJgjDZyCbxFa03yq77UbZn3sQKpwqFInqorHUvn47YWcQoSsohLiuhdB4m9zVa/sQ0f+cT9umemuQjvT7u6aNr1SAtYCN6kTkTh2gUMxhRfr+iy/kU7NK8dKQojBRXw=="

			strs := []string{
				timestamp,
				nonce,
				MockApp.Config.Token,
				encryptMsg,
			}
			sort.Strings(strs)

			h := sha1.New()
			_, _ = io.WriteString(h, strings.Join(strs, ""))
			signature := fmt.Sprintf("%x", h.Sum(nil))

			tt.args.params = url.Values{
				`msg_signature`:[]string{signature},
				`timestamp`:[]string{timestamp},
				`nonce`:[]string{nonce},
			}

			request, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1?"+tt.args.params.Encode(), bytes.NewReader(tt.args.body))

			gotM, err := MockApp.Server.ParseXML(request)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			_, ok := gotM.(type_event.EventMenuView)
			if !ok {
				t.Error("Not type_event.EventMenuView")
				return
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

func TestServer_Response(t *testing.T) {

	mockRequest := httptest.NewRequest("GET", "http://example.com/foo", nil)
	recorder := httptest.NewRecorder()

	type args struct {
		writer  http.ResponseWriter
		request *http.Request
		reply   interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case1",args: args{writer:recorder, request: mockRequest,reply: ""},wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MockApp.Server.Response(tt.args.writer, tt.args.request, tt.args.reply)
			if  (err != nil) != tt.wantErr {
				t.Errorf("Response() error = %v, wantErr %v", err, tt.wantErr)
			}

			resp := recorder.Result()
			body, _ := ioutil.ReadAll(resp.Body)

			encryptReply := type_message.ReplyEncryptMessage{}
			_ = xml.Unmarshal(body, &encryptReply)

			if  encryptReply.Encrypt == "" {
				t.Errorf("Encrypt not found")
			}
		})
	}
}