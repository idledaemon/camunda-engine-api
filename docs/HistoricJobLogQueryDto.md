# HistoricJobLogQueryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogId** | Pointer to **NullableString** | Filter by historic job log id. | [optional] 
**JobId** | Pointer to **NullableString** | Filter by job id. | [optional] 
**JobExceptionMessage** | Pointer to **NullableString** | Filter by job exception message. | [optional] 
**JobDefinitionId** | Pointer to **NullableString** | Filter by job definition id. | [optional] 
**JobDefinitionType** | Pointer to **NullableString** | Filter by job definition type. See the [User Guide](https://docs.camunda.org/manual/7.21/user-guide/process-engine/the-job-executor/#job-creation) for more information about job definition types. | [optional] 
**JobDefinitionConfiguration** | Pointer to **NullableString** | Filter by job definition configuration. | [optional] 
**ActivityIdIn** | Pointer to **[]string** | Only include historic job logs which belong to one of the passed activity ids. | [optional] 
**FailedActivityIdIn** | Pointer to **[]string** | Only include historic job logs which belong to failures of one of the passed activity ids. | [optional] 
**ExecutionIdIn** | Pointer to **[]string** | Only include historic job logs which belong to one of the passed execution ids. | [optional] 
**ProcessInstanceId** | Pointer to **NullableString** | Filter by process instance id. | [optional] 
**ProcessDefinitionId** | Pointer to **NullableString** | Filter by process definition id. | [optional] 
**ProcessDefinitionKey** | Pointer to **NullableString** | Filter by process definition key. | [optional] 
**DeploymentId** | Pointer to **NullableString** | Filter by deployment id. | [optional] 
**TenantIdIn** | Pointer to **[]string** | Only include historic job log entries which belong to one of the passed and comma- separated tenant ids. | [optional] 
**WithoutTenantId** | Pointer to **NullableBool** | Only include historic job log entries that belong to no tenant. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Hostname** | Pointer to **NullableString** | Filter by hostname. | [optional] 
**JobPriorityLowerThanOrEquals** | Pointer to **NullableInt64** | Only include logs for which the associated job had a priority lower than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | [optional] 
**JobPriorityHigherThanOrEquals** | Pointer to **NullableInt64** | Only include logs for which the associated job had a priority higher than or equal to the given value. Value must be a valid &#x60;long&#x60; value. | [optional] 
**CreationLog** | Pointer to **NullableBool** | Only include creation logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**FailureLog** | Pointer to **NullableBool** | Only include failure logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**SuccessLog** | Pointer to **NullableBool** | Only include success logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**DeletionLog** | Pointer to **NullableBool** | Only include deletion logs. Value may only be &#x60;true&#x60;, as &#x60;false&#x60; is the default behavior. | [optional] 
**Sorting** | Pointer to [**[]HistoricJobLogQueryDtoSortingInner**](HistoricJobLogQueryDtoSortingInner.md) | An array of criteria to sort the result by. Each element of the array is                        an object that specifies one ordering. The position in the array                        identifies the rank of an ordering, i.e., whether it is primary, secondary,                        etc. Sorting has no effect for &#x60;count&#x60; endpoints | [optional] 

## Methods

### NewHistoricJobLogQueryDto

`func NewHistoricJobLogQueryDto() *HistoricJobLogQueryDto`

NewHistoricJobLogQueryDto instantiates a new HistoricJobLogQueryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricJobLogQueryDtoWithDefaults

`func NewHistoricJobLogQueryDtoWithDefaults() *HistoricJobLogQueryDto`

NewHistoricJobLogQueryDtoWithDefaults instantiates a new HistoricJobLogQueryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogId

`func (o *HistoricJobLogQueryDto) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *HistoricJobLogQueryDto) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *HistoricJobLogQueryDto) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *HistoricJobLogQueryDto) HasLogId() bool`

HasLogId returns a boolean if a field has been set.

### SetLogIdNil

`func (o *HistoricJobLogQueryDto) SetLogIdNil(b bool)`

 SetLogIdNil sets the value for LogId to be an explicit nil

### UnsetLogId
`func (o *HistoricJobLogQueryDto) UnsetLogId()`

UnsetLogId ensures that no value is present for LogId, not even an explicit nil
### GetJobId

`func (o *HistoricJobLogQueryDto) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *HistoricJobLogQueryDto) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *HistoricJobLogQueryDto) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *HistoricJobLogQueryDto) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *HistoricJobLogQueryDto) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *HistoricJobLogQueryDto) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobExceptionMessage

`func (o *HistoricJobLogQueryDto) GetJobExceptionMessage() string`

GetJobExceptionMessage returns the JobExceptionMessage field if non-nil, zero value otherwise.

### GetJobExceptionMessageOk

`func (o *HistoricJobLogQueryDto) GetJobExceptionMessageOk() (*string, bool)`

GetJobExceptionMessageOk returns a tuple with the JobExceptionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobExceptionMessage

`func (o *HistoricJobLogQueryDto) SetJobExceptionMessage(v string)`

SetJobExceptionMessage sets JobExceptionMessage field to given value.

### HasJobExceptionMessage

`func (o *HistoricJobLogQueryDto) HasJobExceptionMessage() bool`

HasJobExceptionMessage returns a boolean if a field has been set.

### SetJobExceptionMessageNil

`func (o *HistoricJobLogQueryDto) SetJobExceptionMessageNil(b bool)`

 SetJobExceptionMessageNil sets the value for JobExceptionMessage to be an explicit nil

### UnsetJobExceptionMessage
`func (o *HistoricJobLogQueryDto) UnsetJobExceptionMessage()`

UnsetJobExceptionMessage ensures that no value is present for JobExceptionMessage, not even an explicit nil
### GetJobDefinitionId

`func (o *HistoricJobLogQueryDto) GetJobDefinitionId() string`

GetJobDefinitionId returns the JobDefinitionId field if non-nil, zero value otherwise.

### GetJobDefinitionIdOk

`func (o *HistoricJobLogQueryDto) GetJobDefinitionIdOk() (*string, bool)`

GetJobDefinitionIdOk returns a tuple with the JobDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDefinitionId

`func (o *HistoricJobLogQueryDto) SetJobDefinitionId(v string)`

SetJobDefinitionId sets JobDefinitionId field to given value.

### HasJobDefinitionId

`func (o *HistoricJobLogQueryDto) HasJobDefinitionId() bool`

HasJobDefinitionId returns a boolean if a field has been set.

### SetJobDefinitionIdNil

`func (o *HistoricJobLogQueryDto) SetJobDefinitionIdNil(b bool)`

 SetJobDefinitionIdNil sets the value for JobDefinitionId to be an explicit nil

### UnsetJobDefinitionId
`func (o *HistoricJobLogQueryDto) UnsetJobDefinitionId()`

UnsetJobDefinitionId ensures that no value is present for JobDefinitionId, not even an explicit nil
### GetJobDefinitionType

`func (o *HistoricJobLogQueryDto) GetJobDefinitionType() string`

GetJobDefinitionType returns the JobDefinitionType field if non-nil, zero value otherwise.

### GetJobDefinitionTypeOk

`func (o *HistoricJobLogQueryDto) GetJobDefinitionTypeOk() (*string, bool)`

GetJobDefinitionTypeOk returns a tuple with the JobDefinitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDefinitionType

`func (o *HistoricJobLogQueryDto) SetJobDefinitionType(v string)`

SetJobDefinitionType sets JobDefinitionType field to given value.

### HasJobDefinitionType

`func (o *HistoricJobLogQueryDto) HasJobDefinitionType() bool`

HasJobDefinitionType returns a boolean if a field has been set.

### SetJobDefinitionTypeNil

`func (o *HistoricJobLogQueryDto) SetJobDefinitionTypeNil(b bool)`

 SetJobDefinitionTypeNil sets the value for JobDefinitionType to be an explicit nil

### UnsetJobDefinitionType
`func (o *HistoricJobLogQueryDto) UnsetJobDefinitionType()`

UnsetJobDefinitionType ensures that no value is present for JobDefinitionType, not even an explicit nil
### GetJobDefinitionConfiguration

`func (o *HistoricJobLogQueryDto) GetJobDefinitionConfiguration() string`

GetJobDefinitionConfiguration returns the JobDefinitionConfiguration field if non-nil, zero value otherwise.

### GetJobDefinitionConfigurationOk

`func (o *HistoricJobLogQueryDto) GetJobDefinitionConfigurationOk() (*string, bool)`

GetJobDefinitionConfigurationOk returns a tuple with the JobDefinitionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobDefinitionConfiguration

`func (o *HistoricJobLogQueryDto) SetJobDefinitionConfiguration(v string)`

SetJobDefinitionConfiguration sets JobDefinitionConfiguration field to given value.

### HasJobDefinitionConfiguration

`func (o *HistoricJobLogQueryDto) HasJobDefinitionConfiguration() bool`

HasJobDefinitionConfiguration returns a boolean if a field has been set.

### SetJobDefinitionConfigurationNil

`func (o *HistoricJobLogQueryDto) SetJobDefinitionConfigurationNil(b bool)`

 SetJobDefinitionConfigurationNil sets the value for JobDefinitionConfiguration to be an explicit nil

### UnsetJobDefinitionConfiguration
`func (o *HistoricJobLogQueryDto) UnsetJobDefinitionConfiguration()`

UnsetJobDefinitionConfiguration ensures that no value is present for JobDefinitionConfiguration, not even an explicit nil
### GetActivityIdIn

`func (o *HistoricJobLogQueryDto) GetActivityIdIn() []string`

GetActivityIdIn returns the ActivityIdIn field if non-nil, zero value otherwise.

### GetActivityIdInOk

`func (o *HistoricJobLogQueryDto) GetActivityIdInOk() (*[]string, bool)`

GetActivityIdInOk returns a tuple with the ActivityIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIdIn

`func (o *HistoricJobLogQueryDto) SetActivityIdIn(v []string)`

SetActivityIdIn sets ActivityIdIn field to given value.

### HasActivityIdIn

`func (o *HistoricJobLogQueryDto) HasActivityIdIn() bool`

HasActivityIdIn returns a boolean if a field has been set.

### SetActivityIdInNil

`func (o *HistoricJobLogQueryDto) SetActivityIdInNil(b bool)`

 SetActivityIdInNil sets the value for ActivityIdIn to be an explicit nil

### UnsetActivityIdIn
`func (o *HistoricJobLogQueryDto) UnsetActivityIdIn()`

UnsetActivityIdIn ensures that no value is present for ActivityIdIn, not even an explicit nil
### GetFailedActivityIdIn

`func (o *HistoricJobLogQueryDto) GetFailedActivityIdIn() []string`

GetFailedActivityIdIn returns the FailedActivityIdIn field if non-nil, zero value otherwise.

### GetFailedActivityIdInOk

`func (o *HistoricJobLogQueryDto) GetFailedActivityIdInOk() (*[]string, bool)`

GetFailedActivityIdInOk returns a tuple with the FailedActivityIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedActivityIdIn

`func (o *HistoricJobLogQueryDto) SetFailedActivityIdIn(v []string)`

SetFailedActivityIdIn sets FailedActivityIdIn field to given value.

### HasFailedActivityIdIn

`func (o *HistoricJobLogQueryDto) HasFailedActivityIdIn() bool`

HasFailedActivityIdIn returns a boolean if a field has been set.

### SetFailedActivityIdInNil

`func (o *HistoricJobLogQueryDto) SetFailedActivityIdInNil(b bool)`

 SetFailedActivityIdInNil sets the value for FailedActivityIdIn to be an explicit nil

### UnsetFailedActivityIdIn
`func (o *HistoricJobLogQueryDto) UnsetFailedActivityIdIn()`

UnsetFailedActivityIdIn ensures that no value is present for FailedActivityIdIn, not even an explicit nil
### GetExecutionIdIn

`func (o *HistoricJobLogQueryDto) GetExecutionIdIn() []string`

GetExecutionIdIn returns the ExecutionIdIn field if non-nil, zero value otherwise.

### GetExecutionIdInOk

`func (o *HistoricJobLogQueryDto) GetExecutionIdInOk() (*[]string, bool)`

GetExecutionIdInOk returns a tuple with the ExecutionIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionIdIn

`func (o *HistoricJobLogQueryDto) SetExecutionIdIn(v []string)`

SetExecutionIdIn sets ExecutionIdIn field to given value.

### HasExecutionIdIn

`func (o *HistoricJobLogQueryDto) HasExecutionIdIn() bool`

HasExecutionIdIn returns a boolean if a field has been set.

### SetExecutionIdInNil

`func (o *HistoricJobLogQueryDto) SetExecutionIdInNil(b bool)`

 SetExecutionIdInNil sets the value for ExecutionIdIn to be an explicit nil

### UnsetExecutionIdIn
`func (o *HistoricJobLogQueryDto) UnsetExecutionIdIn()`

UnsetExecutionIdIn ensures that no value is present for ExecutionIdIn, not even an explicit nil
### GetProcessInstanceId

`func (o *HistoricJobLogQueryDto) GetProcessInstanceId() string`

GetProcessInstanceId returns the ProcessInstanceId field if non-nil, zero value otherwise.

### GetProcessInstanceIdOk

`func (o *HistoricJobLogQueryDto) GetProcessInstanceIdOk() (*string, bool)`

GetProcessInstanceIdOk returns a tuple with the ProcessInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessInstanceId

`func (o *HistoricJobLogQueryDto) SetProcessInstanceId(v string)`

SetProcessInstanceId sets ProcessInstanceId field to given value.

### HasProcessInstanceId

`func (o *HistoricJobLogQueryDto) HasProcessInstanceId() bool`

HasProcessInstanceId returns a boolean if a field has been set.

### SetProcessInstanceIdNil

`func (o *HistoricJobLogQueryDto) SetProcessInstanceIdNil(b bool)`

 SetProcessInstanceIdNil sets the value for ProcessInstanceId to be an explicit nil

### UnsetProcessInstanceId
`func (o *HistoricJobLogQueryDto) UnsetProcessInstanceId()`

UnsetProcessInstanceId ensures that no value is present for ProcessInstanceId, not even an explicit nil
### GetProcessDefinitionId

`func (o *HistoricJobLogQueryDto) GetProcessDefinitionId() string`

GetProcessDefinitionId returns the ProcessDefinitionId field if non-nil, zero value otherwise.

### GetProcessDefinitionIdOk

`func (o *HistoricJobLogQueryDto) GetProcessDefinitionIdOk() (*string, bool)`

GetProcessDefinitionIdOk returns a tuple with the ProcessDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionId

`func (o *HistoricJobLogQueryDto) SetProcessDefinitionId(v string)`

SetProcessDefinitionId sets ProcessDefinitionId field to given value.

### HasProcessDefinitionId

`func (o *HistoricJobLogQueryDto) HasProcessDefinitionId() bool`

HasProcessDefinitionId returns a boolean if a field has been set.

### SetProcessDefinitionIdNil

`func (o *HistoricJobLogQueryDto) SetProcessDefinitionIdNil(b bool)`

 SetProcessDefinitionIdNil sets the value for ProcessDefinitionId to be an explicit nil

### UnsetProcessDefinitionId
`func (o *HistoricJobLogQueryDto) UnsetProcessDefinitionId()`

UnsetProcessDefinitionId ensures that no value is present for ProcessDefinitionId, not even an explicit nil
### GetProcessDefinitionKey

`func (o *HistoricJobLogQueryDto) GetProcessDefinitionKey() string`

GetProcessDefinitionKey returns the ProcessDefinitionKey field if non-nil, zero value otherwise.

### GetProcessDefinitionKeyOk

`func (o *HistoricJobLogQueryDto) GetProcessDefinitionKeyOk() (*string, bool)`

GetProcessDefinitionKeyOk returns a tuple with the ProcessDefinitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDefinitionKey

`func (o *HistoricJobLogQueryDto) SetProcessDefinitionKey(v string)`

SetProcessDefinitionKey sets ProcessDefinitionKey field to given value.

### HasProcessDefinitionKey

`func (o *HistoricJobLogQueryDto) HasProcessDefinitionKey() bool`

HasProcessDefinitionKey returns a boolean if a field has been set.

### SetProcessDefinitionKeyNil

`func (o *HistoricJobLogQueryDto) SetProcessDefinitionKeyNil(b bool)`

 SetProcessDefinitionKeyNil sets the value for ProcessDefinitionKey to be an explicit nil

### UnsetProcessDefinitionKey
`func (o *HistoricJobLogQueryDto) UnsetProcessDefinitionKey()`

UnsetProcessDefinitionKey ensures that no value is present for ProcessDefinitionKey, not even an explicit nil
### GetDeploymentId

`func (o *HistoricJobLogQueryDto) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *HistoricJobLogQueryDto) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *HistoricJobLogQueryDto) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *HistoricJobLogQueryDto) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *HistoricJobLogQueryDto) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *HistoricJobLogQueryDto) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetTenantIdIn

`func (o *HistoricJobLogQueryDto) GetTenantIdIn() []string`

GetTenantIdIn returns the TenantIdIn field if non-nil, zero value otherwise.

### GetTenantIdInOk

`func (o *HistoricJobLogQueryDto) GetTenantIdInOk() (*[]string, bool)`

GetTenantIdInOk returns a tuple with the TenantIdIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdIn

`func (o *HistoricJobLogQueryDto) SetTenantIdIn(v []string)`

SetTenantIdIn sets TenantIdIn field to given value.

### HasTenantIdIn

`func (o *HistoricJobLogQueryDto) HasTenantIdIn() bool`

HasTenantIdIn returns a boolean if a field has been set.

### SetTenantIdInNil

`func (o *HistoricJobLogQueryDto) SetTenantIdInNil(b bool)`

 SetTenantIdInNil sets the value for TenantIdIn to be an explicit nil

### UnsetTenantIdIn
`func (o *HistoricJobLogQueryDto) UnsetTenantIdIn()`

UnsetTenantIdIn ensures that no value is present for TenantIdIn, not even an explicit nil
### GetWithoutTenantId

`func (o *HistoricJobLogQueryDto) GetWithoutTenantId() bool`

GetWithoutTenantId returns the WithoutTenantId field if non-nil, zero value otherwise.

