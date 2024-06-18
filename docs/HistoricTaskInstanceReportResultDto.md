# HistoricTaskInstanceReportResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskName** | Pointer to **NullableString** | The name of the task. It is only available when the &#x60;groupBy&#x60; parameter is set to &#x60;taskName&#x60;. Else the value is &#x60;null&#x60;.  **Note:** This property is only set for a historic task report object. In these cases, the value of the &#x60;reportType&#x60; query parameter is &#x60;count&#x60;. | [optional] 
**Count** | Pointer to **NullableInt64** | The number of tasks which have the given definition.  **Note:** This property is only set for a historic task report object. In these cases, the value of the &#x60;reportType&#x60; query parameter is &#x60;count&#x60;. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition.  **Note:** This property is only set for a historic task report object. In these cases, the value of the &#x60;reportType&#x60; query parameter is &#x60;count&#x60;. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition.  **Note:** This property is only set for a historic task report object. In these cases, the value of the &#x60;reportType&#x60; query parameter is &#x60;count&#x60;. | [optional] 
**ProcessDefinitionName** | Pointer to **NullableString** | The name of the process definition.  **Note:** This property is only set for a historic task report object. In these cases, the value of the &#x60;reportType&#x60; query parameter is &#x60;count&#x60;. | [optional] 
**Period** | Pointer to **NullableInt32** | Specifies a span of time within a year. **Note:** The period must be interpreted in conjunction with the returned &#x60;periodUnit&#x60;.  **Note:** This property is only set for a duration report object. In these cases, the value of the &#x60;reportType&#x60; query parameter is &#x60;duration&#x60;. | [optional] 
**PeriodUnit** | Pointer to **NullableString** | The unit of the given period. Possible values are &#x60;MONTH&#x60; and &#x60;QUARTER&#x60;.  **Note:** This property is only set for a duration report object. In these cases, the value of the &#x60;reportType&#x60; query parameter is &#x60;duration&#x60;. | [optional] 
**Minimum** | Pointer to **NullableInt64** | The smallest duration in milliseconds of all completed process instances which were started in the given period.  **Note:** This property is only set for a duration report object. In these cases, the value of the &#x60;reportType&#x60; query parameter is &#x60;duration&#x60;. | [optional] 
**Maximum** | Pointer to **NullableInt64** | The greatest duration in milliseconds of all completed process instances which were started in the given period.  **Note:** This property is only set for a duration report object. In these cases, the value of the &#x60;reportType&#x60; query parameter is &#x60;duration&#x60;. | [optional] 
**Average** | Pointer to **NullableInt64** | The average duration in milliseconds of all completed process instances which were started in the given period.  **Note:** This property is only set for a duration report object. In these cases, the value of the &#x60;reportType&#x60; query parameter is &#x60;duration&#x60;. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant. | [optional] 

## Methods

### NewHistoricTaskInstanceReportResultDto

`func NewHistoricTaskInstanceReportResultDto() *HistoricTaskInstanceReportResultDto`

NewHistoricTaskInstanceReportResultDto instantiates a new HistoricTaskInstanceReportResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricTaskInstanceReportResultDtoWithDefaults

`func NewHistoricTaskInstanceReportResultDtoWithDefaults() *HistoricTaskInstanceReportResultDto`

NewHistoricTaskInstanceReportResultDtoWithDefaults instantiates a new HistoricTaskInstanceReportResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskName

`func (o *HistoricTaskInstanceReportResultDto) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *HistoricTaskInstanceReportResultDto) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *HistoricTaskInstanceReportResultDto) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *HistoricTaskInstanceReportResultDto) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### SetTaskNameNil

`func (o *HistoricTaskInstanceReportResultDto) SetTaskNameNil(b bool)`

 SetTaskNameNil sets the value for TaskName to be an explicit nil

### UnsetTaskName
`func (o *HistoricTaskInstanceReportResultDto) UnsetTaskName()`

UnsetTaskName ensures that no value is present for TaskName, not even an explicit nil
### GetCount

`func (o *HistoricTaskInstanceReportResultDto) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HistoricTaskInstanceReportResultDto) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HistoricTaskInstanceReportResultDto) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *HistoricTaskInstanceReportResultDto) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *HistoricTaskInstanceReportResultDto) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *HistoricTaskInstanceReportResultDto) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricTaskInstanceReportResultDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricTaskInstanceReportResultDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricTaskInstanceReportResultDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricTaskInstanceReportResultDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricTaskInstanceReportResultDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricTaskInstanceReportResultDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricTaskInstanceReportResultDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricTaskInstanceReportResultDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricTaskInstanceReportResultDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricTaskInstanceReportResultDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricTaskInstanceReportResultDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricTaskInstanceReportResultDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionName

`func (o *HistoricTaskInstanceReportResultDto) GetProcessDefinitionName() string`

GetProcessDefinitionName returns the ProcessDefinitionName field if non-nil, zero value otherwise.

