# FileVersionUploadStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadId** | **string** |  | 
**FileName** | **string** |  | 
**NextPartNumber** | **float32** |  | 
**MaxParts** | **float32** |  | 
**Parts** | **[]map[string]interface{}** |  | 
**Etags** | **[]map[string]interface{}** | Unknown | 

## Methods

### NewFileVersionUploadStatus

`func NewFileVersionUploadStatus(uploadId string, fileName string, nextPartNumber float32, maxParts float32, parts []map[string]interface{}, etags []map[string]interface{}, ) *FileVersionUploadStatus`

NewFileVersionUploadStatus instantiates a new FileVersionUploadStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileVersionUploadStatusWithDefaults

`func NewFileVersionUploadStatusWithDefaults() *FileVersionUploadStatus`

NewFileVersionUploadStatusWithDefaults instantiates a new FileVersionUploadStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadId

`func (o *FileVersionUploadStatus) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *FileVersionUploadStatus) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *FileVersionUploadStatus) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.


### GetFileName

`func (o *FileVersionUploadStatus) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FileVersionUploadStatus) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FileVersionUploadStatus) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetNextPartNumber

`func (o *FileVersionUploadStatus) GetNextPartNumber() float32`

GetNextPartNumber returns the NextPartNumber field if non-nil, zero value otherwise.

### GetNextPartNumberOk

`func (o *FileVersionUploadStatus) GetNextPartNumberOk() (*float32, bool)`

GetNextPartNumberOk returns a tuple with the NextPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPartNumber

`func (o *FileVersionUploadStatus) SetNextPartNumber(v float32)`

SetNextPartNumber sets NextPartNumber field to given value.


### GetMaxParts

`func (o *FileVersionUploadStatus) GetMaxParts() float32`

GetMaxParts returns the MaxParts field if non-nil, zero value otherwise.

### GetMaxPartsOk

`func (o *FileVersionUploadStatus) GetMaxPartsOk() (*float32, bool)`

GetMaxPartsOk returns a tuple with the MaxParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParts

`func (o *FileVersionUploadStatus) SetMaxParts(v float32)`

SetMaxParts sets MaxParts field to given value.


### GetParts

`func (o *FileVersionUploadStatus) GetParts() []map[string]interface{}`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *FileVersionUploadStatus) GetPartsOk() (*[]map[string]interface{}, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *FileVersionUploadStatus) SetParts(v []map[string]interface{})`

SetParts sets Parts field to given value.


### GetEtags

`func (o *FileVersionUploadStatus) GetEtags() []map[string]interface{}`

GetEtags returns the Etags field if non-nil, zero value otherwise.

### GetEtagsOk

`func (o *FileVersionUploadStatus) GetEtagsOk() (*[]map[string]interface{}, bool)`

GetEtagsOk returns a tuple with the Etags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtags

`func (o *FileVersionUploadStatus) SetEtags(v []map[string]interface{})`

SetEtags sets Etags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


