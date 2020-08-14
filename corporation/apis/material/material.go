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

// Package material 素材管理
package material

import (
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiUpload    = "/cgi-bin/media/upload"
	apiUploadImg = "/cgi-bin/media/uploadimg"
	apiGet       = "/cgi-bin/media/get"
	apiJssdk     = "/cgi-bin/media/get/jssdk"
)

/*
上传临时素材



See: https://work.weixin.qq.com/api/doc/90000/90135/90253

POST(@media) https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
*/
func Upload(ctx *corporation.App, media string, params url.Values) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(media))
		if err != nil {
			return
		}
		file, err := os.Open(media)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

	}()
	return ctx.Client.HTTPPost(apiUpload+"?"+params.Encode(), r, m.FormDataContentType())
}

/*
上传图片



See: https://work.weixin.qq.com/api/doc/90000/90135/90256

POST(@media) https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN
*/
func UploadImg(ctx *corporation.App, media string) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(media))
		if err != nil {
			return
		}
		file, err := os.Open(media)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

	}()
	return ctx.Client.HTTPPost(apiUploadImg, r, m.FormDataContentType())
}

/*
获取临时素材



See: https://work.weixin.qq.com/api/doc/90000/90135/90254

GET https://qyapi.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
*/
func Get(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGet + "?" + params.Encode())
}

/*
获取高清语音素材



See: https://work.weixin.qq.com/api/doc/90000/90135/90255

GET https://qyapi.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
*/
func Jssdk(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiJssdk + "?" + params.Encode())
}
