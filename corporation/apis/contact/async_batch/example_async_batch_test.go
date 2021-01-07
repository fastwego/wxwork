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

package async_batch_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/contact/async_batch"
)

func ExampleUser() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := async_batch.User(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReplaceUser() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := async_batch.ReplaceUser(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReplaceParty() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := async_batch.ReplaceParty(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetResult() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := async_batch.GetResult(ctx, params)

	fmt.Println(resp, err)
}
