// Copyright 2020 FastWeGo
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

package material_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/wechat4work/corporation"
	"github.com/fastwego/wechat4work/corporation/apis/material"
)

func ExampleUpload() {
	var ctx *corporation.App

	media := ""
	params := url.Values{}
	resp, err := material.Upload(ctx, media, params)

	fmt.Println(resp, err)
}

func ExampleUploadImg() {
	var ctx *corporation.App

	media := ""
	resp, err := material.UploadImg(ctx, media)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := material.Get(ctx, params)

	fmt.Println(resp, err)
}

func ExampleJssdk() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := material.Jssdk(ctx, params)

	fmt.Println(resp, err)
}
