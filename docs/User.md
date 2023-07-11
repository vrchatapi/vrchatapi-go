# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAvatarCopying** | **bool** |  | [default to true]
**Bio** | **string** |  | 
**BioLinks** | **[]string** |  | 
**CurrentAvatarImageUrl** | **string** | When profilePicOverride is not empty, use it instead. | 
**CurrentAvatarThumbnailImageUrl** | **string** | When profilePicOverride is not empty, use it instead. | 
**DateJoined** | **string** |  | 
**DeveloperType** | [**DeveloperType**](DeveloperType.md) |  | [default to NONE]
**DisplayName** | **string** | A users visual display name. This is what shows up in-game, and can different from their &#x60;username&#x60;. Changing display name is restricted to a cooldown period. | 
**FriendKey** | **string** |  | 
**FriendRequestStatus** | Pointer to **string** |  | [optional] 
**Id** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**InstanceId** | Pointer to **string** | InstanceID can be \&quot;offline\&quot; on User profiles if you are not friends with that user and \&quot;private\&quot; if you are friends and user is in private instance. | [optional] 
**IsFriend** | **bool** | Either their &#x60;friendKey&#x60;, or empty string if you are not friends. Unknown usage. | 
**LastActivity** | **string** | Either a date-time or empty string. | 
**LastLogin** | **string** | Either a date-time or empty string. | 
**LastPlatform** | **string** | This can be &#x60;standalonewindows&#x60; or &#x60;android&#x60;, but can also pretty much be any random Unity verison such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | 
**Location** | Pointer to **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**ProfilePicOverride** | **string** |  | 
**State** | [**UserState**](UserState.md) |  | [default to OFFLINE]
**Status** | [**UserStatus**](UserStatus.md) |  | [default to OFFLINE]
**StatusDescription** | **string** |  | 
**Tags** | **[]string** |   | 
**TravelingToInstance** | Pointer to **string** |  | [optional] 
**TravelingToLocation** | Pointer to **string** |  | [optional] 
**TravelingToWorld** | Pointer to **string** |  | [optional] 
**UserIcon** | **string** |  | 
**Username** | Pointer to **string** | -| A users unique name, used during login. This is different from &#x60;displayName&#x60; which is what shows up in-game. A users &#x60;username&#x60; can never be changed.&#39; **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429). | [optional] 
**WorldId** | Pointer to **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | [optional] 

## Methods

### NewUser

