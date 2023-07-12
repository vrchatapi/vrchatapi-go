# AccountDeletionLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Typically \&quot;Deletion requested\&quot; or \&quot;Deletion canceled\&quot;. Other messages like \&quot;Deletion completed\&quot; may exist, but are these are not possible to see as a regular user. | [optional] [default to "Deletion requested"]
**DeletionScheduled** | Pointer to **NullableTime** | When the deletion is scheduled to happen, standard is 14 days after the request. | [optional] 
**DateTime** | Pointer to **time.Time** | Date and time of the deletion request. | [optional] 

## Methods

### NewAccountDeletionLog

`func NewAccountDeletionLog() *AccountDeletionLog`

NewAccountDeletionLog instantiates a new AccountDeletionLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDeletionLogWithDefaults

`func NewAccountDeletionLogWithDefaults() *AccountDeletionLog`

NewAccountDeletionLogWithDefaults instantiates a new AccountDeletionLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AccountDeletionLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AccountDeletionLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AccountDeletionLog) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AccountDeletionLog) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDeletionScheduled

`func (o *AccountDeletionLog) GetDeletionScheduled() time.Time`

GetDeletionScheduled returns the DeletionScheduled field if non-nil, zero value otherwise.

### GetDeletionScheduledOk

`func (o *AccountDeletionLog) GetDeletionScheduledOk() (*time.Time, bool)`

GetDeletionScheduledOk returns a tuple with the DeletionScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionScheduled

`func (o *AccountDeletionLog) SetDeletionScheduled(v time.Time)`

SetDeletionScheduled sets DeletionScheduled field to given value.

### HasDeletionScheduled

`func (o *AccountDeletionLog) HasDeletionScheduled() bool`

HasDeletionScheduled returns a boolean if a field has been set.

### SetDeletionScheduledNil

`func (o *AccountDeletionLog) SetDeletionScheduledNil(b bool)`

 SetDeletionScheduledNil sets the value for DeletionScheduled to be an explicit nil

### UnsetDeletionScheduled
`func (o *AccountDeletionLog) UnsetDeletionScheduled()`

UnsetDeletionScheduled ensures that no value is present for DeletionScheduled, not even an explicit nil
### GetDateTime

`func (o *AccountDeletionLog) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *AccountDeletionLog) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *AccountDeletionLog) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *AccountDeletionLog) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


