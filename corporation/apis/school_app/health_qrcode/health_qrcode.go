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

// Package health_qrcode 家校应用/复学码
package health_qrcode

import (
	"bytes"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetTeacherCustomizeHealthInfo = "/cgi-bin/school/user/get_teacher_customize_health_info"
	apiGetStudentCustomizeHealthInfo = "/cgi-bin/school/user/get_student_customize_health_info"
	apiGetHealthQrcode               = "/cgi-bin/school/user/get_health_qrcode"
)

/*
获取老师健康信息



See: https://work.weixin.qq.com/api/doc/90000/90135/93744

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/get_teacher_customize_health_info?access_token=ACCESS_TOKEN
*/
func GetTeacherCustomizeHealthInfo(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetTeacherCustomizeHealthInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取学生健康信息



See: https://work.weixin.qq.com/api/doc/90000/90135/93745

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/get_student_customize_health_info?access_token=ACCESS_TOKEN
*/
func GetStudentCustomizeHealthInfo(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetStudentCustomizeHealthInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取师生健康码



See: https://work.weixin.qq.com/api/doc/90000/90135/93746

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/get_health_qrcode?access_token=ACCESS_TOKEN
*/
func GetHealthQrcode(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetHealthQrcode, bytes.NewReader(payload), "application/json;charset=utf-8")
}
