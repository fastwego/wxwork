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

package health_test

import (
	"fmt"

	"github.com/fastwego/wechat4work/corporation"
	"github.com/fastwego/wechat4work/corporation/apis/health"
)

func ExampleGetHealthReportStat() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := health.GetHealthReportStat(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetReportJobids() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := health.GetReportJobids(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetReportJobInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := health.GetReportJobInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetReportAnswer() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := health.GetReportAnswer(ctx, payload)

	fmt.Println(resp, err)
}
