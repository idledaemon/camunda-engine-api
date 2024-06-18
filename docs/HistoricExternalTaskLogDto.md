# HistoricExternalTaskLogDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the log entry. | [optional] 
**ExternalTaskId** | Pointer to **NullableString** | The id of the external task. | [optional] 
**Timestamp** | Pointer to **NullableTime** | The time when the log entry has been written. | [optional] 
**TopicName** | Pointer to **NullableString** | The topic name of the associated external task. | [optional] 
**WorkerId** | Pointer to **NullableString** | The id of the worker that posessed the most recent lock. | [optional] 
**Retries** | Pointer to **NullableInt32** | The number of retries the associated external task has left. | [optional] 
**Priority** | Pointer to **NullableInt64** | The execution priority the external task had when the log entry was created. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | The message of the error that occurred by executing the associated external task. | [optional] 
**ActivityId** | Pointer to **NullableString** | The id of the activity on which the associated external task was created. | [optional] 
**ActivityInstanceId** | Pointer to **NullableString** | The id of the activity instance on which the associated external task was created. | [optional] 
**ExecutionId** | Pointer to **NullableString** | The execution id on which the associated external task was created. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the process instance on which the associated external task was created. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition which the associated external task belongs to. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition which the associated external task belongs to. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant that this historic external task log entry belongs to. | [optional] 
**CreationLog** | Pointer to **NullableBool** | A flag indicating whether this log represents the creation of the associated external task. | [optional] 
**FailureLog** | Pointer to **NullableBool** | A flag indicating whether this log represents the failed execution of the associated external task. | [optional] 
**SuccessLog** | Pointer to **NullableBool** | A flag indicating whether this log represents the successful execution of the associated external task. | [optional] 
**DeletionLog** | Pointer to **NullableBool** | A flag indicating whether this log represents the deletion of the associated external task. | [optional] 
**RemovalTime** | Pointer to **NullableTime** | The time after which this log should be removed by the History Cleanup job. Default format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;.  For further information, please see the [documentation](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) | [optional] 
**RootProcessInstanceId** | Pointer to **NullableString** | The process instance id of the root process instance that initiated the process containing this log. | [optional] 

## Methods

### NewHistoricExternalTaskLogDto

`func NewHistoricExternalTaskLogDto() *HistoricExternalTaskLogDto`

NewHistoricExternalTaskLogDto instantiates a new HistoricExternalTaskLogDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricExternalTaskLogDtoWithDefaults

`func NewHistoricExternalTaskLogDtoWithDefaults() *HistoricExternalTaskLogDto`

NewHistoricExternalTaskLogDtoWithDefaults instantiates a new HistoricExternalTaskLogDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoricExternalTaskLogDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoricExternalTaskLogDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoricExternalTaskLogDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoricExternalTaskLogDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HistoricExternalTaskLogDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HistoricExternalTaskLogDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetExternalTaskId

`func (o *HistoricExternalTaskLogDto) GetExternalTaskId() string`

GetExternalTaskId returns the ExternalTaskId field if non-nil, zero value otherwise.

### GetExternalTaskIdOk

`func (o *HistoricExternalTaskLogDto) GetExternalTaskIdOk() (*string, bool)`

GetExternalTaskIdOk returns a tuple with the ExternalTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTaskId

`func (o *HistoricExternalTaskLogDto) SetExternalTaskId(v string)`

SetExternalTaskId sets ExternalTaskId field to given value.

### HasExternalTaskId

`func (o *HistoricExternalTaskLogDto) HasExternalTaskId() bool`

HasExternalTaskId returns a boolean if a field has been set.

### SetExternalTaskIdNil

`func (o *HistoricExternalTaskLogDto) SetExternalTaskIdNil(b bool)`

 SetExternalTaskIdNil sets the value for ExternalTaskId to be an explicit nil

### UnsetExternalTaskId
`func (o *HistoricExternalTaskLogDto) UnsetExternalTaskId()`

UnsetExternalTaskId ensures that no value is present for ExternalTaskId, not even an explicit nil
### GetTimestamp

`func (o *HistoricExternalTaskLogDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HistoricExternalTaskLogDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HistoricExternalTaskLogDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *HistoricExternalTaskLogDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *HistoricExternalTaskLogDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *HistoricExternalTaskLogDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTopicName

`func (o *HistoricExternalTaskLogDto) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *HistoricExternalTaskLogDto) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *HistoricExternalTaskLogDto) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *HistoricExternalTaskLogDto) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.

### SetTopicNameNil

`func (o *HistoricExternalTaskLogDto) SetTopicNameNil(b bool)`

 SetTopicNameNil sets the value for TopicName to be an explicit nil

### UnsetTopicName
`func (o *HistoricExternalTaskLogDto) UnsetTopicName()`

UnsetTopicName ensures that no value is present for TopicName, not even an explicit nil
### GetWorkerId

