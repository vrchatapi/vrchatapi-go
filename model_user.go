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

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User struct for User
type User struct {
	AllowAvatarCopying bool `json:"allowAvatarCopying"`
	Bio string `json:"bio"`
	BioLinks []string `json:"bioLinks"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarImageUrl string `json:"currentAvatarImageUrl"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarThumbnailImageUrl string `json:"currentAvatarThumbnailImageUrl"`
	DateJoined string `json:"date_joined"`
	DeveloperType DeveloperType `json:"developerType"`
	// A users visual display name. This is what shows up in-game, and can different from their `username`. Changing display name is restricted to a cooldown period.
	DisplayName string `json:"displayName"`
	FriendKey string `json:"friendKey"`
	FriendRequestStatus *string `json:"friendRequestStatus,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id string `json:"id"`
	// InstanceID can be \"offline\" on User profiles if you are not friends with that user and \"private\" if you are friends and user is in private instance.
	InstanceId *string `json:"instanceId,omitempty"`
	// Either their `friendKey`, or empty string if you are not friends. Unknown usage.
	IsFriend bool `json:"isFriend"`
	// Either a date-time or empty string.
	LastActivity string `json:"last_activity"`
	// Either a date-time or empty string.
	LastLogin string `json:"last_login"`
	// This can be `standalonewindows` or `android`, but can also pretty much be any random Unity verison such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	LastPlatform string `json:"last_platform"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	Location *string `json:"location,omitempty"`
	Note *string `json:"note,omitempty"`
	ProfilePicOverride string `json:"profilePicOverride"`
	State UserState `json:"state"`
	Status UserStatus `json:"status"`
	StatusDescription string `json:"statusDescription"`
	//  
	Tags []string `json:"tags"`
	TravelingToInstance *string `json:"travelingToInstance,omitempty"`
	TravelingToLocation *string `json:"travelingToLocation,omitempty"`
	TravelingToWorld *string `json:"travelingToWorld,omitempty"`
	UserIcon string `json:"userIcon"`
	// -| A users unique name, used during login. This is different from `displayName` which is what shows up in-game. A users `username` can never be changed.' **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429).
	// Deprecated
	Username *string `json:"username,omitempty"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	WorldId *string `json:"worldId,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(allowAvatarCopying bool, bio string, bioLinks []string, currentAvatarImageUrl string, currentAvatarThumbnailImageUrl string, dateJoined string, developerType DeveloperType, displayName string, friendKey string, id string, isFriend bool, lastActivity string, lastLogin string, lastPlatform string, profilePicOverride string, state UserState, status UserStatus, statusDescription string, tags []string, userIcon string) *User {
	this := User{}
	this.AllowAvatarCopying = allowAvatarCopying
	this.Bio = bio
	this.BioLinks = bioLinks
	this.CurrentAvatarImageUrl = currentAvatarImageUrl
	this.CurrentAvatarThumbnailImageUrl = currentAvatarThumbnailImageUrl
	this.DateJoined = dateJoined
	this.DeveloperType = developerType
	this.DisplayName = displayName
	this.FriendKey = friendKey
	this.Id = id
	this.IsFriend = isFriend
	this.LastActivity = lastActivity
	this.LastLogin = lastLogin
	this.LastPlatform = lastPlatform
	this.ProfilePicOverride = profilePicOverride
	this.State = state
	this.Status = status
	this.StatusDescription = statusDescription
	this.Tags = tags
	this.UserIcon = userIcon
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	var allowAvatarCopying bool = true
	this.AllowAvatarCopying = allowAvatarCopying
	var developerType DeveloperType = NONE
	this.DeveloperType = developerType
	var state UserState = OFFLINE
	this.State = state
	var status UserStatus = OFFLINE
	this.Status = status
	return &this
}

// GetAllowAvatarCopying returns the AllowAvatarCopying field value
func (o *User) GetAllowAvatarCopying() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowAvatarCopying
}

// GetAllowAvatarCopyingOk returns a tuple with the AllowAvatarCopying field value
// and a boolean to check if the value has been set.
func (o *User) GetAllowAvatarCopyingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowAvatarCopying, true
}

// SetAllowAvatarCopying sets field value
func (o *User) SetAllowAvatarCopying(v bool) {
	o.AllowAvatarCopying = v
}

