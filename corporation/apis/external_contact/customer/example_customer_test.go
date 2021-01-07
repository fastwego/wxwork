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

package customer_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/external_contact/customer"
)

func ExampleList() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := customer.List(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := customer.Get(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetByUser() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer.GetByUser(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleRemark() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer.Remark(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetMobileHashcode() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer.GetMobileHashcode(ctx, payload)

	fmt.Println(resp, err)
}
