# CreateGroupGalleryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the gallery. | 
**Description** | Pointer to **string** | Description of the gallery. | [optional] 
**MembersOnly** | Pointer to **bool** | Whether the gallery is members only. | [optional] [default to false]
**RoleIdsToView** | Pointer to **[]string** |   | [optional] 
**RoleIdsToSubmit** | Pointer to **[]string** |   | [optional] 
**RoleIdsToAutoApprove** | Pointer to **[]string** |   | [optional] 
**RoleIdsToManage** | Pointer to **[]string** |   | [optional] 

## Methods

### NewCreateGroupGalleryRequest

`func NewCreateGroupGalleryRequest(name string, ) *CreateGroupGalleryRequest`

NewCreateGroupGalleryRequest instantiates a new CreateGroupGalleryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupGalleryRequestWithDefaults

`func NewCreateGroupGalleryRequestWithDefaults() *CreateGroupGalleryRequest`

NewCreateGroupGalleryRequestWithDefaults instantiates a new CreateGroupGalleryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGroupGalleryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroupGalleryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroupGalleryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateGroupGalleryRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGroupGalleryRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGroupGalleryRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGroupGalleryRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMembersOnly

`func (o *CreateGroupGalleryRequest) GetMembersOnly() bool`

GetMembersOnly returns the MembersOnly field if non-nil, zero value otherwise.

### GetMembersOnlyOk

`func (o *CreateGroupGalleryRequest) GetMembersOnlyOk() (*bool, bool)`

GetMembersOnlyOk returns a tuple with the MembersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersOnly

`func (o *CreateGroupGalleryRequest) SetMembersOnly(v bool)`

SetMembersOnly sets MembersOnly field to given value.

### HasMembersOnly

`func (o *CreateGroupGalleryRequest) HasMembersOnly() bool`

HasMembersOnly returns a boolean if a field has been set.

### GetRoleIdsToView

`func (o *CreateGroupGalleryRequest) GetRoleIdsToView() []string`

GetRoleIdsToView returns the RoleIdsToView field if non-nil, zero value otherwise.

### GetRoleIdsToViewOk

`func (o *CreateGroupGalleryRequest) GetRoleIdsToViewOk() (*[]string, bool)`

GetRoleIdsToViewOk returns a tuple with the RoleIdsToView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIdsToView

`func (o *CreateGroupGalleryRequest) SetRoleIdsToView(v []string)`

SetRoleIdsToView sets RoleIdsToView field to given value.

### HasRoleIdsToView

`func (o *CreateGroupGalleryRequest) HasRoleIdsToView() bool`

HasRoleIdsToView returns a boolean if a field has been set.

### SetRoleIdsToViewNil

`func (o *CreateGroupGalleryRequest) SetRoleIdsToViewNil(b bool)`

 SetRoleIdsToViewNil sets the value for RoleIdsToView to be an explicit nil

### UnsetRoleIdsToView
`func (o *CreateGroupGalleryRequest) UnsetRoleIdsToView()`

UnsetRoleIdsToView ensures that no value is present for RoleIdsToView, not even an explicit nil
### GetRoleIdsToSubmit

`func (o *CreateGroupGalleryRequest) GetRoleIdsToSubmit() []string`

GetRoleIdsToSubmit returns the RoleIdsToSubmit field if non-nil, zero value otherwise.

### GetRoleIdsToSubmitOk

`func (o *CreateGroupGalleryRequest) GetRoleIdsToSubmitOk() (*[]string, bool)`

GetRoleIdsToSubmitOk returns a tuple with the RoleIdsToSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIdsToSubmit

`func (o *CreateGroupGalleryRequest) SetRoleIdsToSubmit(v []string)`

SetRoleIdsToSubmit sets RoleIdsToSubmit field to given value.

### HasRoleIdsToSubmit

`func (o *CreateGroupGalleryRequest) HasRoleIdsToSubmit() bool`

HasRoleIdsToSubmit returns a boolean if a field has been set.

### SetRoleIdsToSubmitNil

`func (o *CreateGroupGalleryRequest) SetRoleIdsToSubmitNil(b bool)`

 SetRoleIdsToSubmitNil sets the value for RoleIdsToSubmit to be an explicit nil

### UnsetRoleIdsToSubmit
`func (o *CreateGroupGalleryRequest) UnsetRoleIdsToSubmit()`

UnsetRoleIdsToSubmit ensures that no value is present for RoleIdsToSubmit, not even an explicit nil
### GetRoleIdsToAutoApprove

`func (o *CreateGroupGalleryRequest) GetRoleIdsToAutoApprove() []string`

GetRoleIdsToAutoApprove returns the RoleIdsToAutoApprove field if non-nil, zero value otherwise.

### GetRoleIdsToAutoApproveOk

`func (o *CreateGroupGalleryRequest) GetRoleIdsToAutoApproveOk() (*[]string, bool)`

GetRoleIdsToAutoApproveOk returns a tuple with the RoleIdsToAutoApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIdsToAutoApprove

`func (o *CreateGroupGalleryRequest) SetRoleIdsToAutoApprove(v []string)`

SetRoleIdsToAutoApprove sets RoleIdsToAutoApprove field to given value.

### HasRoleIdsToAutoApprove

`func (o *CreateGroupGalleryRequest) HasRoleIdsToAutoApprove() bool`

HasRoleIdsToAutoApprove returns a boolean if a field has been set.

### SetRoleIdsToAutoApproveNil

`func (o *CreateGroupGalleryRequest) SetRoleIdsToAutoApproveNil(b bool)`

 SetRoleIdsToAutoApproveNil sets the value for RoleIdsToAutoApprove to be an explicit nil

### UnsetRoleIdsToAutoApprove
`func (o *CreateGroupGalleryRequest) UnsetRoleIdsToAutoApprove()`

UnsetRoleIdsToAutoApprove ensures that no value is present for RoleIdsToAutoApprove, not even an explicit nil
### GetRoleIdsToManage

`func (o *CreateGroupGalleryRequest) GetRoleIdsToManage() []string`

GetRoleIdsToManage returns the RoleIdsToManage field if non-nil, zero value otherwise.

### GetRoleIdsToManageOk

`func (o *CreateGroupGalleryRequest) GetRoleIdsToManageOk() (*[]string, bool)`

GetRoleIdsToManageOk returns a tuple with the RoleIdsToManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIdsToManage

`func (o *CreateGroupGalleryRequest) SetRoleIdsToManage(v []string)`

SetRoleIdsToManage sets RoleIdsToManage field to given value.

### HasRoleIdsToManage

`func (o *CreateGroupGalleryRequest) HasRoleIdsToManage() bool`

HasRoleIdsToManage returns a boolean if a field has been set.

### SetRoleIdsToManageNil

`func (o *CreateGroupGalleryRequest) SetRoleIdsToManageNil(b bool)`

 SetRoleIdsToManageNil sets the value for RoleIdsToManage to be an explicit nil

### UnsetRoleIdsToManage
`func (o *CreateGroupGalleryRequest) UnsetRoleIdsToManage()`

UnsetRoleIdsToManage ensures that no value is present for RoleIdsToManage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


