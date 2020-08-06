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

package externalcontact_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/wechat4work/corporation"
	"github.com/fastwego/wechat4work/corporation/apis/externalcontact"
)

func ExampleGetFollowUserList() {
	var ctx *corporation.App

	resp, err := externalcontact.GetFollowUserList(ctx)

	fmt.Println(resp, err)
}

func ExampleAddContactWay() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.AddContactWay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetContactWay() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.GetContactWay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateContactWay() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.UpdateContactWay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelContactWay() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.DelContactWay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCloseTempChat() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.CloseTempChat(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleList() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := externalcontact.List(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := externalcontact.Get(ctx, params)

	fmt.Println(resp, err)
}

func ExampleRemark() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.Remark(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCorpTagList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.GetCorpTagList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddCorpTag() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.AddCorpTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleEditCorpTag() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.EditCorpTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelCorpTag() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.DelCorpTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMarkTag() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.MarkTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupchatList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.GroupchatList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupchatGet() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.GroupchatGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddMsgTemplate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.AddMsgTemplate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGroupMsgResult() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.GetGroupMsgResult(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSendWelcomeMsg() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.SendWelcomeMsg(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupWelcomeTemplateAdd() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.GroupWelcomeTemplateAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupWelcomeTemplateEdit() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.GroupWelcomeTemplateEdit(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupWelcomeTemplateGet() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.GroupWelcomeTemplateGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupWelcomeTemplateDel() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.GroupWelcomeTemplateDel(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUnassignedList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.GetUnassignedList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTransfer() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.Transfer(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupchatTransfer() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.GroupchatTransfer(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUserBehaviorData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.GetUserBehaviorData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupchatStatistic() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := externalcontact.GroupchatStatistic(ctx, payload)

	fmt.Println(resp, err)
}
