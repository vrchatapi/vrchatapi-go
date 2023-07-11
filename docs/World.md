# World

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**AuthorName** | **string** |  | 
**Capacity** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**Description** | **string** |  | 
**Favorites** | Pointer to **int32** |  | [optional] [default to 0]
**Featured** | **bool** |  | [default to false]
**Heat** | **int32** |  | [default to 0]
**Id** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**ImageUrl** | **string** |  | 
**Instances** | Pointer to **[][]interface{}** | Will always be an empty list when unauthenticated. | [optional] 
**LabsPublicationDate** | **string** |  | 
**Name** | **string** |  | 
**Namespace** | **string** |  | 
**Occupants** | Pointer to **int32** | Will always be &#x60;0&#x60; when unauthenticated. | [optional] [default to 0]
**Organization** | **string** |  | [default to "vrchat"]
**Popularity** | **int32** |  | [default to 0]
**PreviewYoutubeId** | Pointer to **NullableString** |  | [optional] 
**PrivateOccupants** | Pointer to **int32** | Will always be &#x60;0&#x60; when unauthenticated. | [optional] [default to 0]
**PublicOccupants** | Pointer to **int32** | Will always be &#x60;0&#x60; when unauthenticated. | [optional] [default to 0]
**PublicationDate** | **string** |  | 
**ReleaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) |  | [default to PUBLIC]
**Tags** | **[]string** |   | 
**ThumbnailImageUrl** | **string** |  | 
**UnityPackages** | [**[]UnityPackage**](UnityPackage.md) | Empty if unauthenticated. | 
**UpdatedAt** | **time.Time** |  | 
**Version** | **int32** |  | [default to 0]
**Visits** | **int32** |  | [default to 0]

## Methods

### NewWorld

`func NewWorld(authorId string, authorName string, capacity int32, createdAt time.Time, description string, featured bool, heat int32, id string, imageUrl string, labsPublicationDate string, name string, namespace string, organization string, popularity int32, publicationDate string, releaseStatus ReleaseStatus, tags []string, thumbnailImageUrl string, unityPackages []UnityPackage, updatedAt time.Time, version int32, visits int32, ) *World`

NewWorld instantiates a new World object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorldWithDefaults

`func NewWorldWithDefaults() *World`

NewWorldWithDefaults instantiates a new World object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorId

`func (o *World) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *World) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *World) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### GetAuthorName

`func (o *World) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *World) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *World) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetCapacity

`func (o *World) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *World) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *World) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### GetCreatedAt

`func (o *World) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *World) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *World) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *World) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *World) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *World) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFavorites

`func (o *World) GetFavorites() int32`

GetFavorites returns the Favorites field if non-nil, zero value otherwise.

### GetFavoritesOk

`func (o *World) GetFavoritesOk() (*int32, bool)`

GetFavoritesOk returns a tuple with the Favorites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorites

`func (o *World) SetFavorites(v int32)`

SetFavorites sets Favorites field to given value.

### HasFavorites

`func (o *World) HasFavorites() bool`

HasFavorites returns a boolean if a field has been set.

### GetFeatured

`func (o *World) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *World) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *World) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.


### GetHeat

`func (o *World) GetHeat() int32`

GetHeat returns the Heat field if non-nil, zero value otherwise.

### GetHeatOk

`func (o *World) GetHeatOk() (*int32, bool)`

GetHeatOk returns a tuple with the Heat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeat

`func (o *World) SetHeat(v int32)`

SetHeat sets Heat field to given value.


### GetId

`func (o *World) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *World) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *World) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *World) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *World) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *World) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetInstances

`func (o *World) GetInstances() [][]interface{}`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *World) GetInstancesOk() (*[][]interface{}, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *World) SetInstances(v [][]interface{})`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *World) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetLabsPublicationDate

`func (o *World) GetLabsPublicationDate() string`

GetLabsPublicationDate returns the LabsPublicationDate field if non-nil, zero value otherwise.

### GetLabsPublicationDateOk

`func (o *World) GetLabsPublicationDateOk() (*string, bool)`

GetLabsPublicationDateOk returns a tuple with the LabsPublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabsPublicationDate

`func (o *World) SetLabsPublicationDate(v string)`

SetLabsPublicationDate sets LabsPublicationDate field to given value.


### GetName

`func (o *World) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *World) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *World) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *World) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *World) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *World) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetOccupants

`func (o *World) GetOccupants() int32`

GetOccupants returns the Occupants field if non-nil, zero value otherwise.

### GetOccupantsOk

`func (o *World) GetOccupantsOk() (*int32, bool)`

GetOccupantsOk returns a tuple with the Occupants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupants

`func (o *World) SetOccupants(v int32)`

SetOccupants sets Occupants field to given value.

### HasOccupants

`func (o *World) HasOccupants() bool`

HasOccupants returns a boolean if a field has been set.

### GetOrganization