// GetBio returns the Bio field value
func (o *User) GetBio() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bio
}

// GetBioOk returns a tuple with the Bio field value
// and a boolean to check if the value has been set.
func (o *User) GetBioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bio, true
}

// SetBio sets field value
func (o *User) SetBio(v string) {
	o.Bio = v
}

// GetBioLinks returns the BioLinks field value
func (o *User) GetBioLinks() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BioLinks
}

// GetBioLinksOk returns a tuple with the BioLinks field value
// and a boolean to check if the value has been set.
func (o *User) GetBioLinksOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BioLinks, true
}

// SetBioLinks sets field value
func (o *User) SetBioLinks(v []string) {
	o.BioLinks = v
}

// GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field value
func (o *User) GetCurrentAvatarImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAvatarImageUrl
}

// GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field value
// and a boolean to check if the value has been set.
func (o *User) GetCurrentAvatarImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentAvatarImageUrl, true
}

// SetCurrentAvatarImageUrl sets field value
func (o *User) SetCurrentAvatarImageUrl(v string) {
	o.CurrentAvatarImageUrl = v
}

// GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field value
func (o *User) GetCurrentAvatarThumbnailImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAvatarThumbnailImageUrl
}

// GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field value
// and a boolean to check if the value has been set.
func (o *User) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentAvatarThumbnailImageUrl, true
}

// SetCurrentAvatarThumbnailImageUrl sets field value
func (o *User) SetCurrentAvatarThumbnailImageUrl(v string) {
	o.CurrentAvatarThumbnailImageUrl = v
}

// GetDateJoined returns the DateJoined field value
func (o *User) GetDateJoined() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateJoined
}

// GetDateJoinedOk returns a tuple with the DateJoined field value
// and a boolean to check if the value has been set.
func (o *User) GetDateJoinedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateJoined, true
}

// SetDateJoined sets field value
func (o *User) SetDateJoined(v string) {
	o.DateJoined = v
}

// GetDeveloperType returns the DeveloperType field value
func (o *User) GetDeveloperType() DeveloperType {
	if o == nil {
		var ret DeveloperType
		return ret
	}

	return o.DeveloperType
}

// GetDeveloperTypeOk returns a tuple with the DeveloperType field value
// and a boolean to check if the value has been set.
func (o *User) GetDeveloperTypeOk() (*DeveloperType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeveloperType, true
}

// SetDeveloperType sets field value
func (o *User) SetDeveloperType(v DeveloperType) {
	o.DeveloperType = v
}

// GetDisplayName returns the DisplayName field value
func (o *User) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *User) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *User) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetFriendKey returns the FriendKey field value
func (o *User) GetFriendKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FriendKey
}

// GetFriendKeyOk returns a tuple with the FriendKey field value
// and a boolean to check if the value has been set.
func (o *User) GetFriendKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FriendKey, true
}

// SetFriendKey sets field value
func (o *User) SetFriendKey(v string) {
	o.FriendKey = v
}

// GetFriendRequestStatus returns the FriendRequestStatus field value if set, zero value otherwise.
func (o *User) GetFriendRequestStatus() string {
	if o == nil || IsNil(o.FriendRequestStatus) {
		var ret string
		return ret
	}
	return *o.FriendRequestStatus
}

// GetFriendRequestStatusOk returns a tuple with the FriendRequestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFriendRequestStatusOk() (*string, bool) {
	if o == nil || IsNil(o.FriendRequestStatus) {
		return nil, false
	}
	return o.FriendRequestStatus, true
}

// HasFriendRequestStatus returns a boolean if a field has been set.
func (o *User) HasFriendRequestStatus() bool {
	if o != nil && !IsNil(o.FriendRequestStatus) {
		return true
	}

	return false
}

// SetFriendRequestStatus gets a reference to the given string and assigns it to the FriendRequestStatus field.
func (o *User) SetFriendRequestStatus(v string) {
	o.FriendRequestStatus = &v
}

// GetId returns the Id field value
func (o *User) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *User) SetId(v string) {
	o.Id = v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *User) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *User) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *User) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetIsFriend returns the IsFriend field value
func (o *User) GetIsFriend() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFriend
}

