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

// checks if the UpdateAvatarRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAvatarRequest{}

// UpdateAvatarRequest struct for UpdateAvatarRequest
type UpdateAvatarRequest struct {
	AssetUrl *string `json:"assetUrl,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	//  
	Tags []string `json:"tags,omitempty"`
	ImageUrl *string `json:"imageUrl,omitempty"`
	ReleaseStatus *ReleaseStatus `json:"releaseStatus,omitempty"`
	Version *float32 `json:"version,omitempty"`
	UnityPackageUrl *string `json:"unityPackageUrl,omitempty"`
}

// NewUpdateAvatarRequest instantiates a new UpdateAvatarRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAvatarRequest() *UpdateAvatarRequest {
	this := UpdateAvatarRequest{}
	var releaseStatus ReleaseStatus = PUBLIC
	this.ReleaseStatus = &releaseStatus
	var version float32 = 1
	this.Version = &version
	return &this
}

// NewUpdateAvatarRequestWithDefaults instantiates a new UpdateAvatarRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAvatarRequestWithDefaults() *UpdateAvatarRequest {
	this := UpdateAvatarRequest{}
	var releaseStatus ReleaseStatus = PUBLIC
	this.ReleaseStatus = &releaseStatus
	var version float32 = 1
	this.Version = &version
	return &this
}

// GetAssetUrl returns the AssetUrl field value if set, zero value otherwise.
func (o *UpdateAvatarRequest) GetAssetUrl() string {
	if o == nil || IsNil(o.AssetUrl) {
		var ret string
		return ret
	}
	return *o.AssetUrl
}

// GetAssetUrlOk returns a tuple with the AssetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAvatarRequest) GetAssetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AssetUrl) {
		return nil, false
	}
	return o.AssetUrl, true
}

// HasAssetUrl returns a boolean if a field has been set.
func (o *UpdateAvatarRequest) HasAssetUrl() bool {
	if o != nil && !IsNil(o.AssetUrl) {
		return true
	}

	return false
}

// SetAssetUrl gets a reference to the given string and assigns it to the AssetUrl field.
func (o *UpdateAvatarRequest) SetAssetUrl(v string) {
	o.AssetUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateAvatarRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAvatarRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateAvatarRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateAvatarRequest) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateAvatarRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAvatarRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAvatarRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateAvatarRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateAvatarRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAvatarRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateAvatarRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateAvatarRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateAvatarRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAvatarRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateAvatarRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateAvatarRequest) SetTags(v []string) {
	o.Tags = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *UpdateAvatarRequest) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAvatarRequest) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *UpdateAvatarRequest) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *UpdateAvatarRequest) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetReleaseStatus returns the ReleaseStatus field value if set, zero value otherwise.
func (o *UpdateAvatarRequest) GetReleaseStatus() ReleaseStatus {
	if o == nil || IsNil(o.ReleaseStatus) {
		var ret ReleaseStatus
		return ret
	}
	return *o.ReleaseStatus
}

// GetReleaseStatusOk returns a tuple with the ReleaseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAvatarRequest) GetReleaseStatusOk() (*ReleaseStatus, bool) {
	if o == nil || IsNil(o.ReleaseStatus) {
		return nil, false
	}
	return o.ReleaseStatus, true
}

// HasReleaseStatus returns a boolean if a field has been set.
func (o *UpdateAvatarRequest) HasReleaseStatus() bool {
	if o != nil && !IsNil(o.ReleaseStatus) {
		return true
	}

	return false
}

// SetReleaseStatus gets a reference to the given ReleaseStatus and assigns it to the ReleaseStatus field.
func (o *UpdateAvatarRequest) SetReleaseStatus(v ReleaseStatus) {
	o.ReleaseStatus = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UpdateAvatarRequest) GetVersion() float32 {
	if o == nil || IsNil(o.Version) {
		var ret float32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAvatarRequest) GetVersionOk() (*float32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UpdateAvatarRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given float32 and assigns it to the Version field.
func (o *UpdateAvatarRequest) SetVersion(v float32) {
	o.Version = &v
}

// GetUnityPackageUrl returns the UnityPackageUrl field value if set, zero value otherwise.
func (o *UpdateAvatarRequest) GetUnityPackageUrl() string {
	if o == nil || IsNil(o.UnityPackageUrl) {
		var ret string
		return ret
	}
	return *o.UnityPackageUrl
}

// GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAvatarRequest) GetUnityPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.UnityPackageUrl) {
		return nil, false
	}
	return o.UnityPackageUrl, true
}

// HasUnityPackageUrl returns a boolean if a field has been set.
func (o *UpdateAvatarRequest) HasUnityPackageUrl() bool {
	if o != nil && !IsNil(o.UnityPackageUrl) {
		return true
	}

	return false
}

// SetUnityPackageUrl gets a reference to the given string and assigns it to the UnityPackageUrl field.
func (o *UpdateAvatarRequest) SetUnityPackageUrl(v string) {
	o.UnityPackageUrl = &v
}

func (o UpdateAvatarRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAvatarRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetUrl) {
		toSerialize["assetUrl"] = o.AssetUrl
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.ReleaseStatus) {
		toSerialize["releaseStatus"] = o.ReleaseStatus
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.UnityPackageUrl) {
		toSerialize["unityPackageUrl"] = o.UnityPackageUrl
	}
	return toSerialize, nil
}

type NullableUpdateAvatarRequest struct {
	value *UpdateAvatarRequest
	isSet bool
}

func (v NullableUpdateAvatarRequest) Get() *UpdateAvatarRequest {
	return v.value
}

func (v *NullableUpdateAvatarRequest) Set(val *UpdateAvatarRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAvatarRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAvatarRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAvatarRequest(val *UpdateAvatarRequest) *NullableUpdateAvatarRequest {
	return &NullableUpdateAvatarRequest{value: val, isSet: true}
}

func (v NullableUpdateAvatarRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAvatarRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


