# CurrentUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedTOSVersion** | **int32** |  | 
**AcceptedPrivacyVersion** | Pointer to **int32** |  | [optional] 
**AccountDeletionDate** | Pointer to **NullableString** |  | [optional] 
**AccountDeletionLog** | Pointer to [**[]AccountDeletionLog**](AccountDeletionLog.md) |   | [optional] 
**ActiveFriends** | Pointer to **[]string** |   | [optional] 
**AllowAvatarCopying** | **bool** |  | 
**Bio** | **string** |  | 
**BioLinks** | **[]string** |   | 
**CurrentAvatar** | **string** |  | 
**CurrentAvatarAssetUrl** | **string** |  | 
**CurrentAvatarImageUrl** | **string** | When profilePicOverride is not empty, use it instead. | 
**CurrentAvatarThumbnailImageUrl** | **string** | When profilePicOverride is not empty, use it instead. | 
**DateJoined** | **string** |  | 
**DeveloperType** | [**DeveloperType**](DeveloperType.md) |  | [default to NONE]
**DisplayName** | **string** |  | 
**EmailVerified** | **bool** |  | 
**FallbackAvatar** | Pointer to **string** |  | [optional] 
**FriendGroupNames** | **[]string** | Always empty array. | 
**FriendKey** | **string** |  | 
**Friends** | **[]string** |  | 
**HasBirthday** | **bool** |  | 
**HasEmail** | **bool** |  | 
**HasLoggedInFromClient** | **bool** |  | 
**HasPendingEmail** | **bool** |  | 
**HomeLocation** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**Id** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**IsFriend** | **bool** |  | [default to false]
**LastActivity** | Pointer to **time.Time** |  | [optional] 
**LastLogin** | **time.Time** |  | 
**LastPlatform** | **string** | This can be &#x60;standalonewindows&#x60; or &#x60;android&#x60;, but can also pretty much be any random Unity verison such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | 
**ObfuscatedEmail** | **string** |  | 
**ObfuscatedPendingEmail** | **string** |  | 
**OculusId** | **string** |  | 
**OfflineFriends** | Pointer to **[]string** |  | [optional] 
**OnlineFriends** | Pointer to **[]string** |  | [optional] 
**PastDisplayNames** | [**[]PastDisplayName**](PastDisplayName.md) |   | 
**Presence** | Pointer to [**CurrentUserPresence**](CurrentUserPresence.md) |  | [optional] 
**ProfilePicOverride** | **string** |  | 
**State** | [**UserState**](UserState.md) |  | [default to OFFLINE]
**Status** | [**UserStatus**](UserStatus.md) |  | [default to OFFLINE]
**StatusDescription** | **string** |  | 
**StatusFirstTime** | **bool** |  | 
**StatusHistory** | **[]string** |  | 
**SteamDetails** | **map[string]interface{}** |  | 
**SteamId** | **string** |  | 
**Tags** | **[]string** |  | 
**TwoFactorAuthEnabled** | **bool** |  | 
**TwoFactorAuthEnabledDate** | Pointer to **NullableTime** |  | [optional] 
**Unsubscribe** | **bool** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UserIcon** | **string** |  | 
**Username** | Pointer to **string** | -| **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429). | [optional] 

## Methods

### NewCurrentUser

`func NewCurrentUser(acceptedTOSVersion int32, allowAvatarCopying bool, bio string, bioLinks []string, currentAvatar string, currentAvatarAssetUrl string, currentAvatarImageUrl string, currentAvatarThumbnailImageUrl string, dateJoined string, developerType DeveloperType, displayName string, emailVerified bool, friendGroupNames []string, friendKey string, friends []string, hasBirthday bool, hasEmail bool, hasLoggedInFromClient bool, hasPendingEmail bool, homeLocation string, id string, isFriend bool, lastLogin time.Time, lastPlatform string, obfuscatedEmail string, obfuscatedPendingEmail string, oculusId string, pastDisplayNames []PastDisplayName, profilePicOverride string, state UserState, status UserStatus, statusDescription string, statusFirstTime bool, statusHistory []string, steamDetails map[string]interface{}, steamId string, tags []string, twoFactorAuthEnabled bool, unsubscribe bool, userIcon string, ) *CurrentUser`

