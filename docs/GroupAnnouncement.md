# GroupAnnouncement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**AuthorId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Text** | Pointer to **NullableString** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewGroupAnnouncement

`func NewGroupAnnouncement() *GroupAnnouncement`

NewGroupAnnouncement instantiates a new GroupAnnouncement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupAnnouncementWithDefaults

`func NewGroupAnnouncementWithDefaults() *GroupAnnouncement`

NewGroupAnnouncementWithDefaults instantiates a new GroupAnnouncement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupAnnouncement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupAnnouncement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupAnnouncement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupAnnouncement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *GroupAnnouncement) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupAnnouncement) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupAnnouncement) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupAnnouncement) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAuthorId

`func (o *GroupAnnouncement) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *GroupAnnouncement) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *GroupAnnouncement) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *GroupAnnouncement) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetTitle

`func (o *GroupAnnouncement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GroupAnnouncement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GroupAnnouncement) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GroupAnnouncement) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *GroupAnnouncement) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *GroupAnnouncement) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetText

`func (o *GroupAnnouncement) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GroupAnnouncement) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GroupAnnouncement) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *GroupAnnouncement) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *GroupAnnouncement) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *GroupAnnouncement) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetImageId

`func (o *GroupAnnouncement) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *GroupAnnouncement) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *GroupAnnouncement) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *GroupAnnouncement) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetImageUrl

`func (o *GroupAnnouncement) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GroupAnnouncement) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GroupAnnouncement) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GroupAnnouncement) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *GroupAnnouncement) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *GroupAnnouncement) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetCreatedAt

`func (o *GroupAnnouncement) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupAnnouncement) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupAnnouncement) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupAnnouncement) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GroupAnnouncement) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GroupAnnouncement) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GroupAnnouncement) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GroupAnnouncement) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GroupAnnouncement) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GroupAnnouncement) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GroupAnnouncement) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GroupAnnouncement) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


