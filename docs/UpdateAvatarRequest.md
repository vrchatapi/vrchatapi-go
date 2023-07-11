# UpdateAvatarRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |   | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**ReleaseStatus** | Pointer to [**ReleaseStatus**](ReleaseStatus.md) |  | [optional] [default to PUBLIC]
**Version** | Pointer to **float32** |  | [optional] [default to 1]
**UnityPackageUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateAvatarRequest

`func NewUpdateAvatarRequest() *UpdateAvatarRequest`

NewUpdateAvatarRequest instantiates a new UpdateAvatarRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAvatarRequestWithDefaults

`func NewUpdateAvatarRequestWithDefaults() *UpdateAvatarRequest`

NewUpdateAvatarRequestWithDefaults instantiates a new UpdateAvatarRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetUrl

`func (o *UpdateAvatarRequest) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *UpdateAvatarRequest) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *UpdateAvatarRequest) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.

### HasAssetUrl

`func (o *UpdateAvatarRequest) HasAssetUrl() bool`

HasAssetUrl returns a boolean if a field has been set.

### GetId

`func (o *UpdateAvatarRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAvatarRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAvatarRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateAvatarRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateAvatarRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAvatarRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAvatarRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAvatarRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAvatarRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAvatarRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAvatarRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAvatarRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *UpdateAvatarRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateAvatarRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateAvatarRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateAvatarRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetImageUrl

`func (o *UpdateAvatarRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UpdateAvatarRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UpdateAvatarRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UpdateAvatarRequest) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetReleaseStatus

`func (o *UpdateAvatarRequest) GetReleaseStatus() ReleaseStatus`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *UpdateAvatarRequest) GetReleaseStatusOk() (*ReleaseStatus, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *UpdateAvatarRequest) SetReleaseStatus(v ReleaseStatus)`

SetReleaseStatus sets ReleaseStatus field to given value.

### HasReleaseStatus

`func (o *UpdateAvatarRequest) HasReleaseStatus() bool`

HasReleaseStatus returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateAvatarRequest) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateAvatarRequest) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateAvatarRequest) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateAvatarRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUnityPackageUrl

`func (o *UpdateAvatarRequest) GetUnityPackageUrl() string`

GetUnityPackageUrl returns the UnityPackageUrl field if non-nil, zero value otherwise.

### GetUnityPackageUrlOk

`func (o *UpdateAvatarRequest) GetUnityPackageUrlOk() (*string, bool)`

GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackageUrl

`func (o *UpdateAvatarRequest) SetUnityPackageUrl(v string)`

SetUnityPackageUrl sets UnityPackageUrl field to given value.

### HasUnityPackageUrl

`func (o *UpdateAvatarRequest) HasUnityPackageUrl() bool`

HasUnityPackageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