NewCurrentUser instantiates a new CurrentUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentUserWithDefaults

`func NewCurrentUserWithDefaults() *CurrentUser`

NewCurrentUserWithDefaults instantiates a new CurrentUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedTOSVersion

`func (o *CurrentUser) GetAcceptedTOSVersion() int32`

GetAcceptedTOSVersion returns the AcceptedTOSVersion field if non-nil, zero value otherwise.

### GetAcceptedTOSVersionOk

`func (o *CurrentUser) GetAcceptedTOSVersionOk() (*int32, bool)`

GetAcceptedTOSVersionOk returns a tuple with the AcceptedTOSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTOSVersion

`func (o *CurrentUser) SetAcceptedTOSVersion(v int32)`

SetAcceptedTOSVersion sets AcceptedTOSVersion field to given value.


### GetAcceptedPrivacyVersion

`func (o *CurrentUser) GetAcceptedPrivacyVersion() int32`

GetAcceptedPrivacyVersion returns the AcceptedPrivacyVersion field if non-nil, zero value otherwise.

### GetAcceptedPrivacyVersionOk

`func (o *CurrentUser) GetAcceptedPrivacyVersionOk() (*int32, bool)`

GetAcceptedPrivacyVersionOk returns a tuple with the AcceptedPrivacyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedPrivacyVersion

`func (o *CurrentUser) SetAcceptedPrivacyVersion(v int32)`

SetAcceptedPrivacyVersion sets AcceptedPrivacyVersion field to given value.

### HasAcceptedPrivacyVersion

`func (o *CurrentUser) HasAcceptedPrivacyVersion() bool`

HasAcceptedPrivacyVersion returns a boolean if a field has been set.

### GetAccountDeletionDate

`func (o *CurrentUser) GetAccountDeletionDate() string`

GetAccountDeletionDate returns the AccountDeletionDate field if non-nil, zero value otherwise.

### GetAccountDeletionDateOk

`func (o *CurrentUser) GetAccountDeletionDateOk() (*string, bool)`

GetAccountDeletionDateOk returns a tuple with the AccountDeletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDeletionDate

`func (o *CurrentUser) SetAccountDeletionDate(v string)`

SetAccountDeletionDate sets AccountDeletionDate field to given value.

### HasAccountDeletionDate

`func (o *CurrentUser) HasAccountDeletionDate() bool`

HasAccountDeletionDate returns a boolean if a field has been set.

### SetAccountDeletionDateNil

`func (o *CurrentUser) SetAccountDeletionDateNil(b bool)`

 SetAccountDeletionDateNil sets the value for AccountDeletionDate to be an explicit nil

### UnsetAccountDeletionDate
`func (o *CurrentUser) UnsetAccountDeletionDate()`

UnsetAccountDeletionDate ensures that no value is present for AccountDeletionDate, not even an explicit nil
### GetAccountDeletionLog

`func (o *CurrentUser) GetAccountDeletionLog() []AccountDeletionLog`

GetAccountDeletionLog returns the AccountDeletionLog field if non-nil, zero value otherwise.

### GetAccountDeletionLogOk

`func (o *CurrentUser) GetAccountDeletionLogOk() (*[]AccountDeletionLog, bool)`

GetAccountDeletionLogOk returns a tuple with the AccountDeletionLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDeletionLog

`func (o *CurrentUser) SetAccountDeletionLog(v []AccountDeletionLog)`

SetAccountDeletionLog sets AccountDeletionLog field to given value.

### HasAccountDeletionLog

`func (o *CurrentUser) HasAccountDeletionLog() bool`

HasAccountDeletionLog returns a boolean if a field has been set.

### SetAccountDeletionLogNil

`func (o *CurrentUser) SetAccountDeletionLogNil(b bool)`

 SetAccountDeletionLogNil sets the value for AccountDeletionLog to be an explicit nil

### UnsetAccountDeletionLog
`func (o *CurrentUser) UnsetAccountDeletionLog()`

UnsetAccountDeletionLog ensures that no value is present for AccountDeletionLog, not even an explicit nil
### GetActiveFriends

`func (o *CurrentUser) GetActiveFriends() []string`

