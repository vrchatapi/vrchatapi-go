# UpdateGroupGalleryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the gallery. | [optional] 
**Description** | Pointer to **string** | Description of the gallery. | [optional] 
**MembersOnly** | Pointer to **bool** | Whether the gallery is members only. | [optional] [default to false]
**RoleIdsToView** | Pointer to **[]string** |   | [optional] 
**RoleIdsToSubmit** | Pointer to **[]string** |   | [optional] 
**RoleIdsToAutoApprove** | Pointer to **[]string** |   | [optional] 
**RoleIdsToManage** | Pointer to **[]string** |   | [optional] 

## Methods

### NewUpdateGroupGalleryRequest

`func NewUpdateGroupGalleryRequest() *UpdateGroupGalleryRequest`

NewUpdateGroupGalleryRequest instantiates a new UpdateGroupGalleryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupGalleryRequestWithDefaults

`func NewUpdateGroupGalleryRequestWithDefaults() *UpdateGroupGalleryRequest`

NewUpdateGroupGalleryRequestWithDefaults instantiates a new UpdateGroupGalleryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateGroupGalleryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGroupGalleryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGroupGalleryRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGroupGalleryRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateGroupGalleryRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGroupGalleryRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGroupGalleryRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGroupGalleryRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMembersOnly

`func (o *UpdateGroupGalleryRequest) GetMembersOnly() bool`

GetMembersOnly returns the MembersOnly field if non-nil, zero value otherwise.

### GetMembersOnlyOk

`func (o *UpdateGroupGalleryRequest) GetMembersOnlyOk() (*bool, bool)`

GetMembersOnlyOk returns a tuple with the MembersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersOnly

`func (o *UpdateGroupGalleryRequest) SetMembersOnly(v bool)`

SetMembersOnly sets MembersOnly field to given value.

### HasMembersOnly

`func (o *UpdateGroupGalleryRequest) HasMembersOnly() bool`

HasMembersOnly returns a boolean if a field has been set.

### GetRoleIdsToView

`func (o *UpdateGroupGalleryRequest) GetRoleIdsToView() []string`

GetRoleIdsToView returns the RoleIdsToView field if non-nil, zero value otherwise.

### GetRoleIdsToViewOk

`func (o *UpdateGroupGalleryRequest) GetRoleIdsToViewOk() (*[]string, bool)`

GetRoleIdsToViewOk returns a tuple with the RoleIdsToView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIdsToView

`func (o *UpdateGroupGalleryRequest) SetRoleIdsToView(v []string)`

SetRoleIdsToView sets RoleIdsToView field to given value.

### HasRoleIdsToView

`func (o *UpdateGroupGalleryRequest) HasRoleIdsToView() bool`

HasRoleIdsToView returns a boolean if a field has been set.

### SetRoleIdsToViewNil

`func (o *UpdateGroupGalleryRequest) SetRoleIdsToViewNil(b bool)`

 SetRoleIdsToViewNil sets the value for RoleIdsToView to be an explicit nil

### UnsetRoleIdsToView
`func (o *UpdateGroupGalleryRequest) UnsetRoleIdsToView()`

UnsetRoleIdsToView ensures that no value is present for RoleIdsToView, not even an explicit nil
### GetRoleIdsToSubmit

`func (o *UpdateGroupGalleryRequest) GetRoleIdsToSubmit() []string`

GetRoleIdsToSubmit returns the RoleIdsToSubmit field if non-nil, zero value otherwise.

### GetRoleIdsToSubmitOk

`func (o *UpdateGroupGalleryRequest) GetRoleIdsToSubmitOk() (*[]string, bool)`

GetRoleIdsToSubmitOk returns a tuple with the RoleIdsToSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIdsToSubmit

`func (o *UpdateGroupGalleryRequest) SetRoleIdsToSubmit(v []string)`

SetRoleIdsToSubmit sets RoleIdsToSubmit field to given value.

### HasRoleIdsToSubmit

`func (o *UpdateGroupGalleryRequest) HasRoleIdsToSubmit() bool`

HasRoleIdsToSubmit returns a boolean if a field has been set.

### SetRoleIdsToSubmitNil

`func (o *UpdateGroupGalleryRequest) SetRoleIdsToSubmitNil(b bool)`

 SetRoleIdsToSubmitNil sets the value for RoleIdsToSubmit to be an explicit nil

### UnsetRoleIdsToSubmit
`func (o *UpdateGroupGalleryRequest) UnsetRoleIdsToSubmit()`

UnsetRoleIdsToSubmit ensures that no value is present for RoleIdsToSubmit, not even an explicit nil
### GetRoleIdsToAutoApprove

`func (o *UpdateGroupGalleryRequest) GetRoleIdsToAutoApprove() []string`

GetRoleIdsToAutoApprove returns the RoleIdsToAutoApprove field if non-nil, zero value otherwise.

### GetRoleIdsToAutoApproveOk

`func (o *UpdateGroupGalleryRequest) GetRoleIdsToAutoApproveOk() (*[]string, bool)`

GetRoleIdsToAutoApproveOk returns a tuple with the RoleIdsToAutoApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIdsToAutoApprove

`func (o *UpdateGroupGalleryRequest) SetRoleIdsToAutoApprove(v []string)`

SetRoleIdsToAutoApprove sets RoleIdsToAutoApprove field to given value.

### HasRoleIdsToAutoApprove

`func (o *UpdateGroupGalleryRequest) HasRoleIdsToAutoApprove() bool`

HasRoleIdsToAutoApprove returns a boolean if a field has been set.

### SetRoleIdsToAutoApproveNil

`func (o *UpdateGroupGalleryRequest) SetRoleIdsToAutoApproveNil(b bool)`

 SetRoleIdsToAutoApproveNil sets the value for RoleIdsToAutoApprove to be an explicit nil

### UnsetRoleIdsToAutoApprove
`func (o *UpdateGroupGalleryRequest) UnsetRoleIdsToAutoApprove()`

UnsetRoleIdsToAutoApprove ensures that no value is present for RoleIdsToAutoApprove, not even an explicit nil
### GetRoleIdsToManage

`func (o *UpdateGroupGalleryRequest) GetRoleIdsToManage() []string`

GetRoleIdsToManage returns the RoleIdsToManage field if non-nil, zero value otherwise.

### GetRoleIdsToManageOk

`func (o *UpdateGroupGalleryRequest) GetRoleIdsToManageOk() (*[]string, bool)`

GetRoleIdsToManageOk returns a tuple with the RoleIdsToManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIdsToManage

`func (o *UpdateGroupGalleryRequest) SetRoleIdsToManage(v []string)`

SetRoleIdsToManage sets RoleIdsToManage field to given value.

### HasRoleIdsToManage

`func (o *UpdateGroupGalleryRequest) HasRoleIdsToManage() bool`

HasRoleIdsToManage returns a boolean if a field has been set.

### SetRoleIdsToManageNil

`func (o *UpdateGroupGalleryRequest) SetRoleIdsToManageNil(b bool)`

 SetRoleIdsToManageNil sets the value for RoleIdsToManage to be an explicit nil

### UnsetRoleIdsToManage
`func (o *UpdateGroupGalleryRequest) UnsetRoleIdsToManage()`

UnsetRoleIdsToManage ensures that no value is present for RoleIdsToManage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


