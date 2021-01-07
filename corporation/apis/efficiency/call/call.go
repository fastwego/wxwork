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

// Package call 效率工具/公费电话
package call

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetDialRecord = "/cgi-bin/dial/get_dial_record"
)

/*
获取公费电话拨打记录

企业可通过此接口，按时间范围拉取成功接通的公费电话拨打记录。

See: https://work.weixin.qq.com/api/doc/90000/90135/93662

POST https://qyapi.weixin.qq.com/cgi-bin/dial/get_dial_record?access_token=ACCESS_TOKEN
*/
func GetDialRecord(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetDialRecord, bytes.NewReader(payload), "application/json;charset=utf-8")
}