`func NewUser(allowAvatarCopying bool, bio string, bioLinks []string, currentAvatarImageUrl string, currentAvatarThumbnailImageUrl string, dateJoined string, developerType DeveloperType, displayName string, friendKey string, id string, isFriend bool, lastActivity string, lastLogin string, lastPlatform string, profilePicOverride string, state UserState, status UserStatus, statusDescription string, tags []string, userIcon string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAvatarCopying

`func (o *User) GetAllowAvatarCopying() bool`

GetAllowAvatarCopying returns the AllowAvatarCopying field if non-nil, zero value otherwise.

### GetAllowAvatarCopyingOk

`func (o *User) GetAllowAvatarCopyingOk() (*bool, bool)`

GetAllowAvatarCopyingOk returns a tuple with the AllowAvatarCopying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAvatarCopying

`func (o *User) SetAllowAvatarCopying(v bool)`

SetAllowAvatarCopying sets AllowAvatarCopying field to given value.


### GetBio

`func (o *User) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *User) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *User) SetBio(v string)`

SetBio sets Bio field to given value.


### GetBioLinks

`func (o *User) GetBioLinks() []string`

GetBioLinks returns the BioLinks field if non-nil, zero value otherwise.

### GetBioLinksOk

`func (o *User) GetBioLinksOk() (*[]string, bool)`

GetBioLinksOk returns a tuple with the BioLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioLinks

`func (o *User) SetBioLinks(v []string)`

SetBioLinks sets BioLinks field to given value.


### GetCurrentAvatarImageUrl

`func (o *User) GetCurrentAvatarImageUrl() string`

GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarImageUrlOk

`func (o *User) GetCurrentAvatarImageUrlOk() (*string, bool)`

GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarImageUrl

`func (o *User) SetCurrentAvatarImageUrl(v string)`

SetCurrentAvatarImageUrl sets CurrentAvatarImageUrl field to given value.


### GetCurrentAvatarThumbnailImageUrl

`func (o *User) GetCurrentAvatarThumbnailImageUrl() string`

GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarThumbnailImageUrlOk

`func (o *User) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool)`

GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarThumbnailImageUrl

`func (o *User) SetCurrentAvatarThumbnailImageUrl(v string)`

SetCurrentAvatarThumbnailImageUrl sets CurrentAvatarThumbnailImageUrl field to given value.


### GetDateJoined

`func (o *User) GetDateJoined() string`

GetDateJoined returns the DateJoined field if non-nil, zero value otherwise.

### GetDateJoinedOk

`func (o *User) GetDateJoinedOk() (*string, bool)`

GetDateJoinedOk returns a tuple with the DateJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateJoined

`func (o *User) SetDateJoined(v string)`

SetDateJoined sets DateJoined field to given value.


### GetDeveloperType

`func (o *User) GetDeveloperType() DeveloperType`

GetDeveloperType returns the DeveloperType field if non-nil, zero value otherwise.

### GetDeveloperTypeOk

`func (o *User) GetDeveloperTypeOk() (*DeveloperType, bool)`

GetDeveloperTypeOk returns a tuple with the DeveloperType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperType

`func (o *User) SetDeveloperType(v DeveloperType)`

SetDeveloperType sets DeveloperType field to given value.


### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetFriendKey

`func (o *User) GetFriendKey() string`

GetFriendKey returns the FriendKey field if non-nil, zero value otherwise.

### GetFriendKeyOk

`func (o *User) GetFriendKeyOk() (*string, bool)`

GetFriendKeyOk returns a tuple with the FriendKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendKey

`func (o *User) SetFriendKey(v string)`

SetFriendKey sets FriendKey field to given value.


### GetFriendRequestStatus

`func (o *User) GetFriendRequestStatus() string`

GetFriendRequestStatus returns the FriendRequestStatus field if non-nil, zero value otherwise.

### GetFriendRequestStatusOk

`func (o *User) GetFriendRequestStatusOk() (*string, bool)`

GetFriendRequestStatusOk returns a tuple with the FriendRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendRequestStatus

`func (o *User) SetFriendRequestStatus(v string)`

SetFriendRequestStatus sets FriendRequestStatus field to given value.

### HasFriendRequestStatus

`func (o *User) HasFriendRequestStatus() bool`

HasFriendRequestStatus returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *User) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *User) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *User) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *User) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetIsFriend

`func (o *User) GetIsFriend() bool`

GetIsFriend returns the IsFriend field if non-nil, zero value otherwise.

### GetIsFriendOk

`func (o *User) GetIsFriendOk() (*bool, bool)`

GetIsFriendOk returns a tuple with the IsFriend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFriend

`func (o *User) SetIsFriend(v bool)`

SetIsFriend sets IsFriend field to given value.


### GetLastActivity

`func (o *User) GetLastActivity() string`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *User) GetLastActivityOk() (*string, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *User) SetLastActivity(v string)`

SetLastActivity sets LastActivity field to given value.


### GetLastLogin

`func (o *User) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *User) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *User) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.


### GetLastPlatform

`func (o *User) GetLastPlatform() string`

GetLastPlatform returns the LastPlatform field if non-nil, zero value otherwise.

### GetLastPlatformOk

`func (o *User) GetLastPlatformOk() (*string, bool)`

GetLastPlatformOk returns a tuple with the LastPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlatform

`func (o *User) SetLastPlatform(v string)`

SetLastPlatform sets LastPlatform field to given value.


### GetLocation

`func (o *User) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *User) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *User) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *User) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetNote

