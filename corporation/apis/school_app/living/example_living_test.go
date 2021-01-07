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

package living_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/school_app/living"
)

func ExampleGetUserAllLivingId() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := living.GetUserAllLivingId(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetLivingInfo() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := living.GetLivingInfo(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetWatchStat() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := living.GetWatchStat(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUnwatchStat() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := living.GetUnwatchStat(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDeleteReplayData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := living.DeleteReplayData(ctx, payload)

	fmt.Println(resp, err)
}
