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

// Package user 家校沟通/学生与家长管理
package user

import (
	"bytes"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiCreateStudent      = "/cgi-bin/school/user/create_student"
	apiDeleteStudent      = "/cgi-bin/school/user/delete_student"
	apiUpdateStudent      = "/cgi-bin/school/user/update_student"
	apiBatchCreateStudent = "/cgi-bin/school/user/batch_create_student"
	apiBatchDeleteStudent = "/cgi-bin/school/user/batch_delete_student"
	apiBatchUpdateStudent = "/cgi-bin/school/user/batch_update_student"
	apiCreateParent       = "/cgi-bin/school/user/create_parent"
	apiDeleteParent       = "/cgi-bin/school/user/delete_parent"
	apiUpdateParent       = "/cgi-bin/school/user/update_parent"
	apiBatchCreateParent  = "/cgi-bin/school/user/batch_create_parent"
	apiBatchDeleteParent  = "/cgi-bin/school/user/batch_delete_parent"
	apiBatchUpdateParent  = "/cgi-bin/school/user/batch_update_parent"
	apiGet                = "/cgi-bin/school/user/get"
	apiList               = "/cgi-bin/school/user/list"
	apiSetArchSyncMode    = "/cgi-bin/school/set_arch_sync_mode"
	apiListParent         = "/cgi-bin/school/user/list_parent"
)

/*
创建学生



See: https://work.weixin.qq.com/api/doc/90000/90135/92325

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/create_student?access_token=ACCESS_TOKEN
*/
func CreateStudent(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreateStudent, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除学生



See: https://work.weixin.qq.com/api/doc/90000/90135/92326

GET https://qyapi.weixin.qq.com/cgi-bin/school/user/delete_student?access_token=ACCESS_TOKEN&userid=USERID
*/
func DeleteStudent(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiDeleteStudent + "?" + params.Encode())
}

/*
更新学生



See: https://work.weixin.qq.com/api/doc/90000/90135/92327

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/update_student?access_token=ACCESS_TOKEN
*/
func UpdateStudent(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateStudent, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量创建学生



See: https://work.weixin.qq.com/api/doc/90000/90135/92328

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_create_student?access_token=ACCESS_TOKEN
*/
func BatchCreateStudent(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchCreateStudent, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量删除学生



See: https://work.weixin.qq.com/api/doc/90000/90135/92329

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_delete_student?access_token=ACCESS_TOKEN
*/
func BatchDeleteStudent(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchDeleteStudent, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量更新学生



See: https://work.weixin.qq.com/api/doc/90000/90135/92330

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_update_student?access_token=ACCESS_TOKEN
*/
func BatchUpdateStudent(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchUpdateStudent, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
创建家长



See: https://work.weixin.qq.com/api/doc/90000/90135/92331

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/create_parent?access_token=ACCESS_TOKEN
*/
func CreateParent(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreateParent, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除家长



See: https://work.weixin.qq.com/api/doc/90000/90135/92332

GET https://qyapi.weixin.qq.com/cgi-bin/school/user/delete_parent?access_token=ACCESS_TOKEN&userid=USERID
*/
func DeleteParent(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiDeleteParent + "?" + params.Encode())
}

/*
更新家长



See: https://work.weixin.qq.com/api/doc/90000/90135/92333

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/update_parent?access_token=ACCESS_TOKEN
*/
func UpdateParent(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateParent, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量创建家长



See: https://work.weixin.qq.com/api/doc/90000/90135/92334

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_create_parent?access_token=ACCESS_TOKEN
*/
func BatchCreateParent(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchCreateParent, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量删除家长



See: https://work.weixin.qq.com/api/doc/90000/90135/92335

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_delete_parent?access_token=ACCESS_TOKEN
*/
func BatchDeleteParent(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchDeleteParent, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量更新家长



See: https://work.weixin.qq.com/api/doc/90000/90135/92336

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_update_parent?access_token=ACCESS_TOKEN
*/
func BatchUpdateParent(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchUpdateParent, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
读取学生或家长



See: https://work.weixin.qq.com/api/doc/90000/90135/92337

GET https://qyapi.weixin.qq.com/cgi-bin/school/user/get?access_token=ACCESS_TOKEN&userid=USERID
*/
func Get(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGet + "?" + params.Encode())
}

/*
获取部门成员详情



See: https://work.weixin.qq.com/api/doc/90000/90135/92338

GET https://qyapi.weixin.qq.com/cgi-bin/school/user/list?access_token=ACCESS_TOKEN&department_id=DEPARTMENT_ID&fetch_child=FETCH_CHILD
*/
func List(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiList + "?" + params.Encode())
}

/*
设置家校通讯录自动同步模式

企业和第三方可通过此接口修改家校通讯录与班级标签之间的自动同步模式，注意，一旦设置禁止自动同步，将无法再次开启。

See: https://work.weixin.qq.com/api/doc/90000/90135/92345

POST https://qyapi.weixin.qq.com/cgi-bin/school/set_arch_sync_mode?access_token=ACCESS_TOKEN
*/
func SetArchSyncMode(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetArchSyncMode, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取部门家长详情



See: https://work.weixin.qq.com/api/doc/90000/90135/92446

GET https://qyapi.weixin.qq.com/cgi-bin/school/user/list_parent?access_token=ACCESS_TOKEN&department_id=DEPARTMENT_ID
*/
func ListParent(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiListParent + "?" + params.Encode())
}