GetActiveFriends returns the ActiveFriends field if non-nil, zero value otherwise.

### GetActiveFriendsOk

`func (o *CurrentUser) GetActiveFriendsOk() (*[]string, bool)`

GetActiveFriendsOk returns a tuple with the ActiveFriends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFriends

`func (o *CurrentUser) SetActiveFriends(v []string)`

SetActiveFriends sets ActiveFriends field to given value.

### HasActiveFriends

`func (o *CurrentUser) HasActiveFriends() bool`

HasActiveFriends returns a boolean if a field has been set.

### GetAllowAvatarCopying

`func (o *CurrentUser) GetAllowAvatarCopying() bool`

GetAllowAvatarCopying returns the AllowAvatarCopying field if non-nil, zero value otherwise.

### GetAllowAvatarCopyingOk

`func (o *CurrentUser) GetAllowAvatarCopyingOk() (*bool, bool)`

GetAllowAvatarCopyingOk returns a tuple with the AllowAvatarCopying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAvatarCopying

`func (o *CurrentUser) SetAllowAvatarCopying(v bool)`

SetAllowAvatarCopying sets AllowAvatarCopying field to given value.


### GetBio

`func (o *CurrentUser) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *CurrentUser) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *CurrentUser) SetBio(v string)`

SetBio sets Bio field to given value.


### GetBioLinks

`func (o *CurrentUser) GetBioLinks() []string`

GetBioLinks returns the BioLinks field if non-nil, zero value otherwise.

### GetBioLinksOk

`func (o *CurrentUser) GetBioLinksOk() (*[]string, bool)`

GetBioLinksOk returns a tuple with the BioLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioLinks

`func (o *CurrentUser) SetBioLinks(v []string)`

SetBioLinks sets BioLinks field to given value.


### GetCurrentAvatar

`func (o *CurrentUser) GetCurrentAvatar() string`

GetCurrentAvatar returns the CurrentAvatar field if non-nil, zero value otherwise.

### GetCurrentAvatarOk

`func (o *CurrentUser) GetCurrentAvatarOk() (*string, bool)`

GetCurrentAvatarOk returns a tuple with the CurrentAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatar

`func (o *CurrentUser) SetCurrentAvatar(v string)`

SetCurrentAvatar sets CurrentAvatar field to given value.


### GetCurrentAvatarAssetUrl

`func (o *CurrentUser) GetCurrentAvatarAssetUrl() string`

GetCurrentAvatarAssetUrl returns the CurrentAvatarAssetUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarAssetUrlOk

`func (o *CurrentUser) GetCurrentAvatarAssetUrlOk() (*string, bool)`

GetCurrentAvatarAssetUrlOk returns a tuple with the CurrentAvatarAssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarAssetUrl

`func (o *CurrentUser) SetCurrentAvatarAssetUrl(v string)`

SetCurrentAvatarAssetUrl sets CurrentAvatarAssetUrl field to given value.


### GetCurrentAvatarImageUrl

`func (o *CurrentUser) GetCurrentAvatarImageUrl() string`

GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarImageUrlOk

`func (o *CurrentUser) GetCurrentAvatarImageUrlOk() (*string, bool)`

GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarImageUrl

`func (o *CurrentUser) SetCurrentAvatarImageUrl(v string)`

SetCurrentAvatarImageUrl sets CurrentAvatarImageUrl field to given value.


### GetCurrentAvatarThumbnailImageUrl

`func (o *CurrentUser) GetCurrentAvatarThumbnailImageUrl() string`

GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarThumbnailImageUrlOk

`func (o *CurrentUser) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool)`

GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarThumbnailImageUrl

`func (o *CurrentUser) SetCurrentAvatarThumbnailImageUrl(v string)`

SetCurrentAvatarThumbnailImageUrl sets CurrentAvatarThumbnailImageUrl field to given value.


### GetDateJoined

`func (o *CurrentUser) GetDateJoined() string`

GetDateJoined returns the DateJoined field if non-nil, zero value otherwise.

### GetDateJoinedOk

`func (o *CurrentUser) GetDateJoinedOk() (*string, bool)`

GetDateJoinedOk returns a tuple with the DateJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateJoined

