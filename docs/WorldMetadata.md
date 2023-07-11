# WorldMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**Metadata** | **map[string]interface{}** |  | 

## Methods

### NewWorldMetadata

`func NewWorldMetadata(id string, metadata map[string]interface{}, ) *WorldMetadata`

NewWorldMetadata instantiates a new WorldMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorldMetadataWithDefaults

`func NewWorldMetadataWithDefaults() *WorldMetadata`

NewWorldMetadataWithDefaults instantiates a new WorldMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorldMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorldMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorldMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *WorldMetadata) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WorldMetadata) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WorldMetadata) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


