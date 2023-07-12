# LimitedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bio** | Pointer to **string** |  | [optional] 
**CurrentAvatarImageUrl** | **string** | When profilePicOverride is not empty, use it instead. | 
**CurrentAvatarThumbnailImageUrl** | **string** | When profilePicOverride is not empty, use it instead. | 
**DeveloperType** | [**DeveloperType**](DeveloperType.md) |  | [default to NONE]
**DisplayName** | **string** |  | 
**FallbackAvatar** | **string** |  | 
**Id** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**IsFriend** | **bool** |  | 
**LastPlatform** | **string** | This can be &#x60;standalonewindows&#x60; or &#x60;android&#x60;, but can also pretty much be any random Unity verison such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | 
**ProfilePicOverride** | **string** |  | 
**Status** | [**UserStatus**](UserStatus.md) |  | [default to OFFLINE]
**StatusDescription** | **string** |  | 
**Tags** | **[]string** | &lt;- Always empty. | 
**UserIcon** | **string** |  | 
**Username** | Pointer to **string** | -| **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429). | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**FriendKey** | Pointer to **string** |  | [optional] 

## Methods

### NewLimitedUser

`func NewLimitedUser(currentAvatarImageUrl string, currentAvatarThumbnailImageUrl string, developerType DeveloperType, displayName string, fallbackAvatar string, id string, isFriend bool, lastPlatform string, profilePicOverride string, status UserStatus, statusDescription string, tags []string, userIcon string, ) *LimitedUser`

NewLimitedUser instantiates a new LimitedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitedUserWithDefaults

`func NewLimitedUserWithDefaults() *LimitedUser`

NewLimitedUserWithDefaults instantiates a new LimitedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBio

`func (o *LimitedUser) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *LimitedUser) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *LimitedUser) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *LimitedUser) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetCurrentAvatarImageUrl

`func (o *LimitedUser) GetCurrentAvatarImageUrl() string`

GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarImageUrlOk

`func (o *LimitedUser) GetCurrentAvatarImageUrlOk() (*string, bool)`

GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarImageUrl

`func (o *LimitedUser) SetCurrentAvatarImageUrl(v string)`

SetCurrentAvatarImageUrl sets CurrentAvatarImageUrl field to given value.


### GetCurrentAvatarThumbnailImageUrl

`func (o *LimitedUser) GetCurrentAvatarThumbnailImageUrl() string`

GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarThumbnailImageUrlOk

`func (o *LimitedUser) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool)`

GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarThumbnailImageUrl

`func (o *LimitedUser) SetCurrentAvatarThumbnailImageUrl(v string)`

SetCurrentAvatarThumbnailImageUrl sets CurrentAvatarThumbnailImageUrl field to given value.


### GetDeveloperType

`func (o *LimitedUser) GetDeveloperType() DeveloperType`

GetDeveloperType returns the DeveloperType field if non-nil, zero value otherwise.

### GetDeveloperTypeOk

`func (o *LimitedUser) GetDeveloperTypeOk() (*DeveloperType, bool)`

GetDeveloperTypeOk returns a tuple with the DeveloperType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperType

`func (o *LimitedUser) SetDeveloperType(v DeveloperType)`

SetDeveloperType sets DeveloperType field to given value.


### GetDisplayName

`func (o *LimitedUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LimitedUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LimitedUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetFallbackAvatar

`func (o *LimitedUser) GetFallbackAvatar() string`

GetFallbackAvatar returns the FallbackAvatar field if non-nil, zero value otherwise.

### GetFallbackAvatarOk

`func (o *LimitedUser) GetFallbackAvatarOk() (*string, bool)`

GetFallbackAvatarOk returns a tuple with the FallbackAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackAvatar

`func (o *LimitedUser) SetFallbackAvatar(v string)`

SetFallbackAvatar sets FallbackAvatar field to given value.


### GetId

`func (o *LimitedUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LimitedUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LimitedUser) SetId(v string)`

SetId sets Id field to given value.


### GetIsFriend

`func (o *LimitedUser) GetIsFriend() bool`

GetIsFriend returns the IsFriend field if non-nil, zero value otherwise.

### GetIsFriendOk

`func (o *LimitedUser) GetIsFriendOk() (*bool, bool)`

GetIsFriendOk returns a tuple with the IsFriend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFriend

`func (o *LimitedUser) SetIsFriend(v bool)`

SetIsFriend sets IsFriend field to given value.


### GetLastPlatform

`func (o *LimitedUser) GetLastPlatform() string`

GetLastPlatform returns the LastPlatform field if non-nil, zero value otherwise.

### GetLastPlatformOk

`func (o *LimitedUser) GetLastPlatformOk() (*string, bool)`

GetLastPlatformOk returns a tuple with the LastPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlatform

`func (o *LimitedUser) SetLastPlatform(v string)`

SetLastPlatform sets LastPlatform field to given value.


### GetProfilePicOverride

`func (o *LimitedUser) GetProfilePicOverride() string`

GetProfilePicOverride returns the ProfilePicOverride field if non-nil, zero value otherwise.

### GetProfilePicOverrideOk

`func (o *LimitedUser) GetProfilePicOverrideOk() (*string, bool)`

GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicOverride

`func (o *LimitedUser) SetProfilePicOverride(v string)`

SetProfilePicOverride sets ProfilePicOverride field to given value.


### GetStatus

`func (o *LimitedUser) GetStatus() UserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LimitedUser) GetStatusOk() (*UserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LimitedUser) SetStatus(v UserStatus)`

SetStatus sets Status field to given value.


### GetStatusDescription

`func (o *LimitedUser) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *LimitedUser) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *LimitedUser) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.


### GetTags

`func (o *LimitedUser) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LimitedUser) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LimitedUser) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetUserIcon

`func (o *LimitedUser) GetUserIcon() string`

GetUserIcon returns the UserIcon field if non-nil, zero value otherwise.

### GetUserIconOk

`func (o *LimitedUser) GetUserIconOk() (*string, bool)`

GetUserIconOk returns a tuple with the UserIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIcon

`func (o *LimitedUser) SetUserIcon(v string)`

SetUserIcon sets UserIcon field to given value.


### GetUsername

`func (o *LimitedUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LimitedUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LimitedUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LimitedUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLocation

`func (o *LimitedUser) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LimitedUser) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LimitedUser) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *LimitedUser) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetFriendKey

`func (o *LimitedUser) GetFriendKey() string`

GetFriendKey returns the FriendKey field if non-nil, zero value otherwise.

### GetFriendKeyOk

`func (o *LimitedUser) GetFriendKeyOk() (*string, bool)`

GetFriendKeyOk returns a tuple with the FriendKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendKey

`func (o *LimitedUser) SetFriendKey(v string)`

SetFriendKey sets FriendKey field to given value.

### HasFriendKey

`func (o *LimitedUser) HasFriendKey() bool`

HasFriendKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