`func (o *CurrentUser) SetDateJoined(v string)`

SetDateJoined sets DateJoined field to given value.


### GetDeveloperType

`func (o *CurrentUser) GetDeveloperType() DeveloperType`

GetDeveloperType returns the DeveloperType field if non-nil, zero value otherwise.

### GetDeveloperTypeOk

`func (o *CurrentUser) GetDeveloperTypeOk() (*DeveloperType, bool)`

GetDeveloperTypeOk returns a tuple with the DeveloperType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperType

`func (o *CurrentUser) SetDeveloperType(v DeveloperType)`

SetDeveloperType sets DeveloperType field to given value.


### GetDisplayName

`func (o *CurrentUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CurrentUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CurrentUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEmailVerified

`func (o *CurrentUser) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *CurrentUser) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *CurrentUser) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.


### GetFallbackAvatar

`func (o *CurrentUser) GetFallbackAvatar() string`

GetFallbackAvatar returns the FallbackAvatar field if non-nil, zero value otherwise.

### GetFallbackAvatarOk

`func (o *CurrentUser) GetFallbackAvatarOk() (*string, bool)`

GetFallbackAvatarOk returns a tuple with the FallbackAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackAvatar

`func (o *CurrentUser) SetFallbackAvatar(v string)`

SetFallbackAvatar sets FallbackAvatar field to given value.

### HasFallbackAvatar

`func (o *CurrentUser) HasFallbackAvatar() bool`

HasFallbackAvatar returns a boolean if a field has been set.

### GetFriendGroupNames

`func (o *CurrentUser) GetFriendGroupNames() []string`

GetFriendGroupNames returns the FriendGroupNames field if non-nil, zero value otherwise.

### GetFriendGroupNamesOk

`func (o *CurrentUser) GetFriendGroupNamesOk() (*[]string, bool)`

GetFriendGroupNamesOk returns a tuple with the FriendGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendGroupNames

`func (o *CurrentUser) SetFriendGroupNames(v []string)`

SetFriendGroupNames sets FriendGroupNames field to given value.


### GetFriendKey

`func (o *CurrentUser) GetFriendKey() string`

GetFriendKey returns the FriendKey field if non-nil, zero value otherwise.

### GetFriendKeyOk

`func (o *CurrentUser) GetFriendKeyOk() (*string, bool)`

GetFriendKeyOk returns a tuple with the FriendKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendKey

`func (o *CurrentUser) SetFriendKey(v string)`

SetFriendKey sets FriendKey field to given value.


### GetFriends

`func (o *CurrentUser) GetFriends() []string`

GetFriends returns the Friends field if non-nil, zero value otherwise.

### GetFriendsOk

`func (o *CurrentUser) GetFriendsOk() (*[]string, bool)`

GetFriendsOk returns a tuple with the Friends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriends

`func (o *CurrentUser) SetFriends(v []string)`

SetFriends sets Friends field to given value.


### GetHasBirthday

`func (o *CurrentUser) GetHasBirthday() bool`

GetHasBirthday returns the HasBirthday field if non-nil, zero value otherwise.

### GetHasBirthdayOk

`func (o *CurrentUser) GetHasBirthdayOk() (*bool, bool)`

GetHasBirthdayOk returns a tuple with the HasBirthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBirthday

`func (o *CurrentUser) SetHasBirthday(v bool)`

SetHasBirthday sets HasBirthday field to given value.


### GetHasEmail

`func (o *CurrentUser) GetHasEmail() bool`

GetHasEmail returns the HasEmail field if non-nil, zero value otherwise.

### GetHasEmailOk

`func (o *CurrentUser) GetHasEmailOk() (*bool, bool)`

GetHasEmailOk returns a tuple with the HasEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEmail

`func (o *CurrentUser) SetHasEmail(v bool)`

SetHasEmail sets HasEmail field to given value.


### GetHasLoggedInFromClient

`func (o *CurrentUser) GetHasLoggedInFromClient() bool`

GetHasLoggedInFromClient returns the HasLoggedInFromClient field if non-nil, zero value otherwise.

### GetHasLoggedInFromClientOk

