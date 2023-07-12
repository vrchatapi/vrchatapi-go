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

// checks if the CurrentUserPresence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrentUserPresence{}

// CurrentUserPresence struct for CurrentUserPresence
type CurrentUserPresence struct {
	AvatarThumbnail NullableString `json:"avatarThumbnail,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Groups []string `json:"groups,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id *string `json:"id,omitempty"`
	Instance NullableString `json:"instance,omitempty"`
	// either an InstanceType or an empty string
	InstanceType NullableString `json:"instanceType,omitempty"`
	IsRejoining NullableString `json:"isRejoining,omitempty"`
	// either a Platform or an empty string
	Platform NullableString `json:"platform,omitempty"`
	ProfilePicOverride NullableString `json:"profilePicOverride,omitempty"`
	// either a UserStatus or empty string
	Status NullableString `json:"status,omitempty"`
	TravelingToInstance NullableString `json:"travelingToInstance,omitempty"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	TravelingToWorld *string `json:"travelingToWorld,omitempty"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	World *string `json:"world,omitempty"`
}

// NewCurrentUserPresence instantiates a new CurrentUserPresence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentUserPresence() *CurrentUserPresence {
	this := CurrentUserPresence{}
	return &this
}

// NewCurrentUserPresenceWithDefaults instantiates a new CurrentUserPresence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentUserPresenceWithDefaults() *CurrentUserPresence {
	this := CurrentUserPresence{}
	return &this
}

// GetAvatarThumbnail returns the AvatarThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUserPresence) GetAvatarThumbnail() string {
	if o == nil || IsNil(o.AvatarThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.AvatarThumbnail.Get()
}

// GetAvatarThumbnailOk returns a tuple with the AvatarThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUserPresence) GetAvatarThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarThumbnail.Get(), o.AvatarThumbnail.IsSet()
}

// HasAvatarThumbnail returns a boolean if a field has been set.
func (o *CurrentUserPresence) HasAvatarThumbnail() bool {
	if o != nil && o.AvatarThumbnail.IsSet() {
		return true
	}

	return false
}

// SetAvatarThumbnail gets a reference to the given NullableString and assigns it to the AvatarThumbnail field.
func (o *CurrentUserPresence) SetAvatarThumbnail(v string) {
	o.AvatarThumbnail.Set(&v)
}
// SetAvatarThumbnailNil sets the value for AvatarThumbnail to be an explicit nil
func (o *CurrentUserPresence) SetAvatarThumbnailNil() {
	o.AvatarThumbnail.Set(nil)
}

// UnsetAvatarThumbnail ensures that no value is present for AvatarThumbnail, not even an explicit nil
func (o *CurrentUserPresence) UnsetAvatarThumbnail() {
	o.AvatarThumbnail.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CurrentUserPresence) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserPresence) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CurrentUserPresence) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CurrentUserPresence) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUserPresence) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUserPresence) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *CurrentUserPresence) HasGroups() bool {
	if o != nil && IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *CurrentUserPresence) SetGroups(v []string) {
	o.Groups = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CurrentUserPresence) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserPresence) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CurrentUserPresence) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CurrentUserPresence) SetId(v string) {
	o.Id = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUserPresence) GetInstance() string {
	if o == nil || IsNil(o.Instance.Get()) {
		var ret string
		return ret
	}
	return *o.Instance.Get()
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUserPresence) GetInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instance.Get(), o.Instance.IsSet()
}

// HasInstance returns a boolean if a field has been set.
func (o *CurrentUserPresence) HasInstance() bool {
	if o != nil && o.Instance.IsSet() {
		return true
	}

	return false
}

// SetInstance gets a reference to the given NullableString and assigns it to the Instance field.
func (o *CurrentUserPresence) SetInstance(v string) {
	o.Instance.Set(&v)
}
// SetInstanceNil sets the value for Instance to be an explicit nil
func (o *CurrentUserPresence) SetInstanceNil() {
	o.Instance.Set(nil)
}

// UnsetInstance ensures that no value is present for Instance, not even an explicit nil
func (o *CurrentUserPresence) UnsetInstance() {
	o.Instance.Unset()
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUserPresence) GetInstanceType() string {
	if o == nil || IsNil(o.InstanceType.Get()) {
		var ret string
		return ret
	}
	return *o.InstanceType.Get()
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUserPresence) GetInstanceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceType.Get(), o.InstanceType.IsSet()
}

