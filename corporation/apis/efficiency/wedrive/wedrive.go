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

// Package wedrive 效率工具/微盘
package wedrive

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiSpaceCreate  = "/cgi-bin/wedrive/space_create"
	apiSpaceRename  = "/cgi-bin/wedrive/space_rename"
	apiSpaceDismiss = "/cgi-bin/wedrive/space_dismiss"
	apiSpaceInfo    = "/cgi-bin/wedrive/space_info"
	apiSpaceAclAdd  = "/cgi-bin/wedrive/space_acl_add"
	apiSpaceAclDel  = "/cgi-bin/wedrive/space_acl_del"
	apiSpaceSetting = "/cgi-bin/wedrive/space_setting"
	apiSpaceShare   = "/cgi-bin/wedrive/space_share"
	apiFileList     = "/cgi-bin/wedrive/file_list"
	apiFileUpload   = "/cgi-bin/wedrive/file_upload"
	apiFileDownload = "/cgi-bin/wedrive/file_download"
	apiFileCreate   = "/cgi-bin/wedrive/file_create"
	apiFileRename   = "/cgi-bin/wedrive/file_rename"
	apiFileMove     = "/cgi-bin/wedrive/file_move"
	apiFileDelete   = "/cgi-bin/wedrive/file_delete"
	apiFileInfo     = "/cgi-bin/wedrive/file_info"
	apiFileAclAdd   = "/cgi-bin/wedrive/file_acl_add"
	apiFileAclDel   = "/cgi-bin/wedrive/file_acl_del"
	apiFileSetting  = "/cgi-bin/wedrive/file_setting"
	apiFileShare    = "/cgi-bin/wedrive/file_share"
)

/*
新建空间

该接口用于在微盘内新建空间，可以指定人创建空间。

See: https://work.weixin.qq.com/api/doc/90000/90135/93655

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_create?access_token=ACCESS_TOKEN
*/
func SpaceCreate(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpaceCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
重命名空间

该接口用于重命名已有空间，接收userid参数，以空间管理员身份来重命名。

See: https://work.weixin.qq.com/api/doc/90000/90135/93655

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_rename?access_token=ACCESS_TOKEN
*/
func SpaceRename(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpaceRename, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
解散空间

该接口用于解散已有空间，需要以空间管理员身份来解散。

See: https://work.weixin.qq.com/api/doc/90000/90135/93655

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_dismiss?access_token=ACCESS_TOKEN
*/
func SpaceDismiss(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpaceDismiss, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取空间信息

该接口用于获取空间成员列表、信息、权限等信息。

See: https://work.weixin.qq.com/api/doc/90000/90135/93655

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_info?access_token=ACCESS_TOKEN
*/
func SpaceInfo(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpaceInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
添加成员/部门

该接口用于对指定空间添加成员/部门，可一次性添加多个。

See: https://work.weixin.qq.com/api/doc/90000/90135/93656

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_acl_add?access_token=ACCESS_TOKEN
*/
func SpaceAclAdd(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpaceAclAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
移除成员/部门

该接口用于对指定空间移除成员/部门，操作者需要有移除权限。

See: https://work.weixin.qq.com/api/doc/90000/90135/93656

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_acl_del?access_token=ACCESS_TOKEN
*/
func SpaceAclDel(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpaceAclDel, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
权限管理

该接口用于修改空间权限，需要传入userid，修改权限范围继承传入用户的权限范围。

See: https://work.weixin.qq.com/api/doc/90000/90135/93656

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_setting?access_token=ACCESS_TOKEN
*/
func SpaceSetting(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpaceSetting, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取邀请链接

该接口用于获取空间邀请分享链接。

See: https://work.weixin.qq.com/api/doc/90000/90135/93656

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_share?access_token=ACCESS_TOKEN
*/
func SpaceShare(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpaceShare, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取文件列表

该接口用于获取指定地址下的文件列表。

See: https://work.weixin.qq.com/api/doc/90000/90135/93657

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_list?access_token=ACCESS_TOKEN
*/
func FileList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiFileList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
上传文件

该接口用于向微盘中的指定位置上传文件。

See: https://work.weixin.qq.com/api/doc/90000/90135/93657

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_upload?access_token=ACCESS_TOKEN
*/
func FileUpload(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiFileUpload, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
下载文件

该接口用于下载文件，请求的userid需有下载权限。

See: https://work.weixin.qq.com/api/doc/90000/90135/93657

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_download?access_token=ACCESS_TOKEN
*/
func FileDownload(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiFileDownload, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
新建文件/微文档

该接口用于在微盘指定位置新建文件、微文档。

See: https://work.weixin.qq.com/api/doc/90000/90135/93657

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_create?access_token=ACCESS_TOKEN
*/
func FileCreate(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiFileCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
重命名文件

该接口用于对指定文件进行重命名。

See: https://work.weixin.qq.com/api/doc/90000/90135/93657

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_rename?access_token=ACCESS_TOKEN
*/
func FileRename(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiFileRename, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
移动文件

该接口用于将文件移动到指定位置。

See: https://work.weixin.qq.com/api/doc/90000/90135/93657

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_move?access_token=ACCESS_TOKEN
*/
func FileMove(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiFileMove, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除文件

该接口用于删除指定文件。

See: https://work.weixin.qq.com/api/doc/90000/90135/93657

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_delete?access_token=ACCESS_TOKEN
*/
func FileDelete(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiFileDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
文件信息

该接口用于获取指定文件的信息。

See: https://work.weixin.qq.com/api/doc/90000/90135/93657

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_info?access_token=ACCESS_TOKEN
*/
func FileInfo(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiFileInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
新增指定人

该接口用于对指定文件添加指定人/部门。

See: https://work.weixin.qq.com/api/doc/90000/90135/93658

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_acl_add?access_token=ACCESS_TOKEN
*/
func FileAclAdd(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiFileAclAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除指定人

该接口用于删除指定文件的指定人/部门。

See: https://work.weixin.qq.com/api/doc/90000/90135/93658

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_acl_del?access_token=ACCESS_TOKEN
*/
func FileAclDel(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiFileAclDel, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
分享设置

该接口用于文件的分享设置。

See: https://work.weixin.qq.com/api/doc/90000/90135/93658

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_setting?access_token=ACCESS_TOKEN
*/
func FileSetting(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiFileSetting, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取分享链接

该接口用于获取文件的分享链接。

See: https://work.weixin.qq.com/api/doc/90000/90135/93658

POST https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_share?access_token=ACCESS_TOKEN
*/
func FileShare(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiFileShare, bytes.NewReader(payload), "application/json;charset=utf-8")
}
