# License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForId** | **string** | Either a AvatarID, LicenseGroupID, PermissionID or ProductID. This depends on the &#x60;forType&#x60; field. | 
**ForType** | [**LicenseType**](LicenseType.md) |  | [default to PERMISSION]
**ForName** | **string** |  | 
**ForAction** | [**LicenseAction**](LicenseAction.md) |  | [default to HAVE]

## Methods

### NewLicense

`func NewLicense(forId string, forType LicenseType, forName string, forAction LicenseAction, ) *License`

NewLicense instantiates a new License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseWithDefaults

`func NewLicenseWithDefaults() *License`

NewLicenseWithDefaults instantiates a new License object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForId

`func (o *License) GetForId() string`

GetForId returns the ForId field if non-nil, zero value otherwise.

### GetForIdOk

`func (o *License) GetForIdOk() (*string, bool)`

GetForIdOk returns a tuple with the ForId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForId

`func (o *License) SetForId(v string)`

SetForId sets ForId field to given value.


### GetForType

`func (o *License) GetForType() LicenseType`

GetForType returns the ForType field if non-nil, zero value otherwise.

### GetForTypeOk

`func (o *License) GetForTypeOk() (*LicenseType, bool)`

GetForTypeOk returns a tuple with the ForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForType

`func (o *License) SetForType(v LicenseType)`

SetForType sets ForType field to given value.


### GetForName

`func (o *License) GetForName() string`

GetForName returns the ForName field if non-nil, zero value otherwise.

### GetForNameOk

`func (o *License) GetForNameOk() (*string, bool)`

GetForNameOk returns a tuple with the ForName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForName

`func (o *License) SetForName(v string)`

SetForName sets ForName field to given value.


### GetForAction

`func (o *License) GetForAction() LicenseAction`

GetForAction returns the ForAction field if non-nil, zero value otherwise.

### GetForActionOk

`func (o *License) GetForActionOk() (*LicenseAction, bool)`

GetForActionOk returns a tuple with the ForAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForAction

`func (o *License) SetForAction(v LicenseAction)`

SetForAction sets ForAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


