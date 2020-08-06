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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/iancoleman/strcase"
)

const ServerUrl = `https://work.weixin.qq.com/api/docFetch/fetchCnt?id=%s`

var apiUniqMap = map[string]bool{}

func run() {

	docsJson, err := ioutil.ReadFile("./data/docs.json")

	file, err := ioutil.ReadFile("./data/corp_doc_links.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	matched := regexp.MustCompile(`data-id="(\d+)"`).FindAllStringSubmatch(string(file), -1)

	for _, match := range matched {
		//fmt.Println(match)

		category_id := match[1]
		m := regexp.MustCompile(`"category_id": ` + category_id + `,\s+"doc_id": (\d+),`).FindStringSubmatch(string(docsJson))
		if len(m) < 2 {
			continue
		}

		doc_id := m[1]
		if doc_id == "0" {
			continue
		}

		link := fmt.Sprintf(ServerUrl, doc_id)
		resp, err := http.Get(link)
		if err != nil {
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		type Cnt struct {
			Data struct {
				Document struct {
					ID          int    `json:"id"`
					Title       string `json:"title"`
					Time        int    `json:"time"`
					Type        int    `json:"type"`
					ParentID    int    `json:"parent_id"`
					ContentHTML string `json:"content_html"`
					IsDeleted   bool   `json:"is_deleted"`
					ContentMd   string `json:"content_md"`
					Status      int    `json:"status"`
					DocID       int    `json:"doc_id"`
					OrderID     int    `json:"order_id"`
					ContentTxt  string `json:"content_txt"`
					GrayStatus  int    `json:"gray_status"`
					GrayInfo    struct {
						LevelID []interface{} `json:"level_id"`
					} `json:"gray_info"`
				} `json:"document"`
			} `json:"data"`
		}

		cnt := Cnt{}
		err = json.Unmarshal(body, &cnt)
		if err != nil {
			continue
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(cnt.Data.Document.ContentHTML))
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(cnt.Data.Document.Title)

		doc.Find("strong").Each(func(i int, selection *goquery.Selection) {

			if !strings.Contains(selection.Text(), "请求方式") {
				return
			}
			//fmt.Println(selection.Text())
			line := selection.Parent().Text()

			m2 := regexp.MustCompile(`请求方式.+(GET|POST)(.+请求地址.+(https://qyapi.weixin.qq.com/\S+))?`).FindStringSubmatch(line)

			//fmt.Println(line,"\n", m2)

			method := "POST"
			if m2[1] != "" {
				method = m2[1]
			} else {
				return
			}

			apiurl := ""
			if len(m2) == 4 && m2[3] != "" {
				apiurl = m2[3]
			} else {

				p := selection.Parent().Next().Text()
				m3 := regexp.MustCompile(`请求地址.+(https://qyapi.weixin.qq.com/\S+)`).FindStringSubmatch(p)
				if len(m3) < 2 || m3[1] == "" {
					return
				}

				apiurl = m3[1]
			}

			_REQUEST_ := method + " " + apiurl

			parse, _ := url.Parse(apiurl)
			if _, ok := apiUniqMap[parse.Path]; !ok {
				apiUniqMap[parse.Path] = true
			} else {
				return
			}

			_DESCRIPTION_ := ""
			isDesc := selection.Parent().Prev().Is("p")
			if isDesc {
				_DESCRIPTION_ = selection.Parent().Prev().Text()
			}

			_NAME_ := cnt.Data.Document.Title
			isName := selection.Parent().Prev().Prev().Is("h2")
			if isName {
				_NAME_ = selection.Parent().Prev().Prev().Text()
			}

			_SEE_ := fmt.Sprintf("https://work.weixin.qq.com/api/doc/90000/90135/%s", category_id)
			_FUNCNAME_ := strcase.ToCamel(path.Base(parse.Path))

			//GetParams: []Param{
			//	{Name: `appid`, Type: `string`},
			//	{Name: `redirect_uri`, Type: `string`},
			//},

			_GET_PARAMS_ := ""
			fields := strings.Fields(_REQUEST_)
			parse, _ = url.Parse(fields[1])
			for param_name, _ := range parse.Query() {
				if param_name == "access_token" {
					continue
				}
				_GET_PARAMS_ += "\t\t{Name: `" + param_name + "`, Type: `string`},\n"
			}
			if _GET_PARAMS_ != "" {
				_GET_PARAMS_ = `GetParams: []Param{
` + _GET_PARAMS_ + "\t},"
			}

			tpl := strings.ReplaceAll(itemTpl, "_NAME_", _NAME_)
			tpl = strings.ReplaceAll(tpl, "_DESCRIPTION_", _DESCRIPTION_)
			tpl = strings.ReplaceAll(tpl, "_REQUEST_", _REQUEST_)
			tpl = strings.ReplaceAll(tpl, "_SEE_", _SEE_)
			tpl = strings.ReplaceAll(tpl, "_FUNCNAME_", _FUNCNAME_)
			tpl = strings.ReplaceAll(tpl, "_GET_PARAMS_", _GET_PARAMS_)

			fmt.Println(tpl)
		})
	}
}

var itemTpl = `{
	Name:        "_NAME_",
	Description: "_DESCRIPTION_",
	Request:     "_REQUEST_",
	See:         "_SEE_",
	FuncName:    "_FUNCNAME_",
	_GET_PARAMS_
},`
