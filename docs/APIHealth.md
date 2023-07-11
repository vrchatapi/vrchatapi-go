# APIHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | 
**ServerName** | **string** |  | 
**BuildVersionTag** | **string** |  | 

## Methods

### NewAPIHealth

`func NewAPIHealth(ok bool, serverName string, buildVersionTag string, ) *APIHealth`

NewAPIHealth instantiates a new APIHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIHealthWithDefaults

`func NewAPIHealthWithDefaults() *APIHealth`

NewAPIHealthWithDefaults instantiates a new APIHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *APIHealth) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *APIHealth) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *APIHealth) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetServerName

`func (o *APIHealth) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *APIHealth) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *APIHealth) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetBuildVersionTag

`func (o *APIHealth) GetBuildVersionTag() string`

GetBuildVersionTag returns the BuildVersionTag field if non-nil, zero value otherwise.

### GetBuildVersionTagOk

`func (o *APIHealth) GetBuildVersionTagOk() (*string, bool)`

GetBuildVersionTagOk returns a tuple with the BuildVersionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersionTag

`func (o *APIHealth) SetBuildVersionTag(v string)`

SetBuildVersionTag sets BuildVersionTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


