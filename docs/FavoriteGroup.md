# FavoriteGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**OwnerDisplayName** | **string** |  | 
**OwnerId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**Tags** | **[]string** |   | 
**Type** | [**FavoriteType**](FavoriteType.md) |  | [default to FRIEND]
**Visibility** | [**FavoriteGroupVisibility**](FavoriteGroupVisibility.md) |  | [default to PRIVATE]

## Methods

### NewFavoriteGroup

`func NewFavoriteGroup(displayName string, id string, name string, ownerDisplayName string, ownerId string, tags []string, type_ FavoriteType, visibility FavoriteGroupVisibility, ) *FavoriteGroup`

NewFavoriteGroup instantiates a new FavoriteGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFavoriteGroupWithDefaults

`func NewFavoriteGroupWithDefaults() *FavoriteGroup`

NewFavoriteGroupWithDefaults instantiates a new FavoriteGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *FavoriteGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FavoriteGroup) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FavoriteGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetId

`func (o *FavoriteGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FavoriteGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FavoriteGroup) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FavoriteGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FavoriteGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FavoriteGroup) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerDisplayName

`func (o *FavoriteGroup) GetOwnerDisplayName() string`

GetOwnerDisplayName returns the OwnerDisplayName field if non-nil, zero value otherwise.

### GetOwnerDisplayNameOk

`func (o *FavoriteGroup) GetOwnerDisplayNameOk() (*string, bool)`

GetOwnerDisplayNameOk returns a tuple with the OwnerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerDisplayName

`func (o *FavoriteGroup) SetOwnerDisplayName(v string)`

SetOwnerDisplayName sets OwnerDisplayName field to given value.


### GetOwnerId

`func (o *FavoriteGroup) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *FavoriteGroup) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *FavoriteGroup) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetTags

`func (o *FavoriteGroup) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FavoriteGroup) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FavoriteGroup) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetType

`func (o *FavoriteGroup) GetType() FavoriteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FavoriteGroup) GetTypeOk() (*FavoriteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FavoriteGroup) SetType(v FavoriteType)`

SetType sets Type field to given value.


### GetVisibility

`func (o *FavoriteGroup) GetVisibility() FavoriteGroupVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FavoriteGroup) GetVisibilityOk() (*FavoriteGroupVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FavoriteGroup) SetVisibility(v FavoriteGroupVisibility)`

SetVisibility sets Visibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


