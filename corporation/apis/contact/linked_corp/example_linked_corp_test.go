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

package linked_corp_test

import (
	"fmt"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/contact/linked_corp"
)

func ExampleGetPermList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := linked_corp.GetPermList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := linked_corp.Get(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSimpleList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := linked_corp.SimpleList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUserList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := linked_corp.UserList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDepartmentList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := linked_corp.DepartmentList(ctx, payload)

	fmt.Println(resp, err)
}