// GetIsFriendOk returns a tuple with the IsFriend field value
// and a boolean to check if the value has been set.
func (o *User) GetIsFriendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFriend, true
}

// SetIsFriend sets field value
func (o *User) SetIsFriend(v bool) {
	o.IsFriend = v
}

// GetLastActivity returns the LastActivity field value
func (o *User) GetLastActivity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastActivity
}

// GetLastActivityOk returns a tuple with the LastActivity field value
// and a boolean to check if the value has been set.
func (o *User) GetLastActivityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastActivity, true
}

// SetLastActivity sets field value
func (o *User) SetLastActivity(v string) {
	o.LastActivity = v
}

// GetLastLogin returns the LastLogin field value
func (o *User) GetLastLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value
// and a boolean to check if the value has been set.
func (o *User) GetLastLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastLogin, true
}

// SetLastLogin sets field value
func (o *User) SetLastLogin(v string) {
	o.LastLogin = v
}

// GetLastPlatform returns the LastPlatform field value
func (o *User) GetLastPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastPlatform
}

// GetLastPlatformOk returns a tuple with the LastPlatform field value
// and a boolean to check if the value has been set.
func (o *User) GetLastPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastPlatform, true
}

// SetLastPlatform sets field value
func (o *User) SetLastPlatform(v string) {
	o.LastPlatform = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *User) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *User) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *User) SetLocation(v string) {
	o.Location = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *User) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *User) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *User) SetNote(v string) {
	o.Note = &v
}

// GetProfilePicOverride returns the ProfilePicOverride field value
func (o *User) GetProfilePicOverride() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfilePicOverride
}

// GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field value
// and a boolean to check if the value has been set.
func (o *User) GetProfilePicOverrideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfilePicOverride, true
}

// SetProfilePicOverride sets field value
func (o *User) SetProfilePicOverride(v string) {
	o.ProfilePicOverride = v
}

// GetState returns the State field value
func (o *User) GetState() UserState {
	if o == nil {
		var ret UserState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *User) GetStateOk() (*UserState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *User) SetState(v UserState) {
	o.State = v
}

// GetStatus returns the Status field value
func (o *User) GetStatus() UserStatus {
	if o == nil {
		var ret UserStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *User) GetStatusOk() (*UserStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *User) SetStatus(v UserStatus) {
	o.Status = v
}

// GetStatusDescription returns the StatusDescription field value
func (o *User) GetStatusDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value
// and a boolean to check if the value has been set.
func (o *User) GetStatusDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDescription, true
}

// SetStatusDescription sets field value
func (o *User) SetStatusDescription(v string) {
	o.StatusDescription = v
}

// GetTags returns the Tags field value
func (o *User) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *User) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *User) SetTags(v []string) {
	o.Tags = v
}

// GetTravelingToInstance returns the TravelingToInstance field value if set, zero value otherwise.
func (o *User) GetTravelingToInstance() string {
	if o == nil || IsNil(o.TravelingToInstance) {
		var ret string
		return ret
	}
	return *o.TravelingToInstance
}

// GetTravelingToInstanceOk returns a tuple with the TravelingToInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTravelingToInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.TravelingToInstance) {
		return nil, false
	}
	return o.TravelingToInstance, true
}

// HasTravelingToInstance returns a boolean if a field has been set.
func (o *User) HasTravelingToInstance() bool {
	if o != nil && !IsNil(o.TravelingToInstance) {
		return true
	}

	return false
}

// SetTravelingToInstance gets a reference to the given string and assigns it to the TravelingToInstance field.
func (o *User) SetTravelingToInstance(v string) {
	o.TravelingToInstance = &v
}

// GetTravelingToLocation returns the TravelingToLocation field value if set, zero value otherwise.
func (o *User) GetTravelingToLocation() string {
	if o == nil || IsNil(o.TravelingToLocation) {
		var ret string
		return ret
	}
	return *o.TravelingToLocation
}

// GetTravelingToLocationOk returns a tuple with the TravelingToLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTravelingToLocationOk() (*string, bool) {
	if o == nil || IsNil(o.TravelingToLocation) {
		return nil, false
	}
	return o.TravelingToLocation, true
}