`func (o *User) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *User) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *User) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *User) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetProfilePicOverride

`func (o *User) GetProfilePicOverride() string`

GetProfilePicOverride returns the ProfilePicOverride field if non-nil, zero value otherwise.

### GetProfilePicOverrideOk

`func (o *User) GetProfilePicOverrideOk() (*string, bool)`

GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicOverride

`func (o *User) SetProfilePicOverride(v string)`

SetProfilePicOverride sets ProfilePicOverride field to given value.


### GetState

`func (o *User) GetState() UserState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *User) GetStateOk() (*UserState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *User) SetState(v UserState)`

SetState sets State field to given value.


### GetStatus

`func (o *User) GetStatus() UserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *User) GetStatusOk() (*UserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *User) SetStatus(v UserStatus)`

SetStatus sets Status field to given value.


### GetStatusDescription

`func (o *User) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *User) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *User) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.


### GetTags

`func (o *User) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *User) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *User) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTravelingToInstance

`func (o *User) GetTravelingToInstance() string`

GetTravelingToInstance returns the TravelingToInstance field if non-nil, zero value otherwise.

### GetTravelingToInstanceOk

`func (o *User) GetTravelingToInstanceOk() (*string, bool)`

GetTravelingToInstanceOk returns a tuple with the TravelingToInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelingToInstance

`func (o *User) SetTravelingToInstance(v string)`

SetTravelingToInstance sets TravelingToInstance field to given value.

### HasTravelingToInstance

`func (o *User) HasTravelingToInstance() bool`

HasTravelingToInstance returns a boolean if a field has been set.

### GetTravelingToLocation

`func (o *User) GetTravelingToLocation() string`

GetTravelingToLocation returns the TravelingToLocation field if non-nil, zero value otherwise.

### GetTravelingToLocationOk

`func (o *User) GetTravelingToLocationOk() (*string, bool)`

GetTravelingToLocationOk returns a tuple with the TravelingToLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelingToLocation

`func (o *User) SetTravelingToLocation(v string)`

SetTravelingToLocation sets TravelingToLocation field to given value.

### HasTravelingToLocation

`func (o *User) HasTravelingToLocation() bool`

HasTravelingToLocation returns a boolean if a field has been set.

### GetTravelingToWorld

`func (o *User) GetTravelingToWorld() string`

GetTravelingToWorld returns the TravelingToWorld field if non-nil, zero value otherwise.

### GetTravelingToWorldOk

`func (o *User) GetTravelingToWorldOk() (*string, bool)`

GetTravelingToWorldOk returns a tuple with the TravelingToWorld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelingToWorld

`func (o *User) SetTravelingToWorld(v string)`

SetTravelingToWorld sets TravelingToWorld field to given value.

### HasTravelingToWorld

`func (o *User) HasTravelingToWorld() bool`

HasTravelingToWorld returns a boolean if a field has been set.

### GetUserIcon

`func (o *User) GetUserIcon() string`

GetUserIcon returns the UserIcon field if non-nil, zero value otherwise.

### GetUserIconOk

`func (o *User) GetUserIconOk() (*string, bool)`

GetUserIconOk returns a tuple with the UserIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIcon

`func (o *User) SetUserIcon(v string)`

SetUserIcon sets UserIcon field to given value.


### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetWorldId

`func (o *User) GetWorldId() string`

GetWorldId returns the WorldId field if non-nil, zero value otherwise.

### GetWorldIdOk

`func (o *User) GetWorldIdOk() (*string, bool)`

GetWorldIdOk returns a tuple with the WorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldId

`func (o *User) SetWorldId(v string)`

SetWorldId sets WorldId field to given value.

### HasWorldId

`func (o *User) HasWorldId() bool`

HasWorldId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


