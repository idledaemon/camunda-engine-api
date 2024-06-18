# DurationReportResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to **NullableInt32** | Specifies a timespan within a year. **Note:** The period must be interpreted in conjunction with the returned &#x60;periodUnit&#x60;. | [optional] 
**PeriodUnit** | Pointer to **NullableString** | The unit of the given period. Possible values are &#x60;MONTH&#x60; and &#x60;QUARTER&#x60;. | [optional] 
**Minimum** | Pointer to **NullableInt64** | The smallest duration in milliseconds of all completed process instances which were started in the given period. | [optional] 
**Maximum** | Pointer to **NullableInt64** | The greatest duration in milliseconds of all completed process instances which were started in the given period. | [optional] 
**Average** | Pointer to **NullableInt64** | The average duration in milliseconds of all completed process instances which were started in the given period. | [optional] 

## Methods

### NewDurationReportResultDto

`func NewDurationReportResultDto() *DurationReportResultDto`

NewDurationReportResultDto instantiates a new DurationReportResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDurationReportResultDtoWithDefaults

`func NewDurationReportResultDtoWithDefaults() *DurationReportResultDto`

NewDurationReportResultDtoWithDefaults instantiates a new DurationReportResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *DurationReportResultDto) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DurationReportResultDto) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DurationReportResultDto) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DurationReportResultDto) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *DurationReportResultDto) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *DurationReportResultDto) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetPeriodUnit

`func (o *DurationReportResultDto) GetPeriodUnit() string`

GetPeriodUnit returns the PeriodUnit field if non-nil, zero value otherwise.

### GetPeriodUnitOk

`func (o *DurationReportResultDto) GetPeriodUnitOk() (*string, bool)`

GetPeriodUnitOk returns a tuple with the PeriodUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodUnit

`func (o *DurationReportResultDto) SetPeriodUnit(v string)`

SetPeriodUnit sets PeriodUnit field to given value.

### HasPeriodUnit

`func (o *DurationReportResultDto) HasPeriodUnit() bool`

HasPeriodUnit returns a boolean if a field has been set.

### SetPeriodUnitNil

`func (o *DurationReportResultDto) SetPeriodUnitNil(b bool)`

 SetPeriodUnitNil sets the value for PeriodUnit to be an explicit nil

### UnsetPeriodUnit
`func (o *DurationReportResultDto) UnsetPeriodUnit()`

UnsetPeriodUnit ensures that no value is present for PeriodUnit, not even an explicit nil
### GetMinimum

`func (o *DurationReportResultDto) GetMinimum() int64`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *DurationReportResultDto) GetMinimumOk() (*int64, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *DurationReportResultDto) SetMinimum(v int64)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *DurationReportResultDto) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### SetMinimumNil

`func (o *DurationReportResultDto) SetMinimumNil(b bool)`

 SetMinimumNil sets the value for Minimum to be an explicit nil

### UnsetMinimum
`func (o *DurationReportResultDto) UnsetMinimum()`

UnsetMinimum ensures that no value is present for Minimum, not even an explicit nil
### GetMaximum

`func (o *DurationReportResultDto) GetMaximum() int64`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *DurationReportResultDto) GetMaximumOk() (*int64, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *DurationReportResultDto) SetMaximum(v int64)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *DurationReportResultDto) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### SetMaximumNil

`func (o *DurationReportResultDto) SetMaximumNil(b bool)`

 SetMaximumNil sets the value for Maximum to be an explicit nil

### UnsetMaximum
`func (o *DurationReportResultDto) UnsetMaximum()`

UnsetMaximum ensures that no value is present for Maximum, not even an explicit nil
### GetAverage

`func (o *DurationReportResultDto) GetAverage() int64`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *DurationReportResultDto) GetAverageOk() (*int64, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *DurationReportResultDto) SetAverage(v int64)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *DurationReportResultDto) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### SetAverageNil

`func (o *DurationReportResultDto) SetAverageNil(b bool)`

 SetAverageNil sets the value for Average to be an explicit nil

### UnsetAverage
`func (o *DurationReportResultDto) UnsetAverage()`

UnsetAverage ensures that no value is present for Average, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


