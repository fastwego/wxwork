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

package oa_test

import (
	"fmt"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/oa"
)

func ExampleGetCheckinData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := oa.GetCheckinData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCheckinPption() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := oa.GetCheckinPption(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetTemplateDetail() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := oa.GetTemplateDetail(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleApplyEvent() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := oa.ApplyEvent(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetApprovalInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := oa.GetApprovalInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetApprovalDetail() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := oa.GetApprovalDetail(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetApprovalData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := oa.GetApprovalData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetDialRecord() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := oa.GetDialRecord(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetOpenApprovalData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := oa.GetOpenApprovalData(ctx, payload)

	fmt.Println(resp, err)
}
