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

package checkin_test

import (
	"fmt"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/oa/checkin"
)

func ExampleGetCorpCheckinOption() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.GetCorpCheckinOption(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCheckinOption() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.GetCheckinOption(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCheckinData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.GetCheckinData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCheckinDayData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.GetCheckinDayData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCheckinMonthData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.GetCheckinMonthData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCheckinScheduleList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.GetCheckinScheduleList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetCheckinScheduleList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.SetCheckinScheduleList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddCheckinUserFace() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.AddCheckinUserFace(ctx, payload)

	fmt.Println(resp, err)
}
