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

// Package linked_corp 通讯录管理/互联企业
package linked_corp

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetPermList    = "/cgi-bin/linkedcorp/agent/get_perm_list"
	apiGet            = "/cgi-bin/linkedcorp/user/get"
	apiSimpleList     = "/cgi-bin/linkedcorp/user/simplelist"
	apiUserList       = "/cgi-bin/linkedcorp/user/list"
	apiDepartmentList = "/cgi-bin/linkedcorp/department/list"
)

/*
获取应用的可见范围

本接口只返回互联企业中非本企业内的成员和部门的信息，如果要获取本企业的可见范围，请调用“获取应用”接口

See: https://work.weixin.qq.com/api/doc/90000/90135/93172

POST https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/agent/get_perm_list?access_token=ACCESS_TOKEN
*/
func GetPermList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetPermList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取互联企业成员详细信息



See: https://work.weixin.qq.com/api/doc/90000/90135/93171

POST https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/get?access_token=ACCESS_TOKEN
*/
func Get(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取互联企业部门成员



See: https://work.weixin.qq.com/api/doc/90000/90135/93168

POST https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/simplelist?access_token=ACCESS_TOKEN
*/
func SimpleList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSimpleList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取互联企业部门成员详情



See: https://work.weixin.qq.com/api/doc/90000/90135/93169

POST https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/list?access_token=ACCESS_TOKEN
*/
func UserList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUserList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取互联企业部门列表



See: https://work.weixin.qq.com/api/doc/90000/90135/93170

POST https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/department/list?access_token=ACCESS_TOKEN
*/
func DepartmentList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDepartmentList, bytes.NewReader(payload), "application/json;charset=utf-8")
}
