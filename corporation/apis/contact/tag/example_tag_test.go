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

package tag_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/contact/tag"
)

func ExampleCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := tag.Create(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := tag.Update(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := tag.Delete(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := tag.Get(ctx, params)

	fmt.Println(resp, err)
}

func ExampleAddTagUsers() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := tag.AddTagUsers(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelTagUsers() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := tag.DelTagUsers(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleList() {
	var ctx *corporation.App

	resp, err := tag.List(ctx)

	fmt.Println(resp, err)
}