### GetProcessDefinitionNameOk

`func (o *HistoricTaskInstanceReportResultDto) GetProcessDefinitionNameOk() (*string, bool)`

GetProcessDefinitionNameOk returns a tuple with the ProcessDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionName

`func (o *HistoricTaskInstanceReportResultDto) SetProcessDefinitionName(v string)`

SetProcessDefinitionName sets ProcessDefinitionName field to given value.

### HasProcessDefinitionName

`func (o *HistoricTaskInstanceReportResultDto) HasProcessDefinitionName() bool`

HasProcessDefinitionName returns a boolean if a field has been set.

### SetProcessDefinitionNameNil

`func (o *HistoricTaskInstanceReportResultDto) SetProcessDefinitionNameNil(b bool)`

 SetProcessDefinitionNameNil sets the value for ProcessDefinitionName to be an explicit nil

### UnsetProcessDefinitionName
`func (o *HistoricTaskInstanceReportResultDto) UnsetProcessDefinitionName()`

UnsetProcessDefinitionName ensures that no value is present for ProcessDefinitionName, not even an explicit nil
### GetPeriod

`func (o *HistoricTaskInstanceReportResultDto) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *HistoricTaskInstanceReportResultDto) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *HistoricTaskInstanceReportResultDto) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *HistoricTaskInstanceReportResultDto) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *HistoricTaskInstanceReportResultDto) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *HistoricTaskInstanceReportResultDto) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetPeriodUnit

`func (o *HistoricTaskInstanceReportResultDto) GetPeriodUnit() string`

GetPeriodUnit returns the PeriodUnit field if non-nil, zero value otherwise.

### GetPeriodUnitOk

`func (o *HistoricTaskInstanceReportResultDto) GetPeriodUnitOk() (*string, bool)`

GetPeriodUnitOk returns a tuple with the PeriodUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodUnit

`func (o *HistoricTaskInstanceReportResultDto) SetPeriodUnit(v string)`

SetPeriodUnit sets PeriodUnit field to given value.

### HasPeriodUnit

`func (o *HistoricTaskInstanceReportResultDto) HasPeriodUnit() bool`

HasPeriodUnit returns a boolean if a field has been set.

### SetPeriodUnitNil

`func (o *HistoricTaskInstanceReportResultDto) SetPeriodUnitNil(b bool)`

 SetPeriodUnitNil sets the value for PeriodUnit to be an explicit nil

### UnsetPeriodUnit
`func (o *HistoricTaskInstanceReportResultDto) UnsetPeriodUnit()`

UnsetPeriodUnit ensures that no value is present for PeriodUnit, not even an explicit nil
### GetMinimum

`func (o *HistoricTaskInstanceReportResultDto) GetMinimum() int64`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *HistoricTaskInstanceReportResultDto) GetMinimumOk() (*int64, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *HistoricTaskInstanceReportResultDto) SetMinimum(v int64)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *HistoricTaskInstanceReportResultDto) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### SetMinimumNil

`func (o *HistoricTaskInstanceReportResultDto) SetMinimumNil(b bool)`

 SetMinimumNil sets the value for Minimum to be an explicit nil

### UnsetMinimum
`func (o *HistoricTaskInstanceReportResultDto) UnsetMinimum()`

UnsetMinimum ensures that no value is present for Minimum, not even an explicit nil
### GetMaximum

`func (o *HistoricTaskInstanceReportResultDto) GetMaximum() int64`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *HistoricTaskInstanceReportResultDto) GetMaximumOk() (*int64, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *HistoricTaskInstanceReportResultDto) SetMaximum(v int64)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *HistoricTaskInstanceReportResultDto) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### SetMaximumNil

`func (o *HistoricTaskInstanceReportResultDto) SetMaximumNil(b bool)`

 SetMaximumNil sets the value for Maximum to be an explicit nil

### UnsetMaximum
`func (o *HistoricTaskInstanceReportResultDto) UnsetMaximum()`

UnsetMaximum ensures that no value is present for Maximum, not even an explicit nil
### GetAverage

`func (o *HistoricTaskInstanceReportResultDto) GetAverage() int64`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *HistoricTaskInstanceReportResultDto) GetAverageOk() (*int64, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *HistoricTaskInstanceReportResultDto) SetAverage(v int64)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *HistoricTaskInstanceReportResultDto) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### SetAverageNil

`func (o *HistoricTaskInstanceReportResultDto) SetAverageNil(b bool)`

 SetAverageNil sets the value for Average to be an explicit nil

### UnsetAverage
`func (o *HistoricTaskInstanceReportResultDto) UnsetAverage()`

UnsetAverage ensures that no value is present for Average, not even an explicit nil
### GetTenantId

`func (o *HistoricTaskInstanceReportResultDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HistoricTaskInstanceReportResultDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HistoricTaskInstanceReportResultDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *HistoricTaskInstanceReportResultDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *HistoricTaskInstanceReportResultDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *HistoricTaskInstanceReportResultDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


