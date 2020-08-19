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

type EventChangeExternalContact struct {
	Event
	ChangeType string
}

const (
	EventTypeChangeExternalContact                       = "change_external_contact"   //企业客户变更
	EventTypeChangeExternalContactAddExternalContact     = "add_external_contact"      // 添加企业客户事件
	EventTypeChangeExternalContactEditExternalContact    = "edit_external_contact"     //编辑企业客户事件
	EventTypeChangeExternalContactAddHalfExternalContact = "add_half_external_contact" //外部联系人免验证添加成员事件
	EventTypeChangeExternalContactDelExternalContact     = "del_external_contact"      //删除企业客户事件
	EventTypeChangeExternalContactDelFollowUser          = "del_follow_user"           // 删除跟进成员事件
	EventTypeChangeExternalContactChangeExternalChat     = "change_external_chat"      //客户群变更事件
)

/**
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[sys]]></FromUserName>
    <CreateTime>1403610513</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[change_external_contact]]></Event>
    <ChangeType><![CDATA[add_external_contact]]></ChangeType>
    <UserID><![CDATA[zhangsan]]></UserID>
    <ExternalUserID><![CDATA[woAJ2GCAAAXtWyujaWJHDDGi0mAAAA]]></ExternalUserID>
    <State><![CDATA[teststate]]></State>
    <WelcomeCode><![CDATA[WELCOMECODE]]></WelcomeCode>
</xml>
*/
type EventChangeExternalContactAddExternalContact struct {
	EventChangeExternalContact
	UserID         string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
	State          string `xml:"State"`
	WelcomeCode    string `xml:"WelcomeCode"`
}

/**
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[sys]]></FromUserName>
    <CreateTime>1403610513</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[change_external_contact]]></Event>
    <ChangeType><![CDATA[edit_external_contact]]></ChangeType>
    <UserID><![CDATA[zhangsan]]></UserID>
    <ExternalUserID><![CDATA[woAJ2GCAAAXtWyujaWJHDDGi0mAAAA]]></ExternalUserID>
    <State><![CDATA[teststate]]></State>
</xml>
*/
type EventChangeExternalContactEditExternalContact struct {
	EventChangeExternalContact
	UserID         string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
	State          string `xml:"State"`
}

/**
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[sys]]></FromUserName>
    <CreateTime>1403610513</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[change_external_contact]]></Event>
    <ChangeType><![CDATA[edit_external_contact]]></ChangeType>
    <UserID><![CDATA[zhangsan]]></UserID>
    <ExternalUserID><![CDATA[woAJ2GCAAAXtWyujaWJHDDGi0mAAAA]]></ExternalUserID>
    <State><![CDATA[teststate]]></State>
</xml>
*/
type EventChangeExternalContactAddHalfExternalContact struct {
	EventChangeExternalContact
	UserID         string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
	State          string `xml:"State"`
}

/**
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[sys]]></FromUserName>
    <CreateTime>1403610513</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[change_external_contact]]></Event>
    <ChangeType><![CDATA[del_external_contact]]></ChangeType>
    <UserID><![CDATA[zhangsan]]></UserID>
    <ExternalUserID><![CDATA[woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA]]></ExternalUserID>
</xml>
*/
type EventChangeExternalContactDelExternalContact struct {
	EventChangeExternalContact
	UserID         string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
}

/*
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[sys]]></FromUserName>
    <CreateTime>1403610513</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[change_external_contact]]></Event>
    <ChangeType><![CDATA[del_follow_user]]></ChangeType>
    <UserID><![CDATA[zhangsan]]></UserID>
    <ExternalUserID><![CDATA[woAJ2GCAAAXtWyujaWJHDDGi0mACHAAA]]></ExternalUserID>
</xml>
*/
type EventChangeExternalContactDelFollowUser struct {
	EventChangeExternalContact
	UserID         string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
}

/**
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[sys]]></FromUserName>
    <CreateTime>1403610513</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[change_external_chat]]></Event>
    <ChatId><![CDATA[CHAT_ID]]></ChatId>
</xml>
*/
type EventChangeExternalContactChangeExternalChat struct {
	EventChangeExternalContact
	ChatId string `xml:"ChatId"`
}
