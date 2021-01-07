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

// Package customer_assign 客户联系/客户分配
package customer_assign

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetUnassignedList = "/cgi-bin/externalcontact/get_unassigned_list"
	apiTransfer          = "/cgi-bin/externalcontact/transfer"
	apiGetTransferResult = "/cgi-bin/externalcontact/get_transfer_result"
	apiGroupChatTransfer = "/cgi-bin/externalcontact/groupchat/transfer"
)

/*
获取离职成员列表

企业和第三方可通过此接口，获取所有离职成员的客户列表，并可进一步调用分配在职或离职成员的客户接口将这些客户重新分配给其他企业成员。

See: https://work.weixin.qq.com/api/doc/90000/90135/92124

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_unassigned_list?access_token=ACCESS_TOKEN
*/
func GetUnassignedList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUnassignedList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
分配在职或离职成员的客户

企业可通过此接口，转接在职成员的客户或分配离职成员的客户给其他成员。

See: https://work.weixin.qq.com/api/doc/90000/90135/92125

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer?access_token=ACCESS_TOKEN
*/
func Transfer(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTransfer, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询客户接替结果

企业和第三方可通过此接口查询客户的接替情况。

See: https://work.weixin.qq.com/api/doc/90000/90135/92973

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_transfer_result?access_token=ACCESS_TOKEN
*/
func GetTransferResult(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetTransferResult, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
分配离职成员的客户群

企业可通过此接口，将已离职成员为群主的群，分配给另一个客服成员。

See: https://work.weixin.qq.com/api/doc/90000/90135/92127

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/transfer?access_token=ACCESS_TOKEN
*/
func GroupChatTransfer(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGroupChatTransfer, bytes.NewReader(payload), "application/json;charset=utf-8")
}
