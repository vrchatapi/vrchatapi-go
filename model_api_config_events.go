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

// checks if the APIConfigEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIConfigEvents{}

// APIConfigEvents struct for APIConfigEvents
type APIConfigEvents struct {
	// Unknown
	DistanceClose int32 `json:"distanceClose"`
	// Unknown
	DistanceFactor int32 `json:"distanceFactor"`
	// Unknown
	DistanceFar int32 `json:"distanceFar"`
	// Unknown
	GroupDistance int32 `json:"groupDistance"`
	// Unknown
	MaximumBunchSize int32 `json:"maximumBunchSize"`
	// Unknown
	NotVisibleFactor int32 `json:"notVisibleFactor"`
	// Unknown
	PlayerOrderBucketSize int32 `json:"playerOrderBucketSize"`
	// Unknown
	PlayerOrderFactor int32 `json:"playerOrderFactor"`
	// Unknown
	SlowUpdateFactorThreshold int32 `json:"slowUpdateFactorThreshold"`
	// Unknown
	ViewSegmentLength int32 `json:"viewSegmentLength"`
}

// NewAPIConfigEvents instantiates a new APIConfigEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIConfigEvents(distanceClose int32, distanceFactor int32, distanceFar int32, groupDistance int32, maximumBunchSize int32, notVisibleFactor int32, playerOrderBucketSize int32, playerOrderFactor int32, slowUpdateFactorThreshold int32, viewSegmentLength int32) *APIConfigEvents {
	this := APIConfigEvents{}
	this.DistanceClose = distanceClose
	this.DistanceFactor = distanceFactor
	this.DistanceFar = distanceFar
	this.GroupDistance = groupDistance
	this.MaximumBunchSize = maximumBunchSize
	this.NotVisibleFactor = notVisibleFactor
	this.PlayerOrderBucketSize = playerOrderBucketSize
	this.PlayerOrderFactor = playerOrderFactor
	this.SlowUpdateFactorThreshold = slowUpdateFactorThreshold
	this.ViewSegmentLength = viewSegmentLength
	return &this
}

// NewAPIConfigEventsWithDefaults instantiates a new APIConfigEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIConfigEventsWithDefaults() *APIConfigEvents {
	this := APIConfigEvents{}
	return &this
}

// GetDistanceClose returns the DistanceClose field value
func (o *APIConfigEvents) GetDistanceClose() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DistanceClose
}

// GetDistanceCloseOk returns a tuple with the DistanceClose field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetDistanceCloseOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DistanceClose, true
}

// SetDistanceClose sets field value
func (o *APIConfigEvents) SetDistanceClose(v int32) {
	o.DistanceClose = v
}

// GetDistanceFactor returns the DistanceFactor field value
func (o *APIConfigEvents) GetDistanceFactor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DistanceFactor
}

// GetDistanceFactorOk returns a tuple with the DistanceFactor field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetDistanceFactorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DistanceFactor, true
}

// SetDistanceFactor sets field value
func (o *APIConfigEvents) SetDistanceFactor(v int32) {
	o.DistanceFactor = v
}

// GetDistanceFar returns the DistanceFar field value
func (o *APIConfigEvents) GetDistanceFar() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DistanceFar
}

// GetDistanceFarOk returns a tuple with the DistanceFar field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetDistanceFarOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DistanceFar, true
}

// SetDistanceFar sets field value
func (o *APIConfigEvents) SetDistanceFar(v int32) {
	o.DistanceFar = v
}

// GetGroupDistance returns the GroupDistance field value
func (o *APIConfigEvents) GetGroupDistance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupDistance
}

// GetGroupDistanceOk returns a tuple with the GroupDistance field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetGroupDistanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupDistance, true
}

// SetGroupDistance sets field value
func (o *APIConfigEvents) SetGroupDistance(v int32) {
	o.GroupDistance = v
}

// GetMaximumBunchSize returns the MaximumBunchSize field value
func (o *APIConfigEvents) GetMaximumBunchSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaximumBunchSize
}

// GetMaximumBunchSizeOk returns a tuple with the MaximumBunchSize field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetMaximumBunchSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumBunchSize, true
}

