# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ShortCode** | Pointer to **string** |  | [optional] 
**Discriminator** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **NullableString** |  | [optional] 
**BannerUrl** | Pointer to **NullableString** |  | [optional] 
**Privacy** | Pointer to [**GroupPrivacy**](GroupPrivacy.md) |  | [optional] [default to DEFAULT]
**OwnerId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**Rules** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to **[]string** |  | [optional] 
**Languages** | Pointer to **[]string** |  | [optional] 
**IconId** | Pointer to **NullableString** |  | [optional] 
**BannerId** | Pointer to **NullableString** |  | [optional] 
**MemberCount** | Pointer to **int32** |  | [optional] 
**MemberCountSyncedAt** | Pointer to **time.Time** |  | [optional] 
**IsVerified** | Pointer to **bool** |  | [optional] [default to false]
**JoinState** | Pointer to [**GroupJoinState**](GroupJoinState.md) |  | [optional] [default to OPEN]
**Tags** | Pointer to **[]string** |   | [optional] 
**Galleries** | Pointer to [**[]GroupGallery**](GroupGallery.md) |   | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**OnlineMemberCount** | Pointer to **int32** |  | [optional] 
**MembershipStatus** | Pointer to [**GroupMemberStatus**](GroupMemberStatus.md) |  | [optional] [default to INACTIVE]
**MyMember** | Pointer to [**GroupMyMember**](GroupMyMember.md) |  | [optional] 
**Roles** | Pointer to [**[]GroupRole**](GroupRole.md) | Only returned if ?includeRoles&#x3D;true is specified. | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Group) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortCode

`func (o *Group) GetShortCode() string`

GetShortCode returns the ShortCode field if non-nil, zero value otherwise.

### GetShortCodeOk

`func (o *Group) GetShortCodeOk() (*string, bool)`

GetShortCodeOk returns a tuple with the ShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCode

`func (o *Group) SetShortCode(v string)`

SetShortCode sets ShortCode field to given value.

### HasShortCode

`func (o *Group) HasShortCode() bool`

HasShortCode returns a boolean if a field has been set.

### GetDiscriminator

`func (o *Group) GetDiscriminator() string`

GetDiscriminator returns the Discriminator field if non-nil, zero value otherwise.

### GetDiscriminatorOk

`func (o *Group) GetDiscriminatorOk() (*string, bool)`

GetDiscriminatorOk returns a tuple with the Discriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscriminator

`func (o *Group) SetDiscriminator(v string)`

SetDiscriminator sets Discriminator field to given value.

### HasDiscriminator

`func (o *Group) HasDiscriminator() bool`

HasDiscriminator returns a boolean if a field has been set.

### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Group) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconUrl

`func (o *Group) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *Group) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *Group) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *Group) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *Group) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *Group) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetBannerUrl

`func (o *Group) GetBannerUrl() string`

GetBannerUrl returns the BannerUrl field if non-nil, zero value otherwise.

### GetBannerUrlOk

`func (o *Group) GetBannerUrlOk() (*string, bool)`

GetBannerUrlOk returns a tuple with the BannerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerUrl

`func (o *Group) SetBannerUrl(v string)`

SetBannerUrl sets BannerUrl field to given value.

### HasBannerUrl

`func (o *Group) HasBannerUrl() bool`

HasBannerUrl returns a boolean if a field has been set.

### SetBannerUrlNil

`func (o *Group) SetBannerUrlNil(b bool)`

 SetBannerUrlNil sets the value for BannerUrl to be an explicit nil

### UnsetBannerUrl
`func (o *Group) UnsetBannerUrl()`

UnsetBannerUrl ensures that no value is present for BannerUrl, not even an explicit nil
### GetPrivacy

`func (o *Group) GetPrivacy() GroupPrivacy`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *Group) GetPrivacyOk() (*GroupPrivacy, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *Group) SetPrivacy(v GroupPrivacy)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *Group) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetOwnerId

`func (o *Group) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Group) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Group) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Group) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetRules

`func (o *Group) GetRules() string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Group) GetRulesOk() (*string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Group) SetRules(v string)`

SetRules sets Rules field to given value.

### HasRules

`func (o *Group) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *Group) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *Group) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetLinks

`func (o *Group) GetLinks() []string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Group) GetLinksOk() (*[]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Group) SetLinks(v []string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Group) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetLanguages

