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

// Package oa OA
package oa

import (
	"bytes"

	"github.com/fastwego/wechat4work/corporation"
)

const (
	apiGetCheckinData      = "/cgi-bin/checkin/getcheckindata"
	apiGetCheckinPption    = "/cgi-bin/checkin/getcheckinoption"
	apiGetTemplateDetail   = "/cgi-bin/oa/gettemplatedetail"
	apiApplyEvent          = "/cgi-bin/oa/applyevent"
	apiGetApprovalInfo     = "/cgi-bin/oa/getapprovalinfo"
	apiGetApprovalDetail   = "/cgi-bin/oa/getapprovaldetail"
	apiGetApprovalData     = "/cgi-bin/corp/getapprovaldata"
	apiGetDialRecord       = "/cgi-bin/dial/get_dial_record"
	apiGetOpenApprovalData = "/cgi-bin/corp/getopenapprovaldata"
)

/*
获取打卡数据



See: https://work.weixin.qq.com/api/doc/90000/90135/90262

POST https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckindata?access_token=ACCESS_TOKEN
*/
func GetCheckinData(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCheckinData, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取打卡规则



See: https://work.weixin.qq.com/api/doc/90000/90135/90263

POST https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckinoption?access_token=ACCESS_TOKEN
*/
func GetCheckinPption(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCheckinPption, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取审批模板详情

企业可通过审批应用或自建应用Secret调用本接口，获取企业微信“审批应用”内指定审批模板的详情。

See: https://work.weixin.qq.com/api/doc/90000/90135/91982

POST https://qyapi.weixin.qq.com/cgi-bin/oa/gettemplatedetail?access_token=ACCESS_TOKEN
*/
func GetTemplateDetail(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetTemplateDetail, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
提交审批申请

企业可通过审批应用或自建应用Secret调用本接口，代应用可见范围内员工在企业微信“审批应用”内提交指定类型的审批申请。

See: https://work.weixin.qq.com/api/doc/90000/90135/91853

POST https://qyapi.weixin.qq.com/cgi-bin/oa/applyevent?access_token=ACCESS_TOKEN
*/
func ApplyEvent(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiApplyEvent, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量获取审批单号



See: https://work.weixin.qq.com/api/doc/90000/90135/91816

POST https://qyapi.weixin.qq.com/cgi-bin/oa/getapprovalinfo?access_token=ACCESS_TOKEN
*/
func GetApprovalInfo(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetApprovalInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取审批申请详情

企业可通过审批应用或自建应用Secret调用本接口，根据审批单号查询企业微信“审批应用”的审批申请详情。

See: https://work.weixin.qq.com/api/doc/90000/90135/91983

POST https://qyapi.weixin.qq.com/cgi-bin/oa/getapprovaldetail?access_token=ACCESS_TOKEN
*/
func GetApprovalDetail(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetApprovalDetail, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取审批数据（旧）

通过本接口来获取公司一段时间内的审批记录。一次拉取调用最多拉取100个审批记录，可以通过多次拉取的方式来满足需求，但调用频率不可超过600次/分。

See: https://work.weixin.qq.com/api/doc/90000/90135/91530

POST https://qyapi.weixin.qq.com/cgi-bin/corp/getapprovaldata?access_token=ACCESS_TOKEN
*/
func GetApprovalData(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetApprovalData, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取公费电话拨打记录

企业可通过此接口，按时间范围拉取成功接通的公费电话拨打记录。

See: https://work.weixin.qq.com/api/doc/90000/90135/90267

POST https://qyapi.weixin.qq.com/cgi-bin/dial/get_dial_record?access_token=ACCESS_TOKEN
*/
func GetDialRecord(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetDialRecord, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询自建应用审批单当前状态

开发者也可主动查询审批单的当前审批状态。

See: https://work.weixin.qq.com/api/doc/90000/90135/90269

POST https://qyapi.weixin.qq.com/cgi-bin/corp/getopenapprovaldata?access_token=ACCESS_TOKEN
*/
func GetOpenApprovalData(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetOpenApprovalData, bytes.NewReader(payload), "application/json;charset=utf-8")
}
