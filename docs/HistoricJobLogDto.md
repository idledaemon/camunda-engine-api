# HistoricJobLogDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the log entry. | [optional] 
**Timestamp** | Pointer to **NullableTime** | The time when the log entry has been written. | [optional] 
**RemovalTime** | Pointer to **NullableTime** | The time after which the log entry should be removed by the History Cleanup job. Default format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;. For further info see the [docs](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/) | [optional] 
**JobId** | Pointer to **NullableString** | The id of the associated job. | [optional] 
**JobDueDate** | Pointer to **NullableTime** | The date on which the associated job is supposed to be processed. | [optional] 
**JobRetries** | Pointer to **NullableInt32** | The number of retries the associated job has left. | [optional] 
**JobPriority** | Pointer to **NullableInt64** | The execution priority the job had when the log entry was created. | [optional] 
**JobExceptionMessage** | Pointer to **NullableString** | The message of the exception that occurred by executing the associated job. | [optional] 
**FailedActivityId** | Pointer to **NullableString** | The id of the activity on which the last exception occurred by executing the associated job. | [optional] 
**JobDefinitionId** | Pointer to **NullableString** | The id of the job definition on which the associated job was created. | [optional] 
**JobDefinitionType** | Pointer to **NullableString** | The job definition type of the associated job. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/the-job-executor/#job-creation) for more information about job definition types. | [optional] 
**JobDefinitionConfiguration** | Pointer to **NullableString** | The job definition configuration type of the associated job. | [optional] 
**ActivityId** | Pointer to **NullableString** | The id of the activity on which the associated job was created. | [optional] 
**ExecutionId** | Pointer to **NullableString** | The execution id on which the associated job was created. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | The id of the process instance on which the associated job was created. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | The id of the process definition which the associated job belongs to. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | The key of the process definition which the associated job belongs to. | [optional] 
**DeploymentId** | Pointer to **NullableString** | The id of the deployment which the associated job belongs to. | [optional] 
**RootProcessInstanceId** | Pointer to **NullableString** | The process instance id of the root process instance that initiated the process which the associated job belongs to. | [optional] 
**TenantId** | Pointer to **NullableString** | The id of the tenant that this historic job log entry belongs to. | [optional] 
**Hostname** | Pointer to **NullableString** |  The name of the host of the Process Engine where the job of this historic job log entry was executed. | [optional] 
**CreationLog** | Pointer to **NullableBool** | A flag indicating whether this log represents the creation of the associated job. | [optional] 
**FailureLog** | Pointer to **NullableBool** | A flag indicating whether this log represents the failed execution of the associated job. | [optional] 
**SuccessLog** | Pointer to **NullableBool** | A flag indicating whether this log represents the successful execution of the associated job. | [optional] 
**DeletionLog** | Pointer to **NullableBool** | A flag indicating whether this log represents the deletion of the associated job. | [optional] 

## Methods

### NewHistoricJobLogDto

`func NewHistoricJobLogDto() *HistoricJobLogDto`

NewHistoricJobLogDto instantiates a new HistoricJobLogDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricJobLogDtoWithDefaults

`func NewHistoricJobLogDtoWithDefaults() *HistoricJobLogDto`

NewHistoricJobLogDtoWithDefaults instantiates a new HistoricJobLogDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoricJobLogDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoricJobLogDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoricJobLogDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoricJobLogDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HistoricJobLogDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HistoricJobLogDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *HistoricJobLogDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HistoricJobLogDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HistoricJobLogDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *HistoricJobLogDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *HistoricJobLogDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *HistoricJobLogDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetRemovalTime

`func (o *HistoricJobLogDto) GetRemovalTime() time.Time`

GetRemovalTime returns the RemovalTime field if non-nil, zero value otherwise.

### GetRemovalTimeOk

`func (o *HistoricJobLogDto) GetRemovalTimeOk() (*time.Time, bool)`

GetRemovalTimeOk returns a tuple with the RemovalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalTime

`func (o *HistoricJobLogDto) SetRemovalTime(v time.Time)`

SetRemovalTime sets RemovalTime field to given value.

### HasRemovalTime

`func (o *HistoricJobLogDto) HasRemovalTime() bool`

