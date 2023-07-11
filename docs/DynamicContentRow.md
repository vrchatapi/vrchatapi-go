# DynamicContentRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Platform** | **string** | Usually \&quot;ThisPlatformSupported\&quot;, but can also be other values such as \&quot;all\&quot; or platform specific identifiers. | 
**SortHeading** | **string** |  | 
**SortOrder** | **string** |  | 
**SortOwnership** | **string** |  | 
**Tag** | Pointer to **string** | Tag to filter content for this row. | [optional] 
**Type** | Pointer to **string** | Type is not present if it is a world. | [optional] 

## Methods

### NewDynamicContentRow

`func NewDynamicContentRow(name string, platform string, sortHeading string, sortOrder string, sortOwnership string, ) *DynamicContentRow`

NewDynamicContentRow instantiates a new DynamicContentRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicContentRowWithDefaults

`func NewDynamicContentRowWithDefaults() *DynamicContentRow`

NewDynamicContentRowWithDefaults instantiates a new DynamicContentRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *DynamicContentRow) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *DynamicContentRow) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *DynamicContentRow) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *DynamicContentRow) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *DynamicContentRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicContentRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicContentRow) SetName(v string)`

SetName sets Name field to given value.


### GetPlatform

`func (o *DynamicContentRow) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DynamicContentRow) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DynamicContentRow) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetSortHeading

`func (o *DynamicContentRow) GetSortHeading() string`

GetSortHeading returns the SortHeading field if non-nil, zero value otherwise.

### GetSortHeadingOk

`func (o *DynamicContentRow) GetSortHeadingOk() (*string, bool)`

GetSortHeadingOk returns a tuple with the SortHeading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortHeading

`func (o *DynamicContentRow) SetSortHeading(v string)`

SetSortHeading sets SortHeading field to given value.


### GetSortOrder

`func (o *DynamicContentRow) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *DynamicContentRow) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *DynamicContentRow) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.


### GetSortOwnership

`func (o *DynamicContentRow) GetSortOwnership() string`

GetSortOwnership returns the SortOwnership field if non-nil, zero value otherwise.

### GetSortOwnershipOk

`func (o *DynamicContentRow) GetSortOwnershipOk() (*string, bool)`

GetSortOwnershipOk returns a tuple with the SortOwnership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOwnership

`func (o *DynamicContentRow) SetSortOwnership(v string)`

SetSortOwnership sets SortOwnership field to given value.


### GetTag

`func (o *DynamicContentRow) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DynamicContentRow) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DynamicContentRow) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DynamicContentRow) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetType

`func (o *DynamicContentRow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DynamicContentRow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DynamicContentRow) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DynamicContentRow) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


