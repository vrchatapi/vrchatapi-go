# UpdateGroupMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visibility** | Pointer to [**GroupUserVisibility**](GroupUserVisibility.md) |  | [optional] 
**IsSubscribedToAnnouncements** | Pointer to **bool** |  | [optional] 
**ManagerNotes** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateGroupMemberRequest

`func NewUpdateGroupMemberRequest() *UpdateGroupMemberRequest`

NewUpdateGroupMemberRequest instantiates a new UpdateGroupMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupMemberRequestWithDefaults

`func NewUpdateGroupMemberRequestWithDefaults() *UpdateGroupMemberRequest`

NewUpdateGroupMemberRequestWithDefaults instantiates a new UpdateGroupMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisibility

`func (o *UpdateGroupMemberRequest) GetVisibility() GroupUserVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateGroupMemberRequest) GetVisibilityOk() (*GroupUserVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateGroupMemberRequest) SetVisibility(v GroupUserVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateGroupMemberRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetIsSubscribedToAnnouncements

`func (o *UpdateGroupMemberRequest) GetIsSubscribedToAnnouncements() bool`

GetIsSubscribedToAnnouncements returns the IsSubscribedToAnnouncements field if non-nil, zero value otherwise.

### GetIsSubscribedToAnnouncementsOk

`func (o *UpdateGroupMemberRequest) GetIsSubscribedToAnnouncementsOk() (*bool, bool)`

GetIsSubscribedToAnnouncementsOk returns a tuple with the IsSubscribedToAnnouncements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribedToAnnouncements

`func (o *UpdateGroupMemberRequest) SetIsSubscribedToAnnouncements(v bool)`

SetIsSubscribedToAnnouncements sets IsSubscribedToAnnouncements field to given value.

### HasIsSubscribedToAnnouncements

`func (o *UpdateGroupMemberRequest) HasIsSubscribedToAnnouncements() bool`

HasIsSubscribedToAnnouncements returns a boolean if a field has been set.

### GetManagerNotes

`func (o *UpdateGroupMemberRequest) GetManagerNotes() string`

GetManagerNotes returns the ManagerNotes field if non-nil, zero value otherwise.

### GetManagerNotesOk

`func (o *UpdateGroupMemberRequest) GetManagerNotesOk() (*string, bool)`

GetManagerNotesOk returns a tuple with the ManagerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerNotes

`func (o *UpdateGroupMemberRequest) SetManagerNotes(v string)`

SetManagerNotes sets ManagerNotes field to given value.

### HasManagerNotes

`func (o *UpdateGroupMemberRequest) HasManagerNotes() bool`

HasManagerNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


