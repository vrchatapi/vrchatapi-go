# GroupRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsSelfAssignable** | Pointer to **bool** |  | [optional] [default to false]
**Permissions** | Pointer to **[]string** |  | [optional] 
**IsManagementRole** | Pointer to **bool** |  | [optional] [default to false]
**RequiresTwoFactor** | Pointer to **bool** |  | [optional] [default to false]
**RequiresPurchase** | Pointer to **bool** |  | [optional] [default to false]
**Order** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGroupRole

`func NewGroupRole() *GroupRole`

NewGroupRole instantiates a new GroupRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRoleWithDefaults

`func NewGroupRoleWithDefaults() *GroupRole`

NewGroupRoleWithDefaults instantiates a new GroupRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *GroupRole) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupRole) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupRole) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupRole) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *GroupRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GroupRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsSelfAssignable

`func (o *GroupRole) GetIsSelfAssignable() bool`

GetIsSelfAssignable returns the IsSelfAssignable field if non-nil, zero value otherwise.

### GetIsSelfAssignableOk

`func (o *GroupRole) GetIsSelfAssignableOk() (*bool, bool)`

GetIsSelfAssignableOk returns a tuple with the IsSelfAssignable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfAssignable

`func (o *GroupRole) SetIsSelfAssignable(v bool)`

SetIsSelfAssignable sets IsSelfAssignable field to given value.

### HasIsSelfAssignable

`func (o *GroupRole) HasIsSelfAssignable() bool`

HasIsSelfAssignable returns a boolean if a field has been set.

### GetPermissions

`func (o *GroupRole) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GroupRole) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GroupRole) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GroupRole) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetIsManagementRole

`func (o *GroupRole) GetIsManagementRole() bool`

GetIsManagementRole returns the IsManagementRole field if non-nil, zero value otherwise.

### GetIsManagementRoleOk

`func (o *GroupRole) GetIsManagementRoleOk() (*bool, bool)`

GetIsManagementRoleOk returns a tuple with the IsManagementRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManagementRole

`func (o *GroupRole) SetIsManagementRole(v bool)`

SetIsManagementRole sets IsManagementRole field to given value.

### HasIsManagementRole

`func (o *GroupRole) HasIsManagementRole() bool`

HasIsManagementRole returns a boolean if a field has been set.

### GetRequiresTwoFactor

`func (o *GroupRole) GetRequiresTwoFactor() bool`

GetRequiresTwoFactor returns the RequiresTwoFactor field if non-nil, zero value otherwise.

### GetRequiresTwoFactorOk

`func (o *GroupRole) GetRequiresTwoFactorOk() (*bool, bool)`

GetRequiresTwoFactorOk returns a tuple with the RequiresTwoFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresTwoFactor

`func (o *GroupRole) SetRequiresTwoFactor(v bool)`

SetRequiresTwoFactor sets RequiresTwoFactor field to given value.

### HasRequiresTwoFactor

`func (o *GroupRole) HasRequiresTwoFactor() bool`

HasRequiresTwoFactor returns a boolean if a field has been set.

### GetRequiresPurchase

`func (o *GroupRole) GetRequiresPurchase() bool`

GetRequiresPurchase returns the RequiresPurchase field if non-nil, zero value otherwise.

### GetRequiresPurchaseOk

`func (o *GroupRole) GetRequiresPurchaseOk() (*bool, bool)`

GetRequiresPurchaseOk returns a tuple with the RequiresPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPurchase

`func (o *GroupRole) SetRequiresPurchase(v bool)`

SetRequiresPurchase sets RequiresPurchase field to given value.

### HasRequiresPurchase

`func (o *GroupRole) HasRequiresPurchase() bool`

HasRequiresPurchase returns a boolean if a field has been set.

### GetOrder

`func (o *GroupRole) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GroupRole) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GroupRole) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GroupRole) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GroupRole) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupRole) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupRole) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupRole) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GroupRole) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GroupRole) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GroupRole) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GroupRole) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


