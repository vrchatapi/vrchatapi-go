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

// checks if the UpdateGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGroupRequest{}

// UpdateGroupRequest struct for UpdateGroupRequest
type UpdateGroupRequest struct {
	Name *string `json:"name,omitempty"`
	ShortCode *string `json:"shortCode,omitempty"`
	Description *string `json:"description,omitempty"`
	JoinState *GroupJoinState `json:"joinState,omitempty"`
	IconId NullableString `json:"iconId,omitempty"`
	BannerId NullableString `json:"bannerId,omitempty"`
	// 3 letter language code
	Languages []string `json:"languages,omitempty"`
	Links []string `json:"links,omitempty"`
	Rules *string `json:"rules,omitempty"`
	//  
	Tags []string `json:"tags,omitempty"`
}

// NewUpdateGroupRequest instantiates a new UpdateGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGroupRequest() *UpdateGroupRequest {
	this := UpdateGroupRequest{}
	var joinState GroupJoinState = OPEN
	this.JoinState = &joinState
	return &this
}

// NewUpdateGroupRequestWithDefaults instantiates a new UpdateGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGroupRequestWithDefaults() *UpdateGroupRequest {
	this := UpdateGroupRequest{}
	var joinState GroupJoinState = OPEN
	this.JoinState = &joinState
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetShortCode returns the ShortCode field value if set, zero value otherwise.
func (o *UpdateGroupRequest) GetShortCode() string {
	if o == nil || IsNil(o.ShortCode) {
		var ret string
		return ret
	}
	return *o.ShortCode
}

// GetShortCodeOk returns a tuple with the ShortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupRequest) GetShortCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ShortCode) {
		return nil, false
	}
	return o.ShortCode, true
}

// HasShortCode returns a boolean if a field has been set.
func (o *UpdateGroupRequest) HasShortCode() bool {
	if o != nil && !IsNil(o.ShortCode) {
		return true
	}

	return false
}

// SetShortCode gets a reference to the given string and assigns it to the ShortCode field.
func (o *UpdateGroupRequest) SetShortCode(v string) {
	o.ShortCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetJoinState returns the JoinState field value if set, zero value otherwise.
func (o *UpdateGroupRequest) GetJoinState() GroupJoinState {
	if o == nil || IsNil(o.JoinState) {
		var ret GroupJoinState
		return ret
	}
	return *o.JoinState
}

// GetJoinStateOk returns a tuple with the JoinState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupRequest) GetJoinStateOk() (*GroupJoinState, bool) {
	if o == nil || IsNil(o.JoinState) {
		return nil, false
	}
	return o.JoinState, true
}

// HasJoinState returns a boolean if a field has been set.
func (o *UpdateGroupRequest) HasJoinState() bool {
	if o != nil && !IsNil(o.JoinState) {
		return true
	}

	return false
}

// SetJoinState gets a reference to the given GroupJoinState and assigns it to the JoinState field.
func (o *UpdateGroupRequest) SetJoinState(v GroupJoinState) {
	o.JoinState = &v
}

// GetIconId returns the IconId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGroupRequest) GetIconId() string {
	if o == nil || IsNil(o.IconId.Get()) {
		var ret string
		return ret
	}
	return *o.IconId.Get()
}

// GetIconIdOk returns a tuple with the IconId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGroupRequest) GetIconIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconId.Get(), o.IconId.IsSet()
}

// HasIconId returns a boolean if a field has been set.
func (o *UpdateGroupRequest) HasIconId() bool {
	if o != nil && o.IconId.IsSet() {
		return true
	}

	return false
}

// SetIconId gets a reference to the given NullableString and assigns it to the IconId field.
func (o *UpdateGroupRequest) SetIconId(v string) {
	o.IconId.Set(&v)
}
// SetIconIdNil sets the value for IconId to be an explicit nil
func (o *UpdateGroupRequest) SetIconIdNil() {
	o.IconId.Set(nil)
}

// UnsetIconId ensures that no value is present for IconId, not even an explicit nil
func (o *UpdateGroupRequest) UnsetIconId() {
	o.IconId.Unset()
}

// GetBannerId returns the BannerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGroupRequest) GetBannerId() string {
	if o == nil || IsNil(o.BannerId.Get()) {
		var ret string
		return ret
	}
	return *o.BannerId.Get()
}

// GetBannerIdOk returns a tuple with the BannerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGroupRequest) GetBannerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BannerId.Get(), o.BannerId.IsSet()
}

// HasBannerId returns a boolean if a field has been set.
func (o *UpdateGroupRequest) HasBannerId() bool {
	if o != nil && o.BannerId.IsSet() {
		return true
	}

	return false
}

// SetBannerId gets a reference to the given NullableString and assigns it to the BannerId field.
func (o *UpdateGroupRequest) SetBannerId(v string) {
	o.BannerId.Set(&v)
}
// SetBannerIdNil sets the value for BannerId to be an explicit nil
func (o *UpdateGroupRequest) SetBannerIdNil() {
	o.BannerId.Set(nil)
}

// UnsetBannerId ensures that no value is present for BannerId, not even an explicit nil
func (o *UpdateGroupRequest) UnsetBannerId() {
	o.BannerId.Unset()
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *UpdateGroupRequest) GetLanguages() []string {
	if o == nil || IsNil(o.Languages) {
		var ret []string
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupRequest) GetLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *UpdateGroupRequest) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []string and assigns it to the Languages field.
func (o *UpdateGroupRequest) SetLanguages(v []string) {
	o.Languages = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UpdateGroupRequest) GetLinks() []string {
	if o == nil || IsNil(o.Links) {
		var ret []string
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupRequest) GetLinksOk() ([]string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UpdateGroupRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []string and assigns it to the Links field.
func (o *UpdateGroupRequest) SetLinks(v []string) {
	o.Links = v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *UpdateGroupRequest) GetRules() string {
	if o == nil || IsNil(o.Rules) {
		var ret string
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupRequest) GetRulesOk() (*string, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *UpdateGroupRequest) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given string and assigns it to the Rules field.
func (o *UpdateGroupRequest) SetRules(v string) {
	o.Rules = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateGroupRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateGroupRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateGroupRequest) SetTags(v []string) {
	o.Tags = v
}

func (o UpdateGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ShortCode) {
		toSerialize["shortCode"] = o.ShortCode
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.JoinState) {
		toSerialize["joinState"] = o.JoinState
	}
	if o.IconId.IsSet() {
		toSerialize["iconId"] = o.IconId.Get()
	}
	if o.BannerId.IsSet() {
		toSerialize["bannerId"] = o.BannerId.Get()
	}
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableUpdateGroupRequest struct {
	value *UpdateGroupRequest
	isSet bool
}

func (v NullableUpdateGroupRequest) Get() *UpdateGroupRequest {
	return v.value
}

func (v *NullableUpdateGroupRequest) Set(val *UpdateGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGroupRequest(val *UpdateGroupRequest) *NullableUpdateGroupRequest {
	return &NullableUpdateGroupRequest{value: val, isSet: true}
}

func (v NullableUpdateGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


