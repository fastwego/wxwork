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

// Package customer_service 客户联系/企业服务人员管理
package customer_service

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetFollowUserList = "/cgi-bin/externalcontact/get_follow_user_list"
	apiAddContactWay     = "/cgi-bin/externalcontact/add_contact_way"
	apiGetContactWay     = "/cgi-bin/externalcontact/get_contact_way"
	apiUpdateContactWay  = "/cgi-bin/externalcontact/update_contact_way"
	apiDelContactWay     = "/cgi-bin/externalcontact/del_contact_way"
	apiCloseTempChat     = "/cgi-bin/externalcontact/close_temp_chat"
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
配置客户联系「联系我」方式

注意:通过API添加的「联系我」不会在管理端进行展示，每个企业可通过API最多配置50万个「联系我」。用户需要妥善存储返回的config_id，config_id丢失可能导致用户无法编辑或删除「联系我」。临时会话模式不占用「联系我」数量，但每日最多添加10万个，并且仅支持单人。临时会话模式的二维码，添加好友完成后该二维码即刻失效。

See: https://work.weixin.qq.com/api/doc/90000/90135/92572

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_contact_way?access_token=ACCESS_TOKEN
*/
func AddContactWay(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddContactWay, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取企业已配置的「联系我」方式



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
删除企业已配置的「联系我」方式



See: https://work.weixin.qq.com/api/doc/90000/90135/92572

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_contact_way?access_token=ACCESS_TOKEN
*/
func DelContactWay(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelContactWay, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
结束临时会话



See: https://work.weixin.qq.com/api/doc/90000/90135/92572

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/close_temp_chat?access_token=ACCESS_TOKEN
*/
func CloseTempChat(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCloseTempChat, bytes.NewReader(payload), "application/json;charset=utf-8")
}
