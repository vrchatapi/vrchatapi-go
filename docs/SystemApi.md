# \SystemApi

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCSS**](SystemApi.md#GetCSS) | **Get** /css/app.css | Download CSS
[**GetConfig**](SystemApi.md#GetConfig) | **Get** /config | Fetch API Config
[**GetCurrentOnlineUsers**](SystemApi.md#GetCurrentOnlineUsers) | **Get** /visits | Current Online Users
[**GetHealth**](SystemApi.md#GetHealth) | **Get** /health | Check API Health
[**GetInfoPush**](SystemApi.md#GetInfoPush) | **Get** /infoPush | Show Information Notices
[**GetJavaScript**](SystemApi.md#GetJavaScript) | **Get** /js/app.js | Download JavaScript
[**GetSystemTime**](SystemApi.md#GetSystemTime) | **Get** /time | Current System Time



## GetCSS

> string GetCSS(ctx).Variant(variant).Branch(branch).Execute()

Download CSS



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
    variant := "variant_example" // string | Specifies which `variant` of the site. Public is the end-user site, while `internal` is the staff-only site with special pages for moderation and management. (optional) (default to "public")
    branch := "branch_example" // string | Specifies which git branch the site should load frontend source code from. (optional) (default to "main")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetCSS(context.Background()).Variant(variant).Branch(branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetCSS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCSS`: string
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.GetCSS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCSSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variant** | **string** | Specifies which &#x60;variant&#x60; of the site. Public is the end-user site, while &#x60;internal&#x60; is the staff-only site with special pages for moderation and management. | [default to &quot;public&quot;]
 **branch** | **string** | Specifies which git branch the site should load frontend source code from. | [default to &quot;main&quot;]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/css, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfig

> APIConfig GetConfig(ctx).Execute()

Fetch API Config



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
    resp, r, err := apiClient.SystemApi.GetConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfig`: APIConfig
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


### Return type

[**APIConfig**](APIConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentOnlineUsers

> int32 GetCurrentOnlineUsers(ctx).Execute()

Current Online Users



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
    resp, r, err := apiClient.SystemApi.GetCurrentOnlineUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetCurrentOnlineUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentOnlineUsers`: int32
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.GetCurrentOnlineUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentOnlineUsersRequest struct via the builder pattern


### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHealth

> APIHealth GetHealth(ctx).Execute()

Check API Health



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
    resp, r, err := apiClient.SystemApi.GetHealth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealth`: APIHealth
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.GetHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthRequest struct via the builder pattern


### Return type

[**APIHealth**](APIHealth.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfoPush

> []InfoPush GetInfoPush(ctx).Require(require).Include(include).Execute()

Show Information Notices



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
    require := "require_example" // string | Tags to include (comma-separated). All of the tags needs to be present. (optional)
    include := "include_example" // string | Tags to include (comma-separated). Any of the tags needs to be present. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetInfoPush(context.Background()).Require(require).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetInfoPush``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfoPush`: []InfoPush
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.GetInfoPush`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInfoPushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **require** | **string** | Tags to include (comma-separated). All of the tags needs to be present. | 
 **include** | **string** | Tags to include (comma-separated). Any of the tags needs to be present. | 

### Return type

[**[]InfoPush**](InfoPush.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJavaScript

> string GetJavaScript(ctx).Variant(variant).Branch(branch).Execute()

Download JavaScript



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
    variant := "variant_example" // string | Specifies which `variant` of the site. Public is the end-user site, while `internal` is the staff-only site with special pages for moderation and management. (optional) (default to "public")
    branch := "branch_example" // string | Specifies which git branch the site should load frontend source code from. (optional) (default to "main")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.GetJavaScript(context.Background()).Variant(variant).Branch(branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetJavaScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJavaScript`: string
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.GetJavaScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJavaScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variant** | **string** | Specifies which &#x60;variant&#x60; of the site. Public is the end-user site, while &#x60;internal&#x60; is the staff-only site with special pages for moderation and management. | [default to &quot;public&quot;]
 **branch** | **string** | Specifies which git branch the site should load frontend source code from. | [default to &quot;main&quot;]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/javascript, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemTime

> time.Time GetSystemTime(ctx).Execute()

Current System Time



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
    resp, r, err := apiClient.SystemApi.GetSystemTime(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.GetSystemTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemTime`: time.Time
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.GetSystemTime`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemTimeRequest struct via the builder pattern


### Return type

[**time.Time**](time.Time.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

