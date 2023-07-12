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

// checks if the InfoPushData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfoPushData{}

// InfoPushData 
type InfoPushData struct {
	ContentList *DynamicContentRow `json:"contentList,omitempty"`
	Description *string `json:"description,omitempty"`
	ImageUrl *string `json:"imageUrl,omitempty"`
	Name *string `json:"name,omitempty"`
	OnPressed *InfoPushDataClickable `json:"onPressed,omitempty"`
	Template *string `json:"template,omitempty"`
	Version *string `json:"version,omitempty"`
	Article *InfoPushDataArticle `json:"article,omitempty"`
}

// NewInfoPushData instantiates a new InfoPushData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfoPushData() *InfoPushData {
	this := InfoPushData{}
	return &this
}

// NewInfoPushDataWithDefaults instantiates a new InfoPushData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfoPushDataWithDefaults() *InfoPushData {
	this := InfoPushData{}
	return &this
}

// GetContentList returns the ContentList field value if set, zero value otherwise.
func (o *InfoPushData) GetContentList() DynamicContentRow {
	if o == nil || IsNil(o.ContentList) {
		var ret DynamicContentRow
		return ret
	}
	return *o.ContentList
}

// GetContentListOk returns a tuple with the ContentList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoPushData) GetContentListOk() (*DynamicContentRow, bool) {
	if o == nil || IsNil(o.ContentList) {
		return nil, false
	}
	return o.ContentList, true
}

// HasContentList returns a boolean if a field has been set.
func (o *InfoPushData) HasContentList() bool {
	if o != nil && !IsNil(o.ContentList) {
		return true
	}

	return false
}

// SetContentList gets a reference to the given DynamicContentRow and assigns it to the ContentList field.
func (o *InfoPushData) SetContentList(v DynamicContentRow) {
	o.ContentList = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InfoPushData) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoPushData) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InfoPushData) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InfoPushData) SetDescription(v string) {
	o.Description = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *InfoPushData) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoPushData) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *InfoPushData) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *InfoPushData) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InfoPushData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoPushData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InfoPushData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InfoPushData) SetName(v string) {
	o.Name = &v
}

// GetOnPressed returns the OnPressed field value if set, zero value otherwise.
func (o *InfoPushData) GetOnPressed() InfoPushDataClickable {
	if o == nil || IsNil(o.OnPressed) {
		var ret InfoPushDataClickable
		return ret
	}
	return *o.OnPressed
}

// GetOnPressedOk returns a tuple with the OnPressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoPushData) GetOnPressedOk() (*InfoPushDataClickable, bool) {
	if o == nil || IsNil(o.OnPressed) {
		return nil, false
	}
	return o.OnPressed, true
}

// HasOnPressed returns a boolean if a field has been set.
func (o *InfoPushData) HasOnPressed() bool {
	if o != nil && !IsNil(o.OnPressed) {
		return true
	}

	return false
}

// SetOnPressed gets a reference to the given InfoPushDataClickable and assigns it to the OnPressed field.
func (o *InfoPushData) SetOnPressed(v InfoPushDataClickable) {
	o.OnPressed = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *InfoPushData) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoPushData) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *InfoPushData) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *InfoPushData) SetTemplate(v string) {
	o.Template = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InfoPushData) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoPushData) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InfoPushData) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InfoPushData) SetVersion(v string) {
	o.Version = &v
}

// GetArticle returns the Article field value if set, zero value otherwise.
func (o *InfoPushData) GetArticle() InfoPushDataArticle {
	if o == nil || IsNil(o.Article) {
		var ret InfoPushDataArticle
		return ret
	}
	return *o.Article
}

// GetArticleOk returns a tuple with the Article field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoPushData) GetArticleOk() (*InfoPushDataArticle, bool) {
	if o == nil || IsNil(o.Article) {
		return nil, false
	}
	return o.Article, true
}

// HasArticle returns a boolean if a field has been set.
func (o *InfoPushData) HasArticle() bool {
	if o != nil && !IsNil(o.Article) {
		return true
	}

	return false
}

// SetArticle gets a reference to the given InfoPushDataArticle and assigns it to the Article field.
func (o *InfoPushData) SetArticle(v InfoPushDataArticle) {
	o.Article = &v
}

func (o InfoPushData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfoPushData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentList) {
		toSerialize["contentList"] = o.ContentList
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OnPressed) {
		toSerialize["onPressed"] = o.OnPressed
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Article) {
		toSerialize["article"] = o.Article
	}
	return toSerialize, nil
}

type NullableInfoPushData struct {
	value *InfoPushData
	isSet bool
}

func (v NullableInfoPushData) Get() *InfoPushData {
	return v.value
}

func (v *NullableInfoPushData) Set(val *InfoPushData) {
	v.value = val
	v.isSet = true
}

func (v NullableInfoPushData) IsSet() bool {
	return v.isSet
}

func (v *NullableInfoPushData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfoPushData(val *InfoPushData) *NullableInfoPushData {
	return &NullableInfoPushData{value: val, isSet: true}
}

func (v NullableInfoPushData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfoPushData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


