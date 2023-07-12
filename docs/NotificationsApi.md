# \NotificationsApi

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptFriendRequest**](NotificationsApi.md#AcceptFriendRequest) | **Put** /auth/user/notifications/{notificationId}/accept | Accept Friend Request
[**ClearNotifications**](NotificationsApi.md#ClearNotifications) | **Put** /auth/user/notifications/clear | Clear All Notifications
[**DeleteNotification**](NotificationsApi.md#DeleteNotification) | **Put** /auth/user/notifications/{notificationId}/hide | Delete Notification
[**GetNotifications**](NotificationsApi.md#GetNotifications) | **Get** /auth/user/notifications | List Notifications
[**MarkNotificationAsRead**](NotificationsApi.md#MarkNotificationAsRead) | **Put** /auth/user/notifications/{notificationId}/see | Mark Notification As Read



## AcceptFriendRequest

> Success AcceptFriendRequest(ctx, notificationId).Execute()

Accept Friend Request



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
    notificationId := "notificationId_example" // string | Must be a valid notification ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.AcceptFriendRequest(context.Background(), notificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.AcceptFriendRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceptFriendRequest`: Success
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.AcceptFriendRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptFriendRequestRequest struct via the builder pattern


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


## ClearNotifications

> Success ClearNotifications(ctx).Execute()

Clear All Notifications



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
    resp, r, err := apiClient.NotificationsApi.ClearNotifications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.ClearNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClearNotifications`: Success
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.ClearNotifications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClearNotificationsRequest struct via the builder pattern


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


## DeleteNotification

> Notification DeleteNotification(ctx, notificationId).Execute()

Delete Notification



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
    notificationId := "notificationId_example" // string | Must be a valid notification ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.DeleteNotification(context.Background(), notificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.DeleteNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNotification`: Notification
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.DeleteNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Notification**](Notification.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotifications

> []Notification GetNotifications(ctx).Type_(type_).Sent(sent).Hidden(hidden).After(after).N(n).Offset(offset).Execute()

List Notifications



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
    type_ := "all" // string | Only send notifications of this type (can use `all` for all). This parameter no longer does anything, and is deprecated. (optional)
    sent := true // bool | Return notifications sent by the user. Must be false or omitted. (optional)
    hidden := true // bool | Whether to return hidden or non-hidden notifications. True only allowed on type `friendRequest`. (optional)
    after := "five_minutes_ago" // string | Only return notifications sent after this Date. Ignored if type is `friendRequest`. (optional)
    n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
    offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.GetNotifications(context.Background()).Type_(type_).Sent(sent).Hidden(hidden).After(after).N(n).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.GetNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotifications`: []Notification
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.GetNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Only send notifications of this type (can use &#x60;all&#x60; for all). This parameter no longer does anything, and is deprecated. | 
 **sent** | **bool** | Return notifications sent by the user. Must be false or omitted. | 
 **hidden** | **bool** | Whether to return hidden or non-hidden notifications. True only allowed on type &#x60;friendRequest&#x60;. | 
 **after** | **string** | Only return notifications sent after this Date. Ignored if type is &#x60;friendRequest&#x60;. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**[]Notification**](Notification.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkNotificationAsRead

> Notification MarkNotificationAsRead(ctx, notificationId).Execute()

Mark Notification As Read



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
    notificationId := "notificationId_example" // string | Must be a valid notification ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.MarkNotificationAsRead(context.Background(), notificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.MarkNotificationAsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkNotificationAsRead`: Notification
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.MarkNotificationAsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkNotificationAsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Notification**](Notification.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

