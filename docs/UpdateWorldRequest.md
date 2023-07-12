# UpdateWorldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetUrl** | Pointer to **string** |  | [optional] 
**AssetVersion** | Pointer to **string** |  | [optional] 
**AuthorId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**AuthorName** | Pointer to **string** |  | [optional] 
**Capacity** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** | This can be &#x60;standalonewindows&#x60; or &#x60;android&#x60;, but can also pretty much be any random Unity verison such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | [optional] 
**ReleaseStatus** | Pointer to [**ReleaseStatus**](ReleaseStatus.md) |  | [optional] [default to PUBLIC]
**Tags** | Pointer to **[]string** |   | [optional] 
**UnityPackageUrl** | Pointer to **string** |  | [optional] 
**UnityVersion** | Pointer to **string** |  | [optional] [default to "5.3.4p1"]

## Methods

### NewUpdateWorldRequest

`func NewUpdateWorldRequest() *UpdateWorldRequest`

NewUpdateWorldRequest instantiates a new UpdateWorldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorldRequestWithDefaults

`func NewUpdateWorldRequestWithDefaults() *UpdateWorldRequest`

NewUpdateWorldRequestWithDefaults instantiates a new UpdateWorldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetUrl

`func (o *UpdateWorldRequest) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *UpdateWorldRequest) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *UpdateWorldRequest) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.

### HasAssetUrl

`func (o *UpdateWorldRequest) HasAssetUrl() bool`

HasAssetUrl returns a boolean if a field has been set.

### GetAssetVersion

`func (o *UpdateWorldRequest) GetAssetVersion() string`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *UpdateWorldRequest) GetAssetVersionOk() (*string, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *UpdateWorldRequest) SetAssetVersion(v string)`

SetAssetVersion sets AssetVersion field to given value.

### HasAssetVersion

`func (o *UpdateWorldRequest) HasAssetVersion() bool`

HasAssetVersion returns a boolean if a field has been set.

### GetAuthorId

`func (o *UpdateWorldRequest) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *UpdateWorldRequest) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *UpdateWorldRequest) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *UpdateWorldRequest) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetAuthorName

`func (o *UpdateWorldRequest) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *UpdateWorldRequest) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *UpdateWorldRequest) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *UpdateWorldRequest) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### GetCapacity

`func (o *UpdateWorldRequest) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *UpdateWorldRequest) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *UpdateWorldRequest) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *UpdateWorldRequest) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateWorldRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateWorldRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateWorldRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateWorldRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImageUrl

`func (o *UpdateWorldRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UpdateWorldRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UpdateWorldRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UpdateWorldRequest) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetName

`func (o *UpdateWorldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWorldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWorldRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateWorldRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *UpdateWorldRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdateWorldRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdateWorldRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *UpdateWorldRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetReleaseStatus

`func (o *UpdateWorldRequest) GetReleaseStatus() ReleaseStatus`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *UpdateWorldRequest) GetReleaseStatusOk() (*ReleaseStatus, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *UpdateWorldRequest) SetReleaseStatus(v ReleaseStatus)`

SetReleaseStatus sets ReleaseStatus field to given value.

### HasReleaseStatus

`func (o *UpdateWorldRequest) HasReleaseStatus() bool`

HasReleaseStatus returns a boolean if a field has been set.

### GetTags

`func (o *UpdateWorldRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateWorldRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateWorldRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateWorldRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnityPackageUrl

`func (o *UpdateWorldRequest) GetUnityPackageUrl() string`

GetUnityPackageUrl returns the UnityPackageUrl field if non-nil, zero value otherwise.

### GetUnityPackageUrlOk

`func (o *UpdateWorldRequest) GetUnityPackageUrlOk() (*string, bool)`

GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackageUrl

`func (o *UpdateWorldRequest) SetUnityPackageUrl(v string)`

SetUnityPackageUrl sets UnityPackageUrl field to given value.

### HasUnityPackageUrl

`func (o *UpdateWorldRequest) HasUnityPackageUrl() bool`

HasUnityPackageUrl returns a boolean if a field has been set.

### GetUnityVersion

`func (o *UpdateWorldRequest) GetUnityVersion() string`

GetUnityVersion returns the UnityVersion field if non-nil, zero value otherwise.

### GetUnityVersionOk

`func (o *UpdateWorldRequest) GetUnityVersionOk() (*string, bool)`

GetUnityVersionOk returns a tuple with the UnityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityVersion

`func (o *UpdateWorldRequest) SetUnityVersion(v string)`

SetUnityVersion sets UnityVersion field to given value.

### HasUnityVersion

`func (o *UpdateWorldRequest) HasUnityVersion() bool`

HasUnityVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