// HasTravelingToLocation returns a boolean if a field has been set.
func (o *User) HasTravelingToLocation() bool {
	if o != nil && !IsNil(o.TravelingToLocation) {
		return true
	}

	return false
}

// SetTravelingToLocation gets a reference to the given string and assigns it to the TravelingToLocation field.
func (o *User) SetTravelingToLocation(v string) {
	o.TravelingToLocation = &v
}

// GetTravelingToWorld returns the TravelingToWorld field value if set, zero value otherwise.
func (o *User) GetTravelingToWorld() string {
	if o == nil || IsNil(o.TravelingToWorld) {
		var ret string
		return ret
	}
	return *o.TravelingToWorld
}

// GetTravelingToWorldOk returns a tuple with the TravelingToWorld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTravelingToWorldOk() (*string, bool) {
	if o == nil || IsNil(o.TravelingToWorld) {
		return nil, false
	}
	return o.TravelingToWorld, true
}

// HasTravelingToWorld returns a boolean if a field has been set.
func (o *User) HasTravelingToWorld() bool {
	if o != nil && !IsNil(o.TravelingToWorld) {
		return true
	}

	return false
}

// SetTravelingToWorld gets a reference to the given string and assigns it to the TravelingToWorld field.
func (o *User) SetTravelingToWorld(v string) {
	o.TravelingToWorld = &v
}

// GetUserIcon returns the UserIcon field value
func (o *User) GetUserIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserIcon
}

// GetUserIconOk returns a tuple with the UserIcon field value
// and a boolean to check if the value has been set.
func (o *User) GetUserIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserIcon, true
}

// SetUserIcon sets field value
func (o *User) SetUserIcon(v string) {
	o.UserIcon = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
// Deprecated
func (o *User) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
// Deprecated
func (o *User) SetUsername(v string) {
	o.Username = &v
}

// GetWorldId returns the WorldId field value if set, zero value otherwise.
func (o *User) GetWorldId() string {
	if o == nil || IsNil(o.WorldId) {
		var ret string
		return ret
	}
	return *o.WorldId
}

// GetWorldIdOk returns a tuple with the WorldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetWorldIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorldId) {
		return nil, false
	}
	return o.WorldId, true
}

// HasWorldId returns a boolean if a field has been set.
func (o *User) HasWorldId() bool {
	if o != nil && !IsNil(o.WorldId) {
		return true
	}

	return false
}

// SetWorldId gets a reference to the given string and assigns it to the WorldId field.
func (o *User) SetWorldId(v string) {
	o.WorldId = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowAvatarCopying"] = o.AllowAvatarCopying
	toSerialize["bio"] = o.Bio
	toSerialize["bioLinks"] = o.BioLinks
	toSerialize["currentAvatarImageUrl"] = o.CurrentAvatarImageUrl
	toSerialize["currentAvatarThumbnailImageUrl"] = o.CurrentAvatarThumbnailImageUrl
	toSerialize["date_joined"] = o.DateJoined
	toSerialize["developerType"] = o.DeveloperType
	toSerialize["displayName"] = o.DisplayName
	toSerialize["friendKey"] = o.FriendKey
	if !IsNil(o.FriendRequestStatus) {
		toSerialize["friendRequestStatus"] = o.FriendRequestStatus
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	toSerialize["isFriend"] = o.IsFriend
	toSerialize["last_activity"] = o.LastActivity
	toSerialize["last_login"] = o.LastLogin
	toSerialize["last_platform"] = o.LastPlatform
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	toSerialize["profilePicOverride"] = o.ProfilePicOverride
	toSerialize["state"] = o.State
	toSerialize["status"] = o.Status
	toSerialize["statusDescription"] = o.StatusDescription
	toSerialize["tags"] = o.Tags
	if !IsNil(o.TravelingToInstance) {
		toSerialize["travelingToInstance"] = o.TravelingToInstance
	}
	if !IsNil(o.TravelingToLocation) {
		toSerialize["travelingToLocation"] = o.TravelingToLocation
	}
	if !IsNil(o.TravelingToWorld) {
		toSerialize["travelingToWorld"] = o.TravelingToWorld
	}
	toSerialize["userIcon"] = o.UserIcon
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.WorldId) {
		toSerialize["worldId"] = o.WorldId
	}
	return toSerialize, nil
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


