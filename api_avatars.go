/*
VRChat API Documentation

![VRChat API Banner](https://vrchatapi.github.io/assets/img/api_banner_1500x400.png)  # Welcome to the VRChat API  Before we begin, we would like to state this is a **COMMUNITY DRIVEN PROJECT**. This means that everything you read on here was written by the community itself and is **not** officially supported by VRChat. The documentation is provided \"AS IS\", and any action you take towards VRChat is completely your own responsibility.  The documentation and additional libraries SHALL ONLY be used for applications interacting with VRChat's API in accordance with their [Terms of Service](https://hello.vrchat.com/legal) and [Community Guidelines](https://hello.vrchat.com/community-guidelines), and MUST NOT be used for modifying the client, \"avatar ripping\", or other illegal activities. Malicious usage or spamming the API may result in account termination. Certain parts of the API are also more sensitive than others, for example moderation, so please tread extra carefully and read the warnings when present.  ![Tupper Policy on API](https://i.imgur.com/yLlW7Ok.png)  Finally, use of the API using applications other than the approved methods (website, VRChat application, Unity SDK) is not officially supported. VRChat provides no guarantee or support for external applications using the API. Access to API endpoints may break **at any time, without notice**. Therefore, please **do not ping** VRChat Staff in the VRChat Discord if you are having API problems, as they do not provide API support. We will make a best effort in keeping this documentation and associated language libraries up to date, but things might be outdated or missing. If you find that something is no longer valid, please contact us on Discord or [create an issue](https://github.com/vrchatapi/specification/issues) and tell us so we can fix it.  # Getting Started  The VRChat API can be used to programmatically retrieve or update information regarding your profile, friends, avatars, worlds and more. The API consists of two parts, \"Photon\" which is only used in-game, and the \"Web API\" which is used by both the game and the website. This documentation focuses only on the Web API.  The API is designed around the REST ideology, providing semi-simple and usually predictable URIs to access and modify objects. Requests support standard HTTP methods like GET, PUT, POST, and DELETE and standard status codes. Response bodies are always UTF-8 encoded JSON objects, unless explicitly documented otherwise.  <div class=\"callout callout-error\">   <strong>🛑 Warning! Do not touch Photon!</strong><br>   Photon is only used by the in-game client and should <b>not</b> be touched. Doing so may result in permanent account termination. </div>  <div class=\"callout callout-info\">   <strong>ℹ️ API Key and Authentication</strong><br>   The API Key has always been the same and is currently <code>JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26</code>.   Read <a href=\"#tag--authentication\">Authentication</a> for how to log in. </div>  # Using the API  For simply exploring what the API can do it is strongly recommended to download [Insomnia](https://insomnia.rest/download), a free and open-source API client that's great for sending requests to the API in an orderly fashion. Insomnia allows you to send data in the format that's required for VRChat's API. It is also possible to try out the API in your browser, by first logging in at [vrchat.com/home](https://vrchat.com/home/) and then going to [vrchat.com/api/1/auth/user](https://vrchat.com/api/1/auth/user), but the information will be much harder to work with.  For more permanent operation such as software development it is instead recommended to use one of the existing language SDKs. This community project maintains API libraries in several languages, which allows you to interact with the API with simple function calls rather than having to implement the HTTP protocol yourself. Most of these libraries are automatically generated from the API specification, sometimes with additional helpful wrapper code to make usage easier. This allows them to be almost automatically updated and expanded upon as soon as a new feature is introduced in the specification itself. The libraries can be found on [GitHub](https://github.com/vrchatapi) or following:  * [NodeJS (JavaScript)](https://www.npmjs.com/package/vrchat) * [Dart](https://pub.dev/packages/vrchat_dart) * [Rust](https://crates.io/crates/vrchatapi) * [C#](https://github.com/vrchatapi/vrchatapi-csharp) * [Python](https://github.com/vrchatapi/vrchatapi-python)  # Pagination  Most endpoints enforce pagination, meaning they will only return 10 entries by default, and never more than 100.<br> Using both the limit and offset parameters allows you to easily paginate through a large number of objects.  | Query Parameter | Type | Description | | ----------|--|------- | | `n` | integer  | The number of objects to return. This value often defaults to 10. Highest limit is always 100.| | `offset` | integer  | A zero-based offset from the default object sorting.|  If a request returns fewer objects than the `limit` parameter, there are no more items available to return.  # Contribution  Do you want to get involved in the documentation effort? Do you want to help improve one of the language API libraries? This project is an [OPEN Open Source Project](https://openopensource.org)! This means that individuals making significant and valuable contributions are given commit-access to the project. It also means we are very open and welcoming of new people making contributions, unlike some more guarded open-source projects.  [![Discord](https://img.shields.io/static/v1?label=vrchatapi&message=discord&color=blueviolet&style=for-the-badge)](https://discord.gg/qjZE9C9fkB)

API version: 1.12.0
Contact: vrchatapi.lpv0t@aries.fyi
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vrchatapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// AvatarsApiService AvatarsApi service
type AvatarsApiService service

type ApiCreateAvatarRequest struct {
	ctx context.Context
	ApiService *AvatarsApiService
	createAvatarRequest *CreateAvatarRequest
}

func (r ApiCreateAvatarRequest) CreateAvatarRequest(createAvatarRequest CreateAvatarRequest) ApiCreateAvatarRequest {
	r.createAvatarRequest = &createAvatarRequest
	return r
}

func (r ApiCreateAvatarRequest) Execute() (*Avatar, *http.Response, error) {
	return r.ApiService.CreateAvatarExecute(r)
}

/*
CreateAvatar Create Avatar

Create an avatar. It's possible to optionally specify a ID if you want a custom one. Attempting to create an Avatar with an already claimed ID will result in a DB error.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateAvatarRequest
*/
func (a *AvatarsApiService) CreateAvatar(ctx context.Context) ApiCreateAvatarRequest {
	return ApiCreateAvatarRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Avatar
func (a *AvatarsApiService) CreateAvatarExecute(r ApiCreateAvatarRequest) (*Avatar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Avatar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvatarsApiService.CreateAvatar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/avatars"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createAvatarRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteAvatarRequest struct {
	ctx context.Context
	ApiService *AvatarsApiService
	avatarId string
}

func (r ApiDeleteAvatarRequest) Execute() (*Avatar, *http.Response, error) {
	return r.ApiService.DeleteAvatarExecute(r)
}

/*
DeleteAvatar Delete Avatar

Delete an avatar. Notice an avatar is never fully "deleted", only its ReleaseStatus is set to "hidden" and the linked Files are deleted. The AvatarID is permanently reserved.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param avatarId Must be a valid avatar ID.
 @return ApiDeleteAvatarRequest
*/
func (a *AvatarsApiService) DeleteAvatar(ctx context.Context, avatarId string) ApiDeleteAvatarRequest {
	return ApiDeleteAvatarRequest{
		ApiService: a,
		ctx: ctx,
		avatarId: avatarId,
	}
}

// Execute executes the request
//  @return Avatar
func (a *AvatarsApiService) DeleteAvatarExecute(r ApiDeleteAvatarRequest) (*Avatar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Avatar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvatarsApiService.DeleteAvatar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/avatars/{avatarId}"
	localVarPath = strings.Replace(localVarPath, "{"+"avatarId"+"}", url.PathEscape(parameterValueToString(r.avatarId, "avatarId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAvatarRequest struct {
	ctx context.Context
	ApiService *AvatarsApiService
	avatarId string
}

func (r ApiGetAvatarRequest) Execute() (*Avatar, *http.Response, error) {
	return r.ApiService.GetAvatarExecute(r)
}

/*
GetAvatar Get Avatar

Get information about a specific Avatar.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param avatarId Must be a valid avatar ID.
 @return ApiGetAvatarRequest
*/
func (a *AvatarsApiService) GetAvatar(ctx context.Context, avatarId string) ApiGetAvatarRequest {
	return ApiGetAvatarRequest{
		ApiService: a,
		ctx: ctx,
		avatarId: avatarId,
	}
}

// Execute executes the request
//  @return Avatar
func (a *AvatarsApiService) GetAvatarExecute(r ApiGetAvatarRequest) (*Avatar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Avatar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvatarsApiService.GetAvatar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/avatars/{avatarId}"
	localVarPath = strings.Replace(localVarPath, "{"+"avatarId"+"}", url.PathEscape(parameterValueToString(r.avatarId, "avatarId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFavoritedAvatarsRequest struct {
	ctx context.Context
	ApiService *AvatarsApiService
	featured *bool
	sort *SortOption
	n *int32
	order *OrderOption
	offset *int32
	search *string
	tag *string
	notag *string
	releaseStatus *ReleaseStatus
	maxUnityVersion *string
	minUnityVersion *string
	platform *string
	userId *string
}

// Filters on featured results.
func (r ApiGetFavoritedAvatarsRequest) Featured(featured bool) ApiGetFavoritedAvatarsRequest {
	r.featured = &featured
	return r
}

// The sort order of the results.
func (r ApiGetFavoritedAvatarsRequest) Sort(sort SortOption) ApiGetFavoritedAvatarsRequest {
	r.sort = &sort
	return r
}

// The number of objects to return.
func (r ApiGetFavoritedAvatarsRequest) N(n int32) ApiGetFavoritedAvatarsRequest {
	r.n = &n
	return r
}

// Result ordering
func (r ApiGetFavoritedAvatarsRequest) Order(order OrderOption) ApiGetFavoritedAvatarsRequest {
	r.order = &order
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiGetFavoritedAvatarsRequest) Offset(offset int32) ApiGetFavoritedAvatarsRequest {
	r.offset = &offset
	return r
}

// Filters by world name.
func (r ApiGetFavoritedAvatarsRequest) Search(search string) ApiGetFavoritedAvatarsRequest {
	r.search = &search
	return r
}

// Tags to include (comma-separated). Any of the tags needs to be present.
func (r ApiGetFavoritedAvatarsRequest) Tag(tag string) ApiGetFavoritedAvatarsRequest {
	r.tag = &tag
	return r
}

// Tags to exclude (comma-separated).
func (r ApiGetFavoritedAvatarsRequest) Notag(notag string) ApiGetFavoritedAvatarsRequest {
	r.notag = &notag
	return r
}

// Filter by ReleaseStatus.
func (r ApiGetFavoritedAvatarsRequest) ReleaseStatus(releaseStatus ReleaseStatus) ApiGetFavoritedAvatarsRequest {
	r.releaseStatus = &releaseStatus
	return r
}

// The maximum Unity version supported by the asset.
func (r ApiGetFavoritedAvatarsRequest) MaxUnityVersion(maxUnityVersion string) ApiGetFavoritedAvatarsRequest {
	r.maxUnityVersion = &maxUnityVersion
	return r
}

// The minimum Unity version supported by the asset.
func (r ApiGetFavoritedAvatarsRequest) MinUnityVersion(minUnityVersion string) ApiGetFavoritedAvatarsRequest {
	r.minUnityVersion = &minUnityVersion
	return r
}

// The platform the asset supports.
func (r ApiGetFavoritedAvatarsRequest) Platform(platform string) ApiGetFavoritedAvatarsRequest {
	r.platform = &platform
	return r
}

// Target user to see information on, admin-only.
func (r ApiGetFavoritedAvatarsRequest) UserId(userId string) ApiGetFavoritedAvatarsRequest {
	r.userId = &userId
	return r
}

func (r ApiGetFavoritedAvatarsRequest) Execute() ([]Avatar, *http.Response, error) {
	return r.ApiService.GetFavoritedAvatarsExecute(r)
}

/*
GetFavoritedAvatars List Favorited Avatars

Search and list favorited avatars by query filters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFavoritedAvatarsRequest
*/
func (a *AvatarsApiService) GetFavoritedAvatars(ctx context.Context) ApiGetFavoritedAvatarsRequest {
	return ApiGetFavoritedAvatarsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Avatar
func (a *AvatarsApiService) GetFavoritedAvatarsExecute(r ApiGetFavoritedAvatarsRequest) ([]Avatar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Avatar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvatarsApiService.GetFavoritedAvatars")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/avatars/favorites"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.featured != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "featured", r.featured, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	if r.n != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "n", r.n, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "")
	}
	if r.tag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tag", r.tag, "")
	}
	if r.notag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "notag", r.notag, "")
	}
	if r.releaseStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "releaseStatus", r.releaseStatus, "")
	}
	if r.maxUnityVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxUnityVersion", r.maxUnityVersion, "")
	}
	if r.minUnityVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "minUnityVersion", r.minUnityVersion, "")
	}
	if r.platform != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "platform", r.platform, "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOwnAvatarRequest struct {
	ctx context.Context
	ApiService *AvatarsApiService
	userId string
}

