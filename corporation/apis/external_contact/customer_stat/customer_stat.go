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

// Package customer_stat 客户联系/统计管理
package customer_stat

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetUserBehaviorData = "/cgi-bin/externalcontact/get_user_behavior_data"
	apiStatistic           = "/cgi-bin/externalcontact/groupchat/statistic"
	apiStatisticGroupByDay = "/cgi-bin/externalcontact/groupchat/statistic_group_by_day"
)

/*
获取「联系客户统计」数据

企业可通过此接口获取成员联系客户的数据，包括发起申请数、新增客户数、聊天数、发送消息数和删除/拉黑成员的客户数等指标。

See: https://work.weixin.qq.com/api/doc/90000/90135/92132

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_user_behavior_data?access_token=ACCESS_TOKEN
*/
func GetUserBehaviorData(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUserBehaviorData, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取「群聊数据统计」数据/按群主聚合的方式



See: https://work.weixin.qq.com/api/doc/90000/90135/92133

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic?access_token=ACCESS_TOKEN
*/
func Statistic(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiStatistic, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取「群聊数据统计」数据/按自然日聚合的方式



See: https://work.weixin.qq.com/api/doc/90000/90135/92133

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic_group_by_day?access_token=ACCESS_TOKEN
*/
func StatisticGroupByDay(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiStatisticGroupByDay, bytes.NewReader(payload), "application/json;charset=utf-8")
}
