# LicenseGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Licenses** | [**[]License**](License.md) |  | 

## Methods

### NewLicenseGroup

`func NewLicenseGroup(id string, name string, description string, licenses []License, ) *LicenseGroup`

NewLicenseGroup instantiates a new LicenseGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseGroupWithDefaults

`func NewLicenseGroupWithDefaults() *LicenseGroup`

NewLicenseGroupWithDefaults instantiates a new LicenseGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LicenseGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicenseGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicenseGroup) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LicenseGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LicenseGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LicenseGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LicenseGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LicenseGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LicenseGroup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLicenses

`func (o *LicenseGroup) GetLicenses() []License`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *LicenseGroup) GetLicensesOk() (*[]License, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *LicenseGroup) SetLicenses(v []License)`

SetLicenses sets Licenses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


