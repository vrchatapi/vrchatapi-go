# CreateAvatarRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |   | [optional] 
**ImageUrl** | **string** |  | 
**ReleaseStatus** | Pointer to [**ReleaseStatus**](ReleaseStatus.md) |  | [optional] [default to PUBLIC]
**Version** | Pointer to **float32** |  | [optional] [default to 1]
**UnityPackageUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateAvatarRequest

`func NewCreateAvatarRequest(name string, imageUrl string, ) *CreateAvatarRequest`

NewCreateAvatarRequest instantiates a new CreateAvatarRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAvatarRequestWithDefaults

`func NewCreateAvatarRequestWithDefaults() *CreateAvatarRequest`

NewCreateAvatarRequestWithDefaults instantiates a new CreateAvatarRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetUrl

`func (o *CreateAvatarRequest) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *CreateAvatarRequest) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *CreateAvatarRequest) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.

### HasAssetUrl

`func (o *CreateAvatarRequest) HasAssetUrl() bool`

HasAssetUrl returns a boolean if a field has been set.

### GetId

`func (o *CreateAvatarRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAvatarRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAvatarRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateAvatarRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateAvatarRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAvatarRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAvatarRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAvatarRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAvatarRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAvatarRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAvatarRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *CreateAvatarRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateAvatarRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateAvatarRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateAvatarRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetImageUrl

`func (o *CreateAvatarRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateAvatarRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateAvatarRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetReleaseStatus

`func (o *CreateAvatarRequest) GetReleaseStatus() ReleaseStatus`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *CreateAvatarRequest) GetReleaseStatusOk() (*ReleaseStatus, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *CreateAvatarRequest) SetReleaseStatus(v ReleaseStatus)`

SetReleaseStatus sets ReleaseStatus field to given value.

### HasReleaseStatus

`func (o *CreateAvatarRequest) HasReleaseStatus() bool`

HasReleaseStatus returns a boolean if a field has been set.

### GetVersion

`func (o *CreateAvatarRequest) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateAvatarRequest) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateAvatarRequest) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateAvatarRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUnityPackageUrl

`func (o *CreateAvatarRequest) GetUnityPackageUrl() string`

GetUnityPackageUrl returns the UnityPackageUrl field if non-nil, zero value otherwise.

### GetUnityPackageUrlOk

`func (o *CreateAvatarRequest) GetUnityPackageUrlOk() (*string, bool)`

GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackageUrl

`func (o *CreateAvatarRequest) SetUnityPackageUrl(v string)`

SetUnityPackageUrl sets UnityPackageUrl field to given value.

### HasUnityPackageUrl

`func (o *CreateAvatarRequest) HasUnityPackageUrl() bool`

HasUnityPackageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


