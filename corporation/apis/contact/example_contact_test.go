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

package contact_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/wechat4work/corporation"
	"github.com/fastwego/wechat4work/corporation/apis/contact"
)

func ExampleUserCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.UserCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUserGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := contact.UserGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleUserUpdate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.UserUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUserDelete() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := contact.UserDelete(ctx, params)

	fmt.Println(resp, err)
}

func ExampleUserBatchDelete() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.UserBatchDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUserSimpleList() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := contact.UserSimpleList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleUserList() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := contact.UserList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleConvertToOpenid() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.ConvertToOpenid(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleConvertToUserid() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.ConvertToUserid(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAuthSucc() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := contact.AuthSucc(ctx, params)

	fmt.Println(resp, err)
}

func ExampleBatchInvite() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.BatchInvite(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetJoinQrcode() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := contact.GetJoinQrcode(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetActiveStat() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.GetActiveStat(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDepartmentCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.DepartmentCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDepartmentUpdate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.DepartmentUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDepartmentDelete() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := contact.DepartmentDelete(ctx, params)

	fmt.Println(resp, err)
}

func ExampleDepartmentList() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := contact.DepartmentList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleTagCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.TagCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTagUpdate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.TagUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTagDelete() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := contact.TagDelete(ctx, params)

	fmt.Println(resp, err)
}

func ExampleTagGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := contact.TagGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleTagAddTagUsers() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.TagAddTagUsers(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTagDelTagUsers() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.TagDelTagUsers(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTagList() {
	var ctx *corporation.App

	resp, err := contact.TagList(ctx)

	fmt.Println(resp, err)
}

func ExampleSyncUser() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.SyncUser(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReplaceUser() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.ReplaceUser(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReplaceParty() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := contact.ReplaceParty(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleJobGetResult() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := contact.JobGetResult(ctx, params)

	fmt.Println(resp, err)
}
