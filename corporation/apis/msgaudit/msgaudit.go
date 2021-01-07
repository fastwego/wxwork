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

// Package msgaudit 会话内容存档
package msgaudit

import (
	"bytes"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetRobotInfo      = "/cgi-bin/msgaudit/get_robot_info"
	apiGetPermitUserList = "/cgi-bin/msgaudit/get_permit_user_list"
	apiCheckSingleAgree  = "/cgi-bin/msgaudit/check_single_agree"
	apiGroupchatGet      = "/cgi-bin/msgaudit/groupchat/get"
)

/*
获取机器人信息

通过robot_id获取机器人的名称和创建者

See: https://work.weixin.qq.com/api/doc/90000/90135/91774

GET https://qyapi.weixin.qq.com/cgi-bin/msgaudit/get_robot_info?access_token=ACCESS_TOKEN&robot_id=ROBOT_ID
*/
func GetRobotInfo(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetRobotInfo + "?" + params.Encode())
}

/*
获取会话内容存档开启成员列表

企业可通过此接口，获取企业开启会话内容存档的成员列表

See: https://work.weixin.qq.com/api/doc/90000/90135/91614

GET https://qyapi.weixin.qq.com/cgi-bin/msgaudit/get_permit_user_list?access_token=ACCESS_TOKEN
*/
func GetPermitUserList(ctx *corporation.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetPermitUserList)
}

/*
获取会话同意情况

企业可通过下述接口，获取会话中外部成员的同意情况

See: https://work.weixin.qq.com/api/doc/90000/90135/91782

POST https://qyapi.weixin.qq.com/cgi-bin/msgaudit/check_single_agree?access_token=ACCESS_TOKEN
*/
func CheckSingleAgree(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCheckSingleAgree, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取会话内容存档内部群信息

企业可通过此接口，获取会话内容存档本企业的内部群信息，包括群名称、群主id、公告、群创建时间以及所有群成员的id与加入时间。

See: https://work.weixin.qq.com/api/doc/90000/90135/92951

POST https://qyapi.weixin.qq.com/cgi-bin/msgaudit/groupchat/get?access_token=ACCESS_TOKEN
*/
func GroupchatGet(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGroupchatGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}
