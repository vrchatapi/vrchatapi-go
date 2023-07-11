# InfoPush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**IsEnabled** | **bool** |  | [default to true]
**ReleaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) |  | [default to PUBLIC]
**Priority** | **int32** |  | 
**Tags** | **[]string** |   | 
**Data** | [**InfoPushData**](InfoPushData.md) |  | 
**Hash** | **string** | Unknown usage, MD5 | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInfoPush

`func NewInfoPush(id string, isEnabled bool, releaseStatus ReleaseStatus, priority int32, tags []string, data InfoPushData, hash string, createdAt time.Time, updatedAt time.Time, ) *InfoPush`

NewInfoPush instantiates a new InfoPush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoPushWithDefaults

`func NewInfoPushWithDefaults() *InfoPush`

NewInfoPushWithDefaults instantiates a new InfoPush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InfoPush) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InfoPush) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InfoPush) SetId(v string)`

SetId sets Id field to given value.


### GetIsEnabled

`func (o *InfoPush) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *InfoPush) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *InfoPush) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetReleaseStatus

`func (o *InfoPush) GetReleaseStatus() ReleaseStatus`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *InfoPush) GetReleaseStatusOk() (*ReleaseStatus, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *InfoPush) SetReleaseStatus(v ReleaseStatus)`

SetReleaseStatus sets ReleaseStatus field to given value.


### GetPriority

`func (o *InfoPush) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InfoPush) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InfoPush) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetTags

`func (o *InfoPush) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InfoPush) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InfoPush) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetData

`func (o *InfoPush) GetData() InfoPushData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InfoPush) GetDataOk() (*InfoPushData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InfoPush) SetData(v InfoPushData)`

SetData sets Data field to given value.


### GetHash

`func (o *InfoPush) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *InfoPush) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *InfoPush) SetHash(v string)`

SetHash sets Hash field to given value.


### GetCreatedAt

`func (o *InfoPush) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InfoPush) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InfoPush) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InfoPush) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InfoPush) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InfoPush) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStartDate

`func (o *InfoPush) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InfoPush) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InfoPush) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InfoPush) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InfoPush) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InfoPush) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InfoPush) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InfoPush) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


