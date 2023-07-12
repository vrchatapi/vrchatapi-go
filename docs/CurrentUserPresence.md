# CurrentUserPresence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarThumbnail** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**Instance** | Pointer to **NullableString** |  | [optional] 
**InstanceType** | Pointer to **NullableString** | either an InstanceType or an empty string | [optional] 
**IsRejoining** | Pointer to **NullableString** |  | [optional] 
**Platform** | Pointer to **NullableString** | either a Platform or an empty string | [optional] 
**ProfilePicOverride** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** | either a UserStatus or empty string | [optional] 
**TravelingToInstance** | Pointer to **NullableString** |  | [optional] 
**TravelingToWorld** | Pointer to **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | [optional] 
**World** | Pointer to **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | [optional] 

## Methods

### NewCurrentUserPresence

`func NewCurrentUserPresence() *CurrentUserPresence`

NewCurrentUserPresence instantiates a new CurrentUserPresence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentUserPresenceWithDefaults

`func NewCurrentUserPresenceWithDefaults() *CurrentUserPresence`

NewCurrentUserPresenceWithDefaults instantiates a new CurrentUserPresence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarThumbnail

`func (o *CurrentUserPresence) GetAvatarThumbnail() string`

GetAvatarThumbnail returns the AvatarThumbnail field if non-nil, zero value otherwise.

### GetAvatarThumbnailOk

`func (o *CurrentUserPresence) GetAvatarThumbnailOk() (*string, bool)`

GetAvatarThumbnailOk returns a tuple with the AvatarThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarThumbnail

`func (o *CurrentUserPresence) SetAvatarThumbnail(v string)`

SetAvatarThumbnail sets AvatarThumbnail field to given value.

### HasAvatarThumbnail

`func (o *CurrentUserPresence) HasAvatarThumbnail() bool`

HasAvatarThumbnail returns a boolean if a field has been set.

### SetAvatarThumbnailNil

`func (o *CurrentUserPresence) SetAvatarThumbnailNil(b bool)`

 SetAvatarThumbnailNil sets the value for AvatarThumbnail to be an explicit nil

### UnsetAvatarThumbnail
`func (o *CurrentUserPresence) UnsetAvatarThumbnail()`

UnsetAvatarThumbnail ensures that no value is present for AvatarThumbnail, not even an explicit nil
### GetDisplayName

`func (o *CurrentUserPresence) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CurrentUserPresence) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CurrentUserPresence) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CurrentUserPresence) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGroups

`func (o *CurrentUserPresence) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CurrentUserPresence) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CurrentUserPresence) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *CurrentUserPresence) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *CurrentUserPresence) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *CurrentUserPresence) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetId

`func (o *CurrentUserPresence) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrentUserPresence) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrentUserPresence) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CurrentUserPresence) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstance

`func (o *CurrentUserPresence) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *CurrentUserPresence) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *CurrentUserPresence) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *CurrentUserPresence) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *CurrentUserPresence) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *CurrentUserPresence) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetInstanceType

`func (o *CurrentUserPresence) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *CurrentUserPresence) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *CurrentUserPresence) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *CurrentUserPresence) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *CurrentUserPresence) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *CurrentUserPresence) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
### GetIsRejoining

`func (o *CurrentUserPresence) GetIsRejoining() string`

GetIsRejoining returns the IsRejoining field if non-nil, zero value otherwise.

### GetIsRejoiningOk

`func (o *CurrentUserPresence) GetIsRejoiningOk() (*string, bool)`

GetIsRejoiningOk returns a tuple with the IsRejoining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRejoining

`func (o *CurrentUserPresence) SetIsRejoining(v string)`

SetIsRejoining sets IsRejoining field to given value.

### HasIsRejoining

`func (o *CurrentUserPresence) HasIsRejoining() bool`

HasIsRejoining returns a boolean if a field has been set.

### SetIsRejoiningNil

`func (o *CurrentUserPresence) SetIsRejoiningNil(b bool)`

 SetIsRejoiningNil sets the value for IsRejoining to be an explicit nil

### UnsetIsRejoining
`func (o *CurrentUserPresence) UnsetIsRejoining()`

UnsetIsRejoining ensures that no value is present for IsRejoining, not even an explicit nil
### GetPlatform

`func (o *CurrentUserPresence) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CurrentUserPresence) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CurrentUserPresence) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CurrentUserPresence) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *CurrentUserPresence) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *CurrentUserPresence) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetProfilePicOverride

