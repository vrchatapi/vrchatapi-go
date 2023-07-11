# \AvatarsApi

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAvatar**](AvatarsApi.md#CreateAvatar) | **Post** /avatars | Create Avatar
[**DeleteAvatar**](AvatarsApi.md#DeleteAvatar) | **Delete** /avatars/{avatarId} | Delete Avatar
[**GetAvatar**](AvatarsApi.md#GetAvatar) | **Get** /avatars/{avatarId} | Get Avatar
[**GetFavoritedAvatars**](AvatarsApi.md#GetFavoritedAvatars) | **Get** /avatars/favorites | List Favorited Avatars
[**GetOwnAvatar**](AvatarsApi.md#GetOwnAvatar) | **Get** /users/{userId}/avatar | Get Own Avatar
[**SearchAvatars**](AvatarsApi.md#SearchAvatars) | **Get** /avatars | Search Avatars
[**SelectAvatar**](AvatarsApi.md#SelectAvatar) | **Put** /avatars/{avatarId}/select | Select Avatar
[**SelectFallbackAvatar**](AvatarsApi.md#SelectFallbackAvatar) | **Put** /avatars/{avatarId}/selectFallback | Select Fallback Avatar
[**UpdateAvatar**](AvatarsApi.md#UpdateAvatar) | **Put** /avatars/{avatarId} | Update Avatar



## CreateAvatar

> Avatar CreateAvatar(ctx).CreateAvatarRequest(createAvatarRequest).Execute()

Create Avatar



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
    createAvatarRequest := *openapiclient.NewCreateAvatarRequest("Name_example", "ImageUrl_example") // CreateAvatarRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvatarsApi.CreateAvatar(context.Background()).CreateAvatarRequest(createAvatarRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvatarsApi.CreateAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAvatar`: Avatar
    fmt.Fprintf(os.Stdout, "Response from `AvatarsApi.CreateAvatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAvatarRequest** | [**CreateAvatarRequest**](CreateAvatarRequest.md) |  | 

### Return type

[**Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAvatar

> Avatar DeleteAvatar(ctx, avatarId).Execute()

Delete Avatar



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
    avatarId := "avatarId_example" // string | Must be a valid avatar ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvatarsApi.DeleteAvatar(context.Background(), avatarId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvatarsApi.DeleteAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAvatar`: Avatar
    fmt.Fprintf(os.Stdout, "Response from `AvatarsApi.DeleteAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avatarId** | **string** | Must be a valid avatar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvatar

> Avatar GetAvatar(ctx, avatarId).Execute()

Get Avatar



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
    avatarId := "avatarId_example" // string | Must be a valid avatar ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvatarsApi.GetAvatar(context.Background(), avatarId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvatarsApi.GetAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvatar`: Avatar
    fmt.Fprintf(os.Stdout, "Response from `AvatarsApi.GetAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avatarId** | **string** | Must be a valid avatar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavoritedAvatars

> []Avatar GetFavoritedAvatars(ctx).Featured(featured).Sort(sort).N(n).Order(order).Offset(offset).Search(search).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).UserId(userId).Execute()

List Favorited Avatars



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
    featured := true // bool | Filters on featured results. (optional)
    sort := openapiclient.SortOption("popularity") // SortOption | The sort order of the results. (optional) (default to "popularity")
    n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
    order := openapiclient.OrderOption("ascending") // OrderOption | Result ordering (optional) (default to "descending")
    offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
    search := "search_example" // string | Filters by world name. (optional)
    tag := "tag_example" // string | Tags to include (comma-separated). Any of the tags needs to be present. (optional)
    notag := "notag_example" // string | Tags to exclude (comma-separated). (optional)
    releaseStatus := openapiclient.ReleaseStatus("public") // ReleaseStatus | Filter by ReleaseStatus. (optional) (default to "public")
    maxUnityVersion := "maxUnityVersion_example" // string | The maximum Unity version supported by the asset. (optional)
    minUnityVersion := "minUnityVersion_example" // string | The minimum Unity version supported by the asset. (optional)
    platform := "platform_example" // string | The platform the asset supports. (optional)
    userId := "userId_example" // string | Target user to see information on, admin-only. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvatarsApi.GetFavoritedAvatars(context.Background()).Featured(featured).Sort(sort).N(n).Order(order).Offset(offset).Search(search).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvatarsApi.GetFavoritedAvatars``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFavoritedAvatars`: []Avatar
    fmt.Fprintf(os.Stdout, "Response from `AvatarsApi.GetFavoritedAvatars`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoritedAvatarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featured** | **bool** | Filters on featured results. | 
 **sort** | [**SortOption**](SortOption.md) | The sort order of the results. | [default to &quot;popularity&quot;]
 **n** | **int32** | The number of objects to return. | [default to 60]
 **order** | [**OrderOption**](OrderOption.md) | Result ordering | [default to &quot;descending&quot;]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **search** | **string** | Filters by world name. | 
 **tag** | **string** | Tags to include (comma-separated). Any of the tags needs to be present. | 
 **notag** | **string** | Tags to exclude (comma-separated). | 
 **releaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) | Filter by ReleaseStatus. | [default to &quot;public&quot;]
 **maxUnityVersion** | **string** | The maximum Unity version supported by the asset. | 
 **minUnityVersion** | **string** | The minimum Unity version supported by the asset. | 
 **platform** | **string** | The platform the asset supports. | 
 **userId** | **string** | Target user to see information on, admin-only. | 

### Return type

[**[]Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnAvatar

> Avatar GetOwnAvatar(ctx, userId).Execute()

Get Own Avatar



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
    userId := "userId_example" // string | Must be a valid user ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvatarsApi.GetOwnAvatar(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvatarsApi.GetOwnAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOwnAvatar`: Avatar
    fmt.Fprintf(os.Stdout, "Response from `AvatarsApi.GetOwnAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAvatars

> []Avatar SearchAvatars(ctx).Featured(featured).Sort(sort).User(user).UserId(userId).N(n).Order(order).Offset(offset).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).Execute()

Search Avatars



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
    featured := true // bool | Filters on featured results. (optional)
    sort := openapiclient.SortOption("popularity") // SortOption | The sort order of the results. (optional) (default to "popularity")
    user := "user_example" // string | Set to `me` for searching own avatars. (optional)
    userId := "userId_example" // string | Filter by UserID. (optional)
    n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
    order := openapiclient.OrderOption("ascending") // OrderOption | Result ordering (optional) (default to "descending")
    offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
    tag := "tag_example" // string | Tags to include (comma-separated). Any of the tags needs to be present. (optional)
    notag := "notag_example" // string | Tags to exclude (comma-separated). (optional)
    releaseStatus := openapiclient.ReleaseStatus("public") // ReleaseStatus | Filter by ReleaseStatus. (optional) (default to "public")
    maxUnityVersion := "maxUnityVersion_example" // string | The maximum Unity version supported by the asset. (optional)
    minUnityVersion := "minUnityVersion_example" // string | The minimum Unity version supported by the asset. (optional)
    platform := "platform_example" // string | The platform the asset supports. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvatarsApi.SearchAvatars(context.Background()).Featured(featured).Sort(sort).User(user).UserId(userId).N(n).Order(order).Offset(offset).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvatarsApi.SearchAvatars``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAvatars`: []Avatar
    fmt.Fprintf(os.Stdout, "Response from `AvatarsApi.SearchAvatars`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAvatarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featured** | **bool** | Filters on featured results. | 
 **sort** | [**SortOption**](SortOption.md) | The sort order of the results. | [default to &quot;popularity&quot;]
 **user** | **string** | Set to &#x60;me&#x60; for searching own avatars. | 
 **userId** | **string** | Filter by UserID. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **order** | [**OrderOption**](OrderOption.md) | Result ordering | [default to &quot;descending&quot;]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **tag** | **string** | Tags to include (comma-separated). Any of the tags needs to be present. | 
 **notag** | **string** | Tags to exclude (comma-separated). | 
 **releaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) | Filter by ReleaseStatus. | [default to &quot;public&quot;]
 **maxUnityVersion** | **string** | The maximum Unity version supported by the asset. | 
 **minUnityVersion** | **string** | The minimum Unity version supported by the asset. | 
 **platform** | **string** | The platform the asset supports. | 

### Return type

[**[]Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectAvatar

> CurrentUser SelectAvatar(ctx, avatarId).Execute()

Select Avatar



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
    avatarId := "avatarId_example" // string | Must be a valid avatar ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvatarsApi.SelectAvatar(context.Background(), avatarId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvatarsApi.SelectAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SelectAvatar`: CurrentUser
    fmt.Fprintf(os.Stdout, "Response from `AvatarsApi.SelectAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avatarId** | **string** | Must be a valid avatar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSelectAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CurrentUser**](CurrentUser.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectFallbackAvatar

> CurrentUser SelectFallbackAvatar(ctx, avatarId).Execute()

Select Fallback Avatar



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
    avatarId := "avatarId_example" // string | Must be a valid avatar ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvatarsApi.SelectFallbackAvatar(context.Background(), avatarId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvatarsApi.SelectFallbackAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SelectFallbackAvatar`: CurrentUser
    fmt.Fprintf(os.Stdout, "Response from `AvatarsApi.SelectFallbackAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avatarId** | **string** | Must be a valid avatar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSelectFallbackAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CurrentUser**](CurrentUser.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAvatar

> Avatar UpdateAvatar(ctx, avatarId).UpdateAvatarRequest(updateAvatarRequest).Execute()

Update Avatar



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
    avatarId := "avatarId_example" // string | Must be a valid avatar ID.
    updateAvatarRequest := *openapiclient.NewUpdateAvatarRequest() // UpdateAvatarRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvatarsApi.UpdateAvatar(context.Background(), avatarId).UpdateAvatarRequest(updateAvatarRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvatarsApi.UpdateAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAvatar`: Avatar
    fmt.Fprintf(os.Stdout, "Response from `AvatarsApi.UpdateAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avatarId** | **string** | Must be a valid avatar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAvatarRequest** | [**UpdateAvatarRequest**](UpdateAvatarRequest.md) |  | 

### Return type

[**Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

