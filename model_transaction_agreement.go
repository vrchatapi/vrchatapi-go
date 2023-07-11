/*
VRChat API Documentation

![VRChat API Banner](https://vrchatapi.github.io/assets/img/api_banner_1500x400.png)  # Welcome to the VRChat API  Before we begin, we would like to state this is a **COMMUNITY DRIVEN PROJECT**. This means that everything you read on here was written by the community itself and is **not** officially supported by VRChat. The documentation is provided \"AS IS\", and any action you take towards VRChat is completely your own responsibility.  The documentation and additional libraries SHALL ONLY be used for applications interacting with VRChat's API in accordance with their [Terms of Service](https://hello.vrchat.com/legal) and [Community Guidelines](https://hello.vrchat.com/community-guidelines), and MUST NOT be used for modifying the client, \"avatar ripping\", or other illegal activities. Malicious usage or spamming the API may result in account termination. Certain parts of the API are also more sensitive than others, for example moderation, so please tread extra carefully and read the warnings when present.  ![Tupper Policy on API](https://i.imgur.com/yLlW7Ok.png)  Finally, use of the API using applications other than the approved methods (website, VRChat application, Unity SDK) is not officially supported. VRChat provides no guarantee or support for external applications using the API. Access to API endpoints may break **at any time, without notice**. Therefore, please **do not ping** VRChat Staff in the VRChat Discord if you are having API problems, as they do not provide API support. We will make a best effort in keeping this documentation and associated language libraries up to date, but things might be outdated or missing. If you find that something is no longer valid, please contact us on Discord or [create an issue](https://github.com/vrchatapi/specification/issues) and tell us so we can fix it.  # Getting Started  The VRChat API can be used to programmatically retrieve or update information regarding your profile, friends, avatars, worlds and more. The API consists of two parts, \"Photon\" which is only used in-game, and the \"Web API\" which is used by both the game and the website. This documentation focuses only on the Web API.  The API is designed around the REST ideology, providing semi-simple and usually predictable URIs to access and modify objects. Requests support standard HTTP methods like GET, PUT, POST, and DELETE and standard status codes. Response bodies are always UTF-8 encoded JSON objects, unless explicitly documented otherwise.  <div class=\"callout callout-error\">   <strong>🛑 Warning! Do not touch Photon!</strong><br>   Photon is only used by the in-game client and should <b>not</b> be touched. Doing so may result in permanent account termination. </div>  <div class=\"callout callout-info\">   <strong>ℹ️ API Key and Authentication</strong><br>   The API Key has always been the same and is currently <code>JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26</code>.   Read <a href=\"#tag--authentication\">Authentication</a> for how to log in. </div>  # Using the API  For simply exploring what the API can do it is strongly recommended to download [Insomnia](https://insomnia.rest/download), a free and open-source API client that's great for sending requests to the API in an orderly fashion. Insomnia allows you to send data in the format that's required for VRChat's API. It is also possible to try out the API in your browser, by first logging in at [vrchat.com/home](https://vrchat.com/home/) and then going to [vrchat.com/api/1/auth/user](https://vrchat.com/api/1/auth/user), but the information will be much harder to work with.  For more permanent operation such as software development it is instead recommended to use one of the existing language SDKs. This community project maintains API libraries in several languages, which allows you to interact with the API with simple function calls rather than having to implement the HTTP protocol yourself. Most of these libraries are automatically generated from the API specification, sometimes with additional helpful wrapper code to make usage easier. This allows them to be almost automatically updated and expanded upon as soon as a new feature is introduced in the specification itself. The libraries can be found on [GitHub](https://github.com/vrchatapi) or following:  * [NodeJS (JavaScript)](https://www.npmjs.com/package/vrchat) * [Dart](https://pub.dev/packages/vrchat_dart) * [Rust](https://crates.io/crates/vrchatapi) * [C#](https://github.com/vrchatapi/vrchatapi-csharp) * [Python](https://github.com/vrchatapi/vrchatapi-python)  # Pagination  Most endpoints enforce pagination, meaning they will only return 10 entries by default, and never more than 100.<br> Using both the limit and offset parameters allows you to easily paginate through a large number of objects.  | Query Parameter | Type | Description | | ----------|--|------- | | `n` | integer  | The number of objects to return. This value often defaults to 10. Highest limit is always 100.| | `offset` | integer  | A zero-based offset from the default object sorting.|  If a request returns fewer objects than the `limit` parameter, there are no more items available to return.  # Contribution  Do you want to get involved in the documentation effort? Do you want to help improve one of the language API libraries? This project is an [OPEN Open Source Project](https://openopensource.org)! This means that individuals making significant and valuable contributions are given commit-access to the project. It also means we are very open and welcoming of new people making contributions, unlike some more guarded open-source projects.  [![Discord](https://img.shields.io/static/v1?label=vrchatapi&message=discord&color=blueviolet&style=for-the-badge)](https://discord.gg/qjZE9C9fkB)

API version: 1.12.0
Contact: vrchatapi.lpv0t@aries.fyi
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vrchatapi

import (
	"encoding/json"
)

// checks if the TransactionAgreement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionAgreement{}

// TransactionAgreement 
type TransactionAgreement struct {
	AgreementId string `json:"agreementId"`
	ItemId float32 `json:"itemId"`
	// This is NOT TransactionStatus, but whatever Steam return.
	Status string `json:"status"`
	Period string `json:"period"`
	Frequency float32 `json:"frequency"`
	BillingType string `json:"billingType"`
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
	RecurringAmt float32 `json:"recurringAmt"`
	Currency string `json:"currency"`
	TimeCreated string `json:"timeCreated"`
	NextPayment string `json:"nextPayment"`
	LastPayment string `json:"lastPayment"`
	LastAmount float32 `json:"lastAmount"`
	LastAmountVat float32 `json:"lastAmountVat"`
	Outstanding float32 `json:"outstanding"`
	FailedAttempts float32 `json:"failedAttempts"`
}

// NewTransactionAgreement instantiates a new TransactionAgreement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionAgreement(agreementId string, itemId float32, status string, period string, frequency float32, billingType string, startDate string, endDate string, recurringAmt float32, currency string, timeCreated string, nextPayment string, lastPayment string, lastAmount float32, lastAmountVat float32, outstanding float32, failedAttempts float32) *TransactionAgreement {
	this := TransactionAgreement{}
	this.AgreementId = agreementId
	this.ItemId = itemId
	this.Status = status
	this.Period = period
	this.Frequency = frequency
	this.BillingType = billingType
	this.StartDate = startDate
	this.EndDate = endDate
	this.RecurringAmt = recurringAmt
	this.Currency = currency
	this.TimeCreated = timeCreated
	this.NextPayment = nextPayment
	this.LastPayment = lastPayment
	this.LastAmount = lastAmount
	this.LastAmountVat = lastAmountVat
	this.Outstanding = outstanding
	this.FailedAttempts = failedAttempts
	return &this
}

// NewTransactionAgreementWithDefaults instantiates a new TransactionAgreement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionAgreementWithDefaults() *TransactionAgreement {
	this := TransactionAgreement{}
	return &this
}

// GetAgreementId returns the AgreementId field value
func (o *TransactionAgreement) GetAgreementId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetAgreementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgreementId, true
}

// SetAgreementId sets field value
func (o *TransactionAgreement) SetAgreementId(v string) {
	o.AgreementId = v
}

// GetItemId returns the ItemId field value
func (o *TransactionAgreement) GetItemId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetItemIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *TransactionAgreement) SetItemId(v float32) {
	o.ItemId = v
}

// GetStatus returns the Status field value
func (o *TransactionAgreement) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransactionAgreement) SetStatus(v string) {
	o.Status = v
}

// GetPeriod returns the Period field value
func (o *TransactionAgreement) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *TransactionAgreement) SetPeriod(v string) {
	o.Period = v
}

// GetFrequency returns the Frequency field value
func (o *TransactionAgreement) GetFrequency() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetFrequencyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *TransactionAgreement) SetFrequency(v float32) {
	o.Frequency = v
}

// GetBillingType returns the BillingType field value
func (o *TransactionAgreement) GetBillingType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingType
}

// GetBillingTypeOk returns a tuple with the BillingType field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetBillingTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingType, true
}

// SetBillingType sets field value
func (o *TransactionAgreement) SetBillingType(v string) {
	o.BillingType = v
}

// GetStartDate returns the StartDate field value
func (o *TransactionAgreement) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *TransactionAgreement) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *TransactionAgreement) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *TransactionAgreement) SetEndDate(v string) {
	o.EndDate = v
}

// GetRecurringAmt returns the RecurringAmt field value
func (o *TransactionAgreement) GetRecurringAmt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RecurringAmt
}

// GetRecurringAmtOk returns a tuple with the RecurringAmt field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetRecurringAmtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecurringAmt, true
}

// SetRecurringAmt sets field value
func (o *TransactionAgreement) SetRecurringAmt(v float32) {
	o.RecurringAmt = v
}

// GetCurrency returns the Currency field value
func (o *TransactionAgreement) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransactionAgreement) SetCurrency(v string) {
	o.Currency = v
}

// GetTimeCreated returns the TimeCreated field value
func (o *TransactionAgreement) GetTimeCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeCreated
}

// GetTimeCreatedOk returns a tuple with the TimeCreated field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetTimeCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeCreated, true
}

// SetTimeCreated sets field value
func (o *TransactionAgreement) SetTimeCreated(v string) {
	o.TimeCreated = v
}

// GetNextPayment returns the NextPayment field value
func (o *TransactionAgreement) GetNextPayment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPayment
}

// GetNextPaymentOk returns a tuple with the NextPayment field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetNextPaymentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPayment, true
}

// SetNextPayment sets field value
func (o *TransactionAgreement) SetNextPayment(v string) {
	o.NextPayment = v
}

// GetLastPayment returns the LastPayment field value
func (o *TransactionAgreement) GetLastPayment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastPayment
}

// GetLastPaymentOk returns a tuple with the LastPayment field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetLastPaymentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastPayment, true
}

// SetLastPayment sets field value
func (o *TransactionAgreement) SetLastPayment(v string) {
	o.LastPayment = v
}

// GetLastAmount returns the LastAmount field value
func (o *TransactionAgreement) GetLastAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LastAmount
}

// GetLastAmountOk returns a tuple with the LastAmount field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetLastAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastAmount, true
}

// SetLastAmount sets field value
func (o *TransactionAgreement) SetLastAmount(v float32) {
	o.LastAmount = v
}

// GetLastAmountVat returns the LastAmountVat field value
func (o *TransactionAgreement) GetLastAmountVat() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LastAmountVat
}

// GetLastAmountVatOk returns a tuple with the LastAmountVat field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetLastAmountVatOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastAmountVat, true
}

// SetLastAmountVat sets field value
func (o *TransactionAgreement) SetLastAmountVat(v float32) {
	o.LastAmountVat = v
}

// GetOutstanding returns the Outstanding field value
func (o *TransactionAgreement) GetOutstanding() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Outstanding
}

// GetOutstandingOk returns a tuple with the Outstanding field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetOutstandingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outstanding, true
}

// SetOutstanding sets field value
func (o *TransactionAgreement) SetOutstanding(v float32) {
	o.Outstanding = v
}

// GetFailedAttempts returns the FailedAttempts field value
func (o *TransactionAgreement) GetFailedAttempts() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.FailedAttempts
}

// GetFailedAttemptsOk returns a tuple with the FailedAttempts field value
// and a boolean to check if the value has been set.
func (o *TransactionAgreement) GetFailedAttemptsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedAttempts, true
}

// SetFailedAttempts sets field value
func (o *TransactionAgreement) SetFailedAttempts(v float32) {
	o.FailedAttempts = v
}

func (o TransactionAgreement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionAgreement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agreementId"] = o.AgreementId
	toSerialize["itemId"] = o.ItemId
	toSerialize["status"] = o.Status
	toSerialize["period"] = o.Period
	toSerialize["frequency"] = o.Frequency
	toSerialize["billingType"] = o.BillingType
	toSerialize["startDate"] = o.StartDate
	toSerialize["endDate"] = o.EndDate
	toSerialize["recurringAmt"] = o.RecurringAmt
	toSerialize["currency"] = o.Currency
	toSerialize["timeCreated"] = o.TimeCreated
	toSerialize["nextPayment"] = o.NextPayment
	toSerialize["lastPayment"] = o.LastPayment
	toSerialize["lastAmount"] = o.LastAmount
	toSerialize["lastAmountVat"] = o.LastAmountVat
	toSerialize["outstanding"] = o.Outstanding
	toSerialize["failedAttempts"] = o.FailedAttempts
	return toSerialize, nil
}

type NullableTransactionAgreement struct {
	value *TransactionAgreement
	isSet bool
}

func (v NullableTransactionAgreement) Get() *TransactionAgreement {
	return v.value
}

func (v *NullableTransactionAgreement) Set(val *TransactionAgreement) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionAgreement) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionAgreement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionAgreement(val *TransactionAgreement) *NullableTransactionAgreement {
	return &NullableTransactionAgreement{value: val, isSet: true}
}

func (v NullableTransactionAgreement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionAgreement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


