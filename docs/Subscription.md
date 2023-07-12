# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SteamItemId** | **string** |  | 
**Amount** | **float32** |  | 
**Description** | **string** |  | 
**Period** | [**SubscriptionPeriod**](SubscriptionPeriod.md) |  | [default to MONTH]
**Tier** | **float32** |  | 

## Methods

### NewSubscription

`func NewSubscription(id string, steamItemId string, amount float32, description string, period SubscriptionPeriod, tier float32, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.


### GetSteamItemId

`func (o *Subscription) GetSteamItemId() string`

GetSteamItemId returns the SteamItemId field if non-nil, zero value otherwise.

### GetSteamItemIdOk

`func (o *Subscription) GetSteamItemIdOk() (*string, bool)`

GetSteamItemIdOk returns a tuple with the SteamItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamItemId

`func (o *Subscription) SetSteamItemId(v string)`

SetSteamItemId sets SteamItemId field to given value.


### GetAmount

`func (o *Subscription) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Subscription) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Subscription) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *Subscription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Subscription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Subscription) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPeriod

`func (o *Subscription) GetPeriod() SubscriptionPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Subscription) GetPeriodOk() (*SubscriptionPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Subscription) SetPeriod(v SubscriptionPeriod)`

SetPeriod sets Period field to given value.


### GetTier

`func (o *Subscription) GetTier() float32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *Subscription) GetTierOk() (*float32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *Subscription) SetTier(v float32)`

SetTier sets Tier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