HasRemovalTime returns a boolean if a field has been set.

### SetRemovalTimeNil

`func (o *HistoricJobLogDto) SetRemovalTimeNil(b bool)`

 SetRemovalTimeNil sets the value for RemovalTime to be an explicit nil

### UnsetRemovalTime
`func (o *HistoricJobLogDto) UnsetRemovalTime()`

UnsetRemovalTime ensures that no value is present for RemovalTime, not even an explicit nil
### GetJobId

`func (o *HistoricJobLogDto) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *HistoricJobLogDto) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *HistoricJobLogDto) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *HistoricJobLogDto) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *HistoricJobLogDto) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *HistoricJobLogDto) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobDueDate

`func (o *HistoricJobLogDto) GetJobDueDate() time.Time`

GetJobDueDate returns the JobDueDate field if non-nil, zero value otherwise.

### GetJobDueDateOk

`func (o *HistoricJobLogDto) GetJobDueDateOk() (*time.Time, bool)`

GetJobDueDateOk returns a tuple with the JobDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDueDate

`func (o *HistoricJobLogDto) SetJobDueDate(v time.Time)`

SetJobDueDate sets JobDueDate field to given value.

### HasJobDueDate

`func (o *HistoricJobLogDto) HasJobDueDate() bool`

HasJobDueDate returns a boolean if a field has been set.

### SetJobDueDateNil

`func (o *HistoricJobLogDto) SetJobDueDateNil(b bool)`

 SetJobDueDateNil sets the value for JobDueDate to be an explicit nil

### UnsetJobDueDate
`func (o *HistoricJobLogDto) UnsetJobDueDate()`

UnsetJobDueDate ensures that no value is present for JobDueDate, not even an explicit nil
### GetJobRetries

`func (o *HistoricJobLogDto) GetJobRetries() int32`

GetJobRetries returns the JobRetries field if non-nil, zero value otherwise.

### GetJobRetriesOk

`func (o *HistoricJobLogDto) GetJobRetriesOk() (*int32, bool)`

GetJobRetriesOk returns a tuple with the JobRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRetries

`func (o *HistoricJobLogDto) SetJobRetries(v int32)`

SetJobRetries sets JobRetries field to given value.

### HasJobRetries

`func (o *HistoricJobLogDto) HasJobRetries() bool`

HasJobRetries returns a boolean if a field has been set.

### SetJobRetriesNil

`func (o *HistoricJobLogDto) SetJobRetriesNil(b bool)`

 SetJobRetriesNil sets the value for JobRetries to be an explicit nil

### UnsetJobRetries
`func (o *HistoricJobLogDto) UnsetJobRetries()`

UnsetJobRetries ensures that no value is present for JobRetries, not even an explicit nil
### GetJobPriority

`func (o *HistoricJobLogDto) GetJobPriority() int64`

GetJobPriority returns the JobPriority field if non-nil, zero value otherwise.

### GetJobPriorityOk

`func (o *HistoricJobLogDto) GetJobPriorityOk() (*int64, bool)`

GetJobPriorityOk returns a tuple with the JobPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobPriority

`func (o *HistoricJobLogDto) SetJobPriority(v int64)`

SetJobPriority sets JobPriority field to given value.

### HasJobPriority

`func (o *HistoricJobLogDto) HasJobPriority() bool`

HasJobPriority returns a boolean if a field has been set.

### SetJobPriorityNil

`func (o *HistoricJobLogDto) SetJobPriorityNil(b bool)`

 SetJobPriorityNil sets the value for JobPriority to be an explicit nil

### UnsetJobPriority
`func (o *HistoricJobLogDto) UnsetJobPriority()`

UnsetJobPriority ensures that no value is present for JobPriority, not even an explicit nil
### GetJobExceptionMessage

`func (o *HistoricJobLogDto) GetJobExceptionMessage() string`

GetJobExceptionMessage returns the JobExceptionMessage field if non-nil, zero value otherwise.

### GetJobExceptionMessageOk

`func (o *HistoricJobLogDto) GetJobExceptionMessageOk() (*string, bool)`

GetJobExceptionMessageOk returns a tuple with the JobExceptionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobExceptionMessage

`func (o *HistoricJobLogDto) SetJobExceptionMessage(v string)`

SetJobExceptionMessage sets JobExceptionMessage field to given value.

### HasJobExceptionMessage

`func (o *HistoricJobLogDto) HasJobExceptionMessage() bool`

HasJobExceptionMessage returns a boolean if a field has been set.

### SetJobExceptionMessageNil

`func (o *HistoricJobLogDto) SetJobExceptionMessageNil(b bool)`

 SetJobExceptionMessageNil sets the value for JobExceptionMessage to be an explicit nil

### UnsetJobExceptionMessage
`func (o *HistoricJobLogDto) UnsetJobExceptionMessage()`

UnsetJobExceptionMessage ensures that no value is present for JobExceptionMessage, not even an explicit nil
### GetFailedActivityId

`func (o *HistoricJobLogDto) GetFailedActivityId() string`

GetFailedActivityId returns the FailedActivityId field if non-nil, zero value otherwise.

### GetFailedActivityIdOk

`func (o *HistoricJobLogDto) GetFailedActivityIdOk() (*string, bool)`

GetFailedActivityIdOk returns a tuple with the FailedActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedActivityId

`func (o *HistoricJobLogDto) SetFailedActivityId(v string)`

SetFailedActivityId sets FailedActivityId field to given value.

### HasFailedActivityId

`func (o *HistoricJobLogDto) HasFailedActivityId() bool`

HasFailedActivityId returns a boolean if a field has been set.

### SetFailedActivityIdNil

`func (o *HistoricJobLogDto) SetFailedActivityIdNil(b bool)`

 SetFailedActivityIdNil sets the value for FailedActivityId to be an explicit nil

### UnsetFailedActivityId
`func (o *HistoricJobLogDto) UnsetFailedActivityId()`

UnsetFailedActivityId ensures that no value is present for FailedActivityId, not even an explicit nil
### GetJobDefinitionId

`func (o *HistoricJobLogDto) GetJobDefinitionId() string`

GetJobDefinitionId returns the JobDefinitionId field if non-nil, zero value otherwise.

### GetJobDefinitionIdOk

`func (o *HistoricJobLogDto) GetJobDefinitionIdOk() (*string, bool)`

GetJobDefinitionIdOk returns a tuple with the JobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDefinitionId

`func (o *HistoricJobLogDto) SetJobDefinitionId(v string)`

SetJobDefinitionId sets JobDefinitionId field to given value.

### HasJobDefinitionId

`func (o *HistoricJobLogDto) HasJobDefinitionId() bool`

HasJobDefinitionId returns a boolean if a field has been set.

### SetJobDefinitionIdNil

`func (o *HistoricJobLogDto) SetJobDefinitionIdNil(b bool)`

 SetJobDefinitionIdNil sets the value for JobDefinitionId to be an explicit nil

### UnsetJobDefinitionId
`func (o *HistoricJobLogDto) UnsetJobDefinitionId()`

UnsetJobDefinitionId ensures that no value is present for JobDefinitionId, not even an explicit nil
### GetJobDefinitionType

`func (o *HistoricJobLogDto) GetJobDefinitionType() string`

GetJobDefinitionType returns the JobDefinitionType field if non-nil, zero value otherwise.

### GetJobDefinitionTypeOk

`func (o *HistoricJobLogDto) GetJobDefinitionTypeOk() (*string, bool)`

GetJobDefinitionTypeOk returns a tuple with the JobDefinitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDefinitionType

`func (o *HistoricJobLogDto) SetJobDefinitionType(v string)`

SetJobDefinitionType sets JobDefinitionType field to given value.

### HasJobDefinitionType

`func (o *HistoricJobLogDto) HasJobDefinitionType() bool`

HasJobDefinitionType returns a boolean if a field has been set.

### SetJobDefinitionTypeNil

`func (o *HistoricJobLogDto) SetJobDefinitionTypeNil(b bool)`

 SetJobDefinitionTypeNil sets the value for JobDefinitionType to be an explicit nil

### UnsetJobDefinitionType
`func (o *HistoricJobLogDto) UnsetJobDefinitionType()`

UnsetJobDefinitionType ensures that no value is present for JobDefinitionType, not even an explicit nil
### GetJobDefinitionConfiguration

`func (o *HistoricJobLogDto) GetJobDefinitionConfiguration() string`

GetJobDefinitionConfiguration returns the JobDefinitionConfiguration field if non-nil, zero value otherwise.

### GetJobDefinitionConfigurationOk

`func (o *HistoricJobLogDto) GetJobDefinitionConfigurationOk() (*string, bool)`

GetJobDefinitionConfigurationOk returns a tuple with the JobDefinitionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDefinitionConfiguration

`func (o *HistoricJobLogDto) SetJobDefinitionConfiguration(v string)`

SetJobDefinitionConfiguration sets JobDefinitionConfiguration field to given value.

### HasJobDefinitionConfiguration

`func (o *HistoricJobLogDto) HasJobDefinitionConfiguration() bool`

HasJobDefinitionConfiguration returns a boolean if a field has been set.

### SetJobDefinitionConfigurationNil

`func (o *HistoricJobLogDto) SetJobDefinitionConfigurationNil(b bool)`

 SetJobDefinitionConfigurationNil sets the value for JobDefinitionConfiguration to be an explicit nil

### UnsetJobDefinitionConfiguration
`func (o *HistoricJobLogDto) UnsetJobDefinitionConfiguration()`

UnsetJobDefinitionConfiguration ensures that no value is present for JobDefinitionConfiguration, not even an explicit nil
### GetActivityId

`func (o *HistoricJobLogDto) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *HistoricJobLogDto) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *HistoricJobLogDto) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *HistoricJobLogDto) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### SetActivityIdNil

