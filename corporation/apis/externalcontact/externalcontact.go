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

// Package externalcontact 外部联系人管理
package externalcontact

import (
	"bytes"
	"net/url"

	"github.com/fastwego/wechat4work/corporation"
)

const (
	apiGetFollowUserList        = "/cgi-bin/externalcontact/get_follow_user_list"
	apiAddContactWay            = "/cgi-bin/externalcontact/add_contact_way"
	apiGetContactWay            = "/cgi-bin/externalcontact/get_contact_way"
	apiUpdateContactWay         = "/cgi-bin/externalcontact/update_contact_way"
	apiDelContactWay            = "/cgi-bin/externalcontact/del_contact_way"
	apiCloseTempChat            = "/cgi-bin/externalcontact/close_temp_chat"
	apiList                     = "/cgi-bin/externalcontact/list"
	apiGet                      = "/cgi-bin/externalcontact/get"
	apiRemark                   = "/cgi-bin/externalcontact/remark"
	apiGetCorpTagList           = "/cgi-bin/externalcontact/get_corp_tag_list"
	apiAddCorpTag               = "/cgi-bin/externalcontact/add_corp_tag"
	apiEditCorpTag              = "/cgi-bin/externalcontact/edit_corp_tag"
	apiDelCorpTag               = "/cgi-bin/externalcontact/del_corp_tag"
	apiMarkTag                  = "/cgi-bin/externalcontact/mark_tag"
	apiGroupchatList            = "/cgi-bin/externalcontact/groupchat/list"
	apiGroupchatGet             = "/cgi-bin/externalcontact/groupchat/get"
	apiAddMsgTemplate           = "/cgi-bin/externalcontact/add_msg_template"
	apiGetGroupMsgResult        = "/cgi-bin/externalcontact/get_group_msg_result"
	apiSendWelcomeMsg           = "/cgi-bin/externalcontact/send_welcome_msg"
	apiGroupWelcomeTemplateAdd  = "/cgi-bin/externalcontact/group_welcome_template/add"
	apiGroupWelcomeTemplateEdit = "/cgi-bin/externalcontact/group_welcome_template/edit"
	apiGroupWelcomeTemplateGet  = "/cgi-bin/externalcontact/group_welcome_template/get"
	apiGroupWelcomeTemplateDel  = "/cgi-bin/externalcontact/group_welcome_template/del"
	apiGetUnassignedList        = "/cgi-bin/externalcontact/get_unassigned_list"
	apiTransfer                 = "/cgi-bin/externalcontact/transfer"
	apiGroupchatTransfer        = "/cgi-bin/externalcontact/groupchat/transfer"
	apiGetUserBehaviorData      = "/cgi-bin/externalcontact/get_user_behavior_data"
	apiGroupchatStatistic       = "/cgi-bin/externalcontact/groupchat/statistic"
)

/*
获取配置了客户联系功能的成员列表

企业和第三方服务商可通过此接口获取配置了客户联系功能的成员列表。

See: https://work.weixin.qq.com/api/doc/90000/90135/92571

GET https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_follow_user_list?access_token=ACCESS_TOKEN
*/
func GetFollowUserList(ctx *corporation.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetFollowUserList)
}

