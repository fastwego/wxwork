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

package util_test

import (
	"fmt"

	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/util"
)

func ExampleGetApiDomainIp() {
	var ctx *corporation.App

	resp, err := util.GetApiDomainIp(ctx)

	fmt.Println(resp, err)
}

func ExampleGetCallbackIp() {
	var ctx *corporation.App

	resp, err := util.GetCallbackIp(ctx)

	fmt.Println(resp, err)
}
