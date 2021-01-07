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

// Package payment 企业支付
package payment

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiAddMerchant    = "/cgi-bin/externalpay/addmerchant"
	apiGetMerchant    = "/cgi-bin/externalpay/getmerchant"
	apiDelMerchant    = "/cgi-bin/externalpay/delmerchant"
	apiSetMchUseScope = "/cgi-bin/externalpay/set_mch_use_scope"
	apiGetBillList    = "/cgi-bin/externalpay/get_bill_list"
)

/*
新增商户号

使用该接口时需要开通对外收款能力并绑定第一个商户号。（说明：第一个商户号的绑定需要由企业管理员手动绑定，绑定入口在：手机端工作台-对外收款，或管理后台-应用管理-对外收款-启用）通过该接口可以导入其他商户号，同一个企业最多绑定200个商户号。

See: https://work.weixin.qq.com/api/doc/90000/90135/93666

POST https://qyapi.weixin.qq.com/cgi-bin/externalpay/addmerchant?access_token=ACCESS_TOKEN
*/
func AddMerchant(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddMerchant, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询商户号详情

通过该接口可以查询商户号信息以及商户号使用范围

See: https://work.weixin.qq.com/api/doc/90000/90135/93666

POST https://qyapi.weixin.qq.com/cgi-bin/externalpay/getmerchant?access_token=ACCESS_TOKEN
*/
func GetMerchant(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetMerchant, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除商户号

通过该接口可以删除未绑定的商户号

See: https://work.weixin.qq.com/api/doc/90000/90135/93666

POST https://qyapi.weixin.qq.com/cgi-bin/externalpay/delmerchant?access_token=ACCESS_TOKEN
*/
func DelMerchant(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelMerchant, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
设置商户号使用范围

通过该接口可以设置商户号的使用范围

See: https://work.weixin.qq.com/api/doc/90000/90135/93666

POST https://qyapi.weixin.qq.com/cgi-bin/externalpay/set_mch_use_scope?access_token=ACCESS_TOKEN
*/
func SetMchUseScope(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetMchUseScope, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取对外收款记录

企业和服务商可通过此接口获取企业的对外收款记录。

See: https://work.weixin.qq.com/api/doc/90000/90135/93667

POST https://qyapi.weixin.qq.com/cgi-bin/externalpay/get_bill_list?access_token=ACCESS_TOKEN
*/
func GetBillList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetBillList, bytes.NewReader(payload), "application/json;charset=utf-8")
}