`func (o *HistoricJobLogDto) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *HistoricJobLogDto) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetExecutionId

`func (o *HistoricJobLogDto) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HistoricJobLogDto) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HistoricJobLogDto) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HistoricJobLogDto) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HistoricJobLogDto) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HistoricJobLogDto) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetProcessInstanceId

`func (o *HistoricJobLogDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricJobLogDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricJobLogDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricJobLogDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricJobLogDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricJobLogDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricJobLogDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricJobLogDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricJobLogDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricJobLogDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricJobLogDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricJobLogDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricJobLogDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricJobLogDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricJobLogDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricJobLogDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricJobLogDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricJobLogDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetDeploymentId

`func (o *HistoricJobLogDto) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *HistoricJobLogDto) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *HistoricJobLogDto) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *HistoricJobLogDto) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *HistoricJobLogDto) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *HistoricJobLogDto) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetRootProcessInstanceId

`func (o *HistoricJobLogDto) GetRootProcessInstanceId() string`

GetRootProcessInstanceId returns the RootProcessInstanceId field if non-nil, zero value otherwise.

### GetRootProcessInstanceIdOk

`func (o *HistoricJobLogDto) GetRootProcessInstanceIdOk() (*string, bool)`

GetRootProcessInstanceIdOk returns a tuple with the RootProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProcessInstanceId

