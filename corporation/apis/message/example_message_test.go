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

package message_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/message"
)

func ExampleSend() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := message.Send(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateTaskcard() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := message.UpdateTaskcard(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAppchatCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := message.AppchatCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAppchatUpdate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := message.AppchatUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAppchatGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := message.AppchatGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleAppchatSend() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := message.AppchatSend(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleLinkedcorpMessageSend() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := message.LinkedcorpMessageSend(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetStatistics() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := message.GetStatistics(ctx, payload)

	fmt.Println(resp, err)
}
