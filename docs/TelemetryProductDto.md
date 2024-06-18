# TelemetryProductDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the product (i.e., Camunda BPM Runtime). | [optional] 
**Version** | Pointer to **NullableString** | The version of the process engine (i.e., 7.X.Y). | [optional] 
**Edition** | Pointer to **NullableString** | The edition of the product (i.e., either community or enterprise). | [optional] 
**Internals** | Pointer to [**TelemetryInternalsDto**](TelemetryInternalsDto.md) |  | [optional] 

## Methods

### NewTelemetryProductDto

`func NewTelemetryProductDto() *TelemetryProductDto`

NewTelemetryProductDto instantiates a new TelemetryProductDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryProductDtoWithDefaults

`func NewTelemetryProductDtoWithDefaults() *TelemetryProductDto`

NewTelemetryProductDtoWithDefaults instantiates a new TelemetryProductDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TelemetryProductDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryProductDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryProductDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryProductDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TelemetryProductDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TelemetryProductDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVersion

`func (o *TelemetryProductDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TelemetryProductDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TelemetryProductDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TelemetryProductDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *TelemetryProductDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *TelemetryProductDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetEdition

`func (o *TelemetryProductDto) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *TelemetryProductDto) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *TelemetryProductDto) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *TelemetryProductDto) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### SetEditionNil

`func (o *TelemetryProductDto) SetEditionNil(b bool)`

 SetEditionNil sets the value for Edition to be an explicit nil

### UnsetEdition
`func (o *TelemetryProductDto) UnsetEdition()`

UnsetEdition ensures that no value is present for Edition, not even an explicit nil
### GetInternals

`func (o *TelemetryProductDto) GetInternals() TelemetryInternalsDto`

GetInternals returns the Internals field if non-nil, zero value otherwise.

### GetInternalsOk

`func (o *TelemetryProductDto) GetInternalsOk() (*TelemetryInternalsDto, bool)`

GetInternalsOk returns a tuple with the Internals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternals

`func (o *TelemetryProductDto) SetInternals(v TelemetryInternalsDto)`

SetInternals sets Internals field to given value.

### HasInternals

`func (o *TelemetryProductDto) HasInternals() bool`

HasInternals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


