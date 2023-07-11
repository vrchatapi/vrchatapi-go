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

// checks if the Instance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Instance{}

// Instance * `hidden` field is only present if InstanceType is `hidden` aka \"Friends+\", and is instance creator. * `friends` field is only present if InstanceType is `friends` aka \"Friends\", and is instance creator. * `private` field is only present if InstanceType is `private` aka \"Invite\" or \"Invite+\", and is instance creator.
type Instance struct {
	Active bool `json:"active"`
	CanRequestInvite bool `json:"canRequestInvite"`
	Capacity int32 `json:"capacity"`
	// Always returns \"unknown\".
	// Deprecated
	ClientNumber string `json:"clientNumber"`
	Full bool `json:"full"`
	// InstanceID can be \"offline\" on User profiles if you are not friends with that user and \"private\" if you are friends and user is in private instance.
	Id string `json:"id"`
	InstanceId string `json:"instanceId"`
	// InstanceID can be \"offline\" on User profiles if you are not friends with that user and \"private\" if you are friends and user is in private instance.
	Location string `json:"location"`
	NUsers int32 `json:"n_users"`
	Name string `json:"name"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	OwnerId *string `json:"ownerId,omitempty"`
	Permanent bool `json:"permanent"`
	PhotonRegion Region `json:"photonRegion"`
	Platforms InstancePlatforms `json:"platforms"`
	Region Region `json:"region"`
	SecureName string `json:"secureName"`
	ShortName *string `json:"shortName,omitempty"`
	// The tags array on Instances usually contain the language tags of the people in the instance. 
	Tags []string `json:"tags"`
	Type InstanceType `json:"type"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	WorldId string `json:"worldId"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Hidden *string `json:"hidden,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Friends *string `json:"friends,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Private *string `json:"private,omitempty"`
}

// NewInstance instantiates a new Instance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstance(active bool, canRequestInvite bool, capacity int32, clientNumber string, full bool, id string, instanceId string, location string, nUsers int32, name string, permanent bool, photonRegion Region, platforms InstancePlatforms, region Region, secureName string, tags []string, type_ InstanceType, worldId string) *Instance {
	this := Instance{}
	this.Active = active
	this.CanRequestInvite = canRequestInvite
	this.Capacity = capacity
	this.ClientNumber = clientNumber
	this.Full = full
	this.Id = id
	this.InstanceId = instanceId
	this.Location = location
	this.NUsers = nUsers
	this.Name = name
	this.Permanent = permanent
	this.PhotonRegion = photonRegion
	this.Platforms = platforms
	this.Region = region
	this.SecureName = secureName
	this.Tags = tags
	this.Type = type_
	this.WorldId = worldId
	return &this
}

// NewInstanceWithDefaults instantiates a new Instance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceWithDefaults() *Instance {
	this := Instance{}
	var active bool = true
	this.Active = active
	var canRequestInvite bool = true
	this.CanRequestInvite = canRequestInvite
	var full bool = false
	this.Full = full
	var permanent bool = false
	this.Permanent = permanent
	var photonRegion Region = US
	this.PhotonRegion = photonRegion
	var region Region = US
	this.Region = region
	return &this
}

// GetActive returns the Active field value
func (o *Instance) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *Instance) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *Instance) SetActive(v bool) {
	o.Active = v
}

// GetCanRequestInvite returns the CanRequestInvite field value
func (o *Instance) GetCanRequestInvite() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanRequestInvite
}

// GetCanRequestInviteOk returns a tuple with the CanRequestInvite field value
// and a boolean to check if the value has been set.
func (o *Instance) GetCanRequestInviteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanRequestInvite, true
}

// SetCanRequestInvite sets field value
func (o *Instance) SetCanRequestInvite(v bool) {
	o.CanRequestInvite = v
}

// GetCapacity returns the Capacity field value
func (o *Instance) GetCapacity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
func (o *Instance) GetCapacityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capacity, true
}

// SetCapacity sets field value
func (o *Instance) SetCapacity(v int32) {
	o.Capacity = v
}

// GetClientNumber returns the ClientNumber field value
// Deprecated
func (o *Instance) GetClientNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientNumber
}

// GetClientNumberOk returns a tuple with the ClientNumber field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *Instance) GetClientNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientNumber, true
}

// SetClientNumber sets field value
// Deprecated
func (o *Instance) SetClientNumber(v string) {
	o.ClientNumber = v
}

// GetFull returns the Full field value
func (o *Instance) GetFull() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Full
}

// GetFullOk returns a tuple with the Full field value
// and a boolean to check if the value has been set.
func (o *Instance) GetFullOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Full, true
}

// SetFull sets field value
func (o *Instance) SetFull(v bool) {
	o.Full = v
}

// GetId returns the Id field value
func (o *Instance) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Instance) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Instance) SetId(v string) {
	o.Id = v
}

// GetInstanceId returns the InstanceId field value
func (o *Instance) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *Instance) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *Instance) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetLocation returns the Location field value
func (o *Instance) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Instance) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Instance) SetLocation(v string) {
	o.Location = v
}

// GetNUsers returns the NUsers field value
func (o *Instance) GetNUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NUsers
}

// GetNUsersOk returns a tuple with the NUsers field value
// and a boolean to check if the value has been set.
func (o *Instance) GetNUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NUsers, true
}

// SetNUsers sets field value
func (o *Instance) SetNUsers(v int32) {
	o.NUsers = v
}

