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

// Package meeting_room OA/会议室
package meeting_room

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiAdd            = "/cgi-bin/oa/meetingroom/add"
	apiList           = "/cgi-bin/oa/meetingroom/list"
	apiEdit           = "/cgi-bin/oa/meetingroom/edit"
	apiDel            = "/cgi-bin/oa/meetingroom/del"
	apiGetBookingInfo = "/cgi-bin/oa/meetingroom/get_booking_info"
	apiBook           = "/cgi-bin/oa/meetingroom/book"
	apiCancelBook     = "/cgi-bin/oa/meetingroom/cancel_book"
)

/*
添加会议室

企业可通过此接口添加一个会议室。

See: https://work.weixin.qq.com/api/doc/90000/90135/92793

POST https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/add?access_token=ACCESS_TOKEN
*/
func Add(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询会议室

企业可通过此接口查询满足条件的会议室。

See: https://work.weixin.qq.com/api/doc/90000/90135/92793

POST https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/list?access_token=ACCESS_TOKEN
*/
func List(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
编辑会议室

企业可通过此接口编辑相关会议室的基本信息。

See: https://work.weixin.qq.com/api/doc/90000/90135/92793

POST https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/edit?access_token=ACCESS_TOKEN
*/
func Edit(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiEdit, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除会议室

企业可通过此接口删除指定的会议室。

See: https://work.weixin.qq.com/api/doc/90000/90135/92793

POST https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/del?access_token=ACCESS_TOKEN
*/
func Del(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDel, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询会议室的预定信息

企业可通过此接口查询相关会议室在指定时间段的预定情况，如是否已被预定，预定者的userid等信息，不支持跨天查询。

See: https://work.weixin.qq.com/api/doc/90000/90135/92794

POST https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/get_booking_info?access_token=ACCESS_TOKEN
*/
func GetBookingInfo(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetBookingInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
预定会议室

企业可通过此接口预定会议室并自动关联日程。

See: https://work.weixin.qq.com/api/doc/90000/90135/92794

POST https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/book?access_token=ACCESS_TOKEN
*/
func Book(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBook, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
取消预定会议室

企业可通过此接口取消会议室的预定

See: https://work.weixin.qq.com/api/doc/90000/90135/92794

POST https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/cancel_book?access_token=ACCESS_TOKEN
*/
func CancelBook(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCancelBook, bytes.NewReader(payload), "application/json;charset=utf-8")
}