/*
客户联系「联系我」管理

注意:通过API添加的「联系我」不会在管理端进行展示，每个企业可通过API最多配置50万个「联系我」。用户需要妥善存储返回的config_id，config_id丢失可能导致用户无法编辑或删除「联系我」。临时会话模式不占用「联系我」数量，但每日最多添加10万个，并且仅支持单人。临时会话模式的二维码，添加好友完成后该二维码即刻失效。

See: https://work.weixin.qq.com/api/doc/90000/90135/92572

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_contact_way?access_token=ACCESS_TOKEN
*/
func AddContactWay(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddContactWay, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
客户联系「联系我」管理



See: https://work.weixin.qq.com/api/doc/90000/90135/92572

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_contact_way?access_token=ACCESS_TOKEN
*/
func GetContactWay(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetContactWay, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新企业已配置的「联系我」方式

更新企业配置的「联系我」二维码和「联系我」小程序按钮中的信息，如使用人员和备注等。

See: https://work.weixin.qq.com/api/doc/90000/90135/92572

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/update_contact_way?access_token=ACCESS_TOKEN
*/
func UpdateContactWay(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateContactWay, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
客户联系「联系我」管理



See: https://work.weixin.qq.com/api/doc/90000/90135/92572

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_contact_way?access_token=ACCESS_TOKEN
*/
func DelContactWay(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelContactWay, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
客户联系「联系我」管理



See: https://work.weixin.qq.com/api/doc/90000/90135/92572

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/close_temp_chat?access_token=ACCESS_TOKEN
*/
func CloseTempChat(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCloseTempChat, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取客户列表

企业可通过此接口获取指定成员添加的客户列表。客户是指配置了客户联系功能的成员所添加的外部联系人。没有配置客户联系功能的成员，所添加的外部联系人将不会作为客户返回。

See: https://work.weixin.qq.com/api/doc/90000/90135/92113

GET https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list?access_token=ACCESS_TOKEN&userid=USERID
*/
func List(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiList + "?" + params.Encode())
}

/*
获取客户详情

企业可通过此接口，根据外部联系人的userid（如何获取?），拉取客户详情。

See: https://work.weixin.qq.com/api/doc/90000/90135/92114

GET https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get?access_token=ACCESS_TOKEN&external_userid=EXTERNAL_USERID
*/
func Get(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGet + "?" + params.Encode())
}

/*
修改客户备注信息

企业可通过此接口修改指定用户添加的客户的备注信息。

See: https://work.weixin.qq.com/api/doc/90000/90135/92115

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark?access_token=ACCESS_TOKEN
*/
func Remark(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiRemark, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业标签库

企业可通过此接口获取企业客户标签详情。

See: https://work.weixin.qq.com/api/doc/90000/90135/92117

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_corp_tag_list?access_token=ACCESS_TOKEN
*/
func GetCorpTagList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCorpTagList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
添加企业客户标签

企业可通过此接口向客户标签库中添加新的标签组和标签，每个企业最多可配置3000个企业标签。暂不支持第三方调用。

See: https://work.weixin.qq.com/api/doc/90000/90135/92117

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag?access_token=ACCESS_TOKEN
*/
func AddCorpTag(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddCorpTag, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
编辑企业客户标签

企业可通过此接口编辑客户标签/标签组的名称或次序值。暂不支持第三方调用。

See: https://work.weixin.qq.com/api/doc/90000/90135/92117

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_corp_tag?access_token=ACCESS_TOKEN
*/
func EditCorpTag(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiEditCorpTag, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除企业客户标签

企业可通过此接口删除客户标签库中的标签，或删除整个标签组。暂不支持第三方调用。

See: https://work.weixin.qq.com/api/doc/90000/90135/92117

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_corp_tag?access_token=ACCESS_TOKEN
*/
func DelCorpTag(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelCorpTag, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
编辑客户企业标签

企业可通过此接口为指定成员的客户添加上由企业统一配置的标签。

See: https://work.weixin.qq.com/api/doc/90000/90135/92118

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/mark_tag?access_token=ACCESS_TOKEN
*/
func MarkTag(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMarkTag, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取客户群列表

该接口用于获取配置过客户群管理的客户群列表。

See: https://work.weixin.qq.com/api/doc/90000/90135/92120

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/list?access_token=ACCESS_TOKEN
*/
func GroupchatList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGroupchatList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取客户群详情

通过客户群ID，获取详情。包括群名、群成员列表、群成员入群时间、入群方式。（客户群是由具有客户群使用权限的成员创建的外部群）

See: https://work.weixin.qq.com/api/doc/90000/90135/92122

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/get?access_token=ACCESS_TOKEN
*/
func GroupchatGet(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGroupchatGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
添加企业群发消息任务

企业可通过此接口添加企业群发消息的任务并通知客服人员发送给相关客户或客户群。（注：企业微信终端需升级到2.7.5版本及以上）注意：调用该接口并不会直接发送消息给客户/客户群，需要相关的客服人员操作以后才会实际发送（客服人员的企业微信需要升级到2.7.5及以上版本）同一个企业每个自然月内仅可针对一个客户/客户群发送4条消息，超过限制的用户将会被忽略。

See: https://work.weixin.qq.com/api/doc/90000/90135/92135

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_msg_template?access_token=ACCESS_TOKEN
*/
func AddMsgTemplate(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddMsgTemplate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业群发消息发送结果

企业可通过该接口获取到添加企业群发消息任务的群发发送结果。

See: https://work.weixin.qq.com/api/doc/90000/90135/92136

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_group_msg_result?access_token=ACCESS_TOKEN
*/
func GetGroupMsgResult(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGroupMsgResult, bytes.NewReader(payload), "application/json;charset=utf-8")
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
添加群欢迎语素材

企业可通过此API向企业的群欢迎语素材库中添加素材。

See: https://work.weixin.qq.com/api/doc/90000/90135/92366

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/add?access_token=ACCESS_TOKEN
*/
func GroupWelcomeTemplateAdd(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGroupWelcomeTemplateAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
编辑群欢迎语素材

企业可通过此API编辑欢迎语素材库中的素材。

See: https://work.weixin.qq.com/api/doc/90000/90135/92366

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/edit?access_token=ACCESS_TOKEN
*/
func GroupWelcomeTemplateEdit(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGroupWelcomeTemplateEdit, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取群欢迎语素材

企业可通过此API获取群欢迎语素材。

See: https://work.weixin.qq.com/api/doc/90000/90135/92366

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/get?access_token=ACCESS_TOKEN
*/
func GroupWelcomeTemplateGet(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGroupWelcomeTemplateGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除群欢迎语素材

企业可通过此API删除群欢迎语素材。

See: https://work.weixin.qq.com/api/doc/90000/90135/92366

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/del?access_token=ACCESS_TOKEN
*/
func GroupWelcomeTemplateDel(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGroupWelcomeTemplateDel, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取离职成员的客户列表

企业和第三方可通过此接口，获取所有离职成员的客户列表，并可进一步调用离职成员的外部联系人再分配接口将这些客户重新分配给其他企业成员。

See: https://work.weixin.qq.com/api/doc/90000/90135/92124

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_unassigned_list?access_token=ACCESS_TOKEN
*/
func GetUnassignedList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUnassignedList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
离职成员的外部联系人再分配

企业可通过此接口，将已离职成员的外部联系人分配给另一个成员接替联系。

See: https://work.weixin.qq.com/api/doc/90000/90135/92125

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer?access_token=ACCESS_TOKEN
*/
func Transfer(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTransfer, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
离职成员的群再分配

企业可通过此接口，将已离职成员为群主的群，分配给另一个客服成员。

See: https://work.weixin.qq.com/api/doc/90000/90135/92127

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/transfer?access_token=ACCESS_TOKEN
*/
func GroupchatTransfer(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGroupchatTransfer, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取联系客户统计数据

企业可通过此接口获取成员联系客户的数据，包括发起申请数、新增客户数、聊天数、发送消息数和删除/拉黑成员的客户数等指标。

See: https://work.weixin.qq.com/api/doc/90000/90135/92132

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_user_behavior_data?access_token=ACCESS_TOKEN
*/
func GetUserBehaviorData(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUserBehaviorData, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取客户群统计数据

获取指定日期全天的统计数据。注意，企业微信仅存储180天的数据。

See: https://work.weixin.qq.com/api/doc/90000/90135/92133

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic?access_token=ACCESS_TOKEN
*/
func GroupchatStatistic(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGroupchatStatistic, bytes.NewReader(payload), "application/json;charset=utf-8")
}
