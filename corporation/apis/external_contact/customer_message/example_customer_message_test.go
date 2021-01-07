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

package customer_message_test

import (
	"fmt"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/external_contact/customer_message"
)

func ExampleAddMsgTemplate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.AddMsgTemplate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGroupmsgList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.GetGroupmsgList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGroupmsgTask() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.GetGroupmsgTask(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGroupmsgSendResult() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.GetGroupmsgSendResult(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSendWelcomeMsg() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.SendWelcomeMsg(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupWelcomeTemplateAdd() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.GroupWelcomeTemplateAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupWelcomeTemplateEdit() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.GroupWelcomeTemplateEdit(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupWelcomeTemplateGet() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.GroupWelcomeTemplateGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupWelcomeTemplateDel() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.GroupWelcomeTemplateDel(ctx, payload)

	fmt.Println(resp, err)
}
