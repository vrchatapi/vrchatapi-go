# Avatar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetUrl** | Pointer to **string** | Not present from general serach &#x60;/avatars&#x60;, only on specific requests &#x60;/avatars/{avatarId}&#x60;. | [optional] 
**AssetUrlObject** | Pointer to **map[string]interface{}** | Not present from general serach &#x60;/avatars&#x60;, only on specific requests &#x60;/avatars/{avatarId}&#x60;. **Deprecation:** &#x60;Object&#x60; has unknown usage/fields, and is always empty. Use normal &#x60;Url&#x60; field instead. | [optional] 
**AuthorId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**AuthorName** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Description** | **string** |  | 
**Featured** | **bool** |  | [default to false]
**Id** | **string** |  | 
**ImageUrl** | **string** |  | 
**Name** | **string** |  | 
**ReleaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) |  | [default to PUBLIC]
**Tags** | **[]string** |   | 
**ThumbnailImageUrl** | **string** |  | 
**UnityPackageUrl** | **string** |  | 
**UnityPackageUrlObject** | [**AvatarUnityPackageUrlObject**](AvatarUnityPackageUrlObject.md) |  | 
**UnityPackages** | [**[]UnityPackage**](UnityPackage.md) |  | 
**UpdatedAt** | **time.Time** |  | 
**Version** | **int32** |  | [default to 0]

## Methods

### NewAvatar

`func NewAvatar(authorId string, authorName string, createdAt time.Time, description string, featured bool, id string, imageUrl string, name string, releaseStatus ReleaseStatus, tags []string, thumbnailImageUrl string, unityPackageUrl string, unityPackageUrlObject AvatarUnityPackageUrlObject, unityPackages []UnityPackage, updatedAt time.Time, version int32, ) *Avatar`

NewAvatar instantiates a new Avatar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvatarWithDefaults

`func NewAvatarWithDefaults() *Avatar`

NewAvatarWithDefaults instantiates a new Avatar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetUrl

`func (o *Avatar) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *Avatar) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *Avatar) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.

### HasAssetUrl

`func (o *Avatar) HasAssetUrl() bool`

HasAssetUrl returns a boolean if a field has been set.

### GetAssetUrlObject

`func (o *Avatar) GetAssetUrlObject() map[string]interface{}`

GetAssetUrlObject returns the AssetUrlObject field if non-nil, zero value otherwise.

### GetAssetUrlObjectOk

`func (o *Avatar) GetAssetUrlObjectOk() (*map[string]interface{}, bool)`

GetAssetUrlObjectOk returns a tuple with the AssetUrlObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrlObject

`func (o *Avatar) SetAssetUrlObject(v map[string]interface{})`

SetAssetUrlObject sets AssetUrlObject field to given value.

### HasAssetUrlObject

`func (o *Avatar) HasAssetUrlObject() bool`

HasAssetUrlObject returns a boolean if a field has been set.

### GetAuthorId

`func (o *Avatar) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *Avatar) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *Avatar) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### GetAuthorName

`func (o *Avatar) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *Avatar) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *Avatar) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetCreatedAt

`func (o *Avatar) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Avatar) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Avatar) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *Avatar) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Avatar) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Avatar) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFeatured

`func (o *Avatar) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *Avatar) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *Avatar) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.


### GetId

`func (o *Avatar) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Avatar) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Avatar) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *Avatar) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Avatar) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Avatar) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetName

`func (o *Avatar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Avatar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Avatar) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseStatus

`func (o *Avatar) GetReleaseStatus() ReleaseStatus`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *Avatar) GetReleaseStatusOk() (*ReleaseStatus, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *Avatar) SetReleaseStatus(v ReleaseStatus)`

SetReleaseStatus sets ReleaseStatus field to given value.


### GetTags

`func (o *Avatar) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Avatar) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Avatar) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetThumbnailImageUrl

`func (o *Avatar) GetThumbnailImageUrl() string`

GetThumbnailImageUrl returns the ThumbnailImageUrl field if non-nil, zero value otherwise.

### GetThumbnailImageUrlOk

`func (o *Avatar) GetThumbnailImageUrlOk() (*string, bool)`

GetThumbnailImageUrlOk returns a tuple with the ThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailImageUrl

`func (o *Avatar) SetThumbnailImageUrl(v string)`

SetThumbnailImageUrl sets ThumbnailImageUrl field to given value.


### GetUnityPackageUrl

`func (o *Avatar) GetUnityPackageUrl() string`

GetUnityPackageUrl returns the UnityPackageUrl field if non-nil, zero value otherwise.

### GetUnityPackageUrlOk

`func (o *Avatar) GetUnityPackageUrlOk() (*string, bool)`

GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackageUrl

`func (o *Avatar) SetUnityPackageUrl(v string)`

SetUnityPackageUrl sets UnityPackageUrl field to given value.


### GetUnityPackageUrlObject

`func (o *Avatar) GetUnityPackageUrlObject() AvatarUnityPackageUrlObject`

GetUnityPackageUrlObject returns the UnityPackageUrlObject field if non-nil, zero value otherwise.

### GetUnityPackageUrlObjectOk

`func (o *Avatar) GetUnityPackageUrlObjectOk() (*AvatarUnityPackageUrlObject, bool)`

GetUnityPackageUrlObjectOk returns a tuple with the UnityPackageUrlObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackageUrlObject

`func (o *Avatar) SetUnityPackageUrlObject(v AvatarUnityPackageUrlObject)`

SetUnityPackageUrlObject sets UnityPackageUrlObject field to given value.


### GetUnityPackages

`func (o *Avatar) GetUnityPackages() []UnityPackage`

GetUnityPackages returns the UnityPackages field if non-nil, zero value otherwise.

### GetUnityPackagesOk

`func (o *Avatar) GetUnityPackagesOk() (*[]UnityPackage, bool)`

GetUnityPackagesOk returns a tuple with the UnityPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackages

`func (o *Avatar) SetUnityPackages(v []UnityPackage)`

SetUnityPackages sets UnityPackages field to given value.


### GetUpdatedAt

`func (o *Avatar) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Avatar) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Avatar) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *Avatar) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Avatar) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Avatar) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