// HasInstanceType returns a boolean if a field has been set.
func (o *CurrentUserPresence) HasInstanceType() bool {
	if o != nil && o.InstanceType.IsSet() {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given NullableString and assigns it to the InstanceType field.
func (o *CurrentUserPresence) SetInstanceType(v string) {
	o.InstanceType.Set(&v)
}
// SetInstanceTypeNil sets the value for InstanceType to be an explicit nil
func (o *CurrentUserPresence) SetInstanceTypeNil() {
	o.InstanceType.Set(nil)
}

// UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
func (o *CurrentUserPresence) UnsetInstanceType() {
	o.InstanceType.Unset()
}

// GetIsRejoining returns the IsRejoining field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUserPresence) GetIsRejoining() string {
	if o == nil || IsNil(o.IsRejoining.Get()) {
		var ret string
		return ret
	}
	return *o.IsRejoining.Get()
}

// GetIsRejoiningOk returns a tuple with the IsRejoining field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUserPresence) GetIsRejoiningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRejoining.Get(), o.IsRejoining.IsSet()
}

// HasIsRejoining returns a boolean if a field has been set.
func (o *CurrentUserPresence) HasIsRejoining() bool {
	if o != nil && o.IsRejoining.IsSet() {
		return true
	}

	return false
}

// SetIsRejoining gets a reference to the given NullableString and assigns it to the IsRejoining field.
func (o *CurrentUserPresence) SetIsRejoining(v string) {
	o.IsRejoining.Set(&v)
}
// SetIsRejoiningNil sets the value for IsRejoining to be an explicit nil
func (o *CurrentUserPresence) SetIsRejoiningNil() {
	o.IsRejoining.Set(nil)
}

// UnsetIsRejoining ensures that no value is present for IsRejoining, not even an explicit nil
func (o *CurrentUserPresence) UnsetIsRejoining() {
	o.IsRejoining.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUserPresence) GetPlatform() string {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret string
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUserPresence) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *CurrentUserPresence) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableString and assigns it to the Platform field.
func (o *CurrentUserPresence) SetPlatform(v string) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *CurrentUserPresence) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *CurrentUserPresence) UnsetPlatform() {
	o.Platform.Unset()
}

// GetProfilePicOverride returns the ProfilePicOverride field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUserPresence) GetProfilePicOverride() string {
	if o == nil || IsNil(o.ProfilePicOverride.Get()) {
		var ret string
		return ret
	}
	return *o.ProfilePicOverride.Get()
}

// GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUserPresence) GetProfilePicOverrideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfilePicOverride.Get(), o.ProfilePicOverride.IsSet()
}

// HasProfilePicOverride returns a boolean if a field has been set.
func (o *CurrentUserPresence) HasProfilePicOverride() bool {
	if o != nil && o.ProfilePicOverride.IsSet() {
		return true
	}

	return false
}

// SetProfilePicOverride gets a reference to the given NullableString and assigns it to the ProfilePicOverride field.
func (o *CurrentUserPresence) SetProfilePicOverride(v string) {
	o.ProfilePicOverride.Set(&v)
}
// SetProfilePicOverrideNil sets the value for ProfilePicOverride to be an explicit nil
func (o *CurrentUserPresence) SetProfilePicOverrideNil() {
	o.ProfilePicOverride.Set(nil)
}

// UnsetProfilePicOverride ensures that no value is present for ProfilePicOverride, not even an explicit nil
func (o *CurrentUserPresence) UnsetProfilePicOverride() {
	o.ProfilePicOverride.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUserPresence) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUserPresence) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *CurrentUserPresence) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *CurrentUserPresence) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *CurrentUserPresence) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *CurrentUserPresence) UnsetStatus() {
	o.Status.Unset()
}

// GetTravelingToInstance returns the TravelingToInstance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUserPresence) GetTravelingToInstance() string {
	if o == nil || IsNil(o.TravelingToInstance.Get()) {
		var ret string
		return ret
	}
	return *o.TravelingToInstance.Get()
}

// GetTravelingToInstanceOk returns a tuple with the TravelingToInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUserPresence) GetTravelingToInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TravelingToInstance.Get(), o.TravelingToInstance.IsSet()
}

