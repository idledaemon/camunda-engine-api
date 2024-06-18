# JobDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the job. | [optional] 
**JobDefinitionId** | Pointer to **NullableString** | The id of the associated job definition. | [optional] 
**DueDate** | Pointer to **NullableTime** | The date on which this job is supposed to be processed. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the process instance which execution created the job. | [optional] 
**ExecutionId** | Pointer to **NullableString** | The specific execution id on which the job was created. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition which this job belongs to. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition which this job belongs to. | [optional] 
**Retries** | Pointer to **NullableInt32** | The number of retries this job has left. | [optional] 
**ExceptionMessage** | Pointer to **NullableString** | The message of the exception that occurred, the last time the job was executed. Is null when no exception occurred. | [optional] 
**FailedActivityId** | Pointer to **NullableString** | The id of the activity on which the last exception occurred, the last time the job was executed. Is null when no exception occurred. | [optional] 
**Suspended** | Pointer to **NullableBool** | A flag indicating whether the job is suspended or not. | [optional] 
**Priority** | Pointer to **NullableInt64** | The job&#39;s priority for execution. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant which this job belongs to. | [optional] 
**CreateTime** | Pointer to **NullableTime** | The date on which this job has been created. | [optional] 

## Methods

### NewJobDto

`func NewJobDto() *JobDto`

NewJobDto instantiates a new JobDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDtoWithDefaults

`func NewJobDtoWithDefaults() *JobDto`

NewJobDtoWithDefaults instantiates a new JobDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JobDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JobDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetJobDefinitionId

`func (o *JobDto) GetJobDefinitionId() string`

GetJobDefinitionId returns the JobDefinitionId field if non-nil, zero value otherwise.

### GetJobDefinitionIdOk

`func (o *JobDto) GetJobDefinitionIdOk() (*string, bool)`

GetJobDefinitionIdOk returns a tuple with the JobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDefinitionId

`func (o *JobDto) SetJobDefinitionId(v string)`

SetJobDefinitionId sets JobDefinitionId field to given value.

### HasJobDefinitionId

`func (o *JobDto) HasJobDefinitionId() bool`

HasJobDefinitionId returns a boolean if a field has been set.

### SetJobDefinitionIdNil

`func (o *JobDto) SetJobDefinitionIdNil(b bool)`

 SetJobDefinitionIdNil sets the value for JobDefinitionId to be an explicit nil

### UnsetJobDefinitionId
`func (o *JobDto) UnsetJobDefinitionId()`

UnsetJobDefinitionId ensures that no value is present for JobDefinitionId, not even an explicit nil
### GetDueDate

`func (o *JobDto) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *JobDto) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *JobDto) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *JobDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *JobDto) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *JobDto) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetProcessInstanceId

`func (o *JobDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *JobDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *JobDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *JobDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *JobDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *JobDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetExecutionId

`func (o *JobDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *JobDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *JobDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *JobDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *JobDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *JobDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetProcessDefinitionId

`func (o *JobDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *JobDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *JobDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *JobDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *JobDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *JobDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *JobDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *JobDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *JobDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *JobDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *JobDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *JobDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetRetries

`func (o *JobDto) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *JobDto) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *JobDto) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *JobDto) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *JobDto) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *JobDto) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetExceptionMessage

`func (o *JobDto) GetExceptionMessage() string`

GetExceptionMessage returns the ExceptionMessage field if non-nil, zero value otherwise.

### GetExceptionMessageOk

`func (o *JobDto) GetExceptionMessageOk() (*string, bool)`

GetExceptionMessageOk returns a tuple with the ExceptionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionMessage

`func (o *JobDto) SetExceptionMessage(v string)`

SetExceptionMessage sets ExceptionMessage field to given value.

### HasExceptionMessage

`func (o *JobDto) HasExceptionMessage() bool`

HasExceptionMessage returns a boolean if a field has been set.

### SetExceptionMessageNil

`func (o *JobDto) SetExceptionMessageNil(b bool)`

 SetExceptionMessageNil sets the value for ExceptionMessage to be an explicit nil

### UnsetExceptionMessage
`func (o *JobDto) UnsetExceptionMessage()`

UnsetExceptionMessage ensures that no value is present for ExceptionMessage, not even an explicit nil
### GetFailedActivityId

`func (o *JobDto) GetFailedActivityId() string`

GetFailedActivityId returns the FailedActivityId field if non-nil, zero value otherwise.

### GetFailedActivityIdOk

`func (o *JobDto) GetFailedActivityIdOk() (*string, bool)`

GetFailedActivityIdOk returns a tuple with the FailedActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedActivityId

`func (o *JobDto) SetFailedActivityId(v string)`

SetFailedActivityId sets FailedActivityId field to given value.

### HasFailedActivityId

`func (o *JobDto) HasFailedActivityId() bool`

HasFailedActivityId returns a boolean if a field has been set.

### SetFailedActivityIdNil

`func (o *JobDto) SetFailedActivityIdNil(b bool)`

 SetFailedActivityIdNil sets the value for FailedActivityId to be an explicit nil

### UnsetFailedActivityId
`func (o *JobDto) UnsetFailedActivityId()`

UnsetFailedActivityId ensures that no value is present for FailedActivityId, not even an explicit nil
### GetSuspended

`func (o *JobDto) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *JobDto) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *JobDto) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *JobDto) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *JobDto) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *JobDto) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetPriority

`func (o *JobDto) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *JobDto) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *JobDto) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *JobDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *JobDto) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *JobDto) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetTenantId

`func (o *JobDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *JobDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *JobDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *JobDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *JobDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *JobDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetCreateTime

`func (o *JobDto) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *JobDto) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *JobDto) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *JobDto) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### SetCreateTimeNil

`func (o *JobDto) SetCreateTimeNil(b bool)`

 SetCreateTimeNil sets the value for CreateTime to be an explicit nil

### UnsetCreateTime
`func (o *JobDto) UnsetCreateTime()`

UnsetCreateTime ensures that no value is present for CreateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


