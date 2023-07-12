# CreateGroupAnnouncementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Announcement title | 
**Text** | Pointer to **string** | Announcement text | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**SendNotification** | Pointer to **bool** | Send notification to group members. | [optional] [default to false]

## Methods

### NewCreateGroupAnnouncementRequest

`func NewCreateGroupAnnouncementRequest(title string, ) *CreateGroupAnnouncementRequest`

NewCreateGroupAnnouncementRequest instantiates a new CreateGroupAnnouncementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupAnnouncementRequestWithDefaults

`func NewCreateGroupAnnouncementRequestWithDefaults() *CreateGroupAnnouncementRequest`

NewCreateGroupAnnouncementRequestWithDefaults instantiates a new CreateGroupAnnouncementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateGroupAnnouncementRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateGroupAnnouncementRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateGroupAnnouncementRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetText

`func (o *CreateGroupAnnouncementRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateGroupAnnouncementRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateGroupAnnouncementRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CreateGroupAnnouncementRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetImageId

`func (o *CreateGroupAnnouncementRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateGroupAnnouncementRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateGroupAnnouncementRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CreateGroupAnnouncementRequest) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetSendNotification

`func (o *CreateGroupAnnouncementRequest) GetSendNotification() bool`

GetSendNotification returns the SendNotification field if non-nil, zero value otherwise.

### GetSendNotificationOk

`func (o *CreateGroupAnnouncementRequest) GetSendNotificationOk() (*bool, bool)`

GetSendNotificationOk returns a tuple with the SendNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotification

`func (o *CreateGroupAnnouncementRequest) SetSendNotification(v bool)`

SetSendNotification sets SendNotification field to given value.

### HasSendNotification

`func (o *CreateGroupAnnouncementRequest) HasSendNotification() bool`

HasSendNotification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