func (r ApiGetOwnAvatarRequest) Execute() (*Avatar, *http.Response, error) {
	return r.ApiService.GetOwnAvatarExecute(r)
}

/*
GetOwnAvatar Get Own Avatar

Get the current avatar for the user. This will return an error for any other user than the one logged in.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId Must be a valid user ID.
 @return ApiGetOwnAvatarRequest
*/
func (a *AvatarsApiService) GetOwnAvatar(ctx context.Context, userId string) ApiGetOwnAvatarRequest {
	return ApiGetOwnAvatarRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return Avatar
func (a *AvatarsApiService) GetOwnAvatarExecute(r ApiGetOwnAvatarRequest) (*Avatar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Avatar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvatarsApiService.GetOwnAvatar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{userId}/avatar"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSearchAvatarsRequest struct {
	ctx context.Context
	ApiService *AvatarsApiService
	featured *bool
	sort *SortOption
	user *string
	userId *string
	n *int32
	order *OrderOption
	offset *int32
	tag *string
	notag *string
	releaseStatus *ReleaseStatus
	maxUnityVersion *string
	minUnityVersion *string
	platform *string
}

// Filters on featured results.
func (r ApiSearchAvatarsRequest) Featured(featured bool) ApiSearchAvatarsRequest {
	r.featured = &featured
	return r
}

// The sort order of the results.
func (r ApiSearchAvatarsRequest) Sort(sort SortOption) ApiSearchAvatarsRequest {
	r.sort = &sort
	return r
}

// Set to &#x60;me&#x60; for searching own avatars.
func (r ApiSearchAvatarsRequest) User(user string) ApiSearchAvatarsRequest {
	r.user = &user
	return r
}

// Filter by UserID.
func (r ApiSearchAvatarsRequest) UserId(userId string) ApiSearchAvatarsRequest {
	r.userId = &userId
	return r
}

// The number of objects to return.
func (r ApiSearchAvatarsRequest) N(n int32) ApiSearchAvatarsRequest {
	r.n = &n
	return r
}

// Result ordering
func (r ApiSearchAvatarsRequest) Order(order OrderOption) ApiSearchAvatarsRequest {
	r.order = &order
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiSearchAvatarsRequest) Offset(offset int32) ApiSearchAvatarsRequest {
	r.offset = &offset
	return r
}

// Tags to include (comma-separated). Any of the tags needs to be present.
func (r ApiSearchAvatarsRequest) Tag(tag string) ApiSearchAvatarsRequest {
	r.tag = &tag
	return r
}

// Tags to exclude (comma-separated).
func (r ApiSearchAvatarsRequest) Notag(notag string) ApiSearchAvatarsRequest {
	r.notag = &notag
	return r
}

// Filter by ReleaseStatus.
func (r ApiSearchAvatarsRequest) ReleaseStatus(releaseStatus ReleaseStatus) ApiSearchAvatarsRequest {
	r.releaseStatus = &releaseStatus
	return r
}

// The maximum Unity version supported by the asset.
func (r ApiSearchAvatarsRequest) MaxUnityVersion(maxUnityVersion string) ApiSearchAvatarsRequest {
	r.maxUnityVersion = &maxUnityVersion
	return r
}

// The minimum Unity version supported by the asset.
func (r ApiSearchAvatarsRequest) MinUnityVersion(minUnityVersion string) ApiSearchAvatarsRequest {
	r.minUnityVersion = &minUnityVersion
	return r
}

// The platform the asset supports.
func (r ApiSearchAvatarsRequest) Platform(platform string) ApiSearchAvatarsRequest {
	r.platform = &platform
	return r
}

func (r ApiSearchAvatarsRequest) Execute() ([]Avatar, *http.Response, error) {
	return r.ApiService.SearchAvatarsExecute(r)
}

/*
SearchAvatars Search Avatars

Search and list avatars by query filters. You can only search your own or featured avatars. It is not possible as a normal user to search other peoples avatars.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchAvatarsRequest
*/
func (a *AvatarsApiService) SearchAvatars(ctx context.Context) ApiSearchAvatarsRequest {
	return ApiSearchAvatarsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Avatar
func (a *AvatarsApiService) SearchAvatarsExecute(r ApiSearchAvatarsRequest) ([]Avatar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Avatar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvatarsApiService.SearchAvatars")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/avatars"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.featured != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "featured", r.featured, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	if r.user != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user", r.user, "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
	}
	if r.n != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "n", r.n, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.tag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tag", r.tag, "")
	}
	if r.notag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "notag", r.notag, "")
	}
	if r.releaseStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "releaseStatus", r.releaseStatus, "")
	}
	if r.maxUnityVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxUnityVersion", r.maxUnityVersion, "")
	}
	if r.minUnityVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "minUnityVersion", r.minUnityVersion, "")
	}
	if r.platform != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "platform", r.platform, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSelectAvatarRequest struct {
	ctx context.Context
	ApiService *AvatarsApiService
	avatarId string
}

