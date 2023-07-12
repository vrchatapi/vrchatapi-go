# CreateFileVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignatureMd5** | **string** |  | 
**SignatureSizeInBytes** | **float32** |  | 
**FileMd5** | Pointer to **string** |  | [optional] 
**FileSizeInBytes** | Pointer to **float32** |  | [optional] 

## Methods

### NewCreateFileVersionRequest

`func NewCreateFileVersionRequest(signatureMd5 string, signatureSizeInBytes float32, ) *CreateFileVersionRequest`

NewCreateFileVersionRequest instantiates a new CreateFileVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFileVersionRequestWithDefaults

`func NewCreateFileVersionRequestWithDefaults() *CreateFileVersionRequest`

NewCreateFileVersionRequestWithDefaults instantiates a new CreateFileVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignatureMd5

`func (o *CreateFileVersionRequest) GetSignatureMd5() string`

GetSignatureMd5 returns the SignatureMd5 field if non-nil, zero value otherwise.

### GetSignatureMd5Ok

`func (o *CreateFileVersionRequest) GetSignatureMd5Ok() (*string, bool)`

GetSignatureMd5Ok returns a tuple with the SignatureMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureMd5

`func (o *CreateFileVersionRequest) SetSignatureMd5(v string)`

SetSignatureMd5 sets SignatureMd5 field to given value.


### GetSignatureSizeInBytes

`func (o *CreateFileVersionRequest) GetSignatureSizeInBytes() float32`

GetSignatureSizeInBytes returns the SignatureSizeInBytes field if non-nil, zero value otherwise.

### GetSignatureSizeInBytesOk

`func (o *CreateFileVersionRequest) GetSignatureSizeInBytesOk() (*float32, bool)`

GetSignatureSizeInBytesOk returns a tuple with the SignatureSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureSizeInBytes

`func (o *CreateFileVersionRequest) SetSignatureSizeInBytes(v float32)`

SetSignatureSizeInBytes sets SignatureSizeInBytes field to given value.


### GetFileMd5

`func (o *CreateFileVersionRequest) GetFileMd5() string`

GetFileMd5 returns the FileMd5 field if non-nil, zero value otherwise.

### GetFileMd5Ok

`func (o *CreateFileVersionRequest) GetFileMd5Ok() (*string, bool)`

GetFileMd5Ok returns a tuple with the FileMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMd5

`func (o *CreateFileVersionRequest) SetFileMd5(v string)`

SetFileMd5 sets FileMd5 field to given value.

### HasFileMd5

`func (o *CreateFileVersionRequest) HasFileMd5() bool`

HasFileMd5 returns a boolean if a field has been set.

### GetFileSizeInBytes

`func (o *CreateFileVersionRequest) GetFileSizeInBytes() float32`

GetFileSizeInBytes returns the FileSizeInBytes field if non-nil, zero value otherwise.

### GetFileSizeInBytesOk

`func (o *CreateFileVersionRequest) GetFileSizeInBytesOk() (*float32, bool)`

GetFileSizeInBytesOk returns a tuple with the FileSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeInBytes

`func (o *CreateFileVersionRequest) SetFileSizeInBytes(v float32)`

SetFileSizeInBytes sets FileSizeInBytes field to given value.

### HasFileSizeInBytes

`func (o *CreateFileVersionRequest) HasFileSizeInBytes() bool`

HasFileSizeInBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


