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


// FilesApiService FilesApi service
type FilesApiService service

type ApiCreateFileRequest struct {
	ctx context.Context
	ApiService *FilesApiService
	createFileRequest *CreateFileRequest
}

func (r ApiCreateFileRequest) CreateFileRequest(createFileRequest CreateFileRequest) ApiCreateFileRequest {
	r.createFileRequest = &createFileRequest
	return r
}

func (r ApiCreateFileRequest) Execute() (*File, *http.Response, error) {
	return r.ApiService.CreateFileExecute(r)
}

/*
CreateFile Create File

Creates a new File object

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateFileRequest
*/
func (a *FilesApiService) CreateFile(ctx context.Context) ApiCreateFileRequest {
	return ApiCreateFileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return File
func (a *FilesApiService) CreateFileExecute(r ApiCreateFileRequest) (*File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.CreateFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/file"

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
	localVarPostBody = r.createFileRequest
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

type ApiCreateFileVersionRequest struct {
	ctx context.Context
	ApiService *FilesApiService
	fileId string
	createFileVersionRequest *CreateFileVersionRequest
}

func (r ApiCreateFileVersionRequest) CreateFileVersionRequest(createFileVersionRequest CreateFileVersionRequest) ApiCreateFileVersionRequest {
	r.createFileVersionRequest = &createFileVersionRequest
	return r
}

func (r ApiCreateFileVersionRequest) Execute() (*File, *http.Response, error) {
	return r.ApiService.CreateFileVersionExecute(r)
}

/*
CreateFileVersion Create File Version

Creates a new FileVersion. Once a Version has been created, proceed to the `/file/{fileId}/{versionId}/file/start` endpoint to start a file upload.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileId Must be a valid file ID.
 @return ApiCreateFileVersionRequest
*/
func (a *FilesApiService) CreateFileVersion(ctx context.Context, fileId string) ApiCreateFileVersionRequest {
	return ApiCreateFileVersionRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
	}
}

// Execute executes the request
//  @return File
func (a *FilesApiService) CreateFileVersionExecute(r ApiCreateFileVersionRequest) (*File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.CreateFileVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/file/{fileId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)

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
	localVarPostBody = r.createFileVersionRequest
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

type ApiDeleteFileRequest struct {
	ctx context.Context
	ApiService *FilesApiService
	fileId string
}

func (r ApiDeleteFileRequest) Execute() (*Success, *http.Response, error) {
	return r.ApiService.DeleteFileExecute(r)
}

/*
DeleteFile Delete File

Deletes a File object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileId Must be a valid file ID.
 @return ApiDeleteFileRequest
*/
func (a *FilesApiService) DeleteFile(ctx context.Context, fileId string) ApiDeleteFileRequest {
	return ApiDeleteFileRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
	}
}

// Execute executes the request
//  @return Success
func (a *FilesApiService) DeleteFileExecute(r ApiDeleteFileRequest) (*Success, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Success
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.DeleteFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/file/{fileId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)

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

type ApiDeleteFileVersionRequest struct {
	ctx context.Context
	ApiService *FilesApiService
	fileId string
	versionId int32
}

func (r ApiDeleteFileVersionRequest) Execute() (*File, *http.Response, error) {
	return r.ApiService.DeleteFileVersionExecute(r)
}

/*
DeleteFileVersion Delete File Version

Delete a specific version of a file. You can only delete the latest version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileId Must be a valid file ID.
 @param versionId Version ID of the asset.
 @return ApiDeleteFileVersionRequest
*/
func (a *FilesApiService) DeleteFileVersion(ctx context.Context, fileId string, versionId int32) ApiDeleteFileVersionRequest {
	return ApiDeleteFileVersionRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
		versionId: versionId,
	}
}

// Execute executes the request
//  @return File
func (a *FilesApiService) DeleteFileVersionExecute(r ApiDeleteFileVersionRequest) (*File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.DeleteFileVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/file/{fileId}/{versionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"versionId"+"}", url.PathEscape(parameterValueToString(r.versionId, "versionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.versionId < 1 {
		return localVarReturnValue, nil, reportError("versionId must be greater than 1")
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
		if localVarHTTPResponse.StatusCode == 500 {
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

type ApiDownloadFileVersionRequest struct {
	ctx context.Context
	ApiService *FilesApiService
	fileId string
	versionId int32
}

func (r ApiDownloadFileVersionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DownloadFileVersionExecute(r)
}

/*
DownloadFileVersion Download File Version

Downloads the file with the provided version number.

**Version Note:** Version 0 is always when the file was created. The real data is usually always located in version 1 and up.

**Extension Note:** Files are not guaranteed to have a file extensions. UnityPackage files tends to have it, images through this endpoint do not. You are responsible for appending file extension from the `extension` field when neccesary.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileId Must be a valid file ID.
 @param versionId Version ID of the asset.
 @return ApiDownloadFileVersionRequest
*/
func (a *FilesApiService) DownloadFileVersion(ctx context.Context, fileId string, versionId int32) ApiDownloadFileVersionRequest {
	return ApiDownloadFileVersionRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
		versionId: versionId,
	}
}

// Execute executes the request
func (a *FilesApiService) DownloadFileVersionExecute(r ApiDownloadFileVersionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.DownloadFileVersion")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/file/{fileId}/{versionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"versionId"+"}", url.PathEscape(parameterValueToString(r.versionId, "versionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.versionId < 1 {
		return nil, reportError("versionId must be greater than 1")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFinishFileDataUploadRequest struct {
	ctx context.Context
	ApiService *FilesApiService
	fileId string
	versionId int32
	fileType string
	finishFileDataUploadRequest *FinishFileDataUploadRequest
}

// Please see documentation on ETag&#39;s: [https://teppen.io/2018/06/23/aws_s3_etags/](https://teppen.io/2018/06/23/aws_s3_etags/)  ETag&#39;s should NOT be present when uploading a &#x60;signature&#x60;.
func (r ApiFinishFileDataUploadRequest) FinishFileDataUploadRequest(finishFileDataUploadRequest FinishFileDataUploadRequest) ApiFinishFileDataUploadRequest {
	r.finishFileDataUploadRequest = &finishFileDataUploadRequest
	return r
}

func (r ApiFinishFileDataUploadRequest) Execute() (*File, *http.Response, error) {
	return r.ApiService.FinishFileDataUploadExecute(r)
}

/*
FinishFileDataUpload Finish FileData Upload

Finish an upload of a FileData. This will mark it as "complete". After uploading the `file` for Avatars and Worlds you then have to upload a `signature` file.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileId Must be a valid file ID.
 @param versionId Version ID of the asset.
 @param fileType Type of file.
 @return ApiFinishFileDataUploadRequest
*/
func (a *FilesApiService) FinishFileDataUpload(ctx context.Context, fileId string, versionId int32, fileType string) ApiFinishFileDataUploadRequest {
	return ApiFinishFileDataUploadRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
		versionId: versionId,
		fileType: fileType,
	}
}

// Execute executes the request
//  @return File
func (a *FilesApiService) FinishFileDataUploadExecute(r ApiFinishFileDataUploadRequest) (*File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.FinishFileDataUpload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/file/{fileId}/{versionId}/{fileType}/finish"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"versionId"+"}", url.PathEscape(parameterValueToString(r.versionId, "versionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fileType"+"}", url.PathEscape(parameterValueToString(r.fileType, "fileType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.versionId < 1 {
		return localVarReturnValue, nil, reportError("versionId must be greater than 1")
	}

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
	localVarPostBody = r.finishFileDataUploadRequest
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

type ApiGetFileRequest struct {
	ctx context.Context
	ApiService *FilesApiService
	fileId string
}

func (r ApiGetFileRequest) Execute() (*File, *http.Response, error) {
	return r.ApiService.GetFileExecute(r)
}

/*
GetFile Show File

Shows general information about the "File" object. Each File can have several "Version"'s, and each Version can have multiple real files or "Data" blobs.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileId Must be a valid file ID.
 @return ApiGetFileRequest
*/
func (a *FilesApiService) GetFile(ctx context.Context, fileId string) ApiGetFileRequest {
	return ApiGetFileRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
	}
}

// Execute executes the request
//  @return File
func (a *FilesApiService) GetFileExecute(r ApiGetFileRequest) (*File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.GetFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/file/{fileId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)

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

type ApiGetFileDataUploadStatusRequest struct {
	ctx context.Context
	ApiService *FilesApiService
	fileId string
	versionId int32
	fileType string
}

func (r ApiGetFileDataUploadStatusRequest) Execute() (*FileVersionUploadStatus, *http.Response, error) {
	return r.ApiService.GetFileDataUploadStatusExecute(r)
}

/*
GetFileDataUploadStatus Check FileData Upload Status

Retrieves the upload status for file upload. Can currently only be accessed when `status` is `waiting`. Trying to access it on a file version already uploaded currently times out.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileId Must be a valid file ID.
 @param versionId Version ID of the asset.
 @param fileType Type of file.
 @return ApiGetFileDataUploadStatusRequest
*/
func (a *FilesApiService) GetFileDataUploadStatus(ctx context.Context, fileId string, versionId int32, fileType string) ApiGetFileDataUploadStatusRequest {
	return ApiGetFileDataUploadStatusRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
		versionId: versionId,
		fileType: fileType,
	}
}

// Execute executes the request
//  @return FileVersionUploadStatus
func (a *FilesApiService) GetFileDataUploadStatusExecute(r ApiGetFileDataUploadStatusRequest) (*FileVersionUploadStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileVersionUploadStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.GetFileDataUploadStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/file/{fileId}/{versionId}/{fileType}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"versionId"+"}", url.PathEscape(parameterValueToString(r.versionId, "versionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fileType"+"}", url.PathEscape(parameterValueToString(r.fileType, "fileType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.versionId < 1 {
		return localVarReturnValue, nil, reportError("versionId must be greater than 1")
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

type ApiGetFilesRequest struct {
	ctx context.Context
	ApiService *FilesApiService
	tag *string
	userId *string
	n *int32
	offset *int32
}

// Tag, for example \&quot;icon\&quot; or \&quot;gallery\&quot;, not included by default.
func (r ApiGetFilesRequest) Tag(tag string) ApiGetFilesRequest {
	r.tag = &tag
	return r
}

// UserID, will always generate a 500 permission error.
// Deprecated
func (r ApiGetFilesRequest) UserId(userId string) ApiGetFilesRequest {
	r.userId = &userId
	return r
}

// The number of objects to return.
func (r ApiGetFilesRequest) N(n int32) ApiGetFilesRequest {
	r.n = &n
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiGetFilesRequest) Offset(offset int32) ApiGetFilesRequest {
	r.offset = &offset
	return r
}

func (r ApiGetFilesRequest) Execute() ([]File, *http.Response, error) {
	return r.ApiService.GetFilesExecute(r)
}

/*
GetFiles List Files

Returns a list of files

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFilesRequest
*/
func (a *FilesApiService) GetFiles(ctx context.Context) ApiGetFilesRequest {
	return ApiGetFilesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []File
func (a *FilesApiService) GetFilesExecute(r ApiGetFilesRequest) ([]File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.GetFiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.tag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tag", r.tag, "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
	}
	if r.n != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "n", r.n, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
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

type ApiStartFileDataUploadRequest struct {
	ctx context.Context
	ApiService *FilesApiService
	fileId string
	versionId int32
	fileType string
	partNumber *int32
}

// The part number to start uploading. If not provided, the first part will be started.
// Deprecated
func (r ApiStartFileDataUploadRequest) PartNumber(partNumber int32) ApiStartFileDataUploadRequest {
	r.partNumber = &partNumber
	return r
}

func (r ApiStartFileDataUploadRequest) Execute() (*FileUploadURL, *http.Response, error) {
	return r.ApiService.StartFileDataUploadExecute(r)
}

/*
StartFileDataUpload Start FileData Upload

Starts an upload of a specific FilePart. This endpoint will return an AWS URL which you can PUT data to. You need to call this and receive a new AWS API URL for each `partNumber`. Please see AWS's REST documentation on "PUT Object to S3" on how to upload. Once all parts has been uploaded, proceed to `/finish` endpoint.

**Note:** `nextPartNumber` seems like it is always ignored. Despite it returning 0, first partNumber is always 1.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileId Must be a valid file ID.
 @param versionId Version ID of the asset.
 @param fileType Type of file.
 @return ApiStartFileDataUploadRequest
*/
func (a *FilesApiService) StartFileDataUpload(ctx context.Context, fileId string, versionId int32, fileType string) ApiStartFileDataUploadRequest {
	return ApiStartFileDataUploadRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
		versionId: versionId,
		fileType: fileType,
	}
}

// Execute executes the request
//  @return FileUploadURL
func (a *FilesApiService) StartFileDataUploadExecute(r ApiStartFileDataUploadRequest) (*FileUploadURL, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileUploadURL
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiService.StartFileDataUpload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/file/{fileId}/{versionId}/{fileType}/start"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"versionId"+"}", url.PathEscape(parameterValueToString(r.versionId, "versionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fileType"+"}", url.PathEscape(parameterValueToString(r.fileType, "fileType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.versionId < 1 {
		return localVarReturnValue, nil, reportError("versionId must be greater than 1")
	}

	if r.partNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "partNumber", r.partNumber, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
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
