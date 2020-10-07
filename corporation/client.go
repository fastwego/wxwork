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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var (
	WXServerUrl      = "https://qyapi.weixin.qq.com" // 微信 api 服务器地址
	UserAgent        = "fastwego/wxwork"
	ErrorAccessToken = errors.New("access token error")
	ErrorSystemBusy  = errors.New("system busy")
)

/*
HttpClient 用于向微信接口发送请求
*/
type Client struct {
	Ctx *App
}

// HTTPGet GET 请求
func (client *Client) HTTPGet(uri string) (resp []byte, err error) {
	newUrl, err := client.applyAccessToken(uri)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodGet, WXServerUrl+newUrl, nil)
	if err != nil {
		return
	}

	return client.httpDo(req)
}

//HTTPPost POST 请求
func (client *Client) HTTPPost(uri string, payload io.Reader, contentType string) (resp []byte, err error) {
	newUrl, err := client.applyAccessToken(uri)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, WXServerUrl+newUrl, payload)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", contentType)

	return client.httpDo(req)
}

//httpDo 执行 请求
func (client *Client) httpDo(req *http.Request) (resp []byte, err error) {
	req.Header.Add("User-Agent", UserAgent)

	if client.Ctx.Corporation.Logger != nil {
		client.Ctx.Corporation.Logger.Printf("%s %s Headers %v", req.Method, req.URL.String(), req.Header)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	resp, err = responseFilter(response)

	// 发现 access_token 过期
	if err == ErrorAccessToken {

		// 主动 通知 access_token 过期
		err = client.Ctx.AccessToken.NoticeAccessTokenExpireHandler(client.Ctx)
		if err != nil {
			return
		}

		// 通知到位后 access_token 会被刷新，那么可以 retry 了
		var accessToken string
		accessToken, err = client.Ctx.AccessToken.GetAccessTokenHandler(client.Ctx)
		if err != nil {
			return
		}

		// 换新
		q := req.URL.Query()
		q.Set("access_token", accessToken)
		req.URL.RawQuery = q.Encode()

		if client.Ctx.Corporation.Logger != nil {
			client.Ctx.Corporation.Logger.Printf("%v retry %s %s Headers %v", ErrorAccessToken, req.Method, req.URL.String(), req.Header)
		}

		response, err = http.DefaultClient.Do(req)
		if err != nil {
			return
		}
		defer response.Body.Close()

		resp, err = responseFilter(response)
	}

	// -1 系统繁忙，此时请开发者稍候再试
	// 重试一次
	if err == ErrorSystemBusy {

		if client.Ctx.Corporation.Logger != nil {
			client.Ctx.Corporation.Logger.Printf("%v : retry %s %s Headers %v", ErrorSystemBusy, req.Method, req.URL.String(), req.Header)
		}

		response, err = http.DefaultClient.Do(req)
		if err != nil {
			return
		}
		defer response.Body.Close()

		resp, err = responseFilter(response)
	}

	return
}

/*
在请求地址上附加上 access_token
*/
func (client *Client) applyAccessToken(oldUrl string) (newUrl string, err error) {
	accessToken, err := client.Ctx.AccessToken.GetAccessTokenHandler(client.Ctx)
	if err != nil {
		return
	}
	if strings.Contains(oldUrl, "?") {
		newUrl = oldUrl + "&access_token=" + accessToken
	} else {
		newUrl = oldUrl + "?access_token=" + accessToken
	}
	return
}

/*
筛查微信 api 服务器响应，判断以下错误：

- http 状态码 不为 200

- 接口响应错误码 errcode 不为 0
*/
func responseFilter(response *http.Response) (resp []byte, err error) {
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("Status %s", response.Status)
		return
	}

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	errorResponse := struct {
		Errcode int64  `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}{}
	err = json.Unmarshal(resp, &errorResponse)
	if err != nil {
		return
	}

	if errorResponse.Errcode == 40014 {
		err = ErrorAccessToken
		return
	}

	//  -1	系统繁忙，此时请开发者稍候再试
	if errorResponse.Errcode == -1 {
		err = ErrorSystemBusy
		return
	}

	if errorResponse.Errcode != 0 {
		err = errors.New(string(resp))
		return
	}
	return
}

// 防止多个 goroutine 并发刷新冲突
var refreshAccessTokenLock sync.Mutex

/*
从 公众号实例 的 AccessToken 管理器 获取 access_token

如果没有 access_token 或者 已过期，那么刷新
*/
func GetAccessToken(app *App) (accessToken string, err error) {
	cacheKey := app.Config.AgentId + app.Config.Secret // 企业微信-系统应用没有 agentid ，所以需要secret 辅助
	accessToken, err = app.AccessToken.Cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}

	refreshAccessTokenLock.Lock()
	defer refreshAccessTokenLock.Unlock()

	accessToken, err = app.AccessToken.Cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}

	accessToken, expiresIn, err := refreshAccessTokenFromWXServer(app.Corporation.Config.Corpid, app.Config.Secret)
	if err != nil {
		return
	}

	d := time.Duration(expiresIn) * time.Second
	_ = app.AccessToken.Cache.Save(cacheKey, accessToken, d)

	if app.Corporation.Logger != nil {
		app.Corporation.Logger.Printf("%s %s %d\n", "refreshAccessTokenFromWXServer", accessToken, expiresIn)
	}

	return
}

/*
NoticeAccessTokenExpire 只需将本地存储的 access_token 删除，即完成了 access_token 已过期的 主动通知

retry 请求的时候，会发现本地没有 access_token ，从而触发refresh
*/
func NoticeAccessTokenExpire(app *App) (err error) {
	if app.Corporation.Logger != nil {
		app.Corporation.Logger.Println("NoticeAccessTokenExpire")
	}

	cacheKey := app.Config.AgentId + app.Config.Secret // 企业微信-系统应用没有 agentid ，所以需要secret 辅助
	err = app.AccessToken.Cache.Delete(cacheKey)
	return
}

/*
从微信服务器获取新的 AccessToken

See: https://developers.weixin.qq.com/doc/corporation/Basic_Information/Get_access_token.html
*/
func refreshAccessTokenFromWXServer(appid string, secret string) (accessToken string, expiresIn int, err error) {
	params := url.Values{}
	params.Add("corpid", appid)
	params.Add("corpsecret", secret)
	url := WXServerUrl + "/cgi-bin/gettoken?" + params.Encode()

	response, err := http.Get(url)
	if err != nil {
		return
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("GET %s RETURN %s", url, response.Status)
		return
	}

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var result = struct {
		AccessToken string  `json:"access_token"`
		ExpiresIn   int     `json:"expires_in"`
		Errcode     float64 `json:"errcode"`
		Errmsg      string  `json:"errmsg"`
	}{}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		err = fmt.Errorf("Unmarshal error %s", string(resp))
		return
	}

	if result.AccessToken == "" {
		err = fmt.Errorf("%s", string(resp))
		return
	}

	return result.AccessToken, result.ExpiresIn, nil
}
