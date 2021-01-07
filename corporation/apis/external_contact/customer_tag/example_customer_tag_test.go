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

package customer_tag_test

import (
	"fmt"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/external_contact/customer_tag"
)

func ExampleGetCorpTagList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_tag.GetCorpTagList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddCorpTag() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_tag.AddCorpTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleEditCorpTag() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_tag.EditCorpTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelCorpTag() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_tag.DelCorpTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMarkTag() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_tag.MarkTag(ctx, payload)

	fmt.Println(resp, err)
}
