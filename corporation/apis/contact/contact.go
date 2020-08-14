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

// Package contact 通讯录管理
package contact

import (
	"bytes"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiUserCreate       = "/cgi-bin/user/create"
	apiUserGet          = "/cgi-bin/user/get"
	apiUserUpdate       = "/cgi-bin/user/update"
	apiUserDelete       = "/cgi-bin/user/delete"
	apiUserBatchDelete  = "/cgi-bin/user/batchdelete"
	apiUserSimpleList   = "/cgi-bin/user/simplelist"
	apiUserList         = "/cgi-bin/user/list"
	apiConvertToOpenid  = "/cgi-bin/user/convert_to_openid"
	apiConvertToUserid  = "/cgi-bin/user/convert_to_userid"
	apiAuthSucc         = "/cgi-bin/user/authsucc"
	apiBatchInvite      = "/cgi-bin/batch/invite"
	apiGetJoinQrcode    = "/cgi-bin/corp/get_join_qrcode"
	apiGetActiveStat    = "/cgi-bin/user/get_active_stat"
	apiDepartmentCreate = "/cgi-bin/department/create"
	apiDepartmentUpdate = "/cgi-bin/department/update"
	apiDepartmentDelete = "/cgi-bin/department/delete"
	apiDepartmentList   = "/cgi-bin/department/list"
	apiTagCreate        = "/cgi-bin/tag/create"
	apiTagUpdate        = "/cgi-bin/tag/update"
	apiTagDelete        = "/cgi-bin/tag/delete"
	apiTagGet           = "/cgi-bin/tag/get"
	apiTagAddTagUsers   = "/cgi-bin/tag/addtagusers"
	apiTagDelTagUsers   = "/cgi-bin/tag/deltagusers"
	apiTagList          = "/cgi-bin/tag/list"
	apiSyncUser         = "/cgi-bin/batch/syncuser"
	apiReplaceUser      = "/cgi-bin/batch/replaceuser"
	apiReplaceParty     = "/cgi-bin/batch/replaceparty"
	apiJobGetResult     = "/cgi-bin/batch/getresult"
)

/*
创建成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90195

POST https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=ACCESS_TOKEN
*/
func UserCreate(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUserCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
读取成员

在通讯录同步助手中此接口可以读取企业通讯录的所有成员信息，而自建应用可以读取该应用设置的可见范围内的成员信息。

See: https://work.weixin.qq.com/api/doc/90000/90135/90196

GET https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&userid=USERID
*/
func UserGet(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiUserGet + "?" + params.Encode())
}

/*
更新成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90197

POST https://qyapi.weixin.qq.com/cgi-bin/user/update?access_token=ACCESS_TOKEN
*/
func UserUpdate(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUserUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90198

GET https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=ACCESS_TOKEN&userid=USERID
*/
func UserDelete(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiUserDelete + "?" + params.Encode())
}

/*
批量删除成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90199

POST https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete?access_token=ACCESS_TOKEN
*/
func UserBatchDelete(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUserBatchDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取部门成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90200

GET https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=ACCESS_TOKEN&department_id=DEPARTMENT_ID&fetch_child=FETCH_CHILD
*/
func UserSimpleList(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiUserSimpleList + "?" + params.Encode())
}

/*
获取部门成员详情



See: https://work.weixin.qq.com/api/doc/90000/90135/90201

GET https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=ACCESS_TOKEN&department_id=DEPARTMENT_ID&fetch_child=FETCH_CHILD
*/
func UserList(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiUserList + "?" + params.Encode())
}

