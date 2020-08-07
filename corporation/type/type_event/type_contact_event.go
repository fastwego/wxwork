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

package type_event

type EventChangeContact struct {
	Event
	ChangeType string
}

const (
	EventTypeChangeContact = "change_contact" // 通讯录变更

	EventTypeChangeContactCreateUser = "create_user" // 新增成员事件
	EventTypeChangeContactUpdateUser = "update_user" // 更新成员事件
	EventTypeChangeContactDeleteUser ="delete_user" // 删除成员事件

	EventTypeChangeContactCreateParty = "create_party" // 新增部门事件
	EventTypeChangeContactUpdateParty = "update_party" // 更新部门事件
	EventTypeChangeContactDeleteParty = "delete_party" // 删除部门事件

	EventTypeChangeContactUpdateTag = "update_tag" // 标签成员变更事件

	EventTypeBatchJobResult = "batch_job_result" // 异步任务完成通知
)

/*
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[sys]]></FromUserName>
    <CreateTime>1403610513</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[change_contact]]></Event>
    <ChangeType>create_user</ChangeType>
    <UserID><![CDATA[zhangsan]]></UserID>
    <Name><![CDATA[张三]]></Name>
    <Department><![CDATA[1,2,3]]></Department>
    <IsLeaderInDept><![CDATA[1,0,0]]></IsLeaderInDept>
    <Position><![CDATA[产品经理]]></Position>
    <Mobile>13800000000</Mobile>
    <Gender>1</Gender>
    <Email><![CDATA[zhangsan@gzdev.com]]></Email>
    <Status>1</Status>
    <Avatar><![CDATA[http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/0]]></Avatar>
    <Alias><![CDATA[zhangsan]]></Alias>
    <Telephone><![CDATA[020-123456]]></Telephone>
    <Address><![CDATA[广州市]]></Address>
    <ExtAttr>
        <Item>
        <Name><![CDATA[爱好]]></Name>
        <Type>0</Type>
        <Text>
            <Value><![CDATA[旅游]]></Value>
        </Text>
        </Item>
        <Item>
        <Name><![CDATA[卡号]]></Name>
        <Type>1</Type>
        <Web>
            <Title><![CDATA[企业微信]]></Title>
            <Url><![CDATA[https://work.weixin.qq.com]]></Url>
        </Web>
        </Item>
    </ExtAttr>
</xml>


*/
type EventChangeContactCreateUser struct {
	EventChangeContact
	UserID         string   `xml:"UserID"`
	Name           string   `xml:"Name"`
	Department     string   `xml:"Department"`
	IsLeaderInDept string   `xml:"IsLeaderInDept"`
	Position       string   `xml:"Position"`
	Mobile         string   `xml:"Mobile"`
	Gender         string   `xml:"Gender"`
	Email          string   `xml:"Email"`
	Status         string   `xml:"Status"`
	Avatar         string   `xml:"Avatar"`
	Alias          string   `xml:"Alias"`
	Telephone      string   `xml:"Telephone"`
	Address        string   `xml:"Address"`
	ExtAttr        struct {
		Text string `xml:",chardata"`
		Item []struct {
			Chardata string `xml:",chardata"`
			Name     string `xml:"Name"`
			Type     string `xml:"Type"`
			Text     struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value"`
			} `xml:"Text"`
			Web struct {
				Text  string `xml:",chardata"`
				Title string `xml:"Title"`
				URL   string `xml:"Url"`
			} `xml:"Web"`
		} `xml:"Item"`
	} `xml:"ExtAttr"`
}

/*
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[sys]]></FromUserName>
    <CreateTime>1403610513</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[change_contact]]></Event>
    <ChangeType>update_user</ChangeType>
    <UserID><![CDATA[zhangsan]]></UserID>
    <NewUserID><![CDATA[zhangsan001]]></NewUserID>
    <Name><![CDATA[张三]]></Name>
    <Department><![CDATA[1,2,3]]></Department>
    <IsLeaderInDept><![CDATA[1,0,0]]></IsLeaderInDept>
    <Position><![CDATA[产品经理]]></Position>
    <Mobile>13800000000</Mobile>
    <Gender>1</Gender>
    <Email><![CDATA[zhangsan@gzdev.com]]></Email>
    <Status>1</Status>
    <Avatar><![CDATA[http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/0]]></Avatar>
    <Alias><![CDATA[zhangsan]]></Alias>
    <Telephone><![CDATA[020-123456]]></Telephone>
    <Address><![CDATA[广州市]]></Address>
    <ExtAttr>
        <Item>
        <Name><![CDATA[爱好]]></Name>
        <Type>0</Type>
        <Text>
            <Value><![CDATA[旅游]]></Value>
        </Text>
        </Item>
        <Item>
        <Name><![CDATA[卡号]]></Name>
        <Type>1</Type>
        <Web>
            <Title><![CDATA[企业微信]]></Title>
            <Url><![CDATA[https://work.weixin.qq.com]]></Url>
        </Web>
        </Item>
    </ExtAttr>
</xml>
 */