// GetName returns the Name field value
func (o *Instance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Instance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Instance) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *Instance) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *Instance) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *Instance) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetPermanent returns the Permanent field value
func (o *Instance) GetPermanent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Permanent
}

// GetPermanentOk returns a tuple with the Permanent field value
// and a boolean to check if the value has been set.
func (o *Instance) GetPermanentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permanent, true
}

// SetPermanent sets field value
func (o *Instance) SetPermanent(v bool) {
	o.Permanent = v
}

// GetPhotonRegion returns the PhotonRegion field value
func (o *Instance) GetPhotonRegion() Region {
	if o == nil {
		var ret Region
		return ret
	}

	return o.PhotonRegion
}

// GetPhotonRegionOk returns a tuple with the PhotonRegion field value
// and a boolean to check if the value has been set.
func (o *Instance) GetPhotonRegionOk() (*Region, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhotonRegion, true
}

// SetPhotonRegion sets field value
func (o *Instance) SetPhotonRegion(v Region) {
	o.PhotonRegion = v
}

// GetPlatforms returns the Platforms field value
func (o *Instance) GetPlatforms() InstancePlatforms {
	if o == nil {
		var ret InstancePlatforms
		return ret
	}

	return o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value
// and a boolean to check if the value has been set.
func (o *Instance) GetPlatformsOk() (*InstancePlatforms, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platforms, true
}

// SetPlatforms sets field value
func (o *Instance) SetPlatforms(v InstancePlatforms) {
	o.Platforms = v
}

// GetRegion returns the Region field value
func (o *Instance) GetRegion() Region {
	if o == nil {
		var ret Region
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *Instance) GetRegionOk() (*Region, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *Instance) SetRegion(v Region) {
	o.Region = v
}

// GetSecureName returns the SecureName field value
func (o *Instance) GetSecureName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecureName
}

// GetSecureNameOk returns a tuple with the SecureName field value
// and a boolean to check if the value has been set.
func (o *Instance) GetSecureNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecureName, true
}

// SetSecureName sets field value
func (o *Instance) SetSecureName(v string) {
	o.SecureName = v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *Instance) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *Instance) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *Instance) SetShortName(v string) {
	o.ShortName = &v
}

// GetTags returns the Tags field value
func (o *Instance) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Instance) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Instance) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *Instance) GetType() InstanceType {
	if o == nil {
		var ret InstanceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Instance) GetTypeOk() (*InstanceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Instance) SetType(v InstanceType) {
	o.Type = v
}

// GetWorldId returns the WorldId field value
func (o *Instance) GetWorldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorldId
}

// GetWorldIdOk returns a tuple with the WorldId field value
// and a boolean to check if the value has been set.
func (o *Instance) GetWorldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorldId, true
}

// SetWorldId sets field value
func (o *Instance) SetWorldId(v string) {
	o.WorldId = v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *Instance) GetHidden() string {
	if o == nil || IsNil(o.Hidden) {
		var ret string
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetHiddenOk() (*string, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *Instance) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given string and assigns it to the Hidden field.
func (o *Instance) SetHidden(v string) {
	o.Hidden = &v
}

// GetFriends returns the Friends field value if set, zero value otherwise.
func (o *Instance) GetFriends() string {
	if o == nil || IsNil(o.Friends) {
		var ret string
		return ret
	}
	return *o.Friends
}

// GetFriendsOk returns a tuple with the Friends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetFriendsOk() (*string, bool) {
	if o == nil || IsNil(o.Friends) {
		return nil, false
	}
	return o.Friends, true
}

// HasFriends returns a boolean if a field has been set.
func (o *Instance) HasFriends() bool {
	if o != nil && !IsNil(o.Friends) {
		return true
	}

	return false
}

// SetFriends gets a reference to the given string and assigns it to the Friends field.
func (o *Instance) SetFriends(v string) {
	o.Friends = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *Instance) GetPrivate() string {
	if o == nil || IsNil(o.Private) {
		var ret string
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetPrivateOk() (*string, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *Instance) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given string and assigns it to the Private field.
func (o *Instance) SetPrivate(v string) {
	o.Private = &v
}

func (o Instance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Instance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["canRequestInvite"] = o.CanRequestInvite
	toSerialize["capacity"] = o.Capacity
	toSerialize["clientNumber"] = o.ClientNumber
	toSerialize["full"] = o.Full
	toSerialize["id"] = o.Id
	toSerialize["instanceId"] = o.InstanceId
	toSerialize["location"] = o.Location
	toSerialize["n_users"] = o.NUsers
	toSerialize["name"] = o.Name
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	toSerialize["permanent"] = o.Permanent
	toSerialize["photonRegion"] = o.PhotonRegion
	toSerialize["platforms"] = o.Platforms
	toSerialize["region"] = o.Region
	toSerialize["secureName"] = o.SecureName
	if !IsNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	toSerialize["tags"] = o.Tags
	toSerialize["type"] = o.Type
	toSerialize["worldId"] = o.WorldId
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.Friends) {
		toSerialize["friends"] = o.Friends
	}
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	return toSerialize, nil
}

type NullableInstance struct {
	value *Instance
	isSet bool
}

func (v NullableInstance) Get() *Instance {
	return v.value
}

func (v *NullableInstance) Set(val *Instance) {
	v.value = val
	v.isSet = true
}

func (v NullableInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstance(val *Instance) *NullableInstance {
	return &NullableInstance{value: val, isSet: true}
}

func (v NullableInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