`func (o *Group) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *Group) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *Group) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *Group) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetIconId

`func (o *Group) GetIconId() string`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *Group) GetIconIdOk() (*string, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *Group) SetIconId(v string)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *Group) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### SetIconIdNil

`func (o *Group) SetIconIdNil(b bool)`

 SetIconIdNil sets the value for IconId to be an explicit nil

### UnsetIconId
`func (o *Group) UnsetIconId()`

UnsetIconId ensures that no value is present for IconId, not even an explicit nil
### GetBannerId

`func (o *Group) GetBannerId() string`

GetBannerId returns the BannerId field if non-nil, zero value otherwise.

### GetBannerIdOk

`func (o *Group) GetBannerIdOk() (*string, bool)`

GetBannerIdOk returns a tuple with the BannerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerId

`func (o *Group) SetBannerId(v string)`

SetBannerId sets BannerId field to given value.

### HasBannerId

`func (o *Group) HasBannerId() bool`

HasBannerId returns a boolean if a field has been set.

### SetBannerIdNil

`func (o *Group) SetBannerIdNil(b bool)`

 SetBannerIdNil sets the value for BannerId to be an explicit nil

### UnsetBannerId
`func (o *Group) UnsetBannerId()`

UnsetBannerId ensures that no value is present for BannerId, not even an explicit nil
### GetMemberCount

`func (o *Group) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *Group) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *Group) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *Group) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetMemberCountSyncedAt

`func (o *Group) GetMemberCountSyncedAt() time.Time`

GetMemberCountSyncedAt returns the MemberCountSyncedAt field if non-nil, zero value otherwise.

### GetMemberCountSyncedAtOk

`func (o *Group) GetMemberCountSyncedAtOk() (*time.Time, bool)`

GetMemberCountSyncedAtOk returns a tuple with the MemberCountSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCountSyncedAt

`func (o *Group) SetMemberCountSyncedAt(v time.Time)`

SetMemberCountSyncedAt sets MemberCountSyncedAt field to given value.

### HasMemberCountSyncedAt

`func (o *Group) HasMemberCountSyncedAt() bool`

HasMemberCountSyncedAt returns a boolean if a field has been set.

### GetIsVerified

`func (o *Group) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *Group) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *Group) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *Group) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetJoinState

`func (o *Group) GetJoinState() GroupJoinState`

GetJoinState returns the JoinState field if non-nil, zero value otherwise.

### GetJoinStateOk

`func (o *Group) GetJoinStateOk() (*GroupJoinState, bool)`

GetJoinStateOk returns a tuple with the JoinState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinState

`func (o *Group) SetJoinState(v GroupJoinState)`

SetJoinState sets JoinState field to given value.

### HasJoinState

`func (o *Group) HasJoinState() bool`

HasJoinState returns a boolean if a field has been set.

### GetTags

`func (o *Group) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Group) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Group) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Group) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetGalleries

`func (o *Group) GetGalleries() []GroupGallery`

GetGalleries returns the Galleries field if non-nil, zero value otherwise.

### GetGalleriesOk

`func (o *Group) GetGalleriesOk() (*[]GroupGallery, bool)`

GetGalleriesOk returns a tuple with the Galleries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGalleries

`func (o *Group) SetGalleries(v []GroupGallery)`

SetGalleries sets Galleries field to given value.

### HasGalleries

`func (o *Group) HasGalleries() bool`

HasGalleries returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Group) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Group) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Group) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Group) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOnlineMemberCount

`func (o *Group) GetOnlineMemberCount() int32`

GetOnlineMemberCount returns the OnlineMemberCount field if non-nil, zero value otherwise.

### GetOnlineMemberCountOk

`func (o *Group) GetOnlineMemberCountOk() (*int32, bool)`

GetOnlineMemberCountOk returns a tuple with the OnlineMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineMemberCount

`func (o *Group) SetOnlineMemberCount(v int32)`

SetOnlineMemberCount sets OnlineMemberCount field to given value.

### HasOnlineMemberCount

`func (o *Group) HasOnlineMemberCount() bool`

HasOnlineMemberCount returns a boolean if a field has been set.

### GetMembershipStatus

`func (o *Group) GetMembershipStatus() GroupMemberStatus`

GetMembershipStatus returns the MembershipStatus field if non-nil, zero value otherwise.

### GetMembershipStatusOk

`func (o *Group) GetMembershipStatusOk() (*GroupMemberStatus, bool)`

GetMembershipStatusOk returns a tuple with the MembershipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipStatus

`func (o *Group) SetMembershipStatus(v GroupMemberStatus)`

SetMembershipStatus sets MembershipStatus field to given value.

### HasMembershipStatus

`func (o *Group) HasMembershipStatus() bool`

HasMembershipStatus returns a boolean if a field has been set.

### GetMyMember

`func (o *Group) GetMyMember() GroupMyMember`

GetMyMember returns the MyMember field if non-nil, zero value otherwise.

### GetMyMemberOk

`func (o *Group) GetMyMemberOk() (*GroupMyMember, bool)`

GetMyMemberOk returns a tuple with the MyMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyMember

`func (o *Group) SetMyMember(v GroupMyMember)`

SetMyMember sets MyMember field to given value.

### HasMyMember

`func (o *Group) HasMyMember() bool`

HasMyMember returns a boolean if a field has been set.

### GetRoles

`func (o *Group) GetRoles() []GroupRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Group) GetRolesOk() (*[]GroupRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Group) SetRoles(v []GroupRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Group) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *Group) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *Group) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


