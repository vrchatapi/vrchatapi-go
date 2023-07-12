/*
VRChat API Documentation

![VRChat API Banner](https://vrchatapi.github.io/assets/img/api_banner_1500x400.png)  # Welcome to the VRChat API  Before we begin, we would like to state this is a **COMMUNITY DRIVEN PROJECT**. This means that everything you read on here was written by the community itself and is **not** officially supported by VRChat. The documentation is provided \"AS IS\", and any action you take towards VRChat is completely your own responsibility.  The documentation and additional libraries SHALL ONLY be used for applications interacting with VRChat's API in accordance with their [Terms of Service](https://hello.vrchat.com/legal) and [Community Guidelines](https://hello.vrchat.com/community-guidelines), and MUST NOT be used for modifying the client, \"avatar ripping\", or other illegal activities. Malicious usage or spamming the API may result in account termination. Certain parts of the API are also more sensitive than others, for example moderation, so please tread extra carefully and read the warnings when present.  ![Tupper Policy on API](https://i.imgur.com/yLlW7Ok.png)  Finally, use of the API using applications other than the approved methods (website, VRChat application, Unity SDK) is not officially supported. VRChat provides no guarantee or support for external applications using the API. Access to API endpoints may break **at any time, without notice**. Therefore, please **do not ping** VRChat Staff in the VRChat Discord if you are having API problems, as they do not provide API support. We will make a best effort in keeping this documentation and associated language libraries up to date, but things might be outdated or missing. If you find that something is no longer valid, please contact us on Discord or [create an issue](https://github.com/vrchatapi/specification/issues) and tell us so we can fix it.  # Getting Started  The VRChat API can be used to programmatically retrieve or update information regarding your profile, friends, avatars, worlds and more. The API consists of two parts, \"Photon\" which is only used in-game, and the \"Web API\" which is used by both the game and the website. This documentation focuses only on the Web API.  The API is designed around the REST ideology, providing semi-simple and usually predictable URIs to access and modify objects. Requests support standard HTTP methods like GET, PUT, POST, and DELETE and standard status codes. Response bodies are always UTF-8 encoded JSON objects, unless explicitly documented otherwise.  <div class=\"callout callout-error\">   <strong>🛑 Warning! Do not touch Photon!</strong><br>   Photon is only used by the in-game client and should <b>not</b> be touched. Doing so may result in permanent account termination. </div>  <div class=\"callout callout-info\">   <strong>ℹ️ API Key and Authentication</strong><br>   The API Key has always been the same and is currently <code>JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26</code>.   Read <a href=\"#tag--authentication\">Authentication</a> for how to log in. </div>  # Using the API  For simply exploring what the API can do it is strongly recommended to download [Insomnia](https://insomnia.rest/download), a free and open-source API client that's great for sending requests to the API in an orderly fashion. Insomnia allows you to send data in the format that's required for VRChat's API. It is also possible to try out the API in your browser, by first logging in at [vrchat.com/home](https://vrchat.com/home/) and then going to [vrchat.com/api/1/auth/user](https://vrchat.com/api/1/auth/user), but the information will be much harder to work with.  For more permanent operation such as software development it is instead recommended to use one of the existing language SDKs. This community project maintains API libraries in several languages, which allows you to interact with the API with simple function calls rather than having to implement the HTTP protocol yourself. Most of these libraries are automatically generated from the API specification, sometimes with additional helpful wrapper code to make usage easier. This allows them to be almost automatically updated and expanded upon as soon as a new feature is introduced in the specification itself. The libraries can be found on [GitHub](https://github.com/vrchatapi) or following:  * [NodeJS (JavaScript)](https://www.npmjs.com/package/vrchat) * [Dart](https://pub.dev/packages/vrchat_dart) * [Rust](https://crates.io/crates/vrchatapi) * [C#](https://github.com/vrchatapi/vrchatapi-csharp) * [Python](https://github.com/vrchatapi/vrchatapi-python)  # Pagination  Most endpoints enforce pagination, meaning they will only return 10 entries by default, and never more than 100.<br> Using both the limit and offset parameters allows you to easily paginate through a large number of objects.  | Query Parameter | Type | Description | | ----------|--|------- | | `n` | integer  | The number of objects to return. This value often defaults to 10. Highest limit is always 100.| | `offset` | integer  | A zero-based offset from the default object sorting.|  If a request returns fewer objects than the `limit` parameter, there are no more items available to return.  # Contribution  Do you want to get involved in the documentation effort? Do you want to help improve one of the language API libraries? This project is an [OPEN Open Source Project](https://openopensource.org)! This means that individuals making significant and valuable contributions are given commit-access to the project. It also means we are very open and welcoming of new people making contributions, unlike some more guarded open-source projects.  [![Discord](https://img.shields.io/static/v1?label=vrchatapi&message=discord&color=blueviolet&style=for-the-badge)](https://discord.gg/qjZE9C9fkB)

API version: 1.12.0
Contact: vrchatapi.lpv0t@aries.fyi
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vrchatapi

import (
	"encoding/json"
	"time"
)

// checks if the SentNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SentNotification{}

// SentNotification 
type SentNotification struct {
	CreatedAt time.Time `json:"created_at"`
	// **NOTICE:** This is not a JSON object, this is a json **encoded** object, meaning you have to json-de-encode to get the NotificationDetail object depending on the NotificationType.
	Details string `json:"details"`
	Id string `json:"id"`
	// 
	Message string `json:"message"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	RecieverUserId string `json:"recieverUserId"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	SenderUserId string `json:"senderUserId"`
	// -| **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429).
	// Deprecated
	SenderUsername *string `json:"senderUsername,omitempty"`
	Type NotificationType `json:"type"`
}

// NewSentNotification instantiates a new SentNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSentNotification(createdAt time.Time, details string, id string, message string, recieverUserId string, senderUserId string, type_ NotificationType) *SentNotification {
	this := SentNotification{}
	this.CreatedAt = createdAt
	this.Details = details
	this.Id = id
	this.Message = message
	this.RecieverUserId = recieverUserId
	this.SenderUserId = senderUserId
	this.Type = type_
	return &this
}

// NewSentNotificationWithDefaults instantiates a new SentNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSentNotificationWithDefaults() *SentNotification {
	this := SentNotification{}
	var details string = "{}"
	this.Details = details
	var type_ NotificationType = FRIEND_REQUEST
	this.Type = type_
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *SentNotification) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SentNotification) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SentNotification) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDetails returns the Details field value
func (o *SentNotification) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SentNotification) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SentNotification) SetDetails(v string) {
	o.Details = v
}

// GetId returns the Id field value
func (o *SentNotification) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SentNotification) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SentNotification) SetId(v string) {
	o.Id = v
}

// GetMessage returns the Message field value
func (o *SentNotification) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SentNotification) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SentNotification) SetMessage(v string) {
	o.Message = v
}

// GetRecieverUserId returns the RecieverUserId field value
func (o *SentNotification) GetRecieverUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecieverUserId
}

// GetRecieverUserIdOk returns a tuple with the RecieverUserId field value
// and a boolean to check if the value has been set.
func (o *SentNotification) GetRecieverUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecieverUserId, true
}

// SetRecieverUserId sets field value
func (o *SentNotification) SetRecieverUserId(v string) {
	o.RecieverUserId = v
}

// GetSenderUserId returns the SenderUserId field value
func (o *SentNotification) GetSenderUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderUserId
}

// GetSenderUserIdOk returns a tuple with the SenderUserId field value
// and a boolean to check if the value has been set.
func (o *SentNotification) GetSenderUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderUserId, true
}

// SetSenderUserId sets field value
func (o *SentNotification) SetSenderUserId(v string) {
	o.SenderUserId = v
}

// GetSenderUsername returns the SenderUsername field value if set, zero value otherwise.
// Deprecated
func (o *SentNotification) GetSenderUsername() string {
	if o == nil || IsNil(o.SenderUsername) {
		var ret string
		return ret
	}
	return *o.SenderUsername
}

// GetSenderUsernameOk returns a tuple with the SenderUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SentNotification) GetSenderUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.SenderUsername) {
		return nil, false
	}
	return o.SenderUsername, true
}

// HasSenderUsername returns a boolean if a field has been set.
func (o *SentNotification) HasSenderUsername() bool {
	if o != nil && !IsNil(o.SenderUsername) {
		return true
	}

	return false
}

// SetSenderUsername gets a reference to the given string and assigns it to the SenderUsername field.
// Deprecated
func (o *SentNotification) SetSenderUsername(v string) {
	o.SenderUsername = &v
}

// GetType returns the Type field value
func (o *SentNotification) GetType() NotificationType {
	if o == nil {
		var ret NotificationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SentNotification) GetTypeOk() (*NotificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SentNotification) SetType(v NotificationType) {
	o.Type = v
}

func (o SentNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SentNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["details"] = o.Details
	toSerialize["id"] = o.Id
	toSerialize["message"] = o.Message
	toSerialize["recieverUserId"] = o.RecieverUserId
	toSerialize["senderUserId"] = o.SenderUserId
	if !IsNil(o.SenderUsername) {
		toSerialize["senderUsername"] = o.SenderUsername
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableSentNotification struct {
	value *SentNotification
	isSet bool
}

func (v NullableSentNotification) Get() *SentNotification {
	return v.value
}

func (v *NullableSentNotification) Set(val *SentNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableSentNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableSentNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSentNotification(val *SentNotification) *NullableSentNotification {
	return &NullableSentNotification{value: val, isSet: true}
}

func (v NullableSentNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSentNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


