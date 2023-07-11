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

// checks if the FileVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileVersion{}

// FileVersion 
type FileVersion struct {
	CreatedAt time.Time `json:"created_at"`
	// Usually only present if `true`
	Deleted *bool `json:"deleted,omitempty"`
	Delta *FileData `json:"delta,omitempty"`
	File *FileData `json:"file,omitempty"`
	Signature *FileData `json:"signature,omitempty"`
	Status FileStatus `json:"status"`
	// Incremental version counter, can only be increased.
	Version int32 `json:"version"`
}

// NewFileVersion instantiates a new FileVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileVersion(createdAt time.Time, status FileStatus, version int32) *FileVersion {
	this := FileVersion{}
	this.CreatedAt = createdAt
	var deleted bool = true
	this.Deleted = &deleted
	this.Status = status
	this.Version = version
	return &this
}

// NewFileVersionWithDefaults instantiates a new FileVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileVersionWithDefaults() *FileVersion {
	this := FileVersion{}
	var deleted bool = true
	this.Deleted = &deleted
	var status FileStatus = WAITING
	this.Status = status
	var version int32 = 0
	this.Version = version
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *FileVersion) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FileVersion) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FileVersion) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *FileVersion) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersion) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *FileVersion) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *FileVersion) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDelta returns the Delta field value if set, zero value otherwise.
func (o *FileVersion) GetDelta() FileData {
	if o == nil || IsNil(o.Delta) {
		var ret FileData
		return ret
	}
	return *o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersion) GetDeltaOk() (*FileData, bool) {
	if o == nil || IsNil(o.Delta) {
		return nil, false
	}
	return o.Delta, true
}

// HasDelta returns a boolean if a field has been set.
func (o *FileVersion) HasDelta() bool {
	if o != nil && !IsNil(o.Delta) {
		return true
	}

	return false
}

// SetDelta gets a reference to the given FileData and assigns it to the Delta field.
func (o *FileVersion) SetDelta(v FileData) {
	o.Delta = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *FileVersion) GetFile() FileData {
	if o == nil || IsNil(o.File) {
		var ret FileData
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersion) GetFileOk() (*FileData, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *FileVersion) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given FileData and assigns it to the File field.
func (o *FileVersion) SetFile(v FileData) {
	o.File = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *FileVersion) GetSignature() FileData {
	if o == nil || IsNil(o.Signature) {
		var ret FileData
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVersion) GetSignatureOk() (*FileData, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *FileVersion) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given FileData and assigns it to the Signature field.
func (o *FileVersion) SetSignature(v FileData) {
	o.Signature = &v
}

// GetStatus returns the Status field value
func (o *FileVersion) GetStatus() FileStatus {
	if o == nil {
		var ret FileStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FileVersion) GetStatusOk() (*FileStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FileVersion) SetStatus(v FileStatus) {
	o.Status = v
}

// GetVersion returns the Version field value
func (o *FileVersion) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FileVersion) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *FileVersion) SetVersion(v int32) {
	o.Version = v
}

func (o FileVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !IsNil(o.Delta) {
		toSerialize["delta"] = o.Delta
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	toSerialize["status"] = o.Status
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableFileVersion struct {
	value *FileVersion
	isSet bool
}

func (v NullableFileVersion) Get() *FileVersion {
	return v.value
}

func (v *NullableFileVersion) Set(val *FileVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableFileVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableFileVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileVersion(val *FileVersion) *NullableFileVersion {
	return &NullableFileVersion{value: val, isSet: true}
}

func (v NullableFileVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


