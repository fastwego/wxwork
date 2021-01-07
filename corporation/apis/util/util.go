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

// Package util 开发辅助
package util

import "github.com/fastwego/wxwork/corporation"

const (
	apiGetApiDomainIp = "/cgi-bin/get_api_domain_ip"
	apiGetCallbackIp  = "/cgi-bin/getcallbackip"
)

/*
获取企业微信API域名IP段

API域名IP即qyapi.weixin.qq.com的解析地址，由开发者调用企业微信侧的接入IP。如果企业需要做防火墙配置，那么可以通过这个接口获取到所有相关的IP段。IP段有变更可能，当IP段变更时，新旧IP段会同时保留一段时间。建议企业每天定时拉取IP段，更新防火墙设置，避免因IP段变更导致网络不通。

See: https://work.weixin.qq.com/api/doc/90000/90135/92520

GET https://qyapi.weixin.qq.com/cgi-bin/get_api_domain_ip?access_token=ACCESS_TOKEN
*/
func GetApiDomainIp(ctx *corporation.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetApiDomainIp)
}

/*
获取企业微信服务器的ip段

企业微信在回调企业指定的URL时，是通过特定的IP发送出去的。如果企业需要做防火墙配置，那么可以通过这个接口获取到所有相关的IP段。

See: https://work.weixin.qq.com/api/doc/90000/90135/90238

GET https://qyapi.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN
*/
func GetCallbackIp(ctx *corporation.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetCallbackIp)
}
