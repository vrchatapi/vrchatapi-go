# FinishFileDataUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Etags** | Pointer to **[]string** | Array of ETags uploaded. | [optional] 
**NextPartNumber** | **string** | Always a zero in string form, despite how many parts uploaded. | [default to "0"]
**MaxParts** | **string** | Always a zero in string form, despite how many parts uploaded. | [default to "0"]

## Methods

### NewFinishFileDataUploadRequest

`func NewFinishFileDataUploadRequest(nextPartNumber string, maxParts string, ) *FinishFileDataUploadRequest`

NewFinishFileDataUploadRequest instantiates a new FinishFileDataUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinishFileDataUploadRequestWithDefaults

`func NewFinishFileDataUploadRequestWithDefaults() *FinishFileDataUploadRequest`

NewFinishFileDataUploadRequestWithDefaults instantiates a new FinishFileDataUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEtags

`func (o *FinishFileDataUploadRequest) GetEtags() []string`

GetEtags returns the Etags field if non-nil, zero value otherwise.

### GetEtagsOk

`func (o *FinishFileDataUploadRequest) GetEtagsOk() (*[]string, bool)`

GetEtagsOk returns a tuple with the Etags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtags

`func (o *FinishFileDataUploadRequest) SetEtags(v []string)`

SetEtags sets Etags field to given value.

### HasEtags

`func (o *FinishFileDataUploadRequest) HasEtags() bool`

HasEtags returns a boolean if a field has been set.

### GetNextPartNumber

`func (o *FinishFileDataUploadRequest) GetNextPartNumber() string`

GetNextPartNumber returns the NextPartNumber field if non-nil, zero value otherwise.

### GetNextPartNumberOk

`func (o *FinishFileDataUploadRequest) GetNextPartNumberOk() (*string, bool)`

GetNextPartNumberOk returns a tuple with the NextPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPartNumber

`func (o *FinishFileDataUploadRequest) SetNextPartNumber(v string)`

SetNextPartNumber sets NextPartNumber field to given value.


### GetMaxParts

`func (o *FinishFileDataUploadRequest) GetMaxParts() string`

GetMaxParts returns the MaxParts field if non-nil, zero value otherwise.

### GetMaxPartsOk

`func (o *FinishFileDataUploadRequest) GetMaxPartsOk() (*string, bool)`

GetMaxPartsOk returns a tuple with the MaxParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParts

`func (o *FinishFileDataUploadRequest) SetMaxParts(v string)`

SetMaxParts sets MaxParts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


