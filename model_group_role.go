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

// checks if the GroupRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupRole{}

// GroupRole struct for GroupRole
type GroupRole struct {
	Id *string `json:"id,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsSelfAssignable *bool `json:"isSelfAssignable,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	IsManagementRole *bool `json:"isManagementRole,omitempty"`
	RequiresTwoFactor *bool `json:"requiresTwoFactor,omitempty"`
	RequiresPurchase *bool `json:"requiresPurchase,omitempty"`
	Order *int32 `json:"order,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewGroupRole instantiates a new GroupRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRole() *GroupRole {
	this := GroupRole{}
	var isSelfAssignable bool = false
	this.IsSelfAssignable = &isSelfAssignable
	var isManagementRole bool = false
	this.IsManagementRole = &isManagementRole
	var requiresTwoFactor bool = false
	this.RequiresTwoFactor = &requiresTwoFactor
	var requiresPurchase bool = false
	this.RequiresPurchase = &requiresPurchase
	return &this
}

// NewGroupRoleWithDefaults instantiates a new GroupRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRoleWithDefaults() *GroupRole {
	this := GroupRole{}
	var isSelfAssignable bool = false
	this.IsSelfAssignable = &isSelfAssignable
	var isManagementRole bool = false
	this.IsManagementRole = &isManagementRole
	var requiresTwoFactor bool = false
	this.RequiresTwoFactor = &requiresTwoFactor
	var requiresPurchase bool = false
	this.RequiresPurchase = &requiresPurchase
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupRole) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupRole) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupRole) SetId(v string) {
	o.Id = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GroupRole) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupRole) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GroupRole) SetGroupId(v string) {
	o.GroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupRole) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupRole) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupRole) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupRole) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupRole) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupRole) SetDescription(v string) {
	o.Description = &v
}

// GetIsSelfAssignable returns the IsSelfAssignable field value if set, zero value otherwise.
func (o *GroupRole) GetIsSelfAssignable() bool {
	if o == nil || IsNil(o.IsSelfAssignable) {
		var ret bool
		return ret
	}
	return *o.IsSelfAssignable
}

// GetIsSelfAssignableOk returns a tuple with the IsSelfAssignable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetIsSelfAssignableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSelfAssignable) {
		return nil, false
	}
	return o.IsSelfAssignable, true
}

// HasIsSelfAssignable returns a boolean if a field has been set.
func (o *GroupRole) HasIsSelfAssignable() bool {
	if o != nil && !IsNil(o.IsSelfAssignable) {
		return true
	}

	return false
}

// SetIsSelfAssignable gets a reference to the given bool and assigns it to the IsSelfAssignable field.
func (o *GroupRole) SetIsSelfAssignable(v bool) {
	o.IsSelfAssignable = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *GroupRole) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GroupRole) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *GroupRole) SetPermissions(v []string) {
	o.Permissions = v
}

// GetIsManagementRole returns the IsManagementRole field value if set, zero value otherwise.
func (o *GroupRole) GetIsManagementRole() bool {
	if o == nil || IsNil(o.IsManagementRole) {
		var ret bool
		return ret
	}
	return *o.IsManagementRole
}

// GetIsManagementRoleOk returns a tuple with the IsManagementRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetIsManagementRoleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsManagementRole) {
		return nil, false
	}
	return o.IsManagementRole, true
}

// HasIsManagementRole returns a boolean if a field has been set.
func (o *GroupRole) HasIsManagementRole() bool {
	if o != nil && !IsNil(o.IsManagementRole) {
		return true
	}

	return false
}

// SetIsManagementRole gets a reference to the given bool and assigns it to the IsManagementRole field.
func (o *GroupRole) SetIsManagementRole(v bool) {
	o.IsManagementRole = &v
}

// GetRequiresTwoFactor returns the RequiresTwoFactor field value if set, zero value otherwise.
func (o *GroupRole) GetRequiresTwoFactor() bool {
	if o == nil || IsNil(o.RequiresTwoFactor) {
		var ret bool
		return ret
	}
	return *o.RequiresTwoFactor
}

// GetRequiresTwoFactorOk returns a tuple with the RequiresTwoFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetRequiresTwoFactorOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiresTwoFactor) {
		return nil, false
	}
	return o.RequiresTwoFactor, true
}

// HasRequiresTwoFactor returns a boolean if a field has been set.
func (o *GroupRole) HasRequiresTwoFactor() bool {
	if o != nil && !IsNil(o.RequiresTwoFactor) {
		return true
	}

	return false
}

// SetRequiresTwoFactor gets a reference to the given bool and assigns it to the RequiresTwoFactor field.
func (o *GroupRole) SetRequiresTwoFactor(v bool) {
	o.RequiresTwoFactor = &v
}

// GetRequiresPurchase returns the RequiresPurchase field value if set, zero value otherwise.
func (o *GroupRole) GetRequiresPurchase() bool {
	if o == nil || IsNil(o.RequiresPurchase) {
		var ret bool
		return ret
	}
	return *o.RequiresPurchase
}

// GetRequiresPurchaseOk returns a tuple with the RequiresPurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetRequiresPurchaseOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiresPurchase) {
		return nil, false
	}
	return o.RequiresPurchase, true
}

// HasRequiresPurchase returns a boolean if a field has been set.
func (o *GroupRole) HasRequiresPurchase() bool {
	if o != nil && !IsNil(o.RequiresPurchase) {
		return true
	}

	return false
}

// SetRequiresPurchase gets a reference to the given bool and assigns it to the RequiresPurchase field.
func (o *GroupRole) SetRequiresPurchase(v bool) {
	o.RequiresPurchase = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GroupRole) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GroupRole) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *GroupRole) SetOrder(v int32) {
	o.Order = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GroupRole) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GroupRole) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GroupRole) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GroupRole) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRole) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GroupRole) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *GroupRole) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o GroupRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsSelfAssignable) {
		toSerialize["isSelfAssignable"] = o.IsSelfAssignable
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.IsManagementRole) {
		toSerialize["isManagementRole"] = o.IsManagementRole
	}
	if !IsNil(o.RequiresTwoFactor) {
		toSerialize["requiresTwoFactor"] = o.RequiresTwoFactor
	}
	if !IsNil(o.RequiresPurchase) {
		toSerialize["requiresPurchase"] = o.RequiresPurchase
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableGroupRole struct {
	value *GroupRole
	isSet bool
}

func (v NullableGroupRole) Get() *GroupRole {
	return v.value
}

func (v *NullableGroupRole) Set(val *GroupRole) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRole) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRole(val *GroupRole) *NullableGroupRole {
	return &NullableGroupRole{value: val, isSet: true}
}

func (v NullableGroupRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


