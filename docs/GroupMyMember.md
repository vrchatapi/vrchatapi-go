# GroupMyMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**RoleIds** | Pointer to **[]string** |  | [optional] 
**ManagerNotes** | Pointer to **string** |  | [optional] 
**MembershipStatus** | Pointer to **string** |  | [optional] 
**IsSubscribedToAnnouncements** | Pointer to **bool** |  | [optional] [default to true]
**Visibility** | Pointer to **string** |  | [optional] 
**IsRepresenting** | Pointer to **bool** |  | [optional] [default to false]
**JoinedAt** | Pointer to **time.Time** |  | [optional] 
**BannedAt** | Pointer to **NullableString** |  | [optional] 
**Has2FA** | Pointer to **bool** |  | [optional] [default to false]
**Permissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGroupMyMember

`func NewGroupMyMember() *GroupMyMember`

NewGroupMyMember instantiates a new GroupMyMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMyMemberWithDefaults

`func NewGroupMyMemberWithDefaults() *GroupMyMember`

NewGroupMyMemberWithDefaults instantiates a new GroupMyMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupMyMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupMyMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupMyMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupMyMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *GroupMyMember) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupMyMember) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupMyMember) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupMyMember) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetUserId

`func (o *GroupMyMember) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GroupMyMember) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GroupMyMember) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GroupMyMember) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRoleIds

`func (o *GroupMyMember) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *GroupMyMember) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *GroupMyMember) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *GroupMyMember) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetManagerNotes

`func (o *GroupMyMember) GetManagerNotes() string`

GetManagerNotes returns the ManagerNotes field if non-nil, zero value otherwise.

### GetManagerNotesOk

`func (o *GroupMyMember) GetManagerNotesOk() (*string, bool)`

GetManagerNotesOk returns a tuple with the ManagerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerNotes

`func (o *GroupMyMember) SetManagerNotes(v string)`

SetManagerNotes sets ManagerNotes field to given value.

### HasManagerNotes

`func (o *GroupMyMember) HasManagerNotes() bool`

HasManagerNotes returns a boolean if a field has been set.

### GetMembershipStatus

`func (o *GroupMyMember) GetMembershipStatus() string`

GetMembershipStatus returns the MembershipStatus field if non-nil, zero value otherwise.

### GetMembershipStatusOk

`func (o *GroupMyMember) GetMembershipStatusOk() (*string, bool)`

GetMembershipStatusOk returns a tuple with the MembershipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipStatus

`func (o *GroupMyMember) SetMembershipStatus(v string)`

SetMembershipStatus sets MembershipStatus field to given value.

### HasMembershipStatus

`func (o *GroupMyMember) HasMembershipStatus() bool`

HasMembershipStatus returns a boolean if a field has been set.

### GetIsSubscribedToAnnouncements

`func (o *GroupMyMember) GetIsSubscribedToAnnouncements() bool`

GetIsSubscribedToAnnouncements returns the IsSubscribedToAnnouncements field if non-nil, zero value otherwise.

### GetIsSubscribedToAnnouncementsOk

`func (o *GroupMyMember) GetIsSubscribedToAnnouncementsOk() (*bool, bool)`

GetIsSubscribedToAnnouncementsOk returns a tuple with the IsSubscribedToAnnouncements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribedToAnnouncements

`func (o *GroupMyMember) SetIsSubscribedToAnnouncements(v bool)`

SetIsSubscribedToAnnouncements sets IsSubscribedToAnnouncements field to given value.

### HasIsSubscribedToAnnouncements

`func (o *GroupMyMember) HasIsSubscribedToAnnouncements() bool`

HasIsSubscribedToAnnouncements returns a boolean if a field has been set.

### GetVisibility

`func (o *GroupMyMember) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GroupMyMember) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GroupMyMember) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GroupMyMember) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetIsRepresenting

`func (o *GroupMyMember) GetIsRepresenting() bool`

GetIsRepresenting returns the IsRepresenting field if non-nil, zero value otherwise.

### GetIsRepresentingOk

`func (o *GroupMyMember) GetIsRepresentingOk() (*bool, bool)`

GetIsRepresentingOk returns a tuple with the IsRepresenting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRepresenting

`func (o *GroupMyMember) SetIsRepresenting(v bool)`

SetIsRepresenting sets IsRepresenting field to given value.

### HasIsRepresenting

`func (o *GroupMyMember) HasIsRepresenting() bool`

HasIsRepresenting returns a boolean if a field has been set.

### GetJoinedAt

`func (o *GroupMyMember) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *GroupMyMember) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *GroupMyMember) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *GroupMyMember) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetBannedAt

`func (o *GroupMyMember) GetBannedAt() string`

GetBannedAt returns the BannedAt field if non-nil, zero value otherwise.

### GetBannedAtOk

`func (o *GroupMyMember) GetBannedAtOk() (*string, bool)`

GetBannedAtOk returns a tuple with the BannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedAt

`func (o *GroupMyMember) SetBannedAt(v string)`

SetBannedAt sets BannedAt field to given value.

### HasBannedAt

`func (o *GroupMyMember) HasBannedAt() bool`

HasBannedAt returns a boolean if a field has been set.

### SetBannedAtNil

`func (o *GroupMyMember) SetBannedAtNil(b bool)`

 SetBannedAtNil sets the value for BannedAt to be an explicit nil

### UnsetBannedAt
`func (o *GroupMyMember) UnsetBannedAt()`

UnsetBannedAt ensures that no value is present for BannedAt, not even an explicit nil
### GetHas2FA

`func (o *GroupMyMember) GetHas2FA() bool`

GetHas2FA returns the Has2FA field if non-nil, zero value otherwise.

### GetHas2FAOk

`func (o *GroupMyMember) GetHas2FAOk() (*bool, bool)`

GetHas2FAOk returns a tuple with the Has2FA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHas2FA

`func (o *GroupMyMember) SetHas2FA(v bool)`

SetHas2FA sets Has2FA field to given value.

### HasHas2FA

`func (o *GroupMyMember) HasHas2FA() bool`

HasHas2FA returns a boolean if a field has been set.

### GetPermissions

`func (o *GroupMyMember) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GroupMyMember) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GroupMyMember) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GroupMyMember) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


