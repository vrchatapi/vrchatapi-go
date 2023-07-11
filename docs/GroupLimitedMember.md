# GroupLimitedMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**IsRepresenting** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewGroupLimitedMember

`func NewGroupLimitedMember() *GroupLimitedMember`

NewGroupLimitedMember instantiates a new GroupLimitedMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupLimitedMemberWithDefaults

`func NewGroupLimitedMemberWithDefaults() *GroupLimitedMember`

NewGroupLimitedMemberWithDefaults instantiates a new GroupLimitedMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupLimitedMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupLimitedMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupLimitedMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupLimitedMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *GroupLimitedMember) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupLimitedMember) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupLimitedMember) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupLimitedMember) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetUserId

`func (o *GroupLimitedMember) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GroupLimitedMember) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GroupLimitedMember) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GroupLimitedMember) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetIsRepresenting

`func (o *GroupLimitedMember) GetIsRepresenting() bool`

GetIsRepresenting returns the IsRepresenting field if non-nil, zero value otherwise.

### GetIsRepresentingOk

`func (o *GroupLimitedMember) GetIsRepresentingOk() (*bool, bool)`

GetIsRepresentingOk returns a tuple with the IsRepresenting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRepresenting

`func (o *GroupLimitedMember) SetIsRepresenting(v bool)`

SetIsRepresenting sets IsRepresenting field to given value.

### HasIsRepresenting

`func (o *GroupLimitedMember) HasIsRepresenting() bool`

HasIsRepresenting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


