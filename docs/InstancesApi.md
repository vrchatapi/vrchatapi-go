# \InstancesApi

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInstance**](InstancesApi.md#GetInstance) | **Get** /instances/{worldId}:{instanceId} | Get Instance
[**GetInstanceByShortName**](InstancesApi.md#GetInstanceByShortName) | **Get** /instances/s/{shortName} | Get Instance By Short Name
[**GetShortName**](InstancesApi.md#GetShortName) | **Get** /instances/{worldId}:{instanceId}/shortName | Get Instance Short Name
[**SendSelfInvite**](InstancesApi.md#SendSelfInvite) | **Post** /instances/{worldId}:{instanceId}/invite | Send Self Invite



## GetInstance

> Instance GetInstance(ctx, worldId, instanceId).Execute()

Get Instance



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
    worldId := "worldId_example" // string | Must be a valid world ID.
    instanceId := "instanceId_example" // string | Must be a valid instance ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstancesApi.GetInstance(context.Background(), worldId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstance`: Instance
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 
**instanceId** | **string** | Must be a valid instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Instance**](Instance.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceByShortName

> Instance GetInstanceByShortName(ctx, shortName).Execute()

Get Instance By Short Name



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
    shortName := "shortName_example" // string | Must be a valid instance short name.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstancesApi.GetInstanceByShortName(context.Background(), shortName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstanceByShortName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstanceByShortName`: Instance
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetInstanceByShortName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shortName** | **string** | Must be a valid instance short name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceByShortNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Instance**](Instance.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShortName

> InstanceShortNameResponse GetShortName(ctx, worldId, instanceId).Execute()

Get Instance Short Name



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
    worldId := "worldId_example" // string | Must be a valid world ID.
    instanceId := "instanceId_example" // string | Must be a valid instance ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstancesApi.GetShortName(context.Background(), worldId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetShortName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShortName`: InstanceShortNameResponse
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetShortName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 
**instanceId** | **string** | Must be a valid instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShortNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InstanceShortNameResponse**](InstanceShortNameResponse.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendSelfInvite

> Success SendSelfInvite(ctx, worldId, instanceId).Execute()

Send Self Invite



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
    worldId := "worldId_example" // string | Must be a valid world ID.
    instanceId := "instanceId_example" // string | Must be a valid instance ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstancesApi.SendSelfInvite(context.Background(), worldId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.SendSelfInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendSelfInvite`: Success
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.SendSelfInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 
**instanceId** | **string** | Must be a valid instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendSelfInviteRequest struct via the builder pattern


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

