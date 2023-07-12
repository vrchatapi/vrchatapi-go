# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Status** | [**TransactionStatus**](TransactionStatus.md) |  | [default to ACTIVE]
**Subscription** | [**Subscription**](Subscription.md) |  | 
**Sandbox** | **bool** |  | [default to false]
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Steam** | Pointer to [**TransactionSteamInfo**](TransactionSteamInfo.md) |  | [optional] 
**Agreement** | Pointer to [**TransactionAgreement**](TransactionAgreement.md) |  | [optional] 
**Error** | **string** |  | 

## Methods

### NewTransaction

`func NewTransaction(id string, status TransactionStatus, subscription Subscription, sandbox bool, createdAt time.Time, updatedAt time.Time, error_ string, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *Transaction) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.


### GetSubscription

`func (o *Transaction) GetSubscription() Subscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *Transaction) GetSubscriptionOk() (*Subscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *Transaction) SetSubscription(v Subscription)`

SetSubscription sets Subscription field to given value.


### GetSandbox

`func (o *Transaction) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *Transaction) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *Transaction) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.


### GetCreatedAt

`func (o *Transaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Transaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Transaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Transaction) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Transaction) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Transaction) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSteam

`func (o *Transaction) GetSteam() TransactionSteamInfo`

GetSteam returns the Steam field if non-nil, zero value otherwise.

### GetSteamOk

`func (o *Transaction) GetSteamOk() (*TransactionSteamInfo, bool)`

GetSteamOk returns a tuple with the Steam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteam

`func (o *Transaction) SetSteam(v TransactionSteamInfo)`

SetSteam sets Steam field to given value.

### HasSteam

`func (o *Transaction) HasSteam() bool`

HasSteam returns a boolean if a field has been set.

### GetAgreement

`func (o *Transaction) GetAgreement() TransactionAgreement`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *Transaction) GetAgreementOk() (*TransactionAgreement, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *Transaction) SetAgreement(v TransactionAgreement)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *Transaction) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetError

`func (o *Transaction) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Transaction) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Transaction) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