`func (o *CurrentUser) GetHasLoggedInFromClientOk() (*bool, bool)`

GetHasLoggedInFromClientOk returns a tuple with the HasLoggedInFromClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLoggedInFromClient

`func (o *CurrentUser) SetHasLoggedInFromClient(v bool)`

SetHasLoggedInFromClient sets HasLoggedInFromClient field to given value.


### GetHasPendingEmail

`func (o *CurrentUser) GetHasPendingEmail() bool`

GetHasPendingEmail returns the HasPendingEmail field if non-nil, zero value otherwise.

### GetHasPendingEmailOk

`func (o *CurrentUser) GetHasPendingEmailOk() (*bool, bool)`

GetHasPendingEmailOk returns a tuple with the HasPendingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPendingEmail

`func (o *CurrentUser) SetHasPendingEmail(v bool)`

SetHasPendingEmail sets HasPendingEmail field to given value.


### GetHomeLocation

`func (o *CurrentUser) GetHomeLocation() string`

GetHomeLocation returns the HomeLocation field if non-nil, zero value otherwise.

### GetHomeLocationOk

`func (o *CurrentUser) GetHomeLocationOk() (*string, bool)`

GetHomeLocationOk returns a tuple with the HomeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeLocation

`func (o *CurrentUser) SetHomeLocation(v string)`

SetHomeLocation sets HomeLocation field to given value.


### GetId

`func (o *CurrentUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrentUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrentUser) SetId(v string)`

SetId sets Id field to given value.


### GetIsFriend

`func (o *CurrentUser) GetIsFriend() bool`

GetIsFriend returns the IsFriend field if non-nil, zero value otherwise.

### GetIsFriendOk

`func (o *CurrentUser) GetIsFriendOk() (*bool, bool)`

GetIsFriendOk returns a tuple with the IsFriend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFriend

`func (o *CurrentUser) SetIsFriend(v bool)`

SetIsFriend sets IsFriend field to given value.


### GetLastActivity

`func (o *CurrentUser) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *CurrentUser) GetLastActivityOk() (*time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *CurrentUser) SetLastActivity(v time.Time)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *CurrentUser) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### GetLastLogin

`func (o *CurrentUser) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *CurrentUser) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *CurrentUser) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.


### GetLastPlatform

`func (o *CurrentUser) GetLastPlatform() string`

GetLastPlatform returns the LastPlatform field if non-nil, zero value otherwise.

### GetLastPlatformOk

`func (o *CurrentUser) GetLastPlatformOk() (*string, bool)`

GetLastPlatformOk returns a tuple with the LastPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlatform

`func (o *CurrentUser) SetLastPlatform(v string)`

SetLastPlatform sets LastPlatform field to given value.


### GetObfuscatedEmail

`func (o *CurrentUser) GetObfuscatedEmail() string`

GetObfuscatedEmail returns the ObfuscatedEmail field if non-nil, zero value otherwise.

### GetObfuscatedEmailOk

`func (o *CurrentUser) GetObfuscatedEmailOk() (*string, bool)`

GetObfuscatedEmailOk returns a tuple with the ObfuscatedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObfuscatedEmail

`func (o *CurrentUser) SetObfuscatedEmail(v string)`

SetObfuscatedEmail sets ObfuscatedEmail field to given value.


### GetObfuscatedPendingEmail

`func (o *CurrentUser) GetObfuscatedPendingEmail() string`

GetObfuscatedPendingEmail returns the ObfuscatedPendingEmail field if non-nil, zero value otherwise.

### GetObfuscatedPendingEmailOk

`func (o *CurrentUser) GetObfuscatedPendingEmailOk() (*string, bool)`

GetObfuscatedPendingEmailOk returns a tuple with the ObfuscatedPendingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObfuscatedPendingEmail

`func (o *CurrentUser) SetObfuscatedPendingEmail(v string)`

SetObfuscatedPendingEmail sets ObfuscatedPendingEmail field to given value.


### GetOculusId

`func (o *CurrentUser) GetOculusId() string`

GetOculusId returns the OculusId field if non-nil, zero value otherwise.

### GetOculusIdOk

`func (o *CurrentUser) GetOculusIdOk() (*string, bool)`

