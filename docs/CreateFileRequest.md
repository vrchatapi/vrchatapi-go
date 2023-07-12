# CreateFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**MimeType** | [**MIMEType**](MIMEType.md) |  | [default to IMAGE_JPEG]
**Extension** | **string** |  | 
**Tags** | Pointer to **[]string** |   | [optional] 

## Methods

### NewCreateFileRequest

`func NewCreateFileRequest(name string, mimeType MIMEType, extension string, ) *CreateFileRequest`

NewCreateFileRequest instantiates a new CreateFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFileRequestWithDefaults

`func NewCreateFileRequestWithDefaults() *CreateFileRequest`

NewCreateFileRequestWithDefaults instantiates a new CreateFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMimeType

`func (o *CreateFileRequest) GetMimeType() MIMEType`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *CreateFileRequest) GetMimeTypeOk() (*MIMEType, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *CreateFileRequest) SetMimeType(v MIMEType)`

SetMimeType sets MimeType field to given value.


### GetExtension

`func (o *CreateFileRequest) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CreateFileRequest) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CreateFileRequest) SetExtension(v string)`

SetExtension sets Extension field to given value.


### GetTags

`func (o *CreateFileRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateFileRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateFileRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateFileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


