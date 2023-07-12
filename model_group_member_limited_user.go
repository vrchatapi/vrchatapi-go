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
)

// checks if the GroupMemberLimitedUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMemberLimitedUser{}

// GroupMemberLimitedUser Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user.
type GroupMemberLimitedUser struct {
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id *string `json:"id,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ThumbnailUrl *string `json:"thumbnailUrl,omitempty"`
	IconUrl *string `json:"iconUrl,omitempty"`
}

// NewGroupMemberLimitedUser instantiates a new GroupMemberLimitedUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMemberLimitedUser() *GroupMemberLimitedUser {
	this := GroupMemberLimitedUser{}
	return &this
}

// NewGroupMemberLimitedUserWithDefaults instantiates a new GroupMemberLimitedUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMemberLimitedUserWithDefaults() *GroupMemberLimitedUser {
	this := GroupMemberLimitedUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupMemberLimitedUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMemberLimitedUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupMemberLimitedUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupMemberLimitedUser) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GroupMemberLimitedUser) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMemberLimitedUser) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GroupMemberLimitedUser) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GroupMemberLimitedUser) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise.
func (o *GroupMemberLimitedUser) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMemberLimitedUser) GetThumbnailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailUrl) {
		return nil, false
	}
	return o.ThumbnailUrl, true
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *GroupMemberLimitedUser) HasThumbnailUrl() bool {
	if o != nil && !IsNil(o.ThumbnailUrl) {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given string and assigns it to the ThumbnailUrl field.
func (o *GroupMemberLimitedUser) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *GroupMemberLimitedUser) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMemberLimitedUser) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *GroupMemberLimitedUser) HasIconUrl() bool {
	if o != nil && !IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *GroupMemberLimitedUser) SetIconUrl(v string) {
	o.IconUrl = &v
}

func (o GroupMemberLimitedUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMemberLimitedUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.ThumbnailUrl) {
		toSerialize["thumbnailUrl"] = o.ThumbnailUrl
	}
	if !IsNil(o.IconUrl) {
		toSerialize["iconUrl"] = o.IconUrl
	}
	return toSerialize, nil
}

type NullableGroupMemberLimitedUser struct {
	value *GroupMemberLimitedUser
	isSet bool
}

func (v NullableGroupMemberLimitedUser) Get() *GroupMemberLimitedUser {
	return v.value
}

func (v *NullableGroupMemberLimitedUser) Set(val *GroupMemberLimitedUser) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMemberLimitedUser) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMemberLimitedUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMemberLimitedUser(val *GroupMemberLimitedUser) *NullableGroupMemberLimitedUser {
	return &NullableGroupMemberLimitedUser{value: val, isSet: true}
}

func (v NullableGroupMemberLimitedUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMemberLimitedUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


