# FileVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Deleted** | Pointer to **bool** | Usually only present if &#x60;true&#x60; | [optional] [default to true]
**Delta** | Pointer to [**FileData**](FileData.md) |  | [optional] 
**File** | Pointer to [**FileData**](FileData.md) |  | [optional] 
**Signature** | Pointer to [**FileData**](FileData.md) |  | [optional] 
**Status** | [**FileStatus**](FileStatus.md) |  | [default to WAITING]
**Version** | **int32** | Incremental version counter, can only be increased. | [default to 0]

## Methods

### NewFileVersion

`func NewFileVersion(createdAt time.Time, status FileStatus, version int32, ) *FileVersion`

NewFileVersion instantiates a new FileVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileVersionWithDefaults

`func NewFileVersionWithDefaults() *FileVersion`

NewFileVersionWithDefaults instantiates a new FileVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FileVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeleted

`func (o *FileVersion) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *FileVersion) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *FileVersion) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *FileVersion) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDelta

`func (o *FileVersion) GetDelta() FileData`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *FileVersion) GetDeltaOk() (*FileData, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *FileVersion) SetDelta(v FileData)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *FileVersion) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetFile

`func (o *FileVersion) GetFile() FileData`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *FileVersion) GetFileOk() (*FileData, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *FileVersion) SetFile(v FileData)`

SetFile sets File field to given value.

### HasFile

`func (o *FileVersion) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetSignature

`func (o *FileVersion) GetSignature() FileData`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *FileVersion) GetSignatureOk() (*FileData, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *FileVersion) SetSignature(v FileData)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *FileVersion) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetStatus

`func (o *FileVersion) GetStatus() FileStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FileVersion) GetStatusOk() (*FileStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FileVersion) SetStatus(v FileStatus)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *FileVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FileVersion) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FileVersion) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


