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

// checks if the GroupGallery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupGallery{}

// GroupGallery struct for GroupGallery
type GroupGallery struct {
	Id *string `json:"id,omitempty"`
	// Name of the gallery.
	Name *string `json:"name,omitempty"`
	// Description of the gallery.
	Description *string `json:"description,omitempty"`
	// Whether the gallery is members only.
	MembersOnly *bool `json:"membersOnly,omitempty"`
	//  
	RoleIdsToView []string `json:"roleIdsToView,omitempty"`
	//  
	RoleIdsToSubmit []string `json:"roleIdsToSubmit,omitempty"`
	//  
	RoleIdsToAutoApprove []string `json:"roleIdsToAutoApprove,omitempty"`
	//  
	RoleIdsToManage []string `json:"roleIdsToManage,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewGroupGallery instantiates a new GroupGallery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupGallery() *GroupGallery {
	this := GroupGallery{}
	var membersOnly bool = false
	this.MembersOnly = &membersOnly
	return &this
}

// NewGroupGalleryWithDefaults instantiates a new GroupGallery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupGalleryWithDefaults() *GroupGallery {
	this := GroupGallery{}
	var membersOnly bool = false
	this.MembersOnly = &membersOnly
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupGallery) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGallery) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupGallery) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupGallery) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupGallery) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGallery) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupGallery) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupGallery) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupGallery) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGallery) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupGallery) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupGallery) SetDescription(v string) {
	o.Description = &v
}

// GetMembersOnly returns the MembersOnly field value if set, zero value otherwise.
func (o *GroupGallery) GetMembersOnly() bool {
	if o == nil || IsNil(o.MembersOnly) {
		var ret bool
		return ret
	}
	return *o.MembersOnly
}

// GetMembersOnlyOk returns a tuple with the MembersOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGallery) GetMembersOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.MembersOnly) {
		return nil, false
	}
	return o.MembersOnly, true
}

// HasMembersOnly returns a boolean if a field has been set.
func (o *GroupGallery) HasMembersOnly() bool {
	if o != nil && !IsNil(o.MembersOnly) {
		return true
	}

	return false
}

// SetMembersOnly gets a reference to the given bool and assigns it to the MembersOnly field.
func (o *GroupGallery) SetMembersOnly(v bool) {
	o.MembersOnly = &v
}

// GetRoleIdsToView returns the RoleIdsToView field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupGallery) GetRoleIdsToView() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RoleIdsToView
}

// GetRoleIdsToViewOk returns a tuple with the RoleIdsToView field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupGallery) GetRoleIdsToViewOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleIdsToView) {
		return nil, false
	}
	return o.RoleIdsToView, true
}

// HasRoleIdsToView returns a boolean if a field has been set.
func (o *GroupGallery) HasRoleIdsToView() bool {
	if o != nil && IsNil(o.RoleIdsToView) {
		return true
	}

	return false
}

// SetRoleIdsToView gets a reference to the given []string and assigns it to the RoleIdsToView field.
func (o *GroupGallery) SetRoleIdsToView(v []string) {
	o.RoleIdsToView = v
}

// GetRoleIdsToSubmit returns the RoleIdsToSubmit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupGallery) GetRoleIdsToSubmit() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RoleIdsToSubmit
}

// GetRoleIdsToSubmitOk returns a tuple with the RoleIdsToSubmit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupGallery) GetRoleIdsToSubmitOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleIdsToSubmit) {
		return nil, false
	}
	return o.RoleIdsToSubmit, true
}

// HasRoleIdsToSubmit returns a boolean if a field has been set.
func (o *GroupGallery) HasRoleIdsToSubmit() bool {
	if o != nil && IsNil(o.RoleIdsToSubmit) {
		return true
	}

	return false
}

// SetRoleIdsToSubmit gets a reference to the given []string and assigns it to the RoleIdsToSubmit field.
func (o *GroupGallery) SetRoleIdsToSubmit(v []string) {
	o.RoleIdsToSubmit = v
}

// GetRoleIdsToAutoApprove returns the RoleIdsToAutoApprove field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupGallery) GetRoleIdsToAutoApprove() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RoleIdsToAutoApprove
}

// GetRoleIdsToAutoApproveOk returns a tuple with the RoleIdsToAutoApprove field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupGallery) GetRoleIdsToAutoApproveOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleIdsToAutoApprove) {
		return nil, false
	}
	return o.RoleIdsToAutoApprove, true
}

// HasRoleIdsToAutoApprove returns a boolean if a field has been set.
func (o *GroupGallery) HasRoleIdsToAutoApprove() bool {
	if o != nil && IsNil(o.RoleIdsToAutoApprove) {
		return true
	}

	return false
}

// SetRoleIdsToAutoApprove gets a reference to the given []string and assigns it to the RoleIdsToAutoApprove field.
func (o *GroupGallery) SetRoleIdsToAutoApprove(v []string) {
	o.RoleIdsToAutoApprove = v
}

// GetRoleIdsToManage returns the RoleIdsToManage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupGallery) GetRoleIdsToManage() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RoleIdsToManage
}

// GetRoleIdsToManageOk returns a tuple with the RoleIdsToManage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupGallery) GetRoleIdsToManageOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleIdsToManage) {
		return nil, false
	}
	return o.RoleIdsToManage, true
}

// HasRoleIdsToManage returns a boolean if a field has been set.
func (o *GroupGallery) HasRoleIdsToManage() bool {
	if o != nil && IsNil(o.RoleIdsToManage) {
		return true
	}

	return false
}

// SetRoleIdsToManage gets a reference to the given []string and assigns it to the RoleIdsToManage field.
func (o *GroupGallery) SetRoleIdsToManage(v []string) {
	o.RoleIdsToManage = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GroupGallery) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGallery) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GroupGallery) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GroupGallery) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GroupGallery) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupGallery) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GroupGallery) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *GroupGallery) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o GroupGallery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupGallery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MembersOnly) {
		toSerialize["membersOnly"] = o.MembersOnly
	}
	if o.RoleIdsToView != nil {
		toSerialize["roleIdsToView"] = o.RoleIdsToView
	}
	if o.RoleIdsToSubmit != nil {
		toSerialize["roleIdsToSubmit"] = o.RoleIdsToSubmit
	}
	if o.RoleIdsToAutoApprove != nil {
		toSerialize["roleIdsToAutoApprove"] = o.RoleIdsToAutoApprove
	}
	if o.RoleIdsToManage != nil {
		toSerialize["roleIdsToManage"] = o.RoleIdsToManage
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableGroupGallery struct {
	value *GroupGallery
	isSet bool
}

func (v NullableGroupGallery) Get() *GroupGallery {
	return v.value
}

func (v *NullableGroupGallery) Set(val *GroupGallery) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupGallery) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupGallery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupGallery(val *GroupGallery) *NullableGroupGallery {
	return &NullableGroupGallery{value: val, isSet: true}
}

func (v NullableGroupGallery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupGallery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


