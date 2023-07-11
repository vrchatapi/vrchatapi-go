# \FilesApi

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFile**](FilesApi.md#CreateFile) | **Post** /file | Create File
[**CreateFileVersion**](FilesApi.md#CreateFileVersion) | **Post** /file/{fileId} | Create File Version
[**DeleteFile**](FilesApi.md#DeleteFile) | **Delete** /file/{fileId} | Delete File
[**DeleteFileVersion**](FilesApi.md#DeleteFileVersion) | **Delete** /file/{fileId}/{versionId} | Delete File Version
[**DownloadFileVersion**](FilesApi.md#DownloadFileVersion) | **Get** /file/{fileId}/{versionId} | Download File Version
[**FinishFileDataUpload**](FilesApi.md#FinishFileDataUpload) | **Put** /file/{fileId}/{versionId}/{fileType}/finish | Finish FileData Upload
[**GetFile**](FilesApi.md#GetFile) | **Get** /file/{fileId} | Show File
[**GetFileDataUploadStatus**](FilesApi.md#GetFileDataUploadStatus) | **Get** /file/{fileId}/{versionId}/{fileType}/status | Check FileData Upload Status
[**GetFiles**](FilesApi.md#GetFiles) | **Get** /files | List Files
[**StartFileDataUpload**](FilesApi.md#StartFileDataUpload) | **Put** /file/{fileId}/{versionId}/{fileType}/start | Start FileData Upload



## CreateFile

> File CreateFile(ctx).CreateFileRequest(createFileRequest).Execute()

Create File



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vrchatapi/vrchatapi-go/vrchatapi"
)

func main() {
    createFileRequest := *openapiclient.NewCreateFileRequest("Name_example", openapiclient.MIMEType("image/jpeg"), "Extension_example") // CreateFileRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.CreateFile(context.Background()).CreateFileRequest(createFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.CreateFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFile`: File
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.CreateFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFileRequest** | [**CreateFileRequest**](CreateFileRequest.md) |  | 

### Return type

[**File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFileVersion

> File CreateFileVersion(ctx, fileId).CreateFileVersionRequest(createFileVersionRequest).Execute()

Create File Version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vrchatapi/vrchatapi-go/vrchatapi"
)

func main() {
    fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
    createFileVersionRequest := *openapiclient.NewCreateFileVersionRequest("SignatureMd5_example", float32(123)) // CreateFileVersionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.CreateFileVersion(context.Background(), fileId).CreateFileVersionRequest(createFileVersionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.CreateFileVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFileVersion`: File
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.CreateFileVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFileVersionRequest** | [**CreateFileVersionRequest**](CreateFileVersionRequest.md) |  | 

### Return type

[**File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFile

> Success DeleteFile(ctx, fileId).Execute()

Delete File



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vrchatapi/vrchatapi-go/vrchatapi"
)

func main() {
    fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.DeleteFile(context.Background(), fileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.DeleteFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFile`: Success
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.DeleteFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFileVersion

> File DeleteFileVersion(ctx, fileId, versionId).Execute()

Delete File Version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vrchatapi/vrchatapi-go/vrchatapi"
)

func main() {
    fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
    versionId := int32(1) // int32 | Version ID of the asset.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.DeleteFileVersion(context.Background(), fileId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.DeleteFileVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFileVersion`: File
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.DeleteFileVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 
**versionId** | **int32** | Version ID of the asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadFileVersion

> DownloadFileVersion(ctx, fileId, versionId).Execute()

Download File Version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vrchatapi/vrchatapi-go/vrchatapi"
)

func main() {
    fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
    versionId := int32(1) // int32 | Version ID of the asset.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FilesApi.DownloadFileVersion(context.Background(), fileId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.DownloadFileVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 
**versionId** | **int32** | Version ID of the asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinishFileDataUpload

> File FinishFileDataUpload(ctx, fileId, versionId, fileType).FinishFileDataUploadRequest(finishFileDataUploadRequest).Execute()

Finish FileData Upload



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vrchatapi/vrchatapi-go/vrchatapi"
)

func main() {
    fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
    versionId := int32(1) // int32 | Version ID of the asset.
    fileType := "file" // string | Type of file.
    finishFileDataUploadRequest := *openapiclient.NewFinishFileDataUploadRequest("0", "0") // FinishFileDataUploadRequest | Please see documentation on ETag's: [https://teppen.io/2018/06/23/aws_s3_etags/](https://teppen.io/2018/06/23/aws_s3_etags/)  ETag's should NOT be present when uploading a `signature`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.FinishFileDataUpload(context.Background(), fileId, versionId, fileType).FinishFileDataUploadRequest(finishFileDataUploadRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.FinishFileDataUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinishFileDataUpload`: File
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.FinishFileDataUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 
**versionId** | **int32** | Version ID of the asset. | 
**fileType** | **string** | Type of file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinishFileDataUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **finishFileDataUploadRequest** | [**FinishFileDataUploadRequest**](FinishFileDataUploadRequest.md) | Please see documentation on ETag&#39;s: [https://teppen.io/2018/06/23/aws_s3_etags/](https://teppen.io/2018/06/23/aws_s3_etags/)  ETag&#39;s should NOT be present when uploading a &#x60;signature&#x60;. | 

### Return type

[**File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFile

> File GetFile(ctx, fileId).Execute()

Show File



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vrchatapi/vrchatapi-go/vrchatapi"
)

func main() {
    fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.GetFile(context.Background(), fileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GetFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFile`: File
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileDataUploadStatus

> FileVersionUploadStatus GetFileDataUploadStatus(ctx, fileId, versionId, fileType).Execute()

Check FileData Upload Status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vrchatapi/vrchatapi-go/vrchatapi"
)

func main() {
    fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
    versionId := int32(1) // int32 | Version ID of the asset.
    fileType := "file" // string | Type of file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.GetFileDataUploadStatus(context.Background(), fileId, versionId, fileType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GetFileDataUploadStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFileDataUploadStatus`: FileVersionUploadStatus
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GetFileDataUploadStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 
**versionId** | **int32** | Version ID of the asset. | 
**fileType** | **string** | Type of file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileDataUploadStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FileVersionUploadStatus**](FileVersionUploadStatus.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiles

> []File GetFiles(ctx).Tag(tag).UserId(userId).N(n).Offset(offset).Execute()

List Files



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vrchatapi/vrchatapi-go/vrchatapi"
)

func main() {
    tag := "tag_example" // string | Tag, for example \"icon\" or \"gallery\", not included by default. (optional)
    userId := "userId_example" // string | UserID, will always generate a 500 permission error. (optional)
    n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
    offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.GetFiles(context.Background()).Tag(tag).UserId(userId).N(n).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GetFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFiles`: []File
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GetFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **string** | Tag, for example \&quot;icon\&quot; or \&quot;gallery\&quot;, not included by default. | 
 **userId** | **string** | UserID, will always generate a 500 permission error. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**[]File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartFileDataUpload

> FileUploadURL StartFileDataUpload(ctx, fileId, versionId, fileType).PartNumber(partNumber).Execute()

Start FileData Upload



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vrchatapi/vrchatapi-go/vrchatapi"
)

func main() {
    fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
    versionId := int32(1) // int32 | Version ID of the asset.
    fileType := "file" // string | Type of file.
    partNumber := int32(1) // int32 | The part number to start uploading. If not provided, the first part will be started. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.StartFileDataUpload(context.Background(), fileId, versionId, fileType).PartNumber(partNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.StartFileDataUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartFileDataUpload`: FileUploadURL
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.StartFileDataUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 
**versionId** | **int32** | Version ID of the asset. | 
**fileType** | **string** | Type of file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartFileDataUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **partNumber** | **int32** | The part number to start uploading. If not provided, the first part will be started. | 

### Return type

[**FileUploadURL**](FileUploadURL.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

