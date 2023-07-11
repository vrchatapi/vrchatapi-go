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

// checks if the Avatar type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Avatar{}

// Avatar 
type Avatar struct {
	// Not present from general serach `/avatars`, only on specific requests `/avatars/{avatarId}`.
	AssetUrl *string `json:"assetUrl,omitempty"`
	// Not present from general serach `/avatars`, only on specific requests `/avatars/{avatarId}`. **Deprecation:** `Object` has unknown usage/fields, and is always empty. Use normal `Url` field instead.
	AssetUrlObject map[string]interface{} `json:"assetUrlObject,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId string `json:"authorId"`
	AuthorName string `json:"authorName"`
	CreatedAt time.Time `json:"created_at"`
	Description string `json:"description"`
	Featured bool `json:"featured"`
	Id string `json:"id"`
	ImageUrl string `json:"imageUrl"`
	Name string `json:"name"`
	ReleaseStatus ReleaseStatus `json:"releaseStatus"`
	//  
	Tags []string `json:"tags"`
	ThumbnailImageUrl string `json:"thumbnailImageUrl"`
	UnityPackageUrl string `json:"unityPackageUrl"`
	// Deprecated
	UnityPackageUrlObject AvatarUnityPackageUrlObject `json:"unityPackageUrlObject"`
	UnityPackages []UnityPackage `json:"unityPackages"`
	UpdatedAt time.Time `json:"updated_at"`
	Version int32 `json:"version"`
}

// NewAvatar instantiates a new Avatar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvatar(authorId string, authorName string, createdAt time.Time, description string, featured bool, id string, imageUrl string, name string, releaseStatus ReleaseStatus, tags []string, thumbnailImageUrl string, unityPackageUrl string, unityPackageUrlObject AvatarUnityPackageUrlObject, unityPackages []UnityPackage, updatedAt time.Time, version int32) *Avatar {
	this := Avatar{}
	this.AuthorId = authorId
	this.AuthorName = authorName
	this.CreatedAt = createdAt
	this.Description = description
	this.Featured = featured
	this.Id = id
	this.ImageUrl = imageUrl
	this.Name = name
	this.ReleaseStatus = releaseStatus
	this.Tags = tags
	this.ThumbnailImageUrl = thumbnailImageUrl
	this.UnityPackageUrl = unityPackageUrl
	this.UnityPackageUrlObject = unityPackageUrlObject
	this.UnityPackages = unityPackages
	this.UpdatedAt = updatedAt
	this.Version = version
	return &this
}

// NewAvatarWithDefaults instantiates a new Avatar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvatarWithDefaults() *Avatar {
	this := Avatar{}
	var featured bool = false
	this.Featured = featured
	var releaseStatus ReleaseStatus = PUBLIC
	this.ReleaseStatus = releaseStatus
	var version int32 = 0
	this.Version = version
	return &this
}

// GetAssetUrl returns the AssetUrl field value if set, zero value otherwise.
func (o *Avatar) GetAssetUrl() string {
	if o == nil || IsNil(o.AssetUrl) {
		var ret string
		return ret
	}
	return *o.AssetUrl
}

// GetAssetUrlOk returns a tuple with the AssetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Avatar) GetAssetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AssetUrl) {
		return nil, false
	}
	return o.AssetUrl, true
}

// HasAssetUrl returns a boolean if a field has been set.
func (o *Avatar) HasAssetUrl() bool {
	if o != nil && !IsNil(o.AssetUrl) {
		return true
	}

	return false
}

// SetAssetUrl gets a reference to the given string and assigns it to the AssetUrl field.
func (o *Avatar) SetAssetUrl(v string) {
	o.AssetUrl = &v
}

// GetAssetUrlObject returns the AssetUrlObject field value if set, zero value otherwise.
func (o *Avatar) GetAssetUrlObject() map[string]interface{} {
	if o == nil || IsNil(o.AssetUrlObject) {
		var ret map[string]interface{}
		return ret
	}
	return o.AssetUrlObject
}

// GetAssetUrlObjectOk returns a tuple with the AssetUrlObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Avatar) GetAssetUrlObjectOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AssetUrlObject) {
		return map[string]interface{}{}, false
	}
	return o.AssetUrlObject, true
}

// HasAssetUrlObject returns a boolean if a field has been set.
func (o *Avatar) HasAssetUrlObject() bool {
	if o != nil && !IsNil(o.AssetUrlObject) {
		return true
	}

	return false
}

// SetAssetUrlObject gets a reference to the given map[string]interface{} and assigns it to the AssetUrlObject field.
func (o *Avatar) SetAssetUrlObject(v map[string]interface{}) {
	o.AssetUrlObject = v
}

// GetAuthorId returns the AuthorId field value
func (o *Avatar) GetAuthorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetAuthorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorId, true
}

// SetAuthorId sets field value
func (o *Avatar) SetAuthorId(v string) {
	o.AuthorId = v
}

// GetAuthorName returns the AuthorName field value
func (o *Avatar) GetAuthorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetAuthorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorName, true
}

// SetAuthorName sets field value
func (o *Avatar) SetAuthorName(v string) {
	o.AuthorName = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Avatar) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Avatar) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value
func (o *Avatar) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Avatar) SetDescription(v string) {
	o.Description = v
}

// GetFeatured returns the Featured field value
func (o *Avatar) GetFeatured() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetFeaturedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Featured, true
}

// SetFeatured sets field value
func (o *Avatar) SetFeatured(v bool) {
	o.Featured = v
}

// GetId returns the Id field value
func (o *Avatar) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Avatar) SetId(v string) {
	o.Id = v
}

// GetImageUrl returns the ImageUrl field value
func (o *Avatar) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *Avatar) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetName returns the Name field value
func (o *Avatar) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Avatar) SetName(v string) {
	o.Name = v
}

// GetReleaseStatus returns the ReleaseStatus field value
func (o *Avatar) GetReleaseStatus() ReleaseStatus {
	if o == nil {
		var ret ReleaseStatus
		return ret
	}

	return o.ReleaseStatus
}

// GetReleaseStatusOk returns a tuple with the ReleaseStatus field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetReleaseStatusOk() (*ReleaseStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseStatus, true
}

// SetReleaseStatus sets field value
func (o *Avatar) SetReleaseStatus(v ReleaseStatus) {
	o.ReleaseStatus = v
}

// GetTags returns the Tags field value
func (o *Avatar) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Avatar) SetTags(v []string) {
	o.Tags = v
}

// GetThumbnailImageUrl returns the ThumbnailImageUrl field value
func (o *Avatar) GetThumbnailImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbnailImageUrl
}

// GetThumbnailImageUrlOk returns a tuple with the ThumbnailImageUrl field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetThumbnailImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThumbnailImageUrl, true
}

// SetThumbnailImageUrl sets field value
func (o *Avatar) SetThumbnailImageUrl(v string) {
	o.ThumbnailImageUrl = v
}

// GetUnityPackageUrl returns the UnityPackageUrl field value
func (o *Avatar) GetUnityPackageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnityPackageUrl
}

// GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetUnityPackageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnityPackageUrl, true
}

// SetUnityPackageUrl sets field value
func (o *Avatar) SetUnityPackageUrl(v string) {
	o.UnityPackageUrl = v
}

// GetUnityPackageUrlObject returns the UnityPackageUrlObject field value
// Deprecated
func (o *Avatar) GetUnityPackageUrlObject() AvatarUnityPackageUrlObject {
	if o == nil {
		var ret AvatarUnityPackageUrlObject
		return ret
	}

	return o.UnityPackageUrlObject
}

// GetUnityPackageUrlObjectOk returns a tuple with the UnityPackageUrlObject field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *Avatar) GetUnityPackageUrlObjectOk() (*AvatarUnityPackageUrlObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnityPackageUrlObject, true
}

// SetUnityPackageUrlObject sets field value
// Deprecated
func (o *Avatar) SetUnityPackageUrlObject(v AvatarUnityPackageUrlObject) {
	o.UnityPackageUrlObject = v
}

// GetUnityPackages returns the UnityPackages field value
func (o *Avatar) GetUnityPackages() []UnityPackage {
	if o == nil {
		var ret []UnityPackage
		return ret
	}

	return o.UnityPackages
}

// GetUnityPackagesOk returns a tuple with the UnityPackages field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetUnityPackagesOk() ([]UnityPackage, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnityPackages, true
}

// SetUnityPackages sets field value
func (o *Avatar) SetUnityPackages(v []UnityPackage) {
	o.UnityPackages = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Avatar) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Avatar) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetVersion returns the Version field value
func (o *Avatar) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Avatar) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Avatar) SetVersion(v int32) {
	o.Version = v
}

func (o Avatar) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Avatar) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetUrl) {
		toSerialize["assetUrl"] = o.AssetUrl
	}
	if !IsNil(o.AssetUrlObject) {
		toSerialize["assetUrlObject"] = o.AssetUrlObject
	}
	toSerialize["authorId"] = o.AuthorId
	toSerialize["authorName"] = o.AuthorName
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["description"] = o.Description
	toSerialize["featured"] = o.Featured
	toSerialize["id"] = o.Id
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["name"] = o.Name
	toSerialize["releaseStatus"] = o.ReleaseStatus
	toSerialize["tags"] = o.Tags
	toSerialize["thumbnailImageUrl"] = o.ThumbnailImageUrl
	toSerialize["unityPackageUrl"] = o.UnityPackageUrl
	toSerialize["unityPackageUrlObject"] = o.UnityPackageUrlObject
	toSerialize["unityPackages"] = o.UnityPackages
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableAvatar struct {
	value *Avatar
	isSet bool
}

func (v NullableAvatar) Get() *Avatar {
	return v.value
}

func (v *NullableAvatar) Set(val *Avatar) {
	v.value = val
	v.isSet = true
}

func (v NullableAvatar) IsSet() bool {
	return v.isSet
}

func (v *NullableAvatar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvatar(val *Avatar) *NullableAvatar {
	return &NullableAvatar{value: val, isSet: true}
}

func (v NullableAvatar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvatar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


