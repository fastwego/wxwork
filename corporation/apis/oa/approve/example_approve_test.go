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

package approve_test

import (
	"fmt"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/oa/approve"
)

func ExampleGetTemplateDetail() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := approve.GetTemplateDetail(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleApplyEvent() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := approve.ApplyEvent(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetApprovalInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := approve.GetApprovalInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetApprovalDetail() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := approve.GetApprovalDetail(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetApprovalData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := approve.GetApprovalData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCorpConf() {
	var ctx *corporation.App

	resp, err := approve.GetCorpConf(ctx)

	fmt.Println(resp, err)
}

func ExampleGetUserVacationQuota() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := approve.GetUserVacationQuota(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetOneUserQuota() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := approve.SetOneUserQuota(ctx, payload)

	fmt.Println(resp, err)
}
