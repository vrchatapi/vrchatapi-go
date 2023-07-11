# GroupMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**IsRepresenting** | Pointer to **bool** | Whether the user is representing the group. This makes the group show up above the name tag in-game. | [optional] [default to false]
**User** | Pointer to [**GroupMemberLimitedUser**](GroupMemberLimitedUser.md) |  | [optional] 
**RoleIds** | Pointer to **[]string** |  | [optional] 
**JoinedAt** | Pointer to **time.Time** |  | [optional] 
**MembershipStatus** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**IsSubscribedToAnnouncements** | Pointer to **bool** |  | [optional] [default to false]
**CreatedAt** | Pointer to **NullableTime** | Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user. | [optional] 
**BannedAt** | Pointer to **NullableTime** | Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user. | [optional] 
**ManagerNotes** | Pointer to **NullableString** | Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user. | [optional] 

## Methods

### NewGroupMember

`func NewGroupMember() *GroupMember`

NewGroupMember instantiates a new GroupMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMemberWithDefaults

`func NewGroupMemberWithDefaults() *GroupMember`

NewGroupMemberWithDefaults instantiates a new GroupMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *GroupMember) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupMember) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupMember) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupMember) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetUserId

`func (o *GroupMember) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GroupMember) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GroupMember) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GroupMember) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetIsRepresenting

`func (o *GroupMember) GetIsRepresenting() bool`

GetIsRepresenting returns the IsRepresenting field if non-nil, zero value otherwise.

### GetIsRepresentingOk

`func (o *GroupMember) GetIsRepresentingOk() (*bool, bool)`

GetIsRepresentingOk returns a tuple with the IsRepresenting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRepresenting

`func (o *GroupMember) SetIsRepresenting(v bool)`

SetIsRepresenting sets IsRepresenting field to given value.

### HasIsRepresenting

`func (o *GroupMember) HasIsRepresenting() bool`

HasIsRepresenting returns a boolean if a field has been set.

### GetUser

`func (o *GroupMember) GetUser() GroupMemberLimitedUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GroupMember) GetUserOk() (*GroupMemberLimitedUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GroupMember) SetUser(v GroupMemberLimitedUser)`

SetUser sets User field to given value.

### HasUser

`func (o *GroupMember) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRoleIds

`func (o *GroupMember) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *GroupMember) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *GroupMember) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *GroupMember) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetJoinedAt

`func (o *GroupMember) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *GroupMember) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *GroupMember) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *GroupMember) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetMembershipStatus

`func (o *GroupMember) GetMembershipStatus() string`

GetMembershipStatus returns the MembershipStatus field if non-nil, zero value otherwise.

### GetMembershipStatusOk

`func (o *GroupMember) GetMembershipStatusOk() (*string, bool)`

GetMembershipStatusOk returns a tuple with the MembershipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipStatus

`func (o *GroupMember) SetMembershipStatus(v string)`

SetMembershipStatus sets MembershipStatus field to given value.

### HasMembershipStatus

`func (o *GroupMember) HasMembershipStatus() bool`

HasMembershipStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *GroupMember) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GroupMember) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GroupMember) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GroupMember) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetIsSubscribedToAnnouncements

`func (o *GroupMember) GetIsSubscribedToAnnouncements() bool`

GetIsSubscribedToAnnouncements returns the IsSubscribedToAnnouncements field if non-nil, zero value otherwise.

### GetIsSubscribedToAnnouncementsOk

`func (o *GroupMember) GetIsSubscribedToAnnouncementsOk() (*bool, bool)`

GetIsSubscribedToAnnouncementsOk returns a tuple with the IsSubscribedToAnnouncements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribedToAnnouncements

`func (o *GroupMember) SetIsSubscribedToAnnouncements(v bool)`

SetIsSubscribedToAnnouncements sets IsSubscribedToAnnouncements field to given value.

### HasIsSubscribedToAnnouncements

`func (o *GroupMember) HasIsSubscribedToAnnouncements() bool`

HasIsSubscribedToAnnouncements returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GroupMember) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupMember) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupMember) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupMember) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GroupMember) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GroupMember) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetBannedAt

`func (o *GroupMember) GetBannedAt() time.Time`

GetBannedAt returns the BannedAt field if non-nil, zero value otherwise.

### GetBannedAtOk

`func (o *GroupMember) GetBannedAtOk() (*time.Time, bool)`

GetBannedAtOk returns a tuple with the BannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedAt

`func (o *GroupMember) SetBannedAt(v time.Time)`

SetBannedAt sets BannedAt field to given value.

### HasBannedAt

`func (o *GroupMember) HasBannedAt() bool`

HasBannedAt returns a boolean if a field has been set.

### SetBannedAtNil

`func (o *GroupMember) SetBannedAtNil(b bool)`

 SetBannedAtNil sets the value for BannedAt to be an explicit nil

### UnsetBannedAt
`func (o *GroupMember) UnsetBannedAt()`

UnsetBannedAt ensures that no value is present for BannedAt, not even an explicit nil
### GetManagerNotes

`func (o *GroupMember) GetManagerNotes() string`

GetManagerNotes returns the ManagerNotes field if non-nil, zero value otherwise.

### GetManagerNotesOk

`func (o *GroupMember) GetManagerNotesOk() (*string, bool)`

GetManagerNotesOk returns a tuple with the ManagerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerNotes

`func (o *GroupMember) SetManagerNotes(v string)`

SetManagerNotes sets ManagerNotes field to given value.

### HasManagerNotes

`func (o *GroupMember) HasManagerNotes() bool`

HasManagerNotes returns a boolean if a field has been set.

### SetManagerNotesNil

`func (o *GroupMember) SetManagerNotesNil(b bool)`

 SetManagerNotesNil sets the value for ManagerNotes to be an explicit nil

### UnsetManagerNotes
`func (o *GroupMember) UnsetManagerNotes()`

UnsetManagerNotes ensures that no value is present for ManagerNotes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


