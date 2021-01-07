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

// Package customer_moment 客户联系/客户朋友圈
package customer_moment

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetMomentList         = "/cgi-bin/externalcontact/get_moment_list"
	apiGetMomentTask         = "/cgi-bin/externalcontact/get_moment_task"
	apiGetMomentCustomerList = "/cgi-bin/externalcontact/get_moment_customer_list"
	apiGetMomentSendResult   = "/cgi-bin/externalcontact/get_moment_send_result"
	apiGetMomentComments     = "/cgi-bin/externalcontact/get_moment_comments"
)

/*
获取企业全部的发表列表

企业和第三方应用可通过该接口获取企业全部的发表内容。

See: https://work.weixin.qq.com/api/doc/90000/90135/93333

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_list?access_token=ACCESS_TOKEN
*/
func GetMomentList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetMomentList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取客户朋友圈企业发表的列表

企业和第三方应用可通过该接口获取企业发表的朋友圈成员执行情况

See: https://work.weixin.qq.com/api/doc/90000/90135/93333

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_task?access_token=ACCESS_TOKEN
*/
func GetMomentTask(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetMomentTask, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取客户朋友圈发表时选择的可见范围



See: https://work.weixin.qq.com/api/doc/90000/90135/93333

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_customer_list?access_token=ACCESS_TOKEN
*/
func GetMomentCustomerList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetMomentCustomerList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取客户朋友圈发表后的可见客户列表



See: https://work.weixin.qq.com/api/doc/90000/90135/93333

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_send_result?access_token=ACCESS_TOKEN
*/
func GetMomentSendResult(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetMomentSendResult, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取客户朋友圈的互动数据

企业和第三方应用可通过此接口获取客户朋友圈的互动数据。

See: https://work.weixin.qq.com/api/doc/90000/90135/93333

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_comments?access_token=ACCESS_TOKEN
*/
func GetMomentComments(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetMomentComments, bytes.NewReader(payload), "application/json;charset=utf-8")
}
