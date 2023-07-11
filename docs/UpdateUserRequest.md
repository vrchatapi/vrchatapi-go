# UpdateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Birthday** | Pointer to **string** |  | [optional] 
**AcceptedTOSVersion** | Pointer to **float32** |  | [optional] 
**Tags** | Pointer to **[]string** |   | [optional] 
**Status** | Pointer to [**UserStatus**](UserStatus.md) |  | [optional] [default to OFFLINE]
**StatusDescription** | Pointer to **string** |  | [optional] 
**Bio** | Pointer to **string** |  | [optional] 
**BioLinks** | Pointer to **[]string** |  | [optional] 
**UserIcon** | Pointer to **string** | MUST be a valid VRChat /file/ url. | [optional] 

## Methods

### NewUpdateUserRequest

`func NewUpdateUserRequest() *UpdateUserRequest`

NewUpdateUserRequest instantiates a new UpdateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequestWithDefaults

`func NewUpdateUserRequestWithDefaults() *UpdateUserRequest`

NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UpdateUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetBirthday

`func (o *UpdateUserRequest) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *UpdateUserRequest) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *UpdateUserRequest) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *UpdateUserRequest) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetAcceptedTOSVersion

`func (o *UpdateUserRequest) GetAcceptedTOSVersion() float32`

GetAcceptedTOSVersion returns the AcceptedTOSVersion field if non-nil, zero value otherwise.

### GetAcceptedTOSVersionOk

`func (o *UpdateUserRequest) GetAcceptedTOSVersionOk() (*float32, bool)`

GetAcceptedTOSVersionOk returns a tuple with the AcceptedTOSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTOSVersion

`func (o *UpdateUserRequest) SetAcceptedTOSVersion(v float32)`

SetAcceptedTOSVersion sets AcceptedTOSVersion field to given value.

### HasAcceptedTOSVersion

`func (o *UpdateUserRequest) HasAcceptedTOSVersion() bool`

HasAcceptedTOSVersion returns a boolean if a field has been set.

### GetTags

`func (o *UpdateUserRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateUserRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateUserRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateUserRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateUserRequest) GetStatus() UserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateUserRequest) GetStatusOk() (*UserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateUserRequest) SetStatus(v UserStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateUserRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *UpdateUserRequest) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *UpdateUserRequest) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *UpdateUserRequest) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *UpdateUserRequest) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetBio

`func (o *UpdateUserRequest) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UpdateUserRequest) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UpdateUserRequest) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *UpdateUserRequest) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetBioLinks

`func (o *UpdateUserRequest) GetBioLinks() []string`

GetBioLinks returns the BioLinks field if non-nil, zero value otherwise.

### GetBioLinksOk

`func (o *UpdateUserRequest) GetBioLinksOk() (*[]string, bool)`

GetBioLinksOk returns a tuple with the BioLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioLinks

`func (o *UpdateUserRequest) SetBioLinks(v []string)`

SetBioLinks sets BioLinks field to given value.

### HasBioLinks

`func (o *UpdateUserRequest) HasBioLinks() bool`

HasBioLinks returns a boolean if a field has been set.

### GetUserIcon

`func (o *UpdateUserRequest) GetUserIcon() string`

GetUserIcon returns the UserIcon field if non-nil, zero value otherwise.

### GetUserIconOk

`func (o *UpdateUserRequest) GetUserIconOk() (*string, bool)`

GetUserIconOk returns a tuple with the UserIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIcon

`func (o *UpdateUserRequest) SetUserIcon(v string)`

SetUserIcon sets UserIcon field to given value.

### HasUserIcon

`func (o *UpdateUserRequest) HasUserIcon() bool`

HasUserIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