GetOculusIdOk returns a tuple with the OculusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOculusId

`func (o *CurrentUser) SetOculusId(v string)`

SetOculusId sets OculusId field to given value.


### GetOfflineFriends

`func (o *CurrentUser) GetOfflineFriends() []string`

GetOfflineFriends returns the OfflineFriends field if non-nil, zero value otherwise.

### GetOfflineFriendsOk

`func (o *CurrentUser) GetOfflineFriendsOk() (*[]string, bool)`

GetOfflineFriendsOk returns a tuple with the OfflineFriends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineFriends

`func (o *CurrentUser) SetOfflineFriends(v []string)`

SetOfflineFriends sets OfflineFriends field to given value.

### HasOfflineFriends

`func (o *CurrentUser) HasOfflineFriends() bool`

HasOfflineFriends returns a boolean if a field has been set.

### GetOnlineFriends

`func (o *CurrentUser) GetOnlineFriends() []string`

GetOnlineFriends returns the OnlineFriends field if non-nil, zero value otherwise.

### GetOnlineFriendsOk

`func (o *CurrentUser) GetOnlineFriendsOk() (*[]string, bool)`

GetOnlineFriendsOk returns a tuple with the OnlineFriends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineFriends

`func (o *CurrentUser) SetOnlineFriends(v []string)`

SetOnlineFriends sets OnlineFriends field to given value.

### HasOnlineFriends

`func (o *CurrentUser) HasOnlineFriends() bool`

HasOnlineFriends returns a boolean if a field has been set.

### GetPastDisplayNames

`func (o *CurrentUser) GetPastDisplayNames() []PastDisplayName`

GetPastDisplayNames returns the PastDisplayNames field if non-nil, zero value otherwise.

### GetPastDisplayNamesOk

`func (o *CurrentUser) GetPastDisplayNamesOk() (*[]PastDisplayName, bool)`

GetPastDisplayNamesOk returns a tuple with the PastDisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastDisplayNames

`func (o *CurrentUser) SetPastDisplayNames(v []PastDisplayName)`

SetPastDisplayNames sets PastDisplayNames field to given value.


### GetPresence

`func (o *CurrentUser) GetPresence() CurrentUserPresence`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *CurrentUser) GetPresenceOk() (*CurrentUserPresence, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *CurrentUser) SetPresence(v CurrentUserPresence)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *CurrentUser) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetProfilePicOverride

`func (o *CurrentUser) GetProfilePicOverride() string`

GetProfilePicOverride returns the ProfilePicOverride field if non-nil, zero value otherwise.

### GetProfilePicOverrideOk

`func (o *CurrentUser) GetProfilePicOverrideOk() (*string, bool)`

GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicOverride

`func (o *CurrentUser) SetProfilePicOverride(v string)`

SetProfilePicOverride sets ProfilePicOverride field to given value.


### GetState

`func (o *CurrentUser) GetState() UserState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CurrentUser) GetStateOk() (*UserState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CurrentUser) SetState(v UserState)`

SetState sets State field to given value.


### GetStatus

`func (o *CurrentUser) GetStatus() UserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CurrentUser) GetStatusOk() (*UserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CurrentUser) SetStatus(v UserStatus)`

SetStatus sets Status field to given value.


### GetStatusDescription

