# SetJobRetriesByProcessDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobIds** | Pointer to **[]string** | A list of job ids to set retries for. | [optional] 
**JobQuery** | Pointer to [**JobQueryDto**](JobQueryDto.md) |  | [optional] 
**DueDate** | Pointer to **NullableTime** | The due date to set for the job. A due date indicates when this job is ready for execution. Jobs with due dates in the past will be scheduled for execution. | [optional] 
**Retries** | Pointer to **NullableInt32** | The number of retries to set for the resource.  Must be &gt;&#x3D; 0. If this is 0, an incident is created and the task, or job, cannot be fetched, or acquired anymore unless the retries are increased again. Can not be null. | [optional] 
**ProcessInstances** | Pointer to **[]string** | A list of process instance ids to fetch jobs, for which retries will be set. | [optional] 
**ProcessInstanceQuery** | Pointer to [**ProcessInstanceQueryDto**](ProcessInstanceQueryDto.md) |  | [optional] 
**HistoricProcessInstanceQuery** | Pointer to [**HistoricProcessInstanceQueryDto**](HistoricProcessInstanceQueryDto.md) |  | [optional] 

## Methods

### NewSetJobRetriesByProcessDto

`func NewSetJobRetriesByProcessDto() *SetJobRetriesByProcessDto`

NewSetJobRetriesByProcessDto instantiates a new SetJobRetriesByProcessDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetJobRetriesByProcessDtoWithDefaults

`func NewSetJobRetriesByProcessDtoWithDefaults() *SetJobRetriesByProcessDto`

NewSetJobRetriesByProcessDtoWithDefaults instantiates a new SetJobRetriesByProcessDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobIds

`func (o *SetJobRetriesByProcessDto) GetJobIds() []string`

GetJobIds returns the JobIds field if non-nil, zero value otherwise.

### GetJobIdsOk

`func (o *SetJobRetriesByProcessDto) GetJobIdsOk() (*[]string, bool)`

GetJobIdsOk returns a tuple with the JobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIds

`func (o *SetJobRetriesByProcessDto) SetJobIds(v []string)`

SetJobIds sets JobIds field to given value.

### HasJobIds

`func (o *SetJobRetriesByProcessDto) HasJobIds() bool`

HasJobIds returns a boolean if a field has been set.

### SetJobIdsNil

`func (o *SetJobRetriesByProcessDto) SetJobIdsNil(b bool)`

 SetJobIdsNil sets the value for JobIds to be an explicit nil

### UnsetJobIds
`func (o *SetJobRetriesByProcessDto) UnsetJobIds()`

UnsetJobIds ensures that no value is present for JobIds, not even an explicit nil
### GetJobQuery

`func (o *SetJobRetriesByProcessDto) GetJobQuery() JobQueryDto`

GetJobQuery returns the JobQuery field if non-nil, zero value otherwise.

### GetJobQueryOk

`func (o *SetJobRetriesByProcessDto) GetJobQueryOk() (*JobQueryDto, bool)`

GetJobQueryOk returns a tuple with the JobQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobQuery

`func (o *SetJobRetriesByProcessDto) SetJobQuery(v JobQueryDto)`

SetJobQuery sets JobQuery field to given value.

### HasJobQuery

`func (o *SetJobRetriesByProcessDto) HasJobQuery() bool`

HasJobQuery returns a boolean if a field has been set.

### GetDueDate

`func (o *SetJobRetriesByProcessDto) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *SetJobRetriesByProcessDto) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *SetJobRetriesByProcessDto) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *SetJobRetriesByProcessDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *SetJobRetriesByProcessDto) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *SetJobRetriesByProcessDto) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetRetries

`func (o *SetJobRetriesByProcessDto) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *SetJobRetriesByProcessDto) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *SetJobRetriesByProcessDto) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *SetJobRetriesByProcessDto) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *SetJobRetriesByProcessDto) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *SetJobRetriesByProcessDto) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetProcessInstances

`func (o *SetJobRetriesByProcessDto) GetProcessInstances() []string`

GetProcessInstances returns the ProcessInstances field if non-nil, zero value otherwise.

### GetProcessInstancesOk

`func (o *SetJobRetriesByProcessDto) GetProcessInstancesOk() (*[]string, bool)`

GetProcessInstancesOk returns a tuple with the ProcessInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstances

`func (o *SetJobRetriesByProcessDto) SetProcessInstances(v []string)`

SetProcessInstances sets ProcessInstances field to given value.

### HasProcessInstances

`func (o *SetJobRetriesByProcessDto) HasProcessInstances() bool`

HasProcessInstances returns a boolean if a field has been set.

### SetProcessInstancesNil

`func (o *SetJobRetriesByProcessDto) SetProcessInstancesNil(b bool)`

 SetProcessInstancesNil sets the value for ProcessInstances to be an explicit nil

### UnsetProcessInstances
`func (o *SetJobRetriesByProcessDto) UnsetProcessInstances()`

UnsetProcessInstances ensures that no value is present for ProcessInstances, not even an explicit nil
### GetProcessInstanceQuery

`func (o *SetJobRetriesByProcessDto) GetProcessInstanceQuery() ProcessInstanceQueryDto`

GetProcessInstanceQuery returns the ProcessInstanceQuery field if non-nil, zero value otherwise.

### GetProcessInstanceQueryOk

`func (o *SetJobRetriesByProcessDto) GetProcessInstanceQueryOk() (*ProcessInstanceQueryDto, bool)`

GetProcessInstanceQueryOk returns a tuple with the ProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceQuery

`func (o *SetJobRetriesByProcessDto) SetProcessInstanceQuery(v ProcessInstanceQueryDto)`

SetProcessInstanceQuery sets ProcessInstanceQuery field to given value.

### HasProcessInstanceQuery

`func (o *SetJobRetriesByProcessDto) HasProcessInstanceQuery() bool`

HasProcessInstanceQuery returns a boolean if a field has been set.

### GetHistoricProcessInstanceQuery

`func (o *SetJobRetriesByProcessDto) GetHistoricProcessInstanceQuery() HistoricProcessInstanceQueryDto`

GetHistoricProcessInstanceQuery returns the HistoricProcessInstanceQuery field if non-nil, zero value otherwise.

### GetHistoricProcessInstanceQueryOk

`func (o *SetJobRetriesByProcessDto) GetHistoricProcessInstanceQueryOk() (*HistoricProcessInstanceQueryDto, bool)`

GetHistoricProcessInstanceQueryOk returns a tuple with the HistoricProcessInstanceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricProcessInstanceQuery

`func (o *SetJobRetriesByProcessDto) SetHistoricProcessInstanceQuery(v HistoricProcessInstanceQueryDto)`

SetHistoricProcessInstanceQuery sets HistoricProcessInstanceQuery field to given value.

### HasHistoricProcessInstanceQuery

`func (o *SetJobRetriesByProcessDto) HasHistoricProcessInstanceQuery() bool`

HasHistoricProcessInstanceQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


