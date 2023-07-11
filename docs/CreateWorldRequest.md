# CreateWorldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetUrl** | **string** |  | 
**AssetVersion** | Pointer to **int32** |  | [optional] 
**AuthorId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**AuthorName** | Pointer to **string** |  | [optional] 
**Capacity** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | [optional] 
**ImageUrl** | **string** |  | 
**Name** | **string** |  | 
**Platform** | Pointer to **string** | This can be &#x60;standalonewindows&#x60; or &#x60;android&#x60;, but can also pretty much be any random Unity verison such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | [optional] 
**ReleaseStatus** | Pointer to [**ReleaseStatus**](ReleaseStatus.md) |  | [optional] [default to PUBLIC]
**Tags** | Pointer to **[]string** |   | [optional] 
**UnityPackageUrl** | Pointer to **string** |  | [optional] 
**UnityVersion** | Pointer to **string** |  | [optional] [default to "5.3.4p1"]

## Methods

### NewCreateWorldRequest

`func NewCreateWorldRequest(assetUrl string, imageUrl string, name string, ) *CreateWorldRequest`

NewCreateWorldRequest instantiates a new CreateWorldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorldRequestWithDefaults

`func NewCreateWorldRequestWithDefaults() *CreateWorldRequest`

NewCreateWorldRequestWithDefaults instantiates a new CreateWorldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetUrl

`func (o *CreateWorldRequest) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *CreateWorldRequest) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *CreateWorldRequest) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.


### GetAssetVersion

`func (o *CreateWorldRequest) GetAssetVersion() int32`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *CreateWorldRequest) GetAssetVersionOk() (*int32, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *CreateWorldRequest) SetAssetVersion(v int32)`

SetAssetVersion sets AssetVersion field to given value.

### HasAssetVersion

`func (o *CreateWorldRequest) HasAssetVersion() bool`

HasAssetVersion returns a boolean if a field has been set.

### GetAuthorId

`func (o *CreateWorldRequest) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *CreateWorldRequest) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *CreateWorldRequest) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *CreateWorldRequest) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetAuthorName

`func (o *CreateWorldRequest) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *CreateWorldRequest) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *CreateWorldRequest) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *CreateWorldRequest) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### GetCapacity

`func (o *CreateWorldRequest) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *CreateWorldRequest) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *CreateWorldRequest) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *CreateWorldRequest) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetDescription

`func (o *CreateWorldRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateWorldRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateWorldRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateWorldRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CreateWorldRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateWorldRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateWorldRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateWorldRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *CreateWorldRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateWorldRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateWorldRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetName

`func (o *CreateWorldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWorldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWorldRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPlatform

`func (o *CreateWorldRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateWorldRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateWorldRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CreateWorldRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetReleaseStatus

`func (o *CreateWorldRequest) GetReleaseStatus() ReleaseStatus`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *CreateWorldRequest) GetReleaseStatusOk() (*ReleaseStatus, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *CreateWorldRequest) SetReleaseStatus(v ReleaseStatus)`

SetReleaseStatus sets ReleaseStatus field to given value.

### HasReleaseStatus

`func (o *CreateWorldRequest) HasReleaseStatus() bool`

HasReleaseStatus returns a boolean if a field has been set.

### GetTags

`func (o *CreateWorldRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateWorldRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateWorldRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateWorldRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnityPackageUrl

`func (o *CreateWorldRequest) GetUnityPackageUrl() string`

GetUnityPackageUrl returns the UnityPackageUrl field if non-nil, zero value otherwise.

### GetUnityPackageUrlOk

`func (o *CreateWorldRequest) GetUnityPackageUrlOk() (*string, bool)`

GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackageUrl

`func (o *CreateWorldRequest) SetUnityPackageUrl(v string)`

SetUnityPackageUrl sets UnityPackageUrl field to given value.

### HasUnityPackageUrl

`func (o *CreateWorldRequest) HasUnityPackageUrl() bool`

HasUnityPackageUrl returns a boolean if a field has been set.

### GetUnityVersion

`func (o *CreateWorldRequest) GetUnityVersion() string`

GetUnityVersion returns the UnityVersion field if non-nil, zero value otherwise.

### GetUnityVersionOk

`func (o *CreateWorldRequest) GetUnityVersionOk() (*string, bool)`

GetUnityVersionOk returns a tuple with the UnityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityVersion

`func (o *CreateWorldRequest) SetUnityVersion(v string)`

SetUnityVersion sets UnityVersion field to given value.

### HasUnityVersion

`func (o *CreateWorldRequest) HasUnityVersion() bool`

HasUnityVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


