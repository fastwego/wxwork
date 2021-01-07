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

package user_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/contact/user"
)

func ExampleCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.Create(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := user.Get(ctx, params)

	fmt.Println(resp, err)
}

func ExampleUpdate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.Update(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := user.Delete(ctx, params)

	fmt.Println(resp, err)
}

func ExampleBatchDelete() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.BatchDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSimpleList() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := user.SimpleList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleList() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := user.List(ctx, params)

	fmt.Println(resp, err)
}

func ExampleConvertToOpenId() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.ConvertToOpenId(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleConvertToUserId() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.ConvertToUserId(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAuthSucc() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := user.AuthSucc(ctx, params)

	fmt.Println(resp, err)
}

func ExampleInvite() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.Invite(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetJoinQrcode() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := user.GetJoinQrcode(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetActiveStat() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.GetActiveStat(ctx, payload)

	fmt.Println(resp, err)
}
