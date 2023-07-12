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

// checks if the DynamicContentRow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DynamicContentRow{}

// DynamicContentRow struct for DynamicContentRow
type DynamicContentRow struct {
	Index *int32 `json:"index,omitempty"`
	Name string `json:"name"`
	// Usually \"ThisPlatformSupported\", but can also be other values such as \"all\" or platform specific identifiers.
	Platform string `json:"platform"`
	SortHeading string `json:"sortHeading"`
	SortOrder string `json:"sortOrder"`
	SortOwnership string `json:"sortOwnership"`
	// Tag to filter content for this row.
	Tag *string `json:"tag,omitempty"`
	// Type is not present if it is a world.
	Type *string `json:"type,omitempty"`
}

// NewDynamicContentRow instantiates a new DynamicContentRow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicContentRow(name string, platform string, sortHeading string, sortOrder string, sortOwnership string) *DynamicContentRow {
	this := DynamicContentRow{}
	this.Name = name
	this.Platform = platform
	this.SortHeading = sortHeading
	this.SortOrder = sortOrder
	this.SortOwnership = sortOwnership
	return &this
}

// NewDynamicContentRowWithDefaults instantiates a new DynamicContentRow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicContentRowWithDefaults() *DynamicContentRow {
	this := DynamicContentRow{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *DynamicContentRow) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicContentRow) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *DynamicContentRow) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *DynamicContentRow) SetIndex(v int32) {
	o.Index = &v
}

// GetName returns the Name field value
func (o *DynamicContentRow) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicContentRow) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicContentRow) SetName(v string) {
	o.Name = v
}

// GetPlatform returns the Platform field value
func (o *DynamicContentRow) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *DynamicContentRow) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *DynamicContentRow) SetPlatform(v string) {
	o.Platform = v
}

// GetSortHeading returns the SortHeading field value
func (o *DynamicContentRow) GetSortHeading() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SortHeading
}

// GetSortHeadingOk returns a tuple with the SortHeading field value
// and a boolean to check if the value has been set.
func (o *DynamicContentRow) GetSortHeadingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortHeading, true
}

// SetSortHeading sets field value
func (o *DynamicContentRow) SetSortHeading(v string) {
	o.SortHeading = v
}

// GetSortOrder returns the SortOrder field value
func (o *DynamicContentRow) GetSortOrder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *DynamicContentRow) GetSortOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *DynamicContentRow) SetSortOrder(v string) {
	o.SortOrder = v
}

// GetSortOwnership returns the SortOwnership field value
func (o *DynamicContentRow) GetSortOwnership() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SortOwnership
}

// GetSortOwnershipOk returns a tuple with the SortOwnership field value
// and a boolean to check if the value has been set.
func (o *DynamicContentRow) GetSortOwnershipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOwnership, true
}

// SetSortOwnership sets field value
func (o *DynamicContentRow) SetSortOwnership(v string) {
	o.SortOwnership = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *DynamicContentRow) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicContentRow) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *DynamicContentRow) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *DynamicContentRow) SetTag(v string) {
	o.Tag = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DynamicContentRow) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicContentRow) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DynamicContentRow) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DynamicContentRow) SetType(v string) {
	o.Type = &v
}

func (o DynamicContentRow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DynamicContentRow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	toSerialize["name"] = o.Name
	toSerialize["platform"] = o.Platform
	toSerialize["sortHeading"] = o.SortHeading
	toSerialize["sortOrder"] = o.SortOrder
	toSerialize["sortOwnership"] = o.SortOwnership
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDynamicContentRow struct {
	value *DynamicContentRow
	isSet bool
}

func (v NullableDynamicContentRow) Get() *DynamicContentRow {
	return v.value
}

func (v *NullableDynamicContentRow) Set(val *DynamicContentRow) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicContentRow) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicContentRow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicContentRow(val *DynamicContentRow) *NullableDynamicContentRow {
	return &NullableDynamicContentRow{value: val, isSet: true}
}

func (v NullableDynamicContentRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicContentRow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


