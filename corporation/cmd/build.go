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
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

type Param struct {
	Name string
	Type string
}

type Api struct {
	Name        string
	Description string
	Request     string
	See         string
	FuncName    string
	GetParams   []Param
}

type ApiGroup struct {
	Name    string
	Apis    []Api
	Package string
}

var apiConfig = []ApiGroup{}

func main() {

	var pkgFlag string
	flag.StringVar(&pkgFlag, "package", "default", "")
	flag.Parse()

	apiConfig = apiConfigCrop

	for _, group := range apiConfig {

		if group.Package == pkgFlag {
			build(group)
		}
	}

	if pkgFlag == "apilist" {
		apilist()
	}
}

func apilist() {
	for _, group := range apiConfig {
		fmt.Printf("- %s(%s)\n", group.Name, group.Package)
		for _, api := range group.Apis {
			split := strings.Split(api.Request, " ")
			parse, _ := url.Parse(split[1])

			if api.FuncName == "" {
				api.FuncName = strcase.ToCamel(path.Base(parse.Path))
			}

			godocLink := fmt.Sprintf("https://pkg.go.dev/github.com/fastwego/wxwork/corporation/apis/%s?tab=doc#%s", group.Package, api.FuncName)
			fmt.Printf("\t- [%s](%s) \n\t\t- [%s (%s)](%s)\n", api.Name, api.See, api.FuncName, parse.Path, godocLink)
		}
	}
}

