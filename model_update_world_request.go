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

// checks if the UpdateWorldRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateWorldRequest{}

// UpdateWorldRequest struct for UpdateWorldRequest
type UpdateWorldRequest struct {
	AssetUrl *string `json:"assetUrl,omitempty"`
	AssetVersion *string `json:"assetVersion,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId *string `json:"authorId,omitempty"`
	AuthorName *string `json:"authorName,omitempty"`
	Capacity *int32 `json:"capacity,omitempty"`
	Description *string `json:"description,omitempty"`
	ImageUrl *string `json:"imageUrl,omitempty"`
	Name *string `json:"name,omitempty"`
	// This can be `standalonewindows` or `android`, but can also pretty much be any random Unity verison such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	Platform *string `json:"platform,omitempty"`
	ReleaseStatus *ReleaseStatus `json:"releaseStatus,omitempty"`
	//  
	Tags []string `json:"tags,omitempty"`
	UnityPackageUrl *string `json:"unityPackageUrl,omitempty"`
	UnityVersion *string `json:"unityVersion,omitempty"`
}

// NewUpdateWorldRequest instantiates a new UpdateWorldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWorldRequest() *UpdateWorldRequest {
	this := UpdateWorldRequest{}
	var releaseStatus ReleaseStatus = PUBLIC
	this.ReleaseStatus = &releaseStatus
	var unityVersion string = "5.3.4p1"
	this.UnityVersion = &unityVersion
	return &this
}

// NewUpdateWorldRequestWithDefaults instantiates a new UpdateWorldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWorldRequestWithDefaults() *UpdateWorldRequest {
	this := UpdateWorldRequest{}
	var releaseStatus ReleaseStatus = PUBLIC
	this.ReleaseStatus = &releaseStatus
	var unityVersion string = "5.3.4p1"
	this.UnityVersion = &unityVersion
	return &this
}

// GetAssetUrl returns the AssetUrl field value if set, zero value otherwise.
func (o *UpdateWorldRequest) GetAssetUrl() string {
	if o == nil || IsNil(o.AssetUrl) {
		var ret string
		return ret
	}
	return *o.AssetUrl
}

// GetAssetUrlOk returns a tuple with the AssetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorldRequest) GetAssetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AssetUrl) {
		return nil, false
	}
	return o.AssetUrl, true
}

// HasAssetUrl returns a boolean if a field has been set.
func (o *UpdateWorldRequest) HasAssetUrl() bool {
	if o != nil && !IsNil(o.AssetUrl) {
		return true
	}

	return false
}

// SetAssetUrl gets a reference to the given string and assigns it to the AssetUrl field.
func (o *UpdateWorldRequest) SetAssetUrl(v string) {
	o.AssetUrl = &v
}

// GetAssetVersion returns the AssetVersion field value if set, zero value otherwise.
func (o *UpdateWorldRequest) GetAssetVersion() string {
	if o == nil || IsNil(o.AssetVersion) {
		var ret string
		return ret
	}
	return *o.AssetVersion
}

// GetAssetVersionOk returns a tuple with the AssetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorldRequest) GetAssetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AssetVersion) {
		return nil, false
	}
	return o.AssetVersion, true
}

// HasAssetVersion returns a boolean if a field has been set.
func (o *UpdateWorldRequest) HasAssetVersion() bool {
	if o != nil && !IsNil(o.AssetVersion) {
		return true
	}

	return false
}

// SetAssetVersion gets a reference to the given string and assigns it to the AssetVersion field.
func (o *UpdateWorldRequest) SetAssetVersion(v string) {
	o.AssetVersion = &v
}

// GetAuthorId returns the AuthorId field value if set, zero value otherwise.
func (o *UpdateWorldRequest) GetAuthorId() string {
	if o == nil || IsNil(o.AuthorId) {
		var ret string
		return ret
	}
	return *o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorldRequest) GetAuthorIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorId) {
		return nil, false
	}
	return o.AuthorId, true
}

// HasAuthorId returns a boolean if a field has been set.
func (o *UpdateWorldRequest) HasAuthorId() bool {
	if o != nil && !IsNil(o.AuthorId) {
		return true
	}

	return false
}

// SetAuthorId gets a reference to the given string and assigns it to the AuthorId field.
func (o *UpdateWorldRequest) SetAuthorId(v string) {
	o.AuthorId = &v
}

// GetAuthorName returns the AuthorName field value if set, zero value otherwise.
func (o *UpdateWorldRequest) GetAuthorName() string {
	if o == nil || IsNil(o.AuthorName) {
		var ret string
		return ret
	}
	return *o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorldRequest) GetAuthorNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorName) {
		return nil, false
	}
	return o.AuthorName, true
}

// HasAuthorName returns a boolean if a field has been set.
func (o *UpdateWorldRequest) HasAuthorName() bool {
	if o != nil && !IsNil(o.AuthorName) {
		return true
	}

	return false
}

// SetAuthorName gets a reference to the given string and assigns it to the AuthorName field.
func (o *UpdateWorldRequest) SetAuthorName(v string) {
	o.AuthorName = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *UpdateWorldRequest) GetCapacity() int32 {
	if o == nil || IsNil(o.Capacity) {
		var ret int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorldRequest) GetCapacityOk() (*int32, bool) {
	if o == nil || IsNil(o.Capacity) {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *UpdateWorldRequest) HasCapacity() bool {
	if o != nil && !IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int32 and assigns it to the Capacity field.
func (o *UpdateWorldRequest) SetCapacity(v int32) {
	o.Capacity = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateWorldRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorldRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateWorldRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateWorldRequest) SetDescription(v string) {
	o.Description = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *UpdateWorldRequest) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorldRequest) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *UpdateWorldRequest) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *UpdateWorldRequest) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateWorldRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorldRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateWorldRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateWorldRequest) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *UpdateWorldRequest) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorldRequest) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *UpdateWorldRequest) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *UpdateWorldRequest) SetPlatform(v string) {
	o.Platform = &v
}

// GetReleaseStatus returns the ReleaseStatus field value if set, zero value otherwise.
func (o *UpdateWorldRequest) GetReleaseStatus() ReleaseStatus {
	if o == nil || IsNil(o.ReleaseStatus) {
		var ret ReleaseStatus
		return ret
	}
	return *o.ReleaseStatus
}

// GetReleaseStatusOk returns a tuple with the ReleaseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorldRequest) GetReleaseStatusOk() (*ReleaseStatus, bool) {
	if o == nil || IsNil(o.ReleaseStatus) {
		return nil, false
	}
	return o.ReleaseStatus, true
}

// HasReleaseStatus returns a boolean if a field has been set.
func (o *UpdateWorldRequest) HasReleaseStatus() bool {
	if o != nil && !IsNil(o.ReleaseStatus) {
		return true
	}

	return false
}

// SetReleaseStatus gets a reference to the given ReleaseStatus and assigns it to the ReleaseStatus field.
func (o *UpdateWorldRequest) SetReleaseStatus(v ReleaseStatus) {
	o.ReleaseStatus = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateWorldRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorldRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateWorldRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateWorldRequest) SetTags(v []string) {
	o.Tags = v
}

// GetUnityPackageUrl returns the UnityPackageUrl field value if set, zero value otherwise.
func (o *UpdateWorldRequest) GetUnityPackageUrl() string {
	if o == nil || IsNil(o.UnityPackageUrl) {
		var ret string
		return ret
	}
	return *o.UnityPackageUrl
}

// GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorldRequest) GetUnityPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.UnityPackageUrl) {
		return nil, false
	}
	return o.UnityPackageUrl, true
}

// HasUnityPackageUrl returns a boolean if a field has been set.
func (o *UpdateWorldRequest) HasUnityPackageUrl() bool {
	if o != nil && !IsNil(o.UnityPackageUrl) {
		return true
	}

	return false
}

// SetUnityPackageUrl gets a reference to the given string and assigns it to the UnityPackageUrl field.
func (o *UpdateWorldRequest) SetUnityPackageUrl(v string) {
	o.UnityPackageUrl = &v
}

// GetUnityVersion returns the UnityVersion field value if set, zero value otherwise.
func (o *UpdateWorldRequest) GetUnityVersion() string {
	if o == nil || IsNil(o.UnityVersion) {
		var ret string
		return ret
	}
	return *o.UnityVersion
}

// GetUnityVersionOk returns a tuple with the UnityVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorldRequest) GetUnityVersionOk() (*string, bool) {
	if o == nil || IsNil(o.UnityVersion) {
		return nil, false
	}
	return o.UnityVersion, true
}

// HasUnityVersion returns a boolean if a field has been set.
func (o *UpdateWorldRequest) HasUnityVersion() bool {
	if o != nil && !IsNil(o.UnityVersion) {
		return true
	}

	return false
}

// SetUnityVersion gets a reference to the given string and assigns it to the UnityVersion field.
func (o *UpdateWorldRequest) SetUnityVersion(v string) {
	o.UnityVersion = &v
}

func (o UpdateWorldRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateWorldRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetUrl) {
		toSerialize["assetUrl"] = o.AssetUrl
	}
	if !IsNil(o.AssetVersion) {
		toSerialize["assetVersion"] = o.AssetVersion
	}
	if !IsNil(o.AuthorId) {
		toSerialize["authorId"] = o.AuthorId
	}
	if !IsNil(o.AuthorName) {
		toSerialize["authorName"] = o.AuthorName
	}
	if !IsNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.ReleaseStatus) {
		toSerialize["releaseStatus"] = o.ReleaseStatus
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UnityPackageUrl) {
		toSerialize["unityPackageUrl"] = o.UnityPackageUrl
	}
	if !IsNil(o.UnityVersion) {
		toSerialize["unityVersion"] = o.UnityVersion
	}
	return toSerialize, nil
}

type NullableUpdateWorldRequest struct {
	value *UpdateWorldRequest
	isSet bool
}

func (v NullableUpdateWorldRequest) Get() *UpdateWorldRequest {
	return v.value
}

func (v *NullableUpdateWorldRequest) Set(val *UpdateWorldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWorldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWorldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWorldRequest(val *UpdateWorldRequest) *NullableUpdateWorldRequest {
	return &NullableUpdateWorldRequest{value: val, isSet: true}
}

func (v NullableUpdateWorldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWorldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


