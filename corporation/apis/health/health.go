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

// Package health 健康上报
package health

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetHealthReportStat = "/cgi-bin/health/get_health_report_stat"
	apiGetReportJobids     = "/cgi-bin/health/get_report_jobids"
	apiGetReportJobInfo    = "/cgi-bin/health/get_report_job_info"
	apiGetReportAnswer     = "/cgi-bin/health/get_report_answer"
)

/*
获取健康上报使用统计



See: https://work.weixin.qq.com/api/doc/90000/90135/92747

POST https://qyapi.weixin.qq.com/cgi-bin/health/get_health_report_stat?access_token=ACCESS_TOKEN
*/
func GetHealthReportStat(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetHealthReportStat, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取健康上报任务ID列表

通过此接口可以获取企业当前正在运行的上报任务ID列表。

See: https://work.weixin.qq.com/api/doc/90000/90135/92749

POST https://qyapi.weixin.qq.com/cgi-bin/health/get_report_jobids?access_token=ACCESS_TOKEN
*/
func GetReportJobids(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetReportJobids, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取健康上报任务详情

通过此接口可以获取指定的健康上报任务详情。

See: https://work.weixin.qq.com/api/doc/90000/90135/92751

POST https://qyapi.weixin.qq.com/cgi-bin/health/get_report_job_info?access_token=ACCESS_TOKEN
*/
func GetReportJobInfo(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetReportJobInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户填写答案

通过此接口可以获取指定的健康上报任务详情。

See: https://work.weixin.qq.com/api/doc/90000/90135/92754

POST https://qyapi.weixin.qq.com/cgi-bin/health/get_report_answer?access_token=ACCESS_TOKEN
*/
func GetReportAnswer(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetReportAnswer, bytes.NewReader(payload), "application/json;charset=utf-8")
}
