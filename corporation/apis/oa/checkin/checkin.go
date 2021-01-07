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

// Package checkin OA/打卡
package checkin

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetCorpCheckinOption   = "/cgi-bin/checkin/getcorpcheckinoption"
	apiGetCheckinOption       = "/cgi-bin/checkin/getcheckinoption"
	apiGetCheckinData         = "/cgi-bin/checkin/getcheckindata"
	apiGetCheckinDayData      = "/cgi-bin/checkin/getcheckin_daydata"
	apiGetCheckinMonthData    = "/cgi-bin/checkin/getcheckin_monthdata"
	apiGetCheckinScheduleList = "/cgi-bin/checkin/getcheckinschedulist"
	apiSetCheckinScheduleList = "/cgi-bin/checkin/setcheckinschedulist"
	apiAddCheckinUserFace     = "/cgi-bin/checkin/addcheckinuserface"
)

/*
获取企业所有打卡规则

企业可通过打卡应用Secret调用本接口，获取企业内所有打卡规则数据。

See: https://work.weixin.qq.com/api/doc/90000/90135/93384

POST https://qyapi.weixin.qq.com/cgi-bin/checkin/getcorpcheckinoption?access_token=ACCESS_TOKEN
*/
func GetCorpCheckinOption(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCorpCheckinOption, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取员工打卡规则

企业可通过打卡应用Secret调用本接口，获取指定员工指定日期的打卡规则。

See: https://work.weixin.qq.com/api/doc/90000/90135/90263

POST https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckinoption?access_token=ACCESS_TOKEN
*/
func GetCheckinOption(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCheckinOption, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取打卡记录数据

企业可通过打卡应用Secret调用本接口，获取指定员工指定时间段内的打卡记录数据。

See: https://work.weixin.qq.com/api/doc/90000/90135/90262

POST https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckindata?access_token=ACCESS_TOKEN
*/
func GetCheckinData(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCheckinData, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取打卡日报数据

企业可通过打卡应用Secret调用本接口，获取指定员工指定时间段内的打卡日报统计数据。

See: https://work.weixin.qq.com/api/doc/90000/90135/93374

POST https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckin_daydata?access_token=ACCESS_TOKEN
*/
func GetCheckinDayData(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCheckinDayData, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取打卡月报数据

企业可通过打卡应用Secret调用本接口，获取指定员工指定时间段内的打卡月报统计数据。

See: https://work.weixin.qq.com/api/doc/90000/90135/93387

POST https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckin_monthdata?access_token=ACCESS_TOKEN
*/
func GetCheckinMonthData(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCheckinMonthData, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取打卡人员排班信息

企业可通过打卡应用Secret调用本接口，获取打卡规则为“按班次上下班”规则的指定员工指定时间段内的排班信息。

See: https://work.weixin.qq.com/api/doc/90000/90135/93380

POST https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckinschedulist?access_token=ACCESS_TOKEN
*/
func GetCheckinScheduleList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCheckinScheduleList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
为打卡人员排班

企业可通过打卡应用Secret调用本接口，为打卡规则为“按班次上下班”规则的指定员工排班。

See: https://work.weixin.qq.com/api/doc/90000/90135/93385

POST https://qyapi.weixin.qq.com/cgi-bin/checkin/setcheckinschedulist?access_token=ACCESS_TOKEN
*/
func SetCheckinScheduleList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetCheckinScheduleList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
录入打卡人员人脸信息

企业可通过打卡应用Secret调用本接口，为企业打卡人员录入人脸信息，人脸信息仅用于人脸打卡。

See: https://work.weixin.qq.com/api/doc/90000/90135/93378

POST https://qyapi.weixin.qq.com/cgi-bin/checkin/addcheckinuserface?access_token=ACCESS_TOKEN
*/
func AddCheckinUserFace(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddCheckinUserFace, bytes.NewReader(payload), "application/json;charset=utf-8")
}
