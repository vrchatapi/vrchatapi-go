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

// checks if the UpdateUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserRequest{}

// UpdateUserRequest struct for UpdateUserRequest
type UpdateUserRequest struct {
	Email *string `json:"email,omitempty"`
	Birthday *string `json:"birthday,omitempty"`
	AcceptedTOSVersion *float32 `json:"acceptedTOSVersion,omitempty"`
	//  
	Tags []string `json:"tags,omitempty"`
	Status *UserStatus `json:"status,omitempty"`
	StatusDescription *string `json:"statusDescription,omitempty"`
	Bio *string `json:"bio,omitempty"`
	BioLinks []string `json:"bioLinks,omitempty"`
	// MUST be a valid VRChat /file/ url.
	UserIcon *string `json:"userIcon,omitempty"`
}

// NewUpdateUserRequest instantiates a new UpdateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserRequest() *UpdateUserRequest {
	this := UpdateUserRequest{}
	var status UserStatus = OFFLINE
	this.Status = &status
	return &this
}

// NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserRequestWithDefaults() *UpdateUserRequest {
	this := UpdateUserRequest{}
	var status UserStatus = OFFLINE
	this.Status = &status
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateUserRequest) SetEmail(v string) {
	o.Email = &v
}

// GetBirthday returns the Birthday field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetBirthday() string {
	if o == nil || IsNil(o.Birthday) {
		var ret string
		return ret
	}
	return *o.Birthday
}

// GetBirthdayOk returns a tuple with the Birthday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetBirthdayOk() (*string, bool) {
	if o == nil || IsNil(o.Birthday) {
		return nil, false
	}
	return o.Birthday, true
}

// HasBirthday returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasBirthday() bool {
	if o != nil && !IsNil(o.Birthday) {
		return true
	}

	return false
}

// SetBirthday gets a reference to the given string and assigns it to the Birthday field.
func (o *UpdateUserRequest) SetBirthday(v string) {
	o.Birthday = &v
}

// GetAcceptedTOSVersion returns the AcceptedTOSVersion field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetAcceptedTOSVersion() float32 {
	if o == nil || IsNil(o.AcceptedTOSVersion) {
		var ret float32
		return ret
	}
	return *o.AcceptedTOSVersion
}

// GetAcceptedTOSVersionOk returns a tuple with the AcceptedTOSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetAcceptedTOSVersionOk() (*float32, bool) {
	if o == nil || IsNil(o.AcceptedTOSVersion) {
		return nil, false
	}
	return o.AcceptedTOSVersion, true
}

// HasAcceptedTOSVersion returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasAcceptedTOSVersion() bool {
	if o != nil && !IsNil(o.AcceptedTOSVersion) {
		return true
	}

	return false
}

// SetAcceptedTOSVersion gets a reference to the given float32 and assigns it to the AcceptedTOSVersion field.
func (o *UpdateUserRequest) SetAcceptedTOSVersion(v float32) {
	o.AcceptedTOSVersion = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateUserRequest) SetTags(v []string) {
	o.Tags = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetStatus() UserStatus {
	if o == nil || IsNil(o.Status) {
		var ret UserStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetStatusOk() (*UserStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given UserStatus and assigns it to the Status field.
func (o *UpdateUserRequest) SetStatus(v UserStatus) {
	o.Status = &v
}

// GetStatusDescription returns the StatusDescription field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetStatusDescription() string {
	if o == nil || IsNil(o.StatusDescription) {
		var ret string
		return ret
	}
	return *o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetStatusDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDescription) {
		return nil, false
	}
	return o.StatusDescription, true
}

// HasStatusDescription returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasStatusDescription() bool {
	if o != nil && !IsNil(o.StatusDescription) {
		return true
	}

	return false
}

// SetStatusDescription gets a reference to the given string and assigns it to the StatusDescription field.
func (o *UpdateUserRequest) SetStatusDescription(v string) {
	o.StatusDescription = &v
}

// GetBio returns the Bio field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *UpdateUserRequest) SetBio(v string) {
	o.Bio = &v
}

// GetBioLinks returns the BioLinks field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetBioLinks() []string {
	if o == nil || IsNil(o.BioLinks) {
		var ret []string
		return ret
	}
	return o.BioLinks
}

// GetBioLinksOk returns a tuple with the BioLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetBioLinksOk() ([]string, bool) {
	if o == nil || IsNil(o.BioLinks) {
		return nil, false
	}
	return o.BioLinks, true
}

// HasBioLinks returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasBioLinks() bool {
	if o != nil && !IsNil(o.BioLinks) {
		return true
	}

	return false
}

// SetBioLinks gets a reference to the given []string and assigns it to the BioLinks field.
func (o *UpdateUserRequest) SetBioLinks(v []string) {
	o.BioLinks = v
}

// GetUserIcon returns the UserIcon field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetUserIcon() string {
	if o == nil || IsNil(o.UserIcon) {
		var ret string
		return ret
	}
	return *o.UserIcon
}

// GetUserIconOk returns a tuple with the UserIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetUserIconOk() (*string, bool) {
	if o == nil || IsNil(o.UserIcon) {
		return nil, false
	}
	return o.UserIcon, true
}

// HasUserIcon returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasUserIcon() bool {
	if o != nil && !IsNil(o.UserIcon) {
		return true
	}

	return false
}

// SetUserIcon gets a reference to the given string and assigns it to the UserIcon field.
func (o *UpdateUserRequest) SetUserIcon(v string) {
	o.UserIcon = &v
}

func (o UpdateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Birthday) {
		toSerialize["birthday"] = o.Birthday
	}
	if !IsNil(o.AcceptedTOSVersion) {
		toSerialize["acceptedTOSVersion"] = o.AcceptedTOSVersion
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusDescription) {
		toSerialize["statusDescription"] = o.StatusDescription
	}
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	if !IsNil(o.BioLinks) {
		toSerialize["bioLinks"] = o.BioLinks
	}
	if !IsNil(o.UserIcon) {
		toSerialize["userIcon"] = o.UserIcon
	}
	return toSerialize, nil
}

type NullableUpdateUserRequest struct {
	value *UpdateUserRequest
	isSet bool
}

func (v NullableUpdateUserRequest) Get() *UpdateUserRequest {
	return v.value
}

func (v *NullableUpdateUserRequest) Set(val *UpdateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserRequest(val *UpdateUserRequest) *NullableUpdateUserRequest {
	return &NullableUpdateUserRequest{value: val, isSet: true}
}

func (v NullableUpdateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


