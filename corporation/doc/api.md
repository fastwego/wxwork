## API 使用

企业微信提供的接口超过 100 个，但无非就是两大类：

- GET 请求：拉取数据，例如数据统计、通讯录列表等
- POST 请求：操作数据或配置，例如设置菜单、用户管理等

每个接口的需要的参数和返回的响应格式都不一样，如果过分关注细节，那么不得不定义种类繁杂的 `struct` 类型来表示这些参数或响应

这会让开发者崩溃 ;(

### 统一参数抽象

为了保持简洁和灵活性，框架将接口参数统一定义为：
- `payload []byte` 用于 `POST` 类请求
- `params url.Values` 用于 `GET` 类请求

例如，设置菜单：
```go
payload := []byte(`
{
     "button":[
     {
           "name":"菜单",
           "sub_button":[
           {	
               "type":"view",
               "name":"搜索",
               "url":"http://www.soso.com/"
            }]
       }]
}`)
resp, err := menu.Create(ctx, payload)
```

获取用户信息：
```go
params := url.Values{}
params.Add("openid", "useropenid")
params.Add("lang","zh_CN")

resp, err := user.GetUserInfo(App, params)
```

### 统一响应抽象

同样地，接口响应的 json 字符串都用 `resp []byte` 来表示

开发者可以根据具体的业务来决定是否要解析 json 为 `struct` 类型


### 文件上传

还有一类特殊的接口，需要上传文件，例如：素材管理、客服头像设置等

这类接口，框架会提供额外的文件路径参数

例如，上传图片素材：
```go
media := "/path/to/image/file.jpg"

resp, err := material.MediaUploadImg(ctx, media)
```

### API 列表

{{#include ./apilist.md}}
