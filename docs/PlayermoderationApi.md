# \PlayermoderationApi

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearAllPlayerModerations**](PlayermoderationApi.md#ClearAllPlayerModerations) | **Delete** /auth/user/playermoderations | Clear All Player Moderations
[**DeletePlayerModeration**](PlayermoderationApi.md#DeletePlayerModeration) | **Delete** /auth/user/playermoderations/{playerModerationId} | Delete Player Moderation
[**GetPlayerModeration**](PlayermoderationApi.md#GetPlayerModeration) | **Get** /auth/user/playermoderations/{playerModerationId} | Get Player Moderation
[**GetPlayerModerations**](PlayermoderationApi.md#GetPlayerModerations) | **Get** /auth/user/playermoderations | Search Player Moderations
[**ModerateUser**](PlayermoderationApi.md#ModerateUser) | **Post** /auth/user/playermoderations | Moderate User
[**UnmoderateUser**](PlayermoderationApi.md#UnmoderateUser) | **Put** /auth/user/unplayermoderate | Unmoderate User



## ClearAllPlayerModerations

> Success ClearAllPlayerModerations(ctx).Execute()

Clear All Player Moderations



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
    resp, r, err := apiClient.PlayermoderationApi.ClearAllPlayerModerations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayermoderationApi.ClearAllPlayerModerations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClearAllPlayerModerations`: Success
    fmt.Fprintf(os.Stdout, "Response from `PlayermoderationApi.ClearAllPlayerModerations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClearAllPlayerModerationsRequest struct via the builder pattern


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


## DeletePlayerModeration

> Success DeletePlayerModeration(ctx, playerModerationId).Execute()

Delete Player Moderation



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
    playerModerationId := "playerModerationId_example" // string | Must be a valid `pmod_` ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayermoderationApi.DeletePlayerModeration(context.Background(), playerModerationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayermoderationApi.DeletePlayerModeration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePlayerModeration`: Success
    fmt.Fprintf(os.Stdout, "Response from `PlayermoderationApi.DeletePlayerModeration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerModerationId** | **string** | Must be a valid &#x60;pmod_&#x60; ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlayerModerationRequest struct via the builder pattern


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


## GetPlayerModeration

> PlayerModeration GetPlayerModeration(ctx, playerModerationId).Execute()

Get Player Moderation



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
    playerModerationId := "playerModerationId_example" // string | Must be a valid `pmod_` ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayermoderationApi.GetPlayerModeration(context.Background(), playerModerationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayermoderationApi.GetPlayerModeration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayerModeration`: PlayerModeration
    fmt.Fprintf(os.Stdout, "Response from `PlayermoderationApi.GetPlayerModeration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerModerationId** | **string** | Must be a valid &#x60;pmod_&#x60; ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerModerationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlayerModeration**](PlayerModeration.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerModerations

> []PlayerModeration GetPlayerModerations(ctx).Type_(type_).TargetUserId(targetUserId).Execute()

Search Player Moderations



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
    type_ := "type__example" // string | Must be one of PlayerModerationType, except unblock. Unblocking simply removes a block. (optional)
    targetUserId := "targetUserId_example" // string | Must be valid UserID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayermoderationApi.GetPlayerModerations(context.Background()).Type_(type_).TargetUserId(targetUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayermoderationApi.GetPlayerModerations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayerModerations`: []PlayerModeration
    fmt.Fprintf(os.Stdout, "Response from `PlayermoderationApi.GetPlayerModerations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerModerationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Must be one of PlayerModerationType, except unblock. Unblocking simply removes a block. | 
 **targetUserId** | **string** | Must be valid UserID. | 

### Return type

[**[]PlayerModeration**](PlayerModeration.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModerateUser

> PlayerModeration ModerateUser(ctx).ModerateUserRequest(moderateUserRequest).Execute()

Moderate User



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
    moderateUserRequest := *openapiclient.NewModerateUserRequest("usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469", openapiclient.PlayerModerationType("mute")) // ModerateUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayermoderationApi.ModerateUser(context.Background()).ModerateUserRequest(moderateUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayermoderationApi.ModerateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModerateUser`: PlayerModeration
    fmt.Fprintf(os.Stdout, "Response from `PlayermoderationApi.ModerateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModerateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moderateUserRequest** | [**ModerateUserRequest**](ModerateUserRequest.md) |  | 

### Return type

[**PlayerModeration**](PlayerModeration.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnmoderateUser

> Success UnmoderateUser(ctx).ModerateUserRequest(moderateUserRequest).Execute()

Unmoderate User



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
    moderateUserRequest := *openapiclient.NewModerateUserRequest("usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469", openapiclient.PlayerModerationType("mute")) // ModerateUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayermoderationApi.UnmoderateUser(context.Background()).ModerateUserRequest(moderateUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayermoderationApi.UnmoderateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnmoderateUser`: Success
    fmt.Fprintf(os.Stdout, "Response from `PlayermoderationApi.UnmoderateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnmoderateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moderateUserRequest** | [**ModerateUserRequest**](ModerateUserRequest.md) |  | 

### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

