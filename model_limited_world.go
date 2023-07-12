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

// checks if the LimitedWorld type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LimitedWorld{}

// LimitedWorld 
type LimitedWorld struct {
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId string `json:"authorId"`
	AuthorName string `json:"authorName"`
	Capacity int32 `json:"capacity"`
	CreatedAt time.Time `json:"created_at"`
	Favorites int32 `json:"favorites"`
	Heat int32 `json:"heat"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	Id string `json:"id"`
	ImageUrl string `json:"imageUrl"`
	LabsPublicationDate string `json:"labsPublicationDate"`
	Name string `json:"name"`
	Occupants int32 `json:"occupants"`
	Organization string `json:"organization"`
	Popularity int32 `json:"popularity"`
	PublicationDate string `json:"publicationDate"`
	ReleaseStatus ReleaseStatus `json:"releaseStatus"`
	//  
	Tags []string `json:"tags"`
	ThumbnailImageUrl string `json:"thumbnailImageUrl"`
	//  
	UnityPackages []LimitedUnityPackage `json:"unityPackages"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewLimitedWorld instantiates a new LimitedWorld object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimitedWorld(authorId string, authorName string, capacity int32, createdAt time.Time, favorites int32, heat int32, id string, imageUrl string, labsPublicationDate string, name string, occupants int32, organization string, popularity int32, publicationDate string, releaseStatus ReleaseStatus, tags []string, thumbnailImageUrl string, unityPackages []LimitedUnityPackage, updatedAt time.Time) *LimitedWorld {
	this := LimitedWorld{}
	this.AuthorId = authorId
	this.AuthorName = authorName
	this.Capacity = capacity
	this.CreatedAt = createdAt
	this.Favorites = favorites
	this.Heat = heat
	this.Id = id
	this.ImageUrl = imageUrl
	this.LabsPublicationDate = labsPublicationDate
	this.Name = name
	this.Occupants = occupants
	this.Organization = organization
	this.Popularity = popularity
	this.PublicationDate = publicationDate
	this.ReleaseStatus = releaseStatus
	this.Tags = tags
	this.ThumbnailImageUrl = thumbnailImageUrl
	this.UnityPackages = unityPackages
	this.UpdatedAt = updatedAt
	return &this
}

// NewLimitedWorldWithDefaults instantiates a new LimitedWorld object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitedWorldWithDefaults() *LimitedWorld {
	this := LimitedWorld{}
	var favorites int32 = 0
	this.Favorites = favorites
	var heat int32 = 0
	this.Heat = heat
	var occupants int32 = 0
	this.Occupants = occupants
	var organization string = "vrchat"
	this.Organization = organization
	var popularity int32 = 0
	this.Popularity = popularity
	var releaseStatus ReleaseStatus = PUBLIC
	this.ReleaseStatus = releaseStatus
	return &this
}

// GetAuthorId returns the AuthorId field value
func (o *LimitedWorld) GetAuthorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetAuthorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorId, true
}

// SetAuthorId sets field value
func (o *LimitedWorld) SetAuthorId(v string) {
	o.AuthorId = v
}

// GetAuthorName returns the AuthorName field value
func (o *LimitedWorld) GetAuthorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetAuthorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorName, true
}

// SetAuthorName sets field value
func (o *LimitedWorld) SetAuthorName(v string) {
	o.AuthorName = v
}

// GetCapacity returns the Capacity field value
func (o *LimitedWorld) GetCapacity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetCapacityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capacity, true
}

// SetCapacity sets field value
func (o *LimitedWorld) SetCapacity(v int32) {
	o.Capacity = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *LimitedWorld) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LimitedWorld) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetFavorites returns the Favorites field value
func (o *LimitedWorld) GetFavorites() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Favorites
}

// GetFavoritesOk returns a tuple with the Favorites field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetFavoritesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Favorites, true
}

// SetFavorites sets field value
func (o *LimitedWorld) SetFavorites(v int32) {
	o.Favorites = v
}

// GetHeat returns the Heat field value
func (o *LimitedWorld) GetHeat() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Heat
}

// GetHeatOk returns a tuple with the Heat field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetHeatOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Heat, true
}

// SetHeat sets field value
func (o *LimitedWorld) SetHeat(v int32) {
	o.Heat = v
}

// GetId returns the Id field value
func (o *LimitedWorld) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LimitedWorld) SetId(v string) {
	o.Id = v
}

// GetImageUrl returns the ImageUrl field value
func (o *LimitedWorld) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *LimitedWorld) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetLabsPublicationDate returns the LabsPublicationDate field value
func (o *LimitedWorld) GetLabsPublicationDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabsPublicationDate
}

// GetLabsPublicationDateOk returns a tuple with the LabsPublicationDate field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetLabsPublicationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabsPublicationDate, true
}

