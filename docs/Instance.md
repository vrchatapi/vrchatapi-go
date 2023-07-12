# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | [default to true]
**CanRequestInvite** | **bool** |  | [default to true]
**Capacity** | **int32** |  | 
**ClientNumber** | **string** | Always returns \&quot;unknown\&quot;. | 
**Full** | **bool** |  | [default to false]
**Id** | **string** | InstanceID can be \&quot;offline\&quot; on User profiles if you are not friends with that user and \&quot;private\&quot; if you are friends and user is in private instance. | 
**InstanceId** | **string** |  | 
**Location** | **string** | InstanceID can be \&quot;offline\&quot; on User profiles if you are not friends with that user and \&quot;private\&quot; if you are friends and user is in private instance. | 
**NUsers** | **int32** |  | 
**Name** | **string** |  | 
**OwnerId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**Permanent** | **bool** |  | [default to false]
**PhotonRegion** | [**Region**](Region.md) |  | [default to US]
**Platforms** | [**InstancePlatforms**](InstancePlatforms.md) |  | 
**Region** | [**Region**](Region.md) |  | [default to US]
**SecureName** | **string** |  | 
**ShortName** | Pointer to **string** |  | [optional] 
**Tags** | **[]string** | The tags array on Instances usually contain the language tags of the people in the instance.  | 
**Type** | [**InstanceType**](InstanceType.md) |  | 
**WorldId** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**Hidden** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**Friends** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**Private** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 

## Methods

### NewInstance

`func NewInstance(active bool, canRequestInvite bool, capacity int32, clientNumber string, full bool, id string, instanceId string, location string, nUsers int32, name string, permanent bool, photonRegion Region, platforms InstancePlatforms, region Region, secureName string, tags []string, type_ InstanceType, worldId string, ) *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Instance) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Instance) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Instance) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCanRequestInvite

`func (o *Instance) GetCanRequestInvite() bool`

GetCanRequestInvite returns the CanRequestInvite field if non-nil, zero value otherwise.

### GetCanRequestInviteOk

`func (o *Instance) GetCanRequestInviteOk() (*bool, bool)`

GetCanRequestInviteOk returns a tuple with the CanRequestInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRequestInvite

`func (o *Instance) SetCanRequestInvite(v bool)`

SetCanRequestInvite sets CanRequestInvite field to given value.


### GetCapacity

`func (o *Instance) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *Instance) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *Instance) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### GetClientNumber

`func (o *Instance) GetClientNumber() string`

GetClientNumber returns the ClientNumber field if non-nil, zero value otherwise.

### GetClientNumberOk

`func (o *Instance) GetClientNumberOk() (*string, bool)`

GetClientNumberOk returns a tuple with the ClientNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNumber

`func (o *Instance) SetClientNumber(v string)`

SetClientNumber sets ClientNumber field to given value.


### GetFull

`func (o *Instance) GetFull() bool`

GetFull returns the Full field if non-nil, zero value otherwise.

### GetFullOk

`func (o *Instance) GetFullOk() (*bool, bool)`

GetFullOk returns a tuple with the Full field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFull

`func (o *Instance) SetFull(v bool)`

SetFull sets Full field to given value.


### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *Instance) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *Instance) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *Instance) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetLocation

`func (o *Instance) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Instance) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Instance) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetNUsers

`func (o *Instance) GetNUsers() int32`

GetNUsers returns the NUsers field if non-nil, zero value otherwise.

### GetNUsersOk

`func (o *Instance) GetNUsersOk() (*int32, bool)`

GetNUsersOk returns a tuple with the NUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNUsers

`func (o *Instance) SetNUsers(v int32)`

SetNUsers sets NUsers field to given value.


### GetName

`func (o *Instance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Instance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Instance) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *Instance) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Instance) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Instance) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Instance) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPermanent

`func (o *Instance) GetPermanent() bool`

GetPermanent returns the Permanent field if non-nil, zero value otherwise.

### GetPermanentOk

`func (o *Instance) GetPermanentOk() (*bool, bool)`

GetPermanentOk returns a tuple with the Permanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanent

`func (o *Instance) SetPermanent(v bool)`

SetPermanent sets Permanent field to given value.


### GetPhotonRegion

`func (o *Instance) GetPhotonRegion() Region`

GetPhotonRegion returns the PhotonRegion field if non-nil, zero value otherwise.

### GetPhotonRegionOk

`func (o *Instance) GetPhotonRegionOk() (*Region, bool)`

GetPhotonRegionOk returns a tuple with the PhotonRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotonRegion

`func (o *Instance) SetPhotonRegion(v Region)`

SetPhotonRegion sets PhotonRegion field to given value.


### GetPlatforms

`func (o *Instance) GetPlatforms() InstancePlatforms`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *Instance) GetPlatformsOk() (*InstancePlatforms, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *Instance) SetPlatforms(v InstancePlatforms)`

SetPlatforms sets Platforms field to given value.


### GetRegion

`func (o *Instance) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Instance) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Instance) SetRegion(v Region)`

SetRegion sets Region field to given value.


### GetSecureName

`func (o *Instance) GetSecureName() string`

GetSecureName returns the SecureName field if non-nil, zero value otherwise.

### GetSecureNameOk

`func (o *Instance) GetSecureNameOk() (*string, bool)`

GetSecureNameOk returns a tuple with the SecureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureName

`func (o *Instance) SetSecureName(v string)`

SetSecureName sets SecureName field to given value.


### GetShortName

`func (o *Instance) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *Instance) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *Instance) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *Instance) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetTags

`func (o *Instance) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Instance) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Instance) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetType

`func (o *Instance) GetType() InstanceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Instance) GetTypeOk() (*InstanceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Instance) SetType(v InstanceType)`

SetType sets Type field to given value.


### GetWorldId

`func (o *Instance) GetWorldId() string`

GetWorldId returns the WorldId field if non-nil, zero value otherwise.

### GetWorldIdOk

`func (o *Instance) GetWorldIdOk() (*string, bool)`

GetWorldIdOk returns a tuple with the WorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldId

`func (o *Instance) SetWorldId(v string)`

SetWorldId sets WorldId field to given value.


### GetHidden

`func (o *Instance) GetHidden() string`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Instance) GetHiddenOk() (*string, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Instance) SetHidden(v string)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Instance) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetFriends

`func (o *Instance) GetFriends() string`

GetFriends returns the Friends field if non-nil, zero value otherwise.

### GetFriendsOk

`func (o *Instance) GetFriendsOk() (*string, bool)`

GetFriendsOk returns a tuple with the Friends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriends

`func (o *Instance) SetFriends(v string)`

SetFriends sets Friends field to given value.

### HasFriends

`func (o *Instance) HasFriends() bool`

HasFriends returns a boolean if a field has been set.

### GetPrivate

`func (o *Instance) GetPrivate() string`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *Instance) GetPrivateOk() (*string, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *Instance) SetPrivate(v string)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *Instance) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


