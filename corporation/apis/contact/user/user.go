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

// Package user 通讯录管理/成员管理
package user

import (
	"bytes"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiCreate          = "/cgi-bin/user/create"
	apiGet             = "/cgi-bin/user/get"
	apiUpdate          = "/cgi-bin/user/update"
	apiDelete          = "/cgi-bin/user/delete"
	apiBatchDelete     = "/cgi-bin/user/batchdelete"
	apiSimpleList      = "/cgi-bin/user/simplelist"
	apiList            = "/cgi-bin/user/list"
	apiConvertToOpenId = "/cgi-bin/user/convert_to_openid"
	apiConvertToUserId = "/cgi-bin/user/convert_to_userid"
	apiAuthSucc        = "/cgi-bin/user/authsucc"
	apiInvite          = "/cgi-bin/batch/invite"
	apiGetJoinQrcode   = "/cgi-bin/corp/get_join_qrcode"
	apiGetActiveStat   = "/cgi-bin/user/get_active_stat"
)

/*
创建成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90195

POST https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
读取成员

在通讯录同步助手中此接口可以读取企业通讯录的所有成员信息，而自建应用可以读取该应用设置的可见范围内的成员信息。

See: https://work.weixin.qq.com/api/doc/90000/90135/90196

GET https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&userid=USERID
*/
func Get(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGet + "?" + params.Encode())
}

/*
更新成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90197

POST https://qyapi.weixin.qq.com/cgi-bin/user/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90198

GET https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=ACCESS_TOKEN&userid=USERID
*/
func Delete(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiDelete + "?" + params.Encode())
}

/*
批量删除成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90199

POST https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete?access_token=ACCESS_TOKEN
*/
func BatchDelete(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取部门成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90200

GET https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=ACCESS_TOKEN&department_id=DEPARTMENT_ID&fetch_child=FETCH_CHILD
*/
func SimpleList(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiSimpleList + "?" + params.Encode())
}

/*
获取部门成员详情



See: https://work.weixin.qq.com/api/doc/90000/90135/90201

GET https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=ACCESS_TOKEN&department_id=DEPARTMENT_ID&fetch_child=FETCH_CHILD
*/
func List(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiList + "?" + params.Encode())
}

/*
userid与openid互换

注：需要成员使用微信登录企业微信或者关注微工作台（原企业号）才能转成openid;如果是外部联系人，请使用外部联系人openid转换转换openid

See: https://work.weixin.qq.com/api/doc/90000/90135/90202

POST https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_openid?access_token=ACCESS_TOKEN
*/
func ConvertToOpenId(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiConvertToOpenId, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
openid转userid

该接口主要应用于使用企业支付之后的结果查询。开发者需要知道某个结果事件的openid对应企业微信内成员的信息时，可以通过调用该接口进行转换查询。

See: https://work.weixin.qq.com/api/doc/90000/90135/90202

POST https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_userid?access_token=ACCESS_TOKEN
*/
func ConvertToUserId(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiConvertToUserId, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
二次验证

此接口可以满足安全性要求高的企业进行成员验证。开启二次验证后，当且仅当成员登录时，需跳转至企业自定义的页面进行验证。验证频率可在设置页面选择。开启二次验证方法如下图：企业在开启二次验证时，必须在管理端填写企业二次验证页面的url。当成员登录企业微信或关注微工作台（原企业号）进入企业时，会自动跳转到企业的验证页面。在跳转到企业的验证页面时，会带上如下参数：code=CODE。企业收到code后，使用“通讯录同步助手”调用接口“根据code获取成员信息”获取成员的userid。如果成员是首次加入企业，企业获取到userid，并验证了成员信息后，调用如下接口即可让成员成功加入企业。

See: https://work.weixin.qq.com/api/doc/90000/90135/90203

GET https://qyapi.weixin.qq.com/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN&userid=USERID
*/
func AuthSucc(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiAuthSucc + "?" + params.Encode())
}

/*
邀请成员

企业可通过接口批量邀请成员使用企业微信，邀请后将通过短信或邮件下发通知。

See: https://work.weixin.qq.com/api/doc/90000/90135/90975

POST https://qyapi.weixin.qq.com/cgi-bin/batch/invite?access_token=ACCESS_TOKEN
*/
func Invite(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiInvite, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取加入企业二维码

支持企业用户获取实时成员加入二维码。

See: https://work.weixin.qq.com/api/doc/90000/90135/91714

GET https://qyapi.weixin.qq.com/cgi-bin/corp/get_join_qrcode?access_token=ACCESS_TOKEN&size_type=SIZE_TYPE
*/
func GetJoinQrcode(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetJoinQrcode + "?" + params.Encode())
}

/*
获取企业活跃成员数



See: https://work.weixin.qq.com/api/doc/90000/90135/92714

POST https://qyapi.weixin.qq.com/cgi-bin/user/get_active_stat?access_token=ACCESS_TOKEN
*/
func GetActiveStat(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetActiveStat, bytes.NewReader(payload), "application/json;charset=utf-8")
}