// SetLabsPublicationDate sets field value
func (o *LimitedWorld) SetLabsPublicationDate(v string) {
	o.LabsPublicationDate = v
}

// GetName returns the Name field value
func (o *LimitedWorld) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LimitedWorld) SetName(v string) {
	o.Name = v
}

// GetOccupants returns the Occupants field value
func (o *LimitedWorld) GetOccupants() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Occupants
}

// GetOccupantsOk returns a tuple with the Occupants field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetOccupantsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Occupants, true
}

// SetOccupants sets field value
func (o *LimitedWorld) SetOccupants(v int32) {
	o.Occupants = v
}

// GetOrganization returns the Organization field value
func (o *LimitedWorld) GetOrganization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *LimitedWorld) SetOrganization(v string) {
	o.Organization = v
}

// GetPopularity returns the Popularity field value
func (o *LimitedWorld) GetPopularity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetPopularityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Popularity, true
}

// SetPopularity sets field value
func (o *LimitedWorld) SetPopularity(v int32) {
	o.Popularity = v
}

// GetPublicationDate returns the PublicationDate field value
func (o *LimitedWorld) GetPublicationDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicationDate
}

// GetPublicationDateOk returns a tuple with the PublicationDate field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetPublicationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicationDate, true
}

// SetPublicationDate sets field value
func (o *LimitedWorld) SetPublicationDate(v string) {
	o.PublicationDate = v
}

// GetReleaseStatus returns the ReleaseStatus field value
func (o *LimitedWorld) GetReleaseStatus() ReleaseStatus {
	if o == nil {
		var ret ReleaseStatus
		return ret
	}

	return o.ReleaseStatus
}

// GetReleaseStatusOk returns a tuple with the ReleaseStatus field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetReleaseStatusOk() (*ReleaseStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseStatus, true
}

// SetReleaseStatus sets field value
func (o *LimitedWorld) SetReleaseStatus(v ReleaseStatus) {
	o.ReleaseStatus = v
}

// GetTags returns the Tags field value
func (o *LimitedWorld) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *LimitedWorld) SetTags(v []string) {
	o.Tags = v
}

// GetThumbnailImageUrl returns the ThumbnailImageUrl field value
func (o *LimitedWorld) GetThumbnailImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbnailImageUrl
}

// GetThumbnailImageUrlOk returns a tuple with the ThumbnailImageUrl field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetThumbnailImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThumbnailImageUrl, true
}

// SetThumbnailImageUrl sets field value
func (o *LimitedWorld) SetThumbnailImageUrl(v string) {
	o.ThumbnailImageUrl = v
}

// GetUnityPackages returns the UnityPackages field value
func (o *LimitedWorld) GetUnityPackages() []LimitedUnityPackage {
	if o == nil {
		var ret []LimitedUnityPackage
		return ret
	}

	return o.UnityPackages
}

// GetUnityPackagesOk returns a tuple with the UnityPackages field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetUnityPackagesOk() ([]LimitedUnityPackage, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnityPackages, true
}

// SetUnityPackages sets field value
func (o *LimitedWorld) SetUnityPackages(v []LimitedUnityPackage) {
	o.UnityPackages = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *LimitedWorld) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *LimitedWorld) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *LimitedWorld) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o LimitedWorld) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LimitedWorld) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorId"] = o.AuthorId
	toSerialize["authorName"] = o.AuthorName
	toSerialize["capacity"] = o.Capacity
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["favorites"] = o.Favorites
	toSerialize["heat"] = o.Heat
	toSerialize["id"] = o.Id
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["labsPublicationDate"] = o.LabsPublicationDate
	toSerialize["name"] = o.Name
	toSerialize["occupants"] = o.Occupants
	toSerialize["organization"] = o.Organization
	toSerialize["popularity"] = o.Popularity
	toSerialize["publicationDate"] = o.PublicationDate
	toSerialize["releaseStatus"] = o.ReleaseStatus
	toSerialize["tags"] = o.Tags
	toSerialize["thumbnailImageUrl"] = o.ThumbnailImageUrl
	toSerialize["unityPackages"] = o.UnityPackages
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableLimitedWorld struct {
	value *LimitedWorld
	isSet bool
}

func (v NullableLimitedWorld) Get() *LimitedWorld {
	return v.value
}

func (v *NullableLimitedWorld) Set(val *LimitedWorld) {
	v.value = val
	v.isSet = true
}

func (v NullableLimitedWorld) IsSet() bool {
	return v.isSet
}

func (v *NullableLimitedWorld) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimitedWorld(val *LimitedWorld) *NullableLimitedWorld {
	return &NullableLimitedWorld{value: val, isSet: true}
}

func (v NullableLimitedWorld) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimitedWorld) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


