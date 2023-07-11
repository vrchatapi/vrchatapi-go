# Favorite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FavoriteId** | **string** | MUST be either AvatarID, UserID or WorldID. | 
**Id** | **string** |  | 
**Tags** | **[]string** |   | 
**Type** | [**FavoriteType**](FavoriteType.md) |  | [default to FRIEND]

## Methods

### NewFavorite

`func NewFavorite(favoriteId string, id string, tags []string, type_ FavoriteType, ) *Favorite`

NewFavorite instantiates a new Favorite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFavoriteWithDefaults

`func NewFavoriteWithDefaults() *Favorite`

NewFavoriteWithDefaults instantiates a new Favorite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFavoriteId

`func (o *Favorite) GetFavoriteId() string`

GetFavoriteId returns the FavoriteId field if non-nil, zero value otherwise.

### GetFavoriteIdOk

`func (o *Favorite) GetFavoriteIdOk() (*string, bool)`

GetFavoriteIdOk returns a tuple with the FavoriteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteId

`func (o *Favorite) SetFavoriteId(v string)`

SetFavoriteId sets FavoriteId field to given value.


### GetId

`func (o *Favorite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Favorite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Favorite) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *Favorite) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Favorite) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Favorite) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetType

`func (o *Favorite) GetType() FavoriteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Favorite) GetTypeOk() (*FavoriteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Favorite) SetType(v FavoriteType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


