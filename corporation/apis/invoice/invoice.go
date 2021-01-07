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

// Package invoice 电子发票
package invoice

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetInvoiceInfo      = "/cgi-bin/card/invoice/reimburse/getinvoiceinfo"
	apiUpdateInvoiceStatus = "/cgi-bin/card/invoice/reimburse/updateinvoicestatus"
	apiUpdateStatusBatch   = "/cgi-bin/card/invoice/reimburse/updatestatusbatch"
	apiGetInvoiceInfoBatch = "/cgi-bin/card/invoice/reimburse/getinvoiceinfobatch"
)

/*
查询电子发票

接口说明：报销方在获得用户选择的电子发票标识参数后，可以通过该接口查询电子发票的结构化信息，并获取发票PDF文件。

See: https://work.weixin.qq.com/api/doc/90000/90135/90284

POST https://qyapi.weixin.qq.com/cgi-bin/card/invoice/reimburse/getinvoiceinfo?access_token=ACCESS_TOKEN
*/
func GetInvoiceInfo(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetInvoiceInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新发票状态

接口说明：报销企业和报销服务商可以通过该接口对某一张发票进行锁定、解锁和报销操作。各操作的业务含义及在用户端的表现如下：锁定：电子发票进入了企业的报销流程时应该执行锁定操作，执行锁定操作后的电子发票仍然会存在于用户卡包内，但无法重复提交报销。解锁：当电子发票由于各种原因，无法完成报销流程时，应执行解锁操作。执行锁定操作后的电子发票将恢复可以被提交的状态。报销：当电子发票报销完成后，应该使用本接口执行报销操作。执行报销操作后的电子发票将从用户的卡包中移除，用户可以在卡包的消息中查看到电子发票的核销信息。注意，报销为不可逆操作，请开发者慎重调用。

See: https://work.weixin.qq.com/api/doc/90000/90135/90285

POST https://qyapi.weixin.qq.com/cgi-bin/card/invoice/reimburse/updateinvoicestatus?access_token=ACCESS_TOKEN
*/
func UpdateInvoiceStatus(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateInvoiceStatus, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量更新发票状态

接口说明：发票平台可以通过该接口对某个成员的一批发票进行锁定、解锁和报销操作。注意，报销状态为不可逆状态，请开发者慎重调用。

See: https://work.weixin.qq.com/api/doc/90000/90135/90286

POST https://qyapi.weixin.qq.com/cgi-bin/card/invoice/reimburse/updatestatusbatch?access_token=ACCESS_TOKEN
*/
func UpdateStatusBatch(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateStatusBatch, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量查询电子发票

接口说明：报销方在获得用户选择的电子发票标识参数后，可以通过该接口批量查询电子发票的结构化信息。

See: https://work.weixin.qq.com/api/doc/90000/90135/90287

POST https://qyapi.weixin.qq.com/cgi-bin/card/invoice/reimburse/getinvoiceinfobatch?access_token=ACCESS_TOKEN
*/
func GetInvoiceInfoBatch(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetInvoiceInfoBatch, bytes.NewReader(payload), "application/json;charset=utf-8")
}
