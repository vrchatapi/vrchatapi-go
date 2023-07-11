# FriendStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomingRequest** | **bool** |  | [default to false]
**IsFriend** | **bool** |  | [default to false]
**OutgoingRequest** | **bool** |  | [default to false]

## Methods

### NewFriendStatus

`func NewFriendStatus(incomingRequest bool, isFriend bool, outgoingRequest bool, ) *FriendStatus`

NewFriendStatus instantiates a new FriendStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFriendStatusWithDefaults

`func NewFriendStatusWithDefaults() *FriendStatus`

NewFriendStatusWithDefaults instantiates a new FriendStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomingRequest

`func (o *FriendStatus) GetIncomingRequest() bool`

GetIncomingRequest returns the IncomingRequest field if non-nil, zero value otherwise.

### GetIncomingRequestOk

`func (o *FriendStatus) GetIncomingRequestOk() (*bool, bool)`

GetIncomingRequestOk returns a tuple with the IncomingRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingRequest

`func (o *FriendStatus) SetIncomingRequest(v bool)`

SetIncomingRequest sets IncomingRequest field to given value.


### GetIsFriend

`func (o *FriendStatus) GetIsFriend() bool`

GetIsFriend returns the IsFriend field if non-nil, zero value otherwise.

### GetIsFriendOk

`func (o *FriendStatus) GetIsFriendOk() (*bool, bool)`

GetIsFriendOk returns a tuple with the IsFriend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFriend

`func (o *FriendStatus) SetIsFriend(v bool)`

SetIsFriend sets IsFriend field to given value.


### GetOutgoingRequest

`func (o *FriendStatus) GetOutgoingRequest() bool`

GetOutgoingRequest returns the OutgoingRequest field if non-nil, zero value otherwise.

### GetOutgoingRequestOk

`func (o *FriendStatus) GetOutgoingRequestOk() (*bool, bool)`

GetOutgoingRequestOk returns a tuple with the OutgoingRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingRequest

`func (o *FriendStatus) SetOutgoingRequest(v bool)`

SetOutgoingRequest sets OutgoingRequest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


