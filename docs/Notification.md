# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Details** | **string** | **NOTICE:** This is not a JSON object when received from the REST API, but it is when received from the Websocket API. When received from the REST API, this is a json **encoded** object, meaning you have to json-de-encode to get the NotificationDetail object depending on the NotificationType. | [default to "{}"]
**Id** | **string** |  | 
**Message** | **string** |  | 
**Seen** | Pointer to **bool** | Not included in notification objects received from the Websocket API | [optional] [default to false]
**ReceiverUserId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**SenderUserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**SenderUsername** | Pointer to **string** | -| **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429). | [optional] 
**Type** | [**NotificationType**](NotificationType.md) |  | [default to FRIEND_REQUEST]

## Methods

### NewNotification

`func NewNotification(createdAt time.Time, details string, id string, message string, senderUserId string, type_ NotificationType, ) *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Notification) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Notification) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Notification) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDetails

`func (o *Notification) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Notification) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Notification) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetId

`func (o *Notification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Notification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Notification) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *Notification) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Notification) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Notification) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSeen

`func (o *Notification) GetSeen() bool`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *Notification) GetSeenOk() (*bool, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *Notification) SetSeen(v bool)`

SetSeen sets Seen field to given value.

### HasSeen

`func (o *Notification) HasSeen() bool`

HasSeen returns a boolean if a field has been set.

### GetReceiverUserId

`func (o *Notification) GetReceiverUserId() string`

GetReceiverUserId returns the ReceiverUserId field if non-nil, zero value otherwise.

### GetReceiverUserIdOk

`func (o *Notification) GetReceiverUserIdOk() (*string, bool)`

GetReceiverUserIdOk returns a tuple with the ReceiverUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverUserId

`func (o *Notification) SetReceiverUserId(v string)`

SetReceiverUserId sets ReceiverUserId field to given value.

### HasReceiverUserId

`func (o *Notification) HasReceiverUserId() bool`

HasReceiverUserId returns a boolean if a field has been set.

### GetSenderUserId

`func (o *Notification) GetSenderUserId() string`

GetSenderUserId returns the SenderUserId field if non-nil, zero value otherwise.

### GetSenderUserIdOk

`func (o *Notification) GetSenderUserIdOk() (*string, bool)`

GetSenderUserIdOk returns a tuple with the SenderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUserId

`func (o *Notification) SetSenderUserId(v string)`

SetSenderUserId sets SenderUserId field to given value.


### GetSenderUsername

`func (o *Notification) GetSenderUsername() string`

GetSenderUsername returns the SenderUsername field if non-nil, zero value otherwise.

### GetSenderUsernameOk

`func (o *Notification) GetSenderUsernameOk() (*string, bool)`

GetSenderUsernameOk returns a tuple with the SenderUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUsername

`func (o *Notification) SetSenderUsername(v string)`

SetSenderUsername sets SenderUsername field to given value.

### HasSenderUsername

`func (o *Notification) HasSenderUsername() bool`

HasSenderUsername returns a boolean if a field has been set.

### GetType

`func (o *Notification) GetType() NotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Notification) GetTypeOk() (*NotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Notification) SetType(v NotificationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


