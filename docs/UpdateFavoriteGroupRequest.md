# UpdateFavoriteGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to [**FavoriteGroupVisibility**](FavoriteGroupVisibility.md) |  | [optional] [default to PRIVATE]
**Tags** | Pointer to **[]string** | Tags on FavoriteGroups are believed to do nothing. | [optional] 

## Methods

### NewUpdateFavoriteGroupRequest

`func NewUpdateFavoriteGroupRequest() *UpdateFavoriteGroupRequest`

NewUpdateFavoriteGroupRequest instantiates a new UpdateFavoriteGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFavoriteGroupRequestWithDefaults

`func NewUpdateFavoriteGroupRequestWithDefaults() *UpdateFavoriteGroupRequest`

NewUpdateFavoriteGroupRequestWithDefaults instantiates a new UpdateFavoriteGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UpdateFavoriteGroupRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateFavoriteGroupRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateFavoriteGroupRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateFavoriteGroupRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateFavoriteGroupRequest) GetVisibility() FavoriteGroupVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateFavoriteGroupRequest) GetVisibilityOk() (*FavoriteGroupVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateFavoriteGroupRequest) SetVisibility(v FavoriteGroupVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateFavoriteGroupRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTags

`func (o *UpdateFavoriteGroupRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateFavoriteGroupRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateFavoriteGroupRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateFavoriteGroupRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


