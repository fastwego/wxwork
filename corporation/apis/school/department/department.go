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

// Package department 家校沟通/部门管理
package department

import (
	"bytes"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiCreate         = "/cgi-bin/school/department/create"
	apiUpdate         = "/cgi-bin/school/department/update"
	apiDelete         = "/cgi-bin/school/department/delete"
	apiList           = "/cgi-bin/school/department/list"
	apiSetUpgradeInfo = "/cgi-bin/school/set_upgrade_info"
)

/*
创建部门



See: https://work.weixin.qq.com/api/doc/90000/90135/92340

POST https://qyapi.weixin.qq.com/cgi-bin/school/department/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新部门



See: https://work.weixin.qq.com/api/doc/90000/90135/92341

POST https://qyapi.weixin.qq.com/cgi-bin/school/department/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除部门



See: https://work.weixin.qq.com/api/doc/90000/90135/92342

GET https://qyapi.weixin.qq.com/cgi-bin/school/department/delete?access_token=ACCESS_TOKEN&id=ID
*/
func Delete(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiDelete + "?" + params.Encode())
}

/*
获取部门列表



See: https://work.weixin.qq.com/api/doc/90000/90135/92343

GET https://qyapi.weixin.qq.com/cgi-bin/school/department/list?access_token=ACCESS_TOKEN&id=ID
*/
func List(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiList + "?" + params.Encode())
}

/*
修改自动升年级的配置



See: https://work.weixin.qq.com/api/doc/90000/90135/92949

POST https://qyapi.weixin.qq.com/cgi-bin/school/set_upgrade_info?access_token=ACCESS_TOKEN
*/
func SetUpgradeInfo(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetUpgradeInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}
