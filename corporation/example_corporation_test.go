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

package corporation_test

import (
	"log"
	"os"

	"github.com/faabiosr/cachego/file"
	"github.com/faabiosr/cachego/sync"
	"github.com/fastwego/wxwork/corporation"
	"github.com/gomodule/redigo/redis"
)

func ExampleCorporation_SetAccessTokenCacheDriver() {
	var App *corporation.App

	// 使用内存
	App.SetAccessTokenCacheDriver(sync.New())

	// 使用指定目录下的 文件
	App.SetAccessTokenCacheDriver(file.New(os.TempDir()))

	// 更多驱动 请查看 https://github.com/faabiosr/cachego
}

func ExampleCorporation_SetLogger() {

	var Ctx *corporation.Corporation

	// 输出日志到控制台
	Ctx.SetLogger(log.New(os.Stdout, "[corporation ]", log.LstdFlags))

	// 记录日志到指定文件
	logFile, _ := os.OpenFile("/path/to/file", os.O_WRONLY, 0644)
	Ctx.SetLogger(log.New(logFile, "[corporation ]", log.LstdFlags))

	// 关闭日志
	Ctx.SetLogger(nil)
}

func ExampleCorporation_SetGetAccessTokenHandler() {
	var Ctx *corporation.App

	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	_, _ = conn.Do("AUTH", "PASSWORD")

	// 从远程 Redis 服务器 获取 AccessToken
	Ctx.SetGetAccessTokenHandler(func(ctx *corporation.App) (accessToken string, err error) {
		accessToken, _ = redis.String(conn.Do("GET", "access_token:"+ctx.Config.AgentId))
		return
	})
}
