# PaginatedGroupAuditLogEntryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]GroupAuditLogEntry**](GroupAuditLogEntry.md) |   | [optional] 
**TotalCount** | Pointer to **int32** | The total number of results that the query would return if there were no pagination. | [optional] 
**HasNext** | Pointer to **bool** | Whether there are more results after this page. | [optional] 

## Methods

### NewPaginatedGroupAuditLogEntryList

`func NewPaginatedGroupAuditLogEntryList() *PaginatedGroupAuditLogEntryList`

NewPaginatedGroupAuditLogEntryList instantiates a new PaginatedGroupAuditLogEntryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedGroupAuditLogEntryListWithDefaults

`func NewPaginatedGroupAuditLogEntryListWithDefaults() *PaginatedGroupAuditLogEntryList`

NewPaginatedGroupAuditLogEntryListWithDefaults instantiates a new PaginatedGroupAuditLogEntryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PaginatedGroupAuditLogEntryList) GetResults() []GroupAuditLogEntry`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedGroupAuditLogEntryList) GetResultsOk() (*[]GroupAuditLogEntry, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedGroupAuditLogEntryList) SetResults(v []GroupAuditLogEntry)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedGroupAuditLogEntryList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalCount

`func (o *PaginatedGroupAuditLogEntryList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedGroupAuditLogEntryList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedGroupAuditLogEntryList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedGroupAuditLogEntryList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetHasNext

`func (o *PaginatedGroupAuditLogEntryList) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PaginatedGroupAuditLogEntryList) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PaginatedGroupAuditLogEntryList) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *PaginatedGroupAuditLogEntryList) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