### GetWithoutTenantIdOk

`func (o *HistoricJobLogQueryDto) GetWithoutTenantIdOk() (*bool, bool)`

GetWithoutTenantIdOk returns a tuple with the WithoutTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutTenantId

`func (o *HistoricJobLogQueryDto) SetWithoutTenantId(v bool)`

SetWithoutTenantId sets WithoutTenantId field to given value.

### HasWithoutTenantId

`func (o *HistoricJobLogQueryDto) HasWithoutTenantId() bool`

HasWithoutTenantId returns a boolean if a field has been set.

### SetWithoutTenantIdNil

`func (o *HistoricJobLogQueryDto) SetWithoutTenantIdNil(b bool)`

 SetWithoutTenantIdNil sets the value for WithoutTenantId to be an explicit nil

### UnsetWithoutTenantId
`func (o *HistoricJobLogQueryDto) UnsetWithoutTenantId()`

UnsetWithoutTenantId ensures that no value is present for WithoutTenantId, not even an explicit nil
### GetHostname

`func (o *HistoricJobLogQueryDto) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HistoricJobLogQueryDto) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HistoricJobLogQueryDto) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *HistoricJobLogQueryDto) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *HistoricJobLogQueryDto) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *HistoricJobLogQueryDto) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetJobPriorityLowerThanOrEquals

`func (o *HistoricJobLogQueryDto) GetJobPriorityLowerThanOrEquals() int64`

GetJobPriorityLowerThanOrEquals returns the JobPriorityLowerThanOrEquals field if non-nil, zero value otherwise.

### GetJobPriorityLowerThanOrEqualsOk

`func (o *HistoricJobLogQueryDto) GetJobPriorityLowerThanOrEqualsOk() (*int64, bool)`

GetJobPriorityLowerThanOrEqualsOk returns a tuple with the JobPriorityLowerThanOrEquals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobPriorityLowerThanOrEquals

`func (o *HistoricJobLogQueryDto) SetJobPriorityLowerThanOrEquals(v int64)`

SetJobPriorityLowerThanOrEquals sets JobPriorityLowerThanOrEquals field to given value.

### HasJobPriorityLowerThanOrEquals

`func (o *HistoricJobLogQueryDto) HasJobPriorityLowerThanOrEquals() bool`

HasJobPriorityLowerThanOrEquals returns a boolean if a field has been set.

### SetJobPriorityLowerThanOrEqualsNil

`func (o *HistoricJobLogQueryDto) SetJobPriorityLowerThanOrEqualsNil(b bool)`

 SetJobPriorityLowerThanOrEqualsNil sets the value for JobPriorityLowerThanOrEquals to be an explicit nil

### UnsetJobPriorityLowerThanOrEquals
`func (o *HistoricJobLogQueryDto) UnsetJobPriorityLowerThanOrEquals()`

UnsetJobPriorityLowerThanOrEquals ensures that no value is present for JobPriorityLowerThanOrEquals, not even an explicit nil
### GetJobPriorityHigherThanOrEquals

`func (o *HistoricJobLogQueryDto) GetJobPriorityHigherThanOrEquals() int64`

GetJobPriorityHigherThanOrEquals returns the JobPriorityHigherThanOrEquals field if non-nil, zero value otherwise.

### GetJobPriorityHigherThanOrEqualsOk

`func (o *HistoricJobLogQueryDto) GetJobPriorityHigherThanOrEqualsOk() (*int64, bool)`

GetJobPriorityHigherThanOrEqualsOk returns a tuple with the JobPriorityHigherThanOrEquals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobPriorityHigherThanOrEquals

`func (o *HistoricJobLogQueryDto) SetJobPriorityHigherThanOrEquals(v int64)`

SetJobPriorityHigherThanOrEquals sets JobPriorityHigherThanOrEquals field to given value.

### HasJobPriorityHigherThanOrEquals

`func (o *HistoricJobLogQueryDto) HasJobPriorityHigherThanOrEquals() bool`

HasJobPriorityHigherThanOrEquals returns a boolean if a field has been set.

### SetJobPriorityHigherThanOrEqualsNil

`func (o *HistoricJobLogQueryDto) SetJobPriorityHigherThanOrEqualsNil(b bool)`

 SetJobPriorityHigherThanOrEqualsNil sets the value for JobPriorityHigherThanOrEquals to be an explicit nil

### UnsetJobPriorityHigherThanOrEquals
`func (o *HistoricJobLogQueryDto) UnsetJobPriorityHigherThanOrEquals()`

UnsetJobPriorityHigherThanOrEquals ensures that no value is present for JobPriorityHigherThanOrEquals, not even an explicit nil
### GetCreationLog

`func (o *HistoricJobLogQueryDto) GetCreationLog() bool`

GetCreationLog returns the CreationLog field if non-nil, zero value otherwise.

### GetCreationLogOk

`func (o *HistoricJobLogQueryDto) GetCreationLogOk() (*bool, bool)`

GetCreationLogOk returns a tuple with the CreationLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationLog

`func (o *HistoricJobLogQueryDto) SetCreationLog(v bool)`

SetCreationLog sets CreationLog field to given value.

### HasCreationLog

`func (o *HistoricJobLogQueryDto) HasCreationLog() bool`

HasCreationLog returns a boolean if a field has been set.

### SetCreationLogNil

`func (o *HistoricJobLogQueryDto) SetCreationLogNil(b bool)`

 SetCreationLogNil sets the value for CreationLog to be an explicit nil

### UnsetCreationLog
`func (o *HistoricJobLogQueryDto) UnsetCreationLog()`

UnsetCreationLog ensures that no value is present for CreationLog, not even an explicit nil
### GetFailureLog

`func (o *HistoricJobLogQueryDto) GetFailureLog() bool`

GetFailureLog returns the FailureLog field if non-nil, zero value otherwise.

### GetFailureLogOk

`func (o *HistoricJobLogQueryDto) GetFailureLogOk() (*bool, bool)`

GetFailureLogOk returns a tuple with the FailureLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLog

`func (o *HistoricJobLogQueryDto) SetFailureLog(v bool)`

SetFailureLog sets FailureLog field to given value.

### HasFailureLog

`func (o *HistoricJobLogQueryDto) HasFailureLog() bool`

HasFailureLog returns a boolean if a field has been set.

### SetFailureLogNil

`func (o *HistoricJobLogQueryDto) SetFailureLogNil(b bool)`

 SetFailureLogNil sets the value for FailureLog to be an explicit nil

### UnsetFailureLog
`func (o *HistoricJobLogQueryDto) UnsetFailureLog()`

UnsetFailureLog ensures that no value is present for FailureLog, not even an explicit nil
### GetSuccessLog

`func (o *HistoricJobLogQueryDto) GetSuccessLog() bool`

GetSuccessLog returns the SuccessLog field if non-nil, zero value otherwise.

### GetSuccessLogOk

`func (o *HistoricJobLogQueryDto) GetSuccessLogOk() (*bool, bool)`

GetSuccessLogOk returns a tuple with the SuccessLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessLog

`func (o *HistoricJobLogQueryDto) SetSuccessLog(v bool)`

SetSuccessLog sets SuccessLog field to given value.

### HasSuccessLog

`func (o *HistoricJobLogQueryDto) HasSuccessLog() bool`

HasSuccessLog returns a boolean if a field has been set.

### SetSuccessLogNil

`func (o *HistoricJobLogQueryDto) SetSuccessLogNil(b bool)`

 SetSuccessLogNil sets the value for SuccessLog to be an explicit nil

### UnsetSuccessLog
`func (o *HistoricJobLogQueryDto) UnsetSuccessLog()`

UnsetSuccessLog ensures that no value is present for SuccessLog, not even an explicit nil
### GetDeletionLog

`func (o *HistoricJobLogQueryDto) GetDeletionLog() bool`

GetDeletionLog returns the DeletionLog field if non-nil, zero value otherwise.

### GetDeletionLogOk

`func (o *HistoricJobLogQueryDto) GetDeletionLogOk() (*bool, bool)`

GetDeletionLogOk returns a tuple with the DeletionLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionLog

`func (o *HistoricJobLogQueryDto) SetDeletionLog(v bool)`

SetDeletionLog sets DeletionLog field to given value.

### HasDeletionLog

`func (o *HistoricJobLogQueryDto) HasDeletionLog() bool`

HasDeletionLog returns a boolean if a field has been set.

### SetDeletionLogNil

`func (o *HistoricJobLogQueryDto) SetDeletionLogNil(b bool)`

 SetDeletionLogNil sets the value for DeletionLog to be an explicit nil

### UnsetDeletionLog
`func (o *HistoricJobLogQueryDto) UnsetDeletionLog()`

UnsetDeletionLog ensures that no value is present for DeletionLog, not even an explicit nil
### GetSorting

`func (o *HistoricJobLogQueryDto) GetSorting() []HistoricJobLogQueryDtoSortingInner`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *HistoricJobLogQueryDto) GetSortingOk() (*[]HistoricJobLogQueryDtoSortingInner, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *HistoricJobLogQueryDto) SetSorting(v []HistoricJobLogQueryDtoSortingInner)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *HistoricJobLogQueryDto) HasSorting() bool`

HasSorting returns a boolean if a field has been set.

### SetSortingNil

`func (o *HistoricJobLogQueryDto) SetSortingNil(b bool)`

 SetSortingNil sets the value for Sorting to be an explicit nil

### UnsetSorting
`func (o *HistoricJobLogQueryDto) UnsetSorting()`

UnsetSorting ensures that no value is present for Sorting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


