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
	"time"
)

// checks if the UserSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSubscription{}

// UserSubscription 
type UserSubscription struct {
	Id string `json:"id"`
	TransactionId string `json:"transactionId"`
	// Which \"Store\" it came from. Right now only Stores are \"Steam\" and \"Admin\".
	Store string `json:"store"`
	SteamItemId *string `json:"steamItemId,omitempty"`
	Amount float32 `json:"amount"`
	Description string `json:"description"`
	Period SubscriptionPeriod `json:"period"`
	Tier float32 `json:"tier"`
	Active bool `json:"active"`
	Status TransactionStatus `json:"status"`
	Expires time.Time `json:"expires"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	LicenseGroups []string `json:"licenseGroups"`
	IsGift bool `json:"isGift"`
}

// NewUserSubscription instantiates a new UserSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSubscription(id string, transactionId string, store string, amount float32, description string, period SubscriptionPeriod, tier float32, active bool, status TransactionStatus, expires time.Time, createdAt time.Time, updatedAt time.Time, licenseGroups []string, isGift bool) *UserSubscription {
	this := UserSubscription{}
	this.Id = id
	this.TransactionId = transactionId
	this.Store = store
	this.Amount = amount
	this.Description = description
	this.Period = period
	this.Tier = tier
	this.Active = active
	this.Status = status
	this.Expires = expires
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.LicenseGroups = licenseGroups
	this.IsGift = isGift
	return &this
}

// NewUserSubscriptionWithDefaults instantiates a new UserSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSubscriptionWithDefaults() *UserSubscription {
	this := UserSubscription{}
	var period SubscriptionPeriod = MONTH
	this.Period = period
	var active bool = true
	this.Active = active
	var status TransactionStatus = ACTIVE
	this.Status = status
	var isGift bool = false
	this.IsGift = isGift
	return &this
}

// GetId returns the Id field value
func (o *UserSubscription) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserSubscription) SetId(v string) {
	o.Id = v
}

// GetTransactionId returns the TransactionId field value
func (o *UserSubscription) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *UserSubscription) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetStore returns the Store field value
func (o *UserSubscription) GetStore() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Store
}

// GetStoreOk returns a tuple with the Store field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetStoreOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Store, true
}

// SetStore sets field value
func (o *UserSubscription) SetStore(v string) {
	o.Store = v
}

// GetSteamItemId returns the SteamItemId field value if set, zero value otherwise.
func (o *UserSubscription) GetSteamItemId() string {
	if o == nil || IsNil(o.SteamItemId) {
		var ret string
		return ret
	}
	return *o.SteamItemId
}

// GetSteamItemIdOk returns a tuple with the SteamItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetSteamItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SteamItemId) {
		return nil, false
	}
	return o.SteamItemId, true
}

// HasSteamItemId returns a boolean if a field has been set.
func (o *UserSubscription) HasSteamItemId() bool {
	if o != nil && !IsNil(o.SteamItemId) {
		return true
	}

	return false
}

// SetSteamItemId gets a reference to the given string and assigns it to the SteamItemId field.
func (o *UserSubscription) SetSteamItemId(v string) {
	o.SteamItemId = &v
}

// GetAmount returns the Amount field value
func (o *UserSubscription) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *UserSubscription) SetAmount(v float32) {
	o.Amount = v
}

// GetDescription returns the Description field value
func (o *UserSubscription) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UserSubscription) SetDescription(v string) {
	o.Description = v
}

// GetPeriod returns the Period field value
func (o *UserSubscription) GetPeriod() SubscriptionPeriod {
	if o == nil {
		var ret SubscriptionPeriod
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetPeriodOk() (*SubscriptionPeriod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *UserSubscription) SetPeriod(v SubscriptionPeriod) {
	o.Period = v
}

// GetTier returns the Tier field value
func (o *UserSubscription) GetTier() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Tier
}

// GetTierOk returns a tuple with the Tier field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetTierOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tier, true
}

// SetTier sets field value
func (o *UserSubscription) SetTier(v float32) {
	o.Tier = v
}

// GetActive returns the Active field value
func (o *UserSubscription) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *UserSubscription) SetActive(v bool) {
	o.Active = v
}

// GetStatus returns the Status field value
func (o *UserSubscription) GetStatus() TransactionStatus {
	if o == nil {
		var ret TransactionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetStatusOk() (*TransactionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UserSubscription) SetStatus(v TransactionStatus) {
	o.Status = v
}

// GetExpires returns the Expires field value
func (o *UserSubscription) GetExpires() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expires, true
}

// SetExpires sets field value
func (o *UserSubscription) SetExpires(v time.Time) {
	o.Expires = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *UserSubscription) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UserSubscription) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *UserSubscription) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *UserSubscription) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetLicenseGroups returns the LicenseGroups field value
func (o *UserSubscription) GetLicenseGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LicenseGroups
}

// GetLicenseGroupsOk returns a tuple with the LicenseGroups field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetLicenseGroupsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LicenseGroups, true
}

// SetLicenseGroups sets field value
func (o *UserSubscription) SetLicenseGroups(v []string) {
	o.LicenseGroups = v
}

// GetIsGift returns the IsGift field value
func (o *UserSubscription) GetIsGift() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsGift
}

// GetIsGiftOk returns a tuple with the IsGift field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetIsGiftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsGift, true
}

// SetIsGift sets field value
func (o *UserSubscription) SetIsGift(v bool) {
	o.IsGift = v
}

func (o UserSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["transactionId"] = o.TransactionId
	toSerialize["store"] = o.Store
	if !IsNil(o.SteamItemId) {
		toSerialize["steamItemId"] = o.SteamItemId
	}
	toSerialize["amount"] = o.Amount
	toSerialize["description"] = o.Description
	toSerialize["period"] = o.Period
	toSerialize["tier"] = o.Tier
	toSerialize["active"] = o.Active
	toSerialize["status"] = o.Status
	toSerialize["expires"] = o.Expires
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["licenseGroups"] = o.LicenseGroups
	toSerialize["isGift"] = o.IsGift
	return toSerialize, nil
}

type NullableUserSubscription struct {
	value *UserSubscription
	isSet bool
}

func (v NullableUserSubscription) Get() *UserSubscription {
	return v.value
}

func (v *NullableUserSubscription) Set(val *UserSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSubscription(val *UserSubscription) *NullableUserSubscription {
	return &NullableUserSubscription{value: val, isSet: true}
}

func (v NullableUserSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


