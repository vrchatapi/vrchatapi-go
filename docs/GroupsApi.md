# \GroupsApi

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupGalleryImage**](GroupsApi.md#AddGroupGalleryImage) | **Post** /groups/{groupId}/galleries/{groupGalleryId}/images | Add Group Gallery Image
[**AddGroupMemberRole**](GroupsApi.md#AddGroupMemberRole) | **Put** /groups/{groupId}/members/{userId}/roles/{groupRoleId} | Add Role to GroupMember
[**BanGroupMember**](GroupsApi.md#BanGroupMember) | **Post** /groups/{groupId}/bans | Ban Group Member
[**CancelGroupRequest**](GroupsApi.md#CancelGroupRequest) | **Delete** /groups/{groupId}/requests | Cancel Group Join Request
[**CreateGroup**](GroupsApi.md#CreateGroup) | **Post** /groups | Create Group
[**CreateGroupAnnouncement**](GroupsApi.md#CreateGroupAnnouncement) | **Post** /groups/{groupId}/announcement | Create Group Announcement
[**CreateGroupGallery**](GroupsApi.md#CreateGroupGallery) | **Post** /groups/{groupId}/galleries | Create Group Gallery
[**CreateGroupInvite**](GroupsApi.md#CreateGroupInvite) | **Post** /groups/{groupId}/invites | Invite User to Group
[**CreateGroupRole**](GroupsApi.md#CreateGroupRole) | **Post** /groups/{groupId}/roles | Create GroupRole
[**DeleteGroup**](GroupsApi.md#DeleteGroup) | **Delete** /groups/{groupId} | Delete Group
[**DeleteGroupAnnouncement**](GroupsApi.md#DeleteGroupAnnouncement) | **Delete** /groups/{groupId}/announcement | Delete Group Announcement
[**DeleteGroupGallery**](GroupsApi.md#DeleteGroupGallery) | **Delete** /groups/{groupId}/galleries/{groupGalleryId} | Delete Group Gallery
[**DeleteGroupGalleryImage**](GroupsApi.md#DeleteGroupGalleryImage) | **Delete** /groups/{groupId}/galleries/{groupGalleryId}/images/{groupGalleryImageId} | Delete Group Gallery Image
[**DeleteGroupInvite**](GroupsApi.md#DeleteGroupInvite) | **Delete** /groups/{groupId}/invites/{userId} | Delete User Invite
[**DeleteGroupRole**](GroupsApi.md#DeleteGroupRole) | **Delete** /groups/{groupId}/roles/{groupRoleId} | Delete Group Role
[**GetGroup**](GroupsApi.md#GetGroup) | **Get** /groups/{groupId} | Get Group by ID
[**GetGroupAnnouncements**](GroupsApi.md#GetGroupAnnouncements) | **Get** /groups/{groupId}/announcement | Get Group Announcement
[**GetGroupAuditLogs**](GroupsApi.md#GetGroupAuditLogs) | **Get** /groups/{groupId}/auditLogs | Get Group Audit Logs
[**GetGroupBans**](GroupsApi.md#GetGroupBans) | **Get** /groups/{groupId}/bans | Get Group Bans
[**GetGroupGalleryImages**](GroupsApi.md#GetGroupGalleryImages) | **Get** /groups/{groupId}/galleries/{groupGalleryId} | Get Group Gallery Images
[**GetGroupInvites**](GroupsApi.md#GetGroupInvites) | **Get** /groups/{groupId}/invites | Get Group Invites Sent
[**GetGroupMember**](GroupsApi.md#GetGroupMember) | **Get** /groups/{groupId}/members/{userId} | Get Group Member
[**GetGroupMembers**](GroupsApi.md#GetGroupMembers) | **Get** /groups/{groupId}/members | List Group Members
[**GetGroupPermissions**](GroupsApi.md#GetGroupPermissions) | **Get** /groups/{groupId}/permissions | List Group Permissions
[**GetGroupRequests**](GroupsApi.md#GetGroupRequests) | **Get** /groups/{groupId}/requests | Get Group Join Requests
[**GetGroupRoles**](GroupsApi.md#GetGroupRoles) | **Get** /groups/{groupId}/roles | Get Group Roles
[**JoinGroup**](GroupsApi.md#JoinGroup) | **Post** /groups/{groupId}/join | Join Group
[**KickGroupMember**](GroupsApi.md#KickGroupMember) | **Delete** /groups/{groupId}/members/{userId} | Kick Group Member
[**LeaveGroup**](GroupsApi.md#LeaveGroup) | **Post** /groups/{groupId}/leave | Leave Group
[**RemoveGroupMemberRole**](GroupsApi.md#RemoveGroupMemberRole) | **Delete** /groups/{groupId}/members/{userId}/roles/{groupRoleId} | Remove Role from GroupMember
[**RespondGroupJoinRequest**](GroupsApi.md#RespondGroupJoinRequest) | **Put** /groups/{groupId}/requests/{userId} | Respond Group Join request
[**UnbanGroupMember**](GroupsApi.md#UnbanGroupMember) | **Delete** /groups/{groupId}/bans/{userId} | Unban Group Member
[**UpdateGroup**](GroupsApi.md#UpdateGroup) | **Put** /groups/{groupId} | Update Group
[**UpdateGroupGallery**](GroupsApi.md#UpdateGroupGallery) | **Put** /groups/{groupId}/galleries/{groupGalleryId} | Update Group Gallery
[**UpdateGroupMember**](GroupsApi.md#UpdateGroupMember) | **Put** /groups/{groupId}/members/{userId} | Update Group Member
[**UpdateGroupRole**](GroupsApi.md#UpdateGroupRole) | **Put** /groups/{groupId}/roles/{groupRoleId} | Update Group Role



## AddGroupGalleryImage

> GroupGalleryImage AddGroupGalleryImage(ctx, groupId, groupGalleryId).AddGroupGalleryImageRequest(addGroupGalleryImageRequest).Execute()

Add Group Gallery Image



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    groupGalleryId := "ggal_00000000-0000-0000-0000-000000000000" // string | Must be a valid group gallery ID.
    addGroupGalleryImageRequest := *openapiclient.NewAddGroupGalleryImageRequest("file_ce35d830-e20a-4df0-a6d4-5aaef4508044") // AddGroupGalleryImageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.AddGroupGalleryImage(context.Background(), groupId, groupGalleryId).AddGroupGalleryImageRequest(addGroupGalleryImageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.AddGroupGalleryImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGroupGalleryImage`: GroupGalleryImage
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.AddGroupGalleryImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**groupGalleryId** | **string** | Must be a valid group gallery ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupGalleryImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addGroupGalleryImageRequest** | [**AddGroupGalleryImageRequest**](AddGroupGalleryImageRequest.md) |  | 

### Return type

[**GroupGalleryImage**](GroupGalleryImage.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddGroupMemberRole

> []string AddGroupMemberRole(ctx, groupId, userId, groupRoleId).Execute()

Add Role to GroupMember



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    userId := "userId_example" // string | Must be a valid user ID.
    groupRoleId := "grol_00000000-0000-0000-0000-000000000000" // string | Must be a valid group role ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.AddGroupMemberRole(context.Background(), groupId, userId, groupRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.AddGroupMemberRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGroupMemberRole`: []string
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.AddGroupMemberRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 
**groupRoleId** | **string** | Must be a valid group role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupMemberRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**[]string**

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BanGroupMember

> GroupMember BanGroupMember(ctx, groupId).BanGroupMemberRequest(banGroupMemberRequest).Execute()

Ban Group Member



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    banGroupMemberRequest := *openapiclient.NewBanGroupMemberRequest("usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469") // BanGroupMemberRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.BanGroupMember(context.Background(), groupId).BanGroupMemberRequest(banGroupMemberRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.BanGroupMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BanGroupMember`: GroupMember
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.BanGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBanGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **banGroupMemberRequest** | [**BanGroupMemberRequest**](BanGroupMemberRequest.md) |  | 

### Return type

[**GroupMember**](GroupMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelGroupRequest

> CancelGroupRequest(ctx, groupId).Execute()

Cancel Group Join Request



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupsApi.CancelGroupRequest(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.CancelGroupRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelGroupRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> Group CreateGroup(ctx).CreateGroupRequest(createGroupRequest).Execute()

Create Group



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
    createGroupRequest := *openapiclient.NewCreateGroupRequest("Name_example", "ShortCode_example", openapiclient.GroupRoleTemplate("default")) // CreateGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.CreateGroup(context.Background()).CreateGroupRequest(createGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGroupRequest** | [**CreateGroupRequest**](CreateGroupRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupAnnouncement

> GroupAnnouncement CreateGroupAnnouncement(ctx, groupId).CreateGroupAnnouncementRequest(createGroupAnnouncementRequest).Execute()

Create Group Announcement



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    createGroupAnnouncementRequest := *openapiclient.NewCreateGroupAnnouncementRequest("Event is starting soon!") // CreateGroupAnnouncementRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.CreateGroupAnnouncement(context.Background(), groupId).CreateGroupAnnouncementRequest(createGroupAnnouncementRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.CreateGroupAnnouncement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroupAnnouncement`: GroupAnnouncement
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.CreateGroupAnnouncement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupAnnouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGroupAnnouncementRequest** | [**CreateGroupAnnouncementRequest**](CreateGroupAnnouncementRequest.md) |  | 

### Return type

[**GroupAnnouncement**](GroupAnnouncement.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupGallery

> GroupGallery CreateGroupGallery(ctx, groupId).CreateGroupGalleryRequest(createGroupGalleryRequest).Execute()

Create Group Gallery



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    createGroupGalleryRequest := *openapiclient.NewCreateGroupGalleryRequest("Example Gallery") // CreateGroupGalleryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.CreateGroupGallery(context.Background(), groupId).CreateGroupGalleryRequest(createGroupGalleryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.CreateGroupGallery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroupGallery`: GroupGallery
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.CreateGroupGallery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupGalleryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGroupGalleryRequest** | [**CreateGroupGalleryRequest**](CreateGroupGalleryRequest.md) |  | 

### Return type

[**GroupGallery**](GroupGallery.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupInvite

> CreateGroupInvite(ctx, groupId).CreateGroupInviteRequest(createGroupInviteRequest).Execute()

Invite User to Group



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    createGroupInviteRequest := *openapiclient.NewCreateGroupInviteRequest("usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469") // CreateGroupInviteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupsApi.CreateGroupInvite(context.Background(), groupId).CreateGroupInviteRequest(createGroupInviteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.CreateGroupInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGroupInviteRequest** | [**CreateGroupInviteRequest**](CreateGroupInviteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupRole

> GroupRole CreateGroupRole(ctx, groupId).CreateGroupRoleRequest(createGroupRoleRequest).Execute()

Create GroupRole



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    createGroupRoleRequest := *openapiclient.NewCreateGroupRoleRequest() // CreateGroupRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.CreateGroupRole(context.Background(), groupId).CreateGroupRoleRequest(createGroupRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.CreateGroupRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroupRole`: GroupRole
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.CreateGroupRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGroupRoleRequest** | [**CreateGroupRoleRequest**](CreateGroupRoleRequest.md) |  | 

### Return type

[**GroupRole**](GroupRole.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> Success DeleteGroup(ctx, groupId).Execute()

Delete Group



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.DeleteGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.DeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGroup`: Success
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.DeleteGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## DeleteGroupAnnouncement

> Success DeleteGroupAnnouncement(ctx, groupId).Execute()

Delete Group Announcement



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.DeleteGroupAnnouncement(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.DeleteGroupAnnouncement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGroupAnnouncement`: Success
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.DeleteGroupAnnouncement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupAnnouncementRequest struct via the builder pattern


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


## DeleteGroupGallery

> Success DeleteGroupGallery(ctx, groupId, groupGalleryId).Execute()

Delete Group Gallery



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    groupGalleryId := "ggal_00000000-0000-0000-0000-000000000000" // string | Must be a valid group gallery ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.DeleteGroupGallery(context.Background(), groupId, groupGalleryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.DeleteGroupGallery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGroupGallery`: Success
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.DeleteGroupGallery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**groupGalleryId** | **string** | Must be a valid group gallery ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupGalleryRequest struct via the builder pattern


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


## DeleteGroupGalleryImage

> Success DeleteGroupGalleryImage(ctx, groupId, groupGalleryId, groupGalleryImageId).Execute()

Delete Group Gallery Image



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    groupGalleryId := "ggal_00000000-0000-0000-0000-000000000000" // string | Must be a valid group gallery ID.
    groupGalleryImageId := "ggim_00000000-0000-0000-0000-000000000000" // string | Must be a valid group gallery image ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.DeleteGroupGalleryImage(context.Background(), groupId, groupGalleryId, groupGalleryImageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.DeleteGroupGalleryImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGroupGalleryImage`: Success
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.DeleteGroupGalleryImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**groupGalleryId** | **string** | Must be a valid group gallery ID. | 
**groupGalleryImageId** | **string** | Must be a valid group gallery image ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupGalleryImageRequest struct via the builder pattern


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


## DeleteGroupInvite

> DeleteGroupInvite(ctx, groupId, userId).Execute()

Delete User Invite



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    userId := "userId_example" // string | Must be a valid user ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupsApi.DeleteGroupInvite(context.Background(), groupId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.DeleteGroupInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupRole

> []GroupRole DeleteGroupRole(ctx, groupId, groupRoleId).Execute()

Delete Group Role



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    groupRoleId := "grol_00000000-0000-0000-0000-000000000000" // string | Must be a valid group role ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.DeleteGroupRole(context.Background(), groupId, groupRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.DeleteGroupRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGroupRole`: []GroupRole
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.DeleteGroupRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**groupRoleId** | **string** | Must be a valid group role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GroupRole**](GroupRole.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroup

> Group GetGroup(ctx, groupId).IncludeRoles(includeRoles).Execute()

Get Group by ID



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    includeRoles := true // bool | Include roles for the Group object. Defaults to false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroup(context.Background(), groupId).IncludeRoles(includeRoles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeRoles** | **bool** | Include roles for the Group object. Defaults to false. | 

### Return type

[**Group**](Group.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupAnnouncements

> GroupAnnouncement GetGroupAnnouncements(ctx, groupId).Execute()

Get Group Announcement



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupAnnouncements(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupAnnouncements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupAnnouncements`: GroupAnnouncement
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupAnnouncements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupAnnouncementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupAnnouncement**](GroupAnnouncement.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupAuditLogs

> PaginatedGroupAuditLogEntryList GetGroupAuditLogs(ctx, groupId).N(n).Offset(offset).StartDate(startDate).EndDate(endDate).Execute()

Get Group Audit Logs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/vrchatapi/vrchatapi-go/vrchatapi"
)

func main() {
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
    offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
    startDate := time.Now() // time.Time | The start date of the search range. (optional)
    endDate := time.Now() // time.Time | The end date of the search range. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupAuditLogs(context.Background(), groupId).N(n).Offset(offset).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupAuditLogs`: PaginatedGroupAuditLogEntryList
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupAuditLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **startDate** | **time.Time** | The start date of the search range. | 
 **endDate** | **time.Time** | The end date of the search range. | 

### Return type

[**PaginatedGroupAuditLogEntryList**](PaginatedGroupAuditLogEntryList.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupBans

> []GroupMember GetGroupBans(ctx, groupId).N(n).Offset(offset).Execute()

Get Group Bans



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
    offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupBans(context.Background(), groupId).N(n).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupBans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupBans`: []GroupMember
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupBans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupBansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**[]GroupMember**](GroupMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupGalleryImages

> []GroupGalleryImage GetGroupGalleryImages(ctx, groupId, groupGalleryId).N(n).Offset(offset).Approved(approved).Execute()

Get Group Gallery Images



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    groupGalleryId := "ggal_00000000-0000-0000-0000-000000000000" // string | Must be a valid group gallery ID.
    n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
    offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
    approved := true // bool | If specified, only returns images that have been approved or not approved. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupGalleryImages(context.Background(), groupId, groupGalleryId).N(n).Offset(offset).Approved(approved).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupGalleryImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupGalleryImages`: []GroupGalleryImage
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupGalleryImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**groupGalleryId** | **string** | Must be a valid group gallery ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupGalleryImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **approved** | **bool** | If specified, only returns images that have been approved or not approved. | 

### Return type

[**[]GroupGalleryImage**](GroupGalleryImage.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupInvites

> []GroupMember GetGroupInvites(ctx, groupId).Execute()

Get Group Invites Sent



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupInvites(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupInvites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupInvites`: []GroupMember
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GroupMember**](GroupMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMember

> GroupLimitedMember GetGroupMember(ctx, groupId, userId).Execute()

Get Group Member



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    userId := "userId_example" // string | Must be a valid user ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupMember(context.Background(), groupId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupMember`: GroupLimitedMember
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupLimitedMember**](GroupLimitedMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMembers

> []GroupMember GetGroupMembers(ctx, groupId).N(n).Offset(offset).Execute()

List Group Members



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
    offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupMembers(context.Background(), groupId).N(n).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupMembers`: []GroupMember
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**[]GroupMember**](GroupMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupPermissions

> []GroupPermission GetGroupPermissions(ctx, groupId).Execute()

List Group Permissions



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupPermissions(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupPermissions`: []GroupPermission
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GroupPermission**](GroupPermission.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupRequests

> []GroupMember GetGroupRequests(ctx, groupId).Execute()

Get Group Join Requests



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupRequests(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupRequests`: []GroupMember
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GroupMember**](GroupMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupRoles

> []GroupRole GetGroupRoles(ctx, groupId).Execute()

Get Group Roles



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupRoles(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupRoles`: []GroupRole
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GroupRole**](GroupRole.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JoinGroup

> GroupMember JoinGroup(ctx, groupId).Execute()

Join Group



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.JoinGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.JoinGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JoinGroup`: GroupMember
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.JoinGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJoinGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupMember**](GroupMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KickGroupMember

> KickGroupMember(ctx, groupId, userId).Execute()

Kick Group Member



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    userId := "userId_example" // string | Must be a valid user ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupsApi.KickGroupMember(context.Background(), groupId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.KickGroupMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKickGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeaveGroup

> LeaveGroup(ctx, groupId).Execute()

Leave Group



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupsApi.LeaveGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.LeaveGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeaveGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveGroupMemberRole

> []string RemoveGroupMemberRole(ctx, groupId, userId, groupRoleId).Execute()

Remove Role from GroupMember



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    userId := "userId_example" // string | Must be a valid user ID.
    groupRoleId := "grol_00000000-0000-0000-0000-000000000000" // string | Must be a valid group role ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.RemoveGroupMemberRole(context.Background(), groupId, userId, groupRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.RemoveGroupMemberRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveGroupMemberRole`: []string
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.RemoveGroupMemberRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 
**groupRoleId** | **string** | Must be a valid group role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupMemberRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**[]string**

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RespondGroupJoinRequest

> RespondGroupJoinRequest(ctx, groupId, userId).RespondGroupJoinRequest(respondGroupJoinRequest).Execute()

Respond Group Join request



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    userId := "userId_example" // string | Must be a valid user ID.
    respondGroupJoinRequest := *openapiclient.NewRespondGroupJoinRequest() // RespondGroupJoinRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupsApi.RespondGroupJoinRequest(context.Background(), groupId, userId).RespondGroupJoinRequest(respondGroupJoinRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.RespondGroupJoinRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRespondGroupJoinRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **respondGroupJoinRequest** | [**RespondGroupJoinRequest**](RespondGroupJoinRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbanGroupMember

> GroupMember UnbanGroupMember(ctx, groupId, userId).Execute()

Unban Group Member



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    userId := "userId_example" // string | Must be a valid user ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.UnbanGroupMember(context.Background(), groupId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.UnbanGroupMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnbanGroupMember`: GroupMember
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.UnbanGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbanGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupMember**](GroupMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> Group UpdateGroup(ctx, groupId).UpdateGroupRequest(updateGroupRequest).Execute()

Update Group



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    updateGroupRequest := *openapiclient.NewUpdateGroupRequest() // UpdateGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.UpdateGroup(context.Background(), groupId).UpdateGroupRequest(updateGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.UpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGroupRequest** | [**UpdateGroupRequest**](UpdateGroupRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupGallery

> GroupGallery UpdateGroupGallery(ctx, groupId, groupGalleryId).UpdateGroupGalleryRequest(updateGroupGalleryRequest).Execute()

Update Group Gallery



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    groupGalleryId := "ggal_00000000-0000-0000-0000-000000000000" // string | Must be a valid group gallery ID.
    updateGroupGalleryRequest := *openapiclient.NewUpdateGroupGalleryRequest() // UpdateGroupGalleryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.UpdateGroupGallery(context.Background(), groupId, groupGalleryId).UpdateGroupGalleryRequest(updateGroupGalleryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.UpdateGroupGallery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroupGallery`: GroupGallery
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.UpdateGroupGallery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**groupGalleryId** | **string** | Must be a valid group gallery ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupGalleryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGroupGalleryRequest** | [**UpdateGroupGalleryRequest**](UpdateGroupGalleryRequest.md) |  | 

### Return type

[**GroupGallery**](GroupGallery.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupMember

> GroupLimitedMember UpdateGroupMember(ctx, groupId, userId).UpdateGroupMemberRequest(updateGroupMemberRequest).Execute()

Update Group Member



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    userId := "userId_example" // string | Must be a valid user ID.
    updateGroupMemberRequest := *openapiclient.NewUpdateGroupMemberRequest() // UpdateGroupMemberRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.UpdateGroupMember(context.Background(), groupId, userId).UpdateGroupMemberRequest(updateGroupMemberRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.UpdateGroupMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroupMember`: GroupLimitedMember
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.UpdateGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGroupMemberRequest** | [**UpdateGroupMemberRequest**](UpdateGroupMemberRequest.md) |  | 

### Return type

[**GroupLimitedMember**](GroupLimitedMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupRole

> []GroupRole UpdateGroupRole(ctx, groupId, groupRoleId).UpdateGroupRoleRequest(updateGroupRoleRequest).Execute()

Update Group Role



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
    groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
    groupRoleId := "grol_00000000-0000-0000-0000-000000000000" // string | Must be a valid group role ID.
    updateGroupRoleRequest := *openapiclient.NewUpdateGroupRoleRequest() // UpdateGroupRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.UpdateGroupRole(context.Background(), groupId, groupRoleId).UpdateGroupRoleRequest(updateGroupRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.UpdateGroupRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroupRole`: []GroupRole
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.UpdateGroupRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**groupRoleId** | **string** | Must be a valid group role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGroupRoleRequest** | [**UpdateGroupRoleRequest**](UpdateGroupRoleRequest.md) |  | 

### Return type

[**[]GroupRole**](GroupRole.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

