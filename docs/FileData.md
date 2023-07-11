# FileData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** |  | [default to "queued"]
**FileName** | **string** |  | 
**Md5** | **string** |  | 
**SizeInBytes** | **int32** |  | 
**Status** | [**FileStatus**](FileStatus.md) |  | [default to WAITING]
**UploadId** | **string** |  | [default to ""]
**Url** | **string** |  | 

## Methods

### NewFileData

`func NewFileData(category string, fileName string, md5 string, sizeInBytes int32, status FileStatus, uploadId string, url string, ) *FileData`

NewFileData instantiates a new FileData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileDataWithDefaults

`func NewFileDataWithDefaults() *FileData`

NewFileDataWithDefaults instantiates a new FileData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *FileData) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FileData) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FileData) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetFileName

`func (o *FileData) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FileData) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FileData) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetMd5

`func (o *FileData) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *FileData) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *FileData) SetMd5(v string)`

SetMd5 sets Md5 field to given value.


### GetSizeInBytes

`func (o *FileData) GetSizeInBytes() int32`

GetSizeInBytes returns the SizeInBytes field if non-nil, zero value otherwise.

### GetSizeInBytesOk

`func (o *FileData) GetSizeInBytesOk() (*int32, bool)`

GetSizeInBytesOk returns a tuple with the SizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInBytes

`func (o *FileData) SetSizeInBytes(v int32)`

SetSizeInBytes sets SizeInBytes field to given value.


### GetStatus

`func (o *FileData) GetStatus() FileStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FileData) GetStatusOk() (*FileStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FileData) SetStatus(v FileStatus)`

SetStatus sets Status field to given value.


### GetUploadId

`func (o *FileData) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *FileData) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *FileData) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.


### GetUrl

`func (o *FileData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileData) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


