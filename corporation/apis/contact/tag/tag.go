// Copyright 2021 FastWeGo
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

// Package tag 通讯录管理/标签管理
package tag

import (
	"bytes"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiCreate      = "/cgi-bin/tag/create"
	apiUpdate      = "/cgi-bin/tag/update"
	apiDelete      = "/cgi-bin/tag/delete"
	apiGet         = "/cgi-bin/tag/get"
	apiAddTagUsers = "/cgi-bin/tag/addtagusers"
	apiDelTagUsers = "/cgi-bin/tag/deltagusers"
	apiList        = "/cgi-bin/tag/list"
)

/*
创建标签



See: https://work.weixin.qq.com/api/doc/90000/90135/90210

POST https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新标签名字



See: https://work.weixin.qq.com/api/doc/90000/90135/90211

POST https://qyapi.weixin.qq.com/cgi-bin/tag/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除标签



See: https://work.weixin.qq.com/api/doc/90000/90135/90212

GET https://qyapi.weixin.qq.com/cgi-bin/tag/delete?access_token=ACCESS_TOKEN&tagid=TAGID
*/
func Delete(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiDelete + "?" + params.Encode())
}

/*
获取标签成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90213

GET https://qyapi.weixin.qq.com/cgi-bin/tag/get?access_token=ACCESS_TOKEN&tagid=TAGID
*/
func Get(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGet + "?" + params.Encode())
}

/*
增加标签成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90214

POST https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers?access_token=ACCESS_TOKEN
*/
func AddTagUsers(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddTagUsers, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除标签成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90215

POST https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers?access_token=ACCESS_TOKEN
*/
func DelTagUsers(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelTagUsers, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取标签列表



See: https://work.weixin.qq.com/api/doc/90000/90135/90216

GET https://qyapi.weixin.qq.com/cgi-bin/tag/list?access_token=ACCESS_TOKEN
*/
func List(ctx *corporation.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiList)
}
