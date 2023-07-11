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


// WorldsApiService WorldsApi service
type WorldsApiService service

type ApiCreateWorldRequest struct {
	ctx context.Context
	ApiService *WorldsApiService
	createWorldRequest *CreateWorldRequest
}

// 
func (r ApiCreateWorldRequest) CreateWorldRequest(createWorldRequest CreateWorldRequest) ApiCreateWorldRequest {
	r.createWorldRequest = &createWorldRequest
	return r
}

func (r ApiCreateWorldRequest) Execute() (*World, *http.Response, error) {
	return r.ApiService.CreateWorldExecute(r)
}

/*
CreateWorld Create World

Create a new world. This endpoint requires `assetUrl` to be a valid File object with `.vrcw` file extension, and `imageUrl` to be a valid File object with an image file extension.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateWorldRequest
*/
func (a *WorldsApiService) CreateWorld(ctx context.Context) ApiCreateWorldRequest {
	return ApiCreateWorldRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return World
func (a *WorldsApiService) CreateWorldExecute(r ApiCreateWorldRequest) (*World, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *World
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorldsApiService.CreateWorld")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/worlds"

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
	localVarPostBody = r.createWorldRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
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

type ApiDeleteWorldRequest struct {
	ctx context.Context
	ApiService *WorldsApiService
	worldId string
}

func (r ApiDeleteWorldRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteWorldExecute(r)
}

/*
DeleteWorld Delete World

Delete a world. Notice a world is never fully "deleted", only its ReleaseStatus is set to "hidden" and the linked Files are deleted. The WorldID is permanently reserved.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param worldId Must be a valid world ID.
 @return ApiDeleteWorldRequest
*/
func (a *WorldsApiService) DeleteWorld(ctx context.Context, worldId string) ApiDeleteWorldRequest {
	return ApiDeleteWorldRequest{
		ApiService: a,
		ctx: ctx,
		worldId: worldId,
	}
}

// Execute executes the request
func (a *WorldsApiService) DeleteWorldExecute(r ApiDeleteWorldRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorldsApiService.DeleteWorld")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/worlds/{worldId}"
	localVarPath = strings.Replace(localVarPath, "{"+"worldId"+"}", url.PathEscape(parameterValueToString(r.worldId, "worldId")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetActiveWorldsRequest struct {
	ctx context.Context
	ApiService *WorldsApiService
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
}

// Filters on featured results.
func (r ApiGetActiveWorldsRequest) Featured(featured bool) ApiGetActiveWorldsRequest {
	r.featured = &featured
	return r
}

// The sort order of the results.
func (r ApiGetActiveWorldsRequest) Sort(sort SortOption) ApiGetActiveWorldsRequest {
	r.sort = &sort
	return r
}

// The number of objects to return.
func (r ApiGetActiveWorldsRequest) N(n int32) ApiGetActiveWorldsRequest {
	r.n = &n
	return r
}

// Result ordering
func (r ApiGetActiveWorldsRequest) Order(order OrderOption) ApiGetActiveWorldsRequest {
	r.order = &order
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiGetActiveWorldsRequest) Offset(offset int32) ApiGetActiveWorldsRequest {
	r.offset = &offset
	return r
}

// Filters by world name.
func (r ApiGetActiveWorldsRequest) Search(search string) ApiGetActiveWorldsRequest {
	r.search = &search
	return r
}

// Tags to include (comma-separated). Any of the tags needs to be present.
func (r ApiGetActiveWorldsRequest) Tag(tag string) ApiGetActiveWorldsRequest {
	r.tag = &tag
	return r
}

// Tags to exclude (comma-separated).
func (r ApiGetActiveWorldsRequest) Notag(notag string) ApiGetActiveWorldsRequest {
	r.notag = &notag
	return r
}

// Filter by ReleaseStatus.
func (r ApiGetActiveWorldsRequest) ReleaseStatus(releaseStatus ReleaseStatus) ApiGetActiveWorldsRequest {
	r.releaseStatus = &releaseStatus
	return r
}

// The maximum Unity version supported by the asset.
func (r ApiGetActiveWorldsRequest) MaxUnityVersion(maxUnityVersion string) ApiGetActiveWorldsRequest {
	r.maxUnityVersion = &maxUnityVersion
	return r
}

// The minimum Unity version supported by the asset.
func (r ApiGetActiveWorldsRequest) MinUnityVersion(minUnityVersion string) ApiGetActiveWorldsRequest {
	r.minUnityVersion = &minUnityVersion
	return r
}

// The platform the asset supports.
func (r ApiGetActiveWorldsRequest) Platform(platform string) ApiGetActiveWorldsRequest {
	r.platform = &platform
	return r
}

func (r ApiGetActiveWorldsRequest) Execute() ([]LimitedWorld, *http.Response, error) {
	return r.ApiService.GetActiveWorldsExecute(r)
}

/*
GetActiveWorlds List Active Worlds

Search and list currently Active worlds by query filters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetActiveWorldsRequest
*/
func (a *WorldsApiService) GetActiveWorlds(ctx context.Context) ApiGetActiveWorldsRequest {
	return ApiGetActiveWorldsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []LimitedWorld
func (a *WorldsApiService) GetActiveWorldsExecute(r ApiGetActiveWorldsRequest) ([]LimitedWorld, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []LimitedWorld
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorldsApiService.GetActiveWorlds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/worlds/active"

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

type ApiGetFavoritedWorldsRequest struct {
	ctx context.Context
	ApiService *WorldsApiService
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
func (r ApiGetFavoritedWorldsRequest) Featured(featured bool) ApiGetFavoritedWorldsRequest {
	r.featured = &featured
	return r
}

// The sort order of the results.
func (r ApiGetFavoritedWorldsRequest) Sort(sort SortOption) ApiGetFavoritedWorldsRequest {
	r.sort = &sort
	return r
}

// The number of objects to return.
func (r ApiGetFavoritedWorldsRequest) N(n int32) ApiGetFavoritedWorldsRequest {
	r.n = &n
	return r
}

// Result ordering
func (r ApiGetFavoritedWorldsRequest) Order(order OrderOption) ApiGetFavoritedWorldsRequest {
	r.order = &order
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiGetFavoritedWorldsRequest) Offset(offset int32) ApiGetFavoritedWorldsRequest {
	r.offset = &offset
	return r
}

// Filters by world name.
func (r ApiGetFavoritedWorldsRequest) Search(search string) ApiGetFavoritedWorldsRequest {
	r.search = &search
	return r
}

// Tags to include (comma-separated). Any of the tags needs to be present.
func (r ApiGetFavoritedWorldsRequest) Tag(tag string) ApiGetFavoritedWorldsRequest {
	r.tag = &tag
	return r
}

// Tags to exclude (comma-separated).
func (r ApiGetFavoritedWorldsRequest) Notag(notag string) ApiGetFavoritedWorldsRequest {
	r.notag = &notag
	return r
}

// Filter by ReleaseStatus.
func (r ApiGetFavoritedWorldsRequest) ReleaseStatus(releaseStatus ReleaseStatus) ApiGetFavoritedWorldsRequest {
	r.releaseStatus = &releaseStatus
	return r
}

// The maximum Unity version supported by the asset.
func (r ApiGetFavoritedWorldsRequest) MaxUnityVersion(maxUnityVersion string) ApiGetFavoritedWorldsRequest {
	r.maxUnityVersion = &maxUnityVersion
	return r
}

// The minimum Unity version supported by the asset.
func (r ApiGetFavoritedWorldsRequest) MinUnityVersion(minUnityVersion string) ApiGetFavoritedWorldsRequest {
	r.minUnityVersion = &minUnityVersion
	return r
}

// The platform the asset supports.
func (r ApiGetFavoritedWorldsRequest) Platform(platform string) ApiGetFavoritedWorldsRequest {
	r.platform = &platform
	return r
}

// Target user to see information on, admin-only.
func (r ApiGetFavoritedWorldsRequest) UserId(userId string) ApiGetFavoritedWorldsRequest {
	r.userId = &userId
	return r
}

func (r ApiGetFavoritedWorldsRequest) Execute() ([]LimitedWorld, *http.Response, error) {
	return r.ApiService.GetFavoritedWorldsExecute(r)
}

/*
GetFavoritedWorlds List Favorited Worlds

Search and list favorited worlds by query filters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFavoritedWorldsRequest
*/
func (a *WorldsApiService) GetFavoritedWorlds(ctx context.Context) ApiGetFavoritedWorldsRequest {
	return ApiGetFavoritedWorldsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []LimitedWorld
func (a *WorldsApiService) GetFavoritedWorldsExecute(r ApiGetFavoritedWorldsRequest) ([]LimitedWorld, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []LimitedWorld
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorldsApiService.GetFavoritedWorlds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/worlds/favorites"

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

type ApiGetRecentWorldsRequest struct {
	ctx context.Context
	ApiService *WorldsApiService
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
func (r ApiGetRecentWorldsRequest) Featured(featured bool) ApiGetRecentWorldsRequest {
	r.featured = &featured
	return r
}

// The sort order of the results.
func (r ApiGetRecentWorldsRequest) Sort(sort SortOption) ApiGetRecentWorldsRequest {
	r.sort = &sort
	return r
}

// The number of objects to return.
func (r ApiGetRecentWorldsRequest) N(n int32) ApiGetRecentWorldsRequest {
	r.n = &n
	return r
}

// Result ordering
func (r ApiGetRecentWorldsRequest) Order(order OrderOption) ApiGetRecentWorldsRequest {
	r.order = &order
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiGetRecentWorldsRequest) Offset(offset int32) ApiGetRecentWorldsRequest {
	r.offset = &offset
	return r
}

// Filters by world name.
func (r ApiGetRecentWorldsRequest) Search(search string) ApiGetRecentWorldsRequest {
	r.search = &search
	return r
}

// Tags to include (comma-separated). Any of the tags needs to be present.
func (r ApiGetRecentWorldsRequest) Tag(tag string) ApiGetRecentWorldsRequest {
	r.tag = &tag
	return r
}

// Tags to exclude (comma-separated).
func (r ApiGetRecentWorldsRequest) Notag(notag string) ApiGetRecentWorldsRequest {
	r.notag = &notag
	return r
}

// Filter by ReleaseStatus.
func (r ApiGetRecentWorldsRequest) ReleaseStatus(releaseStatus ReleaseStatus) ApiGetRecentWorldsRequest {
	r.releaseStatus = &releaseStatus
	return r
}

// The maximum Unity version supported by the asset.
func (r ApiGetRecentWorldsRequest) MaxUnityVersion(maxUnityVersion string) ApiGetRecentWorldsRequest {
	r.maxUnityVersion = &maxUnityVersion
	return r
}

// The minimum Unity version supported by the asset.
func (r ApiGetRecentWorldsRequest) MinUnityVersion(minUnityVersion string) ApiGetRecentWorldsRequest {
	r.minUnityVersion = &minUnityVersion
	return r
}

// The platform the asset supports.
func (r ApiGetRecentWorldsRequest) Platform(platform string) ApiGetRecentWorldsRequest {
	r.platform = &platform
	return r
}

// Target user to see information on, admin-only.
func (r ApiGetRecentWorldsRequest) UserId(userId string) ApiGetRecentWorldsRequest {
	r.userId = &userId
	return r
}

func (r ApiGetRecentWorldsRequest) Execute() ([]LimitedWorld, *http.Response, error) {
	return r.ApiService.GetRecentWorldsExecute(r)
}

/*
GetRecentWorlds List Recent Worlds

Search and list recently visited worlds by query filters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRecentWorldsRequest
*/
func (a *WorldsApiService) GetRecentWorlds(ctx context.Context) ApiGetRecentWorldsRequest {
	return ApiGetRecentWorldsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []LimitedWorld
func (a *WorldsApiService) GetRecentWorldsExecute(r ApiGetRecentWorldsRequest) ([]LimitedWorld, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []LimitedWorld
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorldsApiService.GetRecentWorlds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/worlds/recent"

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

type ApiGetWorldRequest struct {
	ctx context.Context
	ApiService *WorldsApiService
	worldId string
}

func (r ApiGetWorldRequest) Execute() (*World, *http.Response, error) {
	return r.ApiService.GetWorldExecute(r)
}

/*
GetWorld Get World by ID

Get information about a specific World.
Works unauthenticated but when so will always return `0` for certain fields.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param worldId Must be a valid world ID.
 @return ApiGetWorldRequest
*/
func (a *WorldsApiService) GetWorld(ctx context.Context, worldId string) ApiGetWorldRequest {
	return ApiGetWorldRequest{
		ApiService: a,
		ctx: ctx,
		worldId: worldId,
	}
}

// Execute executes the request
//  @return World
func (a *WorldsApiService) GetWorldExecute(r ApiGetWorldRequest) (*World, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *World
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorldsApiService.GetWorld")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/worlds/{worldId}"
	localVarPath = strings.Replace(localVarPath, "{"+"worldId"+"}", url.PathEscape(parameterValueToString(r.worldId, "worldId")), -1)

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

type ApiGetWorldInstanceRequest struct {
	ctx context.Context
	ApiService *WorldsApiService
	worldId string
	instanceId string
}

func (r ApiGetWorldInstanceRequest) Execute() (*Instance, *http.Response, error) {
	return r.ApiService.GetWorldInstanceExecute(r)
}

/*
GetWorldInstance Get World Instance

Returns a worlds instance.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param worldId Must be a valid world ID.
 @param instanceId Must be a valid instance ID.
 @return ApiGetWorldInstanceRequest
*/
func (a *WorldsApiService) GetWorldInstance(ctx context.Context, worldId string, instanceId string) ApiGetWorldInstanceRequest {
	return ApiGetWorldInstanceRequest{
		ApiService: a,
		ctx: ctx,
		worldId: worldId,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return Instance
func (a *WorldsApiService) GetWorldInstanceExecute(r ApiGetWorldInstanceRequest) (*Instance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Instance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorldsApiService.GetWorldInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/worlds/{worldId}/{instanceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"worldId"+"}", url.PathEscape(parameterValueToString(r.worldId, "worldId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterValueToString(r.instanceId, "instanceId")), -1)

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

type ApiGetWorldMetadataRequest struct {
	ctx context.Context
	ApiService *WorldsApiService
	worldId string
}

func (r ApiGetWorldMetadataRequest) Execute() (*WorldMetadata, *http.Response, error) {
	return r.ApiService.GetWorldMetadataExecute(r)
}

/*
GetWorldMetadata Get World Metadata

Return a worlds custom metadata. This is currently believed to be unused. Metadata can be set with `updateWorld` and can be any arbitrary object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param worldId Must be a valid world ID.
 @return ApiGetWorldMetadataRequest

Deprecated
*/
func (a *WorldsApiService) GetWorldMetadata(ctx context.Context, worldId string) ApiGetWorldMetadataRequest {
	return ApiGetWorldMetadataRequest{
		ApiService: a,
		ctx: ctx,
		worldId: worldId,
	}
}

// Execute executes the request
//  @return WorldMetadata
// Deprecated
func (a *WorldsApiService) GetWorldMetadataExecute(r ApiGetWorldMetadataRequest) (*WorldMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorldMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorldsApiService.GetWorldMetadata")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/worlds/{worldId}/metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"worldId"+"}", url.PathEscape(parameterValueToString(r.worldId, "worldId")), -1)

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

type ApiGetWorldPublishStatusRequest struct {
	ctx context.Context
	ApiService *WorldsApiService
	worldId string
}

func (r ApiGetWorldPublishStatusRequest) Execute() (*WorldPublishStatus, *http.Response, error) {
	return r.ApiService.GetWorldPublishStatusExecute(r)
}

/*
GetWorldPublishStatus Get World Publish Status

Returns a worlds publish status.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param worldId Must be a valid world ID.
 @return ApiGetWorldPublishStatusRequest
*/
func (a *WorldsApiService) GetWorldPublishStatus(ctx context.Context, worldId string) ApiGetWorldPublishStatusRequest {
	return ApiGetWorldPublishStatusRequest{
		ApiService: a,
		ctx: ctx,
		worldId: worldId,
	}
}

// Execute executes the request
//  @return WorldPublishStatus
func (a *WorldsApiService) GetWorldPublishStatusExecute(r ApiGetWorldPublishStatusRequest) (*WorldPublishStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorldPublishStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorldsApiService.GetWorldPublishStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/worlds/{worldId}/publish"
	localVarPath = strings.Replace(localVarPath, "{"+"worldId"+"}", url.PathEscape(parameterValueToString(r.worldId, "worldId")), -1)

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

type ApiPublishWorldRequest struct {
	ctx context.Context
	ApiService *WorldsApiService
	worldId string
}

func (r ApiPublishWorldRequest) Execute() (*http.Response, error) {
	return r.ApiService.PublishWorldExecute(r)
}

/*
PublishWorld Publish World

Publish a world. You can only publish one world per week.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param worldId Must be a valid world ID.
 @return ApiPublishWorldRequest
*/
func (a *WorldsApiService) PublishWorld(ctx context.Context, worldId string) ApiPublishWorldRequest {
	return ApiPublishWorldRequest{
		ApiService: a,
		ctx: ctx,
		worldId: worldId,
	}
}

// Execute executes the request
func (a *WorldsApiService) PublishWorldExecute(r ApiPublishWorldRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorldsApiService.PublishWorld")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/worlds/{worldId}/publish"
	localVarPath = strings.Replace(localVarPath, "{"+"worldId"+"}", url.PathEscape(parameterValueToString(r.worldId, "worldId")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSearchWorldsRequest struct {
	ctx context.Context
	ApiService *WorldsApiService
	featured *bool
	sort *SortOption
	user *string
	userId *string
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
}

// Filters on featured results.
func (r ApiSearchWorldsRequest) Featured(featured bool) ApiSearchWorldsRequest {
	r.featured = &featured
	return r
}

// The sort order of the results.
func (r ApiSearchWorldsRequest) Sort(sort SortOption) ApiSearchWorldsRequest {
	r.sort = &sort
	return r
}

// Set to &#x60;me&#x60; for searching own worlds.
func (r ApiSearchWorldsRequest) User(user string) ApiSearchWorldsRequest {
	r.user = &user
	return r
}

// Filter by UserID.
func (r ApiSearchWorldsRequest) UserId(userId string) ApiSearchWorldsRequest {
	r.userId = &userId
	return r
}

// The number of objects to return.
func (r ApiSearchWorldsRequest) N(n int32) ApiSearchWorldsRequest {
	r.n = &n
	return r
}

// Result ordering
func (r ApiSearchWorldsRequest) Order(order OrderOption) ApiSearchWorldsRequest {
	r.order = &order
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiSearchWorldsRequest) Offset(offset int32) ApiSearchWorldsRequest {
	r.offset = &offset
	return r
}

// Filters by world name.
func (r ApiSearchWorldsRequest) Search(search string) ApiSearchWorldsRequest {
	r.search = &search
	return r
}

// Tags to include (comma-separated). Any of the tags needs to be present.
func (r ApiSearchWorldsRequest) Tag(tag string) ApiSearchWorldsRequest {
	r.tag = &tag
	return r
}

// Tags to exclude (comma-separated).
func (r ApiSearchWorldsRequest) Notag(notag string) ApiSearchWorldsRequest {
	r.notag = &notag
	return r
}

// Filter by ReleaseStatus.
func (r ApiSearchWorldsRequest) ReleaseStatus(releaseStatus ReleaseStatus) ApiSearchWorldsRequest {
	r.releaseStatus = &releaseStatus
	return r
}

// The maximum Unity version supported by the asset.
func (r ApiSearchWorldsRequest) MaxUnityVersion(maxUnityVersion string) ApiSearchWorldsRequest {
	r.maxUnityVersion = &maxUnityVersion
	return r
}

// The minimum Unity version supported by the asset.
func (r ApiSearchWorldsRequest) MinUnityVersion(minUnityVersion string) ApiSearchWorldsRequest {
	r.minUnityVersion = &minUnityVersion
	return r
}

// The platform the asset supports.
func (r ApiSearchWorldsRequest) Platform(platform string) ApiSearchWorldsRequest {
	r.platform = &platform
	return r
}

func (r ApiSearchWorldsRequest) Execute() ([]LimitedWorld, *http.Response, error) {
	return r.ApiService.SearchWorldsExecute(r)
}

/*
SearchWorlds Search All Worlds

Search and list any worlds by query filters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchWorldsRequest
*/
func (a *WorldsApiService) SearchWorlds(ctx context.Context) ApiSearchWorldsRequest {
	return ApiSearchWorldsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []LimitedWorld
func (a *WorldsApiService) SearchWorldsExecute(r ApiSearchWorldsRequest) ([]LimitedWorld, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []LimitedWorld
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorldsApiService.SearchWorlds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/worlds"

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

type ApiUnpublishWorldRequest struct {
	ctx context.Context
	ApiService *WorldsApiService
	worldId string
}

func (r ApiUnpublishWorldRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnpublishWorldExecute(r)
}

/*
UnpublishWorld Unpublish World

Unpublish a world.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param worldId Must be a valid world ID.
 @return ApiUnpublishWorldRequest
*/
func (a *WorldsApiService) UnpublishWorld(ctx context.Context, worldId string) ApiUnpublishWorldRequest {
	return ApiUnpublishWorldRequest{
		ApiService: a,
		ctx: ctx,
		worldId: worldId,
	}
}

// Execute executes the request
func (a *WorldsApiService) UnpublishWorldExecute(r ApiUnpublishWorldRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorldsApiService.UnpublishWorld")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/worlds/{worldId}/publish"
	localVarPath = strings.Replace(localVarPath, "{"+"worldId"+"}", url.PathEscape(parameterValueToString(r.worldId, "worldId")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateWorldRequest struct {
	ctx context.Context
	ApiService *WorldsApiService
	worldId string
	updateWorldRequest *UpdateWorldRequest
}

func (r ApiUpdateWorldRequest) UpdateWorldRequest(updateWorldRequest UpdateWorldRequest) ApiUpdateWorldRequest {
	r.updateWorldRequest = &updateWorldRequest
	return r
}

func (r ApiUpdateWorldRequest) Execute() (*World, *http.Response, error) {
	return r.ApiService.UpdateWorldExecute(r)
}

/*
UpdateWorld Update World

Update information about a specific World.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param worldId Must be a valid world ID.
 @return ApiUpdateWorldRequest
*/
func (a *WorldsApiService) UpdateWorld(ctx context.Context, worldId string) ApiUpdateWorldRequest {
	return ApiUpdateWorldRequest{
		ApiService: a,
		ctx: ctx,
		worldId: worldId,
	}
}

// Execute executes the request
//  @return World
func (a *WorldsApiService) UpdateWorldExecute(r ApiUpdateWorldRequest) (*World, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *World
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorldsApiService.UpdateWorld")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/worlds/{worldId}"
	localVarPath = strings.Replace(localVarPath, "{"+"worldId"+"}", url.PathEscape(parameterValueToString(r.worldId, "worldId")), -1)

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
	localVarPostBody = r.updateWorldRequest
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
