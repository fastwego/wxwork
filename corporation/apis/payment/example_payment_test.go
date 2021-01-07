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

package payment_test

import (
	"fmt"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/payment"
)

func ExampleAddMerchant() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := payment.AddMerchant(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetMerchant() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := payment.GetMerchant(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelMerchant() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := payment.DelMerchant(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetMchUseScope() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := payment.SetMchUseScope(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetBillList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := payment.GetBillList(ctx, payload)

	fmt.Println(resp, err)
}
