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

package user_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/school/user"
)

func ExampleCreateStudent() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.CreateStudent(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDeleteStudent() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := user.DeleteStudent(ctx, params)

	fmt.Println(resp, err)
}

func ExampleUpdateStudent() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.UpdateStudent(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBatchCreateStudent() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.BatchCreateStudent(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBatchDeleteStudent() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.BatchDeleteStudent(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBatchUpdateStudent() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.BatchUpdateStudent(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCreateParent() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.CreateParent(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDeleteParent() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := user.DeleteParent(ctx, params)

	fmt.Println(resp, err)
}

func ExampleUpdateParent() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.UpdateParent(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBatchCreateParent() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.BatchCreateParent(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBatchDeleteParent() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.BatchDeleteParent(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBatchUpdateParent() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.BatchUpdateParent(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := user.Get(ctx, params)

	fmt.Println(resp, err)
}

func ExampleList() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := user.List(ctx, params)

	fmt.Println(resp, err)
}

func ExampleSetArchSyncMode() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := user.SetArchSyncMode(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleListParent() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := user.ListParent(ctx, params)

	fmt.Println(resp, err)
}
