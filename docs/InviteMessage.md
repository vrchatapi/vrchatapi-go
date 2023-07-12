# InviteMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanBeUpdated** | **bool** |  | [default to true]
**Id** | **string** |  | 
**Message** | **string** |  | 
**MessageType** | [**InviteMessageType**](InviteMessageType.md) |  | [default to MESSAGE]
**RemainingCooldownMinutes** | **int32** | Changes to 60 when updated, although probably server-side configurable. | [default to 0]
**Slot** | **int32** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewInviteMessage

`func NewInviteMessage(canBeUpdated bool, id string, message string, messageType InviteMessageType, remainingCooldownMinutes int32, slot int32, updatedAt time.Time, ) *InviteMessage`

NewInviteMessage instantiates a new InviteMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteMessageWithDefaults

`func NewInviteMessageWithDefaults() *InviteMessage`

NewInviteMessageWithDefaults instantiates a new InviteMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanBeUpdated

`func (o *InviteMessage) GetCanBeUpdated() bool`

GetCanBeUpdated returns the CanBeUpdated field if non-nil, zero value otherwise.

### GetCanBeUpdatedOk

`func (o *InviteMessage) GetCanBeUpdatedOk() (*bool, bool)`

GetCanBeUpdatedOk returns a tuple with the CanBeUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeUpdated

`func (o *InviteMessage) SetCanBeUpdated(v bool)`

SetCanBeUpdated sets CanBeUpdated field to given value.


### GetId

`func (o *InviteMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InviteMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InviteMessage) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *InviteMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InviteMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InviteMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMessageType

`func (o *InviteMessage) GetMessageType() InviteMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *InviteMessage) GetMessageTypeOk() (*InviteMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *InviteMessage) SetMessageType(v InviteMessageType)`

SetMessageType sets MessageType field to given value.


### GetRemainingCooldownMinutes

`func (o *InviteMessage) GetRemainingCooldownMinutes() int32`

GetRemainingCooldownMinutes returns the RemainingCooldownMinutes field if non-nil, zero value otherwise.

### GetRemainingCooldownMinutesOk

`func (o *InviteMessage) GetRemainingCooldownMinutesOk() (*int32, bool)`

GetRemainingCooldownMinutesOk returns a tuple with the RemainingCooldownMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCooldownMinutes

`func (o *InviteMessage) SetRemainingCooldownMinutes(v int32)`

SetRemainingCooldownMinutes sets RemainingCooldownMinutes field to given value.


### GetSlot

`func (o *InviteMessage) GetSlot() int32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InviteMessage) GetSlotOk() (*int32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InviteMessage) SetSlot(v int32)`

SetSlot sets Slot field to given value.


### GetUpdatedAt

`func (o *InviteMessage) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InviteMessage) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InviteMessage) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