`func (o *World) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *World) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *World) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetPopularity

`func (o *World) GetPopularity() int32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *World) GetPopularityOk() (*int32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *World) SetPopularity(v int32)`

SetPopularity sets Popularity field to given value.


### GetPreviewYoutubeId

`func (o *World) GetPreviewYoutubeId() string`

GetPreviewYoutubeId returns the PreviewYoutubeId field if non-nil, zero value otherwise.

### GetPreviewYoutubeIdOk

`func (o *World) GetPreviewYoutubeIdOk() (*string, bool)`

GetPreviewYoutubeIdOk returns a tuple with the PreviewYoutubeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewYoutubeId

`func (o *World) SetPreviewYoutubeId(v string)`

SetPreviewYoutubeId sets PreviewYoutubeId field to given value.

### HasPreviewYoutubeId

`func (o *World) HasPreviewYoutubeId() bool`

HasPreviewYoutubeId returns a boolean if a field has been set.

### SetPreviewYoutubeIdNil

`func (o *World) SetPreviewYoutubeIdNil(b bool)`

 SetPreviewYoutubeIdNil sets the value for PreviewYoutubeId to be an explicit nil

### UnsetPreviewYoutubeId
`func (o *World) UnsetPreviewYoutubeId()`

UnsetPreviewYoutubeId ensures that no value is present for PreviewYoutubeId, not even an explicit nil
### GetPrivateOccupants

`func (o *World) GetPrivateOccupants() int32`

GetPrivateOccupants returns the PrivateOccupants field if non-nil, zero value otherwise.

### GetPrivateOccupantsOk

`func (o *World) GetPrivateOccupantsOk() (*int32, bool)`

GetPrivateOccupantsOk returns a tuple with the PrivateOccupants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateOccupants

`func (o *World) SetPrivateOccupants(v int32)`

SetPrivateOccupants sets PrivateOccupants field to given value.

### HasPrivateOccupants

`func (o *World) HasPrivateOccupants() bool`

HasPrivateOccupants returns a boolean if a field has been set.

### GetPublicOccupants

`func (o *World) GetPublicOccupants() int32`

GetPublicOccupants returns the PublicOccupants field if non-nil, zero value otherwise.

### GetPublicOccupantsOk

`func (o *World) GetPublicOccupantsOk() (*int32, bool)`

GetPublicOccupantsOk returns a tuple with the PublicOccupants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicOccupants

`func (o *World) SetPublicOccupants(v int32)`

SetPublicOccupants sets PublicOccupants field to given value.

### HasPublicOccupants

`func (o *World) HasPublicOccupants() bool`

HasPublicOccupants returns a boolean if a field has been set.

### GetPublicationDate

`func (o *World) GetPublicationDate() string`

GetPublicationDate returns the PublicationDate field if non-nil, zero value otherwise.

### GetPublicationDateOk

`func (o *World) GetPublicationDateOk() (*string, bool)`

GetPublicationDateOk returns a tuple with the PublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationDate

`func (o *World) SetPublicationDate(v string)`

SetPublicationDate sets PublicationDate field to given value.


### GetReleaseStatus

`func (o *World) GetReleaseStatus() ReleaseStatus`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *World) GetReleaseStatusOk() (*ReleaseStatus, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *World) SetReleaseStatus(v ReleaseStatus)`

SetReleaseStatus sets ReleaseStatus field to given value.


### GetTags

`func (o *World) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *World) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *World) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetThumbnailImageUrl

`func (o *World) GetThumbnailImageUrl() string`

GetThumbnailImageUrl returns the ThumbnailImageUrl field if non-nil, zero value otherwise.

### GetThumbnailImageUrlOk

`func (o *World) GetThumbnailImageUrlOk() (*string, bool)`

GetThumbnailImageUrlOk returns a tuple with the ThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailImageUrl

`func (o *World) SetThumbnailImageUrl(v string)`

SetThumbnailImageUrl sets ThumbnailImageUrl field to given value.


### GetUnityPackages

`func (o *World) GetUnityPackages() []UnityPackage`

GetUnityPackages returns the UnityPackages field if non-nil, zero value otherwise.

### GetUnityPackagesOk

`func (o *World) GetUnityPackagesOk() (*[]UnityPackage, bool)`

GetUnityPackagesOk returns a tuple with the UnityPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackages

`func (o *World) SetUnityPackages(v []UnityPackage)`

SetUnityPackages sets UnityPackages field to given value.


### GetUpdatedAt

`func (o *World) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *World) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *World) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *World) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *World) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *World) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVisits

`func (o *World) GetVisits() int32`

GetVisits returns the Visits field if non-nil, zero value otherwise.

### GetVisitsOk

`func (o *World) GetVisitsOk() (*int32, bool)`

GetVisitsOk returns a tuple with the Visits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisits

`func (o *World) SetVisits(v int32)`

SetVisits sets Visits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


