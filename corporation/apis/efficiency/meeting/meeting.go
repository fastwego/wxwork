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

// Package meeting 效率工具/会议
package meeting

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiCreate           = "/cgi-bin/meeting/create"
	apiUpdate           = "/cgi-bin/meeting/update"
	apiCancel           = "/cgi-bin/meeting/cancel"
	apiGetUserMeetingId = "/cgi-bin/meeting/get_user_meetingid"
	apiGetInfo          = "/cgi-bin/meeting/get_info"
)

/*
创建预约会议



See: https://work.weixin.qq.com/api/doc/90000/90135/93627

POST https://qyapi.weixin.qq.com/cgi-bin/meeting/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改预约会议



See: https://work.weixin.qq.com/api/doc/90000/90135/93631

POST https://qyapi.weixin.qq.com/cgi-bin/meeting/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
取消预约会议



See: https://work.weixin.qq.com/api/doc/90000/90135/93630

POST https://qyapi.weixin.qq.com/cgi-bin/meeting/cancel?access_token=ACCESS_TOKEN
*/
func Cancel(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCancel, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取成员会议ID列表

通过此接口可以获取指定成员指定时间内的所有会议ID

See: https://work.weixin.qq.com/api/doc/90000/90135/93628

POST https://qyapi.weixin.qq.com/cgi-bin/meeting/get_user_meetingid?access_token=ACCESS_TOKEN
*/
func GetUserMeetingId(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUserMeetingId, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取会议详情



See: https://work.weixin.qq.com/api/doc/90000/90135/93629

POST https://qyapi.weixin.qq.com/cgi-bin/meeting/get_info?access_token=ACCESS_TOKEN
*/
func GetInfo(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}
