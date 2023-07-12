# GroupPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the permission. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the permission. | [optional] 
**Help** | Pointer to **string** | Human-readable description of the permission. | [optional] 
**IsManagementPermission** | Pointer to **bool** | Whether this permission is a \&quot;management\&quot; permission. | [optional] [default to false]
**AllowedToAdd** | Pointer to **bool** | Whether the user is allowed to add this permission to a role. | [optional] [default to false]

## Methods

### NewGroupPermission

`func NewGroupPermission() *GroupPermission`

NewGroupPermission instantiates a new GroupPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPermissionWithDefaults

`func NewGroupPermissionWithDefaults() *GroupPermission`

NewGroupPermissionWithDefaults instantiates a new GroupPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupPermission) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupPermission) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *GroupPermission) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GroupPermission) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GroupPermission) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GroupPermission) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHelp

`func (o *GroupPermission) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *GroupPermission) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *GroupPermission) SetHelp(v string)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *GroupPermission) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetIsManagementPermission

`func (o *GroupPermission) GetIsManagementPermission() bool`

GetIsManagementPermission returns the IsManagementPermission field if non-nil, zero value otherwise.

### GetIsManagementPermissionOk

`func (o *GroupPermission) GetIsManagementPermissionOk() (*bool, bool)`

GetIsManagementPermissionOk returns a tuple with the IsManagementPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManagementPermission

`func (o *GroupPermission) SetIsManagementPermission(v bool)`

SetIsManagementPermission sets IsManagementPermission field to given value.

### HasIsManagementPermission

`func (o *GroupPermission) HasIsManagementPermission() bool`

HasIsManagementPermission returns a boolean if a field has been set.

### GetAllowedToAdd

`func (o *GroupPermission) GetAllowedToAdd() bool`

GetAllowedToAdd returns the AllowedToAdd field if non-nil, zero value otherwise.

### GetAllowedToAddOk

`func (o *GroupPermission) GetAllowedToAddOk() (*bool, bool)`

GetAllowedToAddOk returns a tuple with the AllowedToAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedToAdd

`func (o *GroupPermission) SetAllowedToAdd(v bool)`

SetAllowedToAdd sets AllowedToAdd field to given value.

### HasAllowedToAdd

`func (o *GroupPermission) HasAllowedToAdd() bool`

HasAllowedToAdd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


