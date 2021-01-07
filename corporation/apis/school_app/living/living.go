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

// Package living 家校应用/上课直播
package living

import (
	"bytes"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetUserAllLivingId = "/cgi-bin/living/get_user_all_livingid"
	apiGetLivingInfo      = "/cgi-bin/school/living/get_living_info"
	apiGetWatchStat       = "/cgi-bin/school/living/get_watch_stat"
	apiGetUnwatchStat     = "/cgi-bin/school/living/get_unwatch_stat"
	apiDeleteReplayData   = "/cgi-bin/living/delete_replay_data"
)

/*
获取老师直播ID列表

通过此接口可以获取指定老师的所有直播ID

See: https://work.weixin.qq.com/api/doc/90000/90135/93739

POST https://qyapi.weixin.qq.com/cgi-bin/living/get_user_all_livingid?access_token=ACCESS_TOKEN
*/
func GetUserAllLivingId(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUserAllLivingId, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取直播详情



See: https://work.weixin.qq.com/api/doc/90000/90135/93740

GET https://qyapi.weixin.qq.com/cgi-bin/school/living/get_living_info?access_token=ACCESS_TOKEN&livingid=LIVINGID
*/
func GetLivingInfo(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetLivingInfo + "?" + params.Encode())
}

/*
获取观看直播统计

通过该接口可以获取所有观看直播的人员统计

See: https://work.weixin.qq.com/api/doc/90000/90135/93741

POST https://qyapi.weixin.qq.com/cgi-bin/school/living/get_watch_stat?access_token=ACCESS_TOKEN
*/
func GetWatchStat(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetWatchStat, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取未观看直播统计

通过该接口可以获取未观看直播的学生统计，学生的家长必须是已经关注「学校通知」才会纳入统计范围。

See: https://work.weixin.qq.com/api/doc/90000/90135/93742

POST https://qyapi.weixin.qq.com/cgi-bin/school/living/get_unwatch_stat?access_token=ACCESS_TOKEN
*/
func GetUnwatchStat(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUnwatchStat, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除直播回放



See: https://work.weixin.qq.com/api/doc/90000/90135/93743

POST https://qyapi.weixin.qq.com/cgi-bin/living/delete_replay_data?access_token=ACCESS_TOKEN
*/
func DeleteReplayData(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDeleteReplayData, bytes.NewReader(payload), "application/json;charset=utf-8")
}
