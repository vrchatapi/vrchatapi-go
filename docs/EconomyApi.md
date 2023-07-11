# \EconomyApi

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentSubscriptions**](EconomyApi.md#GetCurrentSubscriptions) | **Get** /auth/user/subscription | Get Current Subscriptions
[**GetLicenseGroup**](EconomyApi.md#GetLicenseGroup) | **Get** /licenseGroups/{licenseGroupId} | Get License Group
[**GetSteamTransaction**](EconomyApi.md#GetSteamTransaction) | **Get** /Steam/transactions/{transactionId} | Get Steam Transaction
[**GetSteamTransactions**](EconomyApi.md#GetSteamTransactions) | **Get** /Steam/transactions | List Steam Transactions
[**GetSubscriptions**](EconomyApi.md#GetSubscriptions) | **Get** /subscriptions | List Subscriptions



## GetCurrentSubscriptions

> []UserSubscription GetCurrentSubscriptions(ctx).Execute()

Get Current Subscriptions



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
    resp, r, err := apiClient.EconomyApi.GetCurrentSubscriptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EconomyApi.GetCurrentSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentSubscriptions`: []UserSubscription
    fmt.Fprintf(os.Stdout, "Response from `EconomyApi.GetCurrentSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentSubscriptionsRequest struct via the builder pattern


### Return type

[**[]UserSubscription**](UserSubscription.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseGroup

> LicenseGroup GetLicenseGroup(ctx, licenseGroupId).Execute()

Get License Group



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
    licenseGroupId := "licenseGroupId_example" // string | Must be a valid license group ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EconomyApi.GetLicenseGroup(context.Background(), licenseGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EconomyApi.GetLicenseGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseGroup`: LicenseGroup
    fmt.Fprintf(os.Stdout, "Response from `EconomyApi.GetLicenseGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseGroupId** | **string** | Must be a valid license group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LicenseGroup**](LicenseGroup.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSteamTransaction

> Transaction GetSteamTransaction(ctx, transactionId).Execute()

Get Steam Transaction



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
    transactionId := "transactionId_example" // string | Must be a valid transaction ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EconomyApi.GetSteamTransaction(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EconomyApi.GetSteamTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSteamTransaction`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `EconomyApi.GetSteamTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Must be a valid transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSteamTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSteamTransactions

> []Transaction GetSteamTransactions(ctx).Execute()

List Steam Transactions



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
    resp, r, err := apiClient.EconomyApi.GetSteamTransactions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EconomyApi.GetSteamTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSteamTransactions`: []Transaction
    fmt.Fprintf(os.Stdout, "Response from `EconomyApi.GetSteamTransactions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSteamTransactionsRequest struct via the builder pattern


### Return type

[**[]Transaction**](Transaction.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptions

> []Subscription GetSubscriptions(ctx).Execute()

List Subscriptions



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
    resp, r, err := apiClient.EconomyApi.GetSubscriptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EconomyApi.GetSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptions`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `EconomyApi.GetSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsRequest struct via the builder pattern


### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

