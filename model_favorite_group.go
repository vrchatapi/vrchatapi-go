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

// checks if the FavoriteGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FavoriteGroup{}

// FavoriteGroup 
type FavoriteGroup struct {
	DisplayName string `json:"displayName"`
	Id string `json:"id"`
	Name string `json:"name"`
	OwnerDisplayName string `json:"ownerDisplayName"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	OwnerId string `json:"ownerId"`
	//  
	Tags []string `json:"tags"`
	Type FavoriteType `json:"type"`
	Visibility FavoriteGroupVisibility `json:"visibility"`
}

// NewFavoriteGroup instantiates a new FavoriteGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFavoriteGroup(displayName string, id string, name string, ownerDisplayName string, ownerId string, tags []string, type_ FavoriteType, visibility FavoriteGroupVisibility) *FavoriteGroup {
	this := FavoriteGroup{}
	this.DisplayName = displayName
	this.Id = id
	this.Name = name
	this.OwnerDisplayName = ownerDisplayName
	this.OwnerId = ownerId
	this.Tags = tags
	this.Type = type_
	this.Visibility = visibility
	return &this
}

// NewFavoriteGroupWithDefaults instantiates a new FavoriteGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFavoriteGroupWithDefaults() *FavoriteGroup {
	this := FavoriteGroup{}
	var type_ FavoriteType = FRIEND
	this.Type = type_
	var visibility FavoriteGroupVisibility = PRIVATE
	this.Visibility = visibility
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *FavoriteGroup) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *FavoriteGroup) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *FavoriteGroup) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetId returns the Id field value
func (o *FavoriteGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FavoriteGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FavoriteGroup) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FavoriteGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FavoriteGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FavoriteGroup) SetName(v string) {
	o.Name = v
}

// GetOwnerDisplayName returns the OwnerDisplayName field value
func (o *FavoriteGroup) GetOwnerDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerDisplayName
}

// GetOwnerDisplayNameOk returns a tuple with the OwnerDisplayName field value
// and a boolean to check if the value has been set.
func (o *FavoriteGroup) GetOwnerDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerDisplayName, true
}

// SetOwnerDisplayName sets field value
func (o *FavoriteGroup) SetOwnerDisplayName(v string) {
	o.OwnerDisplayName = v
}

// GetOwnerId returns the OwnerId field value
func (o *FavoriteGroup) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *FavoriteGroup) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *FavoriteGroup) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetTags returns the Tags field value
func (o *FavoriteGroup) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *FavoriteGroup) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *FavoriteGroup) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *FavoriteGroup) GetType() FavoriteType {
	if o == nil {
		var ret FavoriteType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FavoriteGroup) GetTypeOk() (*FavoriteType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FavoriteGroup) SetType(v FavoriteType) {
	o.Type = v
}

// GetVisibility returns the Visibility field value
func (o *FavoriteGroup) GetVisibility() FavoriteGroupVisibility {
	if o == nil {
		var ret FavoriteGroupVisibility
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *FavoriteGroup) GetVisibilityOk() (*FavoriteGroupVisibility, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *FavoriteGroup) SetVisibility(v FavoriteGroupVisibility) {
	o.Visibility = v
}

func (o FavoriteGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FavoriteGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["ownerDisplayName"] = o.OwnerDisplayName
	toSerialize["ownerId"] = o.OwnerId
	toSerialize["tags"] = o.Tags
	toSerialize["type"] = o.Type
	toSerialize["visibility"] = o.Visibility
	return toSerialize, nil
}

type NullableFavoriteGroup struct {
	value *FavoriteGroup
	isSet bool
}

func (v NullableFavoriteGroup) Get() *FavoriteGroup {
	return v.value
}

func (v *NullableFavoriteGroup) Set(val *FavoriteGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableFavoriteGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableFavoriteGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFavoriteGroup(val *FavoriteGroup) *NullableFavoriteGroup {
	return &NullableFavoriteGroup{value: val, isSet: true}
}

func (v NullableFavoriteGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFavoriteGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