`func (o *CurrentUserPresence) GetProfilePicOverride() string`

GetProfilePicOverride returns the ProfilePicOverride field if non-nil, zero value otherwise.

### GetProfilePicOverrideOk

`func (o *CurrentUserPresence) GetProfilePicOverrideOk() (*string, bool)`

GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicOverride

`func (o *CurrentUserPresence) SetProfilePicOverride(v string)`

SetProfilePicOverride sets ProfilePicOverride field to given value.

### HasProfilePicOverride

`func (o *CurrentUserPresence) HasProfilePicOverride() bool`

HasProfilePicOverride returns a boolean if a field has been set.

### SetProfilePicOverrideNil

`func (o *CurrentUserPresence) SetProfilePicOverrideNil(b bool)`

 SetProfilePicOverrideNil sets the value for ProfilePicOverride to be an explicit nil

### UnsetProfilePicOverride
`func (o *CurrentUserPresence) UnsetProfilePicOverride()`

UnsetProfilePicOverride ensures that no value is present for ProfilePicOverride, not even an explicit nil
### GetStatus

`func (o *CurrentUserPresence) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CurrentUserPresence) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CurrentUserPresence) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CurrentUserPresence) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CurrentUserPresence) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CurrentUserPresence) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTravelingToInstance

`func (o *CurrentUserPresence) GetTravelingToInstance() string`

GetTravelingToInstance returns the TravelingToInstance field if non-nil, zero value otherwise.

### GetTravelingToInstanceOk

`func (o *CurrentUserPresence) GetTravelingToInstanceOk() (*string, bool)`

GetTravelingToInstanceOk returns a tuple with the TravelingToInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelingToInstance

`func (o *CurrentUserPresence) SetTravelingToInstance(v string)`

SetTravelingToInstance sets TravelingToInstance field to given value.

### HasTravelingToInstance

`func (o *CurrentUserPresence) HasTravelingToInstance() bool`

HasTravelingToInstance returns a boolean if a field has been set.

### SetTravelingToInstanceNil

`func (o *CurrentUserPresence) SetTravelingToInstanceNil(b bool)`

 SetTravelingToInstanceNil sets the value for TravelingToInstance to be an explicit nil

### UnsetTravelingToInstance
`func (o *CurrentUserPresence) UnsetTravelingToInstance()`

UnsetTravelingToInstance ensures that no value is present for TravelingToInstance, not even an explicit nil
### GetTravelingToWorld

`func (o *CurrentUserPresence) GetTravelingToWorld() string`

GetTravelingToWorld returns the TravelingToWorld field if non-nil, zero value otherwise.

### GetTravelingToWorldOk

`func (o *CurrentUserPresence) GetTravelingToWorldOk() (*string, bool)`

GetTravelingToWorldOk returns a tuple with the TravelingToWorld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelingToWorld

`func (o *CurrentUserPresence) SetTravelingToWorld(v string)`

SetTravelingToWorld sets TravelingToWorld field to given value.

### HasTravelingToWorld

`func (o *CurrentUserPresence) HasTravelingToWorld() bool`

HasTravelingToWorld returns a boolean if a field has been set.

### GetWorld

`func (o *CurrentUserPresence) GetWorld() string`

GetWorld returns the World field if non-nil, zero value otherwise.

### GetWorldOk

`func (o *CurrentUserPresence) GetWorldOk() (*string, bool)`

GetWorldOk returns a tuple with the World field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorld

`func (o *CurrentUserPresence) SetWorld(v string)`

SetWorld sets World field to given value.

### HasWorld

`func (o *CurrentUserPresence) HasWorld() bool`

HasWorld returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


