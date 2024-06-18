# MetricsResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **NullableInt64** | The current sum (count) for the selected metric. | [optional] 

## Methods

### NewMetricsResultDto

`func NewMetricsResultDto() *MetricsResultDto`

NewMetricsResultDto instantiates a new MetricsResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsResultDtoWithDefaults

`func NewMetricsResultDtoWithDefaults() *MetricsResultDto`

NewMetricsResultDtoWithDefaults instantiates a new MetricsResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *MetricsResultDto) GetResult() int64`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MetricsResultDto) GetResultOk() (*int64, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MetricsResultDto) SetResult(v int64)`

SetResult sets Result field to given value.

### HasResult

`func (o *MetricsResultDto) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *MetricsResultDto) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *MetricsResultDto) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


