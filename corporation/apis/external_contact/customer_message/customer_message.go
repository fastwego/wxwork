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

// Package customer_message 客户联系/消息推送
package customer_message

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiAddMsgTemplate           = "/cgi-bin/externalcontact/add_msg_template"
	apiGetGroupmsgList          = "/cgi-bin/externalcontact/get_groupmsg_list"
	apiGetGroupmsgTask          = "/cgi-bin/externalcontact/get_groupmsg_task"
	apiGetGroupmsgSendResult    = "/cgi-bin/externalcontact/get_groupmsg_send_result"
	apiSendWelcomeMsg           = "/cgi-bin/externalcontact/send_welcome_msg"
	apiGroupWelcomeTemplateAdd  = "/cgi-bin/externalcontact/group_welcome_template/add"
	apiGroupWelcomeTemplateEdit = "/cgi-bin/externalcontact/group_welcome_template/edit"
	apiGroupWelcomeTemplateGet  = "/cgi-bin/externalcontact/group_welcome_template/get"
	apiGroupWelcomeTemplateDel  = "/cgi-bin/externalcontact/group_welcome_template/del"
)

/*
创建企业群发

企业跟第三方应用可通过此接口添加企业群发消息的任务并通知成员发送给相关客户或客户群。（注：企业微信终端需升级到2.7.5版本及以上）注意：调用该接口并不会直接发送消息给客户/客户群，需要成员确认后才会执行发送（客服人员的企业微信需要升级到2.7.5及以上版本）同一个企业每个自然月内仅可针对一个客户/客户群发送4条消息，超过接收上限的客户将无法再收到群发消息。

See: https://work.weixin.qq.com/api/doc/90000/90135/92135

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_msg_template?access_token=ACCESS_TOKEN
*/
func AddMsgTemplate(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddMsgTemplate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取群发记录列表

企业和第三方应用可通过此接口获取企业与成员的群发记录。

See: https://work.weixin.qq.com/api/doc/90000/90135/93338

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_list?access_token=ACCESS_TOKEN
*/
func GetGroupmsgList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGroupmsgList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业的全部群发记录



See: https://work.weixin.qq.com/api/doc/90000/90135/93338

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_task?access_token=ACCESS_TOKEN
*/
func GetGroupmsgTask(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGroupmsgTask, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业的全部群发记录



See: https://work.weixin.qq.com/api/doc/90000/90135/93338

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_send_result?access_token=ACCESS_TOKEN
*/
func GetGroupmsgSendResult(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGroupmsgSendResult, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
发送新客户欢迎语



See: https://work.weixin.qq.com/api/doc/90000/90135/92137

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/send_welcome_msg?access_token=ACCESS_TOKEN
*/
func SendWelcomeMsg(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSendWelcomeMsg, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
添加入群欢迎语素材

企业可通过此API向企业的入群欢迎语素材库中添加素材。每个企业的入群欢迎语素材库中，最多容纳100个素材。

See: https://work.weixin.qq.com/api/doc/90000/90135/92366

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/add?access_token=ACCESS_TOKEN
*/
func GroupWelcomeTemplateAdd(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGroupWelcomeTemplateAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
编辑入群欢迎语素材

企业可通过此API编辑入群欢迎语素材库中的素材，且仅能够编辑调用方自己创建的入群欢迎语素材。

See: https://work.weixin.qq.com/api/doc/90000/90135/92366

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/edit?access_token=ACCESS_TOKEN
*/
func GroupWelcomeTemplateEdit(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGroupWelcomeTemplateEdit, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取入群欢迎语素材

企业可通过此API获取入群欢迎语素材。

See: https://work.weixin.qq.com/api/doc/90000/90135/92366

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/get?access_token=ACCESS_TOKEN
*/
func GroupWelcomeTemplateGet(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGroupWelcomeTemplateGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除入群欢迎语素材

企业可通过此API删除入群欢迎语素材，且仅能删除调用方自己创建的入群欢迎语素材。

See: https://work.weixin.qq.com/api/doc/90000/90135/92366

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/del?access_token=ACCESS_TOKEN
*/
func GroupWelcomeTemplateDel(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGroupWelcomeTemplateDel, bytes.NewReader(payload), "application/json;charset=utf-8")
}
