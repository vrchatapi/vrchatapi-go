# \PermissionsApi

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAssignedPermissions**](PermissionsApi.md#GetAssignedPermissions) | **Get** /auth/permissions | Get Assigned Permissions
[**GetPermission**](PermissionsApi.md#GetPermission) | **Get** /permissions/{permissionId} | Get Permission



## GetAssignedPermissions

> []Permission GetAssignedPermissions(ctx).Execute()

Get Assigned Permissions



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.GetAssignedPermissions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.GetAssignedPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssignedPermissions`: []Permission
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.GetAssignedPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssignedPermissionsRequest struct via the builder pattern


### Return type

[**[]Permission**](Permission.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermission

> Permission GetPermission(ctx, permissionId).Execute()

Get Permission



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
    permissionId := "permissionId_example" // string | Must be a valid permission ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.GetPermission(context.Background(), permissionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.GetPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermission`: Permission
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.GetPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionId** | **string** | Must be a valid permission ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Permission**](Permission.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

