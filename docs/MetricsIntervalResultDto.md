# MetricsIntervalResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **NullableTime** | The interval timestamp. | [optional] 
**Name** | Pointer to **NullableString** | The name of the metric. | [optional] 
**Reporter** | Pointer to **NullableString** | The reporter of the metric. &#x60;null&#x60; if the metrics are aggregated by reporter. | [optional] 
**Value** | Pointer to **NullableInt64** | The value of the metric aggregated by the interval. | [optional] 

## Methods

### NewMetricsIntervalResultDto

`func NewMetricsIntervalResultDto() *MetricsIntervalResultDto`

NewMetricsIntervalResultDto instantiates a new MetricsIntervalResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsIntervalResultDtoWithDefaults

`func NewMetricsIntervalResultDtoWithDefaults() *MetricsIntervalResultDto`

NewMetricsIntervalResultDtoWithDefaults instantiates a new MetricsIntervalResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *MetricsIntervalResultDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MetricsIntervalResultDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MetricsIntervalResultDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MetricsIntervalResultDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *MetricsIntervalResultDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *MetricsIntervalResultDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *MetricsIntervalResultDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricsIntervalResultDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricsIntervalResultDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetricsIntervalResultDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MetricsIntervalResultDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MetricsIntervalResultDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReporter

`func (o *MetricsIntervalResultDto) GetReporter() string`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *MetricsIntervalResultDto) GetReporterOk() (*string, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *MetricsIntervalResultDto) SetReporter(v string)`

SetReporter sets Reporter field to given value.

### HasReporter

`func (o *MetricsIntervalResultDto) HasReporter() bool`

HasReporter returns a boolean if a field has been set.

### SetReporterNil

`func (o *MetricsIntervalResultDto) SetReporterNil(b bool)`

 SetReporterNil sets the value for Reporter to be an explicit nil

### UnsetReporter
`func (o *MetricsIntervalResultDto) UnsetReporter()`

UnsetReporter ensures that no value is present for Reporter, not even an explicit nil
### GetValue

`func (o *MetricsIntervalResultDto) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricsIntervalResultDto) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricsIntervalResultDto) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *MetricsIntervalResultDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MetricsIntervalResultDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MetricsIntervalResultDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