/*
userid与openid互换

注：需要成员使用微信登录企业微信或者关注微工作台（原企业号）才能转成openid;如果是外部联系人，请使用外部联系人openid转换转换openid

See: https://work.weixin.qq.com/api/doc/90000/90135/90202

POST https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_openid?access_token=ACCESS_TOKEN
*/
func ConvertToOpenid(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiConvertToOpenid, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
openid转userid

该接口主要应用于使用企业支付之后的结果查询。开发者需要知道某个结果事件的openid对应企业微信内成员的信息时，可以通过调用该接口进行转换查询。

See: https://work.weixin.qq.com/api/doc/90000/90135/90202

POST https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_userid?access_token=ACCESS_TOKEN
*/
func ConvertToUserid(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiConvertToUserid, bytes.NewReader(payload), "application/json;charset=utf-8")
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
func BatchInvite(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchInvite, bytes.NewReader(payload), "application/json;charset=utf-8")
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

/*
创建部门



See: https://work.weixin.qq.com/api/doc/90000/90135/90205

POST https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=ACCESS_TOKEN
*/
func DepartmentCreate(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDepartmentCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新部门



See: https://work.weixin.qq.com/api/doc/90000/90135/90206

POST https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=ACCESS_TOKEN
*/
func DepartmentUpdate(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDepartmentUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除部门



See: https://work.weixin.qq.com/api/doc/90000/90135/90207

GET https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=ACCESS_TOKEN&id=ID
*/
func DepartmentDelete(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiDepartmentDelete + "?" + params.Encode())
}

/*
获取部门列表



See: https://work.weixin.qq.com/api/doc/90000/90135/90208

GET https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=ACCESS_TOKEN&id=ID
*/
func DepartmentList(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiDepartmentList + "?" + params.Encode())
}

/*
创建标签



See: https://work.weixin.qq.com/api/doc/90000/90135/90210

POST https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token=ACCESS_TOKEN
*/
func TagCreate(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTagCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新标签名字



See: https://work.weixin.qq.com/api/doc/90000/90135/90211

POST https://qyapi.weixin.qq.com/cgi-bin/tag/update?access_token=ACCESS_TOKEN
*/
func TagUpdate(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTagUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除标签



See: https://work.weixin.qq.com/api/doc/90000/90135/90212

GET https://qyapi.weixin.qq.com/cgi-bin/tag/delete?access_token=ACCESS_TOKEN&tagid=TAGID
*/
func TagDelete(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiTagDelete + "?" + params.Encode())
}

/*
获取标签成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90213

GET https://qyapi.weixin.qq.com/cgi-bin/tag/get?access_token=ACCESS_TOKEN&tagid=TAGID
*/
func TagGet(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiTagGet + "?" + params.Encode())
}

/*
增加标签成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90214

POST https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers?access_token=ACCESS_TOKEN
*/
func TagAddTagUsers(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTagAddTagUsers, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除标签成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90215

POST https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers?access_token=ACCESS_TOKEN
*/
func TagDelTagUsers(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTagDelTagUsers, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取标签列表



See: https://work.weixin.qq.com/api/doc/90000/90135/90216

GET https://qyapi.weixin.qq.com/cgi-bin/tag/list?access_token=ACCESS_TOKEN
*/
func TagList(ctx *corporation.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiTagList)
}

/*
增量更新成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90980

POST https://qyapi.weixin.qq.com/cgi-bin/batch/syncuser?access_token=ACCESS_TOKEN
*/
func SyncUser(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSyncUser, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
全量覆盖成员



See: https://work.weixin.qq.com/api/doc/90000/90135/90981

POST https://qyapi.weixin.qq.com/cgi-bin/batch/replaceuser?access_token=ACCESS_TOKEN
*/
func ReplaceUser(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiReplaceUser, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
全量覆盖部门



See: https://work.weixin.qq.com/api/doc/90000/90135/90982

POST https://qyapi.weixin.qq.com/cgi-bin/batch/replaceparty?access_token=ACCESS_TOKEN
*/
func ReplaceParty(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiReplaceParty, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取异步任务结果



See: https://work.weixin.qq.com/api/doc/90000/90135/90983

GET https://qyapi.weixin.qq.com/cgi-bin/batch/getresult?access_token=ACCESS_TOKEN&jobid=JOBID
*/
func JobGetResult(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiJobGetResult + "?" + params.Encode())
}