`func (o *HistoricJobLogDto) SetRootProcessInstanceId(v string)`

SetRootProcessInstanceId sets RootProcessInstanceId field to given value.

### HasRootProcessInstanceId

`func (o *HistoricJobLogDto) HasRootProcessInstanceId() bool`

HasRootProcessInstanceId returns a boolean if a field has been set.

### SetRootProcessInstanceIdNil

`func (o *HistoricJobLogDto) SetRootProcessInstanceIdNil(b bool)`

 SetRootProcessInstanceIdNil sets the value for RootProcessInstanceId to be an explicit nil

### UnsetRootProcessInstanceId
`func (o *HistoricJobLogDto) UnsetRootProcessInstanceId()`

UnsetRootProcessInstanceId ensures that no value is present for RootProcessInstanceId, not even an explicit nil
### GetTenantId

`func (o *HistoricJobLogDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HistoricJobLogDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HistoricJobLogDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *HistoricJobLogDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *HistoricJobLogDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *HistoricJobLogDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetHostname

`func (o *HistoricJobLogDto) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HistoricJobLogDto) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HistoricJobLogDto) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *HistoricJobLogDto) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *HistoricJobLogDto) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *HistoricJobLogDto) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetCreationLog

`func (o *HistoricJobLogDto) GetCreationLog() bool`

