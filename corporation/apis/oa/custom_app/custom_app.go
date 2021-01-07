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

// Package custom_app OA/自建应用
package custom_app

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetOpenApprovalData = "/cgi-bin/corp/getopenapprovaldata"
)

/*
查询自建应用审批单当前状态

开发者也可主动查询审批单的当前审批状态。

See: https://work.weixin.qq.com/api/doc/90000/90135/90269

POST https://qyapi.weixin.qq.com/cgi-bin/corp/getopenapprovaldata?access_token=ACCESS_TOKEN
*/
func GetOpenApprovalData(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetOpenApprovalData, bytes.NewReader(payload), "application/json;charset=utf-8")
}
