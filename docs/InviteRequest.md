# InviteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **string** | InstanceID can be \&quot;offline\&quot; on User profiles if you are not friends with that user and \&quot;private\&quot; if you are friends and user is in private instance. | 
**MessageSlot** | Pointer to **int32** |  | [optional] 

## Methods

### NewInviteRequest

`func NewInviteRequest(instanceId string, ) *InviteRequest`

NewInviteRequest instantiates a new InviteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteRequestWithDefaults

`func NewInviteRequestWithDefaults() *InviteRequest`

NewInviteRequestWithDefaults instantiates a new InviteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *InviteRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InviteRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InviteRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetMessageSlot

`func (o *InviteRequest) GetMessageSlot() int32`

GetMessageSlot returns the MessageSlot field if non-nil, zero value otherwise.

### GetMessageSlotOk

`func (o *InviteRequest) GetMessageSlotOk() (*int32, bool)`

GetMessageSlotOk returns a tuple with the MessageSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSlot

`func (o *InviteRequest) SetMessageSlot(v int32)`

SetMessageSlot sets MessageSlot field to given value.

### HasMessageSlot

`func (o *InviteRequest) HasMessageSlot() bool`

HasMessageSlot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


