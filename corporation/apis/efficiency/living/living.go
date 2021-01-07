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

// Package living 效率工具/直播
package living

import (
	"bytes"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiCreate             = "/cgi-bin/living/create"
	apiModify             = "/cgi-bin/living/modify"
	apiCancel             = "/cgi-bin/living/cancel"
	apiDeleteReplayData   = "/cgi-bin/living/delete_replay_data"
	apiGetLivingCode      = "/cgi-bin/living/get_living_code"
	apiGetUserAllLivingId = "/cgi-bin/living/get_user_all_livingid"
	apiGetLivingInfo      = "/cgi-bin/living/get_living_info"
	apiGetWatchStat       = "/cgi-bin/living/get_watch_stat"
)

/*
创建预约直播



See: https://work.weixin.qq.com/api/doc/90000/90135/93637

POST https://qyapi.weixin.qq.com/cgi-bin/living/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改预约直播



See: https://work.weixin.qq.com/api/doc/90000/90135/93640

POST https://qyapi.weixin.qq.com/cgi-bin/living/modify?access_token=ACCESS_TOKEN
*/
func Modify(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiModify, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
取消预约直播



See: https://work.weixin.qq.com/api/doc/90000/90135/93638

POST https://qyapi.weixin.qq.com/cgi-bin/living/cancel?access_token=ACCESS_TOKEN
*/
func Cancel(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCancel, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除直播回放



See: https://work.weixin.qq.com/api/doc/90000/90135/93874

POST https://qyapi.weixin.qq.com/cgi-bin/living/delete_replay_data?access_token=ACCESS_TOKEN
*/
func DeleteReplayData(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDeleteReplayData, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
在微信中观看直播或直播回放



See: https://work.weixin.qq.com/api/doc/90000/90135/93641

POST https://qyapi.weixin.qq.com/cgi-bin/living/get_living_code?access_token=ACCESS_TOKEN
*/
func GetLivingCode(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetLivingCode, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取成员直播ID列表

通过此接口可以获取指定成员的所有直播ID

See: https://work.weixin.qq.com/api/doc/90000/90135/93634

POST https://qyapi.weixin.qq.com/cgi-bin/living/get_user_all_livingid?access_token=ACCESS_TOKEN
*/
func GetUserAllLivingId(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUserAllLivingId, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取直播详情



See: https://work.weixin.qq.com/api/doc/90000/90135/93635

GET https://qyapi.weixin.qq.com/cgi-bin/living/get_living_info?access_token=ACCESS_TOKEN&livingid=LIVINGID
*/
func GetLivingInfo(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetLivingInfo + "?" + params.Encode())
}

/*
获取直播观看明细

通过该接口可以获取所有观看直播的人员统计

See: https://work.weixin.qq.com/api/doc/90000/90135/93636

POST https://qyapi.weixin.qq.com/cgi-bin/living/get_watch_stat?access_token=ACCESS_TOKEN
*/
func GetWatchStat(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetWatchStat, bytes.NewReader(payload), "application/json;charset=utf-8")
}