GetCreationLog returns the CreationLog field if non-nil, zero value otherwise.

### GetCreationLogOk

`func (o *HistoricJobLogDto) GetCreationLogOk() (*bool, bool)`

GetCreationLogOk returns a tuple with the CreationLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationLog

`func (o *HistoricJobLogDto) SetCreationLog(v bool)`

SetCreationLog sets CreationLog field to given value.

### HasCreationLog

`func (o *HistoricJobLogDto) HasCreationLog() bool`

HasCreationLog returns a boolean if a field has been set.

### SetCreationLogNil

`func (o *HistoricJobLogDto) SetCreationLogNil(b bool)`

 SetCreationLogNil sets the value for CreationLog to be an explicit nil

### UnsetCreationLog
`func (o *HistoricJobLogDto) UnsetCreationLog()`

UnsetCreationLog ensures that no value is present for CreationLog, not even an explicit nil
### GetFailureLog

`func (o *HistoricJobLogDto) GetFailureLog() bool`

GetFailureLog returns the FailureLog field if non-nil, zero value otherwise.

### GetFailureLogOk

`func (o *HistoricJobLogDto) GetFailureLogOk() (*bool, bool)`

GetFailureLogOk returns a tuple with the FailureLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLog

`func (o *HistoricJobLogDto) SetFailureLog(v bool)`

SetFailureLog sets FailureLog field to given value.

### HasFailureLog

`func (o *HistoricJobLogDto) HasFailureLog() bool`

HasFailureLog returns a boolean if a field has been set.

### SetFailureLogNil

`func (o *HistoricJobLogDto) SetFailureLogNil(b bool)`

 SetFailureLogNil sets the value for FailureLog to be an explicit nil

### UnsetFailureLog
`func (o *HistoricJobLogDto) UnsetFailureLog()`

UnsetFailureLog ensures that no value is present for FailureLog, not even an explicit nil
### GetSuccessLog

`func (o *HistoricJobLogDto) GetSuccessLog() bool`

GetSuccessLog returns the SuccessLog field if non-nil, zero value otherwise.

### GetSuccessLogOk

`func (o *HistoricJobLogDto) GetSuccessLogOk() (*bool, bool)`

GetSuccessLogOk returns a tuple with the SuccessLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessLog

`func (o *HistoricJobLogDto) SetSuccessLog(v bool)`

SetSuccessLog sets SuccessLog field to given value.

### HasSuccessLog

`func (o *HistoricJobLogDto) HasSuccessLog() bool`

HasSuccessLog returns a boolean if a field has been set.

### SetSuccessLogNil

`func (o *HistoricJobLogDto) SetSuccessLogNil(b bool)`

 SetSuccessLogNil sets the value for SuccessLog to be an explicit nil

### UnsetSuccessLog
`func (o *HistoricJobLogDto) UnsetSuccessLog()`

UnsetSuccessLog ensures that no value is present for SuccessLog, not even an explicit nil
### GetDeletionLog

`func (o *HistoricJobLogDto) GetDeletionLog() bool`

GetDeletionLog returns the DeletionLog field if non-nil, zero value otherwise.

### GetDeletionLogOk

`func (o *HistoricJobLogDto) GetDeletionLogOk() (*bool, bool)`

GetDeletionLogOk returns a tuple with the DeletionLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionLog

`func (o *HistoricJobLogDto) SetDeletionLog(v bool)`

SetDeletionLog sets DeletionLog field to given value.

### HasDeletionLog

`func (o *HistoricJobLogDto) HasDeletionLog() bool`

HasDeletionLog returns a boolean if a field has been set.

### SetDeletionLogNil

`func (o *HistoricJobLogDto) SetDeletionLogNil(b bool)`

 SetDeletionLogNil sets the value for DeletionLog to be an explicit nil

### UnsetDeletionLog
`func (o *HistoricJobLogDto) UnsetDeletionLog()`

UnsetDeletionLog ensures that no value is present for DeletionLog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


