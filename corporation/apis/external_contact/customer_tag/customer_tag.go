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

// Package customer_tag 客户联系/客户标签管理
package customer_tag

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetCorpTagList = "/cgi-bin/externalcontact/get_corp_tag_list"
	apiAddCorpTag     = "/cgi-bin/externalcontact/add_corp_tag"
	apiEditCorpTag    = "/cgi-bin/externalcontact/edit_corp_tag"
	apiDelCorpTag     = "/cgi-bin/externalcontact/del_corp_tag"
	apiMarkTag        = "/cgi-bin/externalcontact/mark_tag"
)

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

企业可通过此接口向客户标签库中添加新的标签组和标签，每个企业最多可配置3000个企业标签。

See: https://work.weixin.qq.com/api/doc/90000/90135/92117

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag?access_token=ACCESS_TOKEN
*/
func AddCorpTag(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddCorpTag, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
编辑企业客户标签

企业可通过此接口编辑客户标签/标签组的名称或次序值。

See: https://work.weixin.qq.com/api/doc/90000/90135/92117

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_corp_tag?access_token=ACCESS_TOKEN
*/
func EditCorpTag(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiEditCorpTag, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除企业客户标签

企业可通过此接口删除客户标签库中的标签，或删除整个标签组。

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
