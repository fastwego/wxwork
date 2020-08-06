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

package app_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/wechat4work/corporation"
	"github.com/fastwego/wechat4work/corporation/apis/app"
)

func ExampleAgentGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := app.AgentGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleAgentList() {
	var ctx *corporation.App

	resp, err := app.AgentList(ctx)

	fmt.Println(resp, err)
}

func ExampleAgentSet() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := app.AgentSet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMenuCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	params := url.Values{}
	resp, err := app.MenuCreate(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleMenuGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := app.MenuGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleMenuDelete() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := app.MenuDelete(ctx, params)

	fmt.Println(resp, err)
}

func ExampleSetWorkbenchTemplate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := app.SetWorkbenchTemplate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetWorkbenchTemplate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := app.GetWorkbenchTemplate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetWorkbenchData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := app.SetWorkbenchData(ctx, payload)

	fmt.Println(resp, err)
}