func (r ApiSelectAvatarRequest) Execute() (*CurrentUser, *http.Response, error) {
	return r.ApiService.SelectAvatarExecute(r)
}

/*
SelectAvatar Select Avatar

Switches into that avatar.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param avatarId Must be a valid avatar ID.
 @return ApiSelectAvatarRequest
*/
func (a *AvatarsApiService) SelectAvatar(ctx context.Context, avatarId string) ApiSelectAvatarRequest {
	return ApiSelectAvatarRequest{
		ApiService: a,
		ctx: ctx,
		avatarId: avatarId,
	}
}

// Execute executes the request
//  @return CurrentUser
func (a *AvatarsApiService) SelectAvatarExecute(r ApiSelectAvatarRequest) (*CurrentUser, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CurrentUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvatarsApiService.SelectAvatar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/avatars/{avatarId}/select"
	localVarPath = strings.Replace(localVarPath, "{"+"avatarId"+"}", url.PathEscape(parameterValueToString(r.avatarId, "avatarId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSelectFallbackAvatarRequest struct {
	ctx context.Context
	ApiService *AvatarsApiService
	avatarId string
}

func (r ApiSelectFallbackAvatarRequest) Execute() (*CurrentUser, *http.Response, error) {
	return r.ApiService.SelectFallbackAvatarExecute(r)
}

/*
SelectFallbackAvatar Select Fallback Avatar

Switches into that avatar as your fallback avatar.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param avatarId Must be a valid avatar ID.
 @return ApiSelectFallbackAvatarRequest
*/
func (a *AvatarsApiService) SelectFallbackAvatar(ctx context.Context, avatarId string) ApiSelectFallbackAvatarRequest {
	return ApiSelectFallbackAvatarRequest{
		ApiService: a,
		ctx: ctx,
		avatarId: avatarId,
	}
}

// Execute executes the request
//  @return CurrentUser
func (a *AvatarsApiService) SelectFallbackAvatarExecute(r ApiSelectFallbackAvatarRequest) (*CurrentUser, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CurrentUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvatarsApiService.SelectFallbackAvatar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/avatars/{avatarId}/selectFallback"
	localVarPath = strings.Replace(localVarPath, "{"+"avatarId"+"}", url.PathEscape(parameterValueToString(r.avatarId, "avatarId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateAvatarRequest struct {
	ctx context.Context
	ApiService *AvatarsApiService
	avatarId string
	updateAvatarRequest *UpdateAvatarRequest
}

func (r ApiUpdateAvatarRequest) UpdateAvatarRequest(updateAvatarRequest UpdateAvatarRequest) ApiUpdateAvatarRequest {
	r.updateAvatarRequest = &updateAvatarRequest
	return r
}

func (r ApiUpdateAvatarRequest) Execute() (*Avatar, *http.Response, error) {
	return r.ApiService.UpdateAvatarExecute(r)
}

/*
UpdateAvatar Update Avatar

Update information about a specific avatar.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param avatarId Must be a valid avatar ID.
 @return ApiUpdateAvatarRequest
*/
func (a *AvatarsApiService) UpdateAvatar(ctx context.Context, avatarId string) ApiUpdateAvatarRequest {
	return ApiUpdateAvatarRequest{
		ApiService: a,
		ctx: ctx,
		avatarId: avatarId,
	}
}

// Execute executes the request
//  @return Avatar
func (a *AvatarsApiService) UpdateAvatarExecute(r ApiUpdateAvatarRequest) (*Avatar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Avatar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvatarsApiService.UpdateAvatar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/avatars/{avatarId}"
	localVarPath = strings.Replace(localVarPath, "{"+"avatarId"+"}", url.PathEscape(parameterValueToString(r.avatarId, "avatarId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateAvatarRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
