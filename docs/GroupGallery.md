# GroupGallery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the gallery. | [optional] 
**Description** | Pointer to **string** | Description of the gallery. | [optional] 
**MembersOnly** | Pointer to **bool** | Whether the gallery is members only. | [optional] [default to false]
**RoleIdsToView** | Pointer to **[]string** |   | [optional] 
**RoleIdsToSubmit** | Pointer to **[]string** |   | [optional] 
**RoleIdsToAutoApprove** | Pointer to **[]string** |   | [optional] 
**RoleIdsToManage** | Pointer to **[]string** |   | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGroupGallery

`func NewGroupGallery() *GroupGallery`

NewGroupGallery instantiates a new GroupGallery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupGalleryWithDefaults

`func NewGroupGalleryWithDefaults() *GroupGallery`

NewGroupGalleryWithDefaults instantiates a new GroupGallery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupGallery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupGallery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupGallery) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupGallery) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GroupGallery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupGallery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupGallery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupGallery) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GroupGallery) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupGallery) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupGallery) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupGallery) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMembersOnly

`func (o *GroupGallery) GetMembersOnly() bool`

GetMembersOnly returns the MembersOnly field if non-nil, zero value otherwise.

### GetMembersOnlyOk

`func (o *GroupGallery) GetMembersOnlyOk() (*bool, bool)`

GetMembersOnlyOk returns a tuple with the MembersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersOnly

`func (o *GroupGallery) SetMembersOnly(v bool)`

SetMembersOnly sets MembersOnly field to given value.

### HasMembersOnly

`func (o *GroupGallery) HasMembersOnly() bool`

HasMembersOnly returns a boolean if a field has been set.

### GetRoleIdsToView

`func (o *GroupGallery) GetRoleIdsToView() []string`

GetRoleIdsToView returns the RoleIdsToView field if non-nil, zero value otherwise.

### GetRoleIdsToViewOk

`func (o *GroupGallery) GetRoleIdsToViewOk() (*[]string, bool)`

GetRoleIdsToViewOk returns a tuple with the RoleIdsToView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIdsToView

`func (o *GroupGallery) SetRoleIdsToView(v []string)`

SetRoleIdsToView sets RoleIdsToView field to given value.

### HasRoleIdsToView

`func (o *GroupGallery) HasRoleIdsToView() bool`

HasRoleIdsToView returns a boolean if a field has been set.

### SetRoleIdsToViewNil

`func (o *GroupGallery) SetRoleIdsToViewNil(b bool)`

 SetRoleIdsToViewNil sets the value for RoleIdsToView to be an explicit nil

### UnsetRoleIdsToView
`func (o *GroupGallery) UnsetRoleIdsToView()`

UnsetRoleIdsToView ensures that no value is present for RoleIdsToView, not even an explicit nil
### GetRoleIdsToSubmit

`func (o *GroupGallery) GetRoleIdsToSubmit() []string`

GetRoleIdsToSubmit returns the RoleIdsToSubmit field if non-nil, zero value otherwise.

### GetRoleIdsToSubmitOk

`func (o *GroupGallery) GetRoleIdsToSubmitOk() (*[]string, bool)`

GetRoleIdsToSubmitOk returns a tuple with the RoleIdsToSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIdsToSubmit

`func (o *GroupGallery) SetRoleIdsToSubmit(v []string)`

SetRoleIdsToSubmit sets RoleIdsToSubmit field to given value.

### HasRoleIdsToSubmit

`func (o *GroupGallery) HasRoleIdsToSubmit() bool`

HasRoleIdsToSubmit returns a boolean if a field has been set.

### SetRoleIdsToSubmitNil

`func (o *GroupGallery) SetRoleIdsToSubmitNil(b bool)`

 SetRoleIdsToSubmitNil sets the value for RoleIdsToSubmit to be an explicit nil

### UnsetRoleIdsToSubmit
`func (o *GroupGallery) UnsetRoleIdsToSubmit()`

UnsetRoleIdsToSubmit ensures that no value is present for RoleIdsToSubmit, not even an explicit nil
### GetRoleIdsToAutoApprove

`func (o *GroupGallery) GetRoleIdsToAutoApprove() []string`

GetRoleIdsToAutoApprove returns the RoleIdsToAutoApprove field if non-nil, zero value otherwise.

### GetRoleIdsToAutoApproveOk

`func (o *GroupGallery) GetRoleIdsToAutoApproveOk() (*[]string, bool)`

GetRoleIdsToAutoApproveOk returns a tuple with the RoleIdsToAutoApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIdsToAutoApprove

`func (o *GroupGallery) SetRoleIdsToAutoApprove(v []string)`

SetRoleIdsToAutoApprove sets RoleIdsToAutoApprove field to given value.

### HasRoleIdsToAutoApprove

`func (o *GroupGallery) HasRoleIdsToAutoApprove() bool`

HasRoleIdsToAutoApprove returns a boolean if a field has been set.

### SetRoleIdsToAutoApproveNil

`func (o *GroupGallery) SetRoleIdsToAutoApproveNil(b bool)`

 SetRoleIdsToAutoApproveNil sets the value for RoleIdsToAutoApprove to be an explicit nil

### UnsetRoleIdsToAutoApprove
`func (o *GroupGallery) UnsetRoleIdsToAutoApprove()`

UnsetRoleIdsToAutoApprove ensures that no value is present for RoleIdsToAutoApprove, not even an explicit nil
### GetRoleIdsToManage

`func (o *GroupGallery) GetRoleIdsToManage() []string`

GetRoleIdsToManage returns the RoleIdsToManage field if non-nil, zero value otherwise.

### GetRoleIdsToManageOk

`func (o *GroupGallery) GetRoleIdsToManageOk() (*[]string, bool)`

GetRoleIdsToManageOk returns a tuple with the RoleIdsToManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIdsToManage

`func (o *GroupGallery) SetRoleIdsToManage(v []string)`

SetRoleIdsToManage sets RoleIdsToManage field to given value.

### HasRoleIdsToManage

`func (o *GroupGallery) HasRoleIdsToManage() bool`

HasRoleIdsToManage returns a boolean if a field has been set.

### SetRoleIdsToManageNil

`func (o *GroupGallery) SetRoleIdsToManageNil(b bool)`

 SetRoleIdsToManageNil sets the value for RoleIdsToManage to be an explicit nil

### UnsetRoleIdsToManage
`func (o *GroupGallery) UnsetRoleIdsToManage()`

UnsetRoleIdsToManage ensures that no value is present for RoleIdsToManage, not even an explicit nil
### GetCreatedAt

`func (o *GroupGallery) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupGallery) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupGallery) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupGallery) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GroupGallery) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GroupGallery) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GroupGallery) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GroupGallery) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


