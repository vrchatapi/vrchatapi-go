# UpdateGroupRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsSelfAssignable** | Pointer to **bool** |  | [optional] [default to false]
**Permissions** | Pointer to **[]string** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateGroupRoleRequest

`func NewUpdateGroupRoleRequest() *UpdateGroupRoleRequest`

NewUpdateGroupRoleRequest instantiates a new UpdateGroupRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupRoleRequestWithDefaults

`func NewUpdateGroupRoleRequestWithDefaults() *UpdateGroupRoleRequest`

NewUpdateGroupRoleRequestWithDefaults instantiates a new UpdateGroupRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateGroupRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGroupRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGroupRoleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGroupRoleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateGroupRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGroupRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGroupRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGroupRoleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsSelfAssignable

`func (o *UpdateGroupRoleRequest) GetIsSelfAssignable() bool`

GetIsSelfAssignable returns the IsSelfAssignable field if non-nil, zero value otherwise.

### GetIsSelfAssignableOk

`func (o *UpdateGroupRoleRequest) GetIsSelfAssignableOk() (*bool, bool)`

GetIsSelfAssignableOk returns a tuple with the IsSelfAssignable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfAssignable

`func (o *UpdateGroupRoleRequest) SetIsSelfAssignable(v bool)`

SetIsSelfAssignable sets IsSelfAssignable field to given value.

### HasIsSelfAssignable

`func (o *UpdateGroupRoleRequest) HasIsSelfAssignable() bool`

HasIsSelfAssignable returns a boolean if a field has been set.

### GetPermissions

`func (o *UpdateGroupRoleRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateGroupRoleRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateGroupRoleRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UpdateGroupRoleRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetOrder

`func (o *UpdateGroupRoleRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *UpdateGroupRoleRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *UpdateGroupRoleRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *UpdateGroupRoleRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


