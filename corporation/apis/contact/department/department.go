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

// Package department 通讯录管理/部门管理
package department

import (
	"bytes"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiCreate = "/cgi-bin/department/create"
	apiUpdate = "/cgi-bin/department/update"
	apiDelete = "/cgi-bin/department/delete"
	apiList   = "/cgi-bin/department/list"
)

/*
创建部门



See: https://work.weixin.qq.com/api/doc/90000/90135/90205

POST https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新部门



See: https://work.weixin.qq.com/api/doc/90000/90135/90206

POST https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除部门



See: https://work.weixin.qq.com/api/doc/90000/90135/90207

GET https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=ACCESS_TOKEN&id=ID
*/
func Delete(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiDelete + "?" + params.Encode())
}

/*
获取部门列表



See: https://work.weixin.qq.com/api/doc/90000/90135/90208

GET https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=ACCESS_TOKEN&id=ID
*/
func List(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiList + "?" + params.Encode())
}
