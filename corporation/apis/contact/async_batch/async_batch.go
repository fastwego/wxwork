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

// Package async_batch 通讯录管理/异步批量接口
package async_batch

import (
	"bytes"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiUser         = "/cgi-bin/batch/syncuser"
	apiReplaceUser  = "/cgi-bin/batch/replaceuser"
	apiReplaceParty = "/cgi-bin/batch/replaceparty"
	apiGetResult    = "/cgi-bin/batch/getresult"
)

/*
增量更新成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90980

POST https://qyapi.weixin.qq.com/cgi-bin/batch/syncuser?access_token=ACCESS_TOKEN
*/
func User(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUser, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
全量覆盖成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90981

POST https://qyapi.weixin.qq.com/cgi-bin/batch/replaceuser?access_token=ACCESS_TOKEN
*/
func ReplaceUser(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiReplaceUser, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
全量覆盖部门



See: https://work.weixin.qq.com/api/doc/90000/90135/90982

POST https://qyapi.weixin.qq.com/cgi-bin/batch/replaceparty?access_token=ACCESS_TOKEN
*/
func ReplaceParty(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiReplaceParty, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取异步任务结果



See: https://work.weixin.qq.com/api/doc/90000/90135/90983

GET https://qyapi.weixin.qq.com/cgi-bin/batch/getresult?access_token=ACCESS_TOKEN&jobid=JOBID
*/
func GetResult(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetResult + "?" + params.Encode())
}
