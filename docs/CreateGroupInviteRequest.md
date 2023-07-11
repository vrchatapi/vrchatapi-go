# CreateGroupInviteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**ConfirmOverrideBlock** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewCreateGroupInviteRequest

`func NewCreateGroupInviteRequest(userId string, ) *CreateGroupInviteRequest`

NewCreateGroupInviteRequest instantiates a new CreateGroupInviteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupInviteRequestWithDefaults

`func NewCreateGroupInviteRequestWithDefaults() *CreateGroupInviteRequest`

NewCreateGroupInviteRequestWithDefaults instantiates a new CreateGroupInviteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CreateGroupInviteRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateGroupInviteRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateGroupInviteRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetConfirmOverrideBlock

`func (o *CreateGroupInviteRequest) GetConfirmOverrideBlock() bool`

GetConfirmOverrideBlock returns the ConfirmOverrideBlock field if non-nil, zero value otherwise.

### GetConfirmOverrideBlockOk

`func (o *CreateGroupInviteRequest) GetConfirmOverrideBlockOk() (*bool, bool)`

GetConfirmOverrideBlockOk returns a tuple with the ConfirmOverrideBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmOverrideBlock

`func (o *CreateGroupInviteRequest) SetConfirmOverrideBlock(v bool)`

SetConfirmOverrideBlock sets ConfirmOverrideBlock field to given value.

### HasConfirmOverrideBlock

`func (o *CreateGroupInviteRequest) HasConfirmOverrideBlock() bool`

HasConfirmOverrideBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


