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

// Package corp_group 企业互联
package corp_group

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiListAppShareInfo = "/cgi-bin/corpgroup/corp/list_app_share_info"
	apiGetToken         = "/cgi-bin/corpgroup/corp/gettoken"
	apiTransferSession  = "/cgi-bin/miniprogram/transfer_session"
)

/*
获取应用共享信息

上级企业通过该接口获取某个应用分享给的所有企业列表。特别注意，对于有敏感权限的应用，需要下级企业确认后才能共享成功，若下级企业未确认，则不会存在于该接口的返回列表

See: https://work.weixin.qq.com/api/doc/90000/90135/93403

POST https://qyapi.weixin.qq.com/cgi-bin/corpgroup/corp/list_app_share_info?access_token=ACCESS_TOKEN
*/
func ListAppShareInfo(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListAppShareInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取下级企业的access_token

获取应用可见范围内下级企业的access_token，该access_token可用于调用下级企业通讯录的只读接口。

See: https://work.weixin.qq.com/api/doc/90000/90135/93359

POST https://qyapi.weixin.qq.com/cgi-bin/corpgroup/corp/gettoken?access_token=ACCESS_TOKEN
*/
func GetToken(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetToken, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取下级企业的小程序session

上级企业通过该接口转换为下级企业的小程序session

See: https://work.weixin.qq.com/api/doc/90000/90135/93355

POST https://qyapi.weixin.qq.com/cgi-bin/miniprogram/transfer_session?access_token=ACCESS_TOKEN
*/
func TransferSession(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTransferSession, bytes.NewReader(payload), "application/json;charset=utf-8")
}