// SetMaximumBunchSize sets field value
func (o *APIConfigEvents) SetMaximumBunchSize(v int32) {
	o.MaximumBunchSize = v
}

// GetNotVisibleFactor returns the NotVisibleFactor field value
func (o *APIConfigEvents) GetNotVisibleFactor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NotVisibleFactor
}

// GetNotVisibleFactorOk returns a tuple with the NotVisibleFactor field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetNotVisibleFactorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotVisibleFactor, true
}

// SetNotVisibleFactor sets field value
func (o *APIConfigEvents) SetNotVisibleFactor(v int32) {
	o.NotVisibleFactor = v
}

// GetPlayerOrderBucketSize returns the PlayerOrderBucketSize field value
func (o *APIConfigEvents) GetPlayerOrderBucketSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PlayerOrderBucketSize
}

// GetPlayerOrderBucketSizeOk returns a tuple with the PlayerOrderBucketSize field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetPlayerOrderBucketSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlayerOrderBucketSize, true
}

// SetPlayerOrderBucketSize sets field value
func (o *APIConfigEvents) SetPlayerOrderBucketSize(v int32) {
	o.PlayerOrderBucketSize = v
}

// GetPlayerOrderFactor returns the PlayerOrderFactor field value
func (o *APIConfigEvents) GetPlayerOrderFactor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PlayerOrderFactor
}

// GetPlayerOrderFactorOk returns a tuple with the PlayerOrderFactor field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetPlayerOrderFactorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlayerOrderFactor, true
}

// SetPlayerOrderFactor sets field value
func (o *APIConfigEvents) SetPlayerOrderFactor(v int32) {
	o.PlayerOrderFactor = v
}

// GetSlowUpdateFactorThreshold returns the SlowUpdateFactorThreshold field value
func (o *APIConfigEvents) GetSlowUpdateFactorThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SlowUpdateFactorThreshold
}

// GetSlowUpdateFactorThresholdOk returns a tuple with the SlowUpdateFactorThreshold field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetSlowUpdateFactorThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowUpdateFactorThreshold, true
}

// SetSlowUpdateFactorThreshold sets field value
func (o *APIConfigEvents) SetSlowUpdateFactorThreshold(v int32) {
	o.SlowUpdateFactorThreshold = v
}

// GetViewSegmentLength returns the ViewSegmentLength field value
func (o *APIConfigEvents) GetViewSegmentLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ViewSegmentLength
}

// GetViewSegmentLengthOk returns a tuple with the ViewSegmentLength field value
// and a boolean to check if the value has been set.
func (o *APIConfigEvents) GetViewSegmentLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewSegmentLength, true
}

// SetViewSegmentLength sets field value
func (o *APIConfigEvents) SetViewSegmentLength(v int32) {
	o.ViewSegmentLength = v
}

func (o APIConfigEvents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIConfigEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["distanceClose"] = o.DistanceClose
	toSerialize["distanceFactor"] = o.DistanceFactor
	toSerialize["distanceFar"] = o.DistanceFar
	toSerialize["groupDistance"] = o.GroupDistance
	toSerialize["maximumBunchSize"] = o.MaximumBunchSize
	toSerialize["notVisibleFactor"] = o.NotVisibleFactor
	toSerialize["playerOrderBucketSize"] = o.PlayerOrderBucketSize
	toSerialize["playerOrderFactor"] = o.PlayerOrderFactor
	toSerialize["slowUpdateFactorThreshold"] = o.SlowUpdateFactorThreshold
	toSerialize["viewSegmentLength"] = o.ViewSegmentLength
	return toSerialize, nil
}

type NullableAPIConfigEvents struct {
	value *APIConfigEvents
	isSet bool
}

func (v NullableAPIConfigEvents) Get() *APIConfigEvents {
	return v.value
}

func (v *NullableAPIConfigEvents) Set(val *APIConfigEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIConfigEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIConfigEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIConfigEvents(val *APIConfigEvents) *NullableAPIConfigEvents {
	return &NullableAPIConfigEvents{value: val, isSet: true}
}

func (v NullableAPIConfigEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIConfigEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


