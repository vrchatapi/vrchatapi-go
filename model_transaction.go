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

// checks if the Transaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transaction{}

// Transaction 
type Transaction struct {
	Id string `json:"id"`
	Status TransactionStatus `json:"status"`
	Subscription Subscription `json:"subscription"`
	Sandbox bool `json:"sandbox"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Steam *TransactionSteamInfo `json:"steam,omitempty"`
	Agreement *TransactionAgreement `json:"agreement,omitempty"`
	Error string `json:"error"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(id string, status TransactionStatus, subscription Subscription, sandbox bool, createdAt time.Time, updatedAt time.Time, error_ string) *Transaction {
	this := Transaction{}
	this.Id = id
	this.Status = status
	this.Subscription = subscription
	this.Sandbox = sandbox
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Error = error_
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	var status TransactionStatus = ACTIVE
	this.Status = status
	var sandbox bool = false
	this.Sandbox = sandbox
	return &this
}

// GetId returns the Id field value
func (o *Transaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Transaction) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *Transaction) GetStatus() TransactionStatus {
	if o == nil {
		var ret TransactionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetStatusOk() (*TransactionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Transaction) SetStatus(v TransactionStatus) {
	o.Status = v
}

// GetSubscription returns the Subscription field value
func (o *Transaction) GetSubscription() Subscription {
	if o == nil {
		var ret Subscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetSubscriptionOk() (*Subscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *Transaction) SetSubscription(v Subscription) {
	o.Subscription = v
}

// GetSandbox returns the Sandbox field value
func (o *Transaction) GetSandbox() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sandbox
}

// GetSandboxOk returns a tuple with the Sandbox field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetSandboxOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sandbox, true
}

// SetSandbox sets field value
func (o *Transaction) SetSandbox(v bool) {
	o.Sandbox = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Transaction) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Transaction) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Transaction) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Transaction) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetSteam returns the Steam field value if set, zero value otherwise.
func (o *Transaction) GetSteam() TransactionSteamInfo {
	if o == nil || IsNil(o.Steam) {
		var ret TransactionSteamInfo
		return ret
	}
	return *o.Steam
}

// GetSteamOk returns a tuple with the Steam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetSteamOk() (*TransactionSteamInfo, bool) {
	if o == nil || IsNil(o.Steam) {
		return nil, false
	}
	return o.Steam, true
}

// HasSteam returns a boolean if a field has been set.
func (o *Transaction) HasSteam() bool {
	if o != nil && !IsNil(o.Steam) {
		return true
	}

	return false
}

// SetSteam gets a reference to the given TransactionSteamInfo and assigns it to the Steam field.
func (o *Transaction) SetSteam(v TransactionSteamInfo) {
	o.Steam = &v
}

// GetAgreement returns the Agreement field value if set, zero value otherwise.
func (o *Transaction) GetAgreement() TransactionAgreement {
	if o == nil || IsNil(o.Agreement) {
		var ret TransactionAgreement
		return ret
	}
	return *o.Agreement
}

// GetAgreementOk returns a tuple with the Agreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetAgreementOk() (*TransactionAgreement, bool) {
	if o == nil || IsNil(o.Agreement) {
		return nil, false
	}
	return o.Agreement, true
}

// HasAgreement returns a boolean if a field has been set.
func (o *Transaction) HasAgreement() bool {
	if o != nil && !IsNil(o.Agreement) {
		return true
	}

	return false
}

// SetAgreement gets a reference to the given TransactionAgreement and assigns it to the Agreement field.
func (o *Transaction) SetAgreement(v TransactionAgreement) {
	o.Agreement = &v
}

// GetError returns the Error field value
func (o *Transaction) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *Transaction) SetError(v string) {
	o.Error = v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["subscription"] = o.Subscription
	toSerialize["sandbox"] = o.Sandbox
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.Steam) {
		toSerialize["steam"] = o.Steam
	}
	if !IsNil(o.Agreement) {
		toSerialize["agreement"] = o.Agreement
	}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


