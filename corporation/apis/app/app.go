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

// Package app 应用管理
package app

import (
	"bytes"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiAgentGet             = "/cgi-bin/agent/get"
	apiAgentList            = "/cgi-bin/agent/list"
	apiAgentSet             = "/cgi-bin/agent/set"
	apiMenuCreate           = "/cgi-bin/menu/create"
	apiMenuGet              = "/cgi-bin/menu/get"
	apiMenuDelete           = "/cgi-bin/menu/delete"
	apiSetWorkbenchTemplate = "/cgi-bin/agent/set_workbench_template"
	apiGetWorkbenchTemplate = "/cgi-bin/agent/get_workbench_template"
	apiSetWorkbenchData     = "/cgi-bin/agent/set_workbench_data"
)

/*
获取应用



See: https://work.weixin.qq.com/api/doc/90000/90135/90227

GET https://qyapi.weixin.qq.com/cgi-bin/agent/get?access_token=ACCESS_TOKEN&agentid=AGENTID参数说明
*/
func AgentGet(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiAgentGet + "?" + params.Encode())
}

/*
获取应用



See: https://work.weixin.qq.com/api/doc/90000/90135/90227

GET https://qyapi.weixin.qq.com/cgi-bin/agent/list?access_token=ACCESS_TOKEN
*/
func AgentList(ctx *corporation.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiAgentList)
}

/*
设置应用



See: https://work.weixin.qq.com/api/doc/90000/90135/90228

POST https://qyapi.weixin.qq.com/cgi-bin/agent/set?access_token=ACCESS_TOKEN
*/
func AgentSet(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAgentSet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
创建菜单



See: https://work.weixin.qq.com/api/doc/90000/90135/90231

POST https://qyapi.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN&agentid=AGENTID
*/
func MenuCreate(ctx *corporation.App, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMenuCreate+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取菜单



See: https://work.weixin.qq.com/api/doc/90000/90135/90232

GET https://qyapi.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN&agentid=AGENTID
*/
func MenuGet(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiMenuGet + "?" + params.Encode())
}

/*
删除菜单



See: https://work.weixin.qq.com/api/doc/90000/90135/90233

GET https://qyapi.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN&agentid=AGENTID
*/
func MenuDelete(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiMenuDelete + "?" + params.Encode())
}

/*
设置工作台自定义展示



See: https://work.weixin.qq.com/api/doc/90000/90135/92535

POST https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_template?access_token=ACCESS_TOKEN
*/
func SetWorkbenchTemplate(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetWorkbenchTemplate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
设置工作台自定义展示



See: https://work.weixin.qq.com/api/doc/90000/90135/92535

POST https://qyapi.weixin.qq.com/cgi-bin/agent/get_workbench_template?access_token=ACCESS_TOKEN
*/
func GetWorkbenchTemplate(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetWorkbenchTemplate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
设置工作台自定义展示



See: https://work.weixin.qq.com/api/doc/90000/90135/92535

POST https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_data?access_token=ACCESS_TOKEN
*/
func SetWorkbenchData(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetWorkbenchData, bytes.NewReader(payload), "application/json;charset=utf-8")
}
