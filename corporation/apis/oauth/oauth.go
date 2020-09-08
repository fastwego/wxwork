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

// Package oauth 网页授权登录(oauth)

package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	apiAuthorize = "https://open.weixin.qq.com/connect/oauth2/authorize"
	apiUserInfo  = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo"
)

/*
构造网页授权链接

如果用户同意授权，页面将跳转至 redirect_uri/?code=CODE&state=STATE

See: https://work.weixin.qq.com/api/doc/90000/90135/91022

GET https://open.weixin.qq.com/connect/oauth2/authorize?appid=CORPID&redirect_uri=REDIRECT_URI&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect
*/
func GetAuthorizeUrl(appid string, redirectUri string, state string) (authorizeUrl string) {
	params := url.Values{}
	params.Add("appid", appid)
	params.Add("redirect_uri", redirectUri)
	params.Add("response_type", "code")
	params.Add("scope", "snsapi_base")
	params.Add("state", state)
	return apiAuthorize + "?" + params.Encode() + "#wechat_redirect"
}

type UserInfo struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	UserID   string `json:"UserId"`
	DeviceID string `json:"DeviceId"`
}

/*
获取访问用户身份

该接口用于根据code获取成员信息

See: https://work.weixin.qq.com/api/doc/90000/90135/91023

GET https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=ACCESS_TOKEN&code=CODE
*/
func GetUserInfo(accessToken string, code string) (userInfo UserInfo, err error) {
	params := url.Values{}
	params.Add("access_token", accessToken)
	params.Add("code", code)

	uri := apiUserInfo + "?" + params.Encode()
	response, err := http.Get(uri)
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		err = fmt.Errorf("%s", string(body))
		return
	}

	return
}
