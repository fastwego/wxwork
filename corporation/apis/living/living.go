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

// Package living 企业直播
package living

import (
	"bytes"
	"net/url"

	"github.com/fastwego/wechat4work/corporation"
)

const (
	apiGetUserLivingid = "/cgi-bin/living/get_user_livingid"
	apiGetLivingInfo   = "/cgi-bin/living/get_living_info"
	apiGetWatchStat    = "/cgi-bin/living/get_watch_stat"
)

/*
获取成员直播ID列表

通过此接口可以获取指定成员指定时间内的所有直播ID

See: https://work.weixin.qq.com/api/doc/90000/90135/92735

POST https://qyapi.weixin.qq.com/cgi-bin/living/get_user_livingid?access_token=ACCESS_TOKEN
*/
func GetUserLivingid(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUserLivingid, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取直播详情



See: https://work.weixin.qq.com/api/doc/90000/90135/92734

GET https://qyapi.weixin.qq.com/cgi-bin/living/get_living_info?access_token=ACCESS_TOKEN&livingid=LIVINGID
*/
func GetLivingInfo(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetLivingInfo + "?" + params.Encode())
}

/*
获取看直播统计

通过该接口可以获取所有观看直播的人员统计

See: https://work.weixin.qq.com/api/doc/90000/90135/92736

POST https://qyapi.weixin.qq.com/cgi-bin/living/get_watch_stat?access_token=ACCESS_TOKEN
*/
func GetWatchStat(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetWatchStat, bytes.NewReader(payload), "application/json;charset=utf-8")
}