func build(group ApiGroup) {
	var funcs []string
	var consts []string
	var testFuncs []string
	var exampleFuncs []string

	for _, api := range group.Apis {
		tpl := postFuncTpl
		_FUNC_NAME_ := ""
		_GET_PARAMS_ := ""
		_GET_SUFFIX_PARAMS_ := ""
		_UPLOAD_ := "media"
		_FIELD_NAME_ := ""
		_FIELDS_ := ""
		_PAYLOAD_ := ""
		switch {
		case strings.Contains(api.Request, "GET http"):
			tpl = getFuncTpl
		case strings.Contains(api.Request, "POST http"):
			tpl = postFuncTpl
		case strings.Contains(api.Request, "POST(@media"):
			tpl = postUploadFuncTpl
			_UPLOAD_ = "media"

			pattern := `POST\(@media\|field=(\S+)\) http`
			reg := regexp.MustCompile(pattern)
			matched := reg.FindAllStringSubmatch(api.Request, -1)
			if matched != nil {
				_FIELD_NAME_ = matched[0][1]
				_PAYLOAD_ = ", payload []byte"
			}
		}
		if len(api.GetParams) > 0 {
			_GET_PARAMS_ = `, params url.Values`
			//if strings.Contains(api.Request, "POST") {
			//	_GET_PARAMS_ = `, ` + _GET_PARAMS_
			//}
			_GET_SUFFIX_PARAMS_ = `+ "?" + params.Encode()`
		}

		split := strings.Split(api.Request, " ")
		parseUrl, _ := url.Parse(split[1])

		if api.FuncName == "" {
			_FUNC_NAME_ = strcase.ToCamel(path.Base(parseUrl.Path))
		} else {
			_FUNC_NAME_ = api.FuncName
		}

		tpl = strings.ReplaceAll(tpl, "_TITLE_", api.Name)
		tpl = strings.ReplaceAll(tpl, "_DESCRIPTION_", api.Description)
		tpl = strings.ReplaceAll(tpl, "_REQUEST_", api.Request)
		tpl = strings.ReplaceAll(tpl, "_SEE_", api.See)
		tpl = strings.ReplaceAll(tpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_UPLOAD_", _UPLOAD_)
		tpl = strings.ReplaceAll(tpl, "_GET_PARAMS_", _GET_PARAMS_)
		tpl = strings.ReplaceAll(tpl, "_GET_SUFFIX_PARAMS_", _GET_SUFFIX_PARAMS_)
		if _FIELD_NAME_ != "" {
			_FIELDS_ = strings.ReplaceAll(fieldTpl, "_FIELD_NAME_", _FIELD_NAME_)
		}
		tpl = strings.ReplaceAll(tpl, "_FIELDS_", _FIELDS_)
		tpl = strings.ReplaceAll(tpl, "_PAYLOAD_", _PAYLOAD_)

		funcs = append(funcs, tpl)

		tpl = strings.ReplaceAll(constTpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_API_PATH_", parseUrl.Path)

		consts = append(consts, tpl)

		// TestFunc
		_TEST_ARGS_STRUCT_ := ""
		switch {
		case strings.Contains(api.Request, "GET http"):
			_TEST_ARGS_STRUCT_ = `ctx *corporation.App , ` + _GET_PARAMS_
		case strings.Contains(api.Request, "POST http"):
			_TEST_ARGS_STRUCT_ = `ctx *corporation.App , payload []byte`
			if _GET_PARAMS_ != "" {
				_TEST_ARGS_STRUCT_ += `,` + _GET_PARAMS_
			}
		case strings.Contains(api.Request, "POST(@media"):
			_TEST_ARGS_STRUCT_ = `ctx *corporation.App , ` + _UPLOAD_ + ` string` + _PAYLOAD_ + _GET_PARAMS_
		}
		_TEST_ARGS_STRUCT_ = strings.ReplaceAll(_TEST_ARGS_STRUCT_, ",", "\n")

		_TEST_FUNC_SIGNATURE_ := ""
		_EXAMPLE_ARGS_STMT_ := ""
		if strings.TrimSpace(_TEST_ARGS_STRUCT_) != "" {
			signatures := strings.Split(_TEST_ARGS_STRUCT_, "\n")
			paramNames := []string{}
			exampleStmt := []string{}
			for _, signature := range signatures {
				signature = strings.TrimSpace(signature)
				tmp := strings.Split(signature, " ")
				//fmt.Println(tmp)
				if len(tmp[0]) > 0 {
					paramNames = append(paramNames, "tt.args."+tmp[0])

					switch tmp[1] {
					case `[]byte`:
						exampleStmt = append(exampleStmt, tmp[0]+" := []byte(\"{}\")")
					case `string`:
						exampleStmt = append(exampleStmt, tmp[0]+" := \"\"")
					case `url.Values`:
						exampleStmt = append(exampleStmt, tmp[0]+" := url.Values{}")
					}
				}
			}
			_TEST_FUNC_SIGNATURE_ = strings.Join(paramNames, ",")
			_EXAMPLE_ARGS_STMT_ = strings.Join(exampleStmt, "\n")
		}

		tpl = strings.ReplaceAll(testFuncTpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_TEST_ARGS_STRUCT_", _TEST_ARGS_STRUCT_)
		tpl = strings.ReplaceAll(tpl, "_TEST_FUNC_SIGNATURE_", _TEST_FUNC_SIGNATURE_)
		testFuncs = append(testFuncs, tpl)

		//Example
		tpl = strings.ReplaceAll(exampleFuncTpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_PACKAGE_", path.Base(group.Package))
		tpl = strings.ReplaceAll(tpl, "_TEST_FUNC_SIGNATURE_", strings.ReplaceAll(_TEST_FUNC_SIGNATURE_, "tt.args.", ""))
		tpl = strings.ReplaceAll(tpl, "_EXAMPLE_ARGS_STMT_", _EXAMPLE_ARGS_STMT_)
		exampleFuncs = append(exampleFuncs, tpl)

	}

	fileContent := fmt.Sprintf(fileTpl, path.Base(group.Package), group.Name, path.Base(group.Package), strings.Join(consts, ``), strings.Join(funcs, ``))
	filename := "./../apis/" + group.Package + "/" + path.Base(group.Package) + ".go"
	_ = os.MkdirAll(path.Dir(filename), 0644)
	ioutil.WriteFile(filename, []byte(fileContent), 0644)

	// output Test
	testFileContent := fmt.Sprintf(testFileTpl, path.Base(group.Package), strings.Join(testFuncs, ``))
	//fmt.Println(testFileContent)
	ioutil.WriteFile("./../apis/"+group.Package+"/"+path.Base(group.Package)+"_test.go", []byte(testFileContent), 0644)

	// output example
	exampleFileContent := fmt.Sprintf(exampleFileTpl, path.Base(group.Package), strings.Join(exampleFuncs, ``))
	//fmt.Println(testFileContent)
	ioutil.WriteFile("./../apis/"+group.Package+"/example_"+path.Base(group.Package)+"_test.go", []byte(exampleFileContent), 0644)

}

var constTpl = `
	api_FUNC_NAME_ = "_API_PATH_"`
var commentTpl = `
/*
_TITLE_

_DESCRIPTION_

See: _SEE_

_REQUEST_
*/`
var postFuncTpl = commentTpl + `
func _FUNC_NAME_(ctx *corporation.App, payload []byte_GET_PARAMS_) (resp []byte, err error) {
	return ctx.Client.HTTPPost(api_FUNC_NAME__GET_SUFFIX_PARAMS_, bytes.NewReader(payload), "application/json;charset=utf-8")
}
`
var getFuncTpl = commentTpl + `
func _FUNC_NAME_(ctx *corporation.App _GET_PARAMS_) (resp []byte, err error) {
	return ctx.Client.HTTPGet(api_FUNC_NAME__GET_SUFFIX_PARAMS_)
}
`
var postUploadFuncTpl = commentTpl + `
func _FUNC_NAME_(ctx *corporation.App, _UPLOAD_ string_PAYLOAD__GET_PARAMS_) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("_UPLOAD_", path.Base(_UPLOAD_))
		if err != nil {
			return
		}
		file, err := os.Open(_UPLOAD_)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

		_FIELDS_
	}()
	return ctx.Client.HTTPPost(api_FUNC_NAME__GET_SUFFIX_PARAMS_, r, m.FormDataContentType())
}
`

var fieldTpl = `
		// field
		err = m.WriteField("_FIELD_NAME_", string(payload))
		if err != nil {
			return
		}
`

var fileTpl = `// Package %s %s
package %s

const (
	%s
)
%s`

var testFileTpl = `package %s

func TestMain(m *testing.M) {
	test.Setup()
	os.Exit(m.Run())
}

%s
`

var testFuncTpl = `
func Test_FUNC_NAME_(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(api_FUNC_NAME_, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		_TEST_ARGS_STRUCT_
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp }, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := _FUNC_NAME_(_TEST_FUNC_SIGNATURE_)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("_FUNC_NAME_() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("_FUNC_NAME_() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}`

var exampleFileTpl = `package %s_test

%s
`
var exampleFuncTpl = `
func Example_FUNC_NAME_() {
	var ctx *corporation.App 

	_EXAMPLE_ARGS_STMT_
	resp, err := _PACKAGE_._FUNC_NAME_(_TEST_FUNC_SIGNATURE_)

	fmt.Println(resp, err)
}

`
