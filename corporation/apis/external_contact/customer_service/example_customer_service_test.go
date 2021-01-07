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

package customer_service_test

import (
	"fmt"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/external_contact/customer_service"
)

func ExampleGetFollowUserList() {
	var ctx *corporation.App

	resp, err := customer_service.GetFollowUserList(ctx)

	fmt.Println(resp, err)
}

func ExampleAddContactWay() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_service.AddContactWay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetContactWay() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_service.GetContactWay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateContactWay() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_service.UpdateContactWay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelContactWay() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_service.DelContactWay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCloseTempChat() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_service.CloseTempChat(ctx, payload)

	fmt.Println(resp, err)
}