`func (o *CurrentUser) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *CurrentUser) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *CurrentUser) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.


### GetStatusFirstTime

`func (o *CurrentUser) GetStatusFirstTime() bool`

GetStatusFirstTime returns the StatusFirstTime field if non-nil, zero value otherwise.

### GetStatusFirstTimeOk

`func (o *CurrentUser) GetStatusFirstTimeOk() (*bool, bool)`

GetStatusFirstTimeOk returns a tuple with the StatusFirstTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusFirstTime

`func (o *CurrentUser) SetStatusFirstTime(v bool)`

SetStatusFirstTime sets StatusFirstTime field to given value.


### GetStatusHistory

`func (o *CurrentUser) GetStatusHistory() []string`

GetStatusHistory returns the StatusHistory field if non-nil, zero value otherwise.

### GetStatusHistoryOk

`func (o *CurrentUser) GetStatusHistoryOk() (*[]string, bool)`

GetStatusHistoryOk returns a tuple with the StatusHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHistory

`func (o *CurrentUser) SetStatusHistory(v []string)`

SetStatusHistory sets StatusHistory field to given value.


### GetSteamDetails

`func (o *CurrentUser) GetSteamDetails() map[string]interface{}`

GetSteamDetails returns the SteamDetails field if non-nil, zero value otherwise.

### GetSteamDetailsOk

`func (o *CurrentUser) GetSteamDetailsOk() (*map[string]interface{}, bool)`

GetSteamDetailsOk returns a tuple with the SteamDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamDetails

`func (o *CurrentUser) SetSteamDetails(v map[string]interface{})`

SetSteamDetails sets SteamDetails field to given value.


### GetSteamId

`func (o *CurrentUser) GetSteamId() string`

GetSteamId returns the SteamId field if non-nil, zero value otherwise.

### GetSteamIdOk

`func (o *CurrentUser) GetSteamIdOk() (*string, bool)`

GetSteamIdOk returns a tuple with the SteamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamId

`func (o *CurrentUser) SetSteamId(v string)`

SetSteamId sets SteamId field to given value.


### GetTags

`func (o *CurrentUser) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CurrentUser) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CurrentUser) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTwoFactorAuthEnabled

`func (o *CurrentUser) GetTwoFactorAuthEnabled() bool`

GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledOk

`func (o *CurrentUser) GetTwoFactorAuthEnabledOk() (*bool, bool)`

GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabled

`func (o *CurrentUser) SetTwoFactorAuthEnabled(v bool)`

SetTwoFactorAuthEnabled sets TwoFactorAuthEnabled field to given value.


### GetTwoFactorAuthEnabledDate

`func (o *CurrentUser) GetTwoFactorAuthEnabledDate() time.Time`

GetTwoFactorAuthEnabledDate returns the TwoFactorAuthEnabledDate field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledDateOk

`func (o *CurrentUser) GetTwoFactorAuthEnabledDateOk() (*time.Time, bool)`

GetTwoFactorAuthEnabledDateOk returns a tuple with the TwoFactorAuthEnabledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabledDate

`func (o *CurrentUser) SetTwoFactorAuthEnabledDate(v time.Time)`

SetTwoFactorAuthEnabledDate sets TwoFactorAuthEnabledDate field to given value.

### HasTwoFactorAuthEnabledDate

`func (o *CurrentUser) HasTwoFactorAuthEnabledDate() bool`

HasTwoFactorAuthEnabledDate returns a boolean if a field has been set.

### SetTwoFactorAuthEnabledDateNil

`func (o *CurrentUser) SetTwoFactorAuthEnabledDateNil(b bool)`

 SetTwoFactorAuthEnabledDateNil sets the value for TwoFactorAuthEnabledDate to be an explicit nil

### UnsetTwoFactorAuthEnabledDate
`func (o *CurrentUser) UnsetTwoFactorAuthEnabledDate()`

UnsetTwoFactorAuthEnabledDate ensures that no value is present for TwoFactorAuthEnabledDate, not even an explicit nil
### GetUnsubscribe

`func (o *CurrentUser) GetUnsubscribe() bool`

GetUnsubscribe returns the Unsubscribe field if non-nil, zero value otherwise.

### GetUnsubscribeOk

`func (o *CurrentUser) GetUnsubscribeOk() (*bool, bool)`

GetUnsubscribeOk returns a tuple with the Unsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribe

`func (o *CurrentUser) SetUnsubscribe(v bool)`

SetUnsubscribe sets Unsubscribe field to given value.


### GetUpdatedAt

`func (o *CurrentUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CurrentUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CurrentUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CurrentUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserIcon

`func (o *CurrentUser) GetUserIcon() string`

GetUserIcon returns the UserIcon field if non-nil, zero value otherwise.

### GetUserIconOk

`func (o *CurrentUser) GetUserIconOk() (*string, bool)`

GetUserIconOk returns a tuple with the UserIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIcon

`func (o *CurrentUser) SetUserIcon(v string)`

SetUserIcon sets UserIcon field to given value.


### GetUsername

`func (o *CurrentUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CurrentUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CurrentUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CurrentUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


