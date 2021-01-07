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

// Package school 家校沟通
package school

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetSubscribeQrCode = "/cgi-bin/externalcontact/get_subscribe_qr_code"
	apiSetSubscribeMode   = "/cgi-bin/externalcontact/set_subscribe_mode"
	apiGetSubscribeMode   = "/cgi-bin/externalcontact/get_subscribe_mode"
	apiSend               = "/cgi-bin/externalcontact/message/send"
	apiConvertToOpenid    = "/cgi-bin/externalcontact/convert_to_openid"
)

/*
获取「学校通知」二维码



See: https://work.weixin.qq.com/api/doc/90000/90135/92320

GET https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_subscribe_qr_code?access_token=ACCESS_TOKEN
*/
func GetSubscribeQrCode(ctx *corporation.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetSubscribeQrCode)
}

/*
设置关注「学校通知」的模式

可通过此接口修改家长关注「学校通知」的模式：“可扫码填写资料加入”或“禁止扫码填写资料加入”

See: https://work.weixin.qq.com/api/doc/90000/90135/92318

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/set_subscribe_mode?access_token=ACCESS_TOKEN
*/
func SetSubscribeMode(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetSubscribeMode, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取关注「学校通知」的模式

可通过此接口获取家长关注「学校通知」的模式：“可扫码填写资料加入”或“禁止扫码填写资料加入”

See: https://work.weixin.qq.com/api/doc/90000/90135/92318

GET https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_subscribe_mode?access_token=ACCESS_TOKEN
*/
func GetSubscribeMode(ctx *corporation.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetSubscribeMode)
}

/*
发送「学校通知」



See: https://work.weixin.qq.com/api/doc/90000/90135/92321

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/message/send?access_token=ACCESS_TOKEN
*/
func Send(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSend, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
外部联系人openid转换

企业和服务商可通过此接口，将微信外部联系人的userid（如何获取?）转为微信openid，用于调用支付相关接口。暂不支持企业微信外部联系人（ExternalUserid为wo开头）的userid转openid。

See: https://work.weixin.qq.com/api/doc/90000/90135/92323

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/convert_to_openid?access_token=ACCESS_TOKEN
*/
func ConvertToOpenid(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiConvertToOpenid, bytes.NewReader(payload), "application/json;charset=utf-8")
}
