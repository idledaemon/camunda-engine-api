# TelemetryDataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Installation** | Pointer to **NullableString** | An id which is unique for each installation of Camunda. It is stored once per database so all engines connected to the same database will have the same installation ID. The ID is used to identify a single installation of Camunda Platform. | [optional] 
**Product** | Pointer to [**TelemetryProductDto**](TelemetryProductDto.md) |  | [optional] 

## Methods

### NewTelemetryDataDto

`func NewTelemetryDataDto() *TelemetryDataDto`

NewTelemetryDataDto instantiates a new TelemetryDataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDataDtoWithDefaults

`func NewTelemetryDataDtoWithDefaults() *TelemetryDataDto`

NewTelemetryDataDtoWithDefaults instantiates a new TelemetryDataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallation

`func (o *TelemetryDataDto) GetInstallation() string`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *TelemetryDataDto) GetInstallationOk() (*string, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *TelemetryDataDto) SetInstallation(v string)`

SetInstallation sets Installation field to given value.

### HasInstallation

`func (o *TelemetryDataDto) HasInstallation() bool`

HasInstallation returns a boolean if a field has been set.

### SetInstallationNil

`func (o *TelemetryDataDto) SetInstallationNil(b bool)`

 SetInstallationNil sets the value for Installation to be an explicit nil

### UnsetInstallation
`func (o *TelemetryDataDto) UnsetInstallation()`

UnsetInstallation ensures that no value is present for Installation, not even an explicit nil
### GetProduct

`func (o *TelemetryDataDto) GetProduct() TelemetryProductDto`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *TelemetryDataDto) GetProductOk() (*TelemetryProductDto, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *TelemetryDataDto) SetProduct(v TelemetryProductDto)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *TelemetryDataDto) HasProduct() bool`

HasProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


