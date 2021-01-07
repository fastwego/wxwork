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

package customer_stat_test

import (
	"fmt"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/external_contact/customer_stat"
)

func ExampleGetUserBehaviorData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_stat.GetUserBehaviorData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleStatistic() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_stat.Statistic(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleStatisticGroupByDay() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_stat.StatisticGroupByDay(ctx, payload)

	fmt.Println(resp, err)
}
