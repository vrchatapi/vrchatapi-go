# \AuthenticationApi

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckUserExists**](AuthenticationApi.md#CheckUserExists) | **Get** /auth/exists | Check User Exists
[**DeleteUser**](AuthenticationApi.md#DeleteUser) | **Put** /users/{userId}/delete | Delete User
[**GetCurrentUser**](AuthenticationApi.md#GetCurrentUser) | **Get** /auth/user | Login and/or Get Current User Info
[**Logout**](AuthenticationApi.md#Logout) | **Put** /logout | Logout
[**Verify2FA**](AuthenticationApi.md#Verify2FA) | **Post** /auth/twofactorauth/totp/verify | Verify 2FA code
[**Verify2FAEmailCode**](AuthenticationApi.md#Verify2FAEmailCode) | **Post** /auth/twofactorauth/emailotp/verify | Verify 2FA email code
[**VerifyAuthToken**](AuthenticationApi.md#VerifyAuthToken) | **Get** /auth | Verify Auth Token
[**VerifyRecoveryCode**](AuthenticationApi.md#VerifyRecoveryCode) | **Post** /auth/twofactorauth/otp/verify | Verify 2FA code with Recovery code



## CheckUserExists

> UserExists CheckUserExists(ctx).Email(email).DisplayName(displayName).UserId(userId).ExcludeUserId(excludeUserId).Execute()

Check User Exists



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
    email := "email_example" // string | Filter by email. (optional)
    displayName := "displayName_example" // string | Filter by displayName. (optional)
    userId := "userId_example" // string | Filter by UserID. (optional)
    excludeUserId := "excludeUserId_example" // string | Exclude by UserID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.CheckUserExists(context.Background()).Email(email).DisplayName(displayName).UserId(userId).ExcludeUserId(excludeUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.CheckUserExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckUserExists`: UserExists
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.CheckUserExists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckUserExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Filter by email. | 
 **displayName** | **string** | Filter by displayName. | 
 **userId** | **string** | Filter by UserID. | 
 **excludeUserId** | **string** | Exclude by UserID. | 

### Return type

[**UserExists**](UserExists.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> CurrentUser DeleteUser(ctx, userId).Execute()

Delete User



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
    resp, r, err := apiClient.AuthenticationApi.DeleteUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUser`: CurrentUser
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## GetCurrentUser

> CurrentUser GetCurrentUser(ctx).Execute()

Login and/or Get Current User Info



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
    resp, r, err := apiClient.AuthenticationApi.GetCurrentUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.GetCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentUser`: CurrentUser
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.GetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserRequest struct via the builder pattern


### Return type

[**CurrentUser**](CurrentUser.md)

### Authorization

[authHeader](../README.md#authHeader), [twoFactorAuthCookie](../README.md#twoFactorAuthCookie), [authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> Success Logout(ctx).Execute()

Logout



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
    resp, r, err := apiClient.AuthenticationApi.Logout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.Logout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Logout`: Success
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.Logout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


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


## Verify2FA

> Verify2FAResult Verify2FA(ctx).TwoFactorAuthCode(twoFactorAuthCode).Execute()

Verify 2FA code



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
    twoFactorAuthCode := *openapiclient.NewTwoFactorAuthCode("Code_example") // TwoFactorAuthCode |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.Verify2FA(context.Background()).TwoFactorAuthCode(twoFactorAuthCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.Verify2FA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Verify2FA`: Verify2FAResult
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.Verify2FA`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerify2FARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorAuthCode** | [**TwoFactorAuthCode**](TwoFactorAuthCode.md) |  | 

### Return type

[**Verify2FAResult**](Verify2FAResult.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Verify2FAEmailCode

> Verify2FAEmailCodeResult Verify2FAEmailCode(ctx).TwoFactorEmailCode(twoFactorEmailCode).Execute()

Verify 2FA email code



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
    twoFactorEmailCode := *openapiclient.NewTwoFactorEmailCode("Code_example") // TwoFactorEmailCode |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.Verify2FAEmailCode(context.Background()).TwoFactorEmailCode(twoFactorEmailCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.Verify2FAEmailCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Verify2FAEmailCode`: Verify2FAEmailCodeResult
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.Verify2FAEmailCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerify2FAEmailCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorEmailCode** | [**TwoFactorEmailCode**](TwoFactorEmailCode.md) |  | 

### Return type

[**Verify2FAEmailCodeResult**](Verify2FAEmailCodeResult.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyAuthToken

> VerifyAuthTokenResult VerifyAuthToken(ctx).Execute()

Verify Auth Token



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
    resp, r, err := apiClient.AuthenticationApi.VerifyAuthToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.VerifyAuthToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyAuthToken`: VerifyAuthTokenResult
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.VerifyAuthToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyAuthTokenRequest struct via the builder pattern


### Return type

[**VerifyAuthTokenResult**](VerifyAuthTokenResult.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyRecoveryCode

> Verify2FAResult VerifyRecoveryCode(ctx).TwoFactorAuthCode(twoFactorAuthCode).Execute()

Verify 2FA code with Recovery code



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
    twoFactorAuthCode := *openapiclient.NewTwoFactorAuthCode("Code_example") // TwoFactorAuthCode |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.VerifyRecoveryCode(context.Background()).TwoFactorAuthCode(twoFactorAuthCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.VerifyRecoveryCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyRecoveryCode`: Verify2FAResult
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.VerifyRecoveryCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyRecoveryCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorAuthCode** | [**TwoFactorAuthCode**](TwoFactorAuthCode.md) |  | 

### Return type

[**Verify2FAResult**](Verify2FAResult.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

