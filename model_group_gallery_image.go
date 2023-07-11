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

// checks if the GroupGalleryImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupGalleryImage{}

// GroupGalleryImage struct for GroupGalleryImage
type GroupGalleryImage struct {
	Id *string `json:"id,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	GalleryId *string `json:"galleryId,omitempty"`
	FileId *string `json:"fileId,omitempty"`
	ImageUrl *string `json:"imageUrl,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	SubmittedByUserId *string `json:"submittedByUserId,omitempty"`
	Approved *bool `json:"approved,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	ApprovedByUserId *string `json:"approvedByUserId,omitempty"`
	ApprovedAt *time.Time `json:"approvedAt,omitempty"`
}

// NewGroupGalleryImage instantiates a new GroupGalleryImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupGalleryImage() *GroupGalleryImage {
	this := GroupGalleryImage{}
	var approved bool = false
	this.Approved = &approved
	return &this
}

// NewGroupGalleryImageWithDefaults instantiates a new GroupGalleryImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupGalleryImageWithDefaults() *GroupGalleryImage {
	this := GroupGalleryImage{}
	var approved bool = false
	this.Approved = &approved
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupGalleryImage) SetId(v string) {
	o.Id = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GroupGalleryImage) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGalleryId returns the GalleryId field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetGalleryId() string {
	if o == nil || IsNil(o.GalleryId) {
		var ret string
		return ret
	}
	return *o.GalleryId
}

// GetGalleryIdOk returns a tuple with the GalleryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetGalleryIdOk() (*string, bool) {
	if o == nil || IsNil(o.GalleryId) {
		return nil, false
	}
	return o.GalleryId, true
}

// HasGalleryId returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasGalleryId() bool {
	if o != nil && !IsNil(o.GalleryId) {
		return true
	}

	return false
}

// SetGalleryId gets a reference to the given string and assigns it to the GalleryId field.
func (o *GroupGalleryImage) SetGalleryId(v string) {
	o.GalleryId = &v
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetFileId() string {
	if o == nil || IsNil(o.FileId) {
		var ret string
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasFileId() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given string and assigns it to the FileId field.
func (o *GroupGalleryImage) SetFileId(v string) {
	o.FileId = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *GroupGalleryImage) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GroupGalleryImage) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetSubmittedByUserId returns the SubmittedByUserId field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetSubmittedByUserId() string {
	if o == nil || IsNil(o.SubmittedByUserId) {
		var ret string
		return ret
	}
	return *o.SubmittedByUserId
}

// GetSubmittedByUserIdOk returns a tuple with the SubmittedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetSubmittedByUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubmittedByUserId) {
		return nil, false
	}
	return o.SubmittedByUserId, true
}

// HasSubmittedByUserId returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasSubmittedByUserId() bool {
	if o != nil && !IsNil(o.SubmittedByUserId) {
		return true
	}

	return false
}

// SetSubmittedByUserId gets a reference to the given string and assigns it to the SubmittedByUserId field.
func (o *GroupGalleryImage) SetSubmittedByUserId(v string) {
	o.SubmittedByUserId = &v
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetApproved() bool {
	if o == nil || IsNil(o.Approved) {
		var ret bool
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetApprovedOk() (*bool, bool) {
	if o == nil || IsNil(o.Approved) {
		return nil, false
	}
	return o.Approved, true
}

// HasApproved returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasApproved() bool {
	if o != nil && !IsNil(o.Approved) {
		return true
	}

	return false
}

// SetApproved gets a reference to the given bool and assigns it to the Approved field.
func (o *GroupGalleryImage) SetApproved(v bool) {
	o.Approved = &v
}

// GetApprovedByUserId returns the ApprovedByUserId field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetApprovedByUserId() string {
	if o == nil || IsNil(o.ApprovedByUserId) {
		var ret string
		return ret
	}
	return *o.ApprovedByUserId
}

// GetApprovedByUserIdOk returns a tuple with the ApprovedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetApprovedByUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovedByUserId) {
		return nil, false
	}
	return o.ApprovedByUserId, true
}

// HasApprovedByUserId returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasApprovedByUserId() bool {
	if o != nil && !IsNil(o.ApprovedByUserId) {
		return true
	}

	return false
}

// SetApprovedByUserId gets a reference to the given string and assigns it to the ApprovedByUserId field.
func (o *GroupGalleryImage) SetApprovedByUserId(v string) {
	o.ApprovedByUserId = &v
}

// GetApprovedAt returns the ApprovedAt field value if set, zero value otherwise.
func (o *GroupGalleryImage) GetApprovedAt() time.Time {
	if o == nil || IsNil(o.ApprovedAt) {
		var ret time.Time
		return ret
	}
	return *o.ApprovedAt
}

// GetApprovedAtOk returns a tuple with the ApprovedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGalleryImage) GetApprovedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ApprovedAt) {
		return nil, false
	}
	return o.ApprovedAt, true
}

// HasApprovedAt returns a boolean if a field has been set.
func (o *GroupGalleryImage) HasApprovedAt() bool {
	if o != nil && !IsNil(o.ApprovedAt) {
		return true
	}

	return false
}

// SetApprovedAt gets a reference to the given time.Time and assigns it to the ApprovedAt field.
func (o *GroupGalleryImage) SetApprovedAt(v time.Time) {
	o.ApprovedAt = &v
}

func (o GroupGalleryImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupGalleryImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.GalleryId) {
		toSerialize["galleryId"] = o.GalleryId
	}
	if !IsNil(o.FileId) {
		toSerialize["fileId"] = o.FileId
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.SubmittedByUserId) {
		toSerialize["submittedByUserId"] = o.SubmittedByUserId
	}
	if !IsNil(o.Approved) {
		toSerialize["approved"] = o.Approved
	}
	if !IsNil(o.ApprovedByUserId) {
		toSerialize["approvedByUserId"] = o.ApprovedByUserId
	}
	if !IsNil(o.ApprovedAt) {
		toSerialize["approvedAt"] = o.ApprovedAt
	}
	return toSerialize, nil
}

type NullableGroupGalleryImage struct {
	value *GroupGalleryImage
	isSet bool
}

func (v NullableGroupGalleryImage) Get() *GroupGalleryImage {
	return v.value
}

func (v *NullableGroupGalleryImage) Set(val *GroupGalleryImage) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupGalleryImage) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupGalleryImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupGalleryImage(val *GroupGalleryImage) *NullableGroupGalleryImage {
	return &NullableGroupGalleryImage{value: val, isSet: true}
}

func (v NullableGroupGalleryImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupGalleryImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


