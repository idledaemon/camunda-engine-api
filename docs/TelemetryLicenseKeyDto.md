# TelemetryLicenseKeyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to **NullableString** | The name of the customer the license was issued for. | [optional] 
**Type** | Pointer to **NullableString** | The license type. | [optional] 
**ValidUntil** | Pointer to **NullableString** | The expiration date of the license. | [optional] 
**Unlimited** | Pointer to **NullableBool** | Flag that indicates if the license is unlimited. | [optional] 
**Features** | Pointer to **map[string]string** | A map of features included in the license. | [optional] 
**Raw** | Pointer to **NullableString** | The raw license information. | [optional] 

## Methods

### NewTelemetryLicenseKeyDto

`func NewTelemetryLicenseKeyDto() *TelemetryLicenseKeyDto`

NewTelemetryLicenseKeyDto instantiates a new TelemetryLicenseKeyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryLicenseKeyDtoWithDefaults

`func NewTelemetryLicenseKeyDtoWithDefaults() *TelemetryLicenseKeyDto`

NewTelemetryLicenseKeyDtoWithDefaults instantiates a new TelemetryLicenseKeyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *TelemetryLicenseKeyDto) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *TelemetryLicenseKeyDto) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *TelemetryLicenseKeyDto) SetCustomer(v string)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *TelemetryLicenseKeyDto) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### SetCustomerNil

`func (o *TelemetryLicenseKeyDto) SetCustomerNil(b bool)`

 SetCustomerNil sets the value for Customer to be an explicit nil

### UnsetCustomer
`func (o *TelemetryLicenseKeyDto) UnsetCustomer()`

UnsetCustomer ensures that no value is present for Customer, not even an explicit nil
### GetType

`func (o *TelemetryLicenseKeyDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryLicenseKeyDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryLicenseKeyDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TelemetryLicenseKeyDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TelemetryLicenseKeyDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TelemetryLicenseKeyDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValidUntil

`func (o *TelemetryLicenseKeyDto) GetValidUntil() string`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *TelemetryLicenseKeyDto) GetValidUntilOk() (*string, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *TelemetryLicenseKeyDto) SetValidUntil(v string)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *TelemetryLicenseKeyDto) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### SetValidUntilNil

`func (o *TelemetryLicenseKeyDto) SetValidUntilNil(b bool)`

 SetValidUntilNil sets the value for ValidUntil to be an explicit nil

### UnsetValidUntil
`func (o *TelemetryLicenseKeyDto) UnsetValidUntil()`

UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil
### GetUnlimited

`func (o *TelemetryLicenseKeyDto) GetUnlimited() bool`

GetUnlimited returns the Unlimited field if non-nil, zero value otherwise.

### GetUnlimitedOk

`func (o *TelemetryLicenseKeyDto) GetUnlimitedOk() (*bool, bool)`

GetUnlimitedOk returns a tuple with the Unlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlimited

`func (o *TelemetryLicenseKeyDto) SetUnlimited(v bool)`

SetUnlimited sets Unlimited field to given value.

### HasUnlimited

`func (o *TelemetryLicenseKeyDto) HasUnlimited() bool`

HasUnlimited returns a boolean if a field has been set.

### SetUnlimitedNil

`func (o *TelemetryLicenseKeyDto) SetUnlimitedNil(b bool)`

 SetUnlimitedNil sets the value for Unlimited to be an explicit nil

### UnsetUnlimited
`func (o *TelemetryLicenseKeyDto) UnsetUnlimited()`

UnsetUnlimited ensures that no value is present for Unlimited, not even an explicit nil
### GetFeatures

`func (o *TelemetryLicenseKeyDto) GetFeatures() map[string]string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *TelemetryLicenseKeyDto) GetFeaturesOk() (*map[string]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *TelemetryLicenseKeyDto) SetFeatures(v map[string]string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *TelemetryLicenseKeyDto) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetRaw

`func (o *TelemetryLicenseKeyDto) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *TelemetryLicenseKeyDto) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *TelemetryLicenseKeyDto) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *TelemetryLicenseKeyDto) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### SetRawNil

`func (o *TelemetryLicenseKeyDto) SetRawNil(b bool)`

 SetRawNil sets the value for Raw to be an explicit nil

### UnsetRaw
`func (o *TelemetryLicenseKeyDto) UnsetRaw()`

UnsetRaw ensures that no value is present for Raw, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


