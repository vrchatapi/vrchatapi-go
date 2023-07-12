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

// checks if the LimitedUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LimitedUser{}

// LimitedUser 
type LimitedUser struct {
	Bio *string `json:"bio,omitempty"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarImageUrl string `json:"currentAvatarImageUrl"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarThumbnailImageUrl string `json:"currentAvatarThumbnailImageUrl"`
	DeveloperType DeveloperType `json:"developerType"`
	DisplayName string `json:"displayName"`
	FallbackAvatar string `json:"fallbackAvatar"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id string `json:"id"`
	IsFriend bool `json:"isFriend"`
	// This can be `standalonewindows` or `android`, but can also pretty much be any random Unity verison such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	LastPlatform string `json:"last_platform"`
	ProfilePicOverride string `json:"profilePicOverride"`
	Status UserStatus `json:"status"`
	StatusDescription string `json:"statusDescription"`
	// <- Always empty.
	Tags []string `json:"tags"`
	UserIcon string `json:"userIcon"`
	// -| **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429).
	// Deprecated
	Username *string `json:"username,omitempty"`
	Location *string `json:"location,omitempty"`
	FriendKey *string `json:"friendKey,omitempty"`
}

// NewLimitedUser instantiates a new LimitedUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimitedUser(currentAvatarImageUrl string, currentAvatarThumbnailImageUrl string, developerType DeveloperType, displayName string, fallbackAvatar string, id string, isFriend bool, lastPlatform string, profilePicOverride string, status UserStatus, statusDescription string, tags []string, userIcon string) *LimitedUser {
	this := LimitedUser{}
	this.CurrentAvatarImageUrl = currentAvatarImageUrl
	this.CurrentAvatarThumbnailImageUrl = currentAvatarThumbnailImageUrl
	this.DeveloperType = developerType
	this.DisplayName = displayName
	this.FallbackAvatar = fallbackAvatar
	this.Id = id
	this.IsFriend = isFriend
	this.LastPlatform = lastPlatform
	this.ProfilePicOverride = profilePicOverride
	this.Status = status
	this.StatusDescription = statusDescription
	this.Tags = tags
	this.UserIcon = userIcon
	return &this
}

// NewLimitedUserWithDefaults instantiates a new LimitedUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitedUserWithDefaults() *LimitedUser {
	this := LimitedUser{}
	var developerType DeveloperType = NONE
	this.DeveloperType = developerType
	var status UserStatus = OFFLINE
	this.Status = status
	return &this
}

// GetBio returns the Bio field value if set, zero value otherwise.
func (o *LimitedUser) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *LimitedUser) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *LimitedUser) SetBio(v string) {
	o.Bio = &v
}

// GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field value
func (o *LimitedUser) GetCurrentAvatarImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAvatarImageUrl
}

// GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field value
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetCurrentAvatarImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentAvatarImageUrl, true
}

// SetCurrentAvatarImageUrl sets field value
func (o *LimitedUser) SetCurrentAvatarImageUrl(v string) {
	o.CurrentAvatarImageUrl = v
}

// GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field value
func (o *LimitedUser) GetCurrentAvatarThumbnailImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAvatarThumbnailImageUrl
}

// GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field value
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentAvatarThumbnailImageUrl, true
}

// SetCurrentAvatarThumbnailImageUrl sets field value
func (o *LimitedUser) SetCurrentAvatarThumbnailImageUrl(v string) {
	o.CurrentAvatarThumbnailImageUrl = v
}

// GetDeveloperType returns the DeveloperType field value
func (o *LimitedUser) GetDeveloperType() DeveloperType {
	if o == nil {
		var ret DeveloperType
		return ret
	}

	return o.DeveloperType
}

// GetDeveloperTypeOk returns a tuple with the DeveloperType field value
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetDeveloperTypeOk() (*DeveloperType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeveloperType, true
}

// SetDeveloperType sets field value
func (o *LimitedUser) SetDeveloperType(v DeveloperType) {
	o.DeveloperType = v
}

// GetDisplayName returns the DisplayName field value
func (o *LimitedUser) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *LimitedUser) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetFallbackAvatar returns the FallbackAvatar field value
func (o *LimitedUser) GetFallbackAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FallbackAvatar
}

// GetFallbackAvatarOk returns a tuple with the FallbackAvatar field value
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetFallbackAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FallbackAvatar, true
}

// SetFallbackAvatar sets field value
func (o *LimitedUser) SetFallbackAvatar(v string) {
	o.FallbackAvatar = v
}

// GetId returns the Id field value
func (o *LimitedUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LimitedUser) SetId(v string) {
	o.Id = v
}

// GetIsFriend returns the IsFriend field value
func (o *LimitedUser) GetIsFriend() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFriend
}

// GetIsFriendOk returns a tuple with the IsFriend field value
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetIsFriendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFriend, true
}

// SetIsFriend sets field value
func (o *LimitedUser) SetIsFriend(v bool) {
	o.IsFriend = v
}

// GetLastPlatform returns the LastPlatform field value
func (o *LimitedUser) GetLastPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastPlatform
}

// GetLastPlatformOk returns a tuple with the LastPlatform field value
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetLastPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastPlatform, true
}

// SetLastPlatform sets field value
func (o *LimitedUser) SetLastPlatform(v string) {
	o.LastPlatform = v
}

// GetProfilePicOverride returns the ProfilePicOverride field value
func (o *LimitedUser) GetProfilePicOverride() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfilePicOverride
}

// GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field value
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetProfilePicOverrideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfilePicOverride, true
}

// SetProfilePicOverride sets field value
func (o *LimitedUser) SetProfilePicOverride(v string) {
	o.ProfilePicOverride = v
}

// GetStatus returns the Status field value
func (o *LimitedUser) GetStatus() UserStatus {
	if o == nil {
		var ret UserStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetStatusOk() (*UserStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LimitedUser) SetStatus(v UserStatus) {
	o.Status = v
}

// GetStatusDescription returns the StatusDescription field value
func (o *LimitedUser) GetStatusDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetStatusDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDescription, true
}

// SetStatusDescription sets field value
func (o *LimitedUser) SetStatusDescription(v string) {
	o.StatusDescription = v
}

// GetTags returns the Tags field value
func (o *LimitedUser) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *LimitedUser) SetTags(v []string) {
	o.Tags = v
}

// GetUserIcon returns the UserIcon field value
func (o *LimitedUser) GetUserIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserIcon
}

// GetUserIconOk returns a tuple with the UserIcon field value
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetUserIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserIcon, true
}

// SetUserIcon sets field value
func (o *LimitedUser) SetUserIcon(v string) {
	o.UserIcon = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
// Deprecated
func (o *LimitedUser) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *LimitedUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *LimitedUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
// Deprecated
func (o *LimitedUser) SetUsername(v string) {
	o.Username = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *LimitedUser) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *LimitedUser) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *LimitedUser) SetLocation(v string) {
	o.Location = &v
}

// GetFriendKey returns the FriendKey field value if set, zero value otherwise.
func (o *LimitedUser) GetFriendKey() string {
	if o == nil || IsNil(o.FriendKey) {
		var ret string
		return ret
	}
	return *o.FriendKey
}

// GetFriendKeyOk returns a tuple with the FriendKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUser) GetFriendKeyOk() (*string, bool) {
	if o == nil || IsNil(o.FriendKey) {
		return nil, false
	}
	return o.FriendKey, true
}

// HasFriendKey returns a boolean if a field has been set.
func (o *LimitedUser) HasFriendKey() bool {
	if o != nil && !IsNil(o.FriendKey) {
		return true
	}

	return false
}

// SetFriendKey gets a reference to the given string and assigns it to the FriendKey field.
func (o *LimitedUser) SetFriendKey(v string) {
	o.FriendKey = &v
}

func (o LimitedUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LimitedUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	toSerialize["currentAvatarImageUrl"] = o.CurrentAvatarImageUrl
	toSerialize["currentAvatarThumbnailImageUrl"] = o.CurrentAvatarThumbnailImageUrl
	toSerialize["developerType"] = o.DeveloperType
	toSerialize["displayName"] = o.DisplayName
	toSerialize["fallbackAvatar"] = o.FallbackAvatar
	toSerialize["id"] = o.Id
	toSerialize["isFriend"] = o.IsFriend
	toSerialize["last_platform"] = o.LastPlatform
	toSerialize["profilePicOverride"] = o.ProfilePicOverride
	toSerialize["status"] = o.Status
	toSerialize["statusDescription"] = o.StatusDescription
	toSerialize["tags"] = o.Tags
	toSerialize["userIcon"] = o.UserIcon
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.FriendKey) {
		toSerialize["friendKey"] = o.FriendKey
	}
	return toSerialize, nil
}

type NullableLimitedUser struct {
	value *LimitedUser
	isSet bool
}

func (v NullableLimitedUser) Get() *LimitedUser {
	return v.value
}

func (v *NullableLimitedUser) Set(val *LimitedUser) {
	v.value = val
	v.isSet = true
}

func (v NullableLimitedUser) IsSet() bool {
	return v.isSet
}

func (v *NullableLimitedUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimitedUser(val *LimitedUser) *NullableLimitedUser {
	return &NullableLimitedUser{value: val, isSet: true}
}

func (v NullableLimitedUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimitedUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


