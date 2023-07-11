# GroupGalleryImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**GalleryId** | Pointer to **string** |  | [optional] 
**FileId** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**SubmittedByUserId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**Approved** | Pointer to **bool** |  | [optional] [default to false]
**ApprovedByUserId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**ApprovedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGroupGalleryImage

`func NewGroupGalleryImage() *GroupGalleryImage`

NewGroupGalleryImage instantiates a new GroupGalleryImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupGalleryImageWithDefaults

`func NewGroupGalleryImageWithDefaults() *GroupGalleryImage`

NewGroupGalleryImageWithDefaults instantiates a new GroupGalleryImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupGalleryImage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupGalleryImage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupGalleryImage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupGalleryImage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *GroupGalleryImage) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupGalleryImage) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupGalleryImage) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupGalleryImage) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGalleryId

`func (o *GroupGalleryImage) GetGalleryId() string`

GetGalleryId returns the GalleryId field if non-nil, zero value otherwise.

### GetGalleryIdOk

`func (o *GroupGalleryImage) GetGalleryIdOk() (*string, bool)`

GetGalleryIdOk returns a tuple with the GalleryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGalleryId

`func (o *GroupGalleryImage) SetGalleryId(v string)`

SetGalleryId sets GalleryId field to given value.

### HasGalleryId

`func (o *GroupGalleryImage) HasGalleryId() bool`

HasGalleryId returns a boolean if a field has been set.

### GetFileId

`func (o *GroupGalleryImage) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *GroupGalleryImage) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *GroupGalleryImage) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *GroupGalleryImage) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetImageUrl

`func (o *GroupGalleryImage) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GroupGalleryImage) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GroupGalleryImage) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GroupGalleryImage) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GroupGalleryImage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupGalleryImage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupGalleryImage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupGalleryImage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSubmittedByUserId

`func (o *GroupGalleryImage) GetSubmittedByUserId() string`

GetSubmittedByUserId returns the SubmittedByUserId field if non-nil, zero value otherwise.

### GetSubmittedByUserIdOk

`func (o *GroupGalleryImage) GetSubmittedByUserIdOk() (*string, bool)`

GetSubmittedByUserIdOk returns a tuple with the SubmittedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedByUserId

`func (o *GroupGalleryImage) SetSubmittedByUserId(v string)`

SetSubmittedByUserId sets SubmittedByUserId field to given value.

### HasSubmittedByUserId

`func (o *GroupGalleryImage) HasSubmittedByUserId() bool`

HasSubmittedByUserId returns a boolean if a field has been set.

### GetApproved

`func (o *GroupGalleryImage) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *GroupGalleryImage) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *GroupGalleryImage) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *GroupGalleryImage) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetApprovedByUserId

`func (o *GroupGalleryImage) GetApprovedByUserId() string`

GetApprovedByUserId returns the ApprovedByUserId field if non-nil, zero value otherwise.

### GetApprovedByUserIdOk

`func (o *GroupGalleryImage) GetApprovedByUserIdOk() (*string, bool)`

GetApprovedByUserIdOk returns a tuple with the ApprovedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedByUserId

`func (o *GroupGalleryImage) SetApprovedByUserId(v string)`

SetApprovedByUserId sets ApprovedByUserId field to given value.

### HasApprovedByUserId

`func (o *GroupGalleryImage) HasApprovedByUserId() bool`

HasApprovedByUserId returns a boolean if a field has been set.

### GetApprovedAt

`func (o *GroupGalleryImage) GetApprovedAt() time.Time`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *GroupGalleryImage) GetApprovedAtOk() (*time.Time, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *GroupGalleryImage) SetApprovedAt(v time.Time)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *GroupGalleryImage) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


