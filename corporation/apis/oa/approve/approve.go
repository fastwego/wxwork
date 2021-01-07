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

// Package approve OA/审批
package approve

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetTemplateDetail    = "/cgi-bin/oa/gettemplatedetail"
	apiApplyEvent           = "/cgi-bin/oa/applyevent"
	apiGetApprovalInfo      = "/cgi-bin/oa/getapprovalinfo"
	apiGetApprovalDetail    = "/cgi-bin/oa/getapprovaldetail"
	apiGetApprovalData      = "/cgi-bin/corp/getapprovaldata"
	apiGetCorpConf          = "/cgi-bin/oa/vacation/getcorpconf"
	apiGetUserVacationQuota = "/cgi-bin/oa/vacation/getuservacationquota"
	apiSetOneUserQuota      = "/cgi-bin/oa/vacation/setoneuserquota"
)

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
获取企业假期管理配置

企业可通过审批应用或自建应用Secret调用本接口，获取可见范围内员工的“假期管理”配置，包括：各个假期的id、名称、请假单位、时长计算方式、发放规则等。

See: https://work.weixin.qq.com/api/doc/90000/90135/93375

GET https://qyapi.weixin.qq.com/cgi-bin/oa/vacation/getcorpconf?access_token=ACCESS_TOKEN
*/
func GetCorpConf(ctx *corporation.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetCorpConf)
}

/*
获取成员假期余额

企业可通过审批应用或自建应用Secret调用本接口，获取可见范围内各个员工的假期余额数据。

See: https://work.weixin.qq.com/api/doc/90000/90135/93376

POST https://qyapi.weixin.qq.com/cgi-bin/oa/vacation/getuservacationquota?access_token=ACCESS_TOKEN
*/
func GetUserVacationQuota(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUserVacationQuota, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改成员假期余额

企业可通过审批应用或自建应用Secret调用本接口，修改可见范围内员工的“假期余额”。

See: https://work.weixin.qq.com/api/doc/90000/90135/93377

POST https://qyapi.weixin.qq.com/cgi-bin/oa/vacation/setoneuserquota?access_token=ACCESS_TOKEN
*/
func SetOneUserQuota(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetOneUserQuota, bytes.NewReader(payload), "application/json;charset=utf-8")
}
