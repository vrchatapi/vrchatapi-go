# AddFavoriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FavoriteType**](FavoriteType.md) |  | [default to FRIEND]
**FavoriteId** | **string** | Must be either AvatarID, WorldID or UserID. | 
**Tags** | **[]string** | Tags indicate which group this favorite belongs to. Adding multiple groups makes it show up in all. Removing it from one in that case removes it from all. | 

## Methods

### NewAddFavoriteRequest

`func NewAddFavoriteRequest(type_ FavoriteType, favoriteId string, tags []string, ) *AddFavoriteRequest`

NewAddFavoriteRequest instantiates a new AddFavoriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFavoriteRequestWithDefaults

`func NewAddFavoriteRequestWithDefaults() *AddFavoriteRequest`

NewAddFavoriteRequestWithDefaults instantiates a new AddFavoriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddFavoriteRequest) GetType() FavoriteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddFavoriteRequest) GetTypeOk() (*FavoriteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddFavoriteRequest) SetType(v FavoriteType)`

SetType sets Type field to given value.


### GetFavoriteId

`func (o *AddFavoriteRequest) GetFavoriteId() string`

GetFavoriteId returns the FavoriteId field if non-nil, zero value otherwise.

### GetFavoriteIdOk

`func (o *AddFavoriteRequest) GetFavoriteIdOk() (*string, bool)`

GetFavoriteIdOk returns a tuple with the FavoriteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteId

`func (o *AddFavoriteRequest) SetFavoriteId(v string)`

SetFavoriteId sets FavoriteId field to given value.


### GetTags

`func (o *AddFavoriteRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddFavoriteRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddFavoriteRequest) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


