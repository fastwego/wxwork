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

// Package customer_group 客户联系/客户群管理
package customer_group

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiList = "/cgi-bin/externalcontact/groupchat/list"
	apiGet  = "/cgi-bin/externalcontact/groupchat/get"
)

/*
获取客户群列表

该接口用于获取配置过客户群管理的客户群列表。

See: https://work.weixin.qq.com/api/doc/90000/90135/92120

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/list?access_token=ACCESS_TOKEN
*/
func List(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取客户群详情

需注意的是，如果发生群信息变动，会立即收到群变更事件，但是部分信息是异步处理，可能需要等一段时间调此接口才能得到最新结果

See: https://work.weixin.qq.com/api/doc/90000/90135/92122

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/get?access_token=ACCESS_TOKEN
*/
func Get(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}
