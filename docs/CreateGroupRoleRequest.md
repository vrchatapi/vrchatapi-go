# CreateGroupRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsSelfAssignable** | Pointer to **bool** |  | [optional] [default to false]
**Permissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateGroupRoleRequest

`func NewCreateGroupRoleRequest() *CreateGroupRoleRequest`

NewCreateGroupRoleRequest instantiates a new CreateGroupRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupRoleRequestWithDefaults

`func NewCreateGroupRoleRequestWithDefaults() *CreateGroupRoleRequest`

NewCreateGroupRoleRequestWithDefaults instantiates a new CreateGroupRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateGroupRoleRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateGroupRoleRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateGroupRoleRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateGroupRoleRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateGroupRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroupRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroupRoleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateGroupRoleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateGroupRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGroupRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGroupRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGroupRoleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsSelfAssignable

`func (o *CreateGroupRoleRequest) GetIsSelfAssignable() bool`

GetIsSelfAssignable returns the IsSelfAssignable field if non-nil, zero value otherwise.

### GetIsSelfAssignableOk

`func (o *CreateGroupRoleRequest) GetIsSelfAssignableOk() (*bool, bool)`

GetIsSelfAssignableOk returns a tuple with the IsSelfAssignable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfAssignable

`func (o *CreateGroupRoleRequest) SetIsSelfAssignable(v bool)`

SetIsSelfAssignable sets IsSelfAssignable field to given value.

### HasIsSelfAssignable

`func (o *CreateGroupRoleRequest) HasIsSelfAssignable() bool`

HasIsSelfAssignable returns a boolean if a field has been set.

### GetPermissions

`func (o *CreateGroupRoleRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateGroupRoleRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateGroupRoleRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateGroupRoleRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


