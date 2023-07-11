# PlayerModeration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**SourceDisplayName** | **string** |  | 
**SourceUserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**TargetDisplayName** | **string** |  | 
**TargetUserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**Type** | [**PlayerModerationType**](PlayerModerationType.md) |  | [default to UNMUTE]

## Methods

### NewPlayerModeration

`func NewPlayerModeration(created time.Time, id string, sourceDisplayName string, sourceUserId string, targetDisplayName string, targetUserId string, type_ PlayerModerationType, ) *PlayerModeration`

NewPlayerModeration instantiates a new PlayerModeration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerModerationWithDefaults

`func NewPlayerModerationWithDefaults() *PlayerModeration`

NewPlayerModerationWithDefaults instantiates a new PlayerModeration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PlayerModeration) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PlayerModeration) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PlayerModeration) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *PlayerModeration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlayerModeration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlayerModeration) SetId(v string)`

SetId sets Id field to given value.


### GetSourceDisplayName

`func (o *PlayerModeration) GetSourceDisplayName() string`

GetSourceDisplayName returns the SourceDisplayName field if non-nil, zero value otherwise.

### GetSourceDisplayNameOk

`func (o *PlayerModeration) GetSourceDisplayNameOk() (*string, bool)`

GetSourceDisplayNameOk returns a tuple with the SourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDisplayName

`func (o *PlayerModeration) SetSourceDisplayName(v string)`

SetSourceDisplayName sets SourceDisplayName field to given value.


### GetSourceUserId

`func (o *PlayerModeration) GetSourceUserId() string`

GetSourceUserId returns the SourceUserId field if non-nil, zero value otherwise.

### GetSourceUserIdOk

`func (o *PlayerModeration) GetSourceUserIdOk() (*string, bool)`

GetSourceUserIdOk returns a tuple with the SourceUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUserId

`func (o *PlayerModeration) SetSourceUserId(v string)`

SetSourceUserId sets SourceUserId field to given value.


### GetTargetDisplayName

`func (o *PlayerModeration) GetTargetDisplayName() string`

GetTargetDisplayName returns the TargetDisplayName field if non-nil, zero value otherwise.

### GetTargetDisplayNameOk

`func (o *PlayerModeration) GetTargetDisplayNameOk() (*string, bool)`

GetTargetDisplayNameOk returns a tuple with the TargetDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDisplayName

`func (o *PlayerModeration) SetTargetDisplayName(v string)`

SetTargetDisplayName sets TargetDisplayName field to given value.


### GetTargetUserId

`func (o *PlayerModeration) GetTargetUserId() string`

GetTargetUserId returns the TargetUserId field if non-nil, zero value otherwise.

### GetTargetUserIdOk

`func (o *PlayerModeration) GetTargetUserIdOk() (*string, bool)`

GetTargetUserIdOk returns a tuple with the TargetUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUserId

`func (o *PlayerModeration) SetTargetUserId(v string)`

SetTargetUserId sets TargetUserId field to given value.


### GetType

`func (o *PlayerModeration) GetType() PlayerModerationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlayerModeration) GetTypeOk() (*PlayerModerationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlayerModeration) SetType(v PlayerModerationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


