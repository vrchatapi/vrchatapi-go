# GroupAuditLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**ActorId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**ActorDisplayname** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **string** | Typically GroupID or GroupRoleID, but could be other types of IDs. | [optional] 
**EventType** | Pointer to **string** | The type of event that occurred. This is a string that is prefixed with the type of object that the event occurred on. For example, a group role update event would be prefixed with &#x60;group.role&#x60;. | [optional] [default to "group.update"]
**Description** | Pointer to **string** | A human-readable description of the event. | [optional] 
**Data** | Pointer to **map[string]interface{}** | The data associated with the event. The format of this data is dependent on the event type. | [optional] 

## Methods

### NewGroupAuditLogEntry

`func NewGroupAuditLogEntry() *GroupAuditLogEntry`

NewGroupAuditLogEntry instantiates a new GroupAuditLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupAuditLogEntryWithDefaults

`func NewGroupAuditLogEntryWithDefaults() *GroupAuditLogEntry`

NewGroupAuditLogEntryWithDefaults instantiates a new GroupAuditLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupAuditLogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupAuditLogEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupAuditLogEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupAuditLogEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GroupAuditLogEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupAuditLogEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupAuditLogEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupAuditLogEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGroupId

`func (o *GroupAuditLogEntry) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupAuditLogEntry) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupAuditLogEntry) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupAuditLogEntry) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetActorId

`func (o *GroupAuditLogEntry) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *GroupAuditLogEntry) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *GroupAuditLogEntry) SetActorId(v string)`

SetActorId sets ActorId field to given value.

### HasActorId

`func (o *GroupAuditLogEntry) HasActorId() bool`

HasActorId returns a boolean if a field has been set.

### GetActorDisplayname

`func (o *GroupAuditLogEntry) GetActorDisplayname() string`

GetActorDisplayname returns the ActorDisplayname field if non-nil, zero value otherwise.

### GetActorDisplaynameOk

`func (o *GroupAuditLogEntry) GetActorDisplaynameOk() (*string, bool)`

GetActorDisplaynameOk returns a tuple with the ActorDisplayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorDisplayname

`func (o *GroupAuditLogEntry) SetActorDisplayname(v string)`

SetActorDisplayname sets ActorDisplayname field to given value.

### HasActorDisplayname

`func (o *GroupAuditLogEntry) HasActorDisplayname() bool`

HasActorDisplayname returns a boolean if a field has been set.

### GetTargetId

`func (o *GroupAuditLogEntry) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *GroupAuditLogEntry) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *GroupAuditLogEntry) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *GroupAuditLogEntry) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetEventType

`func (o *GroupAuditLogEntry) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *GroupAuditLogEntry) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *GroupAuditLogEntry) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *GroupAuditLogEntry) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetDescription

`func (o *GroupAuditLogEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupAuditLogEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupAuditLogEntry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupAuditLogEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetData

`func (o *GroupAuditLogEntry) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GroupAuditLogEntry) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GroupAuditLogEntry) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *GroupAuditLogEntry) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


