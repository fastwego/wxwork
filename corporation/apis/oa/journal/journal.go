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

// Package journal OA/汇报
package journal

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetRecordList   = "/cgi-bin/oa/journal/get_record_list"
	apiGetRecordDetail = "/cgi-bin/oa/journal/get_record_detail"
	apiGetStatList     = "/cgi-bin/oa/journal/get_stat_list"
)

/*
批量获取汇报记录单号

表单ID的获取方式：管理后台—汇报应用—某个汇报的内容设置页—点击“汇报名称”，即可获取

See: https://work.weixin.qq.com/api/doc/90000/90135/93393

POST https://qyapi.weixin.qq.com/cgi-bin/oa/journal/get_record_list?access_token=ACCESS_TOKEN
*/
func GetRecordList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetRecordList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取汇报记录详情

企业可通过access_token调用本接口，根据汇报记录单号查询企业微信“汇报应用”的汇报详情。

See: https://work.weixin.qq.com/api/doc/90000/90135/93394

POST https://qyapi.weixin.qq.com/cgi-bin/oa/journal/get_record_detail?access_token=ACCESS_TOKEN
*/
func GetRecordDetail(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetRecordDetail, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取汇报统计数据

企业可通过access_token调用本接口，根据汇报表单id查询企业微信“汇报应用”的汇报统计汇总信息。该接口只能拉取到已经汇总的统计数据，对于尚未完成汇总的周期不会返回。

See: https://work.weixin.qq.com/api/doc/90000/90135/93395

POST https://qyapi.weixin.qq.com/cgi-bin/oa/journal/get_stat_list?access_token=ACCESS_TOKEN
*/
func GetStatList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetStatList, bytes.NewReader(payload), "application/json;charset=utf-8")
}