// HasTravelingToInstance returns a boolean if a field has been set.
func (o *CurrentUserPresence) HasTravelingToInstance() bool {
	if o != nil && o.TravelingToInstance.IsSet() {
		return true
	}

	return false
}

// SetTravelingToInstance gets a reference to the given NullableString and assigns it to the TravelingToInstance field.
func (o *CurrentUserPresence) SetTravelingToInstance(v string) {
	o.TravelingToInstance.Set(&v)
}
// SetTravelingToInstanceNil sets the value for TravelingToInstance to be an explicit nil
func (o *CurrentUserPresence) SetTravelingToInstanceNil() {
	o.TravelingToInstance.Set(nil)
}

// UnsetTravelingToInstance ensures that no value is present for TravelingToInstance, not even an explicit nil
func (o *CurrentUserPresence) UnsetTravelingToInstance() {
	o.TravelingToInstance.Unset()
}

// GetTravelingToWorld returns the TravelingToWorld field value if set, zero value otherwise.
func (o *CurrentUserPresence) GetTravelingToWorld() string {
	if o == nil || IsNil(o.TravelingToWorld) {
		var ret string
		return ret
	}
	return *o.TravelingToWorld
}

// GetTravelingToWorldOk returns a tuple with the TravelingToWorld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserPresence) GetTravelingToWorldOk() (*string, bool) {
	if o == nil || IsNil(o.TravelingToWorld) {
		return nil, false
	}
	return o.TravelingToWorld, true
}

// HasTravelingToWorld returns a boolean if a field has been set.
func (o *CurrentUserPresence) HasTravelingToWorld() bool {
	if o != nil && !IsNil(o.TravelingToWorld) {
		return true
	}

	return false
}

// SetTravelingToWorld gets a reference to the given string and assigns it to the TravelingToWorld field.
func (o *CurrentUserPresence) SetTravelingToWorld(v string) {
	o.TravelingToWorld = &v
}

// GetWorld returns the World field value if set, zero value otherwise.
func (o *CurrentUserPresence) GetWorld() string {
	if o == nil || IsNil(o.World) {
		var ret string
		return ret
	}
	return *o.World
}

// GetWorldOk returns a tuple with the World field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserPresence) GetWorldOk() (*string, bool) {
	if o == nil || IsNil(o.World) {
		return nil, false
	}
	return o.World, true
}

// HasWorld returns a boolean if a field has been set.
func (o *CurrentUserPresence) HasWorld() bool {
	if o != nil && !IsNil(o.World) {
		return true
	}

	return false
}

// SetWorld gets a reference to the given string and assigns it to the World field.
func (o *CurrentUserPresence) SetWorld(v string) {
	o.World = &v
}

func (o CurrentUserPresence) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentUserPresence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AvatarThumbnail.IsSet() {
		toSerialize["avatarThumbnail"] = o.AvatarThumbnail.Get()
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Instance.IsSet() {
		toSerialize["instance"] = o.Instance.Get()
	}
	if o.InstanceType.IsSet() {
		toSerialize["instanceType"] = o.InstanceType.Get()
	}
	if o.IsRejoining.IsSet() {
		toSerialize["isRejoining"] = o.IsRejoining.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if o.ProfilePicOverride.IsSet() {
		toSerialize["profilePicOverride"] = o.ProfilePicOverride.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.TravelingToInstance.IsSet() {
		toSerialize["travelingToInstance"] = o.TravelingToInstance.Get()
	}
	if !IsNil(o.TravelingToWorld) {
		toSerialize["travelingToWorld"] = o.TravelingToWorld
	}
	if !IsNil(o.World) {
		toSerialize["world"] = o.World
	}
	return toSerialize, nil
}

type NullableCurrentUserPresence struct {
	value *CurrentUserPresence
	isSet bool
}

func (v NullableCurrentUserPresence) Get() *CurrentUserPresence {
	return v.value
}

func (v *NullableCurrentUserPresence) Set(val *CurrentUserPresence) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentUserPresence) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentUserPresence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentUserPresence(val *CurrentUserPresence) *NullableCurrentUserPresence {
	return &NullableCurrentUserPresence{value: val, isSet: true}
}

func (v NullableCurrentUserPresence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentUserPresence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


