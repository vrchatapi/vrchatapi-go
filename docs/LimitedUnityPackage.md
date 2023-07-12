# LimitedUnityPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** | This can be &#x60;standalonewindows&#x60; or &#x60;android&#x60;, but can also pretty much be any random Unity verison such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | 
**UnityVersion** | **string** |  | 

## Methods

### NewLimitedUnityPackage

`func NewLimitedUnityPackage(platform string, unityVersion string, ) *LimitedUnityPackage`

NewLimitedUnityPackage instantiates a new LimitedUnityPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitedUnityPackageWithDefaults

`func NewLimitedUnityPackageWithDefaults() *LimitedUnityPackage`

NewLimitedUnityPackageWithDefaults instantiates a new LimitedUnityPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *LimitedUnityPackage) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *LimitedUnityPackage) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *LimitedUnityPackage) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetUnityVersion

`func (o *LimitedUnityPackage) GetUnityVersion() string`

GetUnityVersion returns the UnityVersion field if non-nil, zero value otherwise.

### GetUnityVersionOk

`func (o *LimitedUnityPackage) GetUnityVersionOk() (*string, bool)`

GetUnityVersionOk returns a tuple with the UnityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityVersion

`func (o *LimitedUnityPackage) SetUnityVersion(v string)`

SetUnityVersion sets UnityVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