`func (o *HistoricExternalTaskLogDto) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *HistoricExternalTaskLogDto) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *HistoricExternalTaskLogDto) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *HistoricExternalTaskLogDto) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### SetWorkerIdNil

`func (o *HistoricExternalTaskLogDto) SetWorkerIdNil(b bool)`

 SetWorkerIdNil sets the value for WorkerId to be an explicit nil

### UnsetWorkerId
`func (o *HistoricExternalTaskLogDto) UnsetWorkerId()`

UnsetWorkerId ensures that no value is present for WorkerId, not even an explicit nil
### GetRetries

`func (o *HistoricExternalTaskLogDto) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *HistoricExternalTaskLogDto) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *HistoricExternalTaskLogDto) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *HistoricExternalTaskLogDto) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *HistoricExternalTaskLogDto) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *HistoricExternalTaskLogDto) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetPriority

`func (o *HistoricExternalTaskLogDto) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *HistoricExternalTaskLogDto) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *HistoricExternalTaskLogDto) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *HistoricExternalTaskLogDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *HistoricExternalTaskLogDto) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *HistoricExternalTaskLogDto) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetErrorMessage

`func (o *HistoricExternalTaskLogDto) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *HistoricExternalTaskLogDto) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *HistoricExternalTaskLogDto) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *HistoricExternalTaskLogDto) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *HistoricExternalTaskLogDto) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *HistoricExternalTaskLogDto) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetActivityId

`func (o *HistoricExternalTaskLogDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *HistoricExternalTaskLogDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *HistoricExternalTaskLogDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *HistoricExternalTaskLogDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *HistoricExternalTaskLogDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *HistoricExternalTaskLogDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetActivityInstanceId

`func (o *HistoricExternalTaskLogDto) GetActivityInstanceId() string`

GetActivityInstanceId returns the ActivityInstanceId field if non-nil, zero value otherwise.

### GetActivityInstanceIdOk

`func (o *HistoricExternalTaskLogDto) GetActivityInstanceIdOk() (*string, bool)`

GetActivityInstanceIdOk returns a tuple with the ActivityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInstanceId

`func (o *HistoricExternalTaskLogDto) SetActivityInstanceId(v string)`

SetActivityInstanceId sets ActivityInstanceId field to given value.

### HasActivityInstanceId

`func (o *HistoricExternalTaskLogDto) HasActivityInstanceId() bool`

HasActivityInstanceId returns a boolean if a field has been set.

### SetActivityInstanceIdNil

`func (o *HistoricExternalTaskLogDto) SetActivityInstanceIdNil(b bool)`

 SetActivityInstanceIdNil sets the value for ActivityInstanceId to be an explicit nil

### UnsetActivityInstanceId
`func (o *HistoricExternalTaskLogDto) UnsetActivityInstanceId()`

UnsetActivityInstanceId ensures that no value is present for ActivityInstanceId, not even an explicit nil
### GetExecutionId

`func (o *HistoricExternalTaskLogDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HistoricExternalTaskLogDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HistoricExternalTaskLogDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HistoricExternalTaskLogDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HistoricExternalTaskLogDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HistoricExternalTaskLogDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetProcessInstanceId

`func (o *HistoricExternalTaskLogDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricExternalTaskLogDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricExternalTaskLogDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricExternalTaskLogDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricExternalTaskLogDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricExternalTaskLogDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricExternalTaskLogDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricExternalTaskLogDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricExternalTaskLogDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricExternalTaskLogDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricExternalTaskLogDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricExternalTaskLogDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricExternalTaskLogDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricExternalTaskLogDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricExternalTaskLogDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricExternalTaskLogDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricExternalTaskLogDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricExternalTaskLogDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetTenantId

`func (o *HistoricExternalTaskLogDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HistoricExternalTaskLogDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HistoricExternalTaskLogDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *HistoricExternalTaskLogDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *HistoricExternalTaskLogDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *HistoricExternalTaskLogDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetCreationLog

`func (o *HistoricExternalTaskLogDto) GetCreationLog() bool`

GetCreationLog returns the CreationLog field if non-nil, zero value otherwise.

### GetCreationLogOk

`func (o *HistoricExternalTaskLogDto) GetCreationLogOk() (*bool, bool)`

GetCreationLogOk returns a tuple with the CreationLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationLog

`func (o *HistoricExternalTaskLogDto) SetCreationLog(v bool)`

SetCreationLog sets CreationLog field to given value.

### HasCreationLog

`func (o *HistoricExternalTaskLogDto) HasCreationLog() bool`

HasCreationLog returns a boolean if a field has been set.

### SetCreationLogNil

`func (o *HistoricExternalTaskLogDto) SetCreationLogNil(b bool)`

 SetCreationLogNil sets the value for CreationLog to be an explicit nil

### UnsetCreationLog
`func (o *HistoricExternalTaskLogDto) UnsetCreationLog()`

UnsetCreationLog ensures that no value is present for CreationLog, not even an explicit nil
### GetFailureLog

`func (o *HistoricExternalTaskLogDto) GetFailureLog() bool`

GetFailureLog returns the FailureLog field if non-nil, zero value otherwise.

### GetFailureLogOk

`func (o *HistoricExternalTaskLogDto) GetFailureLogOk() (*bool, bool)`

GetFailureLogOk returns a tuple with the FailureLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLog

`func (o *HistoricExternalTaskLogDto) SetFailureLog(v bool)`

SetFailureLog sets FailureLog field to given value.

### HasFailureLog

`func (o *HistoricExternalTaskLogDto) HasFailureLog() bool`

HasFailureLog returns a boolean if a field has been set.

### SetFailureLogNil

`func (o *HistoricExternalTaskLogDto) SetFailureLogNil(b bool)`

 SetFailureLogNil sets the value for FailureLog to be an explicit nil

### UnsetFailureLog
`func (o *HistoricExternalTaskLogDto) UnsetFailureLog()`

UnsetFailureLog ensures that no value is present for FailureLog, not even an explicit nil
### GetSuccessLog

`func (o *HistoricExternalTaskLogDto) GetSuccessLog() bool`

GetSuccessLog returns the SuccessLog field if non-nil, zero value otherwise.

### GetSuccessLogOk

`func (o *HistoricExternalTaskLogDto) GetSuccessLogOk() (*bool, bool)`

GetSuccessLogOk returns a tuple with the SuccessLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessLog

`func (o *HistoricExternalTaskLogDto) SetSuccessLog(v bool)`

SetSuccessLog sets SuccessLog field to given value.

### HasSuccessLog

`func (o *HistoricExternalTaskLogDto) HasSuccessLog() bool`

HasSuccessLog returns a boolean if a field has been set.

### SetSuccessLogNil

`func (o *HistoricExternalTaskLogDto) SetSuccessLogNil(b bool)`

 SetSuccessLogNil sets the value for SuccessLog to be an explicit nil

### UnsetSuccessLog
`func (o *HistoricExternalTaskLogDto) UnsetSuccessLog()`

UnsetSuccessLog ensures that no value is present for SuccessLog, not even an explicit nil
### GetDeletionLog

`func (o *HistoricExternalTaskLogDto) GetDeletionLog() bool`

GetDeletionLog returns the DeletionLog field if non-nil, zero value otherwise.

### GetDeletionLogOk

`func (o *HistoricExternalTaskLogDto) GetDeletionLogOk() (*bool, bool)`

GetDeletionLogOk returns a tuple with the DeletionLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionLog

`func (o *HistoricExternalTaskLogDto) SetDeletionLog(v bool)`

SetDeletionLog sets DeletionLog field to given value.

### HasDeletionLog

`func (o *HistoricExternalTaskLogDto) HasDeletionLog() bool`

HasDeletionLog returns a boolean if a field has been set.

### SetDeletionLogNil

`func (o *HistoricExternalTaskLogDto) SetDeletionLogNil(b bool)`

 SetDeletionLogNil sets the value for DeletionLog to be an explicit nil

### UnsetDeletionLog
`func (o *HistoricExternalTaskLogDto) UnsetDeletionLog()`

UnsetDeletionLog ensures that no value is present for DeletionLog, not even an explicit nil
### GetRemovalTime

`func (o *HistoricExternalTaskLogDto) GetRemovalTime() time.Time`

GetRemovalTime returns the RemovalTime field if non-nil, zero value otherwise.

### GetRemovalTimeOk

`func (o *HistoricExternalTaskLogDto) GetRemovalTimeOk() (*time.Time, bool)`

GetRemovalTimeOk returns a tuple with the RemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTime

`func (o *HistoricExternalTaskLogDto) SetRemovalTime(v time.Time)`

SetRemovalTime sets RemovalTime field to given value.

### HasRemovalTime

`func (o *HistoricExternalTaskLogDto) HasRemovalTime() bool`

HasRemovalTime returns a boolean if a field has been set.

### SetRemovalTimeNil

`func (o *HistoricExternalTaskLogDto) SetRemovalTimeNil(b bool)`

 SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil

### UnsetRemovalTime
`func (o *HistoricExternalTaskLogDto) UnsetRemovalTime()`

UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
### GetRootProcessInstanceId

`func (o *HistoricExternalTaskLogDto) GetRootProcessInstanceId() string`

GetRootProcessInstanceId returns the RootProcessInstanceId field if non-nil, zero value otherwise.

### GetRootProcessInstanceIdOk

`func (o *HistoricExternalTaskLogDto) GetRootProcessInstanceIdOk() (*string, bool)`

GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstanceId

`func (o *HistoricExternalTaskLogDto) SetRootProcessInstanceId(v string)`

SetRootProcessInstanceId sets RootProcessInstanceId field to given value.

### HasRootProcessInstanceId

`func (o *HistoricExternalTaskLogDto) HasRootProcessInstanceId() bool`

HasRootProcessInstanceId returns a boolean if a field has been set.

### SetRootProcessInstanceIdNil

`func (o *HistoricExternalTaskLogDto) SetRootProcessInstanceIdNil(b bool)`

 SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil

### UnsetRootProcessInstanceId
`func (o *HistoricExternalTaskLogDto) UnsetRootProcessInstanceId()`

UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