type EventChangeContactUpdateUser struct {
	EventChangeContact
	ChangeType     string   `xml:"ChangeType"`
	UserID         string   `xml:"UserID"`
	NewUserID      string   `xml:"NewUserID"`
	Name           string   `xml:"Name"`
	Department     string   `xml:"Department"`
	IsLeaderInDept string   `xml:"IsLeaderInDept"`
	Position       string   `xml:"Position"`
	Mobile         string   `xml:"Mobile"`
	Gender         string   `xml:"Gender"`
	Email          string   `xml:"Email"`
	Status         string   `xml:"Status"`
	Avatar         string   `xml:"Avatar"`
	Alias          string   `xml:"Alias"`
	Telephone      string   `xml:"Telephone"`
	Address        string   `xml:"Address"`
	ExtAttr        struct {
		Text string `xml:",chardata"`
		Item []struct {
			Chardata string `xml:",chardata"`
			Name     string `xml:"Name"`
			Type     string `xml:"Type"`
			Text     struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value"`
			} `xml:"Text"`
			Web struct {
				Text  string `xml:",chardata"`
				Title string `xml:"Title"`
				URL   string `xml:"Url"`
			} `xml:"Web"`
		} `xml:"Item"`
	} `xml:"ExtAttr"`
}

/*
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[sys]]></FromUserName>
    <CreateTime>1403610513</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[change_contact]]></Event>
    <ChangeType>delete_user</ChangeType>
    <UserID><![CDATA[zhangsan]]></UserID>
</xml>
 */
type EventChangeContactDeleteUser struct {
	EventChangeContact
	ChangeType   string   `xml:"ChangeType"`
	UserID       string   `xml:"UserID"`
}

/*
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[sys]]></FromUserName>
    <CreateTime>1403610513</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[change_contact]]></Event>
    <ChangeType>create_party</ChangeType>
    <Id>2</Id>
    <Name><![CDATA[张三]]></Name>
    <ParentId><![CDATA[1]]></ParentId>
    <Order>1</Order>
</xml>

 */
type EventChangeContactCreateParty struct {
	EventChangeContact
	ID           string   `xml:"Id"`
	Name         string   `xml:"Name"`
	ParentId     string   `xml:"ParentId"`
	Order        string   `xml:"Order"`
}


/*
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[sys]]></FromUserName>
    <CreateTime>1403610513</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[change_contact]]></Event>
    <ChangeType>update_party</ChangeType>
    <Id>2</Id>
    <Name><![CDATA[张三]]></Name>
    <ParentId><![CDATA[1]]></ParentId>
</xml>

*/
type EventChangeContactUpdateParty struct {
	EventChangeContact
	ID           string   `xml:"Id"`
	Name         string   `xml:"Name"`
	ParentId     string   `xml:"ParentId"`
	Order        string   `xml:"Order"`
}

/*
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[sys]]></FromUserName>
    <CreateTime>1403610513</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[change_contact]]></Event>
    <ChangeType>delete_party</ChangeType>
    <Id>2</Id>
</xml>
*/
type EventChangeContactDeleteParty struct {
	EventChangeContact
	ID           string   `xml:"Id"`

}

/*
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[sys]]></FromUserName>
    <CreateTime>1403610513</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[change_contact]]></Event>
    <ChangeType><![CDATA[update_tag]]></ChangeType>
    <TagId>1</TagId>
    <AddUserItems><![CDATA[zhangsan,lisi]]></AddUserItems>
    <DelUserItems><![CDATA[zhangsan1,lisi1]]></DelUserItems>
    <AddPartyItems><![CDATA[1,2]]></AddPartyItems>
    <DelPartyItems><![CDATA[3,4]]></DelPartyItems>
</xml>
 */
type EventChangeContactUpdateTag struct {
	EventChangeContact
	TagId         string   `xml:"TagId"`
	AddUserItems  string   `xml:"AddUserItems"`
	DelUserItems  string   `xml:"DelUserItems"`
	AddPartyItems string   `xml:"AddPartyItems"`
	DelPartyItems string   `xml:"DelPartyItems"`
}

/*
<xml><ToUserName><![CDATA[wx28dbb14e3720FAKE]]></ToUserName>
<FromUserName><![CDATA[FromUser]]></FromUserName>
<CreateTime>1425284517</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[batch_job_result]]></Event>
<BatchJob><JobId><![CDATA[S0MrnndvRG5fadSlLwiBqiDDbM143UqTmKP3152FZk4]]></JobId>
<JobType><![CDATA[sync_user]]></JobType>
<ErrCode>0</ErrCode>
<ErrMsg><![CDATA[ok]]></ErrMsg>
</BatchJob>
</xml>
 */
type EventBatchJobResult struct {
	Event
	BatchJob     struct {
		Text    string `xml:",chardata"`
		JobId   string `xml:"JobId"`
		JobType string `xml:"JobType"`
		ErrCode string `xml:"ErrCode"`
		ErrMsg  string `xml:"ErrMsg"`
	} `xml:"BatchJob"`
}
