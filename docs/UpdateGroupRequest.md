# UpdateGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ShortCode** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**JoinState** | Pointer to [**GroupJoinState**](GroupJoinState.md) |  | [optional] [default to OPEN]
**IconId** | Pointer to **NullableString** |  | [optional] 
**BannerId** | Pointer to **NullableString** |  | [optional] 
**Languages** | Pointer to **[]string** | 3 letter language code | [optional] 
**Links** | Pointer to **[]string** |  | [optional] 
**Rules** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |   | [optional] 

## Methods

### NewUpdateGroupRequest

`func NewUpdateGroupRequest() *UpdateGroupRequest`

NewUpdateGroupRequest instantiates a new UpdateGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupRequestWithDefaults

`func NewUpdateGroupRequestWithDefaults() *UpdateGroupRequest`

NewUpdateGroupRequestWithDefaults instantiates a new UpdateGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortCode

`func (o *UpdateGroupRequest) GetShortCode() string`

GetShortCode returns the ShortCode field if non-nil, zero value otherwise.

### GetShortCodeOk

`func (o *UpdateGroupRequest) GetShortCodeOk() (*string, bool)`

GetShortCodeOk returns a tuple with the ShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCode

`func (o *UpdateGroupRequest) SetShortCode(v string)`

SetShortCode sets ShortCode field to given value.

### HasShortCode

`func (o *UpdateGroupRequest) HasShortCode() bool`

HasShortCode returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJoinState

`func (o *UpdateGroupRequest) GetJoinState() GroupJoinState`

GetJoinState returns the JoinState field if non-nil, zero value otherwise.

### GetJoinStateOk

`func (o *UpdateGroupRequest) GetJoinStateOk() (*GroupJoinState, bool)`

GetJoinStateOk returns a tuple with the JoinState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinState

`func (o *UpdateGroupRequest) SetJoinState(v GroupJoinState)`

SetJoinState sets JoinState field to given value.

### HasJoinState

`func (o *UpdateGroupRequest) HasJoinState() bool`

HasJoinState returns a boolean if a field has been set.

### GetIconId

`func (o *UpdateGroupRequest) GetIconId() string`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *UpdateGroupRequest) GetIconIdOk() (*string, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *UpdateGroupRequest) SetIconId(v string)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *UpdateGroupRequest) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### SetIconIdNil

`func (o *UpdateGroupRequest) SetIconIdNil(b bool)`

 SetIconIdNil sets the value for IconId to be an explicit nil

### UnsetIconId
`func (o *UpdateGroupRequest) UnsetIconId()`

UnsetIconId ensures that no value is present for IconId, not even an explicit nil
### GetBannerId

`func (o *UpdateGroupRequest) GetBannerId() string`

GetBannerId returns the BannerId field if non-nil, zero value otherwise.

### GetBannerIdOk

`func (o *UpdateGroupRequest) GetBannerIdOk() (*string, bool)`

GetBannerIdOk returns a tuple with the BannerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerId

`func (o *UpdateGroupRequest) SetBannerId(v string)`

SetBannerId sets BannerId field to given value.

### HasBannerId

`func (o *UpdateGroupRequest) HasBannerId() bool`

HasBannerId returns a boolean if a field has been set.

### SetBannerIdNil

`func (o *UpdateGroupRequest) SetBannerIdNil(b bool)`

 SetBannerIdNil sets the value for BannerId to be an explicit nil

### UnsetBannerId
`func (o *UpdateGroupRequest) UnsetBannerId()`

UnsetBannerId ensures that no value is present for BannerId, not even an explicit nil
### GetLanguages

`func (o *UpdateGroupRequest) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *UpdateGroupRequest) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *UpdateGroupRequest) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *UpdateGroupRequest) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLinks

`func (o *UpdateGroupRequest) GetLinks() []string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateGroupRequest) GetLinksOk() (*[]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateGroupRequest) SetLinks(v []string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpdateGroupRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRules

`func (o *UpdateGroupRequest) GetRules() string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateGroupRequest) GetRulesOk() (*string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateGroupRequest) SetRules(v string)`

SetRules sets Rules field to given value.

### HasRules

`func (o *UpdateGroupRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTags

`func (o *UpdateGroupRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateGroupRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateGroupRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateGroupRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


