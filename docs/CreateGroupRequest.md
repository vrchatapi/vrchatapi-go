# CreateGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ShortCode** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**JoinState** | Pointer to [**GroupJoinState**](GroupJoinState.md) |  | [optional] [default to OPEN]
**IconId** | Pointer to **NullableString** |  | [optional] 
**BannerId** | Pointer to **NullableString** |  | [optional] 
**Privacy** | Pointer to [**GroupPrivacy**](GroupPrivacy.md) |  | [optional] [default to DEFAULT]
**RoleTemplate** | [**GroupRoleTemplate**](GroupRoleTemplate.md) |  | [default to DEFAULT]

## Methods

### NewCreateGroupRequest

`func NewCreateGroupRequest(name string, shortCode string, roleTemplate GroupRoleTemplate, ) *CreateGroupRequest`

NewCreateGroupRequest instantiates a new CreateGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupRequestWithDefaults

`func NewCreateGroupRequestWithDefaults() *CreateGroupRequest`

NewCreateGroupRequestWithDefaults instantiates a new CreateGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetShortCode

`func (o *CreateGroupRequest) GetShortCode() string`

GetShortCode returns the ShortCode field if non-nil, zero value otherwise.

### GetShortCodeOk

`func (o *CreateGroupRequest) GetShortCodeOk() (*string, bool)`

GetShortCodeOk returns a tuple with the ShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCode

`func (o *CreateGroupRequest) SetShortCode(v string)`

SetShortCode sets ShortCode field to given value.


### GetDescription

`func (o *CreateGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJoinState

`func (o *CreateGroupRequest) GetJoinState() GroupJoinState`

GetJoinState returns the JoinState field if non-nil, zero value otherwise.

### GetJoinStateOk

`func (o *CreateGroupRequest) GetJoinStateOk() (*GroupJoinState, bool)`

GetJoinStateOk returns a tuple with the JoinState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinState

`func (o *CreateGroupRequest) SetJoinState(v GroupJoinState)`

SetJoinState sets JoinState field to given value.

### HasJoinState

`func (o *CreateGroupRequest) HasJoinState() bool`

HasJoinState returns a boolean if a field has been set.

### GetIconId

`func (o *CreateGroupRequest) GetIconId() string`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *CreateGroupRequest) GetIconIdOk() (*string, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *CreateGroupRequest) SetIconId(v string)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *CreateGroupRequest) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### SetIconIdNil

`func (o *CreateGroupRequest) SetIconIdNil(b bool)`

 SetIconIdNil sets the value for IconId to be an explicit nil

### UnsetIconId
`func (o *CreateGroupRequest) UnsetIconId()`

UnsetIconId ensures that no value is present for IconId, not even an explicit nil
### GetBannerId

`func (o *CreateGroupRequest) GetBannerId() string`

GetBannerId returns the BannerId field if non-nil, zero value otherwise.

### GetBannerIdOk

`func (o *CreateGroupRequest) GetBannerIdOk() (*string, bool)`

GetBannerIdOk returns a tuple with the BannerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerId

`func (o *CreateGroupRequest) SetBannerId(v string)`

SetBannerId sets BannerId field to given value.

### HasBannerId

`func (o *CreateGroupRequest) HasBannerId() bool`

HasBannerId returns a boolean if a field has been set.

### SetBannerIdNil

`func (o *CreateGroupRequest) SetBannerIdNil(b bool)`

 SetBannerIdNil sets the value for BannerId to be an explicit nil

### UnsetBannerId
`func (o *CreateGroupRequest) UnsetBannerId()`

UnsetBannerId ensures that no value is present for BannerId, not even an explicit nil
### GetPrivacy

`func (o *CreateGroupRequest) GetPrivacy() GroupPrivacy`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *CreateGroupRequest) GetPrivacyOk() (*GroupPrivacy, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *CreateGroupRequest) SetPrivacy(v GroupPrivacy)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *CreateGroupRequest) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetRoleTemplate

`func (o *CreateGroupRequest) GetRoleTemplate() GroupRoleTemplate`

GetRoleTemplate returns the RoleTemplate field if non-nil, zero value otherwise.

### GetRoleTemplateOk

`func (o *CreateGroupRequest) GetRoleTemplateOk() (*GroupRoleTemplate, bool)`

GetRoleTemplateOk returns a tuple with the RoleTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTemplate

`func (o *CreateGroupRequest) SetRoleTemplate(v GroupRoleTemplate)`

SetRoleTemplate sets RoleTemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


