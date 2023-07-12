# UnityPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetUrl** | Pointer to **string** |  | [optional] 
**AssetUrlObject** | Pointer to **map[string]interface{}** |  | [optional] 
**AssetVersion** | **int32** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | **string** |  | 
**Platform** | **string** | This can be &#x60;standalonewindows&#x60; or &#x60;android&#x60;, but can also pretty much be any random Unity verison such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | 
**PluginUrl** | Pointer to **string** |  | [optional] 
**PluginUrlObject** | Pointer to **map[string]interface{}** |  | [optional] 
**UnitySortNumber** | Pointer to **int64** |  | [optional] 
**UnityVersion** | **string** |  | [default to "5.3.4p1"]

## Methods

### NewUnityPackage

`func NewUnityPackage(assetVersion int32, id string, platform string, unityVersion string, ) *UnityPackage`

NewUnityPackage instantiates a new UnityPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnityPackageWithDefaults

`func NewUnityPackageWithDefaults() *UnityPackage`

NewUnityPackageWithDefaults instantiates a new UnityPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetUrl

`func (o *UnityPackage) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *UnityPackage) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *UnityPackage) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.

### HasAssetUrl

`func (o *UnityPackage) HasAssetUrl() bool`

HasAssetUrl returns a boolean if a field has been set.

### GetAssetUrlObject

`func (o *UnityPackage) GetAssetUrlObject() map[string]interface{}`

GetAssetUrlObject returns the AssetUrlObject field if non-nil, zero value otherwise.

### GetAssetUrlObjectOk

`func (o *UnityPackage) GetAssetUrlObjectOk() (*map[string]interface{}, bool)`

GetAssetUrlObjectOk returns a tuple with the AssetUrlObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrlObject

`func (o *UnityPackage) SetAssetUrlObject(v map[string]interface{})`

SetAssetUrlObject sets AssetUrlObject field to given value.

### HasAssetUrlObject

`func (o *UnityPackage) HasAssetUrlObject() bool`

HasAssetUrlObject returns a boolean if a field has been set.

### GetAssetVersion

`func (o *UnityPackage) GetAssetVersion() int32`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *UnityPackage) GetAssetVersionOk() (*int32, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *UnityPackage) SetAssetVersion(v int32)`

SetAssetVersion sets AssetVersion field to given value.


### GetCreatedAt

`func (o *UnityPackage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UnityPackage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UnityPackage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UnityPackage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *UnityPackage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnityPackage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnityPackage) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *UnityPackage) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UnityPackage) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UnityPackage) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPluginUrl

`func (o *UnityPackage) GetPluginUrl() string`

GetPluginUrl returns the PluginUrl field if non-nil, zero value otherwise.

### GetPluginUrlOk

`func (o *UnityPackage) GetPluginUrlOk() (*string, bool)`

GetPluginUrlOk returns a tuple with the PluginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginUrl

`func (o *UnityPackage) SetPluginUrl(v string)`

SetPluginUrl sets PluginUrl field to given value.

### HasPluginUrl

`func (o *UnityPackage) HasPluginUrl() bool`

HasPluginUrl returns a boolean if a field has been set.

### GetPluginUrlObject

`func (o *UnityPackage) GetPluginUrlObject() map[string]interface{}`

GetPluginUrlObject returns the PluginUrlObject field if non-nil, zero value otherwise.

### GetPluginUrlObjectOk

`func (o *UnityPackage) GetPluginUrlObjectOk() (*map[string]interface{}, bool)`

GetPluginUrlObjectOk returns a tuple with the PluginUrlObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginUrlObject

`func (o *UnityPackage) SetPluginUrlObject(v map[string]interface{})`

SetPluginUrlObject sets PluginUrlObject field to given value.

### HasPluginUrlObject

`func (o *UnityPackage) HasPluginUrlObject() bool`

HasPluginUrlObject returns a boolean if a field has been set.

### GetUnitySortNumber

`func (o *UnityPackage) GetUnitySortNumber() int64`

GetUnitySortNumber returns the UnitySortNumber field if non-nil, zero value otherwise.

### GetUnitySortNumberOk

`func (o *UnityPackage) GetUnitySortNumberOk() (*int64, bool)`

GetUnitySortNumberOk returns a tuple with the UnitySortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitySortNumber

`func (o *UnityPackage) SetUnitySortNumber(v int64)`

SetUnitySortNumber sets UnitySortNumber field to given value.

### HasUnitySortNumber

`func (o *UnityPackage) HasUnitySortNumber() bool`

HasUnitySortNumber returns a boolean if a field has been set.

### GetUnityVersion

`func (o *UnityPackage) GetUnityVersion() string`

GetUnityVersion returns the UnityVersion field if non-nil, zero value otherwise.

### GetUnityVersionOk

`func (o *UnityPackage) GetUnityVersionOk() (*string, bool)`

GetUnityVersionOk returns a tuple with the UnityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityVersion

`func (o *UnityPackage) SetUnityVersion(v string)`

SetUnityVersion sets UnityVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


