# UserSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TransactionId** | **string** |  | 
**Store** | **string** | Which \&quot;Store\&quot; it came from. Right now only Stores are \&quot;Steam\&quot; and \&quot;Admin\&quot;. | 
**SteamItemId** | Pointer to **string** |  | [optional] 
**Amount** | **float32** |  | 
**Description** | **string** |  | 
**Period** | [**SubscriptionPeriod**](SubscriptionPeriod.md) |  | [default to MONTH]
**Tier** | **float32** |  | 
**Active** | **bool** |  | [default to true]
**Status** | [**TransactionStatus**](TransactionStatus.md) |  | [default to ACTIVE]
**Expires** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**LicenseGroups** | **[]string** |  | 
**IsGift** | **bool** |  | [default to false]

## Methods

### NewUserSubscription

`func NewUserSubscription(id string, transactionId string, store string, amount float32, description string, period SubscriptionPeriod, tier float32, active bool, status TransactionStatus, expires time.Time, createdAt time.Time, updatedAt time.Time, licenseGroups []string, isGift bool, ) *UserSubscription`

NewUserSubscription instantiates a new UserSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSubscriptionWithDefaults

`func NewUserSubscriptionWithDefaults() *UserSubscription`

NewUserSubscriptionWithDefaults instantiates a new UserSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSubscription) SetId(v string)`

SetId sets Id field to given value.


### GetTransactionId

`func (o *UserSubscription) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *UserSubscription) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *UserSubscription) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetStore

`func (o *UserSubscription) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *UserSubscription) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *UserSubscription) SetStore(v string)`

SetStore sets Store field to given value.


### GetSteamItemId

`func (o *UserSubscription) GetSteamItemId() string`

GetSteamItemId returns the SteamItemId field if non-nil, zero value otherwise.

### GetSteamItemIdOk

`func (o *UserSubscription) GetSteamItemIdOk() (*string, bool)`

GetSteamItemIdOk returns a tuple with the SteamItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamItemId

`func (o *UserSubscription) SetSteamItemId(v string)`

SetSteamItemId sets SteamItemId field to given value.

### HasSteamItemId

`func (o *UserSubscription) HasSteamItemId() bool`

HasSteamItemId returns a boolean if a field has been set.

### GetAmount

`func (o *UserSubscription) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UserSubscription) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UserSubscription) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *UserSubscription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserSubscription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserSubscription) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPeriod

`func (o *UserSubscription) GetPeriod() SubscriptionPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *UserSubscription) GetPeriodOk() (*SubscriptionPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *UserSubscription) SetPeriod(v SubscriptionPeriod)`

SetPeriod sets Period field to given value.


### GetTier

`func (o *UserSubscription) GetTier() float32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *UserSubscription) GetTierOk() (*float32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *UserSubscription) SetTier(v float32)`

SetTier sets Tier field to given value.


### GetActive

`func (o *UserSubscription) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UserSubscription) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UserSubscription) SetActive(v bool)`

SetActive sets Active field to given value.


### GetStatus

`func (o *UserSubscription) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserSubscription) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserSubscription) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.


### GetExpires

`func (o *UserSubscription) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *UserSubscription) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *UserSubscription) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### GetCreatedAt

`func (o *UserSubscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserSubscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserSubscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserSubscription) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserSubscription) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserSubscription) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLicenseGroups

`func (o *UserSubscription) GetLicenseGroups() []string`

GetLicenseGroups returns the LicenseGroups field if non-nil, zero value otherwise.

### GetLicenseGroupsOk

`func (o *UserSubscription) GetLicenseGroupsOk() (*[]string, bool)`

GetLicenseGroupsOk returns a tuple with the LicenseGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseGroups

`func (o *UserSubscription) SetLicenseGroups(v []string)`

SetLicenseGroups sets LicenseGroups field to given value.


### GetIsGift

`func (o *UserSubscription) GetIsGift() bool`

GetIsGift returns the IsGift field if non-nil, zero value otherwise.

### GetIsGiftOk

`func (o *UserSubscription) GetIsGiftOk() (*bool, bool)`

GetIsGiftOk returns a tuple with the IsGift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGift

`func (o *UserSubscription) SetIsGift(v bool)`

SetIsGift sets IsGift field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


