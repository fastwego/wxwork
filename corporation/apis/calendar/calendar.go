// Copyright 2020 FastWeGo
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

// Package calendar 日程
package calendar

import (
	"bytes"

	"github.com/fastwego/wechat4work/corporation"
)

const (
	apiCalendarAdd           = "/cgi-bin/oa/calendar/add"
	apiCalendarUpdate        = "/cgi-bin/oa/calendar/update"
	apiCalendarGet           = "/cgi-bin/oa/calendar/get"
	apiCalendarDel           = "/cgi-bin/oa/calendar/del"
	apiScheduleAdd           = "/cgi-bin/oa/schedule/add"
	apiScheduleUpdate        = "/cgi-bin/oa/schedule/update"
	apiScheduleGet           = "/cgi-bin/oa/schedule/get"
	apiScheduleDel           = "/cgi-bin/oa/schedule/del"
	apiScheduleGetByCalendar = "/cgi-bin/oa/schedule/get_by_calendar"
)

/*
创建日历



See: https://work.weixin.qq.com/api/doc/90000/90135/92618

POST https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/add?access_token=ACCESS_TOKEN
*/
func CalendarAdd(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCalendarAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新日历



See: https://work.weixin.qq.com/api/doc/90000/90135/92619

POST https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/update?access_token=ACCESS_TOKEN
*/
func CalendarUpdate(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCalendarUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取日历



See: https://work.weixin.qq.com/api/doc/90000/90135/92621

POST https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/get?access_token=ACCESS_TOKEN
*/
func CalendarGet(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCalendarGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除日历



See: https://work.weixin.qq.com/api/doc/90000/90135/92620

POST https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/del?access_token=ACCESS_TOKEN
*/
func CalendarDel(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCalendarDel, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
创建日程

该接口用于在日历中创建一个日程。

See: https://work.weixin.qq.com/api/doc/90000/90135/92622

POST https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/add?access_token=ACCESS_TOKEN
*/
func ScheduleAdd(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiScheduleAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新日程



See: https://work.weixin.qq.com/api/doc/90000/90135/92623

POST https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/update?access_token=ACCESS_TOKEN
*/
func ScheduleUpdate(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiScheduleUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取日程

该接口用于获取指定的日程详情。

See: https://work.weixin.qq.com/api/doc/90000/90135/92624

POST https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/get?access_token=ACCESS_TOKEN
*/
func ScheduleGet(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiScheduleGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
取消日程

该接口用于取消指定的日程。

See: https://work.weixin.qq.com/api/doc/90000/90135/92625

POST https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/del?access_token=ACCESS_TOKEN
*/
func ScheduleDel(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiScheduleDel, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取日历下的日程列表

该接口用于获取指定的日历下的日程列表。

See: https://work.weixin.qq.com/api/doc/90000/90135/92626

POST https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/get_by_calendar?access_token=ACCESS_TOKEN
*/
func ScheduleGetByCalendar(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiScheduleGetByCalendar, bytes.NewReader(payload), "application/json;charset=utf-8")
}
