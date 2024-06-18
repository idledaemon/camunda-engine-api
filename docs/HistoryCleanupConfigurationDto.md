# HistoryCleanupConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchWindowStartTime** | Pointer to **NullableTime** | Start time of the current or next batch window. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**BatchWindowEndTime** | Pointer to **NullableTime** | End time of the current or next batch window. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the date must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., &#x60;2013-01-23T14:42:45.000+0200&#x60;. | [optional] 
**Enabled** | Pointer to **NullableBool** | Indicates whether the engine node participates in history cleanup or not. The default is &#x60;true&#x60;. Participation can be disabled via [Process Engine Configuration](https://docs.camunda.org/manual/7.21/reference/deployment-descriptors/tags/process-engine/#history-cleanup-enabled).  For more details, see [Cleanup Execution Participation per Node](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#cleanup-execution-participation-per-node). | [optional] 

## Methods

### NewHistoryCleanupConfigurationDto

`func NewHistoryCleanupConfigurationDto() *HistoryCleanupConfigurationDto`

NewHistoryCleanupConfigurationDto instantiates a new HistoryCleanupConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryCleanupConfigurationDtoWithDefaults

`func NewHistoryCleanupConfigurationDtoWithDefaults() *HistoryCleanupConfigurationDto`

NewHistoryCleanupConfigurationDtoWithDefaults instantiates a new HistoryCleanupConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchWindowStartTime

`func (o *HistoryCleanupConfigurationDto) GetBatchWindowStartTime() time.Time`

GetBatchWindowStartTime returns the BatchWindowStartTime field if non-nil, zero value otherwise.

### GetBatchWindowStartTimeOk

`func (o *HistoryCleanupConfigurationDto) GetBatchWindowStartTimeOk() (*time.Time, bool)`

GetBatchWindowStartTimeOk returns a tuple with the BatchWindowStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchWindowStartTime

`func (o *HistoryCleanupConfigurationDto) SetBatchWindowStartTime(v time.Time)`

SetBatchWindowStartTime sets BatchWindowStartTime field to given value.

### HasBatchWindowStartTime

`func (o *HistoryCleanupConfigurationDto) HasBatchWindowStartTime() bool`

HasBatchWindowStartTime returns a boolean if a field has been set.

### SetBatchWindowStartTimeNil

`func (o *HistoryCleanupConfigurationDto) SetBatchWindowStartTimeNil(b bool)`

 SetBatchWindowStartTimeNil sets the value for BatchWindowStartTime to be an explicit nil

### UnsetBatchWindowStartTime
`func (o *HistoryCleanupConfigurationDto) UnsetBatchWindowStartTime()`

UnsetBatchWindowStartTime ensures that no value is present for BatchWindowStartTime, not even an explicit nil
### GetBatchWindowEndTime

`func (o *HistoryCleanupConfigurationDto) GetBatchWindowEndTime() time.Time`

GetBatchWindowEndTime returns the BatchWindowEndTime field if non-nil, zero value otherwise.

### GetBatchWindowEndTimeOk

`func (o *HistoryCleanupConfigurationDto) GetBatchWindowEndTimeOk() (*time.Time, bool)`

GetBatchWindowEndTimeOk returns a tuple with the BatchWindowEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchWindowEndTime

`func (o *HistoryCleanupConfigurationDto) SetBatchWindowEndTime(v time.Time)`

SetBatchWindowEndTime sets BatchWindowEndTime field to given value.

### HasBatchWindowEndTime

`func (o *HistoryCleanupConfigurationDto) HasBatchWindowEndTime() bool`

HasBatchWindowEndTime returns a boolean if a field has been set.

### SetBatchWindowEndTimeNil

`func (o *HistoryCleanupConfigurationDto) SetBatchWindowEndTimeNil(b bool)`

 SetBatchWindowEndTimeNil sets the value for BatchWindowEndTime to be an explicit nil

### UnsetBatchWindowEndTime
`func (o *HistoryCleanupConfigurationDto) UnsetBatchWindowEndTime()`

UnsetBatchWindowEndTime ensures that no value is present for BatchWindowEndTime, not even an explicit nil
### GetEnabled

`func (o *HistoryCleanupConfigurationDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HistoryCleanupConfigurationDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HistoryCleanupConfigurationDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HistoryCleanupConfigurationDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *HistoryCleanupConfigurationDto) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *HistoryCleanupConfigurationDto) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


