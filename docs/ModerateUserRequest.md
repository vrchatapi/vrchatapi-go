# ModerateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Moderated** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**Type** | [**PlayerModerationType**](PlayerModerationType.md) |  | [default to UNMUTE]

## Methods

### NewModerateUserRequest

`func NewModerateUserRequest(moderated string, type_ PlayerModerationType, ) *ModerateUserRequest`

NewModerateUserRequest instantiates a new ModerateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerateUserRequestWithDefaults

`func NewModerateUserRequestWithDefaults() *ModerateUserRequest`

NewModerateUserRequestWithDefaults instantiates a new ModerateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModerated

`func (o *ModerateUserRequest) GetModerated() string`

GetModerated returns the Moderated field if non-nil, zero value otherwise.

### GetModeratedOk

`func (o *ModerateUserRequest) GetModeratedOk() (*string, bool)`

GetModeratedOk returns a tuple with the Moderated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerated

`func (o *ModerateUserRequest) SetModerated(v string)`

SetModerated sets Moderated field to given value.


### GetType

`func (o *ModerateUserRequest) GetType() PlayerModerationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModerateUserRequest) GetTypeOk() (*PlayerModerationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModerateUserRequest) SetType(v PlayerModerationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


